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
			name: "with slow & fast urls",
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

func TestRacerWithSelect(t *testing.T) {
	serverA := makeDelayedServer(20 * time.Millisecond)
	serverB := makeDelayedServer(0 * time.Millisecond)

	defer serverA.Close()
	defer serverB.Close()

	type args struct {
		slowURL string
		fastURL string
		timeout time.Duration
	}
	tests := []struct {
		name       string
		args       args
		wantWinner string
	}{
		{
			name: "",
			args: args{
				slowURL: serverA.URL,
				fastURL: serverB.URL,
				timeout: 10 * time.Second,
			},
			wantWinner: serverB.URL,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWinner, err := RacerWithSelect(tt.args.slowURL, tt.args.fastURL)

			if err != nil {
				t.Fatalf("did not expect an error but got one %v", err)
			}

			if gotWinner != tt.wantWinner {
				t.Errorf("got %q, want %q", gotWinner, tt.wantWinner)
			}
		})
	}
}

func TestConfigurableRacer(t *testing.T) {
	server := makeDelayedServer(25 * time.Millisecond)

	defer server.Close()

	type args struct {
		aURL    string
		bURL    string
		timeout time.Duration
	}
	tests := []struct {
		name       string
		args       args
		wantWinner string
		wantErr    bool
	}{
		{
			name: "test configurable racer",
			args: args{
				aURL:    server.URL,
				bURL:    server.URL,
				timeout: 20 * time.Millisecond,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ConfigurableRacer(tt.args.aURL, tt.args.bURL, tt.args.timeout)

			if err == nil {
				t.Error("expected an error but didn't get one")
			}
		})
	}
}
