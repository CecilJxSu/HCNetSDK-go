package hikvision

import "C"
import (
	"unsafe"
)

/************************* SDK 初始化 *************************/

// enum NET_SDK_INIT_CFG_TYPE
type NET_SDK_INIT_CFG_TYPE int32

const (
	NET_SDK_INIT_CFG_TYPE_CHECK_MODULE_COM NET_SDK_INIT_CFG_TYPE = 0 // 增加对必须库的检查
	NET_SDK_INIT_CFG_ABILITY               NET_SDK_INIT_CFG_TYPE = 1 // sdk支持的业务的能力集
	NET_SDK_INIT_CFG_SDK_PATH              NET_SDK_INIT_CFG_TYPE = 2 // 设置组件库加载路径

	// 设置OpenSSL的libeay32.dll/libcrypto.so/libcrypto.dylib所在路径
	NET_SDK_INIT_CFG_LIBEAY_PATH NET_SDK_INIT_CFG_TYPE = 3

	// 设置OpenSSL的ssleay32.dll/libssl.so/libssl.dylib所在路径
	NET_SDK_INIT_CFG_SSLEAY_PATH NET_SDK_INIT_CFG_TYPE = 4
)

// enum INIT_CFG_MAX_NUM
type INIT_CFG_MAX_NUM int32

const (
	INIT_CFG_NUM_2048  INIT_CFG_MAX_NUM = 2048  // 2048路
	INIT_CFG_NUM_5120  INIT_CFG_MAX_NUM = 5120  // 5120路
	INIT_CFG_NUM_10240 INIT_CFG_MAX_NUM = 10240 // 10240路
	INIT_CFG_NUM_15360 INIT_CFG_MAX_NUM = 15360 // 15360路
	INIT_CFG_NUM_20480 INIT_CFG_MAX_NUM = 20480 // 20480路
)

// 初始化能力参数结构体
type NET_DVR_INIT_CFG_ABILITY struct {
	EnumMaxLoginUsersNum INIT_CFG_MAX_NUM // 最多允许的注册用户个数。
	EnumMaxAlarmNum      INIT_CFG_MAX_NUM // 最大的告警路数
	ByRes                [64]byte         // 保留
}
type LPNET_DVR_INIT_CFG_ABILITY *NET_DVR_INIT_CFG_ABILITY

// 组件库加载路径信息结构体
type NET_DVR_LOCAL_SDK_PATH struct {
	SPath [NET_SDK_MAX_FILE_PATH]byte // 组件库地址
	ByRes [128]byte                   // 保留
}
type LPNET_DVR_LOCAL_SDK_PATH *NET_DVR_LOCAL_SDK_PATH

/************************* SDK 本地功能 *************************/
//------------------------ SDK本地参数配置 ------------------------

// enum NET_SDK_LOCAL_CFG_TYPE
type NET_SDK_LOCAL_CFG_TYPE int

