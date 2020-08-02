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

/************************* 网络参数配置 *************************/

// 网络配置回调函数
type FRemoteConfigCallback func(dwType uint32, lpBuffer unsafe.Pointer, dwBufLen uint32, pUserData unsafe.Pointer)

type NET_DVR_CARD_CFG_COND struct {
	DwSize             uint32
	DwCardNum          uint32 //设置或获取卡数量，获取时置为0xffffffff表示获取所有卡信息
	ByCheckCardNo      byte   //设备是否进行卡号校验，0-不校验，1-校验
	ByRes1             [3]byte
	WLocalControllerID uint16 //就地控制器序号，表示往就地控制器下发离线卡参数，0代表是门禁主机
	ByRes2             [2]byte
	DwLockID           uint32 //锁ID
	ByRes3             [20]byte
}

type LPNET_DVR_CARD_CFG_COND *NET_DVR_CARD_CFG_COND

// enum NET_SDK_CALLBACK_TYPE
type NET_SDK_CALLBACK_TYPE int32

const (
	NET_SDK_CALLBACK_TYPE_STATUS   NET_SDK_CALLBACK_TYPE = 0 // 回调状态值
	NET_SDK_CALLBACK_TYPE_PROGRESS NET_SDK_CALLBACK_TYPE = 1 // 回调进度值
	NET_SDK_CALLBACK_TYPE_DATA     NET_SDK_CALLBACK_TYPE = 2 // 回调数据内容
)

// enum NET_SDK_CALLBACK_STATUS_NORMAL
type NET_SDK_CALLBACK_STATUS_NORMAL int32

const (
	NET_SDK_CALLBACK_STATUS_SUCCESS           NET_SDK_CALLBACK_STATUS_NORMAL = 1000 // 成功
	NET_SDK_CALLBACK_STATUS_PROCESSING        NET_SDK_CALLBACK_STATUS_NORMAL = 1001 // 处理中
	NET_SDK_CALLBACK_STATUS_FAILED            NET_SDK_CALLBACK_STATUS_NORMAL = 1002 // 失败
	NET_SDK_CALLBACK_STATUS_EXCEPTION         NET_SDK_CALLBACK_STATUS_NORMAL = 1003 // 异常
	NET_SDK_CALLBACK_STATUS_LANGUAGE_MISMATCH NET_SDK_CALLBACK_STATUS_NORMAL = 1004 //（IPC配置文件导入）语言不匹配
	NET_SDK_CALLBACK_STATUS_DEV_TYPE_MISMATCH NET_SDK_CALLBACK_STATUS_NORMAL = 1005 //（IPC配置文件导入）设备类型不匹配
	NET_DVR_CALLBACK_STATUS_SEND_WAIT         NET_SDK_CALLBACK_STATUS_NORMAL = 1006 // 发送等待
)

type NET_DVR_CARD_CFG struct {
	DwSize            uint32
	DwModifyParamType uint32
	// 需要修改的卡参数，设置卡参数时有效，按位表示，每位代表一种参数，1为需要修改，0为不修改
	ByCardNo        [ACS_CARD_NO_LEN]byte //卡号
	ByCardValid     byte                  //卡是否有效，0-无效，1-有效（用于删除卡，设置时置为0进行删除，获取时此字段始终为1）
	ByCardType      byte                  //卡类型，1-普通卡，2-残疾人卡，3-黑名单卡，4-巡更卡，5-胁迫卡，6-超级卡，7-来宾卡，8-解除卡，默认普通卡
	ByLeaderCard    byte                  //是否为首卡，1-是，0-否
	ByRes1          byte
	DwDoorRight     uint32                                      //门权限，按位表示，1为有权限，0为无权限，从低位到高位表示对门1-N是否有权限
	StruValid       NET_DVR_VALID_PERIOD_CFG                    //有效期参数
	DwBelongGroup   uint32                                      //所属群组，按位表示，1-属于，0-不属于，从低位到高位表示是否从属群组1-N
	ByCardPassword  [CARD_PASSWORD_LEN]byte                     //卡密码
	ByCardRightPlan [MAX_DOOR_NUM][MAX_CARD_RIGHT_PLAN_NUM]byte //卡权限计划，取值为计划模板编号，同个门不同计划模板采用权限或的方式处理
	DwMaxSwipeTime  uint32                                      //最大刷卡次数，0为无次数限制
	DwSwipeTime     uint32                                      //已刷卡次数
	WRoomNumber     uint16                                      //房间号
	WFloorNumber    int16                                       //层号
	ByRes2          [20]byte
}

