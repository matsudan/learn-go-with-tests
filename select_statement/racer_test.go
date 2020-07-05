package select_statement

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	type args struct {
		slowURL string
		fastURL string
	}
	tests := []struct {
		name       string
		args       args
		wantWinner string
	}{
		{
			name: "with two args",
			args: args{
				slowURL: slowServer.URL,
				fastURL: fastServer.URL,
			},
			wantWinner: fastServer.URL,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotWinner := Racer(tt.args.slowURL, tt.args.fastURL); gotWinner != tt.wantWinner {
				t.Errorf("Racer() = %v, want %v", gotWinner, tt.wantWinner)
			}
		})
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
