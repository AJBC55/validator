package validator

// required valued interface used to custom required values and validation logic
type RequiredValue interface {
	valid() bool
}

type RequiredString string
type RequiredInt int
type RequiredBool bool
type RequiredFloat float64
type RequiredSlice []interface{}

func (i RequiredInt) valid() bool {
	return i == 0
}

func (b RequiredBool) valid() bool {
	return b == false
}

func (f RequiredFloat) valid() bool {
	return f == 0.0
}

func (s RequiredString) valid() bool {
	return s == ""
}

func (s RequiredSlice) valid() bool {
	return s == nil
}
