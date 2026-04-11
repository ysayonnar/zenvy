package env

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// TODO: this test MUST BE refactored!
func TestParseFromEnv(t *testing.T) {
	tests := []struct {
		Config any
		envs   [2][]string
	}{
		{
			Config: struct {
				Name    string `env:"TEST_NAME" env-required:"true"`
				Surname string `env:"TEST_SURNAME" env-required:"true"`
			}{},
			envs: [2][]string{
				{"TEST_NAME", "TEST_SURNAME"},
				{"fedor", "borisuk"},
			},
		},
		{
			Config: struct {
				Name    string `env:"TEST_LOGIN" env-required:"true"`
				Surname string `env:"TEST_PASSWORD" env-required:"true"`
			}{},
			envs: [2][]string{
				{"TEST_LOGIN", "TEST_PASSWORD"},
				{"admin", "password123"},
			},
		},
	}

	for _, test := range tests {
		for i := range len(test.envs[0]) {
			if err := os.Setenv(test.envs[0][i], test.envs[1][i]); err != nil {
				panic(fmt.Errorf("os.Setenv() error: %w", err))
			}
		}

		err := Parse(test.Config)
		require.NoError(t, err)

		os.Clearenv()
	}
}
