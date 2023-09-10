package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

const (
	DB_CFG_PREFIX = "newsbotdb"
	TG_CFG_PREFIX = "newsbottg"
)

type Config struct {
	DB       DatabaseConfig
	Telegram TelegramBotConfig

	Messages
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Name     string
	Username string
	Password string
}

type TelegramBotConfig struct {
	BotToken string `envconfig:"BOT_TOKEN"`
	BotURL   string `envconfig:"BOT_URL"`
}

type Messages struct {
	Responses
	Errors
}

type Responses struct {
	Start          string `mapstructure:"start"`
	UnknownCommand string `mapstructure:"unknown_command"`
	SourceAdded    string `mapstructure:"source_added"`
}

type Errors struct {
	Default    string `mapstructure:"default"`
	InvalidURL string `mapstructure:"invalid_url"`
}

// Recieve configuration values from env variables
func InitConfig() (*Config, error) {
	var cfg Config

	// Prepare viper
	if err := setUpViper(); err != nil {
		return nil, err
	}

	// Read database config from env
	if err := envconfig.Process(DB_CFG_PREFIX, &cfg.DB); err != nil {
		return nil, err
	}

	// Read telegram bot config from env
	if err := envconfig.Process(TG_CFG_PREFIX, &cfg.Telegram); err != nil {
		return nil, err
	}

	// Read messages (responses and errors) from yml file
	if err := viper.UnmarshalKey("response", &cfg.Messages.Responses); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("error", &cfg.Messages.Errors); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func setUpViper() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("messages")

	return viper.ReadInConfig()
}
