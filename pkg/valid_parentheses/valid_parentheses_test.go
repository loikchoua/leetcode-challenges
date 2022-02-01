package validparentheses

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Ok",
			args: args{
				s: "()()",
			},
			want: true,
		},
		{
			name: "Invalid",
			args: args{
				s: "(])",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValid(tt.args.s)
			require.Equal(t, tt.want, got)
		})
	}
}
