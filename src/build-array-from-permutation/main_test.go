package build_array_from_permutation

import (
	"reflect"
	"testing"
)

func Test_buildArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[0,2,1,5,3,4] → [0,1,2,4,5,3]",
			args: args{[]int{0, 2, 1, 5, 3, 4}},
			want: []int{0, 1, 2, 4, 5, 3},
		},
		{
			name: "[5,0,1,2,3,4] → [4,5,0,1,2,3]",
			args: args{[]int{5, 0, 1, 2, 3, 4}},
			want: []int{4, 5, 0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildArray1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildArray() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := buildArray2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
