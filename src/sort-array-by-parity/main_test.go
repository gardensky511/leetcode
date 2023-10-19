package sort_array_by_parity

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_sortArrayByParity(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "move all the even integers at the beginning of the array followed by all the odd integers",
			args: args{[]int{3, 1, 2, 4}},
			want: []int{4, 2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortArrayByParity(tt.args.nums)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
		})
	}
}
