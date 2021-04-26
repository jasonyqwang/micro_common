package conf

import (
	"fmt"
	"github.com/asim/go-micro/v3/config"
)

type ElasticConfig struct {
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
	//日志的索引名称
	LogIndexName string `json:"logIndexName"`
}

func GetElasticFromConsul(config config.Config, path ...string) *ElasticConfig {
	elastic := &ElasticConfig{}

	err := config.Get(path...).Scan(elastic)
	if err != nil {
		fmt.Println(path, " 配置，从Consul获取失败，err:", err)
	}
	return elastic
}
