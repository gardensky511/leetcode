package check_if_a_string_is_an_acronym_of_words

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_isAcronym(t *testing.T) {
	type args struct {
		words []string
		s     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Return true if s is an acronym of words",
			args: args{words: []string{"never", "gonna", "give", "up", "on", "you"}, s: "ngguoy"},
			want: true,
		},
		{
			name: "Return false if s is not an acronym of words",
			args: args{words: []string{"never", "gonna", "give", "up", "on", "you"}, s: "ngguob"},
			want: false,
		},
		{
			name: "Return false if s' length and words' length are not the same",
			args: args{words: []string{"never", "gonna", "give", "up", "on", "you"}, s: "abc"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAcronym(tt.args.words, tt.args.s)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
		})
	}
}
