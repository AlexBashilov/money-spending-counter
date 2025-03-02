package cfg

import (
	"net"
	"strings"
)

type Postgres struct {
	Host     string `envconfig:"POSTGRES_HOST" default:"localhost"`
	Port     string `envconfig:"POSTGRES_PORT" default:"5432"`
	Username string `envconfig:"POSTGRES_USER" default:"root"`
	Password string `envconfig:"POSTGRES_PASSWORD"`
	Database string `envconfig:"POSTGRES_DB"`
	SSL      string `envconfig:"POSTGRES_SSL" default:"disable"`
	Debug    string `envconfig:"POSTGRES_DEBUG"`
}

func (c *Postgres) DSN() string {
	dsn := strings.Builder{}
	dsn.WriteString("postgres://")

	if c.Username != "" {
		dsn.WriteString(c.Username)

		if c.Password != "" {
			dsn.WriteString(":" + c.Password)
		}

		dsn.WriteString("@")
	}

	hp := net.JoinHostPort(c.Host, c.Port)

	dsn.WriteString(hp + "/")

	if c.Database != "" {
		dsn.WriteString(c.Database)
	}

	dsn.WriteString("?sslmode=" + c.SSL)

	return dsn.String()
}
