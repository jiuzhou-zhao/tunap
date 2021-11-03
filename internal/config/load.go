package config

import (
	"io/ioutil"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

func LoadConfig(configName string, cfg interface{}) error {
	viper.AddConfigPath("./")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../config/")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("../../config/")
	viper.SetConfigType("yaml")
	viper.SetConfigName(configName)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	bs, err := ioutil.ReadFile(viper.ConfigFileUsed())

	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(bs, cfg); err != nil {
		//if err := viper.Unmarshal(cfg); err != nil {
		return err
	}

	return nil
}
