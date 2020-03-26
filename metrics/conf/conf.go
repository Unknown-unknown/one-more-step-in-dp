package conf

import (
	"flag"
	"github.com/spf13/viper"
)

var (
	file string
	C    *Config
)

type Config struct {
	Email *Email
}

type Email struct {
	Host     string
	Port     int
	Username string
	Password string
}

func init() {
	flag.StringVar(&file, "f", "conf/conf.yaml", "conf file path")
	flag.Parse()
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&C); err != nil {
		panic(err)
	}
}
