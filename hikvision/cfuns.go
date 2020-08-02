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

// 数组指针转成 go 切片
func GoBytes(pointer unsafe.Pointer, size uint32) []byte {
	return C.GoBytes(pointer, C.int(size))
}
