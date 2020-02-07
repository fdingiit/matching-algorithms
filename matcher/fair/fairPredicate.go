package fair

import (
	"fmt"

	"github.com/fdingiit/matching-algorithms/def"
	"github.com/fdingiit/matching-algorithms/matcher"
)

const (
	opLessEqual    = "lessEqual"
	opEqual        = "equal"
	opGreaterEqual = "greaterEqual"
)

const (
	attrFruit  = "fruit"
	attrColor  = "color"
	attrCity   = "city"
	attrWeight = "weight"
)

type predicateValue interface {
	String() string
}

type predicate struct {
	Attribute string
	Operator  string
	Value     predicateValue
}

var (
	predAllFruit = predicate{
		Attribute: attrFruit,
		Operator:  opEqual,
		Value:     def.FruitAll,
	}

	predApple = predicate{
		Attribute: attrFruit,
		Operator:  opEqual,
		Value:     def.Apple,
	}

	predGrape = predicate{
		Attribute: attrFruit,
		Operator:  opEqual,
		Value:     def.Grape,
	}

	predWatermelon = predicate{
		Attribute: attrFruit,
		Operator:  opEqual,
		Value:     def.Watermelon,
	}

	predAllColor = predicate{
		Attribute: attrColor,
		Operator:  opEqual,
		Value:     def.ColorAll,
	}

	predRed = predicate{
		Attribute: attrColor,
		Operator:  opEqual,
		Value:     def.Red,
	}

	predGreen = predicate{
		Attribute: attrColor,
		Operator:  opEqual,
		Value:     def.Green,
	}

	predPurple = predicate{
		Attribute: attrColor,
		Operator:  opEqual,
		Value:     def.Purple,
	}

	predAllCity = predicate{
		Attribute: attrCity,
		Operator:  opEqual,
		Value:     def.CityAll,
	}

	predBeijing = predicate{
		Attribute: attrCity,
		Operator:  opEqual,
		Value:     def.Beijing,
	}

	predShanghai = predicate{
		Attribute: attrCity,
		Operator:  opEqual,
		Value:     def.Shanghai,
	}

	predGuangzhou = predicate{
		Attribute: attrCity,
		Operator:  opEqual,
		Value:     def.Guangzhou,
	}
)

func (pred predicate) String() string {
	return fmt.Sprintf("pred=%s&%s&%s", pred.Attribute, pred.Operator, pred.Value.String())
}

type MatcherFair struct {
	predsEqual        map[string]predicate
	predsLessEqual    map[string]predicate
	predsGreaterEqual map[string]predicate

	predToSub    map[predicate]map[def.Subscription]struct{}
	subNeedPreds int
}

func NewMatcher() matcher.Matcher {
	return &MatcherFair{
		predsEqual:        map[string]predicate{},
		predsLessEqual:    map[string]predicate{},
		predsGreaterEqual: map[string]predicate{},
		predToSub:         map[predicate]map[def.Subscription]struct{}{},
	}
}

func fruitToPred(fruit def.Fruit) predicate {
	switch fruit {
	case def.FruitAll:
		return predAllFruit
	case def.Apple:
		return predApple
	case def.Grape:
		return predGrape
	case def.Watermelon:
		return predWatermelon
	default:
		panic(fmt.Sprintf("no such fruit, %+v", fruit))
	}
}

func colorToPred(color def.Color) predicate {
	switch color {
	case def.ColorAll:
		return predAllColor
	case def.Red:
		return predRed
	case def.Green:
		return predGreen
	case def.Purple:
		return predPurple
	default:
		panic(fmt.Sprintf("no such color, %+v", color))
	}
}

func cityToPred(city def.City) predicate {
	switch city {
	case def.CityAll:
		return predAllCity
	case def.Shanghai:
		return predShanghai
	case def.Beijing:
		return predBeijing
	case def.Guangzhou:
		return predGuangzhou
	default:
		panic(fmt.Sprintf("no such city, %+v", city))
	}
}

func (m *MatcherFair) matchFruit(fruit def.Fruit) []predicate {
	var ret []predicate

	pred := fruitToPred(fruit)
	if _, ok := m.predsEqual[pred.String()]; ok {
		ret = append(ret, pred)
	}
	if fruit != def.FruitAll {
		pred := fruitToPred(def.FruitAll)
		if _, ok := m.predsEqual[pred.String()]; ok {
			ret = append(ret, pred)
		}
	}

	return ret
}

