package romantointeger

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRomanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Ok I",
			args: args{
				s: "I",
			},
			want: 1,
		},
		{
			name: "Ok II",
			args: args{
				s: "II",
			},
			want: 2,
		},
		{
			name: "Ok LVIII",
			args: args{
				s: "LVIII",
			},
			want: 58,
		},
		{
			name: "Ok MCMXCIV",
			args: args{
				s: "MCMXCIV",
			},
			want: 1994,
		},
		{
			name: "Ok MCDLXXVI",
			args: args{
				s: "MCDLXXVI",
			},
			want: 1476,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RomanToInt(tt.args.s)
			require.Equal(t, tt.want, got)
		})
	}
}
