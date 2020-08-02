package hikvision

/*
// 定义系统:
// 1 = Windows
// 2 = Linux
// 3 = Darwin
#cgo windows CFLAGS: -DCGO_OS=1
#cgo linux 	 CFLAGS: -DCGO_OS=2
#cgo darwin  CFLAGS: -DCGO_OS=3

// 头文件
#cgo CFLAGS: -I../include
#include "HCNetSDK.h"

// 引入动态链接库
#if 1==CGO_OS
    #cgo LDFLAGS: -L${SRCDIR}/../lib/win64 -lHCNetSDK
#elif 2==CGO_OS || 3==CGO_OS
    #cgo LDFLAGS: -L${SRCDIR}/../lib/linux64 -lhcnetsdk
#else
    #error("Unknow OS")
#endif

//----------------- See `cfuns.go` file -----------------
void fLoginResultCallBack_cgo (int lUserID, unsigned int dwResult, void* lpDeviceInfo, void* pUser);
void fRemoteConfigCallback_cgo (DWORD dwType, void* lpBuffer, DWORD dwBufLen, void* pUserData);
BOOL msgCallBack_V31_cgo (LONG lCommand, void* pAlarmer, char *pAlarmInfo, DWORD dwBufLen, void* pUser);
*/
import "C"
import (
	"unsafe"
)

/************************* SDK 初始化 *************************/

// 设置SDK初始化参数。
// exapmle:
//    p := hikvision.NET_DVR_LOCAL_SDK_PATH{}
//    copy(p.SPath[:], "/usr/lib/hikvision")
//    result := hik.NET_DVR_SetSDKInitCfg(hikvision.NET_SDK_INIT_CFG_SDK_PATH, unsafe.Pointer(&p))
func NET_DVR_SetSDKInitCfg(enumType NET_SDK_INIT_CFG_TYPE, lpInBuff unsafe.Pointer) bool {
	return goBOOL(C.NET_DVR_SetSDKInitCfg(
		(C.NET_SDK_INIT_CFG_TYPE)(enumType),
		lpInBuff,
	))
}

// 初始化SDK，调用其他SDK函数的前提。
func NET_DVR_Init() bool {
	return goBOOL(C.NET_DVR_Init())
}

// 释放SDK资源，在程序结束之前调用。
func NET_DVR_Cleanup() bool {
	return goBOOL(C.NET_DVR_Cleanup())
}

/************************* SDK 本地功能 *************************/

//------------------------ 设置/获取SDK本地参数配置 ------------------------

// 设置SDK本地参数配置
// example:
//    p := hik.NET_DVR_LOCAL_TCP_PORT_BIND_CFG{}
//    p.WLocalBindTcpMinPort = 2000
//    p.WLocalBindTcpMaxPort = 3000
//    result := hik.NET_DVR_SetSDKLocalCfg(
//        hik.NET_SDK_LOCAL_CFG_TYPE_TCP_PORT_BIND,
//        unsafe.Pointer(&p),
//    )
func NET_DVR_SetSDKLocalCfg(enumType NET_SDK_LOCAL_CFG_TYPE, lpInBuff unsafe.Pointer) bool {
	return goBOOL(C.NET_DVR_SetSDKLocalCfg(
		(C.NET_SDK_LOCAL_CFG_TYPE)(enumType),
		lpInBuff,
	))
}

// 获取SDK本地参数配置
// example:
//    p := hik.NET_DVR_LOCAL_TCP_PORT_BIND_CFG{}
//    result := hik.NET_DVR_GetSDKLocalCfg(
//        hik.NET_SDK_LOCAL_CFG_TYPE_TCP_PORT_BIND,
//        unsafe.Pointer(&p),
//    )
func NET_DVR_GetSDKLocalCfg(enumType NET_SDK_LOCAL_CFG_TYPE, lpOutBuff unsafe.Pointer) bool {
	return goBOOL(C.NET_DVR_GetSDKLocalCfg(
		(C.NET_SDK_LOCAL_CFG_TYPE)(enumType),
		lpOutBuff,
	))
}

//------------------------ 连接、接收超时时间、重连设置 ------------------------

// 设置网络连接超时时间和连接尝试次数
func NET_DVR_SetConnectTime(dwWaitTime int32, dwTryTimes int32) bool {
	return goBOOL(C.NET_DVR_SetConnectTime(
		C.uint(dwWaitTime),
		C.uint(dwTryTimes),
	))
}

// 设置接收超时时间
func NET_DVR_SetRecvTimeOut(nRecvTimeOut int32) bool {
	return goBOOL(C.NET_DVR_SetRecvTimeOut(C.uint(nRecvTimeOut)))
}

