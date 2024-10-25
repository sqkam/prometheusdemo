package ioc

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port        int          `mapstructure:"port"`
	MySQLConfig *MySQLConfig `mapstructure:"mysql"`
}

type MySQLConfig struct {
	Host            string `mapstructure:"host"`
	Port            int64  `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	DB              string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"` //hour

}

var configPath = flag.String("c", "./config.yaml", "配置文件路径")

func InitConfig() *ServerConfig {
	if !flag.Parsed() {
		flag.Parse()
	}

	var serverConfig ServerConfig
	v := viper.New()
	v.SetConfigFile(*configPath)
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("read viper config failed: %s", err.Error())
		panic(err)
	}
	if err := v.Unmarshal(&serverConfig); err != nil {
		fmt.Printf("unmarshal err failed: %s", err.Error())
		panic(err)
	}
	return &serverConfig

}
