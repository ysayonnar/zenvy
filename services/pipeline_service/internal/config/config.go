package config

import "fmt"

type Config struct {
	IsDebug bool `env:"IS_DEBUG"`
	//Server   Server
	Postgres Postgres
}

//type Server struct {
//	Host            string        `env:"HOST"`
//	Port            int           `env:"APP_PORT"`
//	ReadTimeout     time.Duration `env:"READ_TIMEOUT"`
//	WriteTimeout    time.Duration `env:"WRITE_TIMEOUT"`
//	IdleTimeout     time.Duration `env:"IDLE_TIMEOUT"`
//	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT"`
//}

type Postgres struct {
	DbHost     string `env:"PIPELINE_SERVICE_DB_HOST" env-required:"true"`
	DbPort     int    `env:"PIPELINE_SERVICE_DB_PORT" env-required:"true"`
	DbUser     string `env:"PIPELINE_SERVICE_DB_USER" env-required:"true"`
	DbPassword string `env:"PIPELINE_SERVICE_DB_PASSWORD" env-required:"true"`
	DbName     string `env:"PIPELINE_SERVICE_DB_NAME" env-required:"true"`
	DbSslMode  string `env:"PIPELINE_SERVICE_DB_SSL_MODE" env-required:"true"`
}

func (cfg Postgres) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUser,
		cfg.DbPassword,
		cfg.DbName,
		cfg.DbSslMode)
}
