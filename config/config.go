package config

import (
	"AutoBan/internal/errors"
	"AutoBan/pkg/logger"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Environment string `mapstructure:"environment"`
	DB          struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"db"`
	Redis struct {
		Addr     string `mapstructure:"addr"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
	} `mapstructure:"redis"`
	JWT struct {
		Secret string `mapstructure:"secret"`
	} `mapstructure:"jwt"`
	Server struct {
		Address string `mapstructure:"address"`
		Port    string `mapstructure:"port"`
	} `mapstructure:"server"`
}

var (
	config    *Config
	once      sync.Once
	configErr error
)

// LoadConfig returns a singleton instance of Config
func GetConfig() (*Config, error) {
	once.Do(func() {
		config, configErr = loadConfig()
	})
	return config, configErr
}

// loadConfig is the internal function that actually loads the configuration
func loadConfig() (*Config, error) {
	v := viper.New()

	// Get environment from environment variable or default to development
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	// Set default values
	setDefaultValues(v)

	// Read from YAML file first (lower priority)
	readYAMLConfig(v)

	// Read from .env file (higher priority)
	readEnvConfig(v)

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		logger.Error(err, "Failed to unmarshal config")
		return nil, errors.ErrLoadConfig
	}
	return &config, nil
}

func setDefaultValues(v *viper.Viper) {
	v.SetDefault("environment", "development")

	v.SetDefault("server.port", "8080")
	v.SetDefault("server.address", "localhost")

	v.SetDefault("db.port", "5432")
	v.SetDefault("db.host", "localhost")
	v.SetDefault("db.user", "autoban")
	v.SetDefault("db.password", "autoban")
	v.SetDefault("db.name", "autoban")

	v.SetDefault("redis.addr", "localhost:6379")
	v.SetDefault("redis.password", "autoban")
	v.SetDefault("redis.db", 0)

	v.SetDefault("jwt.secret", "mysecretkey")
}

func readYAMLConfig(v *viper.Viper) {
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")

	if err := v.ReadInConfig(); err != nil {
		logger.Info("Failed to read yaml config")
	}
}

func readEnvConfig(v *viper.Viper) {
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		logger.Info("Failed to read env config")
	} else {
		v.Set("environment", v.GetString("ENVIRONMENT"))

		v.Set("server.port", v.GetString("SERVER_PORT"))
		v.Set("server.address", v.GetString("SERVER_ADDRESS"))

		v.Set("db.host", v.GetString("DB_HOST"))
		v.Set("db.port", v.GetString("DB_PORT"))
		v.Set("db.user", v.GetString("DB_USER"))
		v.Set("db.password", v.GetString("DB_PASSWORD"))
		v.Set("db.name", v.GetString("DB_NAME"))

		v.Set("redis.addr", v.GetString("REDIS_ADDR"))
		v.Set("redis.password", v.GetString("REDIS_PASSWORD"))
		v.Set("redis.db", v.GetString("REDIS_DB"))

		v.Set("jwt.secret", v.GetString("JWT_SECRET"))
	}
}
