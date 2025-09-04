package auth

import (
	"net/http"
	"testing"
)

func Test_GetAPIKey(t *testing.T) {
	testCases := []struct {
		header        http.Header
		expectedKey   string
		expectedError error
	}{
		{http.Header{"Authorization": []string{"ApiKey testone"}}, "testone", nil},
		{http.Header{"Authorization": []string{"ApiKey testTwo"}}, "testTwo", ErrNoAuthHeaderIncluded},
		{http.Header{"Apthorization": []string{"ApiKey testThree"}}, "", ErrNoAuthHeaderIncluded},
	}
	for _, tc := range testCases {
		expected := tc.expectedKey
		got, err := GetAPIKey(tc.header)
		if err != tc.expectedError {
			t.Fatalf("expected error: %v, got error: %v", tc.expectedError, err)
		}
		if got != expected {
			t.Fatalf("expected: %v, got: %v", expected, got)
		}
	}
}
