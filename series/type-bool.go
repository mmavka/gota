package series

import (
	"fmt"
	"math"
	"strings"
	"time"
)

type boolElement struct {
	e   bool
	nan bool
}

// force boolElement struct to implement Element interface
var _ Element = (*boolElement)(nil)

func (e *boolElement) Set(value interface{}) {
	e.nan = false
	switch val := value.(type) {
	case string:
		if val == "NaN" {
			e.nan = true
			return
		}
		switch strings.ToLower(value.(string)) {
		case "true", "t", "1":
			e.e = true
		case "false", "f", "0":
			e.e = false
		default:
			e.nan = true
			return
		}
	case int:
		switch val {
		case 1:
			e.e = true
		case 0:
			e.e = false
		default:
			e.nan = true
			return
		}
	case float64:
		switch val {
		case 1:
			e.e = true
		case 0:
			e.e = false
		default:
			e.nan = true
			return
		}
	case bool:
		e.e = val
	case Element:
		b, err := value.(Element).Bool()
		if err != nil {
			e.nan = true
			return
		}
		e.e = b
	default:
		e.nan = true
		return
	}
}

func (e boolElement) Copy() Element {
	if e.IsNA() {
		return &boolElement{false, true}
	}
	return &boolElement{e.e, false}
}

func (e boolElement) IsNA() bool {
	return e.nan
}

func (e boolElement) Type() Type {
	return Bool
}

func (e boolElement) Val() ElementValue {
	if e.IsNA() {
		return nil
	}
	return bool(e.e)
}

func (e boolElement) String() string {
	if e.IsNA() {
		return "NaN"
	}
	if e.e {
		return "true"
	}
	return "false"
}

func (e boolElement) Int() (int, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to int")
	}
	if e.e {
		return 1, nil
	}
	return 0, nil
}

func (e *boolElement) Int64() (int64, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to int64")
	}
	if e.e {
		return 1, nil
	}
	return 0, nil
}

func (e boolElement) Float() float64 {
	if e.IsNA() {
		return math.NaN()
	}
	if e.e {
		return 1.0
	}
	return 0.0
}

func (e boolElement) Bool() (bool, error) {
	if e.IsNA() {
		return false, fmt.Errorf("can't convert NaN to bool")
	}
	return bool(e.e), nil
}

func (e boolElement) Time() (time.Time, error) {
	return time.Time{}, fmt.Errorf("can't convert Bool to time.Time")
}

func (e boolElement) Interface() (interface{}, error) {
	if e.IsNA() {
		return false, fmt.Errorf("can't convert NaN to interface")
	}
	return e.e, nil
}

func (e boolElement) Eq(elem Element) bool {
	b, err := elem.Bool()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e == b
}

func (e boolElement) Neq(elem Element) bool {
	b, err := elem.Bool()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e != b
}

func (e boolElement) Less(elem Element) bool {
	b, err := elem.Bool()
	if err != nil || e.IsNA() {
		return false
	}
	return !e.e && b
}

func (e boolElement) LessEq(elem Element) bool {
	b, err := elem.Bool()
	if err != nil || e.IsNA() {
		return false
	}
	return !e.e || b
}

func (e boolElement) Greater(elem Element) bool {
	b, err := elem.Bool()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e && !b
}

func (e boolElement) GreaterEq(elem Element) bool {
	b, err := elem.Bool()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e || !b
}
