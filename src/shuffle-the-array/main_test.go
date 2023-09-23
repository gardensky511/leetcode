package shuffle_the_array

import (
	"reflect"
	"testing"
)

func Test_shuffle(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1, 2, 3, 4, 4, 3, 2, 1] → [1, 4, 2, 3, 3, 2, 4, 1]",
			args: args{n: 4, nums: []int{1, 2, 3, 4, 4, 3, 2, 1}},
			want: []int{1, 4, 2, 3, 3, 2, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shuffle(tt.args.nums, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}
