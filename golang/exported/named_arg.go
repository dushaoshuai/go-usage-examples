package exported

// https://go.dev/ref/spec#Composite_literals
// https://go.dev/ref/spec#Composite_literals:~:text=For%20struct%20literals%20the%20following%20rules%20apply

// NamedArg sql.NamedArg, sql.Out.
type NamedArg struct {
	_NamedFieldsRequired struct{}

	Name  string
	Value any
}

type NamedArg2 struct {
	NamedFieldsRequired struct{}

	Name  string
	Value any
}

type NamedArg3 struct {
	namedFieldsRequired struct{}

	Name  string
	Value any
}
