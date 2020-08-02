package main

import "C"
import "HCNetSDK-go/example"

func main() {
	example.CardExample()
	example.FaceExample()
	// 暂时不支持指纹
	//example.FingerprintExample()
	//example.FingerCapture()
}
