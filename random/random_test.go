package random

import (
	"testing"

	"github.com/fdingiit/matching-algorithms"
)

func BenchmarkRandSub_P2(b *testing.B) {
	main.log("generating random subscriptions...")
	main.subscriptions = RandSub(100000, 2)
	main.log("random subscriptions generated!")
}

func BenchmarkRandSub_P4(b *testing.B) {
	main.log("generating random subscriptions...")
	main.subscriptions = RandSub(100000, 4)
	main.log("random subscriptions generated!")
}

func BenchmarkRandSub_P8(b *testing.B) {
	main.log("generating random subscriptions...")
	main.subscriptions = RandSub(100000, 8)
	main.log("random subscriptions generated!")
}

func BenchmarkRandSub_P16(b *testing.B) {
	main.log("generating random subscriptions...")
	main.subscriptions = RandSub(100000, 16)
	main.log("random subscriptions generated!")
}

func BenchmarkRandSub_P32(b *testing.B) {
	main.log("generating random subscriptions...")
	main.subscriptions = RandSub(100000, 32)
	main.log("random subscriptions generated!")
}