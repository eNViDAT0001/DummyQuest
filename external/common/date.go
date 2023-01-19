package common

import (
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
