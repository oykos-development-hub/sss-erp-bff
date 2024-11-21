package sse

import (
	"bff/config"
	"bff/structs"
	"net/http"
	"sync"

	"bff/internal/api/middleware"

	"github.com/go-chi/chi/v5"
)

type ServerSentEvent struct {
	Clients map[int][]chan string
	mu      sync.Mutex
}

func NewServerSentEvent() *ServerSentEvent {
	return &ServerSentEvent{
		Clients: make(map[int][]chan string),
	}
}

func (sse *ServerSentEvent) AddClient(userID int) chan string {
	clientChan := make(chan string)

	sse.mu.Lock()
	defer sse.mu.Unlock()
	sse.Clients[userID] = append(sse.Clients[userID], clientChan)

	return clientChan
}

func (sse *ServerSentEvent) RemoveClient(userID int, clientChan chan string) {
	sse.mu.Lock()
	defer sse.mu.Unlock()

	clients := sse.Clients[userID]
	for i, ch := range clients {
		if ch == clientChan {
			sse.Clients[userID] = append(clients[:i], clients[i+1:]...)
			break
		}
	}

	if len(sse.Clients[userID]) == 0 {
		delete(sse.Clients, userID)
	}
}

func (sse *ServerSentEvent) Broadcast(userID int, message string) {
	sse.mu.Lock()
	defer sse.mu.Unlock()

	for _, clientChan := range sse.Clients[userID] {
		clientChan <- message
	}
}

func (sse *ServerSentEvent) Handler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(config.LoggedInAccountKey).(*structs.UserAccounts)

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	clientChan := sse.AddClient(user.ID)
	defer func() {
		sse.RemoveClient(user.ID, clientChan)
		close(clientChan)
	}()

	for msg := range clientChan {
		_, err := w.Write([]byte("data: " + msg + "\n\n"))
		if err != nil {
			break
		}
		w.(http.Flusher).Flush()
	}
}

func (sse *ServerSentEvent) Router(m *middleware.Middleware) http.Handler {
	r := chi.NewRouter()
	r.Get("/stream", sse.Handler)

	return m.GetCorsMiddleware(
		m.AuthMiddlewareSSE(
			m.RequestContextMiddleware(r),
		),
	)
}
