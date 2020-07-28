package main

import (
	hik "HCNetSDK-go/hikvision"
	"encoding/binary"
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	// 初始化
	success := hik.NET_DVR_Init()
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
	loginInfo.CbLoginResult = loginCallback
	// 登陆
	result := hik.NET_DVR_Login_V40(&loginInfo, &deviceInfo)
	if -1 == result {
		printError("Login failed")
	}

	// hang on for testing async callback
	hangOn := make(chan int)
	<-hangOn
}

// 打印错误信息
func printError(desc string) {
	errCode := hik.NET_DVR_GetLastError()
	fmt.Println(desc, ", error code: ", errCode)
	fmt.Println("Error message: " + hik.NET_DVR_GetErrorMsg(errCode))
}

// 登陆异步回调函数
func loginCallback(lUserID int, dwResult uint32, lpDeviceInfo hik.LPNET_DVR_DEVICEINFO_V30, pUser unsafe.Pointer) {
	if 1 != dwResult {
		fmt.Println("异步登陆失败")
		return
	}

	fmt.Println("异步登陆成功，userId: ", strconv.Itoa(lUserID))
	fmt.Println("设备序列号: ", string(lpDeviceInfo.SSerialNumber[:]))

	// 获取卡
	cardInfo := hik.NET_DVR_CARD_CFG_COND{}
	cardInfo.DwSize = uint32(unsafe.Sizeof(cardInfo))
	// 获取所有卡信息，不需要调 NET_DVR_StartRemoteConfig 接口指定查询条件
	// cardInfo.DwCardNum = 0xffffffff

	// 需要调 NET_DVR_StartRemoteConfig 接口，指定查询条件
	cardInfo.DwCardNum = 1
	cardInfo.ByCheckCardNo = 1
	// 回调函数
	cb := remoteConfigCallback

	// 获取卡参数
	lHandle := hik.NET_DVR_StartRemoteConfig(
		lUserID,
		hik.NET_DVR_GET_CARD_CFG_V50,
		unsafe.Pointer(&cardInfo),
		unsafe.Sizeof(cardInfo),
		cb,
		nil,
	)
	// -1 失败，其它值为长连接句柄
	if -1 == lHandle {
		printError("Start remote failed")
	}

	// 发送查询条件
	queryBuf := hik.NET_DVR_CARD_CFG_SEND_DATA{}
	copy(queryBuf.ByCardNo[:], "123")
	structSize := uint32(unsafe.Sizeof(queryBuf))
	queryBuf.DwSize = structSize
	success := hik.NET_DVR_SendRemoteConfig(lHandle, hik.ENUM_ACS_SEND_DATA, (*byte)(unsafe.Pointer(&queryBuf)), structSize)
	if !success {
		printError("Send remote config failed")
	}
}

// 远程配置回调函数
func remoteConfigCallback(dwType uint32, lpBuffer unsafe.Pointer, dwBufLen uint32, pUserData unsafe.Pointer) {
	fmt.Println("dwType: ", dwType)
	if dwType == 0 {
		dwStatusArr := *(*[4]byte)(lpBuffer)
		dwStatus := binary.LittleEndian.Uint32(dwStatusArr[:])
		fmt.Println("dwStatus: ", dwStatus)
		fmt.Println("dwBufLen: ", dwBufLen)
	} else if dwType == 1 {
		progress := *(*int)(lpBuffer)
		fmt.Println("progress: ", progress)
	} else if dwType == 2 {
		cardCfg := *(hik.LPNET_DVR_CARD_CFG_V50)(lpBuffer)
		fmt.Println("byCardNo: ", string(cardCfg.ByCardNo[:]))
		fmt.Println("byCardValid: ", cardCfg.ByCardValid)
		fmt.Println("byCardType: ", cardCfg.ByCardType)
		fmt.Println("dwCardUserId: ", cardCfg.DwCardUserId)
	}
}
