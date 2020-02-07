package test

import (
	"fmt"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set"
	"github.com/stretchr/testify/assert"

	"github.com/fdingiit/matching-algorithms/def"
	"github.com/fdingiit/matching-algorithms/matcher/fair"
	"github.com/fdingiit/matching-algorithms/matcher/naive"
	"github.com/fdingiit/matching-algorithms/random"
)

var subscriptions []def.Subscription
var product def.Product

func log(format string, a ...interface{}) {
	now := time.Now().Format("2006-01-02 15:04:05.999999999")
	fmt.Println("[" + now + "] " + fmt.Sprintf(format, a...))
}

func init() {
	/*
		log("generating random subscriptions...")
		subscriptions = random.RandSub(100, 4)
		log("random subscriptions generated!")

		product = random.RandProduct()
	*/
}

func subsSet(subs []def.Subscription) mapset.Set {
	var set = mapset.NewSet()

	for _, sub := range subs {
		set.Add(sub)
	}

	return set
}

func TestCrossCheck(t *testing.T) {
	log("generating random subscriptions...")
	subscriptions := random.RandSub(1000, 4)
	log("random subscriptions generated!")

	product := random.RandProduct()

	naiveMatcher := naive.NewMatcher()
	fairMatcher := fair.NewMatcher()

	for _, sub := range subscriptions {
		naiveMatcher.Add(sub)
		fairMatcher.Add(sub)
	}

	naiveSubs := naiveMatcher.Match(product)
	fairSubs := fairMatcher.Match(product)

	naiveSubsSet, fairSubsSet := subsSet(naiveSubs), subsSet(fairSubs)

	assert.Equal(t, naiveSubsSet, fairSubsSet)
}

func TestMatch_Compare(t *testing.T) {
	log("generating random subscriptions...")
	subscriptions := random.RandSub(1000000, 4)
	log("random subscriptions generated!")

	product := random.RandProduct()

	naiveMatcher := naive.NewMatcher()
	fairMatcher := fair.NewMatcher()

	for _, sub := range subscriptions {
		naiveMatcher.Add(sub)
		fairMatcher.Add(sub)
	}

	begin := time.Now()
	naiveSubs := naiveMatcher.Match(product)
	elapsed := time.Since(begin)
	fmt.Println("elapsed:", elapsed)

	begin = time.Now()
	fairSubs := fairMatcher.Match(product)
	elapsed = time.Since(begin)
	fmt.Println("elapsed:", elapsed)

	naiveSubsSet, fairSubsSet := subsSet(naiveSubs), subsSet(fairSubs)
	assert.Equal(t, naiveSubsSet, fairSubsSet)
}