type LPNET_DVR_CARD_CFG *NET_DVR_CARD_CFG

type NET_DVR_VALID_PERIOD_CFG struct {
	ByEnable         byte            //使能有效期，0-不使能，1使能
	ByBeginTimeFlag  byte            //是否限制起始时间的标志，0-不限制，1-限制
	ByEnableTimeFlag byte            //是否限制终止时间的标志，0-不限制，1-限制
	ByTimeDurationNo byte            //有效期索引,从0开始（时间段通过SDK设置给锁，后续在制卡时，只需要传递有效期索引即可，以减少数据量）
	StruBeginTime    NET_DVR_TIME_EX //有效期起始时间
	StruEndTime      NET_DVR_TIME_EX //有效期结束时间
	ByTimeType       byte            //时间类型：0-设备本地时间（默认），1-UTC时间（对于struBeginTime，struEndTime字段有效）
	ByRes2           [31]byte
}

type LPNET_DVR_VALID_PERIOD_CFG *NET_DVR_VALID_PERIOD_CFG

type NET_DVR_TIME_EX struct {
	WYear    uint16
	ByMonth  byte
	ByDay    byte
	ByHour   byte
	ByMinute byte
	BySecond byte
	ByRes    byte
}

type LPNET_DVR_TIME_EX *NET_DVR_TIME_EX

type NET_DVR_CARD_CFG_V50 struct {
	DwSize            uint32
	DwModifyParamType uint32
	// 需要修改的卡参数，设置卡参数时有效，按位表示，每位代表一种参数，1为需要修改，0为不修改
	ByCardNo           [ACS_CARD_NO_LEN]byte                             //卡号
	ByCardValid        byte                                              //卡是否有效，0-无效，1-有效（用于删除卡，设置时置为0进行删除，获取时此字段始终为1）
	ByCardType         byte                                              //卡类型，1-普通卡，2-残疾人卡，3-黑名单卡，4-巡更卡，5-胁迫卡，6-超级卡，7-来宾卡，8-解除卡，9-员工卡，10-应急卡，11-应急管理卡（用于授权临时卡权限，本身不能开门），默认普通卡
	ByLeaderCard       byte                                              //是否为首卡，1-是，0-否
	ByUserType         byte                                              // 0 – 普通用户1 - 管理员用户;
	ByDoorRight        [MAX_DOOR_NUM_256]byte                            //门权限(楼层权限、锁权限)，按位表示，1为有权限，0为无权限，从低位到高位表示对门（锁）1-N是否有权限
	StruValid          NET_DVR_VALID_PERIOD_CFG                          //有效期参数
	ByBelongGroup      [MAX_GROUP_NUM_128]byte                           //所属群组，按字节表示，1-属于，0-不属于
	ByCardPassword     [CARD_PASSWORD_LEN]byte                           //卡密码
	WCardRightPlan     [MAX_DOOR_NUM_256][MAX_CARD_RIGHT_PLAN_NUM]uint16 //卡权限计划，取值为计划模板编号，同个门（锁）不同计划模板采用权限或的方式处理
	DwMaxSwipeTime     uint32                                            //最大刷卡次数，0为无次数限制（开锁次数）
	DwSwipeTime        uint32                                            //已刷卡次数
	WRoomNumber        uint16                                            //房间号
	WFloorNumber       int16                                             //层号
	DwEmployeeNo       uint32                                            //工号（用户ID）
	ByName             [NAME_LEN]byte                                    //姓名
	WDepartmentNo      uint16                                            //部门编号
	WSchedulePlanNo    uint16                                            //排班计划编号
	BySchedulePlanType byte                                              //排班计划类型：0-无意义、1-个人、2-部门
	ByRightType        byte                                              //下发权限类型：0-普通发卡权限、1-二维码权限、2-蓝牙权限（可视对讲设备二维码权限配置项：房间号、卡号（虚拟卡号）、最大刷卡次数（开锁次数）、有效期参数；蓝牙权限：卡号（萤石APP账号）、其他参数配置与普通发卡权限一致）
	ByRes2             [2]byte
	DwLockID           uint32                  //锁ID
	ByLockCode         [MAX_LOCK_CODE_LEN]byte //锁代码
	ByRoomCode         [MAX_DOOR_CODE_LEN]byte //房间代码
	//按位表示，0-无权限，1-有权限
	//第0位表示：弱电报警
	//第1位表示：开门提示音
	//第2位表示：限制客卡
	//第3位表示：通道
	//第4位表示：反锁开门
	//第5位表示：巡更功能
	DwCardRight     uint32 //卡权限
	DwPlanTemplate  uint32 //计划模板(每天)各时间段是否启用，按位表示，0--不启用，1-启用
	DwCardUserId    uint32 //持卡人ID
	ByCardModelType byte   //0-空，1- MIFARE S50，2- MIFARE S70，3- FM1208 CPU卡，4- FM1216 CPU卡，5-国密CPU卡，6-身份证，7- NFC
	ByRes3          [51]byte
	BySIMNum        [NAME_LEN] /*32*/ byte //SIM卡号（手机号）
}
type LPNET_DVR_CARD_CFG_V50 *NET_DVR_CARD_CFG_V50

