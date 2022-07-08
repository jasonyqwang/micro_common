package conf

import (
	"github.com/asim/go-micro/v3/config"
	"strings"
)

type AliyunOssConfig struct {
	AccessKeyId     string `json:"accessKeyId" mapstructure:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret" mapstructure:"accessKeySecret"`
	Endpoint        string `json:"endpoint" mapstructure:"endpoint"`
	Host            string `json:"host" mapstructure:"host"`
	Bucket          string `json:"bucket" mapstructure:"bucket"`
}

func (t AliyunOssConfig) GetBaseUrl() string {
	host := t.Host
	if t.Host == "" {
		host = t.Bucket + "." + strings.Replace(t.Endpoint, "http://", "https://", -1)
	}
	return "http://" + host
}

//获取oss的配置
func GetOssFromConsul(config config.Config, path ...string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}
	//获取配置
	config.Get(path...).Scan(mysqlConfig)
	return mysqlConfig
}
