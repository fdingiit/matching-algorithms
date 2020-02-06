package random

import (
	"math/rand"
	"time"

	mapset "github.com/deckarep/golang-set"
	uuid "github.com/satori/go.uuid"

	"github.com/fdingiit/matching-algorithms/def"
)

var fruits = []interface{}{
	nil,
	def.Apple,
	def.Grape,
	def.Watermelon,
}

var colors = []interface{}{
	nil,
	def.Red,
	def.Green,
	def.Purple,
}

var cities = []interface{}{
	nil,
	def.Beijing,
	def.Shanghai,
	def.Guangzhou,
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func choiceOneNotNil(items []interface{}) interface{} {
	for {
		rand.Seed(time.Now().UnixNano())

		item := items[rand.Int()%len(items)]
		if item != nil {
			return item
		}
	}
}

func choice(items []interface{}) mapset.Set {
	rand.Seed(time.Now().UnixNano())

	cnt := rand.Int() % len(items)
	if cnt == 0 {
		return nil
	}

	var set = mapset.NewSet()

	for set.Cardinality() != cnt {
		rand.Seed(time.Now().UnixNano())

		item := items[rand.Int()%len(items)]
		set.Add(item)
	}

	return set
}

func randWeight(n uint) *def.Weight {
	rand.Seed(time.Now().UnixNano())

	for {
		try := uint(rand.Uint32())
		if try > n {
			var w def.Weight
			w = def.Weight(try)
			return &w
		}
	}
}

func randSub() def.Subscription {
	var sub def.Subscription

	sub.Id = uuid.NewV4().String()
	sub.Fruits = choice(fruits)
	sub.Colors = choice(colors)
	sub.Cities = choice(cities)
	sub.WeightBottom = randWeight(0)
	sub.WeightFloor = randWeight(uint(*sub.WeightBottom))

	return sub
}

func RandSub(n int, pall int) []def.Subscription {
	var subs []def.Subscription
	var inC = make(chan interface{}, 1024)
	var outC = make(chan def.Subscription, 1024)

	pall = min(n, pall)

	for i := 0; i < pall; i++ {
		go func() {
			for range inC {
				outC <- randSub()
			}
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			inC <- struct{}{}
		}
	}()

	for sub := range outC {
		subs = append(subs, sub)
		if len(subs) == n {
			break
		}
	}

	return subs
}

func RandProduct() def.Product {
	var product def.Product

	product.Fruit = choiceOneNotNil(fruits).(def.Fruit)
	product.Color = choiceOneNotNil(colors).(def.Color)
	product.City = choiceOneNotNil(cities).(def.City)
	product.Weight = *randWeight(0)

	return product
}
