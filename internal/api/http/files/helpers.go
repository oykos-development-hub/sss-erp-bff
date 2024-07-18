package files

import (
	"bff/internal/api/errors"
	"bff/log"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func MarshalAndWriteJSON(w http.ResponseWriter, obj interface{}) error {
	jsonResponse, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, "Error during JSON marshaling", http.StatusInternalServerError)
		return errors.Wrap(err, "json marshal")
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)

	if err != nil {
		return errors.Wrap(err, "w write")
	}

	return nil
}

func openExcelFile(r *http.Request) (*excelize.File, error) {
	maxFileSize := int64(100 * 1024 * 1024) // file maximum 100 MB

	err := r.ParseMultipartForm(maxFileSize)
	if err != nil {
		return nil, errors.Wrap(err, "r parse multipart form")
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, errors.Wrap(err, "r form file")
	}
	defer file.Close()

	xlsFile, err := excelize.OpenReader(file)

	if err != nil {
		return nil, errors.Wrap(err, "excelize open reade")
	}

	return xlsFile, nil
}

func handleError(w http.ResponseWriter, err error, statusCode int) {
	log.Logger.Printf("Error: %s - %v", err.Error(), err)

	if statusCode < 100 || statusCode > 599 {
		statusCode = http.StatusInternalServerError
	}
	w.WriteHeader(statusCode)
	_ = MarshalAndWriteJSON(w, errorResponse{
		Message: err.Error(),
		Status:  "failed"},
	)
}

func makeBackendRequest(method, url string, body io.Reader, contentType string) (*http.Response, int, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, 0, errors.Wrap(err, "http new request")
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, 0, errors.Wrap(err, "http client do")
	}

	if resp.StatusCode != http.StatusOK {
		decoder := json.NewDecoder(resp.Body)
		var errorStruct errorResponse
		if err := decoder.Decode(&errorStruct); err != nil {
			return nil, resp.StatusCode, errors.Wrap(err, "decoder decode")
		}

		if errorStruct.Message != "" {
			return nil, resp.StatusCode, errors.Wrap(errors.New(errorStruct.Message), "http client do")
		}

		resp.Body.Close()
		return nil, resp.StatusCode, errors.Wrap(fmt.Errorf("backend returned non-OK status: %d", resp.StatusCode), "http client do")
	}

	return resp, 0, nil
}

func ExcelDateToTimeString(excelDate float64) string {
	t := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC).Add(time.Duration(excelDate*86400) * time.Second)
	return t.Format("2006-01-02T15:04:05Z")
}

func ConvertDateFormat(dateString string) (string, error) {
	// Parsiranje originalnog datuma
	t, err := time.Parse("01-02-06", dateString)
	if err != nil {
		return "", errors.Wrap(err, "time parse")
	}

	// Formatiranje u ISO 8601 format
	return t.Format("2006-01-02T15:04:05Z"), nil
}

func parseDate(dateString string) (time.Time, error) {
	var layouts = []string{
		"02.01.2006",                         // dd.mm.yyyy
		"02.01.2006.",                        // dd.mm.yyyy.
		"02/01/2006",                         // dd/mm/yyyy
		"02-01-2006",                         // dd-mm-yyyy
		"2006-01-02",                         // yyyy-mm-dd
		"01/02/2006",                         // mm/dd/yyyy
		"02.01.06",                           // dd.mm.yy
		"02.01.06.",                          // dd.mm.yy.
		"02/01/06",                           // dd/mm/yy
		"01/02/06",                           // mm/dd/yy
		"01-02-06",                           // mm-dd-yy
		"02-01-06",                           // dd-mm-yy
		"06-01-02",                           // yy-mm-dd
		"01/02/06",                           // mm/dd/yy
		"Jan 02, 2006",                       // Mon dd, yyyy
		"January 2, 2006",                    // Month dd, yyyy
		"02 Jan 2006",                        // dd Mon yyyy
		"2006-01-02T15:04:05Z07:00",          // ISO 8601
		"Mon Jan 02 15:04:05 -0700 MST 2006", // Go standard format
		"Jan _2 15:04:05",                    // Without leading zeros in day
		"2006-01-02 15:04:05",                // YYYY-MM-DD HH:MM:SS
		"02 Jan 06 15:04 MST",                // YY instead of YYYY
		"02 Jan 2006 15:04",                  // dd Mon yyyy HH:MM
		"Mon, 02 Jan 2006 15:04:05 MST",      // Day, dd Mon yyyy HH:MM:SS
		"Mon, 02 Jan 2006 15:04:05 -0700",    // Day, dd Mon yyyy HH:MM:SS Timezone Offset
		"Monday, 02-Jan-06 15:04:05 MST",     // Full Day, dd-Mon-yy HH:MM:SS
		"Monday, 02-Jan-06 15:04:05 -0700",   // Full Day, dd-Mon-yy HH:MM:SS Timezone Offset
		"3:04 PM",                            // Short time
		"Jan 02, 2006 at 3:04pm",             // Date with time
	}

	var date time.Time
	var err error

	for _, layout := range layouts {
		date, err = time.Parse(layout, dateString)
		if err == nil {
			return date, nil
		}
	}

	numberOfDays, err := strconv.Atoi(dateString)

	if err == nil {
		startDate := time.Date(1899, time.December, 31, 0, 0, 0, 0, time.UTC)
		daysDuration := time.Duration(numberOfDays) * 24 * time.Hour
		return startDate.Add(daysDuration), nil
	}

	return date, errors.Wrap(fmt.Errorf("date format is not valid: %s", dateString), "time parse")
}