// 设置重连功能
func NET_DVR_SetReconnect(dwInterval int32, bEnableRecon bool) bool {
	var reconnect C.BOOL = 1
	if !bEnableRecon {
		reconnect = 0
	}
	return goBOOL(C.NET_DVR_SetReconnect(C.uint(dwInterval), reconnect))
}

//------------------------ 多网卡绑定 ------------------------

// 获取所有IP，用于支持多网卡接口
// example:
//    ips := [16][16]byte{}
//    var num uint32 = 0
//    bind := false
//    result := hik.NET_DVR_GetLocalIP(&ips, &num, &bind)
func NET_DVR_GetLocalIP(strIP *[16][16]byte, pValidNum *uint32, pEnableBind *bool) bool {
	// GO byte 矩阵转成 C char 矩阵
	buf := [16][16]C.char{}
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			buf[i][j] = C.char(strIP[i][j])
		}
	}

	// 需要变量才能取地址
	var _pEnableBind C.BOOL = cBOOL(*pEnableBind)

	// 调用C
	result := goBOOL(C.NET_DVR_GetLocalIP(
		&buf[0],
		(*C.DWORD)(pValidNum),
		&_pEnableBind,
	))

	// 结果复制到 GO byte 矩阵中
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			strIP[i][j] = byte(buf[i][j])
		}
	}

	// 结果复制到变量中
	*pEnableBind = goBOOL(_pEnableBind)

	return result
}

// 选择使用哪个IP
func NET_DVR_SetValidIP(dwIPIndex uint32, bEnableBind bool) bool {
	return goBOOL(C.NET_DVR_SetValidIP(C.DWORD(dwIPIndex), cBOOL(bEnableBind)))
}

// 获取所有IP_V6，用于支持多网卡接口
// example:
// 参考 NET_DVR_GetLocalIP
func NET_DVR_GetLocalIPv6(strIP *[16][16]byte, pValidNum *uint32, pEnableBind *bool) bool {
	// GO byte 矩阵转成 C char 矩阵
	buf := [16][16]C.BYTE{}
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			buf[i][j] = C.BYTE(strIP[i][j])
		}
	}

	// 需要变量才能取地址
	var _pEnableBind C.BOOL = cBOOL(*pEnableBind)

	// 调用C
	result := goBOOL(C.NET_DVR_GetLocalIPv6(
		&buf[0],
		(*C.DWORD)(pValidNum),
		&_pEnableBind,
	))

	// 结果复制到 GO byte 矩阵中
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			strIP[i][j] = byte(buf[i][j])
		}
	}

	// 结果复制到变量中
	*pEnableBind = goBOOL(_pEnableBind)

	return result
}

// 选择使用哪个IP_V6
func NET_DVR_SetValidIPv6(dwIPIndex uint32, bEnableBind bool) bool {
	return goBOOL(C.NET_DVR_SetValidIPv6(C.DWORD(dwIPIndex), cBOOL(bEnableBind)))
}

//------------------------ SDK版本、状态、能力 ------------------------

// 获取SDK的版本信息
func NET_DVR_GetSDKVersion() uint32 {
	return uint32(C.NET_DVR_GetSDKVersion())
}

// 获取SDK的版本号和build信息
func NET_DVR_GetSDKBuildVersion() uint32 {
	return uint32(C.NET_DVR_GetSDKBuildVersion())
}

// 获取当前SDK的状态信息
// example:
//    p := hik.NET_DVR_SDKSTATE{}
//    result := hik.NET_DVR_GetSDKState(&p)
func NET_DVR_GetSDKState(pSDKState LPNET_DVR_SDKSTATE) bool {
	return goBOOL(C.NET_DVR_GetSDKState(
		C.LPNET_DVR_SDKSTATE(unsafe.Pointer(pSDKState)),
	))
}

// 获取当前SDK的功能信息
// example:
//    p := hik.NET_DVR_SDKABL{}
//    result := hik.NET_DVR_GetSDKAbility(&p)
func NET_DVR_GetSDKAbility(pSDKAbl LPNET_DVR_SDKABL) bool {
	return goBOOL(C.NET_DVR_GetSDKAbility(
		C.LPNET_DVR_SDKABL(unsafe.Pointer(pSDKAbl)),
	))
}

//------------------------ SDK启动写日志 ------------------------

