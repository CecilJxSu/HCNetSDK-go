package example

import (
	hik "HCNetSDK-go/hikvision"
	"fmt"
	"strconv"
	"unsafe"
)

// 报警布防例子
func AlarmMsgExample() {
	// 初始化
	success := hik.NET_DVR_Init()
	// 设置连接超时时间与重连功能
	success = !success || hik.NET_DVR_SetConnectTime(2000, 1)
	success = !success || hik.NET_DVR_SetReconnect(10000, true)
	if !success {
		// 初始化失败
		printError("Initial failed")
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
	loginInfo.CbLoginResult = loginCallbackForAlarmExample
	// 登陆
	result := hik.NET_DVR_Login_V40(&loginInfo, &deviceInfo)
	if -1 == result {
		printError("Login failed")
		hangOn <- false
	}

	// hang on for testing async callback
	<-hangOnForAlarm

	// 释放资源
	hik.NET_DVR_Cleanup()
}

var hangOnForAlarm = make(chan bool)

// 登陆异步回调函数
func loginCallbackForAlarmExample(lUserID int, dwResult uint32, lpDeviceInfo hik.LPNET_DVR_DEVICEINFO_V30, pUser unsafe.Pointer) {
	if 1 != dwResult {
		fmt.Println("异步登陆失败")
		printError("Login failed")
		hangOn <- false
		return
	}

	fmt.Println("异步登陆成功，userId: ", strconv.Itoa(lUserID))
	fmt.Println("设备序列号: ", string(lpDeviceInfo.SSerialNumber[:]))

	success := hik.NET_DVR_SetDVRMessageCallBack_V31(msgCallback, nil)
	if !success {
		printError("Set alarm callback failed")
		hangOnForAlarm <- false
		return
	}

	setupParam := hik.NET_DVR_SETUPALARM_PARAM{}
	setupParam.DwSize = uint32(unsafe.Sizeof(setupParam))
	setupParam.ByLevel = 2

	lAlarmHandle := hik.NET_DVR_SetupAlarmChan_V41(lUserID, &setupParam)
	if -1 == lAlarmHandle {
		printError("Setup alarm channel failed")
	}

	<-setupAlarmFinished

	success = hik.NET_DVR_CloseAlarmChan_V30(lAlarmHandle)
	if !success {
		printError("Close alarm channel failed")
		hangOnForAlarm <- false
	}
	hangOnForAlarm <- true
}

// 设置报警渠道结束标记
var setupAlarmFinished = make(chan bool, 1)

func msgCallback(lCommand int, pAlarmer hik.LPNET_DVR_ALARMER, pAlarmInfo *byte, dwBufLen uint32, pUser unsafe.Pointer) bool {
	fmt.Println("lCommand: ", lCommand)
	fmt.Println("sDeviceName: ", string(pAlarmer.SDeviceName[:]))
	fmt.Println("sSerialNumber: ", string(pAlarmer.SSerialNumber[:]))
	fmt.Println("ByMacAddr: ", pAlarmer.ByMacAddr)
	// 门禁主机报警信息
	if hik.COMM_ALARM_ACS == lCommand {
		info := *(hik.LPNET_DVR_ACS_ALARM_INFO)(unsafe.Pointer(pAlarmInfo))
		fmt.Println("dwMajor: ", info.DwMajor)
	}
	setupAlarmFinished <- true
	return true
}
