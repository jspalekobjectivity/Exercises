package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAreAnagram(t *testing.T) {
	type args struct {
		a string
		b string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true for 'below' and 'elbow'",
			args: args{"below", "elbow"},
			want: true,
		},
		{
			name: "should return false for 'stereo' and 'rodeos'",
			args: args{"stereo", "rodeos"},
			want: false,
		},
		{
			name: "should return false for 'Tatooine' and 'Mustafar'",
			args: args{"Tatooine", "Mustafar"},
			want: false,
		},
		{
			name: "should return false for '' and ' '",
			args: args{"", " "},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := AreAnagram(test.args.a, test.args.b)
			assert.Equal(t, got, test.want)
		})
	}
}
