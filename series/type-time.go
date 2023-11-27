package series

import (
	"fmt"
	"time"
)

// timeElement is the concrete implementation of the Element interface for
// time.Time. If the stored time.Time is zero, it will be considered as a NaN
// element.
type timeElement struct {
	e   time.Time
	nan bool
}

func (e *timeElement) Set(value interface{}) {
	e.nan = false
	switch val := value.(type) {
	case string:
		if val == "NaN" {
			e.nan = true
			return
		}
		t, err := time.Parse(timeformat, val)
		if err != nil {
			e.nan = true
			return
		}
		e.e = t
	case int64:
		e.e = time.UnixMilli(val)
	case int:
		e.e = time.Unix(int64(val), 0)
	case time.Time:
		e.e = val
	case Element:
		t, err := val.Time()
		if err != nil {
			e.nan = true
			return
		}
		e.e = t
	default:
		e.nan = true
		return
	}
}

func (e timeElement) Copy() Element {
	if e.IsNA() {
		return &timeElement{time.Time{}, false}
	}
	return &timeElement{e.e, false}
}

func (e timeElement) IsNA() bool {
	return e.e.IsZero()
}

func (e timeElement) Type() Type {
	return Time
}

func (e timeElement) Val() ElementValue {
	if e.IsNA() {
		return nil
	}
	return e.e
}

func (e timeElement) String() string {
	if e.IsNA() {
		return "NaN"
	}
	return e.e.Format(timeformat)
}

func (e timeElement) Int() (int, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to int")
	}
	return int(e.e.Unix()), nil
}

func (e *timeElement) Int64() (int64, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to int64")
	}
	return e.e.Unix(), nil
}

func (e timeElement) Float() float64 {
	return float64(e.e.Unix())
}

func (e timeElement) Bool() (bool, error) {
	if e.IsNA() {
		return false, fmt.Errorf("can't convert NaN to bool")
	}
	return false, fmt.Errorf("can't convert Time to bool")
}

func (e timeElement) Time() (time.Time, error) {
	if e.IsNA() {
		return time.Time{}, fmt.Errorf("can't convert NaN to time.Time")
	}
	return e.e, nil
}

func (e timeElement) Eq(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	t, _ := elem.Time()
	return e.e.Equal(t)
}

func (e timeElement) Neq(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	t, _ := elem.Time()
	return !e.e.Equal(t)
}

func (e timeElement) Less(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	t, _ := elem.Time()
	return e.e.Before(t)
}

func (e timeElement) LessEq(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	t, _ := elem.Time()
	if e.e.Equal(t) || e.e.Before(t) {
		return true
	}
	return false
}

func (e timeElement) Greater(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	t, _ := elem.Time()
	return e.e.After(t)
}

func (e timeElement) GreaterEq(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	t, _ := elem.Time()
	if e.e.Equal(t) || e.e.After(t) {
		return true
	}
	return false
}
