package configuration

import "github.com/spf13/viper"

type Configuration struct {
	Data   Data
	Broker Broker
	Server Server
}

func Load(path string) (*Configuration, error) {
	cfg := Configuration{}

	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
