package structs

import (
	"encoding/json"
	"fmt"
	"time"
)

type JSONDate string

const jsonDateFormat = "2006-01-02"

func (jd *JSONDate) MarshalJSON() ([]byte, error) {
	if jd == nil {
		return []byte("null"), nil
	}
	if *jd == "" {
		return []byte("null"), nil
	}

	t, err := time.Parse(jsonDateFormat, string(*jd))
	if err != nil {
		return nil, err
	}
	formattedDate := t.Format(time.RFC3339)
	return []byte(fmt.Sprintf(`"%s"`, formattedDate)), nil
}

func (jd *JSONDate) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var dateStr string
	if err := json.Unmarshal(data, &dateStr); err != nil {
		return err
	}

	*jd = JSONDate(dateStr)
	return nil
}

func (jd *JSONDate) ToTime() (time.Time, error) {
	if jd == nil || *jd == "" {
		return time.Time{}, nil
	}
	t, err := time.Parse(time.RFC3339, string(*jd))
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
