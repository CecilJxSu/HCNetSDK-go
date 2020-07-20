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
	BEnable int //端口复用使能，true-开启
	ByRes   [60]byte
}
type LPNET_DVR_LOCAL_PORT_MULTI_CFG *NET_DVR_LOCAL_PORT_MULTI_CFG

//------------------------ SDK版本、状态、能力 ------------------------

//SDK状态信息(9000新增)
type NET_DVR_SDKSTATE struct {
	DwTotalLoginNum         uint32 //当前login用户数
	DwTotalRealPlayNum      uint32 //当前realplay路数
	DwTotalPlayBackNum      uint32 //当前回放或下载路数
	DwTotalAlarmChanNum     uint32 //当前建立报警通道路数
	DwTotalFormatNum        uint32 //当前硬盘格式化路数
	DwTotalFileSearchNum    uint32 //当前日志或文件搜索路数
	DwTotalLogSearchNum     uint32 //当前日志或文件搜索路数
	DwTotalSerialNum        uint32 //当前透明通道路数
	DwTotalUpgradeNum       uint32 //当前升级路数
	DwTotalVoiceComNum      uint32 //当前语音转发路数
	DwTotalBroadCastNum     uint32 //当前语音广播路数
	DwTotalListenNum        uint32 //当前网络监听路数
	DwEmailTestNum          uint32 //当前邮件计数路数
	DwBackupNum             uint32 // 当前文件备份路数
	DwTotalInquestUploadNum uint32 //当前审讯上传路数
	DwRes                   [6]uint32
}
type LPNET_DVR_SDKSTATE *NET_DVR_SDKSTATE

//SDK功能支持信息(9000新增)
type NET_DVR_SDKABL struct {
	DwMaxLoginNum      uint32 //最大login用户数 MAX_LOGIN_USERS
	DwMaxRealPlayNum   uint32 //最大realplay路数 WATCH_NUM
	DwMaxPlayBackNum   uint32 //最大回放或下载路数 WATCH_NUM
	DwMaxAlarmChanNum  uint32 //最大建立报警通道路数 ALARM_NUM
	DwMaxFormatNum     uint32 //最大硬盘格式化路数 SERVER_NUM
	DwMaxFileSearchNum uint32 //最大文件搜索路数 SERVER_NUM
	DwMaxLogSearchNum  uint32 //最大日志搜索路数 SERVER_NUM
	DwMaxSerialNum     uint32 //最大透明通道路数 SERVER_NUM
	DwMaxUpgradeNum    uint32 //最大升级路数 SERVER_NUM
	DwMaxVoiceComNum   uint32 //最大语音转发路数 SERVER_NUM
	DwMaxBroadCastNum  uint32 //最大语音广播路数 MAX_CASTNUM
	DwRes              [10]uint32
}
type LPNET_DVR_SDKABL *NET_DVR_SDKABL

/************************* 用户注册 *************************/

// 设备激活参数结构体
type NET_DVR_ACTIVATECFG struct {
	DwSize    uint32           //结构体大小
	SPassword [PASSWD_LEN]byte //初始密码
	ByRes     [108]byte
}
type LPNET_DVR_ACTIVATECFG *NET_DVR_ACTIVATECFG

