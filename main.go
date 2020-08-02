package main

import "C"
import "HCNetSDK-go/example"

func main() {
	example.CardExample()
	example.FaceCaptureExample()
	example.FaceExample()
	// 暂时不支持指纹
	//example.FingerprintExample()
	//example.FingerCaptureExample()
}
