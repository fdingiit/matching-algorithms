package naive

import (
	"github.com/fdingiit/matching-algorithms/def"
)

func predMatch(sub def.Subscription, product def.Product) bool {
	if sub.Fruits != nil {
		if ok := sub.Fruits.Contains(product.Fruit); !ok {
			return false
		}
	}

	if sub.Colors != nil {
		if ok := sub.Colors.Contains(product.Color); !ok {
			return false
		}
	}

	if sub.Cities != nil {
		if ok := sub.Cities.Contains(product.City); !ok {
			return false
		}
	}

	if sub.WeightBottom != nil && product.Weight < *sub.WeightBottom {
		return false
	}

	if sub.WeightFloor != nil && product.Weight > *sub.WeightFloor {
		return false
	}

	return true
}

func Match(subscriptions []def.Subscription, product def.Product) []def.Subscription {
	var matched []def.Subscription

	for _, sub := range subscriptions {
		if predMatch(sub, product) {
			matched = append(matched, sub)
		}
	}

	return matched
}
