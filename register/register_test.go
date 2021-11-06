package register

import (
	"testing"
)

func TestRegisterSetError(t *testing.T) {
	reg := New()
	err := reg.Set(Size, 0)
	if err == nil {
		t.Fatal("Must return error if pos is too greate")
	}
}

func TestRegisterGetError(t *testing.T) {
	reg := New()
	_, err := reg.Get(Size)
	if err == nil {
		t.Fatal("Must return error if pos is too greate")
	}
}

func TestRegisterSetGet(t *testing.T) {
	reg := register{0}
	for i := uint(0); i < Size; i++ {
		reg.Set(i, 1)
		if value, _ := reg.Get(i); value != 1 {
			t.Fatalf("Get or Set is not working: %v", reg.ToSlice())
		}
	}
}
