package token

type Token interface {
    Value() int16
	Change(val int16)
}

type value struct {
	value int16
}

func (v value) Value() int16 {
	return v.value
}

func New(num int16) Token {
	return &value {
		value: num,
	}
}

func (v *value) Change(val int16) {
	v.value = val
}
