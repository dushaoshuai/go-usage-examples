package align_test

import (
	"fmt"
	"reflect"
)

type foo struct {
	iInt8       int8
	iInt64      int64
	sString     string
	iInt16      int16
	fFloat64    float64
	iInt32      int32
	cComplex64  complex64
	bBool       bool
	fFloat32    float32
	cComplex128 complex128
}

func Example_alignment() {
	fmtStruct(foo{})
	// Output:
	// struct info:
	//
	// name: foo
	// size: 88 bytes
	// alignment: 8 bytes
	//
	// field info:
	//
	// name: iInt8
	// type: int8
	// size: 1 bytes
	// alignment: 1 bytes
	// offset: 0 bytes
	//
	// name: iInt64
	// type: int64
	// size: 8 bytes
	// alignment: 8 bytes
	// offset: 8 bytes
	//
	// name: sString
	// type: string
	// size: 16 bytes
	// alignment: 8 bytes
	// offset: 16 bytes
	//
	// name: iInt16
	// type: int16
	// size: 2 bytes
	// alignment: 2 bytes
	// offset: 32 bytes
	//
	// name: fFloat64
	// type: float64
	// size: 8 bytes
	// alignment: 8 bytes
	// offset: 40 bytes
	//
	// name: iInt32
	// type: int32
	// size: 4 bytes
	// alignment: 4 bytes
	// offset: 48 bytes
	//
	// name: cComplex64
	// type: complex64
	// size: 8 bytes
	// alignment: 4 bytes
	// offset: 52 bytes
	//
	// name: bBool
	// type: bool
	// size: 1 bytes
	// alignment: 1 bytes
	// offset: 60 bytes
	//
	// name: fFloat32
	// type: float32
	// size: 4 bytes
	// alignment: 4 bytes
	// offset: 64 bytes
	//
	// name: cComplex128
	// type: complex128
	// size: 16 bytes
	// alignment: 8 bytes
	// offset: 72 bytes
}

func fmtStruct(st any) {
	stType := reflect.TypeOf(st)
	if stType == nil {
		return
	}
	if stType.Kind() != reflect.Struct {
		return
	}
	fmt.Println("struct info:")
	fmt.Println()
	fmtName(stType.Name())
	fmtSize(stType.Size())
	fmtAlignment(stType.Align())
	fmt.Println()

	fmt.Println("field info:")
	fmt.Println()
	for i := 0; i < stType.NumField(); i++ {
		field := stType.Field(i)
		fmtName(field.Name)
		fmtType(field.Type.Name())
		fmtSize(field.Type.Size())
		fmtAlignment(field.Type.FieldAlign())
		fmtOffset(field.Offset)
		fmt.Println()
	}
}

func fmtName(name string) {
	fmt.Printf("name: %s\n", name)
}

func fmtType(t string) {
	fmt.Printf("type: %s\n", t)
}

func fmtSize(size uintptr) {
	fmtUintptrInfo("size", size)
}

func fmtAlignment(alignment int) {
	fmt.Printf("alignment: %d bytes\n", alignment)
}

func fmtOffset(offset uintptr) {
	fmtUintptrInfo("offset", offset)
}

func fmtUintptrInfo(header string, uiptr uintptr) {
	fmt.Printf("%s: %s\n", header, fmtUintptr(uiptr))
}

func fmtUintptr(uiptr uintptr) string {
	return fmt.Sprintf("%d bytes", uiptr)
}
