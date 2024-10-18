package configs

import "github.com/spf13/viper"

var config *Config

type option struct {
	configFolder []string
	configFile   string
	configType   string
}

func Init() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	config = new(Config)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	config.Service.SecretJWT = viper.GetString("SERVICE_SECRET_KEY")
	config.Database.DataSourceName = viper.GetString("DATABASE_DSN")

	return nil
}

type Option func(*option)

func Get() *Config {
	if config == nil {
		config = &Config{}
	}

	return config
}
