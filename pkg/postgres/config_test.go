//go:build mock

package postgres //nolint:testpackage

import (
	"reflect"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/require"
)

func Test_NewConfigDefault(t *testing.T) {
	t.Parallel()

	var (
		expectedDSN        = "dsn"
		expectedAppName    = "test"
		expectedDriverType = PGXDriver
	)

	cfg := newConfig(expectedDSN, expectedAppName)

	require.NoError(t, cfg.configValidation())

	require.Equal(t, expectedDSN, cfg.dsn)
	require.Empty(t, cfg.databaseName)
	require.Equal(t, dialerNetwork, cfg.networkType)
	require.Equal(t, expectedAppName, cfg.appName)
	require.Equal(t, defaultMaxOpenConnections, cfg.maxOpenConnections)
	require.Zero(t, cfg.maxIdleConnections)
	require.Equal(t, defaultDialTimeout, cfg.dialTimeout)
	require.Equal(
		t,
		reflect.ValueOf(pgconn.ValidateConnectTargetSessionAttrsReadWrite).Pointer(),
		reflect.ValueOf(cfg.validateConnFunc).Pointer(),
	)
	require.Nil(t, cfg.queryExecMode)
	require.Equal(t, expectedDriverType, cfg.driverType)
}

func Test_NewConfigWithOptions(t *testing.T) {
	t.Parallel()

	var (
		expectedDSN     = "dsn"
		expectedAppName = "test"

		expectedOpenConnections    = 5
		expectedMaxIdleConnections = 6
		expectedDialTimeout        = 10 * time.Second
		expectedQueryExecMode      = pgx.QueryExecModeSimpleProtocol
	)

	cfg := newConfig(
		expectedDSN, expectedAppName,
		WithOpenConnections(expectedOpenConnections),
		WithMaxIdleConnections(expectedMaxIdleConnections),
		WithDialTimeout(expectedDialTimeout),
		WithQueryExecMode(expectedQueryExecMode),
		WithoutValidateConnectFunc(),
	)

	require.Equal(t, expectedOpenConnections, cfg.maxOpenConnections)
	require.Equal(t, expectedMaxIdleConnections, cfg.maxIdleConnections)
	require.Equal(t, expectedDialTimeout, cfg.dialTimeout)
	require.Equal(t, &expectedQueryExecMode, cfg.queryExecMode)
	require.Nil(t, cfg.validateConnFunc)
}

func Test_NewConfigInvalidDSN(t *testing.T) {
	t.Parallel()

	cfg := newConfig("", "test")

	err := cfg.configValidation()

	require.Error(t, err)
	require.ErrorIs(t, err, ErrEmptyDSN)
}

func Test_NewConfigInvalidAppName(t *testing.T) {
	t.Parallel()

	cfg := newConfig("dsn", "")

	err := cfg.configValidation()

	require.Error(t, err)
	require.ErrorIs(t, err, ErrEmptyAppName)
}
