package config

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/spf13/viper"
	"github.com/viniosilva/go-boilerplateapi/pkg/logger"
)

type Env string

const (
	envPrefix           = "IB_"
	DBPasswordFieldName = "IB_DB_PASSWORD"
	EnvFieldName        = "IB_ENV"

	EnvDevelopment Env = "dev"
	EnvLocalDocker Env = "docker"
	EnvTest        Env = "test"
	EnvProduction  Env = "prod"
)

type client struct {
	path string
}

func LoadConfig(options ...Option) (Config, error) {
	ctx := context.Background()

	c := &client{path: "."}
	for _, opt := range options {
		opt(c)
	}

	viper.AutomaticEnv()
	env := viper.GetString(EnvFieldName)
	dbPassword := viper.GetString(DBPasswordFieldName)

	viper.SetEnvPrefix(envPrefix)
	viper.SetConfigFile(fmt.Sprintf("%s/.env", c.path))
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		logger.Warn(ctx, "viper.ReadInConfig .env", slog.String("error", err.Error()))
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(fmt.Sprintf("%s/config", c.path))
	if err := viper.MergeInConfig(); err != nil {
		return Config{}, err
	}

	if configName, exists := getEnvConfigName(Env(env)); exists {
		viper.SetConfigName(configName)
		if err := viper.MergeInConfig(); err != nil {
			return Config{}, err
		}
		logger.Info(ctx, "viper.MergeInConfig", slog.String("configName", configName))
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}

	cfg.DB.Password = getEnvFileValue(DBPasswordFieldName, dbPassword)

	return cfg, nil
}

func getEnvConfigName(env Env) (string, bool) {
	if env == "" {
		return "", false
	}

	return fmt.Sprintf("config_%s", env), true
}

func getEnvFileValue(key, envValue string) string {
	value := viper.GetString(key)
	if value != "" {
		return value
	}

	return envValue
}
