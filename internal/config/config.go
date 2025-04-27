package config

import (
	"net"
	"os"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
)

type Config struct {
	App       App
	OTLPTrace OTLPTrace
	HTTP      HTTP
	Postgres  Postgres
}

type App struct {
	Environment      string  `envconfig:"APP_ENV"            default:"local"`
	Name             string  `envconfig:"APP_NAME"           default:"app"`
	LogLevel         string  `envconfig:"LOG_LEVEL"          default:"debug"`
	WithTrace        bool    `envconfig:"WITH_TRACE"`
	TraceSampleRatio float64 `envconfig:"TRACE_SAMPLE_RATIO"`
}

type OTLPTrace struct {
	Endpoint  string `envconfig:"OTLP_TRACE_ENDPOINT"`
	TLS       bool   `envconfig:"OTLP_TRACE_TLS"`
	ByteLimit int    `envconfig:"OTLP_TRACE_REQ_RESP_BYTE_LIMIT" default:"5120"`
}

type HTTP struct {
	Host    string   `envconfig:"HTTP_HOST"`
	Port    int32    `envconfig:"HTTP_PORT"    default:"8080"`
	Schemes []string `envconfig:"HTTP_SCHEMES" default:"http"`
}

func Load() (Config, error) {
	cnf := Config{} //nolint:exhaustruct

	if err := godotenv.Load(".env"); err != nil && !errors.Is(err, os.ErrNotExist) {
		return cnf, errors.Wrap(err, "read .env file")
	}

	if err := envconfig.Process("", &cnf); err != nil {
		return cnf, errors.Wrap(err, "read environment")
	}

	return cnf, nil
}

func (c *Config) LogLevel() (zerolog.Level, error) {
	lvl, err := zerolog.ParseLevel(c.App.LogLevel)
	if err != nil {
		return 0, errors.Wrapf(err, "loading log level from config value %q", c.App.LogLevel)
	}

	return lvl, nil
}

func (c *Config) HTTPAddr() string {
	return net.JoinHostPort(c.HTTP.Host, strconv.Itoa(int(c.HTTP.Port)))
}

func (c *Config) UseTrace() bool {
	return c.App.WithTrace && c.OTLPTrace.Endpoint != ""
}