// 用户登录参数结构体
type NET_DVR_USER_LOGIN_INFO struct {
	SDeviceAddress [NET_DVR_DEV_ADDRESS_MAX_LEN]byte
	ByUseTransport byte //是否启用能力集透传，0--不启用透传，默认，1--启用透传
	WPort          uint16
	SUserName      [NET_DVR_LOGIN_USERNAME_MAX_LEN]byte
	SPassword      [NET_DVR_LOGIN_PASSWD_MAX_LEN]byte
	CbLoginResult  FLoginResultCallBack
	PUser          unsafe.Pointer
	BUseAsynLogin  int
	ByProxyType    byte //0:不使用代理，1：使用标准代理，2：使用EHome代理
	ByUseUTCTime   byte //0-不进行转换，默认,1-接口上输入输出全部使用UTC时间,SDK完成UTC时间与设备时区的转换,2-接口上输入输出全部使用平台本地时间，SDK完成平台本地时间与设备时区的转换
	ByLoginMode    byte //0-Private 1-ISAPI 2-自适应
	ByHttps        byte //0-不适用tls，1-使用tls 2-自适应
	IProxyID       int  //代理服务器序号，添加代理服务器信息时，相对应的服务器数组下表值
	ByVerifyMode   byte //认证方式，0-不认证，1-双向认证，2-单向认证；认证仅在使用TLS的时候生效;
	ByRes3         [119]byte
}
type LPNET_DVR_USER_LOGIN_INFO *NET_DVR_USER_LOGIN_INFO

// 登录状态回调函数
type FLoginResultCallBack func(lUserID int, dwResult uint32, lpDeviceInfo LPNET_DVR_DEVICEINFO_V30, pUser unsafe.Pointer)

