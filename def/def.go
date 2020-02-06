package def

import (
	"fmt"

	mapset "github.com/deckarep/golang-set"
)

type Fruit int

const (
	Apple Fruit = iota
	Grape
	Watermelon
)

var fruitString = map[Fruit]string{
	Apple:      "$apple",
	Grape:      "$grape",
	Watermelon: "$watermelon",
}

func (f Fruit) String() string {
	return fruitString[f]
}

type Color int

const (
	Red Color = iota
	Green
	Purple
)

var colorString = map[Color]string{
	Red:    "$red",
	Green:  "$green",
	Purple: "$purple",
}

func (c Color) String() string {
	return colorString[c]
}

type City int

const (
	Beijing City = iota
	Shanghai
	Guangzhou
)

var cityString = map[City]string{
	Beijing:   "$beijing",
	Shanghai:  "$shanghai",
	Guangzhou: "$guangzhou",
}

func (c City) String() string {
	return cityString[c]
}

type Weight uint

func (w Weight) String() string {
	return fmt.Sprintf("$weight=%d", w)
}

type Subscription struct {
	Id           string
	Fruits       mapset.Set
	Colors       mapset.Set
	Cities       mapset.Set
	WeightBottom *Weight
	WeightFloor  *Weight
}

type Product struct {
	Fruit
	Color
	City
	Weight
}
