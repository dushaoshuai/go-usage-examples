package abstractfactory

type Chair interface {
	IsChair()
}

type artDecoChair struct{}

func (artDecoChair) IsChair() {}

type victorianChair struct{}

func (victorianChair) IsChair() {}

type modernChair struct{}

func (modernChair) IsChair() {}
