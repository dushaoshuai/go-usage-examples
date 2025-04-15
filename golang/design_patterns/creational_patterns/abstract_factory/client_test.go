package abstractfactory_test

import (
	abstractfactory "github.com/dushaoshuai/go-usage-examples/golang/design_patterns/creational_patterns/abstract_factory"
)

const (
	ArtDeco   = "ArtDeco"
	Victorian = "Victorian"
	Modern    = "Modern"
)

type config struct {
	style string
}

type App struct {
	factory abstractfactory.FurnitureFactory
}

func newApp(factory abstractfactory.FurnitureFactory) *App {
	return &App{
		factory: factory,
	}
}

func (app *App) createChair() abstractfactory.Chair {
	return app.factory.CreateChair()
}

func Example_client() {
	c := config{
		style: Modern,
	}

	var factory abstractfactory.FurnitureFactory
	switch c.style {
	case ArtDeco:
		factory = abstractfactory.NewArtDecoFactory()
	case Victorian:
		factory = abstractfactory.NewVictorianFactory()
	case Modern:
		factory = abstractfactory.NewModernFactory()
	}

	app := newApp(factory)
	chair := app.createChair()
	chair.IsChair()
}
