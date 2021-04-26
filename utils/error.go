package utils

import (
	"fmt"
	"log"
)

//错误处理函数
func FailOnError(err error, msg string) {
	if err != nil {
		str := fmt.Sprintf("msg:%s,err:%s", msg, err)
		log.Fatalf(str)
	}
}
