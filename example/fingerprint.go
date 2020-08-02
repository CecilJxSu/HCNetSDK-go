package example

import "C"
import (
	hik "HCNetSDK-go/hikvision"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"unsafe"
)

// 获取和下发指纹例子
func FingerprintExample() {
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
	loginInfo.CbLoginResult = loginCallbackForFingerPrint
	// 登陆
	result := hik.NET_DVR_Login_V40(&loginInfo, &deviceInfo)
	if -1 == result {
		printError("Login failed")
	}

	// hang on for testing async callback
	<-hangOnForFingerPrint

	// 释放资源
	hik.NET_DVR_Cleanup()
}

var hangOnForFingerPrint = make(chan bool)

// 登陆异步回调函数
func loginCallbackForFingerPrint(lUserID int, dwResult uint32, lpDeviceInfo hik.LPNET_DVR_DEVICEINFO_V30, pUser unsafe.Pointer) {
	if 1 != dwResult {
		fmt.Println("异步登陆失败")
		printError("Login failed")
		return
	}

	fmt.Println("异步登陆成功，userId: ", strconv.Itoa(lUserID))
	fmt.Println("设备序列号: ", string(lpDeviceInfo.SSerialNumber[:]))

	// 回调函数复用，因此通过标记区分：1 - 获取指纹参数回调；2 - 设置指纹参数回调
	getOrSet := 2

	// 设置指纹
	setFingerPrint := hik.NET_DVR_FINGER_PRINT_INFO_COND{}
	// 设置数量
	setFingerPrint.DwFingerPrintNum = 1
	copy(setFingerPrint.ByCardNo[:], "123")
	setFingerPrint.DwSize = uint32(unsafe.Sizeof(setFingerPrint))
	setFingerPrint.ByCallbackMode = 0
	setFingerPrint.ByFingerPrintID = 1
	setFingerPrint.ByEnableCardReader[0] = 1
	// 设置指纹
	lHandle := hik.NET_DVR_StartRemoteConfig(
		lUserID,
		hik.NET_DVR_SET_FINGERPRINT_CFG,
		unsafe.Pointer(&setFingerPrint),
		unsafe.Sizeof(setFingerPrint),
		remoteConfigCallbackForFingerPrint,
		nil,
	)

	// -1 失败，其它值为长连接句柄
	if -1 == lHandle {
		printError("Start remote failed")
		hangOnForFingerPrint <- false
		return
	}

	// 下发指纹参数
	setPrintBuf := hik.NET_DVR_FINGER_PRINT_CFG{}
	structSize := uint32(unsafe.Sizeof(setPrintBuf))
	setPrintBuf.DwSize = structSize
	copy(setPrintBuf.ByCardNo[:], "123")
	setPrintBuf.ByEnableCardReader[0] = 1
	setPrintBuf.ByFingerPrintID = 1
	setPrintBuf.ByFingerType = 0

	pwd, _ := os.Getwd()
	// 暂时不支持指纹
	printBuff := readFingerPrintData(pwd + "/res/myprint.bin")
	copy(setPrintBuf.ByFingerData[:], printBuff[:])
	setPrintBuf.DwFingerPrintLen = uint32(len(printBuff))

	// 下发指纹数据
	success := hik.NET_DVR_SendRemoteConfig(
		lHandle,
		hik.ENUM_ACS_SEND_DATA,
		(*byte)(unsafe.Pointer(&setPrintBuf)),
		structSize,
	)
	if !success {
		printError("Send remote config failed")
		hangOnForFingerPrint <- true
		return
	}

	<-sendRemoteFinishedForFingerPrint

	// 关闭长连接
	success = hik.NET_DVR_StopRemoteConfig(lHandle)
	if !success {
		printError("Stop remote config failed")
	}

	fmt.Println("=========================")

	// 回调函数复用，因此通过标记区分：1 - 获取指纹参数回调；2 - 设置指纹参数回调
	getOrSet = 1

	// 获取指纹
	getFingerPrint := hik.NET_DVR_FINGER_PRINT_INFO_COND{}
	getFingerPrint.DwSize = uint32(unsafe.Sizeof(getFingerPrint))
	copy(getFingerPrint.ByCardNo[:], "123")
	getFingerPrint.DwFingerPrintNum = 1
	getFingerPrint.ByFingerPrintID = 1
	getFingerPrint.ByEnableCardReader[0] = 1
	getFingerPrint.ByCallbackMode = 0

	// 回调函数
	cb := remoteConfigCallbackForFingerPrint

	// 获取指纹数据
	lHandle = hik.NET_DVR_StartRemoteConfig(
		lUserID,
		hik.NET_DVR_GET_FINGERPRINT_CFG,
		unsafe.Pointer(&getFingerPrint),
		unsafe.Sizeof(getFingerPrint),
		cb,
		unsafe.Pointer(&getOrSet),
	)
	// -1 失败，其它值为长连接句柄
	if -1 == lHandle {
		printError("Start remote failed")
	}

	// 下发参数是否结束
	<-sendRemoteFinishedForFingerPrint

	// 关闭长连接
	success = hik.NET_DVR_StopRemoteConfig(lHandle)
	if !success {
		printError("Stop remote config failed")
	}

	// 退出应用
	hangOnForFingerPrint <- true
}