// enum CFG_SEND_DATA_TYPE
type CFG_SEND_DATA_TYPE int32

// 数据类型，跟长连接接口NET_DVR_StartRemoteConfig的命令参数（dwCommand）有关
const (
	ENUM_DVR_VEHICLE_CHECK             CFG_SEND_DATA_TYPE = 1  //黑名单车辆数据稽查类型
	ENUM_MSC_SEND_DATA                 CFG_SEND_DATA_TYPE = 2  //屏幕控制器数据类型
	ENUM_ACS_SEND_DATA                 CFG_SEND_DATA_TYPE = 3  //门禁主机数据类型
	ENUM_TME_CARD_SEND_DATA            CFG_SEND_DATA_TYPE = 4  //停车场(出入口控制机)卡片数据类型
	ENUM_TME_VEHICLE_SEND_DATA         CFG_SEND_DATA_TYPE = 5  //停车场(出入口控制机)车辆数据类型
	ENUM_DVR_DEBUG_CMD                 CFG_SEND_DATA_TYPE = 6  //调试命令信息
	ENUM_DVR_SCREEN_CTRL_CMD           CFG_SEND_DATA_TYPE = 7  //屏幕互动命令类型
	ENUM_CVR_PASSBACK_SEND_DATA        CFG_SEND_DATA_TYPE = 8  //CVR获取监控点回传任务可执行性
	ENUM_ACS_INTELLIGENT_IDENTITY_DATA CFG_SEND_DATA_TYPE = 9  //智能身份识别终端数据类型
	ENUM_VIDEO_INTERCOM_SEND_DATA      CFG_SEND_DATA_TYPE = 10 //可视对讲数据类型
	ENUM_SEND_JSON_DATA                CFG_SEND_DATA_TYPE = 11 //透传JSON数据
)

// 获取卡参数的发送数据
type NET_DVR_CARD_CFG_SEND_DATA struct {
	DwSize       uint32
	ByCardNo     [ACS_CARD_NO_LEN]byte //卡号
	DwCardUserId uint32                //持卡人ID
	ByRes        [12]byte
}
type LPNET_DVR_CARD_CFG_SEND_DATA *NET_DVR_CARD_CFG_SEND_DATA

/************************* 人脸参数配置 *************************/

// 人脸参数
// See https://open.hikvision.com/hardware/structures/NET_DVR_FACE_PARAM_COND.html
type NET_DVR_FACE_PARAM_COND struct {
	DwSize             uint32
	ByCardNo           [ACS_CARD_NO_LEN]byte         //人脸关联的卡号（设置时该参数可不设置）
	ByEnableCardReader [MAX_CARD_READER_NUM_512]byte //人脸的读卡器是否有效，0-无效，1-有效（设置时该参数可不设置）
	DwFaceNum          uint32                        //设置或获取人脸数量，获取时置为0xffffffff表示获取所有人脸信息
	ByFaceID           byte                          //人脸编号，有效值范围为1-2  0xff表示该卡所有人脸（设置时该参数可不设置）
	ByFaceDataType     byte                          //人脸数据类型：0-模板（默认），1-图片
	ByRes              [126]byte                     //保留
}
type LPNET_DVR_FACE_PARAM_COND *NET_DVR_FACE_PARAM_COND

