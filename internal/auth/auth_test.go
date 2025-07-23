package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	authHeader := http.Header{
		"Authorization": []string{"ApiKey 12345"},
		"Content-Type":  []string{"text/html"},
	}

	got, err := GetAPIKey(authHeader)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	want := "12345"

	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}

	headers := http.Header{}

	_, err = GetAPIKey(headers)
	if err == nil {
		t.Fatalf("expected an error when authorization header is missing")
	}

	authHeader = http.Header{
		"Authorization": []string{"12345"},
	}

	_, err = GetAPIKey(authHeader)
	if err == nil {
		t.Fatal("expected an error when authorization header len < 2 or does not start with 'ApiKey'")
	}

}