// 启用写日志文件
// example:
//    usr, _ := user.Current()
//    result := hik.NET_DVR_SetLogToFile(3, usr.HomeDir + "/hik.log", true)
func NET_DVR_SetLogToFile(nLogLevel uint32, strLogDir string, bAutoDel bool) bool {
	b := []byte(strLogDir)
	return goBOOL(C.NET_DVR_SetLogToFile(
		C.DWORD(nLogLevel),
		(*C.char)(unsafe.Pointer(&b[0])),
		cBOOL(bAutoDel),
	))
}

//------------------------ 获取错误信息 ------------------------

// 返回最后操作的错误码
func NET_DVR_GetLastError() int {
	return int(C.NET_DVR_GetLastError())
}

// 返回最后操作的错误码信息
func NET_DVR_GetErrorMsg(pErrorNo int) string {
	code := C.LONG(pErrorNo)
	return string(C.GoString(
		C.NET_DVR_GetErrorMsg(&code),
	))
}

/************************* 用户注册 *************************/

// 通过解析服务器，获取设备的动态IP地址和端口号
// example:
//    sGetIP := [16]byte{}
//    var dwPort uint32
//    result := hik.NET_DVR_GetDVRIPByResolveSvr_EX(<your ipserver>, 7071, <device name>, <device serialNumber>, &sGetIP[0], &dwPort)
func NET_DVR_GetDVRIPByResolveSvr_EX(
	sServerIP string, wServerPort uint16, sDVRName string,
	sDVRSerialNumber string, sGetIP *byte, dwPort *uint32,
) bool {
	_sDVRName := []byte(sDVRName)
	_sDVRSerialNumber := []byte(sDVRSerialNumber)

	return goBOOL(C.NET_DVR_GetDVRIPByResolveSvr_EX(
		(*C.char)(unsafe.Pointer(&[]byte(sServerIP)[0])),
		C.WORD(wServerPort),
		(*C.BYTE)(unsafe.Pointer(&_sDVRName[0])),
		C.WORD(len(_sDVRName)),
		(*C.BYTE)(unsafe.Pointer(&_sDVRSerialNumber[0])),
		C.WORD(len(_sDVRSerialNumber)),
		(*C.char)(unsafe.Pointer(sGetIP)),
		(*C.DWORD)(unsafe.Pointer(dwPort)),
	))
}

// 激活设备
// example:
//    config := hik.NET_DVR_ACTIVATECFG{}
//    copy(config.SPassword[:], <initial password>)
//    config.DwSize = uint32(unsafe.Sizeof(config))
//    result := hik.NET_DVR_ActivateDevice("192.168.8.110", 8000, &config)
func NET_DVR_ActivateDevice(sDVRIP string, wDVRPort uint16, lpActivateCfg LPNET_DVR_ACTIVATECFG) bool {
	return goBOOL(C.NET_DVR_ActivateDevice(
		(*C.char)(unsafe.Pointer(&[]byte(sDVRIP)[0])),
		C.WORD(wDVRPort),
		C.LPNET_DVR_ACTIVATECFG(unsafe.Pointer(lpActivateCfg)),
	))
}

// 用户注册设备（支持异步登录）
// example:
//    loginInfo := hik.NET_DVR_USER_LOGIN_INFO{}
//	  deviceInfo := hik.NET_DVR_DEVICEINFO_V40{}
//
//    copy(loginInfo.SDeviceAddress[:], "192.168.8.110")
//    loginInfo.WPort = 8000
//    loginInfo.ByUseTransport = 0
//    loginInfo.BUseAsynLogin = 0
//    copy(loginInfo.SUserName[:], <your username>)
//    copy(loginInfo.SPassword[:], <your password>)
//    loginInfo.CbLoginResult = func(lUserID int, dwResult uint32, lpDeviceInfo hik.LPNET_DVR_DEVICEINFO_V30, pUser unsafe.Pointer) {
//        if 1 == dwResult {
//            fmt.Println("异步登陆成功，userId: ", strconv.Itoa(lUserID))
//            fmt.Println("设备序列号: ", string(lpDeviceInfo.SSerialNumber[:]))
//        } else {
//            fmt.Println("异步登陆失败")
//        }
//    }
//    result := hik.NET_DVR_Login_V40(&loginInfo, &deviceInfo)
func NET_DVR_Login_V40(pLoginInfo LPNET_DVR_USER_LOGIN_INFO, lpDeviceInfo LPNET_DVR_DEVICEINFO_V40) int {
	// 缓存回调函数
	loginChan <- pLoginInfo.CbLoginResult
	// 使用 C 的回调函数，此函数会调用 golang 的函数
	var _pLoginInfo C.LPNET_DVR_USER_LOGIN_INFO = C.LPNET_DVR_USER_LOGIN_INFO(unsafe.Pointer(pLoginInfo))
	(*_pLoginInfo).cbLoginResult = (C.fLoginResultCallBack)(C.fLoginResultCallBack_cgo)
	result := int(C.NET_DVR_Login_V40(
		_pLoginInfo,
		C.LPNET_DVR_DEVICEINFO_V40(unsafe.Pointer(lpDeviceInfo)),
	))
	return result
}

