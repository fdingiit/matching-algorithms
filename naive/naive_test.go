package naive

import (
	"reflect"
	"testing"

	mapset "github.com/deckarep/golang-set"
	uuid "github.com/satori/go.uuid"

	"github.com/fdingiit/matching-algorithms/def"
)

func TestNaiveMatch(t *testing.T) {
	type args struct {
		subscriptions []def.Subscription
		product       def.Product
	}

	subAll := def.Subscription{
		Id: uuid.NewV4().String(),
	}

	subApple := def.Subscription{
		Id:     uuid.NewV4().String(),
		Fruits: mapset.NewSetFromSlice([]interface{}{def.Apple}),
	}

	subMelo := def.Subscription{
		Id:     uuid.NewV4().String(),
		Fruits: mapset.NewSetFromSlice([]interface{}{def.Watermelon}),
	}

	tests := []struct {
		name string
		args args
		want []def.Subscription
	}{
		{
			name: "test-case-001",
			args: struct {
				subscriptions []def.Subscription
				product       def.Product
			}{
				subscriptions: []def.Subscription{
					// sub all
					subAll,
					subApple,
					subMelo,
				},
				product: def.Product{
					Fruit:  def.Apple,
					Color:  def.Green,
					City:   def.Beijing,
					Weight: 100,
				}},
			want: []def.Subscription{
				subAll,
				subApple,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Match(tt.args.subscriptions, tt.args.product); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
