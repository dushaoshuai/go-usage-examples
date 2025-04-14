package abstractfactory

type FurnitureFactory interface {
	CreateChair() Chair
	CreateCoffeeTable() CoffeeTable
	CreateSofa() Sofa
}

const (
	ArtDeco   = "ArtDeco"
	Victorian = "Victorian"
	Modern    = "Modern"
)

func GetFurnitureFactory(style string) FurnitureFactory {
	switch style {
	case ArtDeco:
		return artDecoFactory{}
	case Victorian:
		return victorianFactory{}
	case Modern:
		return modernFactory{}
	default:
		return nil
	}
}

type artDecoFactory struct{}

func (a artDecoFactory) CreateChair() Chair             { return artDecoChair{} }
func (a artDecoFactory) CreateCoffeeTable() CoffeeTable { return artDecoCoffeeTable{} }
func (a artDecoFactory) CreateSofa() Sofa               { return artDecoSofa{} }

type victorianFactory struct{}

func (v victorianFactory) CreateChair() Chair             { return victorianChair{} }
func (v victorianFactory) CreateCoffeeTable() CoffeeTable { return victorianCoffeeTable{} }
func (v victorianFactory) CreateSofa() Sofa               { return victorianSofa{} }

type modernFactory struct{}

func (m modernFactory) CreateChair() Chair             { return modernChair{} }
func (m modernFactory) CreateCoffeeTable() CoffeeTable { return modernCoffeeTable{} }
func (m modernFactory) CreateSofa() Sofa               { return modernSofa{} }
