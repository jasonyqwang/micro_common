package utils

import (
	"encoding/json"
)

//把source数据通过 json tag 对 model 进行结构体赋值
//eg: 把 protobuf模型（source），转成对应的 orm 模型（model）
func SwapTo(source interface{}, model interface{}) (err error) {
	dataByte, err := json.Marshal(source)
	if err != nil {
		return err
	}
	err = json.Unmarshal(dataByte, model)
	return
}

//json 转 map
func JsonToMap(data string) map[string]interface{} {
	var res map[string]interface{}
	if data == "" {
		return res
	}
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		panic(err)
	}
	return res
}
