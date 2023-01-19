package common

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
)

type DateOnly struct {
	time.Time
}

func (t *DateOnly) UnmarshalJSON(b []byte) error {
	date, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return nil
}
func (t *DateOnly) MarshalJSON() ([]byte, error) {
	if y := t.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	date := strings.Split(t.String(), " ")
	return json.Marshal(date[0])
}

func StringToDate(s string) (*time.Time, error) {
	date := strings.Split(s, "-")
	year, err := strconv.Atoi(date[0])
	if err != nil {
		return nil, err
	}
	month, err := strconv.Atoi(date[1])
	if err != nil {
		return nil, err
	}
	day, err := strconv.Atoi(date[2])
	if err != nil {
		return nil, err
	}
	dateTime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return &dateTime, nil
}
