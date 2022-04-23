package functional_programming

import "fmt"

// 函数式选项的一些参考，排列方式按照我认为的参考性从高到低
// https://www.sohamkamani.com/golang/options-pattern/
// https://pkg.go.dev/github.com/robfig/cron/v3#Cron
// https://coolshell.cn/articles/21146.html#%E9%85%8D%E7%BD%AE%E5%AF%B9%E8%B1%A1%E6%96%B9%E6%A1%88
// https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html

// 例子代码来自👆列表中第一个地址
// 第二个地址中，cron 的实现也很标准

type House struct {
	material     string
	HasFireplace bool
	floors       int
}

// define the constructor
func NewHouse(opts ...HouseOption) *House {
	h := &House{
		material:     "wood",
		HasFireplace: true,
		floors:       2,
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

// define the functional options
type HouseOption func(*House)

func WithConcrete() HouseOption {
	return func(house *House) {
		house.material = "concrete"
	}
}

func WithoutFireplace() HouseOption {
	return func(house *House) {
		house.HasFireplace = false
	}
}

func WithFloors(floors int) HouseOption {
	return func(house *House) {
		house.floors = floors
	}
}

func Example_functional_options() {
	house := NewHouse(
		WithoutFireplace(),
		WithFloors(4),
	)
	fmt.Printf("%+v", house)
	// Output:
	// &{material:wood HasFireplace:false floors:4}
}
