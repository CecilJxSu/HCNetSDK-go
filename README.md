# HCNetSDK-go
海康威视

# 配置
+ 拷贝海康威视提供的 linux64 sdk 的 lib 文件到本项目的 lib/linux64 下，文件清单：
```
HCNetSDKCom/
    libanalyzedata.so
    libHCAlarm.so
    libHCCoreDevCfg.so
    libHCDisplay.so
    libHCGeneralCfgMgr.so
    libHCIndustry.so
    libHCPlayBack.so
    libHCPreview.so
    libHCVoiceTalk.so
    libiconv2.so
    libStreamTransClient.so
    libSystemTransform.so
libAudioRender.so
libcrypto.so
libcrypto.so.1.0.0
libHCCore.so
libhcnetsdk.so
libhpr.so
libNPQos.so
libopenal.so.1
libPlayCtrl.so
libssl.so
libSuperRender.so
libz.so
```
注意：以上动态链接库并不是都会用上，不同功能使用不同的动态链接库

+ 放置动态链接库后，进入本项目，即可编译：
```bash
go build .
```

+ 运行时，需要将动态链接库拷贝到 `/usr/lib` 或 `/usr/local/lib` 等目录下
+ 或者在本项目根目录下执行以下命令，配置环境变量，指定依赖库的目录。或者在 `~/.profile` 或 `~/.bashrc`等文件上配置此环境变量：
```bash
export LD_LIBRARY_PATH=`pwd`/lib/linux64:`pwd`/lib/linux64/HCNetSDKCom
```

+ 指定依赖库位置后，即可运行：
```bash
go run .
```