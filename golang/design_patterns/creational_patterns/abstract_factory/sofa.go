package abstractfactory

type Sofa interface {
	IsSofa()
}

type artDecoSofa struct{}

func (artDecoSofa) IsSofa() {}

type victorianSofa struct{}

func (victorianSofa) IsSofa() {}

type modernSofa struct{}

func (modernSofa) IsSofa() {}
