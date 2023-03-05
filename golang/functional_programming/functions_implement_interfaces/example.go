package functions_implement_interfaces

type Iface interface {
	Do()
}

// IImplementIfaceMyself references http.HandlerFunc.
// The IImplementIfaceMyself type is an adapter to allow the use of ordinary functions as Iface.
// If f is a function with the signature func(), IImplementIfaceMyself(f) is an Iface that calls f.
type IImplementIfaceMyself func()

func (f IImplementIfaceMyself) Do() {
	f()
}

// iImplementIfaceViaStruct references grpc.funcServerOption,
// https://github.com/grpc/grpc-go/blob/dba26e15a07f43875ccf806a2dd6cbcbc1c12eab/server.go#L206.
// iImplementIfaceViaStruct wraps a function with the signature func() into an
// implementation of the Iface interface.
type iImplementIfaceViaStruct struct {
	f func()
}

func (i *iImplementIfaceViaStruct) Do() {
	i.f()
}

func newIImplementIfaceViaStruct(f func()) Iface {
	return &iImplementIfaceViaStruct{
		f: f,
	}
}
