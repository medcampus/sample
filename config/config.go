package config

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	var err error
	err = viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s\n", err.Error()))
	}

	log.Info("config", "Initial Configurations")
	allConfigs, _ := json.Marshal(viper.AllSettings())
	log.Info(string(allConfigs))
}
