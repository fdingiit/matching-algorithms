package matcher

import "github.com/fdingiit/matching-algorithms/def"

type Matcher interface {
	Add(subs ...def.Subscription)

	Match(product def.Product) []def.Subscription
}