// 下发参数是否结束
var sendRemoteFinishedForFingerPrint = make(chan bool, 1)

// 远程配置回调函数
func remoteConfigCallbackForFingerPrint(dwType uint32, lpBuffer unsafe.Pointer, dwBufLen uint32, pUserData unsafe.Pointer) {
	fmt.Println("dwType: ", dwType)
	if dwType == 0 {
		dwStatusArr := *(*[4]byte)(lpBuffer)
		dwStatus := binary.LittleEndian.Uint32(dwStatusArr[:])
		fmt.Println("dwStatus: ", dwStatus)
		fmt.Println("dwBufLen: ", dwBufLen)
		// 断开长链接
		sendRemoteFinishedForFingerPrint <- true
	} else if dwType == 1 {
		progress := *(*int)(lpBuffer)
		fmt.Println("progress: ", progress)
	} else if dwType == 2 {
		// 回调类型
		getOrSet := *(*int)(pUserData)

		if getOrSet == 1 {
			cfg := *(hik.LPNET_DVR_FINGER_PRINT_CFG)(lpBuffer)
			fmt.Println("byCardNo: ", string(cfg.ByCardNo[:]))
			fmt.Println("dwFingerPrintLen: ", cfg.DwFingerPrintLen)
			fmt.Println("byFingerType: ", cfg.ByFingerType)
			fmt.Println("byFingerPrintID: ", cfg.ByFingerPrintID)
			// 写指纹到文件中
			writePrintData(cfg.ByFingerData[:], cfg.DwFingerPrintLen)
		} else if getOrSet == 2 {
			cfg := *(hik.LPNET_DVR_FINGER_PRINT_STATUS)(lpBuffer)
			fmt.Println("byCardNo: ", string(cfg.ByCardNo[:]))
			fmt.Println("byCardReaderRecvStatus: ", cfg.ByCardReaderRecvStatus)
		}
	} else {
		sendRemoteFinishedForFingerPrint <- false
	}
}

// 读取指纹数据
func readFingerPrintData(filePath string) []byte {
	imgBuffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		hangOnForFingerPrint <- false
		log.Fatal("Read fingerprint file failed", err)
	}
	return imgBuffer
}

// 写指纹到临时文件
func writePrintData(printData []byte, size uint32) {
	// 临时文件
	tmpFile, err := ioutil.TempFile(os.TempDir()+"/hik_finger", "*.bin")
	if err != nil {
		sendRemoteFinishedForFingerPrint <- false
		log.Fatal("Cannot create temporary file", err)
		return
	}

	// 写文件
	_, err = tmpFile.Write(printData[:size])
	if err != nil {
		sendRemoteFinishedForFingerPrint <- false
		log.Fatal("Cannot write temporary file", err)
	}
	fmt.Println("Output file: ", tmpFile.Name())
}