// 设备参数结构体
type NET_DVR_DEVICEINFO_V30 struct {
	SSerialNumber     [SERIALNO_LEN]byte //序列号
	ByAlarmInPortNum  byte               //报警输入个数
	ByAlarmOutPortNum byte               //报警输出个数
	ByDiskNum         byte               //硬盘个数
	ByDVRType         byte               //设备类型, 1:DVR 2:ATM DVR 3:DVS ......
	ByChanNum         byte               //模拟通道个数
	ByStartChan       byte               //起始通道号,例如DVS-1,DVR - 1
	ByAudioChanNum    byte               //语音通道数
	ByIPChanNum       byte               //最大数字通道个数，低位
	ByZeroChanNum     byte               //零通道编码个数 //2010-01-16
	ByMainProto       byte               //主码流传输协议类型 0-private, 1-rtsp,2-同时支持private和rtsp
	BySubProto        byte               //子码流传输协议类型0-private, 1-rtsp,2-同时支持private和rtsp
	BySupport         byte               //能力，位与结果为0表示不支持，1表示支持，
	//bySupport & 0x1, 表示是否支持智能搜索
	//bySupport & 0x2, 表示是否支持备份
	//bySupport & 0x4, 表示是否支持压缩参数能力获取
	//bySupport & 0x8, 表示是否支持多网卡
	//bySupport & 0x10, 表示支持远程SADP
	//bySupport & 0x20, 表示支持Raid卡功能
	//bySupport & 0x40, 表示支持IPSAN 目录查找
	//bySupport & 0x80, 表示支持rtp over rtsp
	BySupport1 byte // 能力集扩充，位与结果为0表示不支持，1表示支持
	//bySupport1 & 0x1, 表示是否支持snmp v30
	//bySupport1 & 0x2, 支持区分回放和下载
	//bySupport1 & 0x4, 是否支持布防优先级
	//bySupport1 & 0x8, 智能设备是否支持布防时间段扩展
	//bySupport1 & 0x10, 表示是否支持多磁盘数（超过33个）
	//bySupport1 & 0x20, 表示是否支持rtsp over http
	//bySupport1 & 0x80, 表示是否支持车牌新报警信息2012-9-28, 且还表示是否支持NET_DVR_IPPARACFG_V40结构体
	BySupport2 byte /*能力，位与结果为0表示不支持，非0表示支持
	                 bySupport2 & 0x1, 表示解码器是否支持通过URL取流解码
	                 bySupport2 & 0x2,  表示支持FTPV40
	                 bySupport2 & 0x4,  表示支持ANR
	                 bySupport2 & 0x8,  表示支持CCD的通道参数配置
	                 bySupport2 & 0x10,  表示支持布防报警回传信息（仅支持抓拍机报警 新老报警结构）
	                 bySupport2 & 0x20,  表示是否支持单独获取设备状态子项
	bySupport2 & 0x40,  表示是否是码流加密设备*/
	WDevType   uint16 //设备型号
	BySupport3 byte   //能力集扩展，位与结果为0表示不支持，1表示支持
	//bySupport3 & 0x1, 表示是否支持批量配置多码流参数
	// bySupport3 & 0x4 表示支持按组配置， 具体包含 通道图像参数、报警输入参数、IP报警输入、输出接入参数、
	// 用户参数、设备工作状态、JPEG抓图、定时和时间抓图、硬盘盘组管理
	//bySupport3 & 0x8为1 表示支持使用TCP预览、UDP预览、多播预览中的"延时预览"字段来请求延时预览（后续都将使用这种方式请求延时预览）。而当bySupport3 & 0x8为0时，将使用 "私有延时预览"协议。
	//bySupport3 & 0x10 表示支持"获取报警主机主要状态（V40）"。
	//bySupport3 & 0x20 表示是否支持通过DDNS域名解析取流

	ByMultiStreamProto byte //是否支持多码流,按位表示,0-不支持,1-支持,bit1-码流3,bit2-码流4,bit7-主码流，bit-8子码流
	ByStartDChan       byte //起始数字通道号,0表示无效
	ByStartDTalkChan   byte //起始数字对讲通道号，区别于模拟对讲通道号，0表示无效
	ByHighDChanNum     byte //数字通道个数，高位
	BySupport4         byte //能力集扩展，位与结果为0表示不支持，1表示支持
	//bySupport4 & 0x4表示是否支持拼控统一接口
	// bySupport4 & 0x80 支持设备上传中心报警使能。表示判断调用接口是 NET_DVR_PDC_RULE_CFG_V42还是 NET_DVR_PDC_RULE_CFG_V41
	ByLanguageType byte // 支持语种能力,按位表示,每一位0-不支持,1-支持
	//  byLanguageType 等于0 表示 老设备
	//  byLanguageType & 0x1表示支持中文
	//  byLanguageType & 0x2表示支持英文
	ByVoiceInChanNum     byte //音频输入通道数
	ByStartVoiceInChanNo byte //音频输入起始通道号 0表示无效
	BySupport5           byte //按位表示,0-不支持,1-支持,bit0-支持多码流
	//bySupport5 &0x01表示支持wEventTypeEx ,兼容dwEventType 的事件类型（支持行为事件扩展）--先占住，防止冲突
	//bySupport5 &0x04表示是否支持使用扩展的场景模式接口
	/*
	   bySupport5 &0x08 设备返回该值表示是否支持计划录像类型V40接口协议(DVR_SET_RECORDCFG_V40/ DVR_GET_RECORDCFG_V40)(在该协议中设备支持类型类型13的配置)
	   之前的部分发布的设备，支持录像类型13，则配置录像类型13。如果不支持，统一转换成录像类型3兼容处理。SDK通过命令探测处理)
	   bySupport5 &0x10 设备返回改值表示支持超过255个预置点
	*/
	BySupport6 byte //能力，按位表示，0-不支持,1-支持
	//bySupport6 0x1  表示设备是否支持压缩
	//bySupport6 0x2 表示是否支持流ID方式配置流来源扩展命令，DVR_SET_STREAM_SRC_INFO_V40
	//bySupport6 0x4 表示是否支持事件搜索V40接口
	//bySupport6 0x8 表示是否支持扩展智能侦测配置命令
	//bySupport6 0x40表示图片查询结果V40扩展
	ByMirrorChanNum    byte   //镜像通道个数，<录播主机中用于表示导播通道>
	WStartMirrorChanNo uint16 //起始镜像通道号
	BySupport7         byte   //能力,按位表示,0-不支持,1-支持
	// bySupport7 & 0x1  表示设备是否支持 INTER_VCA_RULECFG_V42 扩展
	// bySupport7 & 0x2  表示设备是否支持 IPC HVT 模式扩展
	// bySupport7 & 0x04  表示设备是否支持 返回锁定时间
	// bySupport7 & 0x08  表示设置云台PTZ位置时，是否支持带通道号
	// bySupport7 & 0x10  表示设备是否支持双系统升级备份
	// bySupport7 & 0x20  表示设备是否支持 OSD字符叠加 V50
	// bySupport7 & 0x40  表示设备是否支持 主从跟踪（从摄像机）
	// bySupport7 & 0x80  表示设备是否支持 报文加密
	ByRes2 byte //保留
}
type LPNET_DVR_DEVICEINFO_V30 *NET_DVR_DEVICEINFO_V30

