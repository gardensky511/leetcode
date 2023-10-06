package unique_number_of_occurrences

import "testing"

func Test_uniqueOccurrences(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "return true if the number of occurrences of each value in the array is unique",
			args: args{[]int{1, 2, 2, 1, 1, 3}},
			want: true,
		},
		{
			name: "return false if the number of occurrences of each value in the array is not unique",
			args: args{[]int{1, 2}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueOccurrences(tt.args.arr); got != tt.want {
				t.Errorf("uniqueOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}