// 人脸参数配置结构体
// See https://open.hikvision.com/hardware/structures/NET_DVR_FACE_PARAM_CFG.html
type NET_DVR_FACE_PARAM_CFG struct {
	DwSize             uint32
	ByCardNo           [ACS_CARD_NO_LEN]byte         //人脸关联的卡号
	DwFaceLen          uint32                        //人脸数据长度，最大支持 200k 图片，可选：<DES加密处理>，设备端返回的即加密后的数据
	PFaceBuffer        unsafe.Pointer                //人脸数据指针
	ByEnableCardReader [MAX_CARD_READER_NUM_512]byte //需要下发人脸的读卡器，按数组表示，从低位到高位表示，0-不下发该读卡器，1-下发到该读卡器
	ByFaceID           byte                          //人脸编号，有效值范围为1-2
	ByFaceDataType     byte                          //人脸数据类型：0-模板（默认），1-图片
	ByRes              [126]byte
}
type LPNET_DVR_FACE_PARAM_CFG *NET_DVR_FACE_PARAM_CFG

// 人脸参数下发状态信息结构体
// See https://open.hikvision.com/hardware/structures/NET_DVR_FACE_PARAM_STATUS.html
type NET_DVR_FACE_PARAM_STATUS struct {
	DwSize                 uint32
	ByCardNo               [ACS_CARD_NO_LEN]byte         //人脸关联的卡号
	ByCardReaderRecvStatus [MAX_CARD_READER_NUM_512]byte //人脸读卡器状态，按字节表示，0-失败，1-成功，2-重试或人脸质量差，3-内存已满(人脸数据满)，4-已存在该人脸，5-非法人脸ID
	//,6-算法建模失败，7-未下发卡权限，8-未定义（保留），9-人眼间距小距小，10-图片数据长度小于1KB，11-图片格式不符（png/jpg/bmp）,12-图片像素数量超过上限，13-图片像素数量低于下限，14-图片信息校验失败，15-图片解码失败，16-人脸检测失败，17-人脸评分失败
	ByErrorMsg     [ERROR_MSG_LEN]byte //下发错误信息，当byCardReaderRecvStatus为4时，表示已存在人脸对应的卡号
	DwCardReaderNo uint32              //纹读卡器编号，可用于下发错误返回
	ByTotalStatus  byte                //下发总的状态，0-当前人脸未下完所有读卡器，1-已下完所有读卡器(这里的所有指的是门禁主机往所有的读卡器下发了，不管成功与否)
	ByFaceID       byte                //人脸编号，有效值范围为1-2
	ByRes          [130]byte
}
type LPNET_DVR_FACE_PARAM_STATUS *NET_DVR_FACE_PARAM_STATUS

// 采集人脸信息条件参数结构体
type NET_DVR_CAPTURE_FACE_COND struct {
	DwSize uint32
	ByRes  [128]byte
}
type LPNET_DVR_CAPTURE_FACE_COND *NET_DVR_CAPTURE_FACE_COND

// 人脸信息采集结果结构体
type NET_DVR_CAPTURE_FACE_CFG struct {
	DwSize                   uint32
	DwFaceTemplate1Size      uint32 //人脸模板1数据大小，等于0时，代表无人脸模板1数据
	PFaceTemplate1Buffer     *byte  //人脸模板1数据缓存（不大于2.5k）
	DwFaceTemplate2Size      uint32 //人脸模板2数据大小，等于0时，代表无人脸模板2数据
	PFaceTemplate2Buffer     *byte  //人脸模板2数据缓存（不大于2.5K）
	DwFacePicSize            uint32 //人脸图片数据大小，等于0时，代表无人脸图片数据
	PFacePicBuffer           *byte  //人脸图片数据缓存
	ByFaceQuality1           byte   //人脸质量，范围1-100
	ByFaceQuality2           byte   //人脸质量，范围1-100
	ByCaptureProgress        byte   //采集进度，目前只有两种进度值：0-未采集到人脸，100-采集到人脸（只有在进度为100时，才解析人脸信息）
	ByFacePicQuality         byte   //人脸图片中人脸质量
	DwInfraredFacePicSize    uint32 //红外人脸图片数据大小，等于0时，代表无人脸图片数据
	PInfraredFacePicBuffer   *byte  //红外人脸图片数据缓存
	ByInfraredFacePicQuality byte   //红外人脸图片中人脸质量
	ByRes1                   [3]byte
	StruFeature              NET_DVR_FACE_FEATURE //人脸抠图特征信息
	ByRes                    [56]byte
}
type LPNET_DVR_CAPTURE_FACE_CFG *NET_DVR_CAPTURE_FACE_CFG

