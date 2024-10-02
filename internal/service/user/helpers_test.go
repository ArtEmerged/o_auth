package user

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserService_hashSha256(t *testing.T) {
	tests := []struct {
		name         string
		salt         []byte
		input        string
		expectedHash string
	}{
		{
			name:         "basic test case",
			salt:         []byte("random_salt"),
			input:        "test_input",
			expectedHash: "0bb3ea740a2469ee32953412daa670936702389e9e5a3be49c49a2500d949de6",
		},
		{
			name:         "empty input",
			salt:         []byte("random_salt"),
			input:        "",
			expectedHash: "dd2a22ffc31d7f87c97747104859d031f95ba14b9ed9d0db7b01bb69f5cff004",
		},
		{
			name:         "empty salt",
			salt:         []byte(""),
			input:        "test_input",
			expectedHash: "952822de6a627ea459e1e7a8964191c79fccfb14ea545d93741b5cf3ed71a09a",
		},
		{
			name:         "empty salt and input",
			salt:         []byte(""),
			input:        "",
			expectedHash: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			name:         "different salt",
			salt:         []byte("different_salt"),
			input:        "test_input",
			expectedHash: "5ac65be607b1d2ac35eb8dd86c251e363edaf815521eb2b6e6faf231bdb67713",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &userService{
				salt: tt.salt,
			}

			// Выполнение тестируемой функции
			result := service.hashSha256(tt.input)

			require.Equal(t, tt.expectedHash, result)
		})
	}
}
