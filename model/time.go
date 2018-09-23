package model

import (
	"strings"
	"time"
)

type customTime struct {
	time.Time
}

func (t *customTime) UnmarshalJSON(buf []byte) error {
	tt, err := time.Parse("2006-01-02", strings.Trim(string(buf), `"`))
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}
func (t customTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Time.Format("2006-01-01") + `"`), nil
}