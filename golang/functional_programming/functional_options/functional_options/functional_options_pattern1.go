package functional_options

// https://www.sohamkamani.com/golang/options-pattern/

type House struct {
	material     string
	hasFireplace bool
	floors       int
}

// define the constructor
func NewHouse(opts ...HouseOption) *House {
	h := &House{
		material:     "wood",
		hasFireplace: true,
		floors:       2,
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

// define the functional options
type HouseOption func(*House)

func WithConcrete() HouseOption {
	return func(house *House) {
		house.material = "concrete"
	}
}

func WithoutFireplace() HouseOption {
	return func(house *House) {
		house.hasFireplace = false
	}
}

func WithFloors(floors int) HouseOption {
	return func(house *House) {
		house.floors = floors
	}
}
