package json_test

import (
	"encoding/json"
	"os"
)

func marshal(a any) {
	b, _ := json.Marshal(a)
	_, _ = os.Stdout.Write(b)
}

type anonymousA struct {
	A int
	B string
}

type structA struct {
	anonymousA // 未暴露的匿名字段的内部暴露字段也会被 marshal
}

func ExampleMarshal_unexported_anonymous_struct_fields() {
	marshal(
		structA{
			anonymousA: anonymousA{
				A: 10,
				B: "good",
			},
		},
	)

	// Output:
	// {"A":10,"B":"good"}
}

type anonymousB struct {
	A int
	B string
}

type structB struct {
	anonymousA
	anonymousB `json:"ab"` // 在 JSON tag 中有 name 的匿名字段不会被当作匿名字段进行 marshal
}

func ExampleMarshal_anonymous_struct_fields_with_name_in_JSON_tag() {
	marshal(
		structB{
			anonymousA: anonymousA{
				A: 10,
				B: "Java",
			},
			anonymousB: anonymousB{
				A: 15,
				B: "Python",
			},
		},
	)

	// Output:
	// {"A":10,"B":"Java","ab":{"A":15,"B":"Python"}}
}

type Any any

type structC struct {
	anonymousA
	any // 接口类型的匿名字段不会被当作匿名字段，字段的类型就是它的字段名，any 的类型是 any，name 是 any，小写，因此没有 marshal，
	Any // Any 的类型是 Any，name 是 Any，大写，因此被 marshal.
}

func ExampleMarshal_anonymous_struct_fields_of_interface_type() {
	marshal(
		structC{
			anonymousA: anonymousA{
				A: 90,
				B: "Perl",
			},
			any: anonymousB{
				A: 45,
				B: "C",
			},
			Any: anonymousB{
				A: 56,
				B: "Shell",
			},
		},
	)

	// Output:
	// {"A":90,"B":"Perl","Any":{"A":56,"B":"Shell"}}
}

type structD struct {
	anonymousA // 两个结构体中有相同的字段，冲突，
	anonymousB // 所以没有一个字段被 marshal
}

func ExampleMarshal_anonymous_struct_fields_with_inner_exported_fields_conflict() {
	marshal(
		structD{
			anonymousA: anonymousA{
				A: 67,
				B: "A",
			},
			anonymousB: anonymousB{
				A: 78,
				B: "B",
			},
		},
	)

	// Output:
	// {}
}

type anonymousC struct {
	C int
	D string
}

type structE struct {
	anonymousA // 两个结构体中没有相同的字段，不冲突，
	anonymousC // 4 个 inner exported 字段都被 marshal
}

func ExampleMarshal_anonymous_struct_fields_with_inner_exported_fields_no_conflict() {
	marshal(
		structE{
			anonymousA: anonymousA{
				A: 67,
				B: "A",
			},
			anonymousC: anonymousC{
				C: 78,
				D: "B",
			},
		},
	)

	// Output:
	// {"A":67,"B":"A","C":78,"D":"B"}
}

type anonymousD struct {
	A int `json:"A"`
	B string
}

type anonymousE struct {
	A int
	B string `json:"B"`
}

type structF struct {
	anonymousD // 两个匿名字段的 A 字段 JSON name 冲突，但是 anonymousD 的 A 字段有 JSON tag，
	anonymousE // 所以它被 marshal，同理，anonymousE 的 B 字段也被 marshal
}

func ExampleMarshal_anonymous_struct_fields_with_inner_exported_fields_conflict_with_json_name() {
	marshal(
		structF{
			anonymousD: anonymousD{
				A: 78,
				B: "D",
			},
			anonymousE: anonymousE{
				A: 89,
				B: "E",
			},
		},
	)

	// Output:
	// {"A":78,"B":"E"}
}

type structG struct {
	anonymousA     // anonymousA 的两个字段 A 和 B 被最外层的 A 和 B 遮盖了，
	A          int // 因此没有被 marshal
	B          string
}

func ExampleMarshal_visibility_rules_for_struct_fields() {
	marshal(structG{
		anonymousA: anonymousA{
			A: 33,
			B: "Perl",
		},
		A: 45,
		B: "Go",
	})

	// Output:
	// {"A":45,"B":"Go"}
}

type anonymousF struct {
	A int    `json:"A"`
	B string `json:"B"`
}

type structH struct {
	anonymousF        // anonymousF 的两个字段 A 和 B 被最外层的 A 和 B 遮盖了
	A          int    `json:"A"`
	B          string `json:"B"`
}

func ExampleMarshal_visibility_rules_for_struct_fields_with_json_tags() {
	marshal(structH{
		anonymousF: anonymousF{
			A: 67,
			B: "Perl",
		},
		A: 55,
		B: "Go",
	})

	// Output:
	// {"A":55,"B":"Go"}
}

type structJ struct {
	anonymousF     // anonymousF 的两个字段 A 和 B 被最外层的 A 和 B 遮盖了，
	A          int // 即使它们有 json tag 也没用。
	B          string
}

func ExampleMarshal_visibility_rules_for_struct_fields_with_json_tags_2() {
	marshal(structJ{
		anonymousF: anonymousF{
			A: 13,
			B: "Go",
		},
		A: 55,
		B: "Python",
	})

	// Output:
	// {"A":55,"B":"Python"}
}
