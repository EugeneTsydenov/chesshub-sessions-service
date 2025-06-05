package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

const defaultConfigName = "config"

const (
	LocalEnv   = "local"
	DevelopEnv = "develop"
	ProdEnv    = "prod"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	GeoIp    GeoIp
}

type AppConfig struct {
	Env  string `mapstructure:"env"`
	Port string `mapstructure:"port"`
}

type DatabaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

type GeoIp struct {
	DatabasePath string `mapstructure:"database_path"`
}

func Load(env, cfgPath string) (*Config, error) {
	if cfgPath == "" {
		return nil, errors.New("config file path is empty")
	}

	cfg := &Config{}

	err := loadEnvCfg(env, cfgPath)
	if err != nil {
		return nil, err
	}

	err = loadEnvVariables()
	if err != nil {
		return nil, err
	}

	if err = viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return cfg, err
}

func loadEnvCfg(env string, cfgPath string) error {
	isValidEnv := env == LocalEnv || env == DevelopEnv || env == ProdEnv
	if !isValidEnv {
		return fmt.Errorf("invalid environment: %s", env)
	}

	cfgName := fmt.Sprintf("%s.%s", defaultConfigName, env)
	viper.SetConfigName(cfgName)
	viper.AddConfigPath(cfgPath)

	if err := viper.MergeInConfig(); err != nil {
		var cfgFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &cfgFileNotFoundError) {
			return fmt.Errorf("failed to merge env config: %w", err)
		}
	}

	return nil
}

func loadEnvVariables() error {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := viper.BindEnv("database.password")
	if err != nil {
		return fmt.Errorf("failed to bind env variables: %w", err)
	}

	return nil
}

func (d DatabaseConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSLMode)
}
