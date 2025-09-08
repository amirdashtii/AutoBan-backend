package config

import (
	"os"
	"sync"
	"strings"

	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/amirdashtii/AutoBan/pkg/logger"

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
	SMS struct {
		BaseURL string `mapstructure:"base_url"`
		XAPIKey string `mapstructure:"x_api_key"`
	} `mapstructure:"sms"`
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

	// Read environment variables directly (highest priority after explicit Set)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

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

	v.SetDefault("sms.base_url", "https://api.sms.ir")
	v.SetDefault("sms.x_api_key", "Aklc5AKdy02FdA03TCwEIZeB6gJ2s0fVv80ejWhUyfS4xpbw")
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
	}

	// Map specific uppercase env vars to nested keys (if present)
	if v.IsSet("ENVIRONMENT") {
		v.Set("environment", v.GetString("ENVIRONMENT"))
	}
	if v.IsSet("SERVER_PORT") { v.Set("server.port", v.GetString("SERVER_PORT")) }
	if v.IsSet("SERVER_ADDRESS") { v.Set("server.address", v.GetString("SERVER_ADDRESS")) }

	if v.IsSet("DB_HOST") { v.Set("db.host", v.GetString("DB_HOST")) }
	if v.IsSet("DB_PORT") { v.Set("db.port", v.GetString("DB_PORT")) }
	if v.IsSet("DB_USER") { v.Set("db.user", v.GetString("DB_USER")) }
	if v.IsSet("DB_PASSWORD") { v.Set("db.password", v.GetString("DB_PASSWORD")) }
	if v.IsSet("DB_NAME") { v.Set("db.name", v.GetString("DB_NAME")) }

	if v.IsSet("REDIS_ADDR") { v.Set("redis.addr", v.GetString("REDIS_ADDR")) }
	if v.IsSet("REDIS_PASSWORD") { v.Set("redis.password", v.GetString("REDIS_PASSWORD")) }
	if v.IsSet("REDIS_DB") { v.Set("redis.db", v.GetString("REDIS_DB")) }

	if v.IsSet("JWT_SECRET") { v.Set("jwt.secret", v.GetString("JWT_SECRET")) }

	if v.IsSet("SMS_BASE_URL") { v.Set("sms.base_url", v.GetString("SMS_BASE_URL")) }
	if v.IsSet("SMS_X_API_KEY") { v.Set("sms.x_api_key", v.GetString("SMS_X_API_KEY")) }
}
