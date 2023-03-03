package method_sets_test

import "fmt"

type T struct {
	v int
}

func (t *T) setValue(v int) {
	t.v = v
}

func (t *T) getValue() int {
	return t.v
}

func Example_type_T_value_calls_methods_with_receiver_a_pointer_to_T() {
	// 按照 https://go.dev/ref/spec#Method_sets 的说法，类型 T 是没有方法的，
	// 类型 *T 有两个方法 setValue 和 getValue，
	// 这里 t 的类型为 T，但却可以调用 receiver 为 *T 的两个方法 setValue 和 getValue，
	// 应该是编译器做了工作，调用方法时取了 t 的指针。
	var t T
	t.setValue(10)
	fmt.Println(t.getValue())

	// Output:
	// 10
}
