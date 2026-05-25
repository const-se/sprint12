package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_greeting(t *testing.T) {
	t.Parallel()

	type args struct {
		name *string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "with name",
			args: args{
				name: pointer("John"),
			},
			want: "Hello, John!",
		},
		{
			name: "with empty name",
			args: args{
				name: pointer(""),
			},
			want: "Hello!",
		},
		{
			name: "without name",
			args: args{
				name: nil,
			},
			want: "Hello!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := greeting(tt.args.name)
			require.Equal(t, tt.want, got)
		})
	}
}

func pointer[T any](value T) *T {
	return &value
}
