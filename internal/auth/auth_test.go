package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		header    http.Header
		expect    string
		expectErr error
	}{
		{header: http.Header{"Authorization": []string{"ApiKey abc"}}, expect: "abc", expectErr: nil},
		{header: http.Header{"Authorization": []string{"badApiKey"}}, expect: "", expectErr: ErrMalformedAuthHeader},
		{header: http.Header{}, expect: "", expectErr: ErrNoAuthHeaderIncluded},
	}

	for _, tc := range tests {
		got, gotErr := GetAPIKey(tc.header)
		if !reflect.DeepEqual(got, tc.expect) || !errors.Is(gotErr, tc.expectErr) {
			t.Fatalf("expected: %v, %s, got: %v, %s", tc.expect, tc.expectErr, got, gotErr)
		}
	}
}
