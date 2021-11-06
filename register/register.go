package register

import (
	"errors"
)

const Size = 14

type Register interface {
	Set(pos uint, value uint) error
	Get(pos uint) (uint, error)

	Value() int16
	Change(val int16)

	ToSlice() []byte
}

type register struct {
	bits uint16
}

func New() Register {
	return &register{0}
}

func (reg *register) Set(pos uint, value uint) error {
	if pos > Size-1 {
		return errors.New("No such bit")
	}
	if value != 0 {
		reg.bits = reg.bits | (1 << pos)
	} else {
		reg.bits = reg.bits | (0 << pos)
	}
	return nil
}

func (reg register) Get(pos uint) (uint, error) {
	if pos > Size-1 {
		return 0, errors.New("No such bit")
	}
	if (reg.bits & uint16(1 << pos)) != 0 {
		return 1, nil
	} else {
		return 0, nil
	}
}

func (reg register) ToSlice() []byte {
	var num uint16
	res := make([]byte, Size)

	num = reg.bits
	for i := 0; i < Size && num > 0; i++ {
		res[i] = byte(num % 2)
		num /= 2
	}

	return res
}

func (reg register) Value() int16 {
	return int16(reg.bits)
}


func (reg *register) Change(val int16) {
	reg.bits = uint16(val)
}
