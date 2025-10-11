package exported_test

import (
	"fmt"
	"go/ast"
	"reflect"
	"unicode"

	"github.com/dushaoshuai/go-usage-examples/golang/exported"
)

func Example_is_not_an_upper_case_letter() {
	fmt.Println(ast.IsExported("_NamedFieldsRequired"))
	fmt.Println(unicode.IsUpper('_'))

	// Output:
	// false
	// false
}

func Example_is_exported() {
	_ = exported.NamedArg{Name: "arg0", Value: 0}
	// An element list that does not contain any keys must list an element for each struct field in the order in which the fields are declared.
	// _ = exported.NamedArg{"arg0", 0} // Too few values, Cannot assign a value to the unexported field '_NamedFieldsRequired'

	_ = exported.NamedArg3{Name: "arg1", Value: 1}
	// An element list that does not contain any keys must list an element for each struct field in the order in which the fields are declared.
	// _ = exported.NamedArg3{"arg2", 2} // Too few values, Cannot assign a value to the unexported field 'namedFieldsRequired'

	f := func(rt reflect.Type) {
		if rt.Kind() != reflect.Struct {
			return
		}

		for i := range rt.NumField() {
			field := rt.Field(i)
			fmt.Printf("%+v", field)
			fmt.Printf("IsExported: %t\n", field.IsExported())
		}

		fmt.Println()
	}

	f(reflect.TypeOf(exported.NamedArg{}))
	f(reflect.TypeOf(exported.NamedArg2{}))
	f(reflect.TypeOf(exported.NamedArg3{}))

	// Output:
	// {Name:_NamedFieldsRequired PkgPath:github.com/dushaoshuai/go-usage-examples/golang/exported Type:struct {} Tag: Offset:0 Index:[0] Anonymous:false}IsExported: false
	// {Name:Name PkgPath: Type:string Tag: Offset:0 Index:[1] Anonymous:false}IsExported: true
	// {Name:Value PkgPath: Type:interface {} Tag: Offset:16 Index:[2] Anonymous:false}IsExported: true
	//
	// {Name:NamedFieldsRequired PkgPath: Type:struct {} Tag: Offset:0 Index:[0] Anonymous:false}IsExported: true
	// {Name:Name PkgPath: Type:string Tag: Offset:0 Index:[1] Anonymous:false}IsExported: true
	// {Name:Value PkgPath: Type:interface {} Tag: Offset:16 Index:[2] Anonymous:false}IsExported: true
	//
	// {Name:namedFieldsRequired PkgPath:github.com/dushaoshuai/go-usage-examples/golang/exported Type:struct {} Tag: Offset:0 Index:[0] Anonymous:false}IsExported: false
	// {Name:Name PkgPath: Type:string Tag: Offset:0 Index:[1] Anonymous:false}IsExported: true
	// {Name:Value PkgPath: Type:interface {} Tag: Offset:16 Index:[2] Anonymous:false}IsExported: true
}
