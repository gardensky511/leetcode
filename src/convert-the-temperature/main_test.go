package convert_the_temperature

import (
	"reflect"
	"testing"
)

func Test_convertTemperature(t *testing.T) {
	type args struct {
		celsius float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "return Kelvin and Fahrenheit",
			args: args{36.50},
			want: []float64{309.65000, 97.70000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertTemperature(tt.args.celsius); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertTemperature() = %v, want %v", got, tt.want)
			}
		})
	}
}
