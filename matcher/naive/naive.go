package naive

import (
	"github.com/fdingiit/matching-algorithms/def"
	"github.com/fdingiit/matching-algorithms/matcher"
)

type MatcherNaive struct {
	subs []def.Subscription
}

func NewMatcher() matcher.Matcher {
	return &MatcherNaive{}
}

func (m *MatcherNaive) Add(subs ...def.Subscription) {
	m.subs = append(m.subs, subs...)
}

func (m *MatcherNaive) Match(product def.Product) []def.Subscription {
	var matched []def.Subscription

	for _, sub := range m.subs {
		if match(sub, product) {
			matched = append(matched, sub)
		}
	}

	return matched
}

func match(sub def.Subscription, product def.Product) bool {
	if product.Fruit != sub.Fruit && sub.Fruit != def.FruitAll {
		return false
	}

	if product.City != sub.City && sub.City != def.CityAll {
		return false
	}

	if product.Color != sub.Color && sub.Color != def.ColorAll {
		return false
	}

	if product.Weight < sub.WeightBottom || product.Weight > sub.WeightFloor {
		return false
	}

	return true
}
