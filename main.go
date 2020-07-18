package main

import (
	hik "./hikvision"
	"fmt"
	"strconv"
)

func main() {
	// 初始化
	success := hik.NET_DVR_Init()
	if !success {
		// 初始化失败
		errCode := hik.NET_DVR_GetLastError()
		fmt.Println("Not success, error code: " + strconv.FormatInt(int64(errCode), 10))
		return
	}
}
