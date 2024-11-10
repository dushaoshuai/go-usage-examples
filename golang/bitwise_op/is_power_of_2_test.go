package bitwise_op

import (
	"math"
	"testing"
)

// x > 0 && x&(x-1) == 0
// 判断 x 是不是 2 的幂

func Test_is_power_of_2(t *testing.T) {
	powerOf2 := map[int8]bool{}
	for i := range 7 {
		powerOf2[int8(math.Exp2(float64(i)))] = true
	}

	for x := math.MinInt8; x <= math.MaxInt8; x++ {
		if x > 0 && x&(x-1) == 0 && !powerOf2[int8(x)] {
			t.Errorf("%d should not be power of 2", x)
		}
	}
}
