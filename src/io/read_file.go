package io

import (
	"fmt"
	"golang.org/x/sys/unix"
	_ "golang.org/x/sys/unix"
	"io/ioutil"
)

var basePath = "source/"

func ReadJson(fileName string) string {
	var path = basePath + fileName
	fmt.Printf("%d%s", unix.Getpid(), "--- 准备读取文件:["+path+"]......")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("读取异常", err)
	} else {
		jsonData := string(data)
		fmt.Println("读取完成")
		return jsonData
	}
	return ""
}
