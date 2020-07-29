package hikvision

/*
// 头文件
#cgo CFLAGS: -I../include
#include "HCNetSDK.h"
#include <stdio.h>

// See https://github.com/golang/go/wiki/cgo#function-pointer-callbacks

void fLoginResultCallBack_cgo (int lUserID, unsigned int dwResult, LPNET_DVR_DEVICEINFO_V30 lpDeviceInfo, void* pUser) {
    void fLoginResultCallBackGo(int, unsigned int, void*, void*);
    fLoginResultCallBackGo(lUserID, dwResult, lpDeviceInfo, pUser);
}

void fRemoteConfigCallback_cgo (DWORD dwType, void* lpBuffer, DWORD dwBufLen, void* pUserData) {
    void fRemoteConfigCallbackGo(DWORD, void*, DWORD, void*);
    fRemoteConfigCallbackGo(dwType, lpBuffer, dwBufLen, pUserData);
}
*/
import "C"
import "unsafe"

func convert_NET_DVR_DEVICEINFO_V30(cResult C.NET_DVR_DEVICEINFO_V30) NET_DVR_DEVICEINFO_V30 {
	goResult := NET_DVR_DEVICEINFO_V30{}

	copy(goResult.SSerialNumber[:], (*[SERIALNO_LEN]byte)(unsafe.Pointer(&cResult.sSerialNumber[0]))[:])
	goResult.ByAlarmInPortNum = byte(cResult.byAlarmInPortNum)
	goResult.ByAlarmOutPortNum = byte(cResult.byAlarmOutPortNum)
	goResult.ByDiskNum = byte(cResult.byDiskNum)
	goResult.ByDVRType = byte(cResult.byDVRType)
	goResult.ByChanNum = byte(cResult.byChanNum)
	goResult.ByStartChan = byte(cResult.byStartChan)
	goResult.ByAudioChanNum = byte(cResult.byAudioChanNum)
	goResult.ByIPChanNum = byte(cResult.byIPChanNum)
	goResult.ByZeroChanNum = byte(cResult.byZeroChanNum)
	goResult.ByMainProto = byte(cResult.byMainProto)
	goResult.BySubProto = byte(cResult.bySubProto)
	goResult.BySupport = byte(cResult.bySupport)
	goResult.BySupport1 = byte(cResult.bySupport1)
	goResult.BySupport2 = byte(cResult.bySupport2)
	goResult.WDevType = uint16(cResult.wDevType)
	goResult.BySupport3 = byte(cResult.wDevType)
	goResult.ByMultiStreamProto = byte(cResult.byMultiStreamProto)
	goResult.ByStartDChan = byte(cResult.byStartDChan)
	goResult.ByStartDTalkChan = byte(cResult.byStartDTalkChan)
	goResult.ByHighDChanNum = byte(cResult.byHighDChanNum)
	goResult.BySupport4 = byte(cResult.bySupport4)
	goResult.ByLanguageType = byte(cResult.byLanguageType)
	goResult.ByVoiceInChanNum = byte(cResult.byVoiceInChanNum)
	goResult.ByStartVoiceInChanNo = byte(cResult.byStartVoiceInChanNo)
	goResult.BySupport5 = byte(cResult.bySupport5)
	goResult.BySupport6 = byte(cResult.bySupport6)
	goResult.ByMirrorChanNum = byte(cResult.byMirrorChanNum)
	goResult.WStartMirrorChanNo = uint16(cResult.wStartMirrorChanNo)
	goResult.BySupport7 = byte(cResult.bySupport7)
	goResult.ByRes2 = byte(cResult.byRes2)
	return goResult
}

/******************* golang and cgo type convert each other *******************/
// C.BOOL --> go bool
func goBOOL(flag C.BOOL) bool {
	if flag == 1 {
		return true
	}
	return false
}

// go bool --> C.BOOL
func cBOOL(flag bool) C.BOOL {
	if flag {
		return C.BOOL(1)
	}
	return C.BOOL(0)
}
