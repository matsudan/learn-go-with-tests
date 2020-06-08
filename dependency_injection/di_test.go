package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "withName", args: args{name: "Chris"}, want: "Hello, Chris"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer := bytes.Buffer{}
			Greet(&buffer, tt.args.name)
			got := buffer.String()

			if got != tt.want {
				t.Errorf("got %q want %q", got, tt.want)
			}
		})
	}
}
