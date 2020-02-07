package test

import (
	uuid "github.com/satori/go.uuid"

	"github.com/fdingiit/matching-algorithms/def"
)

type Args struct {
	Subscriptions []def.Subscription
	Product       def.Product
}

type Case struct {
	Name string
	Args
	Wanted []def.Subscription
}

var (
	subAll = def.Subscription{
		Id:           uuid.NewV4().String(),
		Fruit:        def.FruitAll,
		City:         def.CityAll,
		Color:        def.ColorAll,
		WeightBottom: 1000,
		WeightFloor:  2000,
	}

	subApple = def.Subscription{
		Id:           uuid.NewV4().String(),
		Fruit:        def.Apple,
		City:         def.CityAll,
		Color:        def.ColorAll,
		WeightBottom: def.WeightMin,
		WeightFloor:  def.WeightMax,
	}

	subMelo = def.Subscription{
		Id:           uuid.NewV4().String(),
		Fruit:        def.Watermelon,
		City:         def.CityAll,
		Color:        def.ColorAll,
		WeightBottom: def.WeightMin,
		WeightFloor:  def.WeightMax,
	}
)

var BasicCases = []Case{
	{
		Name: "test-case-001",
		Args: Args{
			Subscriptions: []def.Subscription{
				subAll,
				subApple,
				subMelo,
			},
			Product: def.Product{
				Fruit:  def.Apple,
				Color:  def.Green,
				City:   def.Beijing,
				Weight: 100,
			}},
		Wanted: []def.Subscription{
			subApple,
		},
	},
}
