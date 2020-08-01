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

// 获取和下发人脸例子
func FaceExample() {
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
	loginInfo.CbLoginResult = loginCallbackForFaceExample
	// 登陆
	result := hik.NET_DVR_Login_V40(&loginInfo, &deviceInfo)
	if -1 == result {
		printError("Login failed")
	}

	// hang on for testing async callback
	<-hangOnForFaceExample

	// 释放资源
	hik.NET_DVR_Cleanup()
}

var hangOnForFaceExample = make(chan bool)

// 登陆异步回调函数
func loginCallbackForFaceExample(lUserID int, dwResult uint32, lpDeviceInfo hik.LPNET_DVR_DEVICEINFO_V30, pUser unsafe.Pointer) {
	if 1 != dwResult {
		fmt.Println("异步登陆失败")
		printError("Login failed")
		return
	}

	fmt.Println("异步登陆成功，userId: ", strconv.Itoa(lUserID))
	fmt.Println("设备序列号: ", string(lpDeviceInfo.SSerialNumber[:]))

	// 回调函数复用，因此通过标记区分：1 - 获取人脸参数回调；2 - 设置人脸参数回调
	getOrSet := 2

	// 设置人脸参数
	setFaceInfo := hik.NET_DVR_FACE_PARAM_COND{}
	// 设置人数
	setFaceInfo.DwFaceNum = 1
	copy(setFaceInfo.ByCardNo[:], "123")
	setFaceInfo.DwSize = uint32(unsafe.Sizeof(setFaceInfo))
	setFaceInfo.ByFaceDataType = 1
	setFaceInfo.ByFaceID = 1
	setFaceInfo.ByEnableCardReader[0] = 1
	// 设置脸
	lHandle := hik.NET_DVR_StartRemoteConfig(
		lUserID,
		hik.NET_DVR_SET_FACE_PARAM_CFG,
		unsafe.Pointer(&setFaceInfo),
		unsafe.Sizeof(setFaceInfo),
		remoteConfigCallbackForFaceExample,
		unsafe.Pointer(&getOrSet),
	)

	// -1 失败，其它值为长连接句柄
	if -1 == lHandle {
		printError("Start remote failed")
		hangOnForFaceExample <- false
		return
	}

	// 下发人脸参数
	setFaceBuf := hik.NET_DVR_FACE_PARAM_CFG{}
	structSize := uint32(unsafe.Sizeof(setFaceBuf))
	setFaceBuf.DwSize = structSize
	copy(setFaceBuf.ByCardNo[:], "123")
	// 人脸数据类型：0 - 模板, 1 - 图片
	setFaceBuf.ByFaceDataType = 1
	setFaceBuf.ByEnableCardReader[0] = 1
	setFaceBuf.ByFaceID = 1

	// 读取人脸图片，最大支持 200k 的图片，否则接口报错
	pwd, _ := os.Getwd()
	imgBuffer := readFaceImg(pwd + "/res/myface.jpg")
	setFaceBuf.PFaceBuffer = unsafe.Pointer(&imgBuffer[0])
	setFaceBuf.DwFaceLen = uint32(len(imgBuffer))
	// 下发人脸数据
	success := hik.NET_DVR_SendRemoteConfig(
		lHandle,
		hik.ENUM_ACS_INTELLIGENT_IDENTITY_DATA,
		(*byte)(unsafe.Pointer(&setFaceBuf)),
		structSize,
	)
	if !success {
		printError("Send remote config failed")
		hangOnForFaceExample <- true
		return
	}

	<-sendRemoteFinishedForFaceExample

	fmt.Println("=========================")

	// 关闭长连接
	success = hik.NET_DVR_StopRemoteConfig(lHandle)
	if !success {
		printError("Stop remote config failed")
	}

	// 获取人脸
	faceInfo := hik.NET_DVR_FACE_PARAM_COND{}
	faceInfo.DwSize = uint32(unsafe.Sizeof(faceInfo))
	faceInfo.ByFaceDataType = 1
	faceInfo.ByFaceID = 0xff
	faceInfo.ByEnableCardReader[0] = 1
	// 指定卡号和获取数量
	faceInfo.DwFaceNum = 1
	copy(faceInfo.ByCardNo[:], "123")
	// 回调函数
	cb := remoteConfigCallbackForFaceExample

	// 回调函数复用，因此通过标记区分：1 - 获取人脸参数回调；2 - 设置人脸参数回调
	getOrSet = 1

	// 获取人脸数据
	lHandle = hik.NET_DVR_StartRemoteConfig(
		lUserID,
		hik.NET_DVR_GET_FACE_PARAM_CFG,
		unsafe.Pointer(&faceInfo),
		unsafe.Sizeof(faceInfo),
		cb,
		unsafe.Pointer(&getOrSet),
	)
	// -1 失败，其它值为长连接句柄
	if -1 == lHandle {
		printError("Start remote failed")
	}

	// 下发参数是否结束
	<-sendRemoteFinishedForFaceExample

	// 关闭长连接
	success = hik.NET_DVR_StopRemoteConfig(lHandle)
	if !success {
		printError("Stop remote config failed")
	}

	// 退出应用
	hangOnForFaceExample <- true
}

// 下发参数是否结束
var sendRemoteFinishedForFaceExample = make(chan bool, 1)

// 远程配置回调函数
func remoteConfigCallbackForFaceExample(dwType uint32, lpBuffer unsafe.Pointer, dwBufLen uint32, pUserData unsafe.Pointer) {
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
		sendRemoteFinishedForFaceExample <- true
	} else if dwType == 1 {
		progress := *(*int)(lpBuffer)
		fmt.Println("progress: ", progress)
	} else if dwType == 2 {
		// 回调类型
		getOrSet := *(*int)(pUserData)

		if getOrSet == 1 {
			faceCfg := *(hik.LPNET_DVR_FACE_PARAM_CFG)(lpBuffer)
			fmt.Println("byCardNo: ", string(faceCfg.ByCardNo[:]))
			fmt.Println("byFaceDataType: ", faceCfg.ByFaceDataType)
			fmt.Println("dwFaceLen: ", faceCfg.DwFaceLen)
			fmt.Println("byFaceID: ", faceCfg.ByFaceID)
			// 写人脸图片到临时文件
			writeFaceImg(faceCfg.PFaceBuffer, faceCfg.DwFaceLen)
		} else if getOrSet == 2 {
			faceCfg := *(hik.LPNET_DVR_FACE_PARAM_STATUS)(lpBuffer)
			fmt.Println("byCardNo: ", string(faceCfg.ByCardNo[:]))
		}
	} else {
		sendRemoteFinishedForFaceExample <- false
	}
}

// 读取人脸图片数据
func readFaceImg(filePath string) []byte {
	imgBuffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		hangOnForFaceExample <- false
		log.Fatal("Read face img file failed", err)
	}
	return imgBuffer
}

// 写人脸图片到临时文件
func writeFaceImg(faceBuffer unsafe.Pointer, size uint32) {
	buffer := hik.GoBytes(faceBuffer, size)

	// 临时文件
	tmpFile, err := ioutil.TempFile(os.TempDir()+"/hik_face", "*.jpg")
	if err != nil {
		sendRemoteFinishedForFaceExample <- false
		log.Fatal("Cannot create temporary file", err)
		return
	}

	// 写文件
	_, err = tmpFile.Write(buffer)
	if err != nil {
		sendRemoteFinishedForFaceExample <- false
		log.Fatal("Cannot write temporary file", err)
	}
	fmt.Println("Output file: ", tmpFile.Name())
}
