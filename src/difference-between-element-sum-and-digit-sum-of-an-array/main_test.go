package difference_between_element_sum_and_digit_sum_of_an_array

import "testing"

func Test_differenceOfSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Return the absolute difference between the element sum and digit sum of nums",
			args: args{[]int{134, 15, 63, 1999}},
			want: 2160,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := differenceOfSum(tt.args.nums); got != tt.want {
				t.Errorf("differenceOfSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
