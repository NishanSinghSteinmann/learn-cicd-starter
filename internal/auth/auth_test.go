package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		expectedErr error
	}{
		{
			name:        "No Authorization Header",
			headers:     http.Header{},
			expectedKey: "",
			expectedErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:        "Malformed Authorization Header",
			headers:     http.Header{"Authorization": []string{"InvalidHeader"}},
			expectedKey: "",
			expectedErr: errors.New("malformed authorization header"),
		},
		{
			name:        "Valid Authorization Header",
			headers:     http.Header{"Authorization": []string{"ApiKey valid_api_key"}},
			expectedKey: "valid_api_key",
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, gotErr := GetAPIKey(tt.headers)

			if gotKey != tt.expectedKey {
				t.Errorf("GetAPIKey() got key = %v, expected %v", gotKey, tt.expectedKey)
			}

			if (gotErr != nil && tt.expectedErr == nil) || (gotErr == nil && tt.expectedErr != nil) || (gotErr != nil && gotErr.Error() != tt.expectedErr.Error()) {
				t.Errorf("GetAPIKey() got error = %v, expected %v", gotErr, tt.expectedErr)
			}
		})
	}
}