// 人脸抠图特征信息
type NET_DVR_FACE_FEATURE struct {
	StruFace       NET_VCA_RECT  //人脸子图区域
	StruLeftEye    NET_VCA_POINT // 左眼坐标
	StruRightEye   NET_VCA_POINT // 右眼坐标
	StruLeftMouth  NET_VCA_POINT // 嘴左边坐标
	StruRightMouth NET_VCA_POINT // 嘴右边坐标
	StruNoseTip    NET_VCA_POINT // 鼻子坐标
}
type LPNET_DVR_FACE_FEATURE *NET_DVR_FACE_FEATURE

// 区域框结构
type NET_VCA_RECT struct {
	FX      float32 //边界框左上角点的X轴坐标, 0.000~1
	FY      float32 //边界框左上角点的Y轴坐标, 0.000~1
	FWidth  float32 //边界框的宽度, 0.000~1
	FHeight float32 //边界框的高度, 0.000~1
}
type LPNET_VCA_RECT *NET_VCA_RECT

// 点坐标结构
type NET_VCA_POINT struct {
	FX float32 // X轴坐标, 0.000~1
	FY float32 //Y轴坐标, 0.000~1
}
type LPNET_VCA_POINT *NET_VCA_POINT

/************************* 指纹参数配置 *************************/

// 指纹参数配置条件结构体
// See https://open.hikvision.com/hardware/structures/NET_DVR_FINGER_PRINT_INFO_COND.html
type NET_DVR_FINGER_PRINT_INFO_COND struct {
	DwSize             uint32
	ByCardNo           [ACS_CARD_NO_LEN]byte         //指纹关联的卡号
	ByEnableCardReader [MAX_CARD_READER_NUM_512]byte //指纹的读卡器信息，按数组表示
	DwFingerPrintNum   uint32                        //设置或获取卡数量，获取时置为0xffffffff表示获取所有卡信息
	ByFingerPrintID    byte                          //手指编号，有效值范围为-10   0xff表示该卡所有指纹
	ByCallbackMode     byte                          //设备回调方式，0-设备所有读卡器下完了范围，1-在时间段内下了部分也返回
	ByRes1             [26]byte                      //保留
}
type LPNET_DVR_FINGER_PRINT_INFO_COND *NET_DVR_FINGER_PRINT_INFO_COND

// 指纹参数配置结构体
// See https://open.hikvision.com/hardware/structures/NET_DVR_FINGER_PRINT_CFG.html
type NET_DVR_FINGER_PRINT_CFG struct {
	DwSize             uint32
	ByCardNo           [ACS_CARD_NO_LEN]byte         //指纹关联的卡号
	DwFingerPrintLen   uint32                        //指纹数据长度
	ByEnableCardReader [MAX_CARD_READER_NUM_512]byte //需要下发指纹的读卡器，按数组表示，0-不下发该读卡器，1-下发到该读卡器
	ByFingerPrintID    byte                          //手指编号，有效值范围为1-10
	ByFingerType       byte                          //指纹类型  0-普通指纹，1-胁迫指纹
	ByRes1             [30]byte
	ByFingerData       [MAX_FINGER_PRINT_LEN]byte //指纹数据内容
	ByRes              [64]byte
}
type LPNET_DVR_FINGER_PRINT_CFG *NET_DVR_FINGER_PRINT_CFG

// 指纹状态参数结构体
// See https://open.hikvision.com/hardware/structures/NET_DVR_FINGER_PRINT_STATUS.html
type NET_DVR_FINGER_PRINT_STATUS struct {
	DwSize                 uint32
	ByCardNo               [ACS_CARD_NO_LEN]byte         //指纹关联的卡号
	ByCardReaderRecvStatus [MAX_CARD_READER_NUM_512]byte //指纹读卡器状态，按字节表示，0-失败，1-成功，2-该指纹模组不在线，3-重试或指纹质量差，4-内存已满，5-已存在该指纹，6-已存在该指纹ID，7-非法指纹ID，8-该指纹模组无需配置
	ByFingerPrintID        byte                          //手指编号，有效值范围为1-10
	ByFingerType           byte                          //指纹类型  0-普通指纹，1-胁迫指纹
	ByTotalStatus          byte                          //下发总的状态，0-当前指纹未下完所有读卡器，1-已下完所有读卡器(这里的所有指的是门禁主机往所有的读卡器下发了，不管成功与否)
	ByRes1                 byte
	ByErrorMsg             [ERROR_MSG_LEN]byte //下发错误信息，当byCardReaderRecvStatus为5时，表示已存在指纹对应的卡号
	DwCardReaderNo         uint32              //非0表示错误信息byErrMsg有效，其值代表byErrMsg对应的读卡器编号（具体什么错误查看byCardReaderRecvStatus对应编号的值）。0时表示无错误信息
	ByRes                  [24]byte
}
type LPNET_DVR_FINGER_PRINT_STATUS *NET_DVR_FINGER_PRINT_STATUS