const (
	NET_SDK_LOCAL_CFG_TYPE_TCP_PORT_BIND       NET_SDK_LOCAL_CFG_TYPE = 0  //本地TCP端口绑定配置，对应结构体NET_DVR_LOCAL_TCP_PORT_BIND_CFG
	NET_SDK_LOCAL_CFG_TYPE_UDP_PORT_BIND       NET_SDK_LOCAL_CFG_TYPE = 1  //本地UDP端口绑定配置，对应结构体NET_DVR_LOCAL_UDP_PORT_BIND_CFG
	NET_SDK_LOCAL_CFG_TYPE_MEM_POOL            NET_SDK_LOCAL_CFG_TYPE = 2  //内存池本地配置，对应结构体NET_DVR_LOCAL_MEM_POOL_CFG
	NET_SDK_LOCAL_CFG_TYPE_MODULE_RECV_TIMEOUT NET_SDK_LOCAL_CFG_TYPE = 3  //按模块配置超时时间，对应结构体NET_DVR_LOCAL_MODULE_RECV_TIMEOUT_CFG
	NET_SDK_LOCAL_CFG_TYPE_ABILITY_PARSE       NET_SDK_LOCAL_CFG_TYPE = 4  //是否使用能力集解析库，对应结构体NET_DVR_LOCAL_ABILITY_PARSE_CFG
	NET_SDK_LOCAL_CFG_TYPE_TALK_MODE           NET_SDK_LOCAL_CFG_TYPE = 5  //对讲模式，对应结构体NET_DVR_LOCAL_TALK_MODE_CFG
	NET_SDK_LOCAL_CFG_TYPE_PROTECT_KEY         NET_SDK_LOCAL_CFG_TYPE = 6  //密钥设置，对应结构体NET_DVR_LOCAL_PROTECT_KEY_CFG
	NET_SDK_LOCAL_CFG_TYPE_CFG_VERSION         NET_SDK_LOCAL_CFG_TYPE = 7  //用于测试版本头的设备端兼容情NET_DVR_LOCAL_MEM_POOL_CFG况, 只有在设置参数时才起作用。
	NET_SDK_LOCAL_CFG_TYPE_RTSP_PARAMS         NET_SDK_LOCAL_CFG_TYPE = 8  //rtsp参数配置，对于结构体NET_DVR_RTSP_PARAMS_CFG
	NET_SDK_LOCAL_CFG_TYPE_SIMXML_LOGIN        NET_SDK_LOCAL_CFG_TYPE = 9  //在登录时使用模拟能力补充support字段, 对应结构NET_DVR_SIMXML_LOGIN
	NET_SDK_LOCAL_CFG_TYPE_CHECK_DEV           NET_SDK_LOCAL_CFG_TYPE = 10 //心跳交互间隔时间
	NET_SDK_LOCAL_CFG_TYPE_SECURITY            NET_SDK_LOCAL_CFG_TYPE = 11 //SDK本次安全配置，
	NET_SDK_LOCAL_CFG_TYPE_EZVIZLIB_PATH       NET_SDK_LOCAL_CFG_TYPE = 12 //配置萤石云通信库地址，
	NET_SDK_LOCAL_CFG_TYPE_CHAR_ENCODE         NET_SDK_LOCAL_CFG_TYPE = 13 //13.配置字符编码相关处理回调
	NET_SDK_LOCAL_CFG_TYPE_PROXYS              NET_SDK_LOCAL_CFG_TYPE = 14 //设置获取代理
	NET_DVR_LOCAL_CFG_TYPE_LOG                 NET_SDK_LOCAL_CFG_TYPE = 15 //日志参数配置  NET_DVR_LOCAL_LOG_CFG
	NET_DVR_LOCAL_CFG_TYPE_STREAM_CALLBACK     NET_SDK_LOCAL_CFG_TYPE = 16 //码流回调参数配置 NET_DVR_LOCAL_STREAM_CALLBACK_CFG
	NET_DVR_LOCAL_CFG_TYPE_GENERAL             NET_SDK_LOCAL_CFG_TYPE = 17 //通用参数配置 NET_DVR_LOCAL_GENERAL_CFG
	NET_DVR_LOCAL_CFG_TYPE_PTZ                 NET_SDK_LOCAL_CFG_TYPE = 18 //PTZ是否接收设备返回配置
	NET_DVR_LOCAL_CFG_MESSAGE_CALLBACK_V51     NET_SDK_LOCAL_CFG_TYPE = 19 //报警V51回调相关本地配置,对应结构体为NET_DVR_MESSAGE_CALLBACK_PARAM_V51 。(仅对NET_DVR_SetDVRMessageCallBack_V51以上版本有效)
	NET_SDK_LOCAL_CFG_CERTIFICATION            NET_SDK_LOCAL_CFG_TYPE = 20 //配置和证书相关的参数，对应结构体结构体NET_DVR_LOCAL_CERTIFICATION
	NET_SDK_LOCAL_CFG_PORT_MULTIPLEX           NET_SDK_LOCAL_CFG_TYPE = 21 //端口复用，对应结构体NET_DVR_LOCAL_PORT_MULTI_CFG
	NET_SDK_LOCAL_CFG_ASYNC                    NET_SDK_LOCAL_CFG_TYPE = 22 //异步配置，对应结构体NET_DVR_LOCAL_ASYNC_CFG
)

// 本地TCP端口绑定配置结构体
type NET_DVR_LOCAL_TCP_PORT_BIND_CFG struct {
	WLocalBindTcpMinPort int16    //本地绑定Tcp最小端口
	WLocalBindTcpMaxPort int16    //本地绑定Tcp最大端口
	ByRes                [60]byte //保留
}
type LPNET_DVR_LOCAL_TCP_PORT_BIND_CFG *NET_DVR_LOCAL_TCP_PORT_BIND_CFG

// 本地UDP端口绑定配置结构体
type NET_DVR_LOCAL_UDP_PORT_BIND_CFG struct {
	WLocalBindUdpMinPort int16    //本地绑定Udp最小端口
	WLocalBindUdpMaxPort int16    //本地绑定Udp最大端口
	ByRes                [60]byte //保留
}
type LPNET_DVR_LOCAL_UDP_PORT_BIND_CFG *NET_DVR_LOCAL_UDP_PORT_BIND_CFG

