package longestcommonprefix

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// {
		// 	name: "Ok",
		// 	args: args{
		// 		strs: []string{"flower", "flow", "flight"},
		// 	},
		// 	want: "fl",
		// },
		{
			name: "Ok ab",
			args: args{
				strs: []string{"ab", "a"},
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestCommonPrefix(tt.args.strs)
			require.Equal(t, tt.want, got)
		})
	}
}
