package config

type Config struct {
	App struct {
		Name       string
		Env        Env
		Host       string
		Port       string
		TimeoutSec int
	}
	DB struct {
		Host     string
		Port     string
		DBName   string
		User     string
		Password string
		SslMode  string
	}
}
