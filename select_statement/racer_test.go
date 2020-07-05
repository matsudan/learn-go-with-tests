package select_statement

import "testing"

func TestRacer(t *testing.T) {
	type args struct {
		slow string
		fast string
	}
	tests := []struct {
		name       string
		args       args
		wantWinner string
	}{
		{
			name: "with two args",
			args: args{
				"http://www.facebook.com",
				"http://www.quii.co.uk",
			},
			wantWinner: "http://www.quii.co.uk",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotWinner := Racer(tt.args.slow, tt.args.fast); gotWinner != tt.wantWinner {
				t.Errorf("Racer() = %v, want %v", gotWinner, tt.wantWinner)
			}
		})
	}
}