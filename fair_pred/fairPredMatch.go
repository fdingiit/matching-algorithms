package fair_pred

import (
	"fmt"
	"math"

	"github.com/fdingiit/matching-algorithms/def"
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

type PredicateValue interface {
	String() string
}

type Predicate struct {
	Attribute string
	Operator  string
	Value     PredicateValue
}

func (pred Predicate) String() string {
	return fmt.Sprintf("pred=%s&%s&%s", pred.Attribute, pred.Operator, pred.Value.String())
}

type PredicateMatcher struct {
	predsEqualStrs        map[string]Predicate
	predsLessEqualStrs    map[string]Predicate
	predsGreaterEqualStrs map[string]Predicate

	predToSub map[Predicate]map[def.Subscription]struct{}
}

func NewPredicateRecorder() PredicateMatcher {
	return PredicateMatcher{
		predsEqualStrs: map[string]Predicate{},
		predToSub:      map[Predicate]map[def.Subscription]struct{}{},
	}
}

func (pm PredicateMatcher) matchFruit(fruit def.Fruit) *Predicate {
	v, ok := pm.predsEqualStrs[fruit.String()]
	if !ok {
		return nil
	}
	return &v
}

func (pm PredicateMatcher) matchColor(color def.Color) *Predicate {
	v, ok := pm.predsEqualStrs[color.String()]
	if !ok {
		return nil
	}
	return &v
}

func (pm PredicateMatcher) matchCity(city def.City) *Predicate {
	v, ok := pm.predsEqualStrs[city.String()]
	if !ok {
		return nil
	}
	return &v
}

func (pm PredicateMatcher) matchLessEqualWeight(weight def.Weight) []*Predicate {
	var ret []*Predicate

	for _, pred := range pm.predsLessEqualStrs {
		w := pred.Value.(def.Weight)
		if uint(weight) <= uint(w) {
			ret = append(ret, &pred)
		}
	}

	return ret
}

func (pm PredicateMatcher) matchGreaterEqualWeight(weight def.Weight) []*Predicate {
	var ret []*Predicate

	for _, pred := range pm.predsGreaterEqualStrs {
		w := pred.Value.(def.Weight)
		if uint(weight) >= uint(w) {
			ret = append(ret, &pred)
		}
	}

	return ret
}

func (pm PredicateMatcher) satisfiedPred(product def.Product) []Predicate {
	var preds []Predicate

	if pred := pm.matchFruit(product.Fruit); pred != nil {
		preds = append(preds, *pred)
	}
	if pred := pm.matchColor(product.Color); pred != nil {
		preds = append(preds, *pred)
	}
	if pred := pm.matchCity(product.City); pred != nil {
		preds = append(preds, *pred)
	}

	for _, pred := range pm.matchGreaterEqualWeight(product.Weight) {
		preds = append(preds, *pred)
	}

	for _, pred := range pm.matchLessEqualWeight(product.Weight) {
		preds = append(preds, *pred)
	}

	return preds
}

func (pm PredicateMatcher) Match(product def.Product) []def.Subscription {
	var matched []def.Subscription

	preds := pm.satisfiedPred(product)
	
}

func (pm PredicateMatcher) add(pred Predicate, sub def.Subscription) PredicateMatcher {
	str := pred.String()
	if _, ok := pm.predsEqualStrs[str]; !ok {
		if pred.Operator == opEqual {
			pm.predsEqualStrs[str] = pred
		} else {
			if pred.Operator == opLessEqual {
				pm.predsLessEqualStrs[str] = pred
			} else {
				pm.predsGreaterEqualStrs[str] = pred
			}
		}
		pm.predToSub[pred] = map[def.Subscription]struct{}{}
	}

	pm.predToSub[pred][sub] = struct{}{}
	return pm
}

var (
	predApple = Predicate{
		Attribute: attrFruit,
		Operator:  opEqual,
		Value:     def.Apple,
	}

	predGrape = Predicate{
		Attribute: attrFruit,
		Operator:  opEqual,
		Value:     def.Grape,
	}

	predWatermelon = Predicate{
		Attribute: attrFruit,
		Operator:  opEqual,
		Value:     def.Watermelon,
	}

	predRed = Predicate{
		Attribute: attrColor,
		Operator:  opEqual,
		Value:     def.Red,
	}

	predGreen = Predicate{
		Attribute: attrColor,
		Operator:  opEqual,
		Value:     def.Green,
	}

	predPurple = Predicate{
		Attribute: attrColor,
		Operator:  opEqual,
		Value:     def.Purple,
	}

	predBeijing = Predicate{
		Attribute: attrCity,
		Operator:  opEqual,
		Value:     def.Beijing,
	}

	predShanghai = Predicate{
		Attribute: attrCity,
		Operator:  opEqual,
		Value:     def.Shanghai,
	}

	predGuangzhou = Predicate{
		Attribute: attrCity,
		Operator:  opEqual,
		Value:     def.Guangzhou,
	}
)

func (pm PredicateMatcher) Add(sub def.Subscription) {
	if sub.Fruits == nil {
		pm.add(predApple, sub).add(predGrape, sub).add(predWatermelon, sub)
	} else {
		for fruit := range sub.Fruits.ToSlice() {
			switch def.Fruit(fruit) {
			case def.Apple:
				pm.add(predApple, sub)

			case def.Grape:
				pm.add(predGrape, sub)

			case def.Watermelon:
				pm.add(predWatermelon, sub)
			}
		}
	}

	if sub.Colors == nil {
		pm.add(predRed, sub).add(predPurple, sub).add(predGreen, sub)
	} else {
		for color := range sub.Colors.ToSlice() {
			switch def.Color(color) {
			case def.Red:
				pm.add(predRed, sub)

			case def.Green:
				pm.add(predGreen, sub)

			case def.Purple:
				pm.add(predPurple, sub)
			}
		}
	}

	if sub.Cities == nil {
		pm.add(predBeijing, sub).add(predShanghai, sub).add(predGuangzhou, sub)
	} else {
		for city := range sub.Cities.ToSlice() {
			switch def.City(city) {
			case def.Beijing:
				pm.add(predBeijing, sub)

			case def.Shanghai:
				pm.add(predShanghai, sub)

			case def.Guangzhou:
				pm.add(predGuangzhou, sub)
			}
		}
	}

	if sub.WeightBottom == nil {
		pm.add(Predicate{
			Attribute: attrWeight,
			Operator:  opGreaterEqual,
			Value:     def.Weight(0),
		}, sub)
	} else {
		pm.add(Predicate{
			Attribute: attrWeight,
			Operator:  opGreaterEqual,
			Value:     *sub.WeightBottom,
		}, sub)
	}

	if sub.WeightFloor == nil {
		pm.add(Predicate{
			Attribute: attrWeight,
			Operator:  opLessEqual,
			Value:     def.Weight(math.MaxUint32),
		}, sub)
	} else {
		pm.add(Predicate{
			Attribute: attrWeight,
			Operator:  opLessEqual,
			Value:     *sub.WeightFloor,
		}, sub)
	}

}
