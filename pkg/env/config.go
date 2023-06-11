package envs

import "github.com/spf13/viper"

type Config struct {
	TELEGRAM_BOT_TOKEN string `mapstructure:"TELEGRAM_BOT_TOKEN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(path)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
