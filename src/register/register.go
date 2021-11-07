package register

import "errors"

const Size = 14


// Register - interface for representing
// registers in our WM with Size const as
// count of bits in it.
// Implements Token interface.
type Register interface {
	Value() int16
	Change(val int16)

	// Set - change register's bit by its pos.
	// if pos > Size -> return error
	Set(pos uint, val uint) error

	// Get - get bit in register by its pos.
	// if pos > Size -> return error
	Get(pos uint) (uint, error)

	// ToSlice - return slice of bits of register
	ToSlice() []byte
}

type register struct {
	bits uint16
}

// New - Register constructor
// from register struct
func New() Register {
	return &register{0}
}

// ToSlice - return slice of bits of register
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

// Value - return Register init value
func (reg register) Value() int16 {
	return int16(reg.bits)
}

// Change - change register init value
func (reg *register) Change(val int16) {
	reg.bits = uint16(val)
}


// May be will be helpful in future

func (reg *register) Set(pos uint, value uint) error {
	if pos > Size-1 {
		return errors.New("no such bit")
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
		return 0, errors.New("no such bit")
	}
	if (reg.bits & uint16(1 << pos)) != 0 {
		return 1, nil
	} else {
		return 0, nil
	}
}