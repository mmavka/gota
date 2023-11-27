package series

import (
	"fmt"
	"math"
	"time"
)

type interfaceElement struct {
	e   interface{}
	nan bool
}

// force interfaceElement struct to implement Element interface
var _ Element = (*interfaceElement)(nil)

func (e *interfaceElement) Set(value interface{}) {
	e.nan = false
	switch val := value.(type) {
	case string:
		if val == "NaN" {
			e.nan = true
			return
		}
		e.e = val
	case int:
		e.e = val
	case float64:
		e.e = val
	case bool:
		e.e = val
	case Element:
		v, err := val.Interface()
		if err != nil {
			e.nan = true
			return
		}
		e.e = v
	default:
		e.e = value
	}
}

func (e interfaceElement) Copy() Element {
	if e.IsNA() {
		return &interfaceElement{0, true}
	}
	return &interfaceElement{e.e, false}
}

func (e interfaceElement) IsNA() bool {
	return e.nan
}

func (e interfaceElement) Type() Type {
	return Int
}

func (e interfaceElement) Val() ElementValue {
	if e.IsNA() {
		return nil
	}
	return e.e
}

func (e interfaceElement) String() string {
	if e.IsNA() {
		return "NaN"
	}
	return fmt.Sprint(e.e)
}

func (e interfaceElement) Int() (int, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to int")
	}
	v, ok := e.e.(int)
	if !ok {
		return 0, fmt.Errorf("can't convert \"%v\" to int", e.e)
	}
	return v, nil
}

func (e *interfaceElement) Int64() (int64, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to int64")
	}
	v, ok := e.e.(int64)
	if !ok {
		return 0, fmt.Errorf("can't convert \"%v\" to int64", e.e)
	}
	return v, nil
}

func (e interfaceElement) Float() float64 {
	if e.IsNA() {
		return math.NaN()
	}
	v, ok := e.e.(float64)
	if !ok {
		return math.NaN()
	}
	return v
}

func (e interfaceElement) Bool() (bool, error) {
	if e.IsNA() {
		return false, fmt.Errorf("can't convert NaN to bool")
	}
	switch e.e {
	case 1:
		return true, nil
	case 0:
		return false, nil
	}
	return false, fmt.Errorf("can't convert Interface \"%v\" to bool", e.e)
}

func (e interfaceElement) Time() (time.Time, error) {
	return time.Time{}, fmt.Errorf("can't convert Interface to time.Time")
}

func (e interfaceElement) Interface() (interface{}, error) {
	if e.IsNA() {
		return false, fmt.Errorf("can't convert NaN to interface")
	}
	return e.e, nil
}

func (e interfaceElement) Eq(elem Element) bool {
	i, err := elem.Interface()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e == i
}

func (e interfaceElement) Neq(elem Element) bool {
	i, err := elem.Interface()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e != i
}

func (e interfaceElement) Less(elem Element) bool {
	panic("Interface type can't be compared")
	//i, err := elem.Interface()
	//if err != nil || e.IsNA() {
	//	return false
	//}
	//return e.e < i
}

func (e interfaceElement) LessEq(elem Element) bool {
	panic("Interface type can't be compared")
	//i, err := elem.Interface()
	//if err != nil || e.IsNA() {
	//	return false
	//}
	//return e.e <= i
}

func (e interfaceElement) Greater(elem Element) bool {
	panic("Interface type can't be compared")
	//i, err := elem.Interface()
	//if err != nil || e.IsNA() {
	//	return false
	//}
	//return e.e > i
}

func (e interfaceElement) GreaterEq(elem Element) bool {
	panic("Interface type can't be compared")
	//i, err := elem.Interface()
	//if err != nil || e.IsNA() {
	//	return false
	//}
	//return e.e >= i
}
