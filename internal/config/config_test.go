package config_test

import (
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	"booker/internal/config"
)

func TestConfig(t *testing.T) {
	env := map[string]string{
		"APP_ENV":    "test",
		"APP_NAME":   "basic-go-microservice",
		"LOG_LEVEL":  "error",
		"SENTRY_DSN": "test_dsn",
		"WITH_TRACE": "true",
	}

	for k, v := range env {
		err := os.Setenv(k, v)
		require.NoError(t, err)
	}

	conf, err := config.Load()
	require.NoError(t, err)

	lvl, err := conf.LogLevel()
	require.NoError(t, err)
	require.Equal(t, "test", conf.App.ENV)
	require.Equal(t, zerolog.ErrorLevel, lvl)
}
