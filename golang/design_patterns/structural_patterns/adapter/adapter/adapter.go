package adapter

type V5 interface {
	V5Charge()
}

type V220 struct{} // 或许接口也可以?

func NewV220() *V220 {
	return &V220{}
}

func (v220 *V220) V220Charge() {}

type V5Adapter struct {
	v220 *V220
}

func NewV5Adapter(v220 *V220) *V5Adapter {
	return &V5Adapter{
		v220: v220,
	}
}

func (v *V5Adapter) V5Charge() {
	v.v220.V220Charge()
}
