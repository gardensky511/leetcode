package concatenation_of_array

import (
	"reflect"
	"testing"
)

func Test_getConcatenationFirstSolution(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "returns concatenation of two nums arrays",
			args: args{
				[]int{1, 2, 3},
			},
			want: []int{1, 2, 3, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getConcatenationFirstSolution(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConcatenationFirstSolution() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := getConcatenationSecondSolution(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConcatenationSecondSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
