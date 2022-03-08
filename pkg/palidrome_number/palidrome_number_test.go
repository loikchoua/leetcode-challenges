package palidromenumber

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPalidrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Palidrome",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: "Palidrome",
			args: args{
				x: 5845485,
			},
			want: true,
		},
		{
			name: "Not palidrome",
			args: args{
				x: -121,
			},
			want: false,
		},
		{
			name: "Not palidrome",
			args: args{
				x: 102,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPalidrome(tt.args.x)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestIsPalidromeOptimized(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Palidrome",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: "Palidrome",
			args: args{
				x: 5845485,
			},
			want: true,
		},
		{
			name: "Not palidrome",
			args: args{
				x: -121,
			},
			want: false,
		},
		{
			name: "Not palidrome",
			args: args{
				x: 102,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPalidromeOptimized(tt.args.x)
			require.Equal(t, tt.want, got)
		})
	}
}
