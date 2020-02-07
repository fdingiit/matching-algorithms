package naive

import (
	"reflect"
	"testing"

	"github.com/fdingiit/matching-algorithms/matcher"
	"github.com/fdingiit/matching-algorithms/test"
)

func TestNaiveMatch(t *testing.T) {
	var m matcher.Matcher

	m = NewMatcher()
	for _, sub := range test.BasicCases {
		m.Add(sub.Subscriptions...)
	}

	for _, tt := range test.BasicCases {
		t.Run(tt.Name, func(t *testing.T) {
			if got := m.Match(tt.Args.Product); !reflect.DeepEqual(got, tt.Wanted) {
				t.Errorf("Match() = %+v, want %+v", got, tt.Wanted)
			}
		})
	}
}
