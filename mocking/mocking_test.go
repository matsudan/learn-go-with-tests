package mocking

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "countdown",
			want: `3
2
1
Go!`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			spySleeper := &SpySleeper{}

			Countdown(buffer, spySleeper)
			got := buffer.String()

			if got != tt.want {
				t.Errorf("got %q want %q", got, tt.want)
			}
		})
	}
}
