package app

import (
	"GorillaCacher/internal/utils"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func (a *application) SetupViper(path string) error {
	viper.AutomaticEnv()
	viper.SetConfigType("json")

	f, err := os.Open(path)
	if err != nil {
		msg := fmt.Sprintf("cannot read config file: %s", err.Error())
		return errors.New(msg)
	}

	err = viper.ReadConfig(f)
	if err != nil {
		msg := fmt.Sprintf("viper read config error: %s", err.Error())
		return errors.New(msg)
	}

	var c utils.ServiceConfig
	err = viper.Unmarshal(&c)
	if err != nil {
		return err
	}

	a.config = &c
	fmt.Println("Viper loaded successfully")
	fmt.Println(a.config.Nodes)
	return nil
}