// 内存池本地配置结构体
type NET_DVR_LOCAL_MEM_POOL_CFG struct {
	DwAlarmMaxBlockNum      int32     //报警模块内存池最多向系统申请的内存块（block）个数，每个block为64MB, 超过这个上限则不向系统申请，0表示无上限
	DwAlarmReleaseInterval  int32     //报警模块空闲内存释放的间隔，单位秒，为0表示不释放
	DwObjectReleaseInterval int32     //对象申请模块空闲内存释放的间隔，单位秒，为0表示不释放
	ByRes                   [508]byte //保留
}
type LPNET_DVR_LOCAL_MEM_POOL_CFG *NET_DVR_LOCAL_MEM_POOL_CFG

// 按模块配置接收超时时间
type NET_DVR_LOCAL_MODULE_RECV_TIMEOUT_CFG struct {
	DwPreviewTime int32     //预览模块超时时间
	DwAlarmTime   int32     //报警模块超时时间
	DwVodTime     int32     //回放模块超时时间
	DwElse        int32     //其他模块
	ByRes         [512]byte //保留
}
type LPNET_DVR_LOCAL_MODULE_RECV_TIMEOUT_CFG *NET_DVR_LOCAL_MODULE_RECV_TIMEOUT_CFG

// 能力集解析库配置结构体
type NET_DVR_LOCAL_ABILITY_PARSE_CFG struct {
	ByEnableAbilityParse byte //使用能力集解析库,0-不使用,1-使用,默认不使用
	ByRes                [127]byte
}
type LPNET_DVR_LOCAL_ABILITY_PARSE_CFG *NET_DVR_LOCAL_ABILITY_PARSE_CFG

// 对讲模式配置结构体
type NET_DVR_LOCAL_TALK_MODE_CFG struct {
	ByTalkMode byte //对讲模式，0-使用对讲库（默认），1-使用windows api模式
	ByRes      [127]byte
}
type LPNET_DVR_LOCAL_TALK_MODE_CFG *NET_DVR_LOCAL_TALK_MODE_CFG

// 设备在线巡检参数结构体
type NET_DVR_LOCAL_CHECK_DEV struct {
	DwCheckOnlineTimeout    int32 //巡检时间间隔，单位ms  最小值为30s，最大值120s。为0时，表示用默认值(120s)
	DwCheckOnlineNetFailMax int32 //由于网络原因失败的最大累加次数；超过该值SDK才回调用户异常，为0时，表示使用默认值1
	ByRes                   [256]byte
}
type LPNET_DVR_LOCAL_CHECK_DEV *NET_DVR_LOCAL_CHECK_DEV

// 字符编码转换参数结构体
type NET_DVR_LOCAL_BYTE_ENCODE_CONVERT struct {
	FnCharConvertCallBack CHAR_ENCODE_CONVERT
	ByRes                 [256]byte
}
type LPNET_DVR_LOCAL_BYTE_ENCODE_CONVERT *NET_DVR_LOCAL_BYTE_ENCODE_CONVERT

// 字符编码转换回调函数
type CHAR_ENCODE_CONVERT func(*byte, int32, int32, *byte, int32, int32) int

// SDK日志参数配置结构体
type NET_DVR_LOCAL_LOG_CFG struct {
	WSDKLogNum int16     //sdk在覆盖模式下，日志生成的个数 0为默认值（10个）
	ByRes      [254]byte //保留
}
type LPNET_DVR_LOCAL_LOG_CFG *NET_DVR_LOCAL_LOG_CFG

// 通用参数配置结构体
type NET_DVR_LOCAL_GENERAL_CFG struct {
	ByExceptionCbDirectly      byte      //0-通过线程池异常回调，1-直接异常回调给上层
	ByNotSplitRecordFile       byte      //回放和预览中保存到本地录像文件不切片 0-默认切片，1-不切片
	ByResumeUpgradeEnable      byte      //断网续传升级使能，0-关闭（默认），1-开启
	ByAlarmJsonPictureSeparate byte      //控制JSON透传报警数据和图片是否分离，0-不分离，1-分离（分离后走COMM_ISAPI_ALARM回调返回）
	ByRes                      [4]byte   //保留
	I64FileSize                uint64    //单位：Byte
	DwResumeUpgradeTimeout     int32     //断网续传重连超时时间，单位毫秒
	ByAlarmReconnectMode       byte      //0-独立线程重连（默认） 1-线程池重连
	ByStdXmlBufferSize         byte      //设置ISAPI透传接收缓冲区大小，1-1M 其他-默认
	ByMultiplexing             byte      //0-普通链接（非TLS链接）关闭多路复用，1-普通链接（非TLS链接）开启多路复用
	ByRes1                     [233]byte //预留
}
type LPNET_DVR_LOCAL_GENERAL_CFG *NET_DVR_LOCAL_GENERAL_CFG

