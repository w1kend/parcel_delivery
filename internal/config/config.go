package config

import (
	"sync"

	"github.com/caarlos0/env/v6"
)

type config struct {
	Pgdsn           string `env:"PG_DSN,notEmpty"`
	JwtSecret       string `env:"JWT_SECRET,notEmpty"`
	AppPort         string `env:"APP_PORT,notEmpty"`
	PgMaxOpenConns  int    `env:"PG_MAX_OPEN_CONNS" envDefault:"20"`
	PgMaxIddleConns int    `env:"PG_MAX_IDLE_CONNS" envDefault:"20"`

	values map[string]value
}

var (
	cfg  config
	once = sync.Once{}
	mu   sync.RWMutex
)

const (
	PgDsn           = "pgdsn"
	PgMaxOpenConns  = "pg_max_open_conns"
	PgMaxIddleConns = "pg_max_idle_conns"
	JWTSecret       = "jwt_secret"
	AppPort         = "app_port"
)

type value struct {
	source interface{}
}

func (v value) Int() int {
	if i, ok := v.source.(int); ok {
		return i
	}

	return 0
}

func (v value) String() string {
	if s, ok := v.source.(string); ok {
		return s
	}

	return ""
}

func buildConfig() {
	once.Do(func() {
		mu.Lock()
		defer mu.Unlock()

		var c config

		err := env.Parse(&c)
		if err != nil {
			panic("failed to init config: " + err.Error())
		}

		c.values = map[string]value{
			PgDsn:           {source: c.Pgdsn},
			JWTSecret:       {source: c.JwtSecret},
			AppPort:         {source: c.AppPort},
			PgMaxIddleConns: {source: c.PgMaxIddleConns},
			PgMaxOpenConns:  {source: c.PgMaxOpenConns},
		}

		cfg = c
	})
}

func init() {
	buildConfig()
}

func GetValue(name string) value {
	mu.RLock()
	defer mu.RUnlock()
	return cfg.values[name]
}
