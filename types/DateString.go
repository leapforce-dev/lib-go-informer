package facebook

import (
	"encoding/json"
	"fmt"
	"time"

	"cloud.google.com/go/civil"
	errortools "github.com/leapforce-libraries/go_errortools"
)

const (
	DateFormat string = "2006-01-02"
)

type DateString civil.Date

func (d *DateString) UnmarshalJSON(b []byte) error {
	var returnError = func() error {
		errortools.CaptureError(fmt.Sprintf("Cannot parse '%s' to DateString", string(b)))
		return nil
	}

	var s string

	err := json.Unmarshal(b, &s)
	if err != nil {
		return returnError()
	}

	if s == "" || s == "0000-00-00" || s == "0000-00-00 00:00:00" || s == "9999-01-01" {
		d = nil
		return nil
	}

	_t, err := time.Parse(DateFormat, s)
	if err != nil {
		return returnError()
	}

	*d = DateString(civil.DateOf(_t))
	return nil
}

func (d *DateString) ValuePtr() *civil.Date {
	if d == nil {
		return nil
	}

	_d := civil.Date(*d)
	return &_d
}

func (d DateString) Value() civil.Date {
	return civil.Date(d)
}