// 登陆回调函数的缓存通道
var loginChan = make(chan FLoginResultCallBack, 1)

//export fLoginResultCallBackGo
func fLoginResultCallBackGo(lUserID int, dwResult uint, lpDeviceInfo C.LPNET_DVR_DEVICEINFO_V30, pUser unsafe.Pointer) {
	if nil == lpDeviceInfo {
		(<-loginChan)(lUserID, uint32(dwResult), &NET_DVR_DEVICEINFO_V30{}, pUser)
		return
	}
	pDeviceInfo := (LPNET_DVR_DEVICEINFO_V30)(unsafe.Pointer(lpDeviceInfo))
	// 从通道中获取回调函数
	(<-loginChan)(lUserID, uint32(dwResult), pDeviceInfo, pUser)
}

// 用户注销
func NET_DVR_Logout(lUserID int) bool {
	return goBOOL(C.NET_DVR_Logout(C.LONG(lUserID)))
}

/************************* 网络参数配置 *************************/

// key: pUser's pointer
// value: cached function
var remoteConfigCbMap = make(map[uintptr]remoteConfigCbConf)

// 回调函数配置
type remoteConfigCbConf struct {
	cb FRemoteConfigCallback
	// pUserData 是否为内部初始化
	innerInitial bool
}

// 启动远程配置
// See https://open.hikvision.com/hardware/definitions/NET_DVR_StartRemoteConfig_ACS.html
func NET_DVR_StartRemoteConfig(lUserID int, dwCommand uint32, lpInBuffer unsafe.Pointer, dwInBufferLen uintptr, cbStateCallback FRemoteConfigCallback, pUserData unsafe.Pointer) int {
	cache := remoteConfigCbConf{}
	cache.cb = cbStateCallback
	// 初始化 pUserData，该值会在 cgo 回调函数中返回
	if nil == pUserData {
		cbTokenId := 0
		pUserData = unsafe.Pointer(&cbTokenId)
		cache.innerInitial = true
	}
	// 保存 golang 回调函数，通过 cgo 回调函数，通过 pUserData 的地址，获取和调用 golang 的函数
	index := uintptr(pUserData)
	remoteConfigCbMap[index] = cache

	return int(C.NET_DVR_StartRemoteConfig(
		C.LONG(lUserID),
		C.DWORD(dwCommand),
		C.LPVOID(lpInBuffer),
		C.DWORD(dwInBufferLen),
		(C.fRemoteConfigCallback)(C.fRemoteConfigCallback_cgo),
		C.LPVOID(pUserData),
	))
}

// 远程配置回调函数
//export fRemoteConfigCallbackGo
func fRemoteConfigCallbackGo(dwType uint32, lpBuffer unsafe.Pointer, dwBufLen uint32, pUserData unsafe.Pointer) {
	errMsg := "Not found `FRemoteConfigCallback` function in fRemoteConfigCallbackGo function."
	// 根据 pUserData 地址查询 golang 函数，如果为 nil 则无法找到
	if nil == pUserData {
		panic(errMsg)
	}
	// golang 定义的回调函数
	conf := remoteConfigCbMap[uintptr(pUserData)]
	// 回调函数不存在
	if nil == conf.cb {
		panic(errMsg)
	}
	// 内部初始化，则清空 pUserData
	if conf.innerInitial {
		pUserData = nil
	}

	conf.cb(dwType, lpBuffer, dwBufLen, pUserData)
}

// 发送长连接数据
// See https://open.hikvision.com/hardware/definitions/NET_DVR_SendRemoteConfig_ACS.html?_blank
func NET_DVR_SendRemoteConfig(lHandle int, dwDataType CFG_SEND_DATA_TYPE, pSendBuf *byte, dwBufSize uint32) bool {
	return goBOOL(C.NET_DVR_SendRemoteConfig(
		C.LONG(lHandle),
		C.DWORD(dwDataType),
		(*C.char)(unsafe.Pointer(pSendBuf)),
		C.DWORD(dwBufSize),
	))
}

