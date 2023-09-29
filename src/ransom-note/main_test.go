package ransom_note

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "return true if ransomNote can be constructed by using the letters from magazine and false otherwise",
			args: args{ransomNote: "aa", magazine: "aab"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
