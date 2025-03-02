package cfg

import (
	"net"
	"os"
	"strconv"
	"strings"

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

const envProd = "prod"

type App struct {
	ENV              string  `envconfig:"APP_ENV"            default:"local"`
	Name             string  `envconfig:"APP_NAME"           default:"logistic-tracker"`
	LogLevel         string  `envconfig:"LOG_LEVEL"          default:"error"`
	WithTrace        bool    `envconfig:"WITH_TRACE"`
	TraceSampleRatio float64 `envconfig:"TRACE_SAMPLE_RATIO"`
}

func (a App) IsProduction() bool {
	return strings.Contains(a.ENV, envProd)
}

type OTLPTrace struct {
	Endpoint string `envconfig:"OTLP_TRACE_ENDPOINT"`
	TLS      bool   `envconfig:"OTLP_TRACE_TLS"`
	WithDB   bool   `envconfig:"TRACE_DB"`
}

type HTTP struct {
	Host string `envconfig:"HTTP_HOST"`
	Port int32  `envconfig:"HTTP_PORT"    default:"8080"`
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
