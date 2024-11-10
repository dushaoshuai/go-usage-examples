package bitwise_op

import (
	"testing"
)

func Test_integerBits(t *testing.T) {
	// 0

	if want, got := "0", integerBits(uint(0)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "uint(0)", got, want)
	}
	if want, got := "0", integerBits(int(+0)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int(+0)", got, want)
	}
	if want, got := "0", integerBits(int(-0)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int(-0)", got, want)
	}

	if want, got := "0", integerBits(uint8(0)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "uint8(0)", got, want)
	}
	if want, got := "0", integerBits(int8(+0)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int8(+0)", got, want)
	}
	if want, got := "0", integerBits(int8(-0)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int8(-0)", got, want)
	}

	if want, got := "0", integerBits(uint16(0)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "uint16(0)", got, want)
	}
	if want, got := "0", integerBits(int16(+0)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int16(+0)", got, want)
	}
	if want, got := "0", integerBits(int16(-0)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int16(-0)", got, want)
	}

	if want, got := "0", integerBits(uint32(0)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "uint32(0)", got, want)
	}
	if want, got := "0", integerBits(int32(+0)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int32(+0)", got, want)
	}
	if want, got := "0", integerBits(int32(-0)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int32(-0)", got, want)
	}

	if want, got := "0", integerBits(uint64(0)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "uint64(0)", got, want)
	}
	if want, got := "0", integerBits(int64(+0)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int64(+0)", got, want)
	}
	if want, got := "0", integerBits(int64(-0)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int64(-0)", got, want)
	}

	// 1

	if want, got := "1", integerBits(uint(1)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "uint(1)", got, want)
	}
	if want, got := "1", integerBits(int(+1)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int(+1)", got, want)
	}
	if want, got := "1111111111111111111111111111111111111111111111111111111111111111", integerBits(int(-1)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int(-1)", got, want)
	}

	if want, got := "1", integerBits(uint8(1)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "uint8(1)", got, want)
	}
	if want, got := "1", integerBits(int8(+1)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int8(+1)", got, want)
	}
	if want, got := "11111111", integerBits(int8(-1)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int8(-1)", got, want)
	}

	if want, got := "1", integerBits(uint16(1)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "uint16(1)", got, want)
	}
	if want, got := "1", integerBits(int16(+1)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int16(+1)", got, want)
	}
	if want, got := "1111111111111111", integerBits(int16(-1)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int16(-1)", got, want)
	}

	if want, got := "1", integerBits(uint32(1)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "uint32(1)", got, want)
	}
	if want, got := "1", integerBits(int32(+1)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int32(+1)", got, want)
	}
	if want, got := "11111111111111111111111111111111", integerBits(int32(-1)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int32(-1)", got, want)
	}

	if want, got := "1", integerBits(uint64(1)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "uint64(1)", got, want)
	}
	if want, got := "1", integerBits(int64(+1)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int64(+1)", got, want)
	}
	if want, got := "1111111111111111111111111111111111111111111111111111111111111111", integerBits(int64(-1)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int64(-1)", got, want)
	}

	// 45

	if want, got := "101101", integerBits(uint(45)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "uint(45)", got, want)
	}
	if want, got := "101101", integerBits(int(+45)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int(+45)", got, want)
	}
	if want, got := "1111111111111111111111111111111111111111111111111111111111010011", integerBits(int(-45)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int(-45)", got, want)
	}

	if want, got := "101101", integerBits(uint8(45)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "uint8(45)", got, want)
	}
	if want, got := "101101", integerBits(int8(+45)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int8(+45)", got, want)
	}
	if want, got := "11010011", integerBits(int8(-45)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int8(-45)", got, want)
	}

	if want, got := "101101", integerBits(uint16(45)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "uint16(45)", got, want)
	}
	if want, got := "101101", integerBits(int16(+45)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int16(+45)", got, want)
	}
	if want, got := "1111111111010011", integerBits(int16(-45)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int16(-45)", got, want)
	}

	if want, got := "101101", integerBits(uint32(45)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "uint32(45)", got, want)
	}
	if want, got := "101101", integerBits(int32(+45)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int32(+45)", got, want)
	}
	if want, got := "11111111111111111111111111010011", integerBits(int32(-45)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int32(-45)", got, want)
	}

	if want, got := "101101", integerBits(uint64(45)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "uint64(45)", got, want)
	}
	if want, got := "101101", integerBits(int64(+45)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int64(+45)", got, want)
	}
	if want, got := "1111111111111111111111111111111111111111111111111111111111010011", integerBits(int64(-45)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int64(-45)", got, want)
	}

	// 127

	if want, got := "1111111", integerBits(uint(127)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "uint(127)", got, want)
	}
	if want, got := "1111111", integerBits(int(+127)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int(+127)", got, want)
	}
	if want, got := "1111111111111111111111111111111111111111111111111111111110000001", integerBits(int(-127)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int(-127)", got, want)
	}

	if want, got := "1111111", integerBits(uint8(127)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "uint8(127)", got, want)
	}
	if want, got := "1111111", integerBits(int8(+127)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int8(+127)", got, want)
	}
	if want, got := "10000001", integerBits(int8(-127)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int8(-127)", got, want)
	}

	if want, got := "1111111", integerBits(uint16(127)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "uint16(127)", got, want)
	}
	if want, got := "1111111", integerBits(int16(+127)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int16(+127)", got, want)
	}
	if want, got := "1111111110000001", integerBits(int16(-127)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int16(-127)", got, want)
	}

	if want, got := "1111111", integerBits(uint32(127)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "uint32(127)", got, want)
	}
	if want, got := "1111111", integerBits(int32(+127)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int32(+127)", got, want)
	}
	if want, got := "11111111111111111111111110000001", integerBits(int32(-127)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int32(-127)", got, want)
	}

	if want, got := "1111111", integerBits(uint64(127)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "uint64(127)", got, want)
	}
	if want, got := "1111111", integerBits(int64(+127)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int64(+127)", got, want)
	}
	if want, got := "1111111111111111111111111111111111111111111111111111111110000001", integerBits(int64(-127)); got != want {
		t.Errorf("integerBits(%s) = %v, want %v", "int64(-127)", got, want)
	}
}
