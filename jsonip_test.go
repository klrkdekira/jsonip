package jsonip

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupServer(payload string) {
	mockHandler := http.NewServeMux()
	mockHandler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(payload))
	})

	mockServer := httptest.NewServer(mockHandler)

	// overwriting the defaults
	ipv4Endpoint = mockServer.URL
	ipv6Endpoint = mockServer.URL
}

func test(t *testing.T, body, expected string) {
	setupServer(body)

	data, err := IPv4()
	if err != nil {
		t.Fatalf("Encountered unexpected error, %v", err)
	}

	if data.IP != expected {
		t.Fatalf("Expected %s, got %s", expected, data.IP)
	}
}

func TestIPv4(t *testing.T) {
	test(t, `{"ip": "127.0.0.1"}`, "127.0.0.1")
}

func TestIPv6(t *testing.T) {
	test(t, `{"ip": "2001:4860:4860:0:0:0:0:8888"}`, "2001:4860:4860:0:0:0:0:8888")
}

// For the sake of example
func BenchmarkIPv4(b *testing.B) {
	setupServer(`{"ip": "127.0.0.1"}`)
	for i := 0; i < b.N; i++ {
		data, err := IPv4()
		if err != nil {
			b.Fatalf("Encountered unexpected error, %v", err)
		}
		_ = data
	}
}