// 采集指纹信息条件参数结构体
type NET_DVR_CAPTURE_FINGERPRINT_COND struct {
	DwSize               uint32
	ByFingerPrintPicType byte //图片类型：0-无意义
	ByFingerNo           byte //手指编号，范围1-10
	ByRes                [126]byte
}
type LPNET_DVR_CAPTURE_FINGERPRINT_COND *NET_DVR_CAPTURE_FINGERPRINT_COND

// 指纹信息采集结果结构体
type NET_DVR_CAPTURE_FINGERPRINT_CFG struct {
	DwSize                uint32
	DwFingerPrintDataSize uint32                     //指纹数据大小
	ByFingerData          [MAX_FINGER_PRINT_LEN]byte //指纹数据内容
	DwFingerPrintPicSize  uint32                     //指纹图片大小，等于0时，代表无指纹图片数据
	PFingerPrintPicBuffer *byte                      //指纹图片缓存
	ByFingerNo            byte                       //手指编号，范围1-10
	ByFingerPrintQuality  byte                       //指纹质量，范围1-100
	ByRes                 [62]byte
}
type LPNET_DVR_CAPTURE_FINGERPRINT_CFG *NET_DVR_CAPTURE_FINGERPRINT_CFG

/************************* 报警布防 *************************/

// 登录状态回调函数
// true - 数据接收成功；false - 接收失败
type MSGCallBack_V31 func(lCommand int, pAlarmer LPNET_DVR_ALARMER, pAlarmInfo *byte, dwBufLen uint32, pUser unsafe.Pointer) bool

//报警设备信息
type NET_DVR_ALARMER struct {
	ByUserIDValid     byte               /* userid是否有效 0-无效，1-有效 */
	BySerialValid     byte               /* 序列号是否有效 0-无效，1-有效 */
	ByVersionValid    byte               /* 版本号是否有效 0-无效，1-有效 */
	ByDeviceNameValid byte               /* 设备名字是否有效 0-无效，1-有效 */
	ByMacAddrValid    byte               /* MAC地址是否有效 0-无效，1-有效 */
	ByLinkPortValid   byte               /* login端口是否有效 0-无效，1-有效 */
	ByDeviceIPValid   byte               /* 设备IP是否有效 0-无效，1-有效 */
	BySocketIPValid   byte               /* socket ip是否有效 0-无效，1-有效 */
	LUserID           int                /* NET_DVR_Login()返回值, 布防时有效 */
	SSerialNumber     [SERIALNO_LEN]byte /* 序列号 */
	DwDeviceVersion   uint32             /* 版本信息 高16位表示主版本，低16位表示次版本*/
	SDeviceName       [NAME_LEN]byte     /* 设备名字 */
	ByMacAddr         [MACADDR_LEN]byte  /* MAC地址 */
	WLinkPort         uint16             /* link port */
	SDeviceIP         [128]byte          /* IP地址 */
	SSocketIP         [128]byte          /* 报警主动上传时的socket IP地址 */
	ByIpProtocol      byte               /* Ip协议 0-IPV4, 1-IPV6 */
	ByRes1            [2]byte
	BJSONBroken       byte //JSON断网续传标志。0：不续传；1：续传
	WSocketPort       uint32
	ByRes2            [6]byte
}
type LPNET_DVR_ALARMER *NET_DVR_ALARMER

