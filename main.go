package main

import (
	hik "HCNetSDK-go/hikvision"
	"fmt"
	"strconv"
	"unsafe"
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

	// 登陆参数
	loginInfo := hik.NET_DVR_USER_LOGIN_INFO{}
	deviceInfo := hik.NET_DVR_DEVICEINFO_V40{}

	copy(loginInfo.SDeviceAddress[:], "192.168.8.110")
	loginInfo.WPort = 8000
	loginInfo.ByUseTransport = 0
	loginInfo.BUseAsynLogin = 1
	copy(loginInfo.SUserName[:], "<your username>")
	copy(loginInfo.SPassword[:], "<your password>")
	// 异步登陆回调函数
	loginInfo.CbLoginResult = func(lUserID int, dwResult uint32, lpDeviceInfo hik.LPNET_DVR_DEVICEINFO_V30, pUser unsafe.Pointer) {
		if 1 == dwResult {
			fmt.Println("异步登陆成功，userId: ", strconv.Itoa(lUserID))
			fmt.Println("设备序列号: ", string(lpDeviceInfo.SSerialNumber[:]))
		} else {
			fmt.Println("异步登陆失败")
		}
	}
	// 登陆
	hik.NET_DVR_Login_V40(&loginInfo, &deviceInfo)

	hangOn := make(chan int)
	<-hangOn
}
