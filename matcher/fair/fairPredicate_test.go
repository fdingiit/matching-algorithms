package fair

import (
	"reflect"
	"testing"

	mapset "github.com/deckarep/golang-set"

	"github.com/fdingiit/matching-algorithms/test"
)

func TestPredicateMatcher_Match(t *testing.T) {
	matcher := NewMatcher()

	for _, bc := range test.BasicCases {
		matcher.Add(bc.Args.Subscriptions...)
	}

	for _, tt := range test.BasicCases {
		t.Run(tt.Name, func(t *testing.T) {
			var gotSet, wantedSet = mapset.NewSet(), mapset.NewSet()

			for _, got := range matcher.Match(tt.Args.Product) {
				gotSet.Add(got)
			}

			for _, wanted := range tt.Wanted {
				wantedSet.Add(wanted)
			}

			if !reflect.DeepEqual(gotSet, wantedSet) {
				t.Errorf("Match() = %v, want %v", gotSet, wantedSet)
			}
		})
	}
}
