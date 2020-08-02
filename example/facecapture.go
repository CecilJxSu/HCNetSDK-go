package example

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

// 人脸采集例子
func FaceCaptureExample() {
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
	loginInfo.CbLoginResult = loginCallbackForFaceCapture
	// 登陆
	result := hik.NET_DVR_Login_V40(&loginInfo, &deviceInfo)
	if -1 == result {
		printError("Login failed")
		hangOnForFaceCapture <- false
	}

	// hang on for testing async callback
	<-hangOnForFaceCapture

	// 释放资源
	hik.NET_DVR_Cleanup()
}

var hangOnForFaceCapture = make(chan bool)

// 登陆异步回调函数
func loginCallbackForFaceCapture(lUserID int, dwResult uint32, lpDeviceInfo hik.LPNET_DVR_DEVICEINFO_V30, pUser unsafe.Pointer) {
	if 1 != dwResult {
		fmt.Println("异步登陆失败")
		printError("Login failed")
		hangOnForFaceCapture <- false
		return
	}

	fmt.Println("异步登陆成功，userId: ", strconv.Itoa(lUserID))
	fmt.Println("设备序列号: ", string(lpDeviceInfo.SSerialNumber[:]))

	// 采集人脸参数
	captureFaceInfo := hik.NET_DVR_CAPTURE_FACE_COND{}
	captureFaceInfo.DwSize = uint32(unsafe.Sizeof(captureFaceInfo))
	// 采集人脸
	lHandle := hik.NET_DVR_StartRemoteConfig(
		lUserID,
		hik.NET_DVR_CAPTURE_FACE_INFO,
		unsafe.Pointer(&captureFaceInfo),
		unsafe.Sizeof(captureFaceInfo),
		remoteConfigCallbackForFaceCapture,
		nil,
	)

	// -1 失败，其它值为长连接句柄
	if -1 == lHandle {
		printError("Start remote failed")
		hangOnForFaceCapture <- false
		return
	}

	<-sendRemoteFinishedForFaceCapture

	// 关闭长连接
	success := hik.NET_DVR_StopRemoteConfig(lHandle)
	if !success {
		printError("Stop remote config failed")
	}

	// 退出应用
	hangOnForFaceCapture <- true
}

// 下发参数是否结束
var sendRemoteFinishedForFaceCapture = make(chan bool, 1)

// 远程配置回调函数
func remoteConfigCallbackForFaceCapture(dwType uint32, lpBuffer unsafe.Pointer, dwBufLen uint32, pUserData unsafe.Pointer) {
	fmt.Println("dwType: ", dwType)
	if dwType == 0 {
		dwStatusArr := *(*[4]byte)(lpBuffer)
		dwStatus := binary.LittleEndian.Uint32(dwStatusArr[:])
		fmt.Println("dwStatus: ", dwStatus)
		fmt.Println("dwBufLen: ", dwBufLen)
		// 断开长链接
		sendRemoteFinishedForFaceCapture <- true
	} else if dwType == 1 {
		progress := *(*int)(lpBuffer)
		fmt.Println("progress: ", progress)
	} else if dwType == 2 {
		captureResult := *(hik.LPNET_DVR_CAPTURE_FACE_CFG)(lpBuffer)
		fmt.Println("byCaptureProgress: ", captureResult.ByCaptureProgress)
		fmt.Println("byFacePicQuality: ", captureResult.ByFacePicQuality)
		// 写人脸图片到临时文件
		writeCaptureFace(unsafe.Pointer(captureResult.PFacePicBuffer), captureResult.DwFacePicSize)
	} else {
		sendRemoteFinishedForFaceCapture <- false
	}
}

// 写人脸图片到临时文件
func writeCaptureFace(faceBuffer unsafe.Pointer, size uint32) {
	buffer := hik.GoBytes(faceBuffer, size)

	// 目录
	dir := os.TempDir() + "/hik_face/capture"
	err := os.MkdirAll(dir, os.ModePerm)
	if nil != err {
		sendRemoteFinishedForFaceCapture <- false
		log.Fatal("Cannot create temporary dir", err)
		return
	}

	// 临时文件
	tmpFile, err := ioutil.TempFile(dir, "*.jpg")
	if err != nil {
		sendRemoteFinishedForFaceCapture <- false
		log.Fatal("Cannot create temporary file", err)
		return
	}

	// 写文件
	_, err = tmpFile.Write(buffer)
	if err != nil {
		sendRemoteFinishedForFaceCapture <- false
		log.Fatal("Cannot write temporary file", err)
	}
	fmt.Println("Output file: ", tmpFile.Name())
}
