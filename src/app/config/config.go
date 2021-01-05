/**
 * @Author: 张俊ZhangJun
 * @Date: 2020/12/16 上午10:26
 */
package config

import (
	"github.com/caarlos0/env"
	"github.com/caarlos0/env/parsers"
	"net/url"
)

type Config struct {
	Port int `env:"PORT" envDefault:"8050"`
	//MysqlURL url.URL `env:"MYSQL_URL" envDefault:"mysql://root:123@0.0.0.0:3306/test?charset=utf8&parseTime=True"` //?parseTime=true
	MysqlURL url.URL `env:"MYSQL_URL" envDefault:"mysql://root:123@0.0.0.0:3306/test?parseTime=true"` //?parseTime=true

}

func ReadEnv() *Config {
	cfg := Config{}
	if err := env.ParseWithFuncs(&cfg, env.CustomParsers{
		parsers.URLType: parsers.URLFunc,
	}); err != nil {
		panic(err)
	}
	return &cfg
}
