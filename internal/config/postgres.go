package config

import (
	"net"
	"strings"
)

type Postgres struct {
	HostRW   string `envconfig:"POSTGRES_HOST" default:"localhost"`
	HostRO   string `envconfig:"POSTGRES_HOST_RO" default:"localhost"`
	Port     string `envconfig:"POSTGRES_PORT" default:"5432"`
	User     string `envconfig:"POSTGRES_USER" default:"root"`
	Password string `envconfig:"POSTGRES_PASSWORD"`
	Database string `envconfig:"POSTGRES_DB"`
	SSL      string `envconfig:"POSTGRES_SSL" default:"disable"`
}

func (p Postgres) DSN(ro bool) string {
	dsn := strings.Builder{}
	dsn.WriteString("postgres://")

	if p.User != "" {
		dsn.WriteString(p.User)

		if p.Password != "" {
			dsn.WriteString(":" + p.Password)
		}

		dsn.WriteString("@")
	}

	hp := net.JoinHostPort(p.HostRW, p.Port)

	if ro {
		hp = net.JoinHostPort(p.HostRO, p.Port)
	}

	dsn.WriteString(hp + "/")

	if p.Database != "" {
		dsn.WriteString(p.Database)
	}

	dsn.WriteString("?sslmode=" + p.SSL)

	return dsn.String()
}
