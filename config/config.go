package config

import "github.com/spf13/viper"

type Config struct {
	Data   Data
	Events Events
}

func Load(path string) (*Config, error) {
	config := Config{}

	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