// PTZ是否交互配置参数结构体
type NET_DVR_LOCAL_PTZ_CFG struct {
	ByWithoutRecv byte     //是否交互配置，0-接收设备返回，1-不接收设备返回
	ByRes         [63]byte //保留，置为0
}
type LPNET_DVR_LOCAL_PTZ_CFG *NET_DVR_LOCAL_PTZ_CFG

// 证书校验参数配置结构体
type NET_DVR_LOCAL_CERTIFICATION struct {
	SzLoadPath [MAX_FILE_PATH_LEN]byte
	FnCB       FnCertVerifyResultCallBack
	PUserData  unsafe.Pointer
	ByRes      [64]byte
}
type LPNET_DVR_LOCAL_CERTIFICATION *NET_DVR_LOCAL_CERTIFICATION

// 证书校验回调函数
type FnCertVerifyResultCallBack func(int32, LPNET_DVR_CETTIFICATE_INFO, *byte) bool

// 证书校验结构体
type NET_DVR_CETTIFICATE_INFO struct {
	DwSize        int32
	SzIssuer      [MAX_CERTIFICATE_ISSUER_LEN]byte //证书颁发者
	StruStartTime NET_DVR_TIME
	StruEndTime   NET_DVR_TIME
	ByRes1        [1024]byte
}
type LPNET_DVR_CETTIFICATE_INFO *NET_DVR_CETTIFICATE_INFO

// 校时结构参数
type NET_DVR_TIME struct {
	DwYear   int32 //年
	DwMonth  int32 //月
	DwDay    int32 //日
	DwHour   int32 //时
	DwMinute int32 //分
	DwSecond int32 //秒
}
type LPNET_DVR_TIME *NET_DVR_TIME

// 端口复用参数配置结构体
type NET_DVR_LOCAL_PORT_MULTI_CFG struct {
	bEnable int //端口复用使能，true-开启
	byRes   [60]byte
}
type LPNET_DVR_LOCAL_PORT_MULTI_CFG *NET_DVR_LOCAL_PORT_MULTI_CFG

//------------------------ SDK版本、状态、能力 ------------------------

//SDK状态信息(9000新增)
type NET_DVR_SDKSTATE struct {
	dwTotalLoginNum         uint32 //当前login用户数
	dwTotalRealPlayNum      uint32 //当前realplay路数
	dwTotalPlayBackNum      uint32 //当前回放或下载路数
	dwTotalAlarmChanNum     uint32 //当前建立报警通道路数
	dwTotalFormatNum        uint32 //当前硬盘格式化路数
	dwTotalFileSearchNum    uint32 //当前日志或文件搜索路数
	dwTotalLogSearchNum     uint32 //当前日志或文件搜索路数
	dwTotalSerialNum        uint32 //当前透明通道路数
	dwTotalUpgradeNum       uint32 //当前升级路数
	dwTotalVoiceComNum      uint32 //当前语音转发路数
	dwTotalBroadCastNum     uint32 //当前语音广播路数
	dwTotalListenNum        uint32 //当前网络监听路数
	dwEmailTestNum          uint32 //当前邮件计数路数
	dwBackupNum             uint32 // 当前文件备份路数
	dwTotalInquestUploadNum uint32 //当前审讯上传路数
	dwRes                   [6]uint32
}
type LPNET_DVR_SDKSTATE *NET_DVR_SDKSTATE

//SDK功能支持信息(9000新增)
type NET_DVR_SDKABL struct {
	dwMaxLoginNum      uint32 //最大login用户数 MAX_LOGIN_USERS
	dwMaxRealPlayNum   uint32 //最大realplay路数 WATCH_NUM
	dwMaxPlayBackNum   uint32 //最大回放或下载路数 WATCH_NUM
	dwMaxAlarmChanNum  uint32 //最大建立报警通道路数 ALARM_NUM
	dwMaxFormatNum     uint32 //最大硬盘格式化路数 SERVER_NUM
	dwMaxFileSearchNum uint32 //最大文件搜索路数 SERVER_NUM
	dwMaxLogSearchNum  uint32 //最大日志搜索路数 SERVER_NUM
	dwMaxSerialNum     uint32 //最大透明通道路数 SERVER_NUM
	dwMaxUpgradeNum    uint32 //最大升级路数 SERVER_NUM
	dwMaxVoiceComNum   uint32 //最大语音转发路数 SERVER_NUM
	dwMaxBroadCastNum  uint32 //最大语音广播路数 MAX_CASTNUM
	dwRes              [10]uint32
}
type LPNET_DVR_SDKABL *NET_DVR_SDKABL