// 报警布防参数结构体
// See https://open.hikvision.com/hardware/structures/NET_DVR_SETUPALARM_PARAM.html
type NET_DVR_SETUPALARM_PARAM struct {
	DwSize              uint32
	ByLevel             byte //布防优先级，0-一等级（高），1-二等级（中），2-三等级（低）
	ByAlarmInfoType     byte //上传报警信息类型（抓拍机支持），0-老报警信息（NET_DVR_PLATE_RESULT），1-新报警信息(NET_ITS_PLATE_RESULT)2012-9-28
	ByRetAlarmTypeV40   byte //0--返回NET_DVR_ALARMINFO_V30或NET_DVR_ALARMINFO, 1--设备支持NET_DVR_ALARMINFO_V40则返回NET_DVR_ALARMINFO_V40，不支持则返回NET_DVR_ALARMINFO_V30或NET_DVR_ALARMINFO
	ByRetDevInfoVersion byte //CVR上传报警信息回调结构体版本号 0-COMM_ALARM_DEVICE， 1-COMM_ALARM_DEVICE_V40
	ByRetVQDAlarmType   byte //VQD报警上传类型，0-上传报报警NET_DVR_VQD_DIAGNOSE_INFO，1-上传报警NET_DVR_VQD_ALARM
	//1-表示人脸侦测报警扩展(INTER_FACE_DETECTION),0-表示原先支持结构(INTER_FACESNAP_RESULT)
	ByFaceAlarmDetection byte
	//Bit0- 表示二级布防是否上传图片: 0-上传，1-不上传
	//Bit1- 表示开启数据上传确认机制；0-不开启，1-开启
	//Bit6- 表示雷达检测报警(eventType:radarDetection)是否开启实时上传；0-不开启，1-开启（用于web插件实时显示雷达目标轨迹）
	BySupport byte
	//断网续传类型
	//bit0-车牌检测（IPC） （0-不续传，1-续传）
	//bit1-客流统计（IPC）  （0-不续传，1-续传）
	//bit2-热度图统计（IPC） （0-不续传，1-续传）
	//bit3-人脸抓拍（IPC） （0-不续传，1-续传）
	//bit4-人脸对比（IPC） （0-不续传，1-续传）
	ByBrokenNetHttp byte
	WTaskNo         uint16 //任务处理号 和 (上传数据NET_DVR_VEHICLE_RECOG_RESULT中的字段dwTaskNo对应 同时 下发任务结构 NET_DVR_VEHICLE_RECOG_COND中的字段dwTaskNo对应)
	ByDeployType    byte   //布防类型：0-客户端布防，1-实时布防
	BySubScription  byte   //订阅，按位表示，未开启订阅不上报  //占位
	//Bit7-移动侦测人车分类是否传图；0-不传图(V30上报)，1-传图(V40上报)
	ByRes1         [2]byte
	ByAlarmTypeURL byte //bit0-表示人脸抓拍报警上传（INTER_FACESNAP_RESULT）；0-表示二进制传输，1-表示URL传输（设备支持的情况下，设备支持能力根据具体报警能力集判断,同时设备需要支持URL的相关服务，当前是”云存储“）
	//bit1-表示EVENT_JSON中图片数据长传类型；0-表示二进制传输，1-表示URL传输（设备支持的情况下，设备支持能力根据具体报警能力集判断）
	//bit2 - 人脸比对(报警类型为COMM_SNAP_MATCH_ALARM)中图片数据上传类型：0 - 二进制传输，1 - URL传输
	//bit3 - 行为分析(报警类型为COMM_ALARM_RULE)中图片数据上传类型：0 - 二进制传输，1 - URL传输，本字段设备是否支持，对应软硬件能力集中<isSupportBehaviorUploadByCloudStorageURL>节点是否返回且为true
	ByCustomCtrl byte //Bit0- 表示支持副驾驶人脸子图上传: 0-不上传,1-上传
}
type LPNET_DVR_SETUPALARM_PARAM *NET_DVR_SETUPALARM_PARAM

