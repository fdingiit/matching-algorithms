package def

import (
	"fmt"
)

type Fruit int

const (
	FruitAll Fruit = iota
	Apple
	Grape
	Watermelon
)

var fruitString = map[Fruit]string{
	FruitAll:   "$all",
	Apple:      "$apple",
	Grape:      "$grape",
	Watermelon: "$watermelon",
}

func (f Fruit) String() string {
	return fruitString[f]
}

type Color int

const (
	ColorAll Color = iota
	Red
	Green
	Purple
)

var colorString = map[Color]string{
	ColorAll: "$all",
	Red:      "$red",
	Green:    "$green",
	Purple:   "$purple",
}

func (c Color) String() string {
	return colorString[c]
}

type City int

const (
	CityAll City = iota
	Beijing
	Shanghai
	Guangzhou
)

var cityString = map[City]string{
	CityAll:   "$all",
	Beijing:   "$beijing",
	Shanghai:  "$shanghai",
	Guangzhou: "$guangzhou",
}

func (c City) String() string {
	return cityString[c]
}

type Weight uint

const (
	WeightMin Weight = 0
	WeightMax Weight = 2<<32 - 1
)

func (w Weight) String() string {
	return fmt.Sprintf("$weight=%d", w)
}

type Subscription struct {
	Id string
	Fruit
	Color
	City
	WeightBottom Weight
	WeightFloor  Weight
}

type Product struct {
	Fruit
	Color
	City
	Weight
}
