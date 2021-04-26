package conf

import (
	"fmt"
	"github.com/asim/go-micro/v3/config"
)

type RabbitmqConfig struct {
	Host     string `json:"host"`
	Vhost    string `json:"vhost"`
	User     string `json:"user"`
	Password string `json:"password"`
	Port     int    `json:"port"`
}

//获取Rabbitmq的配置
func GetRabbitmqFromConsul(config config.Config, path ...string) *RabbitmqConfig {
	rabbitmqConfig := &RabbitmqConfig{}
	//获取配置
	err := config.Get(path...).Scan(rabbitmqConfig)
	if err != nil {
		fmt.Println(path, " 配置，从Consul获取失败，err:", err)
	}

	return rabbitmqConfig
}