// 设备参数结构体
type NET_DVR_DEVICEINFO_V40 struct {
	StruDeviceV30     NET_DVR_DEVICEINFO_V30
	BySupportLock     byte   //设备支持锁定功能，该字段由SDK根据设备返回值来赋值的。bySupportLock为1时，dwSurplusLockTime和byRetryLoginTime有效
	ByRetryLoginTime  byte   //剩余可尝试登陆的次数，用户名，密码错误时，此参数有效
	ByPasswordLevel   byte   //admin密码安全等级0-无效，1-默认密码，2-有效密码,3-风险较高的密码。当用户的密码为出厂默认密码（12345）或者风险较高的密码时，上层客户端需要提示用户更改密码。4-管理员创建一个普通用户为其设置密码，该普通用户正确登录设备后要提示“请修改初始登录密码”，未修改的情况下，用户每次登入都会进行提醒；5-当普通用户的密码被管理员修改，该普通用户再次正确登录设备后，需要提示“请重新设置登录密码”，未修改的情况下，用户每次登入都会进行提醒;
	ByProxyType       byte   //代理类型，0-不使用代理, 1-使用socks5代理, 2-使用EHome代理
	BwSurplusLockTime uint32 //剩余时间，单位秒，用户锁定时，此参数有效
	ByCharEncodeType  byte   //字符编码类型
	BySupportDev5     byte   //支持v50版本的设备参数获取，设备名称和设备类型名称长度扩展为64字节
	BySupport         byte   //能力集扩展，位与结果：0- 不支持，1- 支持
	// bySupport & 0x1:  保留
	// bySupport & 0x2:  0-不支持变化上报 1-支持变化上报
	ByLoginMode            byte //登录模式 0-Private登录 1-ISAPI登录
	DwOEMCode              uint32
	IResidualValidity      int  //该用户密码剩余有效天数，单位：天，返回负值，表示密码已经超期使用，例如“-3表示密码已经超期使用3天”
	ByResidualValidity     byte // iResidualValidity字段是否有效，0-无效，1-有效
	BySingleStartDTalkChan byte //独立音轨接入的设备，起始接入通道号，0-为保留字节，无实际含义，音轨通道号不能从0开始
	BySingleDTalkChanNums  byte //独立音轨接入的设备的通道总数，0-表示不支持
	ByPassWordResetLevel   byte //0-无效，1-管理员创建一个非管理员用户为其设置密码，该非管理员用户正确登录设备后要提示“请修改初始登录密码”，未修改的情况下，用户每次登入都会进行提醒；2-当非管理员用户的密码被管理员修改，该非管理员用户再次正确登录设备后，需要提示“请重新设置登录密码”，未修改的情况下，用户每次登入都会进行提醒。
	BySupportStreamEncrypt byte //能力集扩展，位与结果：0- 不支持，1- 支持 bySupportStreamEncrypt & 0x1:表示是否支持RTP/TLS取流 bySupportStreamEncrypt & 0x2:  表示是否支持SRTP/UDP取流 bySupportStreamEncrypt & 0x4:  表示是否支持SRTP/MULTICAST取流
	ByMarketType           byte //0-无效（未知类型）,1-经销型，2-行业型
	ByRes2                 [238]byte
}
type LPNET_DVR_DEVICEINFO_V40 *NET_DVR_DEVICEINFO_V40
