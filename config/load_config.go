package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	var err error
	Appconfig, err = LoadConfig()
	if err != nil {
		fmt.Println("Not able to load config file big dick")
	}
	fmt.Println(*Appconfig)
}
func LoadConfig() (*Config, error) {

	var cfg Config

	viper.AddConfigPath("./config")
	//set the file name and path
	viper.SetConfigName("dbconfig")
	viper.SetConfigType("toml")
	viper.ReadInConfig()
	//read the configuration file
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	
	return &cfg, nil
}
