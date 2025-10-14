package type_parameters

type mixedIface interface {
	~string | ~struct{} | ~struct {
		strVal string
		intVal int
	} | ~struct {
		int64Val int64
	} | herStruct

	echo()
}

func domixedIface[T mixedIface](v T) {
	v.echo()
}

func Example_mixedIface() {
	domixedIface(myString(""))
	domixedIface(myEmptyStruct{})
	domixedIface(myStruct{})
	domixedIface(yourStruct{})
	domixedIface(herStruct{})

	// Output:
	// myString implements mixed
	// myEmptyStruct implements mixed
	// myStruct implements mixed
	// yourStruct implements mixed
	// herStruct implements mixed
}
