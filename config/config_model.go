package config

type Config struct {
	App struct {
		Name       string
		Env        Env
		Host       string
		Port       string
		TimeoutSec int
	}
	Swagger struct {
		Addr string
	}
	DB struct {
		Host     string
		Port     string
		DBName   string
		User     string
		Password string
		SslMode  string
	}
	Otel struct {
		Traces struct {
			Endpoint string
		}
		Metrics struct {
			Endpoint string
		}
	}
}
