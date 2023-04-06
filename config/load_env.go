package config

import "github.com/spf13/viper"

type Config struct {
	DBHost     string `mapstructure:"MYSQL_HOST"`
	DBPort     string `mapstructure:"MYSQL_PORT"`
	DBUser     string `mapstructure:"MYSQL_USER"`
	DBPassword string `mapstructure:"MYSQL_PASSWORD"`
	DBName     string `mapstructure:"MYSQL_DB"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("dev")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
