package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/fdingiit/matching-algorithms/def"
	"github.com/fdingiit/matching-algorithms/naive"
	"github.com/fdingiit/matching-algorithms/random"
)

var subscriptions []def.Subscription
var product def.Product

func log(format string, a ...interface{}) {
	now := time.Now().Format("2006-01-02 15:04:05.999999999")
	fmt.Println("[" + now + "] " + fmt.Sprintf(format, a...))
}

func init() {
	log("generating random subscriptions...")
	subscriptions = random.RandSub(100000, 4)
	log("random subscriptions generated!")

	product = random.RandProduct()
}

func TestNaiveMatch_1M(t *testing.T) {
	begin := time.Now()
	naive.Match(subscriptions, product)
	elapsed := time.Since(begin)
	fmt.Println("elapsed:", elapsed)
}
