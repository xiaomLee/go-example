package config

import "github.com/spf13/viper"

var (
	Debug   = false
	Default *viper.Viper
)

func Init(path string) error {
	Default = viper.New()
	Default.SetConfigFile(path)
	if err := Default.ReadInConfig(); err != nil {
		return err
	}
	//open db debug for print sql
	if Default.GetBool("server.debug") {
		Debug = true
	}
	return nil
}
