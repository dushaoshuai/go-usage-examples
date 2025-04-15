package abstractfactory

type FurnitureFactory interface {
	CreateChair() Chair
	CreateCoffeeTable() CoffeeTable
	CreateSofa() Sofa
}

type artDecoFactory struct{}

func NewArtDecoFactory() FurnitureFactory { return artDecoFactory{} }

func (a artDecoFactory) CreateChair() Chair             { return artDecoChair{} }
func (a artDecoFactory) CreateCoffeeTable() CoffeeTable { return artDecoCoffeeTable{} }
func (a artDecoFactory) CreateSofa() Sofa               { return artDecoSofa{} }

type victorianFactory struct{}

func NewVictorianFactory() FurnitureFactory { return victorianFactory{} }

func (v victorianFactory) CreateChair() Chair             { return victorianChair{} }
func (v victorianFactory) CreateCoffeeTable() CoffeeTable { return victorianCoffeeTable{} }
func (v victorianFactory) CreateSofa() Sofa               { return victorianSofa{} }

type modernFactory struct{}

func NewModernFactory() FurnitureFactory { return modernFactory{} }

func (m modernFactory) CreateChair() Chair             { return modernChair{} }
func (m modernFactory) CreateCoffeeTable() CoffeeTable { return modernCoffeeTable{} }
func (m modernFactory) CreateSofa() Sofa               { return modernSofa{} }
