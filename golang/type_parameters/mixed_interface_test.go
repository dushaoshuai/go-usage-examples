package type_parameters

import "fmt"

type mixed interface {
	~string | ~struct{} | ~struct {
		strVal string
		intVal int
	} | ~struct {
		int64Val int64
	} | herStruct

	echo()
}

func doMixed[T mixed](v T) {
	v.echo()
}

type myString string

func (ms myString) echo() {
	fmt.Println("myString implements mixed")
}

type myEmptyStruct struct{}

func (mes myEmptyStruct) echo() {
	fmt.Println("myEmptyStruct implements mixed")
}

type myStruct struct {
	strVal string
	intVal int
}

func (ms myStruct) echo() {
	fmt.Println("myStruct implements mixed")
}

type yourStruct struct {
	int64Val int64
}

func (ys yourStruct) echo() {
	fmt.Println("yourStruct implements mixed")
}

type herStruct struct {
	float64Val float64
}

func (hs herStruct) echo() {
	fmt.Println("herStruct implements mixed")
}

func Example_mixed() {
	doMixed(myString(""))
	doMixed(myEmptyStruct{})
	doMixed(myStruct{})
	doMixed(yourStruct{})
	doMixed(herStruct{})

	// Output:
	// myString implements mixed
	// myEmptyStruct implements mixed
	// myStruct implements mixed
	// yourStruct implements mixed
	// herStruct implements mixed
}
