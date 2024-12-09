package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	Database struct {
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
		DB_name    string `yaml:"db_name"`
		Username   string `yaml:"user"`
		Password   string `yaml:"password"`
		Sqlite_dir string `yaml:"sqllite_dir"`
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	InitDB()
}
