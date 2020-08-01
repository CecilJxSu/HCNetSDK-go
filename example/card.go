package example

import (
	hik "HCNetSDK-go/hikvision"
	"encoding/binary"
	"fmt"
	"strconv"
	"unsafe"
)

// 获取和下发卡例子
func CardExample() {
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
	loginInfo.CbLoginResult = loginCallback
	// 登陆
	result := hik.NET_DVR_Login_V40(&loginInfo, &deviceInfo)
	if -1 == result {
		printError("Login failed")
	}

	// hang on for testing async callback
	<-hangOn

	// 释放资源
	hik.NET_DVR_Cleanup()
}

var hangOn = make(chan bool)

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
		printError("Login failed")
		return
	}

	fmt.Println("异步登陆成功，userId: ", strconv.Itoa(lUserID))
	fmt.Println("设备序列号: ", string(lpDeviceInfo.SSerialNumber[:]))

	// 设置卡参数
	setCardInfo := hik.NET_DVR_CARD_CFG_COND{}
	setCardInfo.DwSize = uint32(unsafe.Sizeof(setCardInfo))
	setCardInfo.DwCardNum = 1
	setCardInfo.ByCheckCardNo = 1
	// 设置卡
	lHandle := hik.NET_DVR_StartRemoteConfig(
		lUserID,
		hik.NET_DVR_SET_CARD_CFG_V50,
		unsafe.Pointer(&setCardInfo),
		unsafe.Sizeof(setCardInfo),
		remoteConfigCallback,
		nil,
	)

	// -1 失败，其它值为长连接句柄
	if -1 == lHandle {
		printError("Start remote failed")
	}

	// 下发卡参数
	setCardBuf := hik.NET_DVR_CARD_CFG_V50{}
	structSize := uint32(unsafe.Sizeof(setCardBuf))
	setCardBuf.DwSize = structSize
	copy(setCardBuf.ByCardNo[:], "957005719")
	// 设置：卡是否有效参数; 首卡参数
	setCardBuf.DwModifyParamType = 0x00000001 | 0x00000010
	setCardBuf.ByCardValid = 1
	setCardBuf.ByLeaderCard = 1
	// 发送卡参数
	success := hik.NET_DVR_SendRemoteConfig(lHandle, hik.ENUM_ACS_SEND_DATA, (*byte)(unsafe.Pointer(&setCardBuf)), structSize)
	if !success {
		printError("Send remote config failed")
		sendRemoteFinished <- true
	}

	<-sendRemoteFinished

	// 关闭长连接
	success = hik.NET_DVR_StopRemoteConfig(lHandle)
	if !success {
		printError("Stop remote config failed")
	}

	fmt.Println("=========================")

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
	lHandle = hik.NET_DVR_StartRemoteConfig(
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
	structSize = uint32(unsafe.Sizeof(queryBuf))
	queryBuf.DwSize = structSize
	copy(queryBuf.ByCardNo[:], "957005719")
	success = hik.NET_DVR_SendRemoteConfig(lHandle, hik.ENUM_ACS_SEND_DATA, (*byte)(unsafe.Pointer(&queryBuf)), structSize)
	if !success {
		printError("Send remote config failed")
		sendRemoteFinished <- true
	}

	// 下发参数是否结束
	<-sendRemoteFinished

	// 关闭长连接
	success = hik.NET_DVR_StopRemoteConfig(lHandle)
	if !success {
		printError("Stop remote config failed")
	}

	// 退出应用
	hangOn <- true
}

// 下发参数是否结束
var sendRemoteFinished = make(chan bool, 1)

// 远程配置回调函数
func remoteConfigCallback(dwType uint32, lpBuffer unsafe.Pointer, dwBufLen uint32, pUserData unsafe.Pointer) {
	fmt.Println("dwType: ", dwType)
	if dwType == 0 {
		dwStatusArr := *(*[4]byte)(lpBuffer)
		dwStatus := binary.LittleEndian.Uint32(dwStatusArr[:])
		fmt.Println("dwStatus: ", dwStatus)
		fmt.Println("dwBufLen: ", dwBufLen)

		switch dwStatus {
		// 处理中
		case uint32(hik.NET_SDK_CALLBACK_STATUS_PROCESSING):
			dwStatusArr := *(*[4 + 32]byte)(lpBuffer)
			fmt.Println("Processing, cardNo: ", string(dwStatusArr[4:]))
			return
		}

		// 断开长链接
		sendRemoteFinished <- true
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
