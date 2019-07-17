package snoo

import (
	"strconv"
	"strings"
	"time"
)

// IsoDate is a type to marshal unix to ISO Datetime
type ISODate struct {
	time.Time
}

func (d *ISODate) UnmarshalJSON(b []byte) (err error) {
	f, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		return err
	}
	i := int64(f)
	d.Time = time.Unix(i, 0)
	return
}

type NumBool struct {
	Val bool
	Num float64
}

func (f *NumBool) UnmarshalJSON(b []byte) (err error) {
	switch str := strings.ToLower(strings.Trim(string(b), `"`)); str {
	case "true":
		f.Val = true
	case "false":
		f.Val = false
	default:
		f.Num, err = strconv.ParseFloat(str, 64)
		if f.Num > 0 {
			f.Val = true
		}
	}
	return err
}
