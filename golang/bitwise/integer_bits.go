package bitwise

import (
	"fmt"
	"reflect"

	"golang.org/x/exp/constraints"
)

func integerBits[Integer constraints.Integer](integer Integer) string {
	switch reflect.TypeOf(integer).Kind() {
	case reflect.Int:
		return fmt.Sprintf("%b", uint(integer))
	case reflect.Int8:
		return fmt.Sprintf("%b", uint8(integer))
	case reflect.Int16:
		return fmt.Sprintf("%b", uint16(integer))
	case reflect.Int32:
		return fmt.Sprintf("%b", uint32(integer))
	case reflect.Int64:
		return fmt.Sprintf("%b", uint64(integer))
	default:
		return fmt.Sprintf("%b", integer)
	}
}
