package base7

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvertToBase7(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Ok positive",
			args: args{
				num: 100,
			},
			want: "202",
		},
		{
			name: "Ok negative",
			args: args{
				num: -7,
			},
			want: "-10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertToBase7(tt.args.num)
			require.Equal(t, tt.want, got)
		})
	}
}
