package abstractfactory_test

import (
	abstractfactory "github.com/dushaoshuai/go-usage-examples/golang/design_patterns/creational_patterns/abstract_factory"
)

func Example_client() {
	arcDecoFactory := abstractfactory.GetFurnitureFactory(abstractfactory.ArtDeco)
	arcDecoChair := arcDecoFactory.CreateChair()
	arcDecoCoffeeTable := arcDecoFactory.CreateCoffeeTable()
	arcDecoSofa := arcDecoFactory.CreateSofa()
	_ = arcDecoChair
	_ = arcDecoCoffeeTable
	_ = arcDecoSofa

	victorianFactory := abstractfactory.GetFurnitureFactory(abstractfactory.Victorian)
	victorianChair := victorianFactory.CreateChair()
	victorianCoffeeTable := victorianFactory.CreateCoffeeTable()
	victorianSofa := victorianFactory.CreateSofa()
	_ = victorianChair
	_ = victorianCoffeeTable
	_ = victorianSofa

	modernFactory := abstractfactory.GetFurnitureFactory(abstractfactory.Modern)
	modernChair := modernFactory.CreateChair()
	modernCoffeeTable := modernFactory.CreateCoffeeTable()
	modernSofa := modernFactory.CreateSofa()
	_ = modernChair
	_ = modernCoffeeTable
	_ = modernSofa
}
