package middleware

import (
	"bff/config"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/log"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"
)

type Middleware struct {
	Repo   repository.MicroserviceRepositoryInterface
	Config *config.Config
}

func NewMiddleware(repo repository.MicroserviceRepositoryInterface, config *config.Config) *Middleware {
	return &Middleware{
		Repo:   repo,
		Config: config,
	}
}

func (m *Middleware) RequestContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), config.Requestkey, r)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) AddResponseWriterToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), config.HTTPResponseWriterKey, w)
		// Retrieve the Authorization header value from the request
		authHeader := r.Header.Get("Authorization")
		// Add the bearer token as a header in the context
		headers := map[string]string{
			"Authorization": authHeader,
		}
		ctx = context.WithValue(ctx, config.HTTPHeadersKey, headers)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) ErrorHandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a buffer to capture the response
		buf := &bytes.Buffer{}
		responseWriter := httptest.NewRecorder()
		// Replace the original response writer with the recorder
		defer func() {
			// Check for errors in the response
			if responseWriter.Code >= http.StatusBadRequest {
				// Handle the error using logrus
				log.Logger.WithFields(logrus.Fields{
					"status_code": responseWriter.Code,
					"response":    buf.String(),
				}).Error("HTTP error detected")
			}
			// Copy the response from the recorder to the original writer
			for key, values := range responseWriter.Header() {
				w.Header()[key] = values
			}
			w.WriteHeader(responseWriter.Code)
			_, _ = buf.WriteTo(w)
		}()
		// Replace the response writer with the buffer
		responseWriter.Body = buf
		// Pass the modified response writer to the next handler
		next.ServeHTTP(responseWriter, r)
	})
}

func extractTokenFromHeader(headerValue string) (string, error) {
	if headerValue == "" {
		return "", fmt.Errorf("no Authorization header provided")
	}

	split := strings.Split(headerValue, " ")
	if len(split) == 2 && strings.EqualFold(split[0], "Bearer") {
		return split[1], nil
	}

	return "", fmt.Errorf("invalid Authorization header format")
}

func (m *Middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read the body to a string
		body, err := io.ReadAll(r.Body)
		if err != nil {
			response, _ := errors.HandleAPPError(fmt.Errorf("unauthorized: read body"))
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK) // This is important as you want to return a 200 status code
			_ = json.NewEncoder(w).Encode(response)

			return
		}
		// Set the body back after reading
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		// Check for the operations that don't require authentication
		if bytes.Contains(body, []byte("login")) ||
			bytes.Contains(body, []byte("refresh")) ||
			bytes.Contains(body, []byte("ForgotPassword")) ||
			bytes.Contains(body, []byte("ValidateMail")) ||
			bytes.Contains(body, []byte("ResetPassword")) {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		// Extract the token value from the header
		token, err := extractTokenFromHeader(authHeader)
		if err != nil {
			response, _ := errors.HandleAPPError(fmt.Errorf("unauthorized: extract token from header"))
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK) // This is important as you want to return a 200 status code
			_ = json.NewEncoder(w).Encode(response)

			return
		}

		// Attempt to retrieve the logged-in user's account using the extracted token
		loggedInAccount, err := m.Repo.GetLoggedInUser(token)
		if err != nil {
			response, _ := errors.HandleAPPError(fmt.Errorf("unauthorized: get user from token"))
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK) // This is important as you want to return a 200 status code
			_ = json.NewEncoder(w).Encode(response)

			return
		}

		userProfile, _ := m.Repo.GetUserProfileByUserAccountID(loggedInAccount.ID)
		organizationUnitID, _ := m.Repo.GetOrganizationUnitIDByUserProfile(userProfile.ID)

		// Create a new context that carries the necessary values
		ctx := context.WithValue(r.Context(), config.TokenKey, token)
		ctx = context.WithValue(ctx, config.LoggedInAccountKey, loggedInAccount)
		ctx = context.WithValue(ctx, config.LoggedInProfileKey, userProfile)
		ctx = context.WithValue(ctx, config.OrganizationUnitIDKey, organizationUnitID)

		// Call the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) GetCorsMiddleware(next http.Handler) http.Handler {
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{
			"http://localhost:3000",
			"http://localhost:3001",
			"http://localhost:3002",
			"http://localhost:3003",
			"http://localhost:3004",
			"http://localhost:3005",
			"https://localhost:3000",
			"https://127.0.0.1:3000",
			"https://localhost:3001",
			"https://localhost:3002",
			"https://localhost:3003",
			"https://localhost:3004",
			"https://localhost:3005",
			m.Config.Frontend.HR,
			m.Config.Frontend.Procurements,
			m.Config.Frontend.Accounting,
			m.Config.Frontend.Inventory,
			m.Config.Frontend.Finance,
			m.Config.Frontend.Core,
		}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)

	return corsMiddleware(next)
}
