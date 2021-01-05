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

	SispUrl url.URL   `env:"SispUrl" envDefault:"http://public-INT-gw.intsit.sfdc.com.cn:1080/userQuery/waybillQueryService/queryWaybillInfo"`
	FxsUrl url.URL    `env:"FxsUrl"  envDefault:"http://sgs-gw-three.intsit.sfdc.com.cn:1080/fxs/appRestQueryService/irCheckResult"`
	DecryptMobileUrl url.URL `env:"DecryptMobileUrl"  envDefault:"http://des-int-gw.intsit.sfdc.com.cn:8080/isic-des-code/code/cryptPhone"`

	SispHeaderSysCode string `env:"SispHeaderSysCode" envDefault:"INC-FCCV-CORE"`
	SispUserId string `env:"SispUserId" envDefault:"676123"`
	SispContextOrder string `env:"SispContextOrderstring" envDefault:"00001"`
	Port             int           `env:"PORT" envDefault:"8050"`
	DecryptMobileAccessToken string `env:"DecryptMobileAccessToken" envDefault:"f9d0bbe61d27bfd4c3ef8cc6d67735f55f71aa0654fa13e0b9f59f5cd2c5a0324fd4cd28f30be0249be141990f33abf993a1bfad40e69c28602673dd4fa89c94"`
	DecryptMobileSysCode string `env:"DecryptMobileSysCode" envDefault:"INC-FCCV-CORE"`

	SIT string `env:"IsSit" envDefault:"1"`  //yes :1   no :0


	MysqlURL url.URL `env:"MYSQL_URL" envDefault:"mysql://root:123@0.0.0.0:3306/test?charset=utf8&parseTime=True"` //?parseTime=true

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
