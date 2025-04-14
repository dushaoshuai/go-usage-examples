package abstractfactory

type CoffeeTable interface {
	IsCoffeeTable()
}

type artDecoCoffeeTable struct{}

func (artDecoCoffeeTable) IsCoffeeTable() {}

type victorianCoffeeTable struct{}

func (victorianCoffeeTable) IsCoffeeTable() {}

type modernCoffeeTable struct{}

func (modernCoffeeTable) IsCoffeeTable() {}
