package token

// Token interface that represents object of operand.
// Can be used for wrapping simple numbers and
// Register interface.
type Token interface {
	// Value method for getting value from operand
    Value() int16

	// Change method for changing inside value
	Change(val int16)
}

// Struct for simple numbers
type simpleValue struct {
	value int16
}

// NewSimpleValue - constructor for simple simpleValue
func NewSimpleValue(num int16) Token {
	return &simpleValue{
		value: num,
	}
}

func (v simpleValue) Value() int16 {
	return v.value
}

func (v *simpleValue) Change(val int16) {
	v.value = val
}
