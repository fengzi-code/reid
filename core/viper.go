package core

import (
	"fmt"
	"github.com/spf13/viper"
	"reid/global"
)

func LoadConfigFromYaml(configName string) *viper.Viper {

	//读取yaml文件
	v := viper.New()
	//设置读取的配置文件
	v.SetConfigName(configName)
	//添加读取的配置文件路径
	v.AddConfigPath("./config/")
	//设置配置文件类型
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
	}

	//反序列化

	if err := v.Unmarshal(&global.GvaConfig); err != nil {
		fmt.Printf("err:%s", err)
	}
	return v
}