// 门禁主机报警信息结构体
// See https://open.hikvision.com/hardware/structures/NET_DVR_ACS_ALARM_INFO.html
type NET_DVR_ACS_ALARM_INFO struct {
	DwSize               uint32
	DwMajor              uint32                 //报警主类型，参考宏定义
	DwMinor              uint32                 //报警次类型，参考宏定义
	StruTime             NET_DVR_TIME           //时间
	SNetUser             [MAX_NAMELEN]byte      //网络操作的用户名
	StruRemoteHostAddr   NET_DVR_IPADDR         //远程主机地址
	StruAcsEventInfo     NET_DVR_ACS_EVENT_INFO //详细参数
	DwPicDataLen         uint32                 //图片数据大小，不为0是表示后面带数据
	PPicData             *byte
	WInductiveEventType  uint16 //归纳事件类型，0-无效，客户端判断该值为非0值后，报警类型通过归纳事件类型区分，否则通过原有报警主次类型（dwMajor、dwMinor）区分
	ByPicTransType       byte   //图片数据传输方式: 0-二进制；1-url
	ByRes1               byte   //保留字节
	DwIOTChannelNo       uint32 //IOT通道号
	PAcsEventInfoExtend  *byte  //byAcsEventInfoExtend为1时，表示指向一个NET_DVR_ACS_EVENT_INFO_EXTEND结构体
	ByAcsEventInfoExtend byte   //pAcsEventInfoExtend是否有效：0-无效，1-有效
	ByTimeType           byte   //时间类型：0-设备本地时间，1-UTC时间（struTime的时间）
	ByRes                [10]byte
}
type LPNET_DVR_ACS_ALARM_INFO *NET_DVR_ACS_ALARM_INFO

// 主机地址
type NET_DVR_IPADDR struct {
	SIpV4  [16]byte  /* IPv4地址 */
	ByIPv6 [128]byte /* 保留 */
}
type LPNET_DVR_IPADDR *NET_DVR_IPADDR

// 详细参数
type NET_DVR_ACS_EVENT_INFO struct {
	DwSize                         uint32
	ByCardNo                       [ACS_CARD_NO_LEN]byte //卡号，为0无效
	ByCardType                     byte                  //卡类型，1-普通卡，2-残疾人卡，3-黑名单卡，4-巡更卡，5-胁迫卡，6-超级卡，7-来宾卡，8-解除卡，为0无效
	ByWhiteListNo                  byte                  //白名单单号,1-8，为0无效
	ByReportChannel                byte                  //报告上传通道，1-布防上传，2-中心组1上传，3-中心组2上传，为0无效
	ByCardReaderKind               byte                  //读卡器属于哪一类，0-无效，1-IC读卡器，2-身份证读卡器，3-二维码读卡器,4-指纹头
	DwCardReaderNo                 uint32                //读卡器编号，为0无效
	DwDoorNo                       uint32                //门编号(楼层编号)，为0无效（当接的设备为人员通道设备时，门1为进方向，门2为出方向）
	DwVerifyNo                     uint32                //多重卡认证序号，为0无效
	DwAlarmInNo                    uint32                //报警输入号，为0无效
	DwAlarmOutNo                   uint32                //报警输出号，为0无效
	DwCaseSensorNo                 uint32                //事件触发器编号
	DwRs485No                      uint32                //RS485通道号，为0无效
	DwMultiCardGroupNo             uint32                //群组编号
	WAccessChannel                 uint16                //人员通道号
	ByDeviceNo                     byte                  //设备编号，为0无效
	ByDistractControlNo            byte                  //分控器编号，为0无效
	DwEmployeeNo                   uint32                //工号，为0无效
	WLocalControllerID             uint16                //就地控制器编号，0-门禁主机，1-64代表就地控制器
	ByInternetAccess               byte                  //网口ID：（1-上行网口1,2-上行网口2,3-下行网口1）
	ByType                         byte                  //防区类型，0:即时防区,1-24小时防区,2-延时防区 ,3-内部防区，4-钥匙防区 5-火警防区 6-周界防区 7-24小时无声防区  8-24小时辅助防区，9-24小时震动防区,10-门禁紧急开门防区，11-门禁紧急关门防区 0xff-无
	ByMACAddr                      [MACADDR_LEN]byte     //物理地址，为0无效
	BySwipeCardType                byte                  //刷卡类型，0-无效，1-二维码
	ByRes2                         byte
	DwSerialNo                     uint16 //事件流水号，为0无效
	ByChannelControllerID          byte   //通道控制器ID，为0无效，1-主通道控制器，2-从通道控制器
	ByChannelControllerLampID      byte   //通道控制器灯板ID，为0无效（有效范围1-255）
	ByChannelControllerIRAdaptorID byte   //通道控制器红外转接板ID，为0无效（有效范围1-255）
	ByChannelControllerIREmitterID byte   //通道控制器红外对射ID，为0无效（有效范围1-255）
	ByRes                          [4]byte
}
type LPNET_DVR_ACS_EVENT_INFO *NET_DVR_ACS_EVENT_INFO