func (m *MatcherFair) matchColor(color def.Color) []predicate {
	var ret []predicate

	pred := colorToPred(color)
	if _, ok := m.predsEqual[pred.String()]; ok {
		ret = append(ret, pred)
	}
	if color != def.ColorAll {
		pred := colorToPred(def.ColorAll)
		if _, ok := m.predsEqual[pred.String()]; ok {
			ret = append(ret, pred)
		}
	}

	return ret
}

func (m *MatcherFair) matchCity(city def.City) []predicate {
	var ret []predicate

	pred := cityToPred(city)
	if _, ok := m.predsEqual[pred.String()]; ok {
		ret = append(ret, pred)
	}
	if city != def.CityAll {
		pred := cityToPred(def.CityAll)
		if _, ok := m.predsEqual[pred.String()]; ok {
			ret = append(ret, pred)
		}
	}

	return ret
}

func (m *MatcherFair) matchLessEqualWeight(weight def.Weight) []predicate {
	var ret []predicate

	for _, pred := range m.predsLessEqual {
		w := pred.Value.(def.Weight)
		if uint(weight) <= uint(w) {
			ret = append(ret, pred)
		}
	}

	return ret
}

func (m *MatcherFair) matchGreaterEqualWeight(weight def.Weight) []predicate {
	var ret []predicate

	for _, pred := range m.predsGreaterEqual {
		w := pred.Value.(def.Weight)
		if uint(weight) >= uint(w) {
			ret = append(ret, pred)
		}
	}

	return ret
}

func (m *MatcherFair) satisfiedPred(product def.Product) []predicate {
	var preds []predicate

	for _, pred := range m.matchFruit(product.Fruit) {
		preds = append(preds, pred)
	}
	for _, pred := range m.matchColor(product.Color) {
		preds = append(preds, pred)
	}
	for _, pred := range m.matchCity(product.City) {
		preds = append(preds, pred)
	}
	for _, pred := range m.matchGreaterEqualWeight(product.Weight) {
		preds = append(preds, pred)
	}
	for _, pred := range m.matchLessEqualWeight(product.Weight) {
		preds = append(preds, pred)
	}

	return preds
}

func (m *MatcherFair) Match(product def.Product) []def.Subscription {
	var matched []def.Subscription
	var hitCount = map[def.Subscription]int{}

	preds := m.satisfiedPred(product)

	for _, pred := range preds {
		subs := m.predToSub[pred]
		for sub := range subs {
			hitCount[sub] += 1
		}
	}

	for sub, cnt := range hitCount {
		if cnt == m.subNeedPreds {
			matched = append(matched, sub)
		}
	}

	return matched
}

func (m *MatcherFair) add(pred predicate, sub def.Subscription) {
	str := pred.String()

	switch pred.Operator {
	case opEqual:
		if _, ok := m.predsEqual[str]; !ok {
			m.predsEqual[str] = pred
		}

	case opLessEqual:
		if _, ok := m.predsLessEqual[str]; !ok {
			m.predsLessEqual[str] = pred
		}

	case opGreaterEqual:
		if _, ok := m.predsGreaterEqual[str]; !ok {
			m.predsGreaterEqual[str] = pred
		}
	}

	if _, ok := m.predToSub[pred]; !ok {
		m.predToSub[pred] = map[def.Subscription]struct{}{}
	}
	m.predToSub[pred][sub] = struct{}{}
}

func (m *MatcherFair) Add(subs ...def.Subscription) {
	for _, sub := range subs {
		m.add(fruitToPred(sub.Fruit), sub)
		m.add(cityToPred(sub.City), sub)
		m.add(colorToPred(sub.Color), sub)
		m.add(predicate{
			Attribute: attrWeight,
			Operator:  opGreaterEqual,
			Value:     sub.WeightBottom,
		}, sub)
		m.add(predicate{
			Attribute: attrWeight,
			Operator:  opLessEqual,
			Value:     sub.WeightFloor,
		}, sub)
		m.subNeedPreds = 5
	}
}