// 关闭长连接配置接口所创建的句柄，释放资源
func NET_DVR_StopRemoteConfig(lHandle int) bool {
	return goBOOL(C.NET_DVR_StopRemoteConfig(C.LONG(lHandle)))
}

/************************* 报警布防 *************************/

// 注册回调函数，接收设备报警消息等
func NET_DVR_SetDVRMessageCallBack_V31(fMessageCallBack MSGCallBack_V31, pUser unsafe.Pointer) bool {
	cache := dvrMsgCbConf{}
	cache.cb = fMessageCallBack
	// 初始化 pUserData，该值会在 cgo 回调函数中返回
	if nil == pUser {
		cbTokenId := 0
		pUser = unsafe.Pointer(&cbTokenId)
		cache.innerInitial = true
	}
	// 保存 golang 回调函数，通过 cgo 回调函数，通过 pUserData 的地址，获取和调用 golang 的函数
	index := uintptr(pUser)
	dvrMsgCbMap[index] = cache

	return goBOOL(C.NET_DVR_SetDVRMessageCallBack_V31(
		(C.MSGCallBack_V31)(C.msgCallBack_V31_cgo),
		pUser,
	))
}

// key: pUser's pointer
// value: cached function
var dvrMsgCbMap = make(map[uintptr]dvrMsgCbConf)

// 回调函数配置
type dvrMsgCbConf struct {
	cb MSGCallBack_V31
	// pUserData 是否为内部初始化
	innerInitial bool
}

// 报警布防回调函数
//export msgCallBack_V31_Go
func msgCallBack_V31_Go(lCommand int, pAlarmer C.LPNET_DVR_ALARMER, pAlarmInfo *byte, dwBufLen uint32, pUser unsafe.Pointer) C.BOOL {
	errMsg := "Not found `MSGCallBack_V31` function in msgCallBack_V31_Go function."
	// 根据 pUserData 地址查询 golang 函数，如果为 nil 则无法找到
	if nil == pUser {
		panic(errMsg)
	}
	// golang 定义的回调函数
	conf := dvrMsgCbMap[uintptr(pUser)]
	// 回调函数不存在
	if nil == conf.cb {
		panic(errMsg)
	}
	// 内部初始化，则清空 pUserData
	if conf.innerInitial {
		pUser = nil
	}

	_pAlarmer := (LPNET_DVR_ALARMER)(unsafe.Pointer(pAlarmer))

	return cBOOL(conf.cb(lCommand, _pAlarmer, pAlarmInfo, dwBufLen, pUser))
}

// 建立报警上传通道，获取报警等信息
func NET_DVR_SetupAlarmChan_V41(lUserID int, lpSetupParam LPNET_DVR_SETUPALARM_PARAM) int {
	return int(C.NET_DVR_SetupAlarmChan_V41(
		C.LONG(lUserID),
		C.LPNET_DVR_SETUPALARM_PARAM(unsafe.Pointer(lpSetupParam)),
	))
}

// 撤销报警上传通道
func NET_DVR_CloseAlarmChan_V30(lAlarmHandle int) bool {
	return goBOOL(C.NET_DVR_CloseAlarmChan_V30(C.LONG(lAlarmHandle)))
}

/************************* 远程反控 *************************/

// 远程门禁控制或梯控控制
func NET_DVR_ControlGateway(lUserID int, lGatewayIndex int, dwStaic uint32) bool {
	return goBOOL(C.NET_DVR_ControlGateway(
		C.LONG(lUserID),
		C.LONG(lGatewayIndex),
		C.DWORD(dwStaic),
	))
}

// 辅助功能控制
func NET_DVR_AlarmHostAssistantControl(lUserID int, dwType uint32, dwNumber uint32, dwCmdParam uint32) bool {
	return goBOOL(C.NET_DVR_AlarmHostAssistantControl(
		C.LONG(lUserID),
		C.DWORD(dwType),
		C.DWORD(dwNumber),
		C.DWORD(dwCmdParam),
	))
}

/************************* 获取能力集 *************************/

// 获取能力集
func NET_DVR_GetDeviceAbility(lUserID int, dwAbilityType uint32, pInBuf string, pOutBuf unsafe.Pointer, dwOutLength uint32) bool {
	_pInBuf := []byte(pInBuf)
	return goBOOL(C.NET_DVR_GetDeviceAbility(
		C.LONG(lUserID),
		C.DWORD(dwAbilityType),
		(*C.char)(unsafe.Pointer(&_pInBuf[0])),
		C.DWORD(len(pInBuf)),
		(*C.char)(pOutBuf),
		C.DWORD(dwOutLength),
	))
}
