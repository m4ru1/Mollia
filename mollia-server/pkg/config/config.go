package config

import "github.com/spf13/viper"

type Config struct {
	DB_DSN     string `mapstructure:"DB_DSN"`
	JWT_SECRET string `mapstructure:"JWT_SECRET"`
	PORT       string `mapstructure:"PORT"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		// 如果没有找到配置文件，可以忽略，因为我们会依赖环境变量
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return
		}
	}

	err = viper.Unmarshal(&config)
	return
}
