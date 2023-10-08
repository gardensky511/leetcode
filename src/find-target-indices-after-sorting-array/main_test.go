package find_target_indices_after_sorting_array

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_targetIndices(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Return a list of the target indices of nums after sorting nums in non-decreasing order",
			args: args{[]int{1, 2, 5, 2, 3}, 2},
			want: []int{1, 2},
		},
		{
			name: "If there are no target indices, return an empty list",
			args: args{[]int{1, 2, 5, 2, 3}, 7},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := targetIndices(tt.args.nums, tt.args.target)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
		})
	}
}
