package hikvision

//宏定义
const (
	MAX_NAMELEN           = 16  //DVR本地登陆名
	MAX_RIGHT             = 32  //设备支持的权限（1-12表示本地权限，13-32表示远程权限）
	NAME_LEN              = 32  //用户名长度
	MIN_PASSWD_LEN        = 8   //最小密码长度
	PASSWD_LEN            = 16  //密码长度
	STREAM_PASSWD_LEN     = 12  //码流加密密钥最大长度
	MAX_PASSWD_LEN_EX     = 64  //密码长度64位
	GUID_LEN              = 16  //GUID长度
	DEV_TYPE_NAME_LEN     = 24  //设备类型名称长度
	SERIALNO_LEN          = 48  //序列号长度
	MACADDR_LEN           = 6   //mac地址长度
	MAC_ADDRESS_NUM       = 48  //Mac地址长度
	MAX_SENCE_NUM         = 16  //场景数
	RULE_REGION_MAX       = 128 //最大区域
	MAX_ETHERNET          = 2   //设备可配以太网络
	MAX_NETWORK_CARD      = 4   //设备可配最大网卡数目
	MAX_NETWORK_CARD_EX   = 12  //设备可配最大网卡数目扩展
	PATHNAME_LEN          = 128 //路径长度
	MAX_PRESET_V13        = 16  //预置点
	MAX_TEST_COMMAND_NUM  = 32  //产线测试保留字段长度
	MAX_NUMBER_LEN        = 32  //号码最大长度
	MAX_NAME_LEN          = 128 //设备名称最大长度
	MAX_INDEX_LED         = 8   //LED索引最大值 2013-11-19
	MAX_CUSTOM_DIR        = 64  //自定义目录最大长度
	URL_LEN_V40           = 256 //最大URL长度
	CLOUD_NAME_LEN        = 48  //云存储服务器用户名长度
	CLOUD_PASSWD_LEN      = 48  //云存储服务器密码长度
	MAX_SENSORNAME_LEN    = 64  //传感器名称长度
	MAX_SENSORCHAN_LEN    = 32  //传感器通道长度
	MAX_DESCRIPTION_LEN   = 32  //传感器描述长度
	MAX_DEVNAME_LEN_EX    = 64  //设备名称长度扩展
	NET_SDK_MAX_FILE_PATH = 256 //文件路径长度
	MAX_TMEVOICE_LEN      = 64  //TME语音播报内容长度
	ISO_8601_LEN          = 32  //ISO_8601时间长度
	MODULE_INFO_LEN       = 32  //模块信息长度
	VERSION_INFO_LEN      = 32  //版本信息长度

	MAX_NUM_INPUT_BOARD      = 512 //输入板最大个数
	MAX_SHIPSDETE_REGION_NUM = 8   // 船只检测区域列表最大数目

	MAX_RES_NUM_ONE_VS_INPUT_CHAN = 8  //一个虚拟屏输入通道支持的分辨率的最大数量
	MAX_VS_INPUT_CHAN_NUM         = 16 //虚拟屏输入通道最大数量

	NET_SDK_MAX_FDID_LEN          = 256 //人脸库ID最大长度
	NET_SDK_MAX_PICID_LEN         = 256 //人脸ID最大长度
	NET_SDK_FDPIC_CUSTOM_INFO_LEN = 96  //人脸库图片自定义信息长度
	NET_DVR_MAX_FACE_ANALYSIS_NUM = 32  //最大支持单张图片识别出的人脸区域个数
	NET_DVR_MAX_FACE_SEARCH_NUM   = 5   //最大支持搜索人脸区域个数
	NET_SDK_SECRETKEY_LEN         = 128 //配置文件密钥长度
	NET_SDK_CUSTOM_LEN            = 512 //自定义信息最大长度
	NET_SDK_CHECK_CODE_LEN        = 128 //校验码长度
	RELATIVE_CHANNEL_LEN          = 2   //报警关联的通道号的数量
	NET_SDK_MAX_CALLEDTARGET_NAME = 32  //呗呼叫目标的用户名
	//小间距LED控制器
	MAX_LEN_TEXT_CONTENT      = 128 //字符内容长度
	MAX_NUM_INPUT_SOURCE_TEXT = 32  //信号源可叠加的文本数量
	MAX_NUM_OUTPUT_CHANNEL    = 512 //LED区域包含的输出口个数

	//子窗口解码OSD
	MAX_LEN_OSD_CONTENT    = 256 //OSD信息最大长度
	MAX_NUM_OSD_ONE_SUBWND = 8   //单个子窗口支持的最大OSD数量
	MAX_NUM_SPLIT_WND      = 64  //单个窗口支持的最大分屏窗口数量（即子窗口数量）
	MAX_NUM_OSD            = 8

	//2013-11-19
	MAX_DEVNAME_LEN      = 32  //设备名称最大长度
	MAX_LED_INFO         = 256 //屏幕字体显示信息最大长度
	MAX_TIME_LEN         = 32  //时间最大长度
	MAX_CARD_LEN         = 24  //卡号最大长度
	MAX_OPERATORNAME_LEN = 32  //操作人员名称最大长度

	THERMOMETRY_ALARMRULE_NUM          = 40  //热成像报警规则数
	MAX_THERMOMETRY_REGION_NUM         = 40  //热度图检测区域最大支持数
	MAX_THERMOMETRY_DIFFCOMPARISON_NUM = 40  //热成像温差报警规则数
	MAX_SHIPS_NUM                      = 20  //船只检测最大船只数
	KEY_WORD_NUM                       = 3   //关键字个数
	KEY_WORD_LEN                       = 128 //关键字长度
	//异步登录回调状态宏定义
	ASYN_LOGIN_SUCC   = 1 //异步登录成功
	ASYN_LOGIN_FAILED = 0 //异步登录失败

	NET_SDK_MAX_VERIFICATION_CODE_LEN = 32  //萤石云验证码长度
	NET_SDK_MAX_OPERATE_CODE_LEN      = 64  //萤石云操作码长度
	MAX_TIMESEGMENT_V30               = 8   //9000设备最大时间段数
	MAX_TIMESEGMENT                   = 4   //8000设备最大时间段数
	MAX_ICR_NUM                       = 8   //抓拍机红外滤光片预置点数2013-07-09
	MAX_VEHICLEFLOW_INFO              = 24  //车流量信息最大个数
	MAX_SHELTERNUM                    = 4   //8000设备最大遮挡区域数
	MAX_DAYS                          = 7   //每周天数
	PHONENUMBER_LEN                   = 32  //pppoe拨号号码最大长度
	MAX_ACCESSORY_CARD                = 256 //配件板信息最大长度
	MAX_DISKNUM_V30                   = 33  //9000设备最大硬盘数/* 最多33个硬盘(包括16个内置SATA硬盘、1个eSATA硬盘和16个NFS盘) */

	NET_SDK_DISK_LOCATION_LEN = 16 //硬盘位置长度
	NET_SDK_SUPPLIER_NAME_LEN = 32 //供应商名称长度
	NET_SDK_DISK_MODEL_LEN    = 64 //硬盘型号长度
	NET_SDK_MAX_DISK_VOLUME   = 33 //最大硬盘卷个数
	NET_SDK_DISK_VOLUME_LEN   = 36 //硬盘卷名称长度

	MAX_DISKNUM             = 16 //8000设备最大硬盘数
	MAX_DISKNUM_V10         = 8  //1.2版本之前版本
	CARD_READER_DESCRIPTION = 32 //读卡器描述
	MAX_FACE_NUM            = 2  //最大人脸数

	MAX_WINDOW_V30 = 32 //9000设备本地显示最大播放窗口数
	MAX_WINDOW_V40 = 64 //Netra 2.3.1扩展
	MAX_WINDOW     = 16 //8000设备最大硬盘数
	MAX_VGA_V30    = 4  //9000设备最大可接VGA数
	MAX_VGA        = 1  //8000设备最大可接VGA数

	MAX_USERNUM_V30        = 32  //9000设备最大用户数
	MAX_USERNUM            = 16  //8000设备最大用户数
	MAX_EXCEPTIONNUM_V30   = 32  //9000设备最大异常处理数
	MAX_EXCEPTIONNUM       = 16  //8000设备最大异常处理数
	MAX_LINK               = 6   //8000设备单通道最大视频流连接数
	MAX_ITC_EXCEPTIONOUT   = 32  //抓拍机最大报警输出
	MAX_SCREEN_DISPLAY_LEN = 512 //屏幕显示字符长度

	MAX_DECPOOLNUM     = 4  //单路解码器每个解码通道最大可循环解码数
	MAX_DECNUM         = 4  //单路解码器的最大解码通道数（实际只有一个，其他三个保留）
	MAX_TRANSPARENTNUM = 2  //单路解码器可配置最大透明通道数
	MAX_CYCLE_CHAN     = 16 //单路解码器最大轮巡通道数
	MAX_CYCLE_CHAN_V30 = 64 //最大轮巡通道数（扩展）
	MAX_DIRNAME_LENGTH = 80 //最大目录长度

	MAX_STRINGNUM_V30        = 8    //9000设备最大OSD字符行数数
	MAX_STRINGNUM            = 4    //8000设备最大OSD字符行数数
	MAX_STRINGNUM_EX         = 8    //8000定制扩展
	MAX_AUXOUT_V30           = 16   //9000设备最大辅助输出数
	MAX_AUXOUT               = 4    //8000设备最大辅助输出数
	MAX_HD_GROUP             = 16   //9000设备最大硬盘组数
	MAX_HD_GROUP_V40         = 32   //设备最大硬盘组数
	MAX_NFS_DISK             = 8    //8000设备最大NFS硬盘数
	NET_SDK_VERSION_LIST_LEN = 64   //算法库版本最大值
	IW_ESSID_MAX_SIZE        = 32   //WIFI的SSID号长度
	IW_ENCODING_TOKEN_MAX    = 32   //WIFI密锁最大字节数
	MAX_SERIAL_NUM           = 64   //最多支持的透明通道路数
	MAX_DDNS_NUMS            = 10   //9000设备最大可配ddns数
	MAX_DOMAIN_NAME          = 64   /* 最大域名长度 */
	MAX_EMAIL_ADDR_LEN       = 48   //最大email地址长度
	MAX_EMAIL_PWD_LEN        = 32   //最大email密码长度
	MAX_SLAVECAMERA_NUM      = 8    //从摄像机个数
	MAX_CALIB_NUM            = 6    //标定点的个数
	MAX_CALIB_NUM_EX         = 20   //扩展标定点的个数
	MAX_LEDDISPLAYINFO_LEN   = 1024 //最大LED屏显示长度
	MAX_PEOPLE_DETECTION_NUM = 8    //最大人员检测区域数
	MAXPROGRESS              = 100  //回放时的最大百分率
	MAX_SERIALNUM            = 2    //8000设备支持的串口数 1-232， 2-485
	CARDNUM_LEN              = 20   //卡号长度
	PATIENTID_LEN            = 64
	CARDNUM_LEN_OUT          = 32 //外部结构体卡号长度
	MAX_VIDEOOUT_V30         = 4  //9000设备的视频输出数
	MAX_VIDEOOUT             = 2  //8000设备的视频输出数

	MAX_PRESET_V30 = 256 /* 9000设备支持的云台预置点数 */
	MAX_TRACK_V30  = 256 /* 9000设备支持的云台轨迹数 */
	MAX_CRUISE_V30 = 256 /* 9000设备支持的云台巡航数 */
	MAX_PRESET     = 128 /* 8000设备支持的云台预置点数 */
	MAX_TRACK      = 128 /* 8000设备支持的云台轨迹数 */
	MAX_CRUISE     = 128 /* 8000设备支持的云台巡航数 */

	MAX_PRESET_V40          = 300 /* 云台支持的最大预置点数 */
	MAX_CRUISE_POINT_NUM    = 128 /* 最大支持的巡航点的个数 */
	MAX_CRUISEPOINT_NUM_V50 = 256 //最大支持的巡航点的个数扩展

	CRUISE_MAX_PRESET_NUMS = 32 /* 一条巡航最多的巡航点 */
	MAX_FACE_PIC_NUM       = 30 /*人脸子图个数*/
	LOCKGATE_TIME_NUM      = 4  //锁闸时间段个数

	MAX_SERIAL_PORT  = 8     //9000设备支持232串口数
	MAX_PREVIEW_MODE = 8     /* 设备支持最大预览模式数目 1画面,4画面,9画面,16画面.... */
	MAX_MATRIXOUT    = 16    /* 最大模拟矩阵输出个数 */
	LOG_INFO_LEN     = 11840 /* 日志附加信息 */
	DESC_LEN         = 16    /* 云台描述字符串长度 */
	PTZ_PROTOCOL_NUM = 200   /* 9000最大支持的云台协议数 */
	IPC_PROTOCOL_NUM = 50    //ipc 协议最大个数

	MAX_AUDIO     = 1  //8000语音对讲通道数
	MAX_AUDIO_V30 = 2  //9000语音对讲通道数
	MAX_CHANNUM   = 16 //8000设备最大通道数
	MAX_ALARMIN   = 16 //8000设备最大报警输入数
	MAX_ALARMOUT  = 4  //8000设备最大报警输出数
	//9000 IPC接入
	MAX_ANALOG_CHANNUM  = 32 //最大32个模拟通道
	MAX_ANALOG_ALARMOUT = 32 //最大32路模拟报警输出
	MAX_ANALOG_ALARMIN  = 32 //最大32路模拟报警输入

	MAX_IP_DEVICE       = 32   //允许接入的最大IP设备数
	MAX_IP_DEVICE_V40   = 64   // 允许接入的最大IP设备数 最多可添加64个 IVMS 2000等新设备
	MAX_IP_CHANNEL      = 32   //允许加入的最多IP通道数
	MAX_IP_ALARMIN      = 128  //允许加入的最多报警输入数
	MAX_IP_ALARMOUT     = 64   //允许加入的最多报警输出数
	MAX_IP_ALARMIN_V40  = 4096 //允许加入的最多报警输入数
	MAX_IP_ALARMOUT_V40 = 4096 //允许加入的最多报警输出数

	MAX_RECORD_FILE_NUM = 20 // 每次删除或者刻录的最大文件数
	//SDK_V31 ATM
	MAX_ACTION_TYPE      = 12   //自定义协议叠加交易行为最大行为个数
	MAX_ATM_PROTOCOL_NUM = 256  //每种输入方式对应的ATM最大协议数
	ATM_CUSTOM_PROTO     = 1025 //自定义协议 值为1025
	ATM_PROTOCOL_SORT    = 4    //ATM协议段数
	ATM_DESC_LEN         = 32   //ATM描述字符串长度
	// SDK_V31 ATM

	MAX_IPV6_LEN    = 64 //IPv6地址最大长度
	MAX_EVENTID_LEN = 64 //事件ID长度

	INVALID_VALUE_UINT32 = 0xffffffff //无效值
	MAX_CHANNUM_V40      = 512
	MAX_MULTI_AREA_NUM   = 24

	//SDK 录播主机
	COURSE_NAME_LEN        = 32  //课程名称
	INSTRUCTOR_NAME_LEN    = 16  //授课教师
	COURSE_DESCRIPTION_LEN = 256 //课程信息

	MAX_TIMESEGMENT_V40 = 16 //每节课信息

	MAX_MIX_CHAN_NUM      = 16 /*目前支持的最大混音通道数，背景通道 + MIC + LINE IN + 最多4个小画面*/
	MAX_LINE_IN_CHAN_NUM  = 16 //最大line in通道数
	MAX_MIC_CHAN_NUM      = 16 //最大MIC通道数
	INQUEST_CASE_NO_LEN   = 64 //审讯案件编号长度
	INQUEST_CASE_NAME_LEN = 64 //审讯案件名称长度
	CUSTOM_INFO_LEN       = 64 //自定义信息长度
	INQUEST_CASE_LEN      = 64 //审讯信息长度

	MAX_FILE_ID_LEN  = 128 //视图库项目中文件ID的最大长度
	MAX_PIC_NAME_LEN = 128 //图片名称长度

	/* 最大支持的通道数 最大模拟加上最大IP支持 */
	MAX_CHANNUM_V30                  = (MAX_ANALOG_CHANNUM + MAX_IP_CHANNEL)       //64
	MAX_ALARMOUT_V40                 = (MAX_IP_ALARMOUT_V40 + MAX_ANALOG_ALARMOUT) //4128
	MAX_ALARMOUT_V30                 = (MAX_ANALOG_ALARMOUT + MAX_IP_ALARMOUT)     //96
	MAX_ALARMIN_V30                  = (MAX_ANALOG_ALARMIN + MAX_IP_ALARMIN)       //160
	MAX_ALARMIN_V40                  = (MAX_IP_ALARMIN_V40 + MAX_ANALOG_ALARMOUT)  //4128
	MAX_ANALOG_ALARM_WITH_VOLT_LIMIT = 16                                          //受电压限定的模拟报警最大输入数

	MAX_ROIDETECT_NUM     = 8  //支持的ROI区域数
	MAX_LANERECT_NUM      = 5  //最大车牌识别区域数
	MAX_FORTIFY_NUM       = 10 //最大布防个数
	MAX_INTERVAL_NUM      = 4  //最大时间间隔个数
	MAX_CHJC_NUM          = 3  //最大车辆省份简称字符个数
	MAX_VL_NUM            = 5  //最大虚拟线圈个数
	MAX_DRIVECHAN_NUM     = 16 //最大车道数
	MAX_COIL_NUM          = 3  //最大线圈个数
	MAX_SIGNALLIGHT_NUM   = 6  //最大信号灯个数
	LEN_16                = 16
	LEN_32                = 32
	LEN_64                = 64
	LEN_31                = 31
	MAX_LINKAGE_CHAN_NUM  = 16 //报警联动的通道的最大数量
	MAX_CABINET_COUNT     = 8  //最大支持机柜数量
	MAX_ID_LEN            = 48
	MAX_PARKNO_LEN        = 16
	MAX_ALARMREASON_LEN   = 32
	MAX_UPGRADE_INFO_LEN  = 48  //获取升级文件匹配信息(模糊升级)
	MAX_CUSTOMDIR_LEN     = 32  //自定义目录长度
	MAX_LED_INFO_LEN      = 512 //LED内容长度
	MAX_VOICE_INFO_LEN    = 128 //语音播报内容长度
	MAX_LITLE_INFO_LEN    = 64  //纸票标题内容长度
	MAX_CUSTOM_INFO_LEN   = 64  //纸票自定义信息内容长度
	MAX_PHONE_NUM_LEN     = 16  //联系电话内容长度
	MAX_APP_SERIALNUM_LEN = 32  //应用序列号长度

	AUDIOTALKTYPE_G722    = 0
	AUDIOTALKTYPE_G711_MU = 1
	AUDIOTALKTYPE_G711_A  = 2
	AUDIOTALKTYPE_MP2L2   = 5
	AUDIOTALKTYPE_G726    = 6
	AUDIOTALKTYPE_AAC     = 7
	AUDIOTALKTYPE_PCM     = 8
	AUDIOTALKTYPE_G722C   = 9
	AUDIOTALKTYPE_MP3     = 15

	//packet type
	FILE_HEAD     = 0  //file head
	VIDEO_I_FRAME = 1  //video I frame
	VIDEO_B_FRAME = 2  //video B frame
	VIDEO_P_FRAME = 3  //video P frame
	AUDIO_PACKET  = 10 //audio packet
	PRIVT_PACKET  = 11 //private packet
	//E frame
	HIK_H264_E_FRAME           = (1 << 6) // 以前E帧不用了,深P帧也没用到
	MAX_TRANSPARENT_CHAN_NUM   = 4        //每个串口允许建立的最大透明通道数
	MAX_TRANSPARENT_ACCESS_NUM = 4        //每个监听端口允许接入的最大主机数

	//ITS
	MAX_PARKING_STATUS = 8 //车位状态 0代表无车，1代表有车，2代表压线(优先级最高), 3特殊车位
	MAX_PARKING_NUM    = 4 //一个通道最大4个车位 (从左到右车位 数组0～3)

	MAX_ITS_SCENE_NUM     = 16  //最大场景数量
	MAX_SCENE_TIMESEG_NUM = 16  //最大场景时间段数量
	MAX_IVMS_IP_CHANNEL   = 128 //最大IP通道数
	DEVICE_ID_LEN         = 48  //设备编号长度
	MONITORSITE_ID_LEN    = 48  //监测点编号长度
	MAX_AUXAREA_NUM       = 16  //辅助区域最大数目
	MAX_SLAVE_CHANNEL_NUM = 16  //最大从通道数量
	MAX_DEVDESC_LEN       = 64  //设备描述信息最大长度
	ILLEGAL_LEN           = 32  //违法代码长度
	MAX_TRUCK_AXLE_NUM    = 10  //货车轴最大数
	MAX_CATEGORY_LEN      = 8   //车牌附加信息最大字符
	SERIAL_NO_LEN         = 16  //泊车位编号

	MAX_SECRETKEY_LEN  = 512 //最大秘钥长度
	MAX_INDEX_CODE_LEN = 64  //最大序号长度
	MAX_ILLEGAL_LEN    = 64  //违法代码最大字符长度
	CODE_LEN           = 64  //授权码
	ALIAS_LEN          = 32  //别名，只读
	MAX_SCH_TASKS_NUM  = 10

	MAX_SERVERID_LEN           = 64  //最大服务器ID的长度
	MAX_SERVERDOMAIN_LEN       = 128 //服务器域名最大长度
	MAX_AUTHENTICATEID_LEN     = 64  //认证ID最大长度
	MAX_AUTHENTICATEPASSWD_LEN = 32  //认证密码最大长度
	MAX_SERVERNAME_LEN         = 64  //最大服务器用户名
	MAX_COMPRESSIONID_LEN      = 64  //编码ID的最大长度
	MAX_SIPSERVER_ADDRESS_LEN  = 128 //SIP服务器地址支持域名和IP地址
	//压线报警
	MAX_PlATE_NO_LEN = 32 //车牌号码最大长度 2013-09-27
	UPNP_PORT_NUM    = 12 //upnp端口映射端口数目

	MAX_NOTICE_NUMBER_LEN = 32   //公告编号最大长度
	MAX_NOTICE_THEME_LEN  = 64   //公告主题最大长度
	MAX_NOTICE_DETAIL_LEN = 1024 //公告详情最大长度
	MAX_NOTICE_PIC_NUM    = 6    //公告信息最大图片数量
	MAX_DEV_NUMBER_LEN    = 32   //设备编号最大长度
	LOCK_NAME_LEN         = 32   //锁名称

	HOLIDAY_GROUP_NAME_LEN          = 32  //假日组名称长度
	MAX_HOLIDAY_PLAN_NUM            = 16  //假日组最大假日计划数
	TEMPLATE_NAME_LEN               = 32  //计划模板名称长度
	MAX_HOLIDAY_GROUP_NUM           = 16  //计划模板最大假日组数
	DOOR_NAME_LEN                   = 32  //门名称
	STRESS_PASSWORD_LEN             = 8   //胁迫密码长度
	SUPER_PASSWORD_LEN              = 8   //胁迫密码长度
	GROUP_NAME_LEN                  = 32  //群组名称长度
	GROUP_COMBINATION_NUM           = 8   //群组组合数
	MULTI_CARD_GROUP_NUM            = 4   //单门最大多重卡组数
	ACS_CARD_NO_LEN                 = 32  //门禁卡号长度
	NET_SDK_EMPLOYEE_NO_LEN         = 32  //工号长度
	NET_SDK_UUID_LEN                = 36  //UUID长度
	NET_SDK_EHOME_KEY_LEN           = 32  //EHome Key长度
	CARD_PASSWORD_LEN               = 8   //卡密码长度
	MAX_DOOR_NUM                    = 32  //最大门数
	MAX_CARD_RIGHT_PLAN_NUM         = 4   //卡权限最大计划个数
	MAX_GROUP_NUM_128               = 128 //最大群组数
	MAX_CARD_READER_NUM             = 64  //最大读卡器数
	MAX_SNEAK_PATH_NODE             = 8   //最大后续读卡器数
	MAX_MULTI_DOOR_INTERLOCK_GROUP  = 8   //最大多门互锁组数
	MAX_INTER_LOCK_DOOR_NUM         = 8   //一个多门互锁组中最大互锁门数
	MAX_CASE_SENSOR_NUM             = 8   //最大case sensor触发器数
	MAX_DOOR_NUM_256                = 256 //最大门数
	MAX_READER_ROUTE_NUM            = 16  //最大刷卡循序路径
	MAX_FINGER_PRINT_NUM            = 10  //最大指纹个数
	MAX_CARD_READER_NUM_512         = 512 //最大读卡器数
	NET_SDK_MULTI_CARD_GROUP_NUM_20 = 20  //单门最大多重卡组数

	ERROR_MSG_LEN       = 32 //下发错误信息
	MAX_DOOR_CODE_LEN   = 8  //房间代码长度
	MAX_LOCK_CODE_LEN   = 8  //锁代码长度
	PER_RING_PORT_NUM   = 2  //每个环的端口数
	SENSORNAME_LEN      = 32 //传感器名称长度
	MAX_SENSORDESCR_LEN = 64 //传感器描述长度
	MAX_DNS_SERVER_NUM  = 2  //最大DNS个数
	SENSORUNIT_LEN      = 32 //最大单位长度

	WEP_KEY_MAX_SIZE = 32 //最大WEP加密密钥长度
	WEP_KEY_MAX_NUM  = 4  //最大WEP加密密钥个数
	WPA_KEY_MAX_SIZE = 64 //最大WPA共享密钥长度

	MAX_SINGLE_FTPPICNAME_LEN = 20 //最大单个FTP通道名称
	MAX_CAMNAME_LEN           = 32 //最大通道名称
	MAX_FTPNAME_NUM           = 12 //TFP名称数

	MAX_IDCODE_LEN  = 128 //  识别码最大长度
	MAX_VERSIIN_LEN = 64  //版本最大长度
	MAX_IDCODE_NUM  = 32  // 识别码个数
	SDK_LEN_2048    = 2048
	SDK_MAX_IP_LEN  = 48

	RECT_POINT_NUM = 4 //矩形角数

	MAX_PUBLIC_KEY_LEN = 512 // 最大公钥长度
	CHIP_SERIALNO_LEN  = 32  //加密芯片序列号长度
	ENCRYPT_DEV_ID_LEN = 20  //设备ID长度

	//MCU相关的
	MAX_SEARCH_ID_LEN = 36  //搜索标识符最大长度
	TERMINAL_NAME_LEN = 64  //终端名称长度
	MAX_URL_LEN       = 512 //URL长度
	REGISTER_NAME_LEN = 64  //终端注册GK名称最大长度

	//光纤
	MAX_PORT_NUM                 = 64  //最大端口数
	MAX_SINGLE_CARD_PORT_NO      = 4   //光纤收发器单卡最大端口数
	MAX_FUNC_CARD_NUM            = 32  //光纤收发器最大功能卡数
	MAX_FC_CARD_NUM              = 33  //光纤收发器最大卡数
	MAX_REMARKS_LEN              = 128 //注释最大长度
	MAX_OUTPUT_PORT_NUM          = 32  //单路输出包含的最大输出端口数
	MAX_SINGLE_PORT_RECVCARD_NUM = 64  //单个端口连接的最大接收卡数
	MAX_GAMMA_X_VALUE            = 256 //GAMMA表X轴取值个数
	NET_DEV_NAME_LEN             = 64  //设备名称长度
	NET_DEV_TYPE_NAME_LEN        = 64  //设备类型名称长度
	ABNORMAL_INFO_NUM            = 4   //异常时间段个数

	PLAYLIST_NAME_LEN = 64 //播放表名称长度
	PLAYLIST_ITEM_NUM = 64 //播放项数目

	//后端相关
	NET_SDK_MAX_LOGIN_PASSWORD_LEN = 128 //用户登录密码最大长度
	NET_SDK_MAX_ANSWER_LEN         = 256 //安全问题答案最大长度
	NET_SDK_MAX_QUESTION_LIST_LEN  = 32  //安全问题列表最大长度

	MAX_SCREEN_AREA_NUM            = 128 //屏幕区域最大数量
	NET_SDK_MAX_THERMOMETRYALGNAME = 128 //测温算法库版本最大长度
	NET_SDK_MAX_SHIPSALGNAME       = 128 //船只算法库版本最大长度
	NET_SDK_MAX_FIRESALGNAME       = 128 //火点算法库版本最大长度

	MAX_PASSPORT_NUM_LEN     = 16   //最大护照证件号长度
	MAX_PASSPORT_INFO_LEN    = 128  //最大护照通用信息长度
	MAX_PASSPORT_NAME_LEN    = 64   //最大护照姓名长度
	MAX_PASSPORT_MONITOR_LEN = 1024 //最大护照监护信息长度
	MAX_NATIONALITY_LEN      = 16   //最大护照国籍长度
	MAX_PASSPORT_TYPE_LEN    = 4    //最大护照证件类型长度

	/*******************全局错误码 begin**********************/
	NET_DVR_NOERROR                 = 0   //没有错误
	NET_DVR_PASSWORD_ERROR          = 1   //用户名密码错误
	NET_DVR_NOENOUGHPRI             = 2   //权限不足
	NET_DVR_NOINIT                  = 3   //没有初始化
	NET_DVR_CHANNEL_ERROR           = 4   //通道号错误
	NET_DVR_OVER_MAXLINK            = 5   //连接到DVR的客户端个数超过最大
	NET_DVR_VERSIONNOMATCH          = 6   //版本不匹配
	NET_DVR_NETWORK_FAIL_CONNECT    = 7   //连接服务器失败
	NET_DVR_NETWORK_SEND_ERROR      = 8   //向服务器发送失败
	NET_DVR_NETWORK_RECV_ERROR      = 9   //从服务器接收数据失败
	NET_DVR_NETWORK_RECV_TIMEOUT    = 10  //从服务器接收数据超时
	NET_DVR_NETWORK_ERRORDATA       = 11  //传送的数据有误
	NET_DVR_ORDER_ERROR             = 12  //调用次序错误
	NET_DVR_OPERNOPERMIT            = 13  //无此权限
	NET_DVR_COMMANDTIMEOUT          = 14  //DVR命令执行超时
	NET_DVR_ERRORSERIALPORT         = 15  //串口号错误
	NET_DVR_ERRORALARMPORT          = 16  //报警端口错误
	NET_DVR_PARAMETER_ERROR         = 17  //参数错误
	NET_DVR_CHAN_EXCEPTION          = 18  //服务器通道处于错误状态
	NET_DVR_NODISK                  = 19  //没有硬盘
	NET_DVR_ERRORDISKNUM            = 20  //硬盘号错误
	NET_DVR_DISK_FULL               = 21  //服务器硬盘满
	NET_DVR_DISK_ERROR              = 22  //服务器硬盘出错
	NET_DVR_NOSUPPORT               = 23  //服务器不支持
	NET_DVR_BUSY                    = 24  //服务器忙
	NET_DVR_MODIFY_FAIL             = 25  //服务器修改不成功
	NET_DVR_PASSWORD_FORMAT_ERROR   = 26  //密码输入格式不正确
	NET_DVR_DISK_FORMATING          = 27  //硬盘正在格式化，不能启动操作
	NET_DVR_DVRNORESOURCE           = 28  //DVR资源不足
	NET_DVR_DVROPRATEFAILED         = 29  //DVR操作失败
	NET_DVR_OPENHOSTSOUND_FAIL      = 30  //打开PC声音失败
	NET_DVR_DVRVOICEOPENED          = 31  //服务器语音对讲被占用
	NET_DVR_TIMEINPUTERROR          = 32  //时间输入不正确
	NET_DVR_NOSPECFILE              = 33  //回放时服务器没有指定的文件
	NET_DVR_CREATEFILE_ERROR        = 34  //创建文件出错
	NET_DVR_FILEOPENFAIL            = 35  //打开文件出错
	NET_DVR_OPERNOTFINISH           = 36  //上次的操作还没有完成
	NET_DVR_GETPLAYTIMEFAIL         = 37  //获取当前播放的时间出错
	NET_DVR_PLAYFAIL                = 38  //播放出错
	NET_DVR_FILEFORMAT_ERROR        = 39  //文件格式不正确
	NET_DVR_DIR_ERROR               = 40  //路径错误
	NET_DVR_ALLOC_RESOURCE_ERROR    = 41  //资源分配错误
	NET_DVR_AUDIO_MODE_ERROR        = 42  //声卡模式错误
	NET_DVR_NOENOUGH_BUF            = 43  //缓冲区太小
	NET_DVR_CREATESOCKET_ERROR      = 44  //创建SOCKET出错
	NET_DVR_SETSOCKET_ERROR         = 45  //设置SOCKET出错
	NET_DVR_MAX_NUM                 = 46  //个数达到最大
	NET_DVR_USERNOTEXIST            = 47  //用户不存在
	NET_DVR_WRITEFLASHERROR         = 48  //写FLASH出错
	NET_DVR_UPGRADEFAIL             = 49  //DVR升级失败
	NET_DVR_CARDHAVEINIT            = 50  //解码卡已经初始化过
	NET_DVR_PLAYERFAILED            = 51  //调用播放库中某个函数失败
	NET_DVR_MAX_USERNUM             = 52  //设备端用户数达到最大
	NET_DVR_GETLOCALIPANDMACFAIL    = 53  //获得客户端的IP地址或物理地址失败
	NET_DVR_NOENCODEING             = 54  //该通道没有编码
	NET_DVR_IPMISMATCH              = 55  //IP地址不匹配
	NET_DVR_MACMISMATCH             = 56  //MAC地址不匹配
	NET_DVR_UPGRADELANGMISMATCH     = 57  //升级文件语言不匹配
	NET_DVR_MAX_PLAYERPORT          = 58  //播放器路数达到最大
	NET_DVR_NOSPACEBACKUP           = 59  //备份设备中没有足够空间进行备份
	NET_DVR_NODEVICEBACKUP          = 60  //没有找到指定的备份设备
	NET_DVR_PICTURE_BITS_ERROR      = 61  //图像素位数不符，限24色
	NET_DVR_PICTURE_DIMENSION_ERROR = 62  //图片高*宽超限， 限128*256
	NET_DVR_PICTURE_SIZ_ERROR       = 63  //图片大小超限，限100K
	NET_DVR_LOADPLAYERSDKFAILED     = 64  //载入当前目录下Player Sdk出错
	NET_DVR_LOADPLAYERSDKPROC_ERROR = 65  //找不到Player Sdk中某个函数入口
	NET_DVR_LOADDSSDKFAILED         = 66  //载入当前目录下DSsdk出错
	NET_DVR_LOADDSSDKPROC_ERROR     = 67  //找不到DsSdk中某个函数入口
	NET_DVR_DSSDK_ERROR             = 68  //调用硬解码库DsSdk中某个函数失败
	NET_DVR_VOICEMONOPOLIZE         = 69  //声卡被独占
	NET_DVR_JOINMULTICASTFAILED     = 70  //加入多播组失败
	NET_DVR_CREATEDIR_ERROR         = 71  //建立日志文件目录失败
	NET_DVR_BINDSOCKET_ERROR        = 72  //绑定套接字失败
	NET_DVR_SOCKETCLOSE_ERROR       = 73  //socket连接中断，此错误通常是由于连接中断或目的地不可达
	NET_DVR_USERID_ISUSING          = 74  //注销时用户ID正在进行某操作
	NET_DVR_SOCKETLISTEN_ERROR      = 75  //监听失败
	NET_DVR_PROGRAM_EXCEPTION       = 76  //程序异常
	NET_DVR_WRITEFILE_FAILED        = 77  //写文件失败
	NET_DVR_FORMAT_READONLY         = 78  //禁止格式化只读硬盘
	NET_DVR_WITHSAMEUSERNAME        = 79  //用户配置结构中存在相同的用户名
	NET_DVR_DEVICETYPE_ERROR        = 80  /*导入参数时设备型号不匹配*/
	NET_DVR_LANGUAGE_ERROR          = 81  /*导入参数时语言不匹配*/
	NET_DVR_PARAVERSION_ERROR       = 82  /*导入参数时软件版本不匹配*/
	NET_DVR_IPCHAN_NOTALIVE         = 83  /*预览时外接IP通道不在线*/
	NET_DVR_RTSP_SDK_ERROR          = 84  /*加载高清IPC通讯库StreamTransClient.dll失败*/
	NET_DVR_CONVERT_SDK_ERROR       = 85  /*加载转码库失败*/
	NET_DVR_IPC_COUNT_OVERFLOW      = 86  /*超出最大的ip接入通道数*/
	NET_DVR_MAX_ADD_NUM             = 87  /*添加标签(一个文件片段64)等个数达到最大*/
	NET_DVR_PARAMMODE_ERROR         = 88  //图像增强仪，参数模式错误（用于硬件设置时，客户端进行软件设置时错误值）
	NET_DVR_CODESPITTER_OFFLINE     = 89  //视频综合平台，码分器不在线
	NET_DVR_BACKUP_COPYING          = 90  //设备正在备份
	NET_DVR_CHAN_NOTSUPPORT         = 91  // 通道不支持该操作
	NET_DVR_CALLINEINVALID          = 92  // 高度线位置太集中或长度线不够倾斜
	NET_DVR_CALCANCELCONFLICT       = 93  // 取消标定冲突，如果设置了规则及全局的实际大小尺寸过滤
	NET_DVR_CALPOINTOUTRANGE        = 94  // 标定点超出范围
	NET_DVR_FILTERRECTINVALID       = 95  // 尺寸过滤器不符合要求
	NET_DVR_DDNS_DEVOFFLINE         = 96  //设备没有注册到ddns上
	NET_DVR_DDNS_INTER_ERROR        = 97  //DDNS 服务器内部错误
	NET_DVR_FUNCTION_NOT_SUPPORT_OS = 98  //此功能不支持该操作系统
	NET_DVR_DEC_CHAN_REBIND         = 99  //解码通道绑定显示输出次数受限
	NET_DVR_INTERCOM_SDK_ERROR      = 100 //加载当前目录下的语音对讲库失败
	NET_DVR_NO_CURRENT_UPDATEFILE   = 101 //没有正确的升级包
	NET_DVR_USER_NOT_SUCC_LOGIN     = 102 //用户还没登陆成功
	NET_DVR_USE_LOG_SWITCH_FILE     = 103 //正在使用日志开关文件
	NET_DVR_POOL_PORT_EXHAUST       = 104 //端口池中用于绑定的端口已耗尽
	NET_DVR_PACKET_TYPE_NOT_SUPPORT = 105 //码流封装格式错误
	NET_DVR_IPPARA_IPID_ERROR       = 106 //IP接入配置时IPID有误

	NET_DVR_LOAD_HCPREVIEW_SDK_ERROR       = 107 //预览组件加载失败
	NET_DVR_LOAD_HCVOICETALK_SDK_ERROR     = 108 //语音组件加载失败
	NET_DVR_LOAD_HCALARM_SDK_ERROR         = 109 //报警组件加载失败
	NET_DVR_LOAD_HCPLAYBACK_SDK_ERROR      = 110 //回放组件加载失败
	NET_DVR_LOAD_HCDISPLAY_SDK_ERROR       = 111 //显示组件加载失败
	NET_DVR_LOAD_HCINDUSTRY_SDK_ERROR      = 112 //行业应用组件加载失败
	NET_DVR_LOAD_HCGENERALCFGMGR_SDK_ERROR = 113 //通用配置管理组件加载失败
	NET_DVR_LOAD_HCCOREDEVCFG_SDK_ERROR    = 114 //设备配置核心组件加载失败
	NET_DVR_LOAD_HCNETUTILS_SDK_ERROR      = 115 //HCNetUtils加载失败

	NET_DVR_CORE_VER_MISMATCH                 = 121 //单独加载组件时，组件与core版本不匹配
	NET_DVR_CORE_VER_MISMATCH_HCPREVIEW       = 122 //预览组件与core版本不匹配
	NET_DVR_CORE_VER_MISMATCH_HCVOICETALK     = 123 //语音组件与core版本不匹配
	NET_DVR_CORE_VER_MISMATCH_HCALARM         = 124 //报警组件与core版本不匹配
	NET_DVR_CORE_VER_MISMATCH_HCPLAYBACK      = 125 //回放组件与core版本不匹配
	NET_DVR_CORE_VER_MISMATCH_HCDISPLAY       = 126 //显示组件与core版本不匹配
	NET_DVR_CORE_VER_MISMATCH_HCINDUSTRY      = 127 //行业应用组件与core版本不匹配
	NET_DVR_CORE_VER_MISMATCH_HCGENERALCFGMGR = 128 //通用配置管理组件与core版本不匹配

	NET_DVR_COM_VER_MISMATCH_HCPREVIEW       = 136 //预览组件与HCNetSDK版本不匹配
	NET_DVR_COM_VER_MISMATCH_HCVOICETALK     = 137 //语音组件与HCNetSDK版本不匹配
	NET_DVR_COM_VER_MISMATCH_HCALARM         = 138 //报警组件与HCNetSDK版本不匹配
	NET_DVR_COM_VER_MISMATCH_HCPLAYBACK      = 139 //回放组件与HCNetSDK版本不匹配
	NET_DVR_COM_VER_MISMATCH_HCDISPLAY       = 140 //显示组件与HCNetSDK版本不匹配
	NET_DVR_COM_VER_MISMATCH_HCINDUSTRY      = 141 //行业应用组件与HCNetSDK版本不匹配
	NET_DVR_COM_VER_MISMATCH_HCGENERALCFGMGR = 142 //通用配置管理组件与HCNetSDK版本不匹配

	NET_ERR_CONFIG_FILE_IMPORT_FAILED = 145 //配置文件导入失败
	NET_ERR_CONFIG_FILE_EXPORT_FAILED = 146 //配置文件导出失败
	NET_DVR_CERTIFICATE_FILE_ERROR    = 147 //证书错误
	NET_DVR_LOAD_SSL_LIB_ERROR        = 148 //加载SSL库失败（可能是版本不匹配，也可能是不存在）
	NET_DVR_SSL_VERSION_NOT_MATCH     = 149 //SSL库版本不匹配

	NET_DVR_ALIAS_DUPLICATE         = 150 //别名重复  //2011-08-31 通过别名或者序列号来访问设备的新版本ddns的配置
	NET_DVR_INVALID_COMMUNICATION   = 151 //无效通信
	NET_DVR_USERNAME_NOT_EXIST      = 152 //用户名不存在（用户名不存在，IPC5.1.7中发布出去了，所以删不掉。后续的产品这个错误码用不上）
	NET_DVR_USER_LOCKED             = 153 //用户被锁定
	NET_DVR_INVALID_USERID          = 154 //无效用户ID
	NET_DVR_LOW_LOGIN_VERSION       = 155 //登录版本低
	NET_DVR_LOAD_LIBEAY32_DLL_ERROR = 156 //加载libeay32.dll库失败
	NET_DVR_LOAD_SSLEAY32_DLL_ERROR = 157 //加载ssleay32.dll库失败
	NET_ERR_LOAD_LIBICONV           = 158 //加载libiconv库失败
	NET_ERR_SSL_CONNECT_FAILED      = 159 //SSL连接失败
	NET_ERR_MCAST_ADDRESS_ERROR     = 160 //获取多播地址错误
	NET_ERR_LOAD_ZLIB               = 161 //加载zlib.dll库失败
	NET_ERR_OPENSSL_NO_INIT         = 162 //Openssl库未初始化

	NET_DVR_SERVER_NOT_EXIST                         = 164 //对应的服务器找不到,查找时输入的国家编号或者服务器类型错误
	NET_DVR_TEST_SERVER_FAIL_CONNECT                 = 165 //连接测试服务器失败
	NET_DVR_NAS_SERVER_INVALID_DIR                   = 166 //NAS服务器挂载目录失败，目录无效
	NET_DVR_NAS_SERVER_NOENOUGH_PRI                  = 167 //NAS服务器挂载目录失败，没有权限
	NET_DVR_EMAIL_SERVER_NOT_CONFIG_DNS              = 168 //服务器使用域名，但是没有配置DNS，可能造成域名无效。
	NET_DVR_EMAIL_SERVER_NOT_CONFIG_GATEWAY          = 169 //没有配置网关，可能造成发送邮件失败。
	NET_DVR_TEST_SERVER_PASSWORD_ERROR               = 170 //用户名密码不正确，测试服务器的用户名或密码错误
	NET_DVR_EMAIL_SERVER_CONNECT_EXCEPTION_WITH_SMTP = 171 //设备和smtp服务器交互异常
	NET_DVR_FTP_SERVER_FAIL_CREATE_DIR               = 172 //FTP服务器创建目录失败
	NET_DVR_FTP_SERVER_NO_WRITE_PIR                  = 173 //FTP服务器没有写入权限
	NET_DVR_IP_CONFLICT                              = 174 //IP冲突
	NET_DVR_INSUFFICIENT_STORAGEPOOL_SPACE           = 175 //存储池空间已满
	NET_DVR_STORAGEPOOL_INVALID                      = 176 //云服务器存储池无效,没有配置存储池或者存储池ID错误
	NET_DVR_EFFECTIVENESS_REBOOT                     = 177 //生效需要重启
	NET_ERR_ANR_ARMING_EXIST                         = 178 //断网续传布防连接已经存在(该错误码是在HIK私有布防连接建立的情况下，重复布防的断网续传功能时，返回。)
	NET_ERR_UPLOADLINK_EXIST                         = 179 //断网续传上传连接已经存在(EHOME协议和HIK SDK协议是不能同时支持断网续传的，当一个协议存在的时候，另外一个连接建立话，报错这个错误码。)
	NET_ERR_INCORRECT_FILE_FORMAT                    = 180 //导入文件格式不正确
	NET_ERR_INCORRECT_FILE_CONTENT                   = 181 //导入文件内容不正确
	NET_ERR_MAX_HRUDP_LINK                           = 182 //HRUDP 连接数 超过设备限制
	NET_SDK_ERR_ACCESSKEY_SECRETKEY                  = 183 // 接入秘钥或加密秘钥不正确
	NET_SDK_ERR_CREATE_PORT_MULTIPLEX                = 184 //创建端口复用失败
	NET_DVR_NONBLOCKING_CAPTURE_NOTSUPPORT           = 185 //不支持无阻塞抓图
	NET_SDK_ERR_FUNCTION_INVALID                     = 186 //已开启异步，该功能无效
	NET_SDK_ERR_MAX_PORT_MULTIPLEX                   = 187 //已达到端口复用最大数目
	// 2010-5-28
	// 阵列错误码
	RAID_ERROR_INDEX                = 200
	NET_DVR_NAME_NOT_ONLY           = (RAID_ERROR_INDEX + 0)  // 名称已存在
	NET_DVR_OVER_MAX_ARRAY          = (RAID_ERROR_INDEX + 1)  // 阵列达到上限
	NET_DVR_OVER_MAX_VD             = (RAID_ERROR_INDEX + 2)  // 虚拟磁盘达到上限
	NET_DVR_VD_SLOT_EXCEED          = (RAID_ERROR_INDEX + 3)  // 虚拟磁盘槽位已满
	NET_DVR_PD_STATUS_INVALID       = (RAID_ERROR_INDEX + 4)  // 重建阵列所需物理磁盘状态错误
	NET_DVR_PD_BE_DEDICATE_SPARE    = (RAID_ERROR_INDEX + 5)  // 重建阵列所需物理磁盘为指定热备
	NET_DVR_PD_NOT_FREE             = (RAID_ERROR_INDEX + 6)  // 重建阵列所需物理磁盘非空闲
	NET_DVR_CANNOT_MIG2NEWMODE      = (RAID_ERROR_INDEX + 7)  // 不能从当前的阵列类型迁移到新的阵列类型
	NET_DVR_MIG_PAUSE               = (RAID_ERROR_INDEX + 8)  // 迁移操作已暂停
	NET_DVR_MIG_CANCEL              = (RAID_ERROR_INDEX + 9)  // 正在执行的迁移操作已取消
	NET_DVR_EXIST_VD                = (RAID_ERROR_INDEX + 10) // 阵列上阵列上存在虚拟磁盘，无法删除阵列
	NET_DVR_TARGET_IN_LD_FUNCTIONAL = (RAID_ERROR_INDEX + 11) // 对象物理磁盘为虚拟磁盘组成部分且工作正常
	NET_DVR_HD_IS_ASSIGNED_ALREADY  = (RAID_ERROR_INDEX + 12) // 指定的物理磁盘被分配为虚拟磁盘
	NET_DVR_INVALID_HD_COUNT        = (RAID_ERROR_INDEX + 13) // 物理磁盘数量与指定的RAID等级不匹配
	NET_DVR_LD_IS_FUNCTIONAL        = (RAID_ERROR_INDEX + 14) // 阵列正常，无法重建
	NET_DVR_BGA_RUNNING             = (RAID_ERROR_INDEX + 15) // 存在正在执行的后台任务
	NET_DVR_LD_NO_ATAPI             = (RAID_ERROR_INDEX + 16) // 无法用ATAPI盘创建虚拟磁盘
	NET_DVR_MIGRATION_NOT_NEED      = (RAID_ERROR_INDEX + 17) // 阵列无需迁移
	NET_DVR_HD_TYPE_MISMATCH        = (RAID_ERROR_INDEX + 18) // 物理磁盘不属于同意类型
	NET_DVR_NO_LD_IN_DG             = (RAID_ERROR_INDEX + 19) // 无虚拟磁盘，无法进行此项操作
	NET_DVR_NO_ROOM_FOR_SPARE       = (RAID_ERROR_INDEX + 20) // 磁盘空间过小，无法被指定为热备盘
	NET_DVR_SPARE_IS_IN_MULTI_DG    = (RAID_ERROR_INDEX + 21) // 磁盘已被分配为某阵列热备盘
	NET_DVR_DG_HAS_MISSING_PD       = (RAID_ERROR_INDEX + 22) // 阵列缺少盘

	// x86 64bit nvr新增 2012-02-04
	NET_DVR_NAME_EMPTY              = (RAID_ERROR_INDEX + 23) /*名称为空*/
	NET_DVR_INPUT_PARAM             = (RAID_ERROR_INDEX + 24) /*输入参数有误*/
	NET_DVR_PD_NOT_AVAILABLE        = (RAID_ERROR_INDEX + 25) /*物理磁盘不可用*/
	NET_DVR_ARRAY_NOT_AVAILABLE     = (RAID_ERROR_INDEX + 26) /*阵列不可用*/
	NET_DVR_PD_COUNT                = (RAID_ERROR_INDEX + 27) /*物理磁盘数不正确*/
	NET_DVR_VD_SMALL                = (RAID_ERROR_INDEX + 28) /*虚拟磁盘太小*/
	NET_DVR_NO_EXIST                = (RAID_ERROR_INDEX + 29) /*不存在*/
	NET_DVR_NOT_SUPPORT             = (RAID_ERROR_INDEX + 30) /*不支持该操作*/
	NET_DVR_NOT_FUNCTIONAL          = (RAID_ERROR_INDEX + 31) /*阵列状态不是正常状态*/
	NET_DVR_DEV_NODE_NOT_FOUND      = (RAID_ERROR_INDEX + 32) /*虚拟磁盘设备节点不存在*/
	NET_DVR_SLOT_EXCEED             = (RAID_ERROR_INDEX + 33) /*槽位达到上限*/
	NET_DVR_NO_VD_IN_ARRAY          = (RAID_ERROR_INDEX + 34) /*阵列上不存在虚拟磁盘*/
	NET_DVR_VD_SLOT_INVALID         = (RAID_ERROR_INDEX + 35) /*虚拟磁盘槽位无效*/
	NET_DVR_PD_NO_ENOUGH_SPACE      = (RAID_ERROR_INDEX + 36) /*所需物理磁盘空间不足*/
	NET_DVR_ARRAY_NONFUNCTION       = (RAID_ERROR_INDEX + 37) /*只有处于正常状态的阵列才能进行迁移*/
	NET_DVR_ARRAY_NO_ENOUGH_SPACE   = (RAID_ERROR_INDEX + 38) /*阵列空间不足*/
	NET_DVR_STOPPING_SCANNING_ARRAY = (RAID_ERROR_INDEX + 39) /*正在执行安全拔盘或重新扫描*/
	NET_DVR_NOT_SUPPORT_16T         = (RAID_ERROR_INDEX + 40) /*不支持创建大于16T的阵列*/
	NET_DVR_ARRAY_FORMATING         = (RAID_ERROR_INDEX + 41) /*正在执行格式化的阵列无法删除*/
	NET_DVR_QUICK_SETUP_PD_COUNT    = (RAID_ERROR_INDEX + 42) /*一键配置至少需要三块空闲盘*/

	//设备未激活时，登录失败，返回错误码
	NET_DVR_ERROR_DEVICE_NOT_ACTIVATED = 250 //设备未激活
	//老SDK接新设备，设置用户密码或者激活的时候为风险密码时，错误码
	NET_DVR_ERROR_RISK_PASSWORD = 251 //有风险的密码
	//已激活的设备，再次激活时返回错误码
	NET_DVR_ERROR_DEVICE_HAS_ACTIVATED = 252 //设备已激活

	// 智能错误码
	VCA_ERROR_INDEX                   = 300                    // 智能错误码索引
	NET_DVR_ID_ERROR                  = (VCA_ERROR_INDEX + 0)  // 配置ID不合理
	NET_DVR_POLYGON_ERROR             = (VCA_ERROR_INDEX + 1)  // 多边形不符合要求
	NET_DVR_RULE_PARAM_ERROR          = (VCA_ERROR_INDEX + 2)  // 规则参数不合理
	NET_DVR_RULE_CFG_CONFLICT         = (VCA_ERROR_INDEX + 3)  // 配置信息冲突
	NET_DVR_CALIBRATE_NOT_READY       = (VCA_ERROR_INDEX + 4)  // 当前没有标定信息
	NET_DVR_CAMERA_DATA_ERROR         = (VCA_ERROR_INDEX + 5)  // 摄像机参数不合理
	NET_DVR_CALIBRATE_DATA_UNFIT      = (VCA_ERROR_INDEX + 6)  // 长度不够倾斜，不利于标定
	NET_DVR_CALIBRATE_DATA_CONFLICT   = (VCA_ERROR_INDEX + 7)  // 标定出错，以为所有点共线或者位置太集中
	NET_DVR_CALIBRATE_CALC_FAIL       = (VCA_ERROR_INDEX + 8)  // 摄像机标定参数值计算失败
	NET_DVR_CALIBRATE_LINE_OUT_RECT   = (VCA_ERROR_INDEX + 9)  // 输入的样本标定线超出了样本外接矩形框
	NET_DVR_ENTER_RULE_NOT_READY      = (VCA_ERROR_INDEX + 10) // 没有设置进入区域
	NET_DVR_AID_RULE_NO_INCLUDE_LANE  = (VCA_ERROR_INDEX + 11) // 交通事件规则中没有包括车道（特值拥堵和逆行）
	NET_DVR_LANE_NOT_READY            = (VCA_ERROR_INDEX + 12) // 当前没有设置车道
	NET_DVR_RULE_INCLUDE_TWO_WAY      = (VCA_ERROR_INDEX + 13) // 事件规则中包含2种不同方向
	NET_DVR_LANE_TPS_RULE_CONFLICT    = (VCA_ERROR_INDEX + 14) // 车道和数据规则冲突
	NET_DVR_NOT_SUPPORT_EVENT_TYPE    = (VCA_ERROR_INDEX + 15) // 不支持的事件类型
	NET_DVR_LANE_NO_WAY               = (VCA_ERROR_INDEX + 16) // 车道没有方向
	NET_DVR_SIZE_FILTER_ERROR         = (VCA_ERROR_INDEX + 17) // 尺寸过滤框不合理
	NET_DVR_LIB_FFL_NO_FACE           = (VCA_ERROR_INDEX + 18) // 特征点定位时输入的图像没有人脸
	NET_DVR_LIB_FFL_IMG_TOO_SMALL     = (VCA_ERROR_INDEX + 19) // 特征点定位时输入的图像太小
	NET_DVR_LIB_FD_IMG_NO_FACE        = (VCA_ERROR_INDEX + 20) // 单张图像人脸检测时输入的图像没有人脸
	NET_DVR_LIB_FACE_TOO_SMALL        = (VCA_ERROR_INDEX + 21) // 建模时人脸太小
	NET_DVR_LIB_FACE_QUALITY_TOO_BAD  = (VCA_ERROR_INDEX + 22) // 建模时人脸图像质量太差
	NET_DVR_KEY_PARAM_ERR             = (VCA_ERROR_INDEX + 23) //高级参数设置错误
	NET_DVR_CALIBRATE_DATA_ERR        = (VCA_ERROR_INDEX + 24) //标定样本数目错误，或数据值错误，或样本点超出地平线
	NET_DVR_CALIBRATE_DISABLE_FAIL    = (VCA_ERROR_INDEX + 25) //所配置规则不允许取消标定
	NET_DVR_VCA_LIB_FD_SCALE_OUTRANGE = (VCA_ERROR_INDEX + 26) //最大过滤框的宽高最小值超过最小过滤框的宽高最大值两倍以上
	NET_DVR_LIB_FD_REGION_TOO_LARGE   = (VCA_ERROR_INDEX + 27) //当前检测区域范围过大。检测区最大为图像的2/3
	NET_DVR_TRIAL_OVERDUE             = (VCA_ERROR_INDEX + 28) //试用版评估期已结束
	NET_DVR_CONFIG_FILE_CONFLICT      = (VCA_ERROR_INDEX + 29) //设备类型与配置文件冲突（加密狗类型与现有分析仪配置不符错误码提示）
	//算法库相关错误码
	NET_DVR_FR_FPL_FAIL          = (VCA_ERROR_INDEX + 30) // 人脸特征点定位失败
	NET_DVR_FR_IQA_FAIL          = (VCA_ERROR_INDEX + 31) // 人脸评分失败
	NET_DVR_FR_FEM_FAIL          = (VCA_ERROR_INDEX + 32) // 人脸特征提取失败
	NET_DVR_FPL_DT_CONF_TOO_LOW  = (VCA_ERROR_INDEX + 33) // 特征点定位时人脸检测置信度过低
	NET_DVR_FPL_CONF_TOO_LOW     = (VCA_ERROR_INDEX + 34) // 特征点定位置信度过低
	NET_DVR_E_DATA_SIZE          = (VCA_ERROR_INDEX + 35) // 数据长度不匹配
	NET_DVR_FR_MODEL_VERSION_ERR = (VCA_ERROR_INDEX + 36) // 人脸模型数据中的模型版本错误
	NET_DVR_FR_FD_FAIL           = (VCA_ERROR_INDEX + 37) // 识别库中人脸检测失败
	NET_DVR_FA_NORMALIZE_ERR     = (VCA_ERROR_INDEX + 38) // 人脸归一化出错
	//其他错误码
	NET_DVR_DOG_PUSTREAM_NOT_MATCH     = (VCA_ERROR_INDEX + 39) // 加密狗与前端取流设备类型不匹配
	NET_DVR_DEV_PUSTREAM_NOT_MATCH     = (VCA_ERROR_INDEX + 40) // 前端取流设备版本不匹配
	NET_DVR_PUSTREAM_ALREADY_EXISTS    = (VCA_ERROR_INDEX + 41) // 设备的其他通道已经添加过该前端设备
	NET_DVR_SEARCH_CONNECT_FAILED      = (VCA_ERROR_INDEX + 42) // 连接检索服务器失败
	NET_DVR_INSUFFICIENT_DISK_SPACE    = (VCA_ERROR_INDEX + 43) // 可存储的硬盘空间不足
	NET_DVR_DATABASE_CONNECTION_FAILED = (VCA_ERROR_INDEX + 44) // 数据库连接失败
	NET_DVR_DATABASE_ADM_PW_ERROR      = (VCA_ERROR_INDEX + 45) // 数据库用户名、密码错误
	NET_DVR_DECODE_YUV                 = (VCA_ERROR_INDEX + 46) // 解码失败
	NET_DVR_IMAGE_RESOLUTION_ERROR     = (VCA_ERROR_INDEX + 47) //
	NET_DVR_CHAN_WORKMODE_ERROR        = (VCA_ERROR_INDEX + 48) //

	NET_DVR_RTSP_ERROR_NOENOUGHPRI    = 401 //无权限：服务器返回401时，转成这个错误码
	NET_DVR_RTSP_ERROR_ALLOC_RESOURCE = 402 //分配资源失败
	NET_DVR_RTSP_ERROR_PARAMETER      = 403 //参数错误
	NET_DVR_RTSP_ERROR_NO_URL         = 404 //指定的URL地址不存在：服务器返回404时，转成这个错误码
	NET_DVR_RTSP_ERROR_FORCE_STOP     = 406 //用户中途强行退出

	NET_DVR_RTSP_GETPORTFAILED        = 407 //rtsp 得到端口错误
	NET_DVR_RTSP_DESCRIBERROR         = 410 //rtsp decribe 交互错误
	NET_DVR_RTSP_DESCRIBESENDTIMEOUT  = 411 //rtsp decribe 发送超时
	NET_DVR_RTSP_DESCRIBESENDERROR    = 412 //rtsp decribe 发送失败
	NET_DVR_RTSP_DESCRIBERECVTIMEOUT  = 413 //rtsp decribe 接收超时
	NET_DVR_RTSP_DESCRIBERECVDATALOST = 414 //rtsp decribe 接收数据错误
	NET_DVR_RTSP_DESCRIBERECVERROR    = 415 //rtsp decribe 接收失败
	NET_DVR_RTSP_DESCRIBESERVERERR    = 416 //rtsp decribe 服务器返回错误状态

	NET_DVR_RTSP_SETUPERROR        = 420 //rtsp setup 交互错误
	NET_DVR_RTSP_SETUPSENDTIMEOUT  = 421 //rtsp setup 发送超时
	NET_DVR_RTSP_SETUPSENDERROR    = 422 //rtsp setup 发送错误
	NET_DVR_RTSP_SETUPRECVTIMEOUT  = 423 //rtsp setup 接收超时
	NET_DVR_RTSP_SETUPRECVDATALOST = 424 //rtsp setup 接收数据错误
	NET_DVR_RTSP_SETUPRECVERROR    = 425 //rtsp setup 接收失败
	NET_DVR_RTSP_OVER_MAX_CHAN     = 426 //超过服务器最大连接数，或者服务器资源不足，服务器返回453时，转成这个错误码。
	NET_DVR_RTSP_SETUPSERVERERR    = 427 //rtsp setup 服务器返回错误状态

	NET_DVR_RTSP_PLAYERROR        = 430 //rtsp play 交互错误
	NET_DVR_RTSP_PLAYSENDTIMEOUT  = 431 //rtsp play 发送超时
	NET_DVR_RTSP_PLAYSENDERROR    = 432 //rtsp play 发送错误
	NET_DVR_RTSP_PLAYRECVTIMEOUT  = 433 //rtsp play 接收超时
	NET_DVR_RTSP_PLAYRECVDATALOST = 434 //rtsp play 接收数据错误
	NET_DVR_RTSP_PLAYRECVERROR    = 435 //rtsp play 接收失败
	NET_DVR_RTSP_PLAYSERVERERR    = 436 //rtsp play 服务器返回错误状态

	NET_DVR_RTSP_TEARDOWNERROR        = 440 //rtsp teardown 交互错误
	NET_DVR_RTSP_TEARDOWNSENDTIMEOUT  = 441 //rtsp teardown 发送超时
	NET_DVR_RTSP_TEARDOWNSENDERROR    = 442 //rtsp teardown 发送错误
	NET_DVR_RTSP_TEARDOWNRECVTIMEOUT  = 443 //rtsp teardown 接收超时
	NET_DVR_RTSP_TEARDOWNRECVDATALOST = 444 //rtsp teardown 接收数据错误
	NET_DVR_RTSP_TEARDOWNRECVERROR    = 445 //rtsp teardown 接收失败
	NET_DVR_RTSP_TEARDOWNSERVERERR    = 446 //rtsp teardown 服务器返回错误状态

	NET_PLAYM4_NOERROR                = 500 //no error
	NET_PLAYM4_PARA_OVER              = 501 //input parameter is invalid;
	NET_PLAYM4_ORDER_ERROR            = 502 //The order of the function to be called is error.
	NET_PLAYM4_TIMER_ERROR            = 503 //Create multimedia clock failed;
	NET_PLAYM4_DEC_VIDEO_ERROR        = 504 //Decode video data failed.
	NET_PLAYM4_DEC_AUDIO_ERROR        = 505 //Decode audio data failed.
	NET_PLAYM4_ALLOC_MEMORY_ERROR     = 506 //Allocate memory failed.
	NET_PLAYM4_OPEN_FILE_ERROR        = 507 //Open the file failed.
	NET_PLAYM4_CREATE_OBJ_ERROR       = 508 //Create thread or event failed
	NET_PLAYM4_CREATE_DDRAW_ERROR     = 509 //Create DirectDraw object failed.
	NET_PLAYM4_CREATE_OFFSCREEN_ERROR = 510 //failed when creating off-screen surface.
	NET_PLAYM4_BUF_OVER               = 511 //buffer is overflow
	NET_PLAYM4_CREATE_SOUND_ERROR     = 512 //failed when creating audio device.
	NET_PLAYM4_SET_VOLUME_ERROR       = 513 //Set volume failed
	NET_PLAYM4_SUPPORT_FILE_ONLY      = 514 //The function only support play file.
	NET_PLAYM4_SUPPORT_STREAM_ONLY    = 515 //The function only support play stream.
	NET_PLAYM4_SYS_NOT_SUPPORT        = 516 //System not support.
	NET_PLAYM4_FILEHEADER_UNKNOWN     = 517 //No file header.
	NET_PLAYM4_VERSION_INCORRECT      = 518 //The version of decoder and encoder is not adapted.
	NET_PALYM4_INIT_DECODER_ERROR     = 519 //Initialize decoder failed.
	NET_PLAYM4_CHECK_FILE_ERROR       = 520 //The file data is unknown.
	NET_PLAYM4_INIT_TIMER_ERROR       = 521 //Initialize multimedia clock failed.
	NET_PLAYM4_BLT_ERROR              = 522 //Blt failed.
	NET_PLAYM4_UPDATE_ERROR           = 523 //Update failed.
	NET_PLAYM4_OPEN_FILE_ERROR_MULTI  = 524 //openfile error, streamtype is multi
	NET_PLAYM4_OPEN_FILE_ERROR_VIDEO  = 525 //openfile error, streamtype is video
	NET_PLAYM4_JPEG_COMPRESS_ERROR    = 526 //JPEG compress error
	NET_PLAYM4_EXTRACT_NOT_SUPPORT    = 527 //Don't support the version of this file.
	NET_PLAYM4_EXTRACT_DATA_ERROR     = 528 //extract video data failed.

	//转封装库错误码
	NET_CONVERT_ERROR_NOT_SUPPORT = 581 //convert not support

	//语音对讲库错误码
	NET_AUDIOINTERCOM_OK              = 600 //无错误
	NET_AUDIOINTECOM_ERR_NOTSUPORT    = 601 //不支持
	NET_AUDIOINTECOM_ERR_ALLOC_MEMERY = 602 //内存申请错误
	NET_AUDIOINTECOM_ERR_PARAMETER    = 603 //参数错误
	NET_AUDIOINTECOM_ERR_CALL_ORDER   = 604 //调用次序错误
	NET_AUDIOINTECOM_ERR_FIND_DEVICE  = 605 //未发现设备
	NET_AUDIOINTECOM_ERR_OPEN_DEVICE  = 606 //不能打开设备诶
	NET_AUDIOINTECOM_ERR_NO_CONTEXT   = 607 //设备上下文出错
	NET_AUDIOINTECOM_ERR_NO_WAVFILE   = 608 //WAV文件出错
	NET_AUDIOINTECOM_ERR_INVALID_TYPE = 609 //无效的WAV参数类型
	NET_AUDIOINTECOM_ERR_ENCODE_FAIL  = 610 //编码失败
	NET_AUDIOINTECOM_ERR_DECODE_FAIL  = 611 //解码失败
	NET_AUDIOINTECOM_ERR_NO_PLAYBACK  = 612 //播放失败
	NET_AUDIOINTECOM_ERR_DENOISE_FAIL = 613 //降噪失败
	NET_AUDIOINTECOM_ERR_UNKOWN       = 619 //未知错误

	NET_QOS_OK                                   = 700               //no error
	NET_QOS_ERROR                                = (NET_QOS_OK - 1)  //qos error
	NET_QOS_ERR_INVALID_ARGUMENTS                = (NET_QOS_OK - 2)  //invalid arguments
	NET_QOS_ERR_SESSION_NOT_FOUND                = (NET_QOS_OK - 3)  //session net found
	NET_QOS_ERR_LIB_NOT_INITIALIZED              = (NET_QOS_OK - 4)  //lib not initialized
	NET_QOS_ERR_OUTOFMEM                         = (NET_QOS_OK - 5)  //outtofmem
	NET_QOS_ERR_PACKET_UNKNOW                    = (NET_QOS_OK - 10) //packet unknow
	NET_QOS_ERR_PACKET_VERSION                   = (NET_QOS_OK - 11) //packet version error
	NET_QOS_ERR_PACKET_LENGTH                    = (NET_QOS_OK - 12) //packet length error
	NET_QOS_ERR_PACKET_TOO_BIG                   = (NET_QOS_OK - 13) //packet too big
	NET_QOS_ERR_SCHEDPARAMS_INVALID_BANDWIDTH    = (NET_QOS_OK - 20) //schedparams invalid bandwidth
	NET_QOS_ERR_SCHEDPARAMS_BAD_FRACTION         = (NET_QOS_OK - 21) //schedparams bad fraction
	NET_QOS_ERR_SCHEDPARAMS_BAD_MINIMUM_INTERVAL = (NET_QOS_OK - 22) //schedparams bad minimum interval

	NET_ERROR_TRUNK_LINE                      = 711 //子系统已被配成干线
	NET_ERROR_MIXED_JOINT                     = 712 //不能进行混合拼接
	NET_ERROR_DISPLAY_SWITCH                  = 713 //不能进行显示通道切换
	NET_ERROR_USED_BY_BIG_SCREEN              = 714 //解码资源被大屏占用
	NET_ERROR_USE_OTHER_DEC_RESOURCE          = 715 //不能使用其他解码子系统资源
	NET_ERROR_DISP_MODE_SWITCH                = 716 //显示通道显示状态切换中
	NET_ERROR_SCENE_USING                     = 717 //场景正在使用
	NET_ERR_NO_ENOUGH_DEC_RESOURCE            = 718 //解码资源不足
	NET_ERR_NO_ENOUGH_FREE_SHOW_RESOURCE      = 719 //畅显资源不足
	NET_ERR_NO_ENOUGH_VIDEO_MEMORY            = 720 //显存资源不足
	NET_ERR_MAX_VIDEO_NUM                     = 721 //一拖多资源不足
	NET_ERR_WIN_COVER_FREE_SHOW_AND_NORMAL    = 722 //窗口跨越了畅显输出口和非畅显输出口
	NET_ERR_FREE_SHOW_WIN_SPLIT               = 723 //畅显窗口不支持分屏
	NET_ERR_INAPPROPRIATE_WIN_FREE_SHOW       = 724 //不是输出口整数倍的窗口不支持开启畅显
	NET_DVR_TRANSPARENT_WIN_NOT_SUPPORT_SPLIT = 725 //开启透明度的窗口不支持分屏
	NET_DVR_SPLIT_WIN_NOT_SUPPORT_TRANSPARENT = 726 //开启多分屏的窗口不支持透明度设置
	NET_ERR_MAX_LOGO_NUM                      = 727 //logo数达到上限
	NET_ERR_MAX_WIN_LOOP_NUM                  = 728 //轮巡窗口数达到上限
	NET_ERR_VIRTUAL_LED_VERTICAL_CROSS        = 729 //虚拟LED不能纵向跨屏
	NET_ERR_MAX_VIRTUAL_LED_HEIGHT            = 730 //虚拟LED高度超限
	NET_ERR_VIRTUAL_LED_ILLEGAL_CHARACTER     = 731 //虚拟LED内容包含非法字符
	NET_ERR_BASEMAP_NOT_EXIST                 = 732 //底图图片不存在
	NET_ERR_LED_NOT_SUPPORT_VIRTUAL_LED       = 733 //LED屏幕不支持虚拟LED
	NET_ERR_LED_RESOLUTION_NOT_SUPPORT        = 734 //LED分辨率不支持
	NET_ERR_PLAN_OVERDUE                      = 735 //预案超期，不能再调用
	NET_ERR_PROCESSER_MAX_SCREEN_BLK          = 736 //单个处理器接入的信号跨越的屏幕个数超限
	NET_ERR_WND_SIZE_TOO_SMALL                = 737 //开窗窗口宽高太小
	NET_ERR_WND_SPLIT_NOT_SUPPORT_ROAM        = 738 //分屏窗口不支持漫游
	NET_ERR_OUTPUT_ONE_BOARD_ONE_WALL         = 739 //同一个子板的输出口只能绑定到同一面墙上
	NET_ERR_WND_CANNOT_LCD_AND_LED_OUTPUT     = 740 //窗口不能同时跨LCD和LED输出口
	NET_ERR_MAX_OSD_NUM                       = 741 //OSD数量达到最大

	NET_SDK_CANCEL_WND_TOPKEEP_ATTR_FIRST     = 751 //先取消置顶保持窗口的置顶保持属性才能进行置底操作
	NET_SDK_ERR_LED_SCREEN_CHECKING           = 752 //正在校正LED屏幕
	NET_SDK_ERR_NOT_SUPPORT_SINGLE_RESOLUTION = 753 //LCD/LED输出口绑定之后不支持单个输出口的分辨率配置
	NET_SDK_ERR_LED_RESOLUTION_MISMATCHED     = 754 //该输出口的LED分辨率和其他输出口的LED分辨率不匹配，需要满足同行等高、同列等宽

	NET_SDK_ERR_MAX_VIRTUAL_LED_WIDTH        = 755 //虚拟LED宽度超限，包括最大值和最小值
	NET_SDK_ERR_MAX_VIRTUAL_LED_IN_SCREEN    = 756 //单屏虚拟LED数量超限
	NET_SDK_ERR_MAX_VIRTUAL_LED_IN_WALL      = 757 //单墙虚拟LED数量超限
	NET_SDK_ERR_VIRTUAL_LED_OVERLAP          = 758 //虚拟LED重叠错误
	NET_SDK_ERR_VIRTUAL_LED_TYPE             = 759 //类型错误
	NET_SDK_ERR_VIRTUAL_LED_COLOUR           = 760 //颜色错误
	NET_SDK_ERR_VIRTUAL_LED_MOVE_DIRECTION   = 761 //移动方向错误
	NET_SDK_ERR_VIRTUAL_LED_MOVE_MODE        = 762 //移动模式错误
	NET_SDK_ERR_VIRTUAL_LED_MOVE_SPEED       = 763 //移动速度错误
	NET_SDK_ERR_VIRTUAL_LED_DISP_MODE        = 764 //显示模式有误
	NET_SDK_ERR_VIRTUAL_LED_NO               = 765 //虚拟LED序号错误
	NET_SDK_ERR_VIRTUAL_LED_PARA             = 766 //虚拟LED参数配置错误，包括结构体内其他参数
	NET_SDK_ERR_BASEMAP_POSITION             = 767 //底图窗口宽高参数错误
	NET_SDK_ERR_BASEMAP_PICTURE_LEN          = 768 //底图图片长度超限
	NET_SDK_ERR_BASEMAP_PICTURE_RESOLUTION   = 769 //底图图片分辨率错误
	NET_SDK_ERR_BASEMAP_PICTURE_FORMAT       = 770 //底图图片格式错误
	NET_SDK_ERR_MAX_VIRTUAL_LED_NUM          = 771 //设备支持的虚拟LED数量超限
	NET_SDK_ERR_MAX_TIME_VIRTUAL_LED_IN_WALL = 772 //单面电视墙支持的时间虚拟LED的数量超限

	NET_ERR_TERMINAL_BUSY = 780 //终端忙，终端处于会议中

	NET_ERR_DATA_RETURNED_ILLEGAL         = 790 //设备返回的数据不合法
	NET_DVR_FUNCTION_RESOURCE_USAGE_ERROR = 791 //设备其它功能占用资源，导致该功能无法开启

	NET_DVR_ERR_IMPORT_EMPTY_FILE        = 792 //导入文件为空
	NET_DVR_ERR_IMPORT_TOO_LARGE_FILE    = 793 //导入文件过大
	NET_DVR_ERR_BAD_IPV4_ADDRESS         = 794 //IPV4地址无效
	NET_DVR_ERR_BAD_NET_MASK             = 795 //子网掩码地址无效
	NET_DVR_ERR_INVALID_NET_GATE_ADDRESS = 796 //网关地址无效
	NET_DVR_ERR_BAD_DNS                  = 797 //DNS地址无效
	NET_DVR_ERR_ILLEGAL_PASSWORD         = 798 //密码不能包含用户名

	NET_DVR_DEV_NET_OVERFLOW                   = 800 //网络流量超过设备能力上限
	NET_DVR_STATUS_RECORDFILE_WRITING_NOT_LOCK = 801 //录像文件在录像，无法被锁定
	NET_DVR_STATUS_CANT_FORMAT_LITTLE_DISK     = 802 //由于硬盘太小无法格式化

	//N+1错误码
	NET_SDK_ERR_REMOTE_DISCONNECT  = 803 //远端无法连接
	NET_SDK_ERR_RD_ADD_RD          = 804 //备机不能添加备机
	NET_SDK_ERR_BACKUP_DISK_EXCEPT = 805 //备份盘异常
	NET_SDK_ERR_RD_LIMIT           = 806 //备机数已达上限
	NET_SDK_ERR_ADDED_RD_IS_WD     = 807 //添加的备机是工作机
	NET_SDK_ERR_ADD_ORDER_WRONG    = 808 //添加顺序出错，比如没有被工作机添加为备机，就添加工作机
	NET_SDK_ERR_WD_ADD_WD          = 809 //工作机不能添加工作机
	NET_SDK_ERR_WD_SERVICE_EXCETP  = 810 //工作机CVR服务异常
	NET_SDK_ERR_RD_SERVICE_EXCETP  = 811 //备机CVR服务异常
	NET_SDK_ERR_ADDED_WD_IS_RD     = 812 //添加的工作机是备机
	NET_SDK_ERR_PERFORMANCE_LIMIT  = 813 //性能达到上限
	NET_SDK_ERR_ADDED_DEVICE_EXIST = 814 //添加的设备已经存在

	//审讯机错误码
	NET_SDK_ERR_INQUEST_RESUMING  = 815 //审讯恢复中
	NET_SDK_ERR_RECORD_BACKUPING  = 816 //审讯备份中
	NET_SDK_ERR_DISK_PLAYING      = 817 //光盘回放中
	NET_SDK_ERR_INQUEST_STARTED   = 818 //审讯已开启
	NET_SDK_ERR_LOCAL_OPERATING   = 819 //本地操作进行中
	NET_SDK_ERR_INQUEST_NOT_START = 820 //审讯未开启
	//Netra3.1.0错误码
	NET_SDK_ERR_CHAN_AUDIO_BIND = 821 //通道未绑定或绑定语音对讲失败
	//云存储错误码
	NET_DVR_N_PLUS_ONE_MODE      = 822 //设备当前处于N+1模式
	NET_DVR_CLOUD_STORAGE_OPENED = 823 //云存储模式已开启

	NET_DVR_ERR_OPER_NOT_ALLOWED = 824 //设备处于N+0被接管状态，不允许该操作
	NET_DVR_ERR_NEED_RELOCATE    = 825 //设备处于N+0被接管状态，需要获取重定向信息，再重新操作

	//庭审主机错误码
	NET_SDK_ERR_IR_PORT_ERROR                        = 830 //红外输出口错误
	NET_SDK_ERR_IR_CMD_ERROR                         = 831 //红外输出口的命令号错误
	NET_SDK_ERR_NOT_INQUESTING                       = 832 //设备处于非审讯状态
	NET_SDK_ERR_INQUEST_NOT_PAUSED                   = 833 //设备处于非暂停状态
	NET_DVR_CHECK_PASSWORD_MISTAKE_ERROR             = 834 //校验密码错误
	NET_DVR_CHECK_PASSWORD_NULL_ERROR                = 835 //校验密码不能为空
	NET_DVR_UNABLE_CALIB_ERROR                       = 836 // 当前无法标定
	NET_DVR_PLEASE_CALIB_ERROR                       = 837 //请先完成标定
	NET_DVR_ERR_PANORAMIC_CAL_EMPTY                  = 838 //Flash中全景标定为空
	NET_DVR_ERR_CALIB_FAIL_PLEASEAGAIN               = 839 //标定失败，请重新标定(Calibration failed. Please calibrate again.)
	NET_DVR_ERR_DETECTION_LINE                       = 840 //规则线配置错误，请重新配置规则线，确保规则线位于红色区域内(Please set detection line again. The detection line should be within the red count area.)
	NET_DVR_ERR_TURN_OFF_IMAGE_PARA                  = 841 //请先关闭图像参数切换功能(Please turn off the image parameters switch first.)
	NET_DVR_EXCEED_FACE_IMAGES_ERROR                 = 843 //超过人脸图片最大张数
	NET_DVR_ANALYSIS_FACE_IMAGES_ERROR               = 844 //图片数据识别失败
	NET_ERR_ALARM_INPUT_OCCUPIED                     = 845 //A<-1报警号已用于触发车辆抓拍Alarm Input No. A<-1 is used to trigger vehicle capture.
	NET_DVR_FACELIB_DATABASE_ERROR                   = 846 //人脸库中数据库版本不匹配
	NET_DVR_FACELIB_DATA_ERROR                       = 847 //人脸库数据错误
	NET_DVR_FACE_DATA_ID_ERROR                       = 848 //人脸数据PID无效
	NET_DVR_FACELIB_ID_ERROR                         = 849 //人脸库ID无效
	NET_DVR_EXCEED_FACE_LIBARY_ERROR                 = 850 //超过人脸库最大个数
	NET_DVR_PIC_ANALYSIS_NO_TARGET_ERROR             = 851 //图片未识别到目标
	NET_DVR_SUBPIC_ANALYSIS_MODELING_ERROR           = 852 //子图建模失败
	NET_DVR_PIC_ANALYSIS_NO_RESOURCE_ERROR           = 853 //无对应智能分析引擎支持图片二次识别
	NET_DVR_ANALYSIS_ENGINES_NO_RESOURCE_ERROR       = 854 //无分析引擎资源
	NET_DVR_ANALYSIS_ENGINES_USAGE_EXCEED_ERROR      = 855 //引擎使用率超负荷，已达100%
	NET_DVR_EXCEED_HUMANMISINFO_FILTER_ENABLED_ERROR = 856 //超过开启人体去误报最大通道个数
	NET_DVR_NAME_ERROR                               = 857 //名称错误
	NET_DVR_NAME_EXIST_ERROR                         = 858 //名称已存在
	NET_DVR_FACELIB_PIC_IMPORTING_ERROR              = 859 //人脸库导入图片中
	NET_DVR_ERR_CALIB_POSITION                       = 860 //标定位置超出摄像机运动范围
	NET_DVR_ERR_DELETE                               = 861 //无法删除
	NET_DVR_ERR_SCENE_ID                             = 862 //场景ID无效
	NET_DVR_ERR_CALIBING                             = 863 //标定中
	NET_DVR_PIC_FORMAT_ERROR                         = 864 //图片格式错误
	NET_DVR_PIC_RESOLUTION_INVALID_ERROR             = 865 //图片分辨率无效错误
	NET_DVR_PIC_SIZE_EXCEED_ERROR                    = 866 //图片过大
	NET_DVR_PIC_ANALYSIS_TARGRT_NUM_EXCEED_ERROR     = 867 //图片目标个数超过上限
	NET_DVR_ANALYSIS_ENGINES_LOADING_ERROR           = 868 //分析引擎初始化中
	NET_DVR_ANALYSIS_ENGINES_ABNORMA_ERROR           = 869 //分析引擎异常
	NET_DVR_ANALYSIS_ENGINES_FACELIB_IMPORTING       = 870 //分析引擎正在导入人脸库
	NET_DVR_NO_DATA_FOR_MODELING_ERROR               = 871 //无待建模数据
	NET_DVR_FACE_DATA_MODELING_ERROR                 = 872 //设备正在进行图片建模操作，不支持并发处理
	NET_ERR_FACELIBDATA_OVERLIMIT                    = 873 //超过设备中支持导入人脸数最大个数限制（导入的人脸库中数据）
	NET_DVR_ANALYSIS_ENGINES_ASSOCIATED_CHANNEL      = 874 //分析引擎已关联通道
	NET_DVR_ERR_CUSTOMID_LEN                         = 875 //上层自定义ID的长度最小32字符长度
	NET_DVR_ERR_CUSTOMFACELIBID_REPEAT               = 876 //上层下发重复的自定义人脸库ID
	NET_DVR_ERR_CUSTOMHUMANID_REPEAT                 = 877 //上层下发重复的自定义人员ID
	NET_DVR_ERR_URL_DOWNLOAD_FAIL                    = 878 //url下载失败
	NET_DVR_ERR_URL_DOWNLOAD_NOTSTART                = 879 //url未开始下载

	NET_DVR_CFG_FILE_SECRETKEY_ERROR = 880 //配置文件安全校验密钥错误
	NET_DVR_WDR_NOTDISABLE_ERROR     = 881 //请先关闭所有通道当前日夜参数转换模式下的宽动态
	NET_DVR_HLC_NOTDISABLE_ERROR     = 882 //请先关闭所有通道当前日夜参数转换模式下的强光抑制

	NET_DVR_THERMOMETRY_REGION_OVERSTEP_ERROR = 883 //测温区域越界

	NET_DVR_ERR_MODELING_DEVICEINTERNAL = 884 //建模失败，设备内部错误
	NET_DVR_ERR_MODELING_FACE           = 885 //建模失败，人脸建模错误
	NET_DVR_ERR_MODELING_FACEGRADING    = 886 //建模失败，人脸质量评分错误
	NET_DVR_ERR_MODELING_FACEGFEATURE   = 887 //建模失败，特征点提取错误
	NET_DVR_ERR_MODELING_FACEGANALYZING = 888 //建模失败，属性提取错误

	NET_DVR_ERR_STREAM_LIMIT            = 889 //码流性能超过上限，请减少取流路数
	NET_DVR_ERR_STREAM_DESCRIPTION      = 890 //请输入码流描述
	NET_DVR_ERR_STREAM_DELETE           = 891 //码流正在使用无法删除
	NET_DVR_ERR_CUSTOMSTREAM_NAME       = 892 //自定义码流名称为空或不合法
	NET_DVR_ERR_CUSTOMSTREAM_NOTEXISTED = 893 //该自定义码流不存在

	NET_DVR_ERR_TOO_SHORT_CALIBRATING_TIME = 894 //标定时间太短
	NET_DVR_ERR_AUTO_CALIBRATE_FAILED      = 895 //自动标定失败
	NET_DVR_ERR_VERIFICATION_FAILED        = 896 //校验失败

	NET_DVR_NO_TEMP_SENSOR_ERROR          = 897 //无温度传感器
	NET_DVR_PUPIL_DISTANCE_OVERSIZE_ERROR = 898 //瞳距过大
	NET_DVR_ERR_UNOPENED_FACE_SNAP        = 899 //操作无效，请先开启人脸抓拍
	//2011-10-25多屏控制器错误码（900-950）
	NET_ERR_CUT_INPUTSTREAM_OVERLIMIT   = 900 //信号源裁剪数值超限
	NET_ERR_WINCHAN_IDX                 = 901 // 开窗通道号错误
	NET_ERR_WIN_LAYER                   = 902 // 窗口层数错误，单个屏幕上最多覆盖的窗口层数
	NET_ERR_WIN_BLK_NUM                 = 903 // 窗口的块数错误，单个窗口可覆盖的屏幕个数
	NET_ERR_OUTPUT_RESOLUTION           = 904 // 输出分辨率错误
	NET_ERR_LAYOUT                      = 905 // 布局号错误
	NET_ERR_INPUT_RESOLUTION            = 906 // 输入分辨率不支持
	NET_ERR_SUBDEVICE_OFFLINE           = 907 // 子设备不在线
	NET_ERR_NO_DECODE_CHAN              = 908 // 没有空闲解码通道
	NET_ERR_MAX_WINDOW_ABILITY          = 909 // 开窗能力上限, 分布式多屏控制器中解码子设备能力上限或者显示处理器能力上限导致
	NET_ERR_ORDER_ERROR                 = 910 // 调用顺序有误
	NET_ERR_PLAYING_PLAN                = 911 // 正在执行预案
	NET_ERR_DECODER_USED                = 912 // 解码板正在使用
	NET_ERR_OUTPUT_BOARD_DATA_OVERFLOW  = 913 // 输出板数据量超限
	NET_ERR_SAME_USER_NAME              = 914 // 用户名相同
	NET_ERR_INVALID_USER_NAME           = 915 // 无效用户名
	NET_ERR_MATRIX_USING                = 916 // 输入矩阵正在使用
	NET_ERR_DIFFERENT_CHAN_TYPE         = 917 // 通道类型不同（矩阵输出通道和控制器的输入为不同的类型）
	NET_ERR_INPUT_CHAN_BINDED           = 918 // 输入通道已经被其他矩阵绑定
	NET_ERR_BINDED_OUTPUT_CHAN_OVERFLOW = 919 // 正在使用的矩阵输出通道个数超过矩阵与控制器绑定的通道个数
	NET_ERR_MAX_SIGNAL_NUM              = 920 // 输入信号源个数达到上限
	NET_ERR_INPUT_CHAN_USING            = 921 // 输入通道正在使用
	NET_ERR_MANAGER_LOGON               = 922 // 管理员已经登陆，操作失败
	NET_ERR_USERALREADY_LOGON           = 923 // 该用户已经登陆，操作失败
	NET_ERR_LAYOUT_INIT                 = 924 // 布局正在初始化，操作失败
	NET_ERR_BASEMAP_SIZE_NOT_MATCH      = 925 // 底图大小不符
	NET_ERR_WINDOW_OPERATING            = 926 // 窗口正在执行其他操作，本次操作失败
	NET_ERR_SIGNAL_UPLIMIT              = 927 // 信号源开窗个数达到上限
	NET_ERR_SIGNAL_MAX_ENLARGE_TIMES    = 928 // 信号源放大倍数超限
	NET_ERR_ONE_SIGNAL_MULTI_CROSS      = 929 // 单个信号源不能多次跨屏
	NET_ERR_ULTRA_HD_SIGNAL_MULTI_WIN   = 930 // 超高清信号源不能重复开窗
	NET_ERR_MAX_VIRTUAL_LED_WIDTH       = 931 //虚拟LED宽度大于限制值
	NET_ERR_MAX_VIRTUAL_LED_WORD_LEN    = 932 //虚拟LED字符数大于限制值
	NET_ERR_SINGLE_OUTPUTPARAM_CONFIG   = 933 //不支持单个显示输出参数设置
	NET_ERR_MULTI_WIN_BE_COVER          = 934 //多分屏窗口被覆盖
	NET_ERR_WIN_NOT_EXIST               = 935 //窗口不存在
	NET_ERR_WIN_MAX_SIGNALSOURCE        = 936 //窗口信号源数超过限制值
	NET_ERR_MULTI_WIN_MOVE              = 937 //对多分屏窗口移动
	NET_ERR_MULTI_WIN_YPBPR_SDI         = 938 // YPBPR 和SDI信号源不支持9/16分屏
	NET_ERR_DIFF_TYPE_OUTPUT_MIXUSE     = 939 //不同类型输出板混插
	NET_ERR_SPLIT_WIN_CROSS             = 940 //对跨屏窗口分屏
	NET_ERR_SPLIT_WIN_NOT_FULL_SCREEN   = 941 //对未满屏窗口分屏
	NET_ERR_SPLIT_WIN_MANY_WIN          = 942 //对单个输出口上有多个窗口的窗口分屏
	NET_ERR_WINDOW_SIZE_OVERLIMIT       = 943 //窗口大小超限
	NET_ERR_INPUTSTREAM_ALREADY_JOINT   = 944 //信号源已加入拼接
	NET_ERR_JOINT_INPUTSTREAM_OVERLIMIT = 945 //拼接信号源个数超限

	NET_ERR_LED_RESOLUTION                 = 946 //LED 分辨率大于输出分辨率
	NET_ERR_JOINT_SCALE_OVERLIMIT          = 947 //拼接信号源的规模超限
	NET_ERR_INPUTSTREAM_ALREADY_DECODE     = 948 //信号源已上墙
	NET_ERR_INPUTSTREAM_NOTSUPPORT_CAPTURE = 949 //信号源不支持抓图
	NET_ERR_JOINT_NOTSUPPORT_SPLITWIN      = 950 //拼接信号源不支持分屏

	//解码器错误码（951-999）
	NET_ERR_MAX_WIN_OVERLAP          = 951 //达到最大窗口重叠数
	NET_ERR_STREAMID_CHAN_BOTH_VALID = 952 //stream ID和通道号同时有效
	NET_ERR_NO_ZERO_CHAN             = 953 //设备无零通道
	NEED_RECONNECT                   = 955 //需要重定向（转码子系统使用）
	NET_ERR_NO_STREAM_ID             = 956 //流ID不存在
	NET_DVR_TRANS_NOT_START          = 957 //转码未启动
	NET_ERR_MAXNUM_STREAM_ID         = 958 //流ID数达到上限
	NET_ERR_WORKMODE_MISMATCH        = 959 //工作模式不匹配
	NET_ERR_MODE_IS_USING            = 960 //已工作在当前模式
	NET_ERR_DEV_PROGRESSING          = 961 //设备正在处理中
	NET_ERR_PASSIVE_TRANSCODING      = 962 //正在被动转码

	NET_ERR_RING_NOT_CONFIGURE = 964 //环网未配置

	NET_ERR_CLOSE_WINDOW_FIRST                  = 971 //切换全帧率畅显时必须先关闭对应的已上墙的窗口
	NET_ERR_SPLIT_WINDOW_NUM_NOT_SUPPORT        = 972 //VGA/DVI/DP/HDMI/HDBase_T输入源在全帧率畅显下不支持9/16画面
	NET_ERR_REACH_ONE_SIGNAL_PREVIEW_MAX_LINK   = 973 //单信号源回显连接数量超限
	NET_ERR_ONLY_SPLITWND_SUPPORT_AMPLIFICATION = 974 //只有分屏窗口支持子窗口放大
	NET_DVR_ERR_WINDOW_SIZE_PLACE               = 975 //窗口位置错误
	NET_DVR_ERR_RGIONAL_RESTRICTIONS            = 976 //屏幕距离超限
	NET_ERR_WNDZOOM_NOT_SUPPORT                 = 977 //单窗口不支持子窗口全屏功能
	NET_ERR_LED_SCREEN_SIZE                     = 978 //LED屏宽高不正确
	NET_ERR_OPEN_WIN_IN_ERROR_AREA              = 979 //在非法区域开窗(包括跨LCD/LED屏)
	NET_ERR_TITLE_WIN_NOT_SUPPORT_MOVE          = 980 //平铺模式不支持漫游
	NET_ERR_TITLE_WIN_NOT_SUPPORT_COVER         = 981 //平铺模式不支持图层覆盖
	NET_ERR_TITLE_WIN_NOT_SUPPORT_SPLIT         = 982 //平铺模式不支持分屏
	NET_DVR_LED_WINDOWS_ALREADY_CLOSED          = 983 //LED区域内输出口的分辨率发生变化，设备已关闭该区域内的所有LED窗口
	NET_DVR_ERR_CLOSE_WINDOWS                   = 984 //操作失败，请先关闭窗口
	NET_DVR_ERR_MATRIX_LOOP_ABILITY             = 985 //超出轮巡解码能力限制
	NET_DVR_ERR_MATRIX_LOOP_TIME                = 986 //轮巡解码时间不支持
	NET_DVR_ERR_LINKED_OUT_ABILITY              = 987 //联动通道数超过上限
	NET_ERR_REACH_SCENE_MAX_NUM                 = 988 //场景数量达到上限
	NET_ERR_SCENE_MEM_NOT_ENOUGH                = 989 //内存不足，无法新建场景
	NET_ERR_RESOLUTION_NOT_SUPPORT_ODD_VOUT     = 990 //奇口不支持该分辨率
	NET_ERR_RESOLUTION_NOT_SUPPORT_EVEN_VOUT    = 991 //偶口不支持该分辨率

	NET_DVR_CANCEL_WND_OPENKEEP_ATTR_FIRST        = 992 //常开窗口需要先取消常开属性才能被关闭
	NET_SDK_LED_MODE_NOT_SUPPORT_SPLIT            = 993 //LED模式下不支持窗口分屏
	NET_ERR_VOICETALK_ONLY_SUPPORT_ONE_TALK       = 994 //同时只支持一路语音对讲
	NET_ERR_WND_POSITION_ADJUSTED                 = 995 //窗口位置被设备调整，上层需要重新获取下窗口位置
	NET_SDK_ERR_STARTTIME_CANNOT_LESSTHAN_CURTIME = 996 //开始时间不能小于当前时间
	NET_SDK_ERR_NEED_ADJUST_PLAN                  = 997 //场景已被预案关联，请先将该场景从预案中删除
	NET_ERR_UnitConfig_Failed                     = 998 //当“启用单位统一”勾选时，测温下配置的单位与系统设置下的单位不同返回单位配置错误

	//能力集解析库错误码
	XML_ABILITY_NOTSUPPORT            = 1000 //不支持能力节点获取
	XML_ANALYZE_NOENOUGH_BUF          = 1001 //输出内存不足
	XML_ANALYZE_FIND_LOCALXML_ERROR   = 1002 //无法找到对应的本地xml
	XML_ANALYZE_LOAD_LOCALXML_ERROR   = 1003 //加载本地xml出错
	XML_NANLYZE_DVR_DATA_FORMAT_ERROR = 1004 //设备能力数据格式错误
	XML_ANALYZE_TYPE_ERROR            = 1005 //能力集类型错误
	XML_ANALYZE_XML_NODE_ERROR        = 1006 //XML能力节点格式错误
	XML_INPUT_PARAM_ERROR             = 1007 //输入的能力XML节点值错误

	NET_DVR_ERR_RETURNED_XML_DATA = 1008 //设备返回的XML数据有误

	//传显错误码
	NET_ERR_LEDAREA_EXIST_WINDOW   = 1051 //LED区域有窗口存在(如果LED区域上已经有窗口存在，不允许修改LED区域）
	NET_ERR_AUDIO_EXIST            = 1052 //输出口上存在音频输出，不允许解除绑定
	NET_ERR_MATERIAL_NAME_EXIST    = 1053 //素材名称已存在
	NET_ERR_MATERIAL_APPROVE_STATE = 1054 //素材审核状态错误
	NET_ERR_DATAHD_SIGNAL_FORMAT   = 1055 //已使用的硬盘不允许单个格式化

	NET_ERR_SCENE_SWITCHING                                 = 1056 //场景正在切换
	NER_ERR_DATA_TRANSFER                                   = 1057 //设备正在数据转移中
	NET_ERR_DATA_RESTORE                                    = 1058 //设备正在数据还原中
	NET_ERR_CHECK_NOT_ENABLE                                = 1059 //校正使能未开启
	NET_ERR_AREA_OFFLINE                                    = 1060 //区域不在线
	NET_ERR_SCREEN_TYPE                                     = 1061 //屏幕类型不匹配
	NET_ERR_MIN_OPERATE_UNIT                                = 1062 //最小操作单元不匹配
	NET_ERR_MAINHD_NOT_BACKUP                               = 1063 //第一槽位上的普通盘（主盘）禁止设置成备份盘
	NET_ERR_ONE_BACKUP_HD                                   = 1064 //设置备份盘时，设备至少有一块普通盘
	NET_ERR_CONNECT_SUB_SYSTEM_ABNORMAL                     = 1065 //连接子系统异常
	NET_ERR_SERIAL_PORT_VEST                                = 1066 //错误的串口归属
	NET_ERR_WHITE_LIST_FULL                                 = 1067 //白名单列表数量已满
	NET_ERR_NOT_MATCH_SOURCE                                = 1068 //不匹配的信号源类型
	NET_ERR_CLOCK_VIRTUAL_LED_FULL                          = 1069 //开启时钟功能的虚拟LED达上限
	NET_ERR_MAX_WIN_SIGNAL_LOOP_NUM                         = 1070 //窗口轮巡信号源个数已达上限
	NET_ERR_RESOLUTION_NO_MATCH_FRAME                       = 1071 //分辨率与当前帧数不匹配
	NET_ERR_NOT_UPDATE_LOW_VERSION                          = 1072 //不支持向低版本升级
	NET_ERR_NO_CUSTOM_TO_UPDATE                             = 1073 //非定制程序无法升级
	NET_ERR_CHAN_RESOLUTION_NOT_SUPPORT_SPLIT               = 1074 //该输出口分辨率不支持分屏
	NET_ERR_HIGH_DEFINITION_NOT_SUPPORT_SPLIT               = 1075 //超高清不支持9/16画面分割
	NET_ERR_MIRROR_IMAGE_BY_VIDEO_WALL                      = 1076 //电视墙镜像出错
	NET_ERR_MAX_OSD_FONT_SIZE                               = 1077 //超过OSD最大支持字符数
	NET_ERR_HIGH_DEFINITION_NOT_SUPPORT_VIDEO_SET           = 1078 //超清不支持视频参数设置
	NET_ERR_TILE_MODE_NOT_SUPPORT_JOINT                     = 1079 //平铺模式不支持拼接窗口
	NET_ERR_ADD_AUDIO_MATRIX_FAILED                         = 1080 //创建音频矩阵失败
	NET_ERR_ONE_VIRTUAL_LED_AREA_BIND_ONE_AUDIO_AREA        = 1081 //一个虚拟LED区域只能绑定一个音频区域
	NET_ERR_NAT_NOT_MODIFY_SERVER_NETWORK_PARAM             = 1082 //NAT下无法修改服务器网络参数
	NET_ERR_ORIGINAL_CHECH_DATA_ERROR                       = 1083 //原始校正数据错误
	NET_ERR_INPUT_BOARD_SPLICED_IN_DIFFERENT_NETWORKAREAS   = 1084 //不同网络区域的输入板不能拼接
	NET_ERR_SPLICINGSOURCE_ONWALL_IN_DIFFERENT_NETWORKAREAS = 1085 //不同网络区域的拼接源不能上墙
	NET_ERR_ONWALL_OUTPUTBOARD_MODIFY_NETWORKAREAS          = 1086 //已绑定在电视墙上的输出板不能修改网络区域
	NET_ERR_LAN_AND_WAN_CANNOT_SAME_NET_SEGMENT             = 1087 //LAN口IP和WAN口IP不能处于同一网段
	NET_ERR_USERNAME_REPETITIVE                             = 1088 //用户名重复
	NET_ERR_ASSOCIATED_SAMEWALL_IN_DIFFERENT_NETWORKAREAS   = 1089 //不同数据网络区域的输出口不能关联到同一电视墙
	NET_ERR_BASEMAP_ROAM_IN_LED_AREA                        = 1090 //LED区域不允许底图漫游
	NET_ERR_VIRTUAL_LED_NOT_SUPPORT_4K_OUTPUT               = 1091 //虚拟LED不支持4K输出口显示
	NET_ERR_BASEMAP_NOT_SUPPORT_4K_OUTPUT                   = 1092 //底图不支持4K输出口显示
	NET_ERR_MIN_BLOCK_IN_VIRTUAL_LED_AND_OUTPUT             = 1093 //虚拟LED与输出口相交出现最小块
	NET_ERR_485FIlE_VERSION_INVALID                         = 1094 //485文件版本无效
	NET_ERR_485FIlE_CHECK_ERROR                             = 1095 //485文件校验错误
	NET_ERR_485FIlE_ABNORMAL_SIZE                           = 1096 //485文件大小异常效
	NET_ERR_MODIFY_SUBBOARD_NETCFG_IN_NAT                   = 1097 //NAT下无法修改子板网络参
	NET_ERR_OSD_CONTENT_WITH_ILLEGAL_CHARACTERS             = 1098 //OSD内容包含非法字符
	NET_ERR_NON_SLAVE_DEVICE_INSERT_SYNC_LINE               = 1099 //非从设备禁止插入同步线
	//民用错误码（1100～1200）
	NET_ERR_PLT_USERID                    = 1100 //验证平台userid错误
	NET_ERR_TRANS_CHAN_START              = 1101 //透明通道已打开，当前操作无法完成
	NET_ERR_DEV_UPGRADING                 = 1102 //设备正在升级
	NET_ERR_MISMATCH_UPGRADE_PACK_TYPE    = 1103 //升级包类型不匹配
	NET_ERR_DEV_FORMATTING                = 1104 //设备正在格式化
	NET_ERR_MISMATCH_UPGRADE_PACK_VERSION = 1105 //升级包版本不匹配
	NET_ERR_PT_LOCKED                     = 1106 //PT锁定功能

	NET_DVR_LOGO_OVERLAY_WITHOUT_UPLOAD_PIC                          = 1110 //logo叠加失败，没有上传logo图片
	NET_DVR_ERR_ILLEGAL_VERIFICATION_CODE                            = 1111 //不合法验证码
	NET_DVR_ERR_LACK_VERIFICATION_CODE                               = 1112 //缺少验证码
	NET_DVR_ERR_FORBIDDEN_IP                                         = 1113 //该IP地址已被禁止，不允许配置(设备支持的IP地址过滤功能)
	NET_DVR_ERR_UNLOCKPTZ                                            = 1114 //操作无效，请先锁定云台
	NET_DVR_ERR_COUNTAREA_LARGE                                      = 1116 //计数区域绘制错误，区域面积大于设备允许值
	NET_DVR_ERR_LABEL_ID_EXCEED                                      = 1117 //标签ID超限
	NET_DVR_ERR_LABEL_TYPE                                           = 1118 //标签类型错误
	NET_DVR_ERR_LABEL_FULL                                           = 1119 //标签满
	NET_DVR_ERR_LABEL_DISABLED                                       = 1120 //标签未使能
	NET_DVR_ERR_DOME_PT_TRANS_TO_DOME_XY                             = 1121 //球机PT转球机XY失败
	NET_DVR_ERR_DOME_PT_TRANS_TO_PANORAMA_XY                         = 1122 //球机PT转全景XY失败
	NET_DVR_ERR_PANORAMA_XY_TRANS_TO_DOME_PT                         = 1123 //全景XY坐标转球机PT错误
	NET_DVR_ERR_SCENE_DUR_TIME_LESS_THAN_INTERV_TIME                 = 1124 //场景停留时间要大于检测周期
	NET_DVR_ERR_HTTP_BKN_EXCEED_ONE                                  = 1125 //断网续传布防只支持一路
	NET_DVR_ERR_DELETING_FAILED_TURN_OFF_HTTPS_ESDK_WEBSOCKETS_FIRST = 1126 //删除失败，请先关闭HTTPS和网络服务中的增强型SDK服务及WebSockets服务。
	NET_DVR_ERR_DELETING_FAILED_TURN_OFF_HTTPS_ESDK_FIRST            = 1127 //删除失败，请先关闭HTTPS和网络服务中的增强型SDK服务
	NET_DVR_ERR_PTZ_OCCUPIED_PRIORITY                                = 1128 // 有高优先级云台控制权限用户操作
	NET_DVR_ERR_INCORRECT_VIDEOAUDIO_ID                              = 1129 // 视频通道编码ID或语音输出通道编码ID错误
	NET_DVR_ERR_REPETITIONTIME_OVER_MAXIMUM                          = 1130 // 去重时间最大不超过最大值
	NET_DVR_ERR_FORMATTING_FAILED                                    = 1131 // 格式化错误，请重新
	NET_DVR_ERR_ENCRYPTED_FORMATTING_FAILED                          = 1132 // 加密格式化失败，请重试
	NET_DVR_ERR_WRONG_PASSWORD                                       = 1133 // 密码错误,请输入正确的密码（SD卡 密码校验失败）
	NET_DVR_ERR_EXPOSURE_SYNC                                        = 1134 // 镜头间曝光同步已开启，不允许配置手动RGB

	//2012-10-16 报警设备错误码（1200~1300）
	NET_ERR_SEARCHING_MODULE                  = 1201 // 正在搜索外接模块
	NET_ERR_REGISTERING_MODULE                = 1202 // 正在注册外接模块
	NET_ERR_GETTING_ZONES                     = 1203 // 正在获取防区参数
	NET_ERR_GETTING_TRIGGERS                  = 1204 // 正在获取触发器
	NET_ERR_ARMED_STATUS                      = 1205 // 系统处于布防状态
	NET_ERR_PROGRAM_MODE_STATUS               = 1206 // 系统处于编程模式
	NET_ERR_WALK_TEST_MODE_STATUS             = 1207 // 系统处于步测模式
	NET_ERR_BYPASS_STATUS                     = 1208 // 旁路状态
	NET_ERR_DISABLED_MODULE_STATUS            = 1209 // 功能未使能
	NET_ERR_NOT_SUPPORT_OPERATE_ZONE          = 1210 // 防区不支持该操作
	NET_ERR_NOT_SUPPORT_MOD_MODULE_ADDR       = 1211 // 模块地址不能被修改
	NET_ERR_UNREGISTERED_MODULE               = 1212 // 模块未注册
	NET_ERR_PUBLIC_SUBSYSTEM_ASSOCIATE_SELF   = 1213 // 公共子系统关联自身
	NET_ERR_EXCEEDS_ASSOCIATE_SUBSYSTEM_NUM   = 1214 // 超过公共子系统最大关联个数
	NET_ERR_BE_ASSOCIATED_BY_PUBLIC_SUBSYSTEM = 1215 // 子系统被其他公共子系统关联
	NET_ERR_ZONE_FAULT_STATUS                 = 1216 // 防区处于故障状态
	NET_ERR_SAME_EVENT_TYPE                   = 1217 // 事件触发报警输出开启和事件触发报警输出关闭中有相同事件类型
	NET_ERR_ZONE_ALARM_STATUS                 = 1218 // 防区处于报警状态
	NET_ERR_EXPANSION_BUS_SHORT_CIRCUIT       = 1219 //扩展总线短路
	NET_ERR_PWD_CONFLICT                      = 1220 //密码冲突
	NET_ERR_DETECTOR_GISTERED_BY_OTHER_ZONE   = 1221 //探测器已被其他防区注册
	NET_ERR_DETECTOR_GISTERED_BY_OTHER_PU     = 1222 //探测器已被其他主机注册
	NET_ERR_DETECTOR_DISCONNECT               = 1223 //探测器不在线
	NET_ERR_CALL_BUSY                         = 1224 //设备正在通话中
	NET_DVR_ERR_ZONE_TAMPER_STAUS             = 1225 //防区处于防拆状态
	NET_DVR_ERR_WIRELESS_DEV_REGISTER         = 1226 //无线外设已被其他主机注册
	NET_DVR_ERR_WIRELESS_DEV_ADDED            = 1227 //无线外设已被添加
	NET_DVR_ERR_WIRELESS_DEV_OFFLINE          = 1228 //无线外设不在线
	NET_DVR_ERR_WIRELESS_DEV_TAMPER_STATUS    = 1229 //无线外设处于防拆状态
	NET_DVR_ERR_GPRS_PHONE_CONFLICT           = 1230 //电话报警与无线报警中心冲突
	//信息发布主机
	NET_ERR_GET_ALL_RETURN_OVER = 1300 //获取所有返回数目超限
	NET_ERR_RESOURCE_USING      = 1301 //信息发布资源正在使用，不能修改
	NET_ERR_FILE_SIZE_OVERLIMIT = 1302 //文件大小超限

	//信息发布服务器错误码
	NET_ERR_MATERIAL_NAME                                                   = 1303 //素材名称非法
	NET_ERR_MATERIAL_NAME_LEN                                               = 1304 //素材名称长度非法
	NET_ERR_MATERIAL_REMARK                                                 = 1305 //素材描述非法
	NET_ERR_MATERIAL_REMARK_LEN                                             = 1306 //素材描述长度非法
	NET_ERR_MATERIAL_SHARE_PROPERTY                                         = 1307 //素材共享属性非法
	NET_ERR_UNSUPPORT_MATERIAL_TYPE                                         = 1308 //素材类型不支持
	NET_ERR_MATERIAL_NOT_EXIST                                              = 1309 //素材不存在
	NET_ERR_READ_FROM_DISK                                                  = 1310 //从硬盘读取素材文件失败
	NET_ERR_WRITE_TO_DISK                                                   = 1311 //向硬盘写素材文件失败
	NET_ERR_WRITE_DATA_BASE                                                 = 1312 //素材写数据库失败
	NET_ERR_NO_APPROVED_NOT_EXPORT                                          = 1313 //未审核内容禁止发布
	NET_ERR_MATERIAL_EXCEPTION                                              = 1314 //素材异常
	NET_ERR_NO_MISINFO                                                      = 1315 //无误报信息
	NET_ERR_LAN_NOT_SUP_DHCP_CLIENT_CONFIGURATION                           = 1316 //网桥在路由模式下，配置DHCP客户端返回错误
	NET_ERR_VIDEOWALL_OPTPORT_RESOLUTION_INCONSISTENT                       = 1317 //电视墙上各输出口分辨率不一致(主要用于设置输出分辨率为4K出现异常时报错)
	NET_ERR_VIDEOWALL_OPTPORT_RESOLUTION_INCONSISTENT_UNBIND_OPTPORT_FIRST  = 1318 //电视墙上各输出口分辨率不一致，请先解绑定输出口(主要用于绑定输出口出现异常时报错)
	NET_ERR_FOUR_K_OUTPUT_RESOLUTION_UNSUPPORT_NINE_TO_SIXTEEN_SPLIT_SCREEN = 1319 //4K输出分辨率不支持9/16分屏
	NET_ERR_SIGNAL_SOURCE_UNSUPPORT_CUSTOM_RESOLUTION                       = 1320 //信号源不支持该自定义分辨率
	NET_ERR_DVI_UNSUPPORT_FOURK_OUTPUT_RESOLUTION                           = 1321 //DVI不支持4K输出分辨率
	NET_ERR_BNC_UNSUPPORT_SOURCE_CROPPING                                   = 1322 //BNC不支持信号源裁剪

	//多屏互动错误码
	NET_ERR_MAX_SCREEN_CTRL_NUM   = 1351 //屏幕控制连接数达到上限
	NET_ERR_FILE_NOT_EXIST        = 1352 //文件不存在
	NET_ERR_THUMBNAIL_NOT_EXIST   = 1353 //缩略图不存在
	NET_ERR_DEV_OPEN_FILE_FAIL    = 1354 //设备端打开文件失败
	NET_ERR_SERVER_READ_FILE_FAIL = 1355 //设备端读取文件失败
	NET_ERR_FILE_SIZE             = 1356 //文件大小错误
	NET_ERR_FILE_NAME             = 1357 //文件名称错误，为空或不合法

	//分段错误码（1351-1400）
	NET_ERR_BROADCAST_BUSY = 1358 //设备正在广播中

	//2012-12-20抓拍机错误码（1400-1499）
	NET_DVR_ERR_LANENUM_EXCEED           = 1400 //车道数超出能力
	NET_DVR_ERR_PRAREA_EXCEED            = 1401 //牌识区域过大
	NET_DVR_ERR_LIGHT_PARAM              = 1402 //信号灯接入参数错误
	NET_DVR_ERR_LANE_LINE_INVALID        = 1403 //车道线配置错误
	NET_DVR_ERR_STOP_LINE_INVALID        = 1404 //停止线配置错误
	NET_DVR_ERR_LEFTORRIGHT_LINE_INVALID = 1405 //左/右转分界线配置错误
	NET_DVR_ERR_LANE_NO_REPEAT           = 1406 //叠加车道号重复
	NET_DVR_ERR_PRAREA_INVALID           = 1407 //牌识多边形不符合要求
	NET_DVR_ERR_LIGHT_NUM_EXCEED         = 1408 //视频检测交通灯信号灯数目超出最大值
	NET_DVR_ERR_SUBLIGHT_NUM_INVALID     = 1409 //视频检测交通灯信号灯子灯数目不合法
	NET_DVR_ERR_LIGHT_AREASIZE_INVALID   = 1410 //视频检测交通灯输入信号灯框大小不合法
	NET_DVR_ERR_LIGHT_COLOR_INVALID      = 1411 //视频检测交通灯输入信号灯颜色不合法
	NET_DVR_ERR_LIGHT_DIRECTION_INVALID  = 1412 //视频检测交通灯输入灯方向属性不合法
	NET_DVR_ERR_LACK_IOABLITY            = 1413 //IO口实际支持的能力不足

	NET_DVR_ERR_FTP_PORT                          = 1414 //FTP端口号非法（端口号重复或者异常）
	NET_DVR_ERR_FTP_CATALOGUE                     = 1415 //FTP目录名非法（启用多级目录，多级目录传值为空）
	NET_DVR_ERR_FTP_UPLOAD_TYPE                   = 1416 //FTP上传类型非法（单ftp只支持全部/双ftp只支持卡口和违章）
	NET_DVR_ERR_FLASH_PARAM_WRITE                 = 1417 //配置参数时写FLASH失败
	NET_DVR_ERR_FLASH_PARAM_READ                  = 1418 //配置参数时读FLASH失败
	NET_DVR_ERR_PICNAME_DELIMITER                 = 1419 //FTP图片命名分隔符非法
	NET_DVR_ERR_PICNAME_ITEM                      = 1420 //FTP图片命名项非法（例如 分隔符）
	NET_DVR_ERR_PLATE_RECOGNIZE_TYPE              = 1421 //牌识区域类型非法 （矩形和多边形有效性校验）
	NET_DVR_ERR_CAPTURE_TIMES                     = 1422 //抓拍次数非法 （有效值是0～5）
	NET_DVR_ERR_LOOP_DISTANCE                     = 1423 //线圈距离非法 （有效值是0～2000ms）
	NET_DVR_ERR_LOOP_INPUT_STATUS                 = 1424 //线圈输入状态非法 （有效值）
	NET_DVR_ERR_RELATE_IO_CONFLICT                = 1425 //测速组IO关联冲突
	NET_DVR_ERR_INTERVAL_TIME                     = 1426 //连拍间隔时间非法 （0～6000ms）
	NET_DVR_ERR_SIGN_SPEED                        = 1427 //标志限速值非法（大车标志限速不能大于小车标志限速 ）
	NET_DVR_ERR_PIC_FLIP                          = 1428 //图像配置翻转 （配置交互影响）
	NET_DVR_ERR_RELATE_LANE_NUMBER                = 1429 //关联车道数错误 (重复 有效值校验1～99)
	NET_DVR_ERR_TRIGGER_MODE                      = 1430 //配置抓拍机触发模式非法
	NET_DVR_ERR_DELAY_TIME                        = 1431 //触发延时时间错误(2000ms)
	NET_DVR_ERR_EXCEED_RS485_COUNT                = 1432 //超过最大485个数限制
	NET_DVR_ERR_RADAR_TYPE                        = 1433 //雷达类型错误
	NET_DVR_ERR_RADAR_ANGLE                       = 1434 //雷达角度错误
	NET_DVR_ERR_RADAR_SPEED_VALID_TIME            = 1435 //雷达有效时间错误
	NET_DVR_ERR_RADAR_LINE_CORRECT                = 1436 //雷达线性矫正参数错误
	NET_DVR_ERR_RADAR_CONST_CORRECT               = 1437 //雷达常量矫正参数错误
	NET_DVR_ERR_RECORD_PARAM                      = 1438 //录像参数无效（预录时间不超过10s）
	NET_DVR_ERR_LIGHT_WITHOUT_COLOR_AND_DIRECTION = 1439 //视频检测信号灯配置信号灯个数，但是没有勾选信号灯方向和颜色的
	NET_DVR_ERR_LIGHT_WITHOUT_DETECTION_REGION    = 1440 //视频检测信号灯配置信号灯个数，但是没有画检测区域
	NET_DVR_ERR_RECOGNIZE_PROVINCE_PARAM          = 1441 //牌识参数省份参数的合法性

	NET_DVR_ERR_SPEED_TIMEOUT        = 1442 //IO测速超时时间非法（有效值大于0）
	NET_DVR_ERR_NTP_TIMEZONE         = 1443 //ntp时区参数错误
	NET_DVR_ERR_NTP_INTERVAL_TIME    = 1444 //ntp校时间隔错误
	NET_DVR_ERR_NETWORK_CARD_NUM     = 1445 //可配置网卡数目错误
	NET_DVR_ERR_DEFAULT_ROUTE        = 1446 //默认路由错误
	NET_DVR_ERR_BONDING_WORK_MODE    = 1447 //bonding网卡工作模式错误
	NET_DVR_ERR_SLAVE_CARD           = 1448 //slave网卡错误
	NET_DVR_ERR_PRIMARY_CARD         = 1449 //Primary网卡错误
	NET_DVR_ERR_DHCP_PPOE_WORK       = 1450 //dhcp和pppoE不能同时启动
	NET_DVR_ERR_NET_INTERFACE        = 1451 //网络接口错误
	NET_DVR_ERR_MTU                  = 1452 //MTU错误
	NET_DVR_ERR_NETMASK              = 1453 //子网掩码错误
	NET_DVR_ERR_IP_INVALID           = 1454 //IP地址不合法
	NET_DVR_ERR_MULTICAST_IP_INVALID = 1455 //多播地址不合法
	NET_DVR_ERR_GATEWAY_INVALID      = 1456 //网关不合法
	NET_DVR_ERR_DNS_INVALID          = 1457 //DNS不合法
	NET_DVR_ERR_ALARMHOST_IP_INVALID = 1458 //告警主机地址不合法
	NET_DVR_ERR_IP_CONFLICT          = 1459 //IP冲突
	NET_DVR_ERR_NETWORK_SEGMENT      = 1460 //IP不支持同网段
	NET_DVR_ERR_NETPORT              = 1461 //端口错误

	NET_DVR_ERR_PPPOE_NOSUPPORT        = 1462 //PPPOE不支持
	NET_DVR_ERR_DOMAINNAME_NOSUPPORT   = 1463 //域名不支持
	NET_DVR_ERR_NO_SPEED               = 1464 //未启用测速功能
	NET_DVR_ERR_IOSTATUS_INVALID       = 1465 //IO状态错误
	NET_DVR_ERR_BURST_INTERVAL_INVALID = 1466 //连拍间隔非法
	NET_DVR_ERR_RESERVE_MODE           = 1467 //备用模式错误

	NET_DVR_ERR_LANE_NO            = 1468 //叠加车道号错误
	NET_DVR_ERR_COIL_AREA_TYPE     = 1469 //线圈区域类型错误
	NET_DVR_ERR_TRIGGER_AREA_PARAM = 1470 //触发区域参数错误
	NET_DVR_ERR_SPEED_LIMIT_PARAM  = 1471 //违章限速参数错误
	NET_DVR_ERR_LANE_PROTOCOL_TYPE = 1472 //车道关联协议类型错误

	NET_DVR_ERR_INTERVAL_TYPE               = 1473 //连拍间隔类型非法
	NET_DVR_ERR_INTERVAL_DISTANCE           = 1474 //连拍间隔距离非法
	NET_DVR_ERR_RS485_ASSOCIATE_DEVTYPE     = 1475 //RS485关联类型非法
	NET_DVR_ERR_RS485_ASSOCIATE_LANENO      = 1476 //RS485关联车道号非法
	NET_DVR_ERR_LANENO_ASSOCIATE_MULTIRS485 = 1477 //车道号关联多个RS485口
	NET_DVR_ERR_LIGHT_DETECTION_REGION      = 1478 //视频检测信号灯配置信号灯个数，但是检测区域宽或高为0

	NET_DVR_ERR_DN2D_NOSUPPORT     = 1479 //不支持抓拍帧2D降噪
	NET_DVR_ERR_IRISMODE_NOSUPPORT = 1480 //不支持的镜头类型
	NET_DVR_ERR_WB_NOSUPPORT       = 1481 //不支持的白平衡模式
	NET_DVR_ERR_IO_EFFECTIVENESS   = 1482 //IO口的有效性
	NET_DVR_ERR_LIGHTNO_MAX        = 1483 //信号灯检测器接入红/黄灯超限(16)
	NET_DVR_ERR_LIGHTNO_CONFLICT   = 1484 //信号灯检测器接入红/黄灯冲突

	NET_DVR_ERR_CANCEL_LINE        = 1485 //直行触发线
	NET_DVR_ERR_STOP_LINE          = 1486 //待行区停止线
	NET_DVR_ERR_RUSH_REDLIGHT_LINE = 1487 //闯红灯触发线
	NET_DVR_ERR_IOOUTNO_MAX        = 1488 //IO输出口编号越界

	NET_DVR_ERR_IOOUTNO_AHEADTIME_MAX             = 1489 //IO输出口提前时间超限
	NET_DVR_ERR_IOOUTNO_IOWORKTIME                = 1490 //IO输出口有效持续时间超限
	NET_DVR_ERR_IOOUTNO_FREQMULTI                 = 1491 //IO输出口脉冲模式下倍频出错
	NET_DVR_ERR_IOOUTNO_DUTYRATE                  = 1492 //IO输出口脉冲模式下占空比出错
	NET_DVR_ERR_VIDEO_WITH_EXPOSURE               = 1493 //以曝闪起效，工作方式不支持视频
	NET_DVR_ERR_PLATE_BRIGHTNESS_WITHOUT_FLASHDET = 1494 //车牌亮度自动使能闪光灯仅在车牌亮度补偿模式下起效

	NET_DVR_ERR_RECOGNIZE_TYPE_PARAM       = 1495 //识别类型非法 车牌识别参数（如大车、小车、背向、正向、车标识别等）
	NET_DVR_ERR_PALTE_RECOGNIZE_AREA_PARAM = 1496 //牌识参数非法 牌识区域配置时判断出错
	NET_DVR_ERR_PORT_CONFLICT              = 1497 //端口有冲突
	NET_DVR_ERR_LOOP_IP                    = 1498 //IP不能设置为回环地址
	NET_DVR_ERR_DRIVELINE_SENSITIVE        = 1499 //压线灵敏度出错(视频电警模式下)

	//2013-3-6VQD错误码（1500～1550）
	NET_ERR_VQD_TIME_CONFLICT = 1500 //VQD诊断时间段冲突
	NET_ERR_VQD_PLAN_NO_EXIST = 1501 //VQD诊断计划不存在
	NET_ERR_VQD_CHAN_NO_EXIST = 1502 //VQD监控点不存在
	NET_ERR_VQD_CHAN_MAX      = 1503 //VQD计划数已达上限
	NET_ERR_VQD_TASK_MAX      = 1504 //VQD任务数已达上限

	NET_SDK_GET_INPUTSTREAMCFG          = 1551 //获取信号源
	NET_SDK_AUDIO_SWITCH_CONTROL        = 1552 //子窗口音频开关控制
	NET_SDK_GET_VIDEOWALLDISPLAYNO      = 1553 //获取设备显示输出号
	NET_SDK_GET_ALLSUBSYSTEM_BASIC_INFO = 1554 //获取所有子系统基本信息
	NET_SDK_SET_ALLSUBSYSTEM_BASIC_INFO = 1555 //设置所有子系统基本信息
	NET_SDK_GET_AUDIO_INFO              = 1556 //获取所有音频信息
	NET_SDK_GET_MATRIX_STATUS_V50       = 1557 // 获取视频综合平台状态_V50
	NET_SDK_DELETE_MONITOR_INFO         = 1558 //删除Monitor信息
	NET_SDK_DELETE_CAMERA_INFO          = 1559 //删除Camaera信息

	//抓拍机错误码新增扩展(1600~1900)
	NET_DVR_ERR_EXCEED_MAX_CAPTURE_TIMES = 1600 //抓拍模式为频闪时最大抓拍张数为2张(IVT模式下)
	NET_DVR_ERR_REDAR_TYPE_CONFLICT      = 1601 //相同485口关联雷达类型冲突
	NET_DVR_ERR_LICENSE_PLATE_NULL       = 1602 //车牌号为空
	NET_DVR_ERR_WRITE_DATABASE           = 1603 //写入数据库失败
	NET_DVR_ERR_LICENSE_EFFECTIVE_TIME   = 1604 //车牌有效时间错误
	//视频电警
	NET_DVR_ERR_PRERECORDED_STARTTIME_LONG = 1605 //预录开始时间大于违法抓拍张数
	//混合卡口
	NET_DVR_ERR_TRIGGER_RULE_LINE                 = 1606 //触发规则线错误
	NET_DVR_ERR_LEFTRIGHT_TRIGGERLINE_NOTVERTICAL = 1607 //左/右触发线不垂直
	NET_DVR_ERR_FLASH_LAMP_MODE                   = 1608 //闪光灯闪烁模式错误
	NET_DVR_ERR_ILLEGAL_SNAPSHOT_NUM              = 1609 //违章抓拍张数错误
	NET_DVR_ERR_ILLEGAL_DETECTION_TYPE            = 1610 //违章检测类型错误
	NET_DVR_ERR_POSITIVEBACK_TRIGGERLINE_HIGH     = 1611 //正背向触发线高度错误
	NET_DVR_ERR_MIXEDMODE_CAPTYPE_ALLTARGETS      = 1612 //混合模式下只支持机非人抓拍类型

	NET_DVR_ERR_CARSIGNSPEED_GREATERTHAN_LIMITSPEED                                 = 1613 //小车标志限速大于限速值
	NET_DVR_ERR_BIGCARSIGNSPEED_GREATERTHAN_LIMITSPEED                              = 1614 //大车标志限速大于限速值
	NET_DVR_ERR_BIGCARSIGNSPEED_GREATERTHAN_CARSIGNSPEED                            = 1615 //大车标志限速大于小车标志限速值
	NET_DVR_ERR_BIGCARLIMITSPEED_GREATERTHAN_CARLIMITSPEED                          = 1616 //大车限速值大于小车限速值
	NET_DVR_ERR_BIGCARLOWSPEEDLIMIT_GREATERTHAN_CARLOWSPEEDLIMIT                    = 1617 //大车低速限速值大于小车低速限速值
	NET_DVR_ERR_CARLIMITSPEED_GREATERTHAN_EXCEPHIGHSPEED                            = 1618 //小车限速大于异常高速值
	NET_DVR_ERR_BIGCARLIMITSPEED_GREATERTHAN_EXCEPHIGHSPEED                         = 1619 //大车限速大于异常高速值
	NET_DVR_ERR_STOPLINE_MORETHAN_TRIGGERLINE                                       = 1620 //停止线超过直行触发线
	NET_DVR_ERR_YELLOWLIGHTTIME_INVALID                                             = 1621 /*视频检测黄灯持续时间不合法报错*/
	NET_DVR_ERR_TRIGGERLINE1_FOR_NOT_YIELD_TO_PEDESTRIAN_CANNOT_EXCEED_TRIGGERLINE2 = 1622 //第一条不礼让行人触发线的位置超过了第二条不礼让行人触发线
	NET_DVR_ERR_TRIGGERLINE2_FOR_NOT_YIELD_TO_PEDESTRIAN_CANNOT_EXCEED_TRIGGERLINE1 = 1623 //第二条不礼让行人触发线的位置超过了第一条不礼让行人触发线

	//门禁主机错误码
	NET_ERR_TIME_OVERLAP                            = 1900 //时间段重叠
	NET_ERR_HOLIDAY_PLAN_OVERLAP                    = 1901 //假日计划重叠
	NET_ERR_CARDNO_NOT_SORT                         = 1902 //卡号未排序
	NET_ERR_CARDNO_NOT_EXIST                        = 1903 //卡号不存在
	NET_ERR_ILLEGAL_CARDNO                          = 1904 //卡号错误
	NET_ERR_ZONE_ALARM                              = 1905 //防区处于布防状态(参数修改不允许)
	NET_ERR_ZONE_OPERATION_NOT_SUPPORT              = 1906 //防区不支持该操作
	NET_ERR_INTERLOCK_ANTI_CONFLICT                 = 1907 //多门互锁和反潜回同时配置错误
	NET_ERR_DEVICE_CARD_FULL                        = 1908 //卡已满（卡达到10W后返回）
	NET_ERR_HOLIDAY_GROUP_DOWNLOAD                  = 1909 //假日组下载失败
	NET_ERR_LOCAL_CONTROL_OFF                       = 1910 //就地控制器离线
	NET_ERR_LOCAL_CONTROL_DISADD                    = 1911 //就地控制器未添加
	NET_ERR_LOCAL_CONTROL_HASADD                    = 1912 //就地控制器已添加
	NET_ERR_LOCAL_CONTROL_DOORNO_CONFLICT           = 1913 //与已添加的就地控制器门编号冲突
	NET_ERR_LOCAL_CONTROL_COMMUNICATION_FAIL        = 1914 //就地控制器通信失败
	NET_ERR_OPERAND_INEXISTENCE                     = 1915 //操作对象不存在（对门、报警输出、报警输入相关操作，当对象未添加时返回）
	NET_ERR_LOCAL_CONTROL_OVER_LIMIT                = 1916 //就地控制器超出设备最大能力（主控对就地数量有限制）
	NET_ERR_DOOR_OVER_LIMIT                         = 1917 //门超出设备最大能力
	NET_ERR_ALARM_OVER_LIMIT                        = 1918 //报警输入输出超出设备最大能力
	NET_ERR_LOCAL_CONTROL_ADDRESS_INCONFORMITY_TYPE = 1919 //就地控制器地址与类型不符
	NET_ERR_NOT_SUPPORT_ONE_MORE_CARD               = 1920 //不支持一人多卡
	NET_ERR_DELETE_NO_EXISTENCE_FACE                = 1921 //删除的人脸不存在
	NET_ERR_DOOR_SPECIAL_PASSWORD_REPEAT            = 1922 //与设备门特殊密码重复
	NET_ERR_AUTH_CODE_REPEAT                        = 1923 //与设备认证码重复
	NET_ERR_DEPLOY_EXCEED_MAX                       = 1924 //布防超过最大连接数
	NET_ERR_NOT_SUPPORT_DEL_FP_BY_ID                = 1925 //读卡器不支持按手指ID删除指纹
	NET_ERR_TIME_RANGE                              = 1926 //有效期参数配置范围有误
	NET_ERR_CAPTURE_TIMEOUT                         = 1927 //采集超时
	NET_ERR_LOW_SCORE                               = 1928 //采集质量低
	NET_ERR_OFFLINE_CAPTURING                       = 1929 //离线采集中，无法响应

	//可视对讲错误码
	NET_DVR_ERR_OUTDOOR_COMMUNICATION       = 1950 //与门口机通信异常
	NET_DVR_ERR_ROOMNO_UNDEFINED            = 1951 //未设置房间号
	NET_DVR_ERR_NO_CALLING                  = 1952 //无呼叫
	NET_DVR_ERR_RINGING                     = 1953 //响铃
	NET_DVR_ERR_IS_CALLING_NOW              = 1954 //正在通话
	NET_DVR_ERR_LOCK_PASSWORD_WRONG         = 1955 //智能锁密码错误
	NET_DVR_ERR_CONTROL_LOCK_FAILURE        = 1956 //开关锁失败
	NET_DVR_ERR_CONTROL_LOCK_OVERTIME       = 1957 //开关锁超时
	NET_DVR_ERR_LOCK_DEVICE_BUSY            = 1958 //智能锁设备繁忙
	NET_DVR_ERR_UNOPEN_REMOTE_LOCK_FUNCTION = 1959 //远程开锁功能未打开

	//后端错误码 （2100 - 3000）
	NET_DVR_ERR_FILE_NOT_COMPLETE   = 2100 //下载的文件不完整
	NET_DVR_ERR_IPC_EXIST           = 2101 //该IPC已经存在
	NET_DVR_ERR_ADD_IPC             = 2102 //该通道已添加IPC
	NET_DVR_ERR_OUT_OF_RES          = 2103 //网络带宽能力不足
	NET_DVR_ERR_CONFLICT_TO_LOCALIP = 2104 //IPC的ip地址跟DVR的ip地址冲突
	NET_DVR_ERR_IP_SET              = 2105 //非法ip地址
	NET_DVR_ERR_PORT_SET            = 2106 //非法的端口号

	NET_ERR_WAN_NOTSUPPORT                     = 2107 //不在同一个局域网，无法设置安全问题或导出GUID文件
	NET_ERR_MUTEX_FUNCTION                     = 2108 //功能互斥
	NET_ERR_QUESTION_CONFIGNUM                 = 2109 //安全问题配置数量错误
	NET_ERR_FACECHAN_NORESOURCE                = 2110 //人脸智能通道资源已用完
	NET_ERR_DATA_CALLBACK                      = 2111 //正在数据回迁
	NET_ERR_ATM_VCA_CHAN_IS_RELATED            = 2112 //ATM智能通道已被关联
	NET_ERR_ATM_VCA_CHAN_IS_OVERLAPED          = 2113 //ATM智能通道已被叠加
	NET_ERR_FACE_CHAN_UNOVERLAP_EACH_OTHER     = 2114 //人脸通道不能互相叠加
	NET_ERR_ACHIEVE_MAX_CHANNLE_LIMIT          = 2115 //达到最大路数限制
	NET_DVR_SMD_ENCODING_NORESOURSE            = 2116 //SMD编码资源不足
	NET_DVR_SMD_DECODING_NORESOURSE            = 2117 //SMD解码资源不足
	NET_DVR_FACELIB_DATA_PROCESSING            = 2118 //人脸库数据正在处理
	NET_DVR_ERR_LARGE_TIME_DIFFRENCE           = 2119 //设备和服务器之间的时间差异太大
	NET_DVR_NO_SUPPORT_WITH_PLAYBACK           = 2120 //已开启回放，不支持本功能
	NET_DVR_CHANNEL_NO_SUPPORT_WITH_SMD        = 2121 //通道已开启SMD，不支持本功能
	NET_DVR_CHANNEL_NO_SUPPORT_WITH_FD         = 2122 //通道已开启人脸抓拍，不支持本功能
	NET_DVR_ILLEGAL_PHONE_NUMBER               = 2123 //非法的电话号码
	NET_DVR_ILLEGAL_CERITIFICATE_NUMBER        = 2124 //非法的证件号码
	NET_DVR_ERR_CHANNEL_RESOLUTION_NO_SUPPORT  = 2125 //通道分辨率不支持
	NET_DVR_ERR_CHANNEL_COMPRESSION_NO_SUPPORT = 2126 //通道编码格式不支持

	NET_DVR_ERR_CLUSTER_DEVICE_TOO_LESS              = 2127 //设备数少，不允许删除
	NET_DVR_ERR_CLUSTER_DEL_DEVICE_CM_PLAYLOAD       = 2128 //该设备是集群主机，不允许删除
	NET_DVR_ERR_CLUSTER_DEVNUM_OVER_UPPER_LIMIT      = 2129 //设备数达到上限，不允许增加
	NET_DVR_ERR_CLUSTER_DEVICE_TYPE_INCONFORMITY     = 2130 //设备类型不一致
	NET_DVR_ERR_CLUSTER_DEVICE_VERSION_INCONFORMITY  = 2131 //设备版本不一致
	NET_DVR_ERR_CLUSTER_IP_CONFLICT                  = 2132 //集群系统IP地址冲突：ipv4地址冲突、ipv6地址冲突
	NET_DVR_ERR_CLUSTER_IP_INVALID                   = 2133 //集群系统IP地址无效：ipv4非法、ipv6非法
	NET_DVR_ERR_CLUSTER_PORT_CONFLICT                = 2134 //集群系统端口冲突
	NET_DVR_ERR_CLUSTER_PORT_INVALID                 = 2135 //集群系统端口非法
	NET_DVR_ERR_CLUSTER_USERNAEM_OR_PASSWORD_INVALID = 2136 //组网用户名或密码非法
	NET_DVR_ERR_CLUSTER_DEVICE_ALREADY_EXIST         = 2137 //存在相同设备
	NET_DVR_ERR_CLUSTER_DEVICE_NOT_EXIST             = 2138 //设备不存在(组网时下发的cs列表中的设备信息在任何一台cs上都找不到，删除的时候下发的id不对)
	NET_DVR_ERR_CLUSTER_NON_CLUSTER_MODE             = 2139 //设备处于非集群模式
	NET_DVR_ERR_CLUSTER_IP_NOT_SAME_LAN              = 2140 //IP地址不在同一局域网，不同区域网不允许组网/扩容

	NET_DVR_ERR_CAPTURE_PACKAGE_FAILED              = 2141 //抓包失败
	NET_DVR_ERR_CAPTURE_PACKAGE_PROCESSING          = 2142 //正在抓包
	NET_DVR_ERR_SAFETY_HELMET_NO_RESOURCE           = 2143 //安全帽检测资源不足
	NET_DVR_NO_SUPPORT_WITH_ABSTRACT                = 2144 //已开启视频摘要，不支持本功能
	NET_DVR_ERR_TAPE_LIB_NEED_STOP_ARCHIVE          = 2145 //磁带库归档需要停止
	NET_DVR_INSUFFICIENT_DEEP_LEARNING_RESOURCES    = 2146 //深度学习资源超限
	NET_DVR_ERR_IDENTITY_KEY                        = 2147 //交互口令错误
	NET_DVR_MISSING_IDENTITY_KEY                    = 2148 //交互口令缺失
	NET_DVR_NO_SUPPORT_WITH_PERSON_DENSITY_DETECT   = 2149 //已开启人员密度检测，不支持本功能
	NET_DVR_IPC_RESOLUTION_OVERFLOW                 = 2150 //IPC分辨率超限
	NET_DVR_IPC_BITRATE_OVERFLOW                    = 2151 //IPC码率超限
	NET_DVR_ERR_INVALID_TASKID                      = 2152 //无效的taskID
	NET_DVR_PANEL_MODE_NOT_CONFIG                   = 2153 //没有配置面板路智能
	NET_DVR_NO_HUMAN_ENGINES_RESOURCE               = 2154 //人体引擎资源不足
	NET_DVR_ERR_TASK_NUMBER_OVERFLOW                = 2155 //任务数据超过上限
	NET_DVR_ERR_COLLISION_TIME_OVERFLOW             = 2156 //碰撞时间超过上限
	NET_DVR_ERR_CAPTURE_PACKAGE_NO_USB              = 2157 //未识别到U盘，请插入U盘或重新插入
	NET_DVR_ERR_NO_SET_SECURITY_EMAIL               = 2158 //未设置安全邮箱
	NET_DVR_ERR_EVENT_NOTSUPPORT                    = 2159 //订阅事件不支持
	NET_DVR_ERR_PASSWORD_FORMAT                     = 2160 //密码格式不对
	NET_DVR_ACCESS_FRONT_DEVICE_PARAM_FAILURE       = 2161 //获取前端设备参数失败
	NET_DVR_ACCESS_FRONT_DEVICE_STREAM_FAILURE      = 2162 //对前端设备取流失败
	NET_DVR_ERR_USERNAME_FORMAT                     = 2163 //用户名格式不对
	NET_DVR_ERR_UNOPENED_HIGH_RESOLUTION_MODE       = 2164 //超高分辨率模式未开启
	NET_DVR_ERR_TOO_SMALL_QUATO                     = 2165 //配额设置太小
	NET_DVR_ERR_EMAIL_FORMAT                        = 2166 //邮箱格式不对
	NET_DVR_ERR_SECURITY_CODE_FORMAT                = 2167 //安全码格式不对
	NET_DVR_PD_SPACE_TOO_SMALL                      = 2168 //阵列硬盘容量太小
	NET_DVR_PD_NUM_TOO_BIG                          = 2169 //阵列硬盘总数超过总盘数的二分之一
	NET_DVR_ERR_USB_IS_FULL                         = 2170 //U盘已满
	NET_DVR_EXCEED_MAX_SMD_TYPE                     = 2171 //达到最大SMD事件种类上限
	NET_DVR_CHANNEL_NO_SUPPORT_WITH_BEHAVIOR        = 2172 //通道已开启行为分析，不支持本功能
	NET_DVR_NO_BEHAVIOR_ENGINES_RESOURCE            = 2173 //行为分析资源不足
	NET_DVR_NO_RETENTION_ENGINES_RESOURCE           = 2174 //人员滞留检测资源不足
	NET_DVR_NO_LEAVE_POSITION_ENGINES_RESOURCE      = 2175 //离岗检测资源不足
	NET_DVR_NO_PEOPLE_NUM_CHANGE_ENGINES_RESOURCE   = 2176 //人数异常资源不足
	NET_DVR_PANEL_MODE_NUM_OVER_LIMIT               = 2177 //超过面板路最大路数
	NET_DVR_SURROUND_MODE_NUM_OVER_LIMIT            = 2178 //超过环境路最大路数
	NET_DVR_FACE_MODE_NUM_OVER_LIMIT                = 2179 //超过人脸路最大路数
	NET_DVR_SAFETYCABIN_MODE_NUM_OVER_LIMIT         = 2180 //超过防护舱路最大路数
	NET_DVR_DETECT_REGION_RANGE_INVALID             = 2181 //检测区域范围非法
	NET_DVR_CHANNEL_CAPTURE_PICTURE_FAILURE         = 2182 //通道抓图失败
	NET_DVR_VCACHAN_IS_NORESOURCE                   = 2183 //智能通道资源用完
	NET_DVR_IPC_NUM_REACHES_LIMIT                   = 2184 // Ipc通道数目达到上限
	NET_DVR_IOT_NUM_REACHES_LIMIT                   = 2185 // IOT通道数目达到上限
	NET_DVR_IOT_CHANNEL_DEVICE_EXIST                = 2186 //当前IOT通道已经添加设备
	NET_DVR_IOT_CHANNEL_DEVICE_NOT_EXIST            = 2187 //当前IOT通道不存在设备
	NET_DVR_INVALID_IOT_PROTOCOL_TYPE               = 2188 //非法的IOT协议类型
	NET_DVR_INVALID_EZVIZ_SECRET_KEY                = 2189 //非法的萤石注册验证码
	NET_DVR_DUPLICATE_IOT_DEVICE                    = 2190 //重复的IOT设备
	NET_DVR_SADP_MODIFY_FALIURE                     = 2191 // SADP修改失败
	NET_DVR_IPC_NETWORK_ABNORMAL                    = 2192 // IPC网络异常
	NET_DVR_IPC_PASSWORD_ERROR                      = 2193 // IPC用户名密码错误
	NET_DVR_ERROR_IPC_TYPE                          = 2194 //IPC类型不对
	NET_DVR_ERROR_IPC_LIST_NOT_EMPTY                = 2195 //已添加IPC列表不为空，不支持一键配置
	NET_DVR_ERROR_IPC_LIST_NOT_MATCH_PAIRING        = 2196 //IPC列表和配单不匹配
	NET_DVR_ERROR_IPC_BAD_LANGUAGE                  = 2197 //IPC语言和设备不匹配
	NET_DVR_ERROR_IPC_IS_LOCKING                    = 2198 //IPC已被锁
	NET_DVR_ERROR_IPC_NOT_ACTIVATED                 = 2199 //IPC未激活
	NET_DVR_FIELD_CODING_NOT_SUPPORT                = 2200 //场编码不支持
	NET_DVR_ERROR_H323_NOT_SUPPORT_H265             = 2201 //H323视频会议就不支持H265码流
	NET_DVR_ERROR_EXPOSURE_TIME_TOO_BIG_IN_MODE_P   = 2202 //P制式下，曝光时间过大
	NET_DVR_ERROR_EXPOSURE_TIME_TOO_BIG_IN_MODE_N   = 2203 //N制式下，曝光时间过大
	NET_DVR_ERROR_PING_PROCESSING                   = 2204 //正在PING
	NET_DVR_ERROR_PING_NOT_START                    = 2205 //Ping功能未开始
	NET_DVR_ERROR_NEED_DOUBLE_VERIFICATION          = 2206 //需要二次认证
	NET_DVR_NO_DOUBLE_VERIFICATION_USER             = 2207 //无二次认证用户
	NET_DVR_CHANNEL_OFFLINE                         = 2208 //通道离线
	NET_DVR_TIMESPAN_NUM_OVER_LIMIT                 = 2209 //时间段超出支持最大数目
	NET_DVR_CHANNEL_NUM_OVER_LIMIT                  = 2210 //通道数目超出支持最大数目
	NET_DVR_NO_SEARCH_ID_RESOURCE                   = 2211 //分页查询的searchID资源不足
	NET_DVR_ERROR_ONEKEY_EXPORT                     = 2212 //正在进行导出操作，请稍后再试
	NET_DVR_NO_CITY_MANAGEMENT_ENGINES_RESOURCE     = 2213 //城管算法引擎资源不足
	NET_DVR_NO_SITUATION_ANALYSIS_ENGINES_RESOURCE  = 2214 //态势分析引擎资源不足
	NET_DVR_INTELLIGENT_ANALYSIS_IPC_CANNT_DELETE   = 2215 //正在进行智能分析的IPC无法删除
	NET_DVR_NOSUPPORT_RESET_PASSWORD                = 2216 //NVR不支持对IPC重置密码
	NET_DVR_ERROR_IPC_NEED_ON_LAN                   = 2217 // IPC需要在局域网内
	NET_DVR_CHANNEL_NO_SUPPORT_WITH_SAFETY_HELMET   = 2218 //通道已开启安全帽检测，不支持本功能
	NET_DVR_ERROR_GET_RESETPASSWORDTYPE_IS_ABNORMAL = 2219 /*IPC重置密码时，获取IPC的重置密码类型异常*/
	NET_DVR_ERROR_IPC_NOSUPPORT_RESET_PASSWORD      = 2220 /* IPC不支持重置密码*/
	NET_DVR_ERROR_IP_IS_NOT_ONLY_ONE                = 2221 /*IPC的IP不唯一，有重复*/
	NET_DVR_NO_SUPPORT_WITH_SMD_OR_SCD              = 2222 //已开启SMD/SCD，不支持本功能（SCD为场景变更）
	NET_DVR_NO_SUPPORT_WITH_FD                      = 2223 //已开启人脸抓拍，不支持本功能
	NET_DVR_NO_FD_ENGINES_RESOURCE                  = 2224 //人脸抓拍资源不足
	NET_DVR_ERROR_ONEKEY_REMOVE                     = 2225 //正在进行删除操作，请稍后再试
	NET_DVR_FACE_PIP_BACKGROUND_CHANNEL_OVERFLOW    = 2226 //人脸画中画背景通道超限
	NET_DVR_MICIN_CHANNEL_OCCUPIED                  = 2227 //micin通道被占用
	NET_DVR_IPC_CHANNEL_IS_IN_PIP                   = 2228 //操作失败，该通道已关联到审讯通道，请先取消画中画配置关联

	NET_DVR_CHANNEL_NO_SUPPORT_WITH_FACE_CONTRAST = 2229 //通道已开启人脸比对，不支持本功能

	NET_DVR_INVALID_RECHARGE_CARD                     = 2230 //无效的充值卡
	NET_DVR_CLOUD_PLATFORM_SERVER_EXCEPTION           = 2231 //云平台服务器异常
	NET_DVR_OPERATION_FAILURE_WITHOUT_LOGIN           = 2232 //未登录操作失败
	NET_DVR_INVALID_ASSOCIATED_SERIAL_NUMBER          = 2233 //关联序列号非法
	NET_DVR_CLOUD_PLATFORM_ACCOUNT_NOT_EXIST          = 2234 //云平台帐号不存在
	NET_DVR_DEVICE_SERIAL_NUMBER_REGISTERED           = 2235 //设备序列号已注册
	NET_DVR_CONFERENCE_ROOM_NOT_EXIST                 = 2236 //会议室不存在
	NET_DVR_NEED_DISABLED_ANALOG_CHANNEL              = 2237 //需禁用模拟通道
	NET_DVR_STUDENT_ROLL_CALL_FAILURE                 = 2238 //学生点名失败
	NET_DVR_SUB_DEVICE_NOT_ENABLE_INDIVIDUAL_BEHAVIOR = 2239 //子设备未启用个体行为模式
	NET_DVR_SUB_DEVICE_CHANNEL_CONTROL_FAILED         = 2240 //子设备通道控制失败
	NET_DVR_DEVICE_NOT_IN_CONFERENCE                  = 2241 //设备不在会议中
	NET_DVR_ALREADY_EXIST_CONFERENCE                  = 2242 //当前已经存在会议
	NET_DVR_NO_SUPPORT_WITH_VIDEO_CONFERENCE          = 2243 //当前正在视频会议中，不支持本功能
	NET_DVR_START_INTERACTION_FAILURE                 = 2244 //互动开始失败
	NET_DVR_ASK_QUESTION_STARTED                      = 2245 //已开始提问
	NET_DVR_ASK_QUESTION_CLOSED                       = 2246 //已结束提问
	NET_DVR_UNABLE_OPERATE_BY_HOST                    = 2247 //已被主持人禁用，无法操作
	NET_DVR_REPEATED_ASK_QUESTION                     = 2248 //重复提问
	NET_DVR_SWITCH_TIMEDIFF_LESS_LIMIT                = 2249 /*开关机时间差小于限制值(10分钟)*/
	NET_DVR_CHANNEL_DEVICE_EXIST                      = 2250 //当前通道已经添加设备
	NET_DVR_CHANNEL_DEVICE_NOT_EXIST                  = 2251 //当前通道不存在设备
	NET_DVR_ERROR_ADJUSTING_RESOLUTION                = 2252 //先关闭摄像机的裁剪，再调整分辨率
	NET_DVR_SSD_FILE_SYSTEM_IS_UPGRADING              = 2253 //SSD文件系统正在升级
	NET_DVR_SSD_FILE_SYSTEM_IS_FORMAT                 = 2254 //SSD正在格式化
	NET_DVR_CHANNEL_IS_CONNECTING                     = 2255 //当前通道正在连接

	NET_DVR_CHANNEL_STREAM_TYPE_NOT_SUPPORT = 2257 //当前通道码流类型不支持
	NET_DVR_CHANNEL_USERNAME_NOT_EXIST      = 2258 //当前通道用户名不存在
	NET_DVR_CHANNEL_ACCESS_PARAM_FAILURE    = 2259 //当前通道获取参数失败
	NET_DVR_CHANNEL_GET_STREAM_FAILURE      = 2260 //当前通道取流失败
	NET_DVR_CHANNEL_RISK_PASSWORD           = 2261 //当前通道密码为风险密码

	NET_DVR_NO_SUPPORT_DELETE_STRANGER_LIB = 2262 //不支持删除陌生人库
	NET_DVR_NO_SUPPORT_CREATE_STRANGER_LIB = 2263 //不支持创建陌生人库

	NET_DVR_NETWORK_PORT_CONFLICT         = 2264 //网络配置端口冲突
	NET_DVR_TRANSCODE_NO_RESOURCES        = 2265 //转码资源不足
	NET_DVR_SSD_FILE_SYSTEM_ERROR         = 2266 //SSD文件系统错误
	NET_DVR_INSUFFICIENT_SSD__FOR_FPD     = 2267 //用于人员频次业务的SSD容量不够
	NET_DVR_ASSOCIATED_FACELIB_OVER_LIMIT = 2268 //关联人脸库已达上限
	NET_DVR_NEED_DELETE_DIGITAL_CHANNEL   = 2269 //需删除数字通道

	//热成像产线相关错误码（3001 - 3500）
	NET_DVR_ERR_NOTSUPPORT_DEICING             = 3001 //设备当前状态不支持除冰功能（只在POE+、AC24V、DC12V供电下支持除冰功能）
	NET_DVR_ERR_THERMENABLE_CLOSE              = 3002 //测温功能总使能未开启。(即NET_DVR_THERMOMETRY_BASICPARAM使能未开启)
	NET_DVR_ERR_NOTMEET_DEICING                = 3003 //当前空腔温度不满足手动除冰开启条件（需空腔温度小于30度才可开启）
	NET_DVR_ERR_PANORAMIC_LIMIT_OPERATED       = 3004 //全景地图和限位不可同时操作
	NET_DVR_ERR_SMARTH264_ROI_OPERATED         = 3005 //SmartH264和ROI不可同时操作
	NET_DVR_ERR_RULENUM_LIMIT                  = 3006 //规则数上限
	NET_DVR_ERR_LASER_DEICING_OPERATED         = 3007 //激光以及除冰功能不可同时操作
	NET_DVR_ERR_OFFDIGITALZOOM_OR_MINZOOMLIMIT = 3008 //请先关闭数据变倍功能或变倍限制设置为最小值。当执行烟火检测、行为分析、船只检测、坏点矫正、测温、烟火屏蔽功能时，若没有关闭数据变倍或者变倍限制没有设置为最小值时，将会提示该错误码。
	NET_DVR_ERR_FIREWAITING                    = 3009 //设备云台正在火点等待中
	NET_DVR_SYNCHRONIZEFOV_ERROR               = 3010 //同步视场角错误
	NET_DVR_CERTIFICATE_VALIDATION_ERROR       = 3011 //证书验证不通过
	NET_DVR_CERTIFICATES_NUM_EXCEED_ERROR      = 3012 //证书个数超过上限
	NET_DVR_RULE_SHIELDMASK_CONFLICT_ERROR     = 3013 //规则区域与屏蔽区域冲突

	//前端产品线错误码（3501-4000）
	NET_DVR_ERR_NO_SAFETY_HELMET_REGION = 3501 //未配置安全帽检测区域
	NET_DVR_ERR_UNCLOSED_SAFETY_HELMET  = 3502 //未关闭安全帽检测使能

	NET_ERR_NPQ_BASE_INDEX   = 8000                         //NPQ库错误码
	NET_ERR_NPQ_PARAM        = (NET_ERR_NPQ_BASE_INDEX + 1) //NPQ库参数有误
	NET_ERR_NPQ_SYSTEM       = (NET_ERR_NPQ_BASE_INDEX + 2) //NPQ库操作系统调用错误(包括资源申请失败或内部错误等）
	NET_ERR_NPQ_GENRAL       = (NET_ERR_NPQ_BASE_INDEX + 3) //NPQ库内部通用错误
	NET_ERR_NPQ_PRECONDITION = (NET_ERR_NPQ_BASE_INDEX + 4) //NPQ库调用顺序错误
	NET_ERR_NPQ_NOTSUPPORT   = (NET_ERR_NPQ_BASE_INDEX + 5) //NPQ库功能不支持

	NET_ERR_NPQ_NOTCALLBACK = (NET_ERR_NPQ_BASE_INDEX + 100) //数据没有回调上来
	NET_ERR_NPQ_LOADLIB     = (NET_ERR_NPQ_BASE_INDEX + 101) //NPQ库加载失败
	NET_ERR_NPQ_STEAM_CLOSE = (NET_ERR_NPQ_BASE_INDEX + 104) //本路码流NPQ功能未开启
	NET_ERR_NPQ_MAX_LINK    = (NET_ERR_NPQ_BASE_INDEX + 110) //NPQ取流路数达到上限
	NET_ERR_NPQ_STREAM_CFG  = (NET_ERR_NPQ_BASE_INDEX + 111) //编码参数存在冲突配置

	//传显错误码 8501~9500
	NET_ERR_UPGRADE_PROG_ERR        = 8501 //程序执行出错
	NET_ERR_UPGRADE_NO_DEVICE       = 8502 //没有设备(指LED控制器没有接接收卡)
	NET_ERR_UPGRADE_NO_FILE         = 8503 //没有找到升级文件
	NET_ERR_UPGRADE_DATA_ERROR      = 8504 //升级文件数据不兼容
	NET_ERR_UPGRADE_LINK_SERVER_ERR = 8505 //与服务器连接失败
	NET_ERR_UPGRADE_OEMCODE_NOMATCH = 8506 //oemCode不匹配
	NET_ERR_UPGRADE_FLASH_NOENOUGH  = 8507 //flash不足
	NET_ERR_UPGRADE_RAM_NOENOUGH    = 8508 //RAM不足
	NET_ERR_UPGRADE_DSPRAM_NOENOUGH = 8509 //DSP RAM不足
	NET_ERR_NOT_SUPPORT_CHECK       = 8510 //该屏幕型号不支持校正
	NET_ERR_LED_DEVICE_BUSY_CHECK   = 8511 //LED设备忙（正在校正）
	NET_ERR_DEVICE_MEM_NOT_ENOUGH   = 8512 //设备内存不足
	NET_ERR_CHECK_PARAM             = 8513 //校正参数错误
	NET_ERR_RESOLUTION_OVER_LIMIT   = 8514 //输入分辨率超过限制
	NET_ERR_NO_CUSTOM_BASE          = 8515 //无自定义底图
	NET_ERR_PRIORITY_LOWER          = 8516 //优先级低于当前模式
	NET_ERR_SEND_MESSAGE_EXCEPT     = 8517 //消息发送异常
	NET_ERR_SENDCARD_UPGRADING      = 8518 //发送卡升级中
	NET_ERR_NO_WIRELESS_NETCARD     = 8519 //未插入无线网卡
	NET_ERR_LOAD_FS_FAIL            = 8520 //从屏幕加载失败
	NET_ERR_FLASH_UNSTORAGE_RECCARD = 8521 //Flash中未存储接收卡参数

	/*******************全局错误码 end**********************/

	/*************************************************
	  NET_DVR_IsSupport()返回值
	  1－9位分别表示以下信息（位与是TRUE)表示支持；
	  **************************************************/
	NET_DVR_SUPPORT_DDRAW       = 0x01  //支持DIRECTDRAW，如果不支持，则播放器不能工作；
	NET_DVR_SUPPORT_BLT         = 0x02  //显卡支持BLT操作，如果不支持，则播放器不能工作；
	NET_DVR_SUPPORT_BLTFOURCC   = 0x04  //显卡BLT支持颜色转换，如果不支持，播放器会用软件方法作RGB转换；
	NET_DVR_SUPPORT_BLTSHRINKX  = 0x08  //显卡BLT支持X轴缩小；如果不支持，系统会用软件方法转换；
	NET_DVR_SUPPORT_BLTSHRINKY  = 0x10  //显卡BLT支持Y轴缩小；如果不支持，系统会用软件方法转换；
	NET_DVR_SUPPORT_BLTSTRETCHX = 0x20  //显卡BLT支持X轴放大；如果不支持，系统会用软件方法转换；
	NET_DVR_SUPPORT_BLTSTRETCHY = 0x40  //显卡BLT支持Y轴放大；如果不支持，系统会用软件方法转换；
	NET_DVR_SUPPORT_SSE         = 0x80  //CPU支持SSE指令，Intel Pentium3以上支持SSE指令；
	NET_DVR_SUPPORT_MMX         = 0x100 //CPU支持MMX指令集，Intel Pentium3以上支持SSE指令；

	/**********************云台控制命令 begin*************************/
	LIGHT_PWRON  = 2 /* 接通灯光电源 */
	WIPER_PWRON  = 3 /* 接通雨刷开关 */
	FAN_PWRON    = 4 /* 接通风扇开关 */
	HEATER_PWRON = 5 /* 接通加热器开关 */
	AUX_PWRON1   = 6 /* 接通辅助设备开关 */
	AUX_PWRON2   = 7 /* 接通辅助设备开关 */
	SET_PRESET   = 8 /* 设置预置点 */
	CLE_PRESET   = 9 /* 清除预置点 */

	ZOOM_IN    = 11 /* 焦距以速度SS变大(倍率变大) */
	ZOOM_OUT   = 12 /* 焦距以速度SS变小(倍率变小) */
	FOCUS_NEAR = 13 /* 焦点以速度SS前调 */
	FOCUS_FAR  = 14 /* 焦点以速度SS后调 */
	IRIS_OPEN  = 15 /* 光圈以速度SS扩大 */
	IRIS_CLOSE = 16 /* 光圈以速度SS缩小 */

	TILT_UP    = 21 /* 云台以SS的速度上仰 */
	TILT_DOWN  = 22 /* 云台以SS的速度下俯 */
	PAN_LEFT   = 23 /* 云台以SS的速度左转 */
	PAN_RIGHT  = 24 /* 云台以SS的速度右转 */
	UP_LEFT    = 25 /* 云台以SS的速度上仰和左转 */
	UP_RIGHT   = 26 /* 云台以SS的速度上仰和右转 */
	DOWN_LEFT  = 27 /* 云台以SS的速度下俯和左转 */
	DOWN_RIGHT = 28 /* 云台以SS的速度下俯和右转 */
	PAN_AUTO   = 29 /* 云台以SS的速度左右自动扫描 */

	FILL_PRE_SEQ   = 30 /* 将预置点加入巡航序列 */
	SET_SEQ_DWELL  = 31 /* 设置巡航点停顿时间 */
	SET_SEQ_SPEED  = 32 /* 设置巡航速度 */
	CLE_PRE_SEQ    = 33 /* 将预置点从巡航序列中删除 */
	STA_MEM_CRUISE = 34 /* 开始记录轨迹 */
	STO_MEM_CRUISE = 35 /* 停止记录轨迹 */
	RUN_CRUISE     = 36 /* 开始轨迹 */
	RUN_SEQ        = 37 /* 开始巡航 */
	STOP_SEQ       = 38 /* 停止巡航 */
	GOTO_PRESET    = 39 /* 快球转到预置点 */

	DEL_SEQ           = 43 /* 删除巡航路径 */
	STOP_CRUISE       = 44 /* 停止轨迹 */
	DELETE_CRUISE     = 45 /* 删除单条轨迹 */
	DELETE_ALL_CRUISE = 46 /* 删除所有轨迹 */

	PAN_CIRCLE     = 50 /* 云台以SS的速度自动圆周扫描 */
	DRAG_PTZ       = 51 /* 拖动PTZ */
	LINEAR_SCAN    = 52 /* 区域扫描 */ //2014-03-15
	CLE_ALL_PRESET = 53 /* 预置点全部清除 */
	CLE_ALL_SEQ    = 54 /* 巡航全部清除 */
	CLE_ALL_CRUISE = 55 /* 轨迹全部清除 */

	POPUP_MENU = 56 /* 显示操作菜单 */

	TILT_DOWN_ZOOM_IN   = 58 /* 云台以SS的速度下俯&&焦距以速度SS变大(倍率变大) */
	TILT_DOWN_ZOOM_OUT  = 59 /* 云台以SS的速度下俯&&焦距以速度SS变小(倍率变小) */
	PAN_LEFT_ZOOM_IN    = 60 /* 云台以SS的速度左转&&焦距以速度SS变大(倍率变大)*/
	PAN_LEFT_ZOOM_OUT   = 61 /* 云台以SS的速度左转&&焦距以速度SS变小(倍率变小)*/
	PAN_RIGHT_ZOOM_IN   = 62 /* 云台以SS的速度右转&&焦距以速度SS变大(倍率变大) */
	PAN_RIGHT_ZOOM_OUT  = 63 /* 云台以SS的速度右转&&焦距以速度SS变小(倍率变小) */
	UP_LEFT_ZOOM_IN     = 64 /* 云台以SS的速度上仰和左转&&焦距以速度SS变大(倍率变大)*/
	UP_LEFT_ZOOM_OUT    = 65 /* 云台以SS的速度上仰和左转&&焦距以速度SS变小(倍率变小)*/
	UP_RIGHT_ZOOM_IN    = 66 /* 云台以SS的速度上仰和右转&&焦距以速度SS变大(倍率变大)*/
	UP_RIGHT_ZOOM_OUT   = 67 /* 云台以SS的速度上仰和右转&&焦距以速度SS变小(倍率变小)*/
	DOWN_LEFT_ZOOM_IN   = 68 /* 云台以SS的速度下俯和左转&&焦距以速度SS变大(倍率变大) */
	DOWN_LEFT_ZOOM_OUT  = 69 /* 云台以SS的速度下俯和左转&&焦距以速度SS变小(倍率变小) */
	DOWN_RIGHT_ZOOM_IN  = 70 /* 云台以SS的速度下俯和右转&&焦距以速度SS变大(倍率变大) */
	DOWN_RIGHT_ZOOM_OUT = 71 /* 云台以SS的速度下俯和右转&&焦距以速度SS变小(倍率变小) */
	TILT_UP_ZOOM_IN     = 72 /* 云台以SS的速度上仰&&焦距以速度SS变大(倍率变大) */
	TILT_UP_ZOOM_OUT    = 73 /* 云台以SS的速度上仰&&焦距以速度SS变小(倍率变小) */
	/**********************云台控制命令 end*************************/

	DVR_VEHICLE_CONTROL_LIST = 0x1 //车辆黑白名单数据类型(发送的数据类型)2013-11-04

	/*************************************************
	  回放时播放控制命令宏定义
	      NET_DVR_PlayBackControl
	   = NET_DVR_PlayControlLocDisplay
	  NET_DVR_DecPlayBackCtrl的宏定义
	  具体支持查看函数说明和代码
	  **************************************************/
	NET_DVR_PLAYSTART            = 1  //开始播放
	NET_DVR_PLAYSTOP             = 2  //停止播放
	NET_DVR_PLAYPAUSE            = 3  //暂停播放
	NET_DVR_PLAYRESTART          = 4  //恢复播放
	NET_DVR_PLAYFAST             = 5  //快放
	NET_DVR_PLAYSLOW             = 6  //慢放
	NET_DVR_PLAYNORMAL           = 7  //正常速度
	NET_DVR_PLAYFRAME            = 8  //单帧放
	NET_DVR_PLAYSTARTAUDIO       = 9  //打开声音
	NET_DVR_PLAYSTOPAUDIO        = 10 //关闭声音
	NET_DVR_PLAYAUDIOVOLUME      = 11 //调节音量
	NET_DVR_PLAYSETPOS           = 12 //改变文件回放的进度
	NET_DVR_PLAYGETPOS           = 13 //获取文件回放的进度
	NET_DVR_PLAYGETTIME          = 14 //获取当前已经播放的时间(按文件回放的时候有效)
	NET_DVR_PLAYGETFRAME         = 15 //获取当前已经播放的帧数(按文件回放的时候有效)
	NET_DVR_GETTOTALFRAMES       = 16 //获取当前播放文件总的帧数(按文件回放的时候有效)
	NET_DVR_GETTOTALTIME         = 17 //获取当前播放文件总的时间(按文件回放的时候有效)
	NET_DVR_THROWBFRAME          = 20 //丢B帧
	NET_DVR_SETSPEED             = 24 //设置码流速度
	NET_DVR_KEEPALIVE            = 25 //保持与设备的心跳(如果回调阻塞，建议2秒发送一次)
	NET_DVR_PLAYSETTIME          = 26 //按绝对时间定位
	NET_DVR_PLAYGETTOTALLEN      = 27 //获取按时间回放对应时间段内的所有文件的总长度
	NET_DVR_PLAYSETTIME_V50      = 28 //按绝对时间定位(支持时区扩展)
	NET_DVR_PLAY_FORWARD         = 29 //倒放切换为正放
	NET_DVR_PLAY_REVERSE         = 30 //正放切换为倒放
	NET_DVR_SET_DECODEFFRAMETYPE = 31
	NET_DVR_SET_TRANS_TYPE       = 32 //设置转码格式
	NET_DVR_PLAY_CONVERT         = 33 //回放转码
	NET_DVR_START_DRAWFRAME      = 34 //开始抽帧回放
	NET_DVR_STOP_DRAWFRAME       = 35 //停止抽帧回放
	NET_DVR_CHANGEWNDRESOLUTION  = 36 //窗口大小改变，通知播放库
	NET_DVR_RESETBUFFER          = 37 //清空矩阵解码缓冲区（远程回放文件）
	NET_DVR_VOD_DRAG_ING         = 38 //回放拖动中
	NET_DVR_VOD_DRAG_END         = 39 //回放拖动结束
	NET_DVR_VOD_RESET_PLAYTIME   = 40 //重设播放时间

	PLAYM4_DECODE_NORMAIL          = 0 //正常解码
	PLAYM4_DECODE_KEY_FRAME        = 1 //只解I帧
	PLAYM4_DECODE_NONE             = 2 //全不解
	PLAYM4_DECODE_TEMPORAL_LAYER_0 = 3 //解1/2
	PLAYM4_DECODE_TEMPORAL_LAYER_1 = 4 //解1/4

	//远程按键定义如下：
	/* key value send to CONFIG program */
	KEY_CODE_1      = 1
	KEY_CODE_2      = 2
	KEY_CODE_3      = 3
	KEY_CODE_4      = 4
	KEY_CODE_5      = 5
	KEY_CODE_6      = 6
	KEY_CODE_7      = 7
	KEY_CODE_8      = 8
	KEY_CODE_9      = 9
	KEY_CODE_0      = 10
	KEY_CODE_POWER  = 11
	KEY_CODE_MENU   = 12
	KEY_CODE_ENTER  = 13
	KEY_CODE_CANCEL = 14
	KEY_CODE_UP     = 15
	KEY_CODE_DOWN   = 16
	KEY_CODE_LEFT   = 17
	KEY_CODE_RIGHT  = 18
	KEY_CODE_EDIT   = 19
	KEY_CODE_ADD    = 20
	KEY_CODE_MINUS  = 21
	KEY_CODE_PLAY   = 22
	KEY_CODE_REC    = 23
	KEY_CODE_PAN    = 24
	KEY_CODE_M      = 25
	KEY_CODE_A      = 26
	KEY_CODE_F1     = 27
	KEY_CODE_F2     = 28

	/* for PTZ control */
	KEY_PTZ_UP_START = KEY_CODE_UP
	KEY_PTZ_UP_STOP  = 32

	KEY_PTZ_DOWN_START = KEY_CODE_DOWN
	KEY_PTZ_DOWN_STOP  = 33

	KEY_PTZ_LEFT_START = KEY_CODE_LEFT
	KEY_PTZ_LEFT_STOP  = 34

	KEY_PTZ_RIGHT_START = KEY_CODE_RIGHT
	KEY_PTZ_RIGHT_STOP  = 35

	KEY_PTZ_AP1_START = KEY_CODE_EDIT /*光圈+*/
	KEY_PTZ_AP1_STOP  = 36

	KEY_PTZ_AP2_START = KEY_CODE_PAN /*光圈-*/
	KEY_PTZ_AP2_STOP  = 37

	KEY_PTZ_FOCUS1_START = KEY_CODE_A /*聚焦+*/
	KEY_PTZ_FOCUS1_STOP  = 38

	KEY_PTZ_FOCUS2_START = KEY_CODE_M /*聚焦-*/
	KEY_PTZ_FOCUS2_STOP  = 39

	KEY_PTZ_B1_START = 40 /*变倍+*/
	KEY_PTZ_B1_STOP  = 41

	KEY_PTZ_B2_START = 42 /*变倍-*/
	KEY_PTZ_B2_STOP  = 43

	//9000新增
	KEY_CODE_11 = 44
	KEY_CODE_12 = 45
	KEY_CODE_13 = 46
	KEY_CODE_14 = 47
	KEY_CODE_15 = 48
	KEY_CODE_16 = 49

	/*************************参数配置命令 begin*******************************/
	//用于NET_DVR_SetDVRConfig和NET_DVR_GetDVRConfig,注意其对应的配置结构

	NET_DVR_GET_DEVICECFG    = 100 //获取设备参数
	NET_DVR_SET_DEVICECFG    = 101 //设置设备参数
	NET_DVR_GET_NETCFG       = 102 //获取网络参数
	NET_DVR_SET_NETCFG       = 103 //设置网络参数
	NET_DVR_GET_PICCFG       = 104 //获取图象参数
	NET_DVR_SET_PICCFG       = 105 //设置图象参数
	NET_DVR_GET_COMPRESSCFG  = 106 //获取压缩参数
	NET_DVR_SET_COMPRESSCFG  = 107 //设置压缩参数
	NET_DVR_GET_RECORDCFG    = 108 //获取录像时间参数
	NET_DVR_SET_RECORDCFG    = 109 //设置录像时间参数
	NET_DVR_GET_DECODERCFG   = 110 //获取解码器参数
	NET_DVR_SET_DECODERCFG   = 111 //设置解码器参数
	NET_DVR_GET_RS232CFG     = 112 //获取232串口参数
	NET_DVR_SET_RS232CFG     = 113 //设置232串口参数
	NET_DVR_GET_ALARMINCFG   = 114 //获取报警输入参数
	NET_DVR_SET_ALARMINCFG   = 115 //设置报警输入参数
	NET_DVR_GET_ALARMOUTCFG  = 116 //获取报警输出参数
	NET_DVR_SET_ALARMOUTCFG  = 117 //设置报警输出参数
	NET_DVR_GET_TIMECFG      = 118 //获取DVR时间
	NET_DVR_SET_TIMECFG      = 119 //设置DVR时间
	NET_DVR_GET_PREVIEWCFG   = 120 //获取预览参数
	NET_DVR_SET_PREVIEWCFG   = 121 //设置预览参数
	NET_DVR_GET_VIDEOOUTCFG  = 122 //获取视频输出参数
	NET_DVR_SET_VIDEOOUTCFG  = 123 //设置视频输出参数
	NET_DVR_GET_USERCFG      = 124 //获取用户参数
	NET_DVR_SET_USERCFG      = 125 //设置用户参数
	NET_DVR_GET_EXCEPTIONCFG = 126 //获取异常参数
	NET_DVR_SET_EXCEPTIONCFG = 127 //设置异常参数
	NET_DVR_GET_ZONEANDDST   = 128 //获取时区和夏时制参数
	NET_DVR_SET_ZONEANDDST   = 129 //设置时区和夏时制参数

	//注：该命令只支持4条OSD的类型，通常用于V30以下的设备版本。
	NET_DVR_GET_SHOWSTRING = 130 //获取叠加字符参数
	NET_DVR_SET_SHOWSTRING = 131 //设置叠加字符参数

	NET_DVR_GET_EVENTCOMPCFG   = 132 //获取事件触发录像参数
	NET_DVR_SET_EVENTCOMPCFG   = 133 //设置事件触发录像参数
	NET_DVR_GET_FTPCFG         = 134 //获取抓图的FTP参数(基线)
	NET_DVR_SET_FTPCFG         = 135 //设置抓图的FTP参数(基线)
	NET_DVR_GET_AUXOUTCFG      = 140 //获取报警触发辅助输出设置(HS设备辅助输出2006-02-28)
	NET_DVR_SET_AUXOUTCFG      = 141 //设置报警触发辅助输出设置(HS设备辅助输出2006-02-28)
	NET_DVR_GET_PREVIEWCFG_AUX = 142 //获取-s系列双输出预览参数(-s系列双输出2006-04-13)
	NET_DVR_SET_PREVIEWCFG_AUX = 143 //设置-s系列双输出预览参数(-s系列双输出2006-04-13)

	NET_DVR_GET_PASSWORD_MANAGE_CFG = 144 //获取密码管理配置
	NET_DVR_SET_PASSWORD_MANAGE_CFG = 145 //设置密码管理配置
	NET_DVR_UNLOCK_USER             = 146 //用户解锁
	NET_DVR_GET_SECURITY_CFG        = 147 //获取安全认证配置
	NET_DVR_SET_SECURITY_CFG        = 148 //设置安全认证配置
	NET_DVR_GET_LOCKED_INFO_LIST    = 149 //获取所有被锁定信息

	/*********************************智能部分接口 begin***************************************/
	//行为对应（NET_VCA_RULECFG）
	NET_DVR_SET_RULECFG = 152 //设置行为分析规则
	NET_DVR_GET_RULECFG = 153 //获取行为分析规则
	//球机标定参数（NET_DVR_TRACK_CFG ）
	NET_DVR_SET_TRACK_CFG = 160 //设置球机的配置参数
	NET_DVR_GET_TRACK_CFG = 161 //获取球机的配置参数

	//智能分析仪取流配置结构
	NET_DVR_SET_IVMS_STREAMCFG = 162 //设置智能分析仪取流参数
	NET_DVR_GET_IVMS_STREAMCFG = 163 //获取智能分析仪取流参数
	//智能控制参数结构
	NET_DVR_SET_VCA_CTRLCFG = 164 //设置智能控制参数
	NET_DVR_GET_VCA_CTRLCFG = 165 //获取智能控制参数
	//屏蔽区域NET_VCA_MASK_REGION_LIST
	NET_DVR_SET_VCA_MASK_REGION = 166 //设置屏蔽区域参数
	NET_DVR_GET_VCA_MASK_REGION = 167 //获取屏蔽区域参数

	//ATM进入区域 NET_VCA_ENTER_REGION
	NET_DVR_SET_VCA_ENTER_REGION = 168 //设置进入区域参数
	NET_DVR_GET_VCA_ENTER_REGION = 169 //获取进入区域参数

	//标定线配置NET_VCA_LINE_SEGMENT_LIST
	NET_DVR_SET_VCA_LINE_SEGMENT = 170 //设置标定线
	NET_DVR_GET_VCA_LINE_SEGMENT = 171 //获取标定线

	// ivms屏蔽区域NET_IVMS_MASK_REGION_LIST
	NET_DVR_SET_IVMS_MASK_REGION = 172 //设置IVMS屏蔽区域参数
	NET_DVR_GET_IVMS_MASK_REGION = 173 //获取IVMS屏蔽区域参数
	// ivms进入检测区域NET_IVMS_ENTER_REGION
	NET_DVR_SET_IVMS_ENTER_REGION = 174 //设置IVMS进入区域参数
	NET_DVR_GET_IVMS_ENTER_REGION = 175 //获取IVMS进入区域参数

	NET_DVR_SET_IVMS_BEHAVIORCFG = 176 //设置智能分析仪行为规则参数
	NET_DVR_GET_IVMS_BEHAVIORCFG = 177 //获取智能分析仪行为规则参数

	// IVMS 回放检索
	NET_DVR_IVMS_SET_SEARCHCFG = 178 //设置IVMS回放检索参数
	NET_DVR_IVMS_GET_SEARCHCFG = 179 //获取IVMS回放检索参数

	NET_DVR_SET_POSITION_TRACK = 180 // 设置场景跟踪配置信息
	NET_DVR_GET_POSITION_TRACK = 181 // 获取场景跟踪配置信息

	NET_DVR_SET_CALIBRATION = 182 // 设置标定信息
	NET_DVR_GET_CALIBRATION = 183 // 获取标定信息

	NET_DVR_SET_PDC_RULECFG = 184 // 设置人流量统计规则
	NET_DVR_GET_PDC_RULECFG = 185 // 获取人流量统计规则

	NET_DVR_SET_PU_STREAMCFG = 186 // 设置前段取流设备信息
	NET_DVR_GET_PU_STREAMCFG = 187 // 获取前段取流设备信息

	NET_VCA_SET_IVMS_BEHAVIOR_CFG = 192 // 设置IVMS行为规则配置 不带时间段
	NET_VCA_GET_IVMS_BEHAVIOR_CFG = 193 // 获取IVMS行为规则配置 不带时间段

	NET_VCA_SET_SIZE_FILTER = 194 // 设置全局尺寸过滤器
	NET_VCA_GET_SIZE_FILTER = 195 // 获取全局尺寸过滤器

	NET_DVR_SET_TRACK_PARAMCFG = 196 // 设置球机本地菜单规则
	NET_DVR_GET_TRACK_PARAMCFG = 197 // 获取球机本地菜单规则

	NET_DVR_SET_DOME_MOVEMENT_PARAM = 198 // 设置球机机芯参数
	NET_DVR_GET_DOME_MOVEMENT_PARAM = 199 // 获取球机机芯参数

	NET_DVR_GET_PICCFG_EX      = 200 //获取图象参数(SDK_V14扩展命令)
	NET_DVR_SET_PICCFG_EX      = 201 //设置图象参数(SDK_V14扩展命令)
	NET_DVR_GET_USERCFG_EX     = 202 //获取用户参数(SDK_V15扩展命令)
	NET_DVR_SET_USERCFG_EX     = 203 //设置用户参数(SDK_V15扩展命令)
	NET_DVR_GET_COMPRESSCFG_EX = 204 //获取压缩参数(SDK_V15扩展命令2006-05-15)
	NET_DVR_SET_COMPRESSCFG_EX = 205 //设置压缩参数(SDK_V15扩展命令2006-05-15)

	NET_DVR_GET_NETAPPCFG = 222 //获取网络应用参数 NTP/DDNS/EMAIL
	NET_DVR_SET_NETAPPCFG = 223 //设置网络应用参数 NTP/DDNS/EMAIL
	NET_DVR_GET_NTPCFG    = 224 //获取网络应用参数 NTP
	NET_DVR_SET_NTPCFG    = 225 //设置网络应用参数 NTP
	NET_DVR_GET_DDNSCFG   = 226 //获取网络应用参数 DDNS
	NET_DVR_SET_DDNSCFG   = 227 //设置网络应用参数 DDNS
	//对应NET_DVR_EMAILPARA
	NET_DVR_GET_EMAILCFG = 228 //获取网络应用参数 EMAIL
	NET_DVR_SET_EMAILCFG = 229 //设置网络应用参数 EMAIL

	NET_DVR_GET_NFSCFG = 230 /* NFS disk config */
	NET_DVR_SET_NFSCFG = 231 /* NFS disk config */

	/*注：该命令为定制，只支持8条OSD的类型，不会兼容V30设备版本之前的
	  NET_DVR_GET_SHOWSTRING  = 、NET_DVR_SET_SHOWSTRING 命令。（不建议使用）*/
	NET_DVR_GET_SHOWSTRING_EX = 238 //获取叠加字符参数扩展(支持8条字符)
	NET_DVR_SET_SHOWSTRING_EX = 239 //设置叠加字符参数扩展(支持8条字符)
	NET_DVR_GET_NETCFG_OTHER  = 244 //获取网络参数
	NET_DVR_SET_NETCFG_OTHER  = 245 //设置网络参数

	//对应NET_DVR_EMAILCFG结构
	NET_DVR_GET_EMAILPARACFG = 250 //Get EMAIL parameters
	NET_DVR_SET_EMAILPARACFG = 251 //Setup EMAIL parameters

	NET_DVR_GET_DDNSCFG_EX = 274 //获取扩展DDNS参数
	NET_DVR_SET_DDNSCFG_EX = 275 //设置扩展DDNS参数

	NET_DVR_SET_PTZPOS   = 292 //云台设置PTZ位置
	NET_DVR_GET_PTZPOS   = 293 //云台获取PTZ位置
	NET_DVR_GET_PTZSCOPE = 294 //云台获取PTZ范围

	NET_DVR_GET_AP_INFO_LIST  = 305 //获取无线网络资源参数
	NET_DVR_SET_WIFI_CFG      = 306 //设置IP监控设备无线参数
	NET_DVR_GET_WIFI_CFG      = 307 //获取IP监控设备无线参数
	NET_DVR_SET_WIFI_WORKMODE = 308 //设置IP监控设备网口工作模式参数
	NET_DVR_GET_WIFI_WORKMODE = 309 //获取IP监控设备网口工作模式参数
	NET_DVR_GET_WIFI_STATUS   = 310 //获取设备当前wifi连接状态
	/*********************************智能交通事件begin***************************************/
	NET_DVR_GET_REFERENCE_REGION = 400 //获取参考区域
	NET_DVR_SET_REFERENCE_REGION = 401 //设置参考区域

	NET_DVR_GET_TRAFFIC_MASK_REGION = 402 //获取交通事件屏蔽区域
	NET_DVR_SET_TRAFFIC_MASK_REGION = 403 //设置交通事件屏蔽区域
	NET_DVR_SET_AID_RULECFG         = 404 //设置交通事件规则参数
	NET_DVR_GET_AID_RULECFG         = 405 //获取交通事件规则参数

	NET_DVR_SET_TPS_RULECFG = 406 //设置交通统计规则参数
	NET_DVR_GET_TPS_RULECFG = 407 //获取交通统计规则参数

	NET_DVR_SET_LANECFG            = 408 //设置车道规则
	NET_DVR_GET_LANECFG            = 409 //获取车道规则
	NET_DVR_GET_VCA_RULE_COLOR_CFG = 410 //获取智能规则关联的颜色参数
	NET_DVR_SET_VCA_RULE_COLOR_CFG = 411 //设置智能规则关联的颜色参数
	NET_DVR_GET_SWITCH_LAMP_CFG    = 412 //获取开关灯检测规则配置参数
	NET_DVR_SET_SWITCH_LAMP_CFG    = 413 //设置开关灯检测规则配置参数

	/*********************************智能交通事件end***************************************/
	NET_DVR_SET_FACEDETECT_RULECFG = 420 // 设置人脸检测规则
	NET_DVR_GET_FACEDETECT_RULECFG = 421 // 获取人脸检测规则

	NET_DVR_SET_VEHICLE_RECOG_TASK = 422 //车辆二次识别任务提交
	NET_DVR_GET_VEHICLE_RECOG_TASK = 423 //车辆二次识别任务获取

	NET_DVR_SET_TIMECORRECT  = 432 //校时配置（只做校时操作，不记录校时配置 eg.NET_DVR_SET_TIMECFG 会修改设备的校时配置（NTP校时，会被修改为手动校时））
	NET_DVR_GET_CONNECT_LIST = 433 //获取连接设备列表信息

	/***************************DS9000新增命令(_V30) begin *****************************/
	//网络(NET_DVR_NETCFG_V30结构)
	NET_DVR_GET_NETCFG_V30 = 1000 //获取网络参数
	NET_DVR_SET_NETCFG_V30 = 1001 //设置网络参数

	//图象(NET_DVR_PICCFG_V30结构)
	NET_DVR_GET_PICCFG_V30 = 1002 //获取图象参数
	NET_DVR_SET_PICCFG_V30 = 1003 //设置图象参数

	//录像时间(NET_DVR_RECORD_V30结构)
	NET_DVR_GET_RECORDCFG_V30 = 1004 //获取录像参数
	NET_DVR_SET_RECORDCFG_V30 = 1005 //设置录像参数

	//用户(NET_DVR_USER_V30结构)
	NET_DVR_GET_USERCFG_V30 = 1006 //获取用户参数
	NET_DVR_SET_USERCFG_V30 = 1007 //设置用户参数

	//录像时间(NET_DVR_RECORD_V40结构)
	NET_DVR_GET_RECORDCFG_V40 = 1008 //获取录像参数(扩展)
	NET_DVR_SET_RECORDCFG_V40 = 1009 //设置录像参数(扩展)

	//9000DDNS参数配置(NET_DVR_DDNSPARA_V30结构)
	NET_DVR_GET_DDNSCFG_V30 = 1010 //获取DDNS(9000扩展)
	NET_DVR_SET_DDNSCFG_V30 = 1011 //设置DDNS(9000扩展)

	//EMAIL功能(NET_DVR_EMAILCFG_V30结构)
	NET_DVR_GET_EMAILCFG_V30 = 1012 //获取EMAIL参数
	NET_DVR_SET_EMAILCFG_V30 = 1013 //设置EMAIL参数

	NET_DVR_GET_NETCFG_V50 = 1015 //获取网络参数配置(V50)
	NET_DVR_SET_NETCFG_V50 = 1016 //设置网络参数配置(V50)

	NET_GET_CRUISEPOINT_V40 = 1018 //获取巡航路径配置

	//巡航参数 (NET_DVR_CRUISE_PARA结构)
	NET_DVR_GET_CRUISE = 1020
	NET_DVR_SET_CRUISE = 1021

	//报警输入结构参数 (NET_DVR_ALARMINCFG_V30结构)
	NET_DVR_GET_ALARMINCFG_V30 = 1024
	NET_DVR_SET_ALARMINCFG_V30 = 1025

	//报警输出结构参数 (NET_DVR_ALARMOUTCFG_V30结构)
	NET_DVR_GET_ALARMOUTCFG_V30 = 1026
	NET_DVR_SET_ALARMOUTCFG_V30 = 1027

	//视频输出结构参数 (NET_DVR_VIDEOOUT_V30结构)
	NET_DVR_GET_VIDEOOUTCFG_V30 = 1028
	NET_DVR_SET_VIDEOOUTCFG_V30 = 1029

	/*该命令支持8条OSD的类型（即设备版本为V30以上时），并会通过设备版本的匹配，
	  同时兼容之前的NET_DVR_GET_SHOWSTRING 、NET_DVR_SET_SHOWSTRING 命令。（建议使用）*/
	//叠加字符结构参数 (NET_DVR_SHOWSTRING_V30结构)
	NET_DVR_GET_SHOWSTRING_V30 = 1030
	NET_DVR_SET_SHOWSTRING_V30 = 1031

	//异常结构参数 (NET_DVR_EXCEPTION_V30结构)
	NET_DVR_GET_EXCEPTIONCFG_V30 = 1034
	NET_DVR_SET_EXCEPTIONCFG_V30 = 1035

	//串口232结构参数 (NET_DVR_RS232CFG_V30结构)
	NET_DVR_GET_RS232CFG_V30 = 1036
	NET_DVR_SET_RS232CFG_V30 = 1037

	//网络硬盘接入结构参数 (NET_DVR_NET_DISKCFG结构)
	NET_DVR_GET_NET_DISKCFG = 1038 //网络硬盘接入获取
	NET_DVR_SET_NET_DISKCFG = 1039 //网络硬盘接入设置
	//压缩参数 (NET_DVR_COMPRESSIONCFG_V30结构)
	NET_DVR_GET_COMPRESSCFG_V30 = 1040
	NET_DVR_SET_COMPRESSCFG_V30 = 1041

	//获取485解码器参数 (NET_DVR_DECODERCFG_V30结构)
	NET_DVR_GET_DECODERCFG_V30 = 1042 //获取解码器参数
	NET_DVR_SET_DECODERCFG_V30 = 1043 //设置解码器参数

	//获取预览参数 (NET_DVR_PREVIEWCFG_V30结构)
	NET_DVR_GET_PREVIEWCFG_V30 = 1044 //获取预览参数
	NET_DVR_SET_PREVIEWCFG_V30 = 1045 //设置预览参数

	//辅助预览参数 (NET_DVR_PREVIEWCFG_AUX_V30结构)
	NET_DVR_GET_PREVIEWCFG_AUX_V30 = 1046 //获取辅助预览参数
	NET_DVR_SET_PREVIEWCFG_AUX_V30 = 1047 //设置辅助预览参数

	//IP接入配置参数 （NET_DVR_IPPARACFG结构）
	NET_DVR_GET_IPPARACFG = 1048 //获取IP接入配置信息
	NET_DVR_SET_IPPARACFG = 1049 //设置IP接入配置信息

	//IP报警输入接入配置参数 （NET_DVR_IPALARMINCFG结构）
	NET_DVR_GET_IPALARMINCFG = 1050 //获取IP报警输入接入配置信息
	NET_DVR_SET_IPALARMINCFG = 1051 //设置IP报警输入接入配置信息

	//IP报警输出接入配置参数 （NET_DVR_IPALARMOUTCFG结构）
	NET_DVR_GET_IPALARMOUTCFG = 1052 //获取IP报警输出接入配置信息
	NET_DVR_SET_IPALARMOUTCFG = 1053 //设置IP报警输出接入配置信息

	//硬盘管理的参数获取 (NET_DVR_HDCFG结构)
	NET_DVR_GET_HDCFG = 1054 //获取硬盘管理配置参数
	NET_DVR_SET_HDCFG = 1055 //设置硬盘管理配置参数
	//盘组管理的参数获取 (NET_DVR_HDGROUP_CFG结构)
	NET_DVR_GET_HDGROUP_CFG = 1056 //获取盘组管理配置参数
	NET_DVR_SET_HDGROUP_CFG = 1057 //设置盘组管理配置参数

	//设备编码类型配置(NET_DVR_COMPRESSION_AUDIO结构)
	NET_DVR_GET_COMPRESSCFG_AUD = 1058 //获取设备语音对讲编码参数
	NET_DVR_SET_COMPRESSCFG_AUD = 1059 //设置设备语音对讲编码参数

	//IP接入配置参数 （NET_DVR_IPPARACFG_V31结构）
	NET_DVR_GET_IPPARACFG_V31 = 1060 //获取IP接入配置信息
	NET_DVR_SET_IPPARACFG_V31 = 1061 //设置IP接入配置信息

	// 通道资源配置 (NET_DVR_IPPARACFG_V40结构)
	NET_DVR_GET_IPPARACFG_V40 = 1062 // 获取IP接入配置
	NET_DVR_SET_IPPARACFG_V40 = 1063 // 设置IP接入配置

	NET_DVR_GET_CCDPARAMCFG = 1067 //IPC获取CCD参数配置
	NET_DVR_SET_CCDPARAMCFG = 1068 //IPC设置CCD参数配置

	NET_DVR_GET_IOINCFG = 1070 //获取抓拍机IO输入参数
	NET_DVR_SET_IOINCFG = 1071 //设置抓拍机IO输入参数

	NET_DVR_GET_IOOUTCFG = 1072 //获取抓拍机IO输出参数
	NET_DVR_SET_IOOUTCFG = 1073 //设置抓拍机IO输出参数

	NET_DVR_GET_FLASHCFG = 1074 //获取IO闪光灯输出参数
	NET_DVR_SET_FLASHCFG = 1075 //设置IO闪光灯输出参数

	NET_DVR_GET_LIGHTSNAPCFG = 1076 //获取抓拍机红绿灯参数
	NET_DVR_SET_LIGHTSNAPCFG = 1077 //设置抓拍机红绿灯参数

	NET_DVR_GET_MEASURESPEEDCFG = 1078 //获取抓拍机测速参数
	NET_DVR_SET_MEASURESPEEDCFG = 1079 //设置抓拍机测速参数

	NET_DVR_GET_IMAGEOVERLAYCFG = 1080 //获取抓拍机图像叠加信息参数
	NET_DVR_SET_IMAGEOVERLAYCFG = 1081 //设置抓拍机图像叠加信息参数

	NET_DVR_GET_SNAPCFG = 1082 //获取单IO触发抓拍功能配置
	NET_DVR_SET_SNAPCFG = 1083 //设置单IO触发抓拍功能配置

	NET_DVR_GET_VTPPARAM = 1084 //获取虚拟线圈参数
	NET_DVR_SET_VTPPARAM = 1085 //设置虚拟线圈参数

	NET_DVR_GET_SNAPENABLECFG = 1086 //获取抓拍机使能参数
	NET_DVR_SET_SNAPENABLECFG = 1087 //设置抓拍机使能参数

	NET_DVR_GET_POSTEPOLICECFG = 1088 //获取卡口电警参数
	NET_DVR_SET_POSTEPOLICECFG = 1089 //设置卡口电警参数

	NET_DVR_GET_JPEGCFG_V30 = 1090 //获取抓图的JPEG参数(基线)
	NET_DVR_SET_JPEGCFG_V30 = 1091 //设置抓图的JPEG参数(基线)

	NET_DVR_GET_SPRCFG = 1092 //获取车牌识别参数
	NET_DVR_SET_SPRCFG = 1093 //设置车牌识别参数
	NET_DVR_GET_PLCCFG = 1094 //获取车牌亮度补偿参数
	NET_DVR_SET_PLCCFG = 1095 //设置车牌亮度补偿参数

	NET_DVR_GET_DEVICESTATECFG = 1096 //获取设备当前状态参数
	NET_DVR_SET_CALIBRATE_TIME = 1097 //设置扩展时间校时
	NET_DVR_GET_CALIBRATE_TIME = 1098 //获取扩展时间校时

	NET_DVR_GET_DEVICECFG_V40 = 1100 //获取扩展设备参数
	NET_DVR_SET_DEVICECFG_V40 = 1101 //设置扩展设备参数

	NET_DVR_GET_ZEROCHANCFG = 1102 //获取零通道压缩参数
	NET_DVR_SET_ZEROCHANCFG = 1103 //设置零通道压缩参数

	NET_DVR_GET_ZERO_PREVIEWCFG_V30 = 1104 // 获取零通道预览参数配置
	NET_DVR_SET_ZERO_PREVIEWCFG_V30 = 1105 // 设置零通道预览参数配置

	NET_DVR_SET_ZERO_ZOOM = 1106 //设置零通道的缩放配置
	NET_DVR_GET_ZERO_ZOOM = 1107 //获取零通道的缩放配置

	NET_DVR_NATASSOCIATECFG_GET = 1110 //获取NAT功能相关信息
	NET_DVR_NATASSOCIATECFG_SET = 1111 //设置NAT功能相关信息

	NET_DVR_GET_SNMPCFG = 1112 //获取SNMP参数
	NET_DVR_SET_SNMPCFG = 1113 //设置SNMP参数

	NET_DVR_GET_SNMPCFG_V30 = 1114 //获取SNMPv30参数
	NET_DVR_SET_SNMPCFG_V30 = 1115 //设置SNMPv30参数

	NET_DVR_VIDEOPLATFORMALARMCFG_GET = 1130 //获取视频综合平台报警配置
	NET_DVR_VIDEOPLATFORMALARMCFG_SET = 1131 //设置视频综合平台报警配置

	NET_DVR_GET_RAID_ADAPTER_INFO = 1134 // 获取适配器信息
	NET_DVR_SET_RAID_ADAPTER_INFO = 1135 // 设置适配器信息

	NET_DVR_MATRIX_BIGSCREENCFG_GET = 1140 //获取大屏拼接参数
	NET_DVR_MATRIX_BIGSCREENCFG_SET = 1141 //设置大屏拼接参数

	NET_DVR_GET_MB_PLATFORMPARA = 1145 //获取平台登录参数
	NET_DVR_SET_MB_PLATFORMPARA = 1146 //设置平台登录参数
	NET_DVR_GET_MB_DEVSTATUS    = 1147 //获取移动设备状态

	NET_DVR_GET_DECODER_JOINT_CHAN = 1151 //获取解码关联通道
	NET_DVR_SET_DECODER_JOINT_CHAN = 1152 //设置解码关联通道

	//多网卡配置
	NET_DVR_GET_NETCFG_MULTI     = 1161 //获取多网卡配置
	NET_DVR_SET_NETCFG_MULTI     = 1162 //设置多网卡配置
	NET_DVR_GET_NETCFG_MULTI_V50 = 1163 //获取多网卡配置(分组)
	NET_DVR_SET_NETCFG_MULTI_V50 = 1164 //设置多网卡配置(分组)

	NET_DVR_GET_CMSPARA                = 1170 //获取平台参数
	NET_DVR_SET_CMSPARA                = 1171 //设置平台参数
	NET_DVR_GET_DIALSTATUS             = 1172 //获取拨号状态参数
	NET_DVR_GET_SMSRELATIVEPARA        = 1173 //获取短信相关参数
	NET_DVR_SET_SMSRELATIVEPARA        = 1174 //设置短信相关参数
	NET_DVR_GET_PINSTATUS              = 1175 //获取Pin状态
	NET_DVR_SET_PINCMD                 = 1176 //设置PIN命令
	NET_DVR_SET_SENSOR_CFG             = 1180 //设置模拟量参数
	NET_DVR_GET_SENSOR_CFG             = 1181 //获取模拟量参数
	NET_DVR_SET_ALARMIN_PARAM          = 1182 //设置报警输入参数
	NET_DVR_GET_ALARMIN_PARAM          = 1183 //获取报警输入参数
	NET_DVR_SET_ALARMOUT_PARAM         = 1184 //设置报警输出参数
	NET_DVR_GET_ALARMOUT_PARAM         = 1185 //获取报警输出参数
	NET_DVR_SET_SIREN_PARAM            = 1186 //设置警号参数
	NET_DVR_GET_SIREN_PARAM            = 1187 //获取警号参数
	NET_DVR_SET_ALARM_RS485CFG         = 1188 //设置报警主机485参数
	NET_DVR_GET_ALARM_RS485CFG         = 1189 //获取报警主机485参数
	NET_DVR_GET_ALARMHOST_MAIN_STATUS  = 1190 //获取报警主机主要状态
	NET_DVR_GET_ALARMHOST_OTHER_STATUS = 1191 //获取报警主机其他状态
	NET_DVR_SET_ALARMHOST_ENABLECFG    = 1192 //获取报警主机使能状态
	NET_DVR_GET_ALARMHOST_ENABLECFG    = 1193 //设置报警主机使能状态
	NET_DVR_SET_ALARM_CAMCFG           = 1194 //设置视频综合平台报警触发CAM操作配置
	NET_DVR_GET_ALARM_CAMCFG           = 1195 //设置视频综合平台报警触发CAM操作配置
	NET_DVR_GET_GATEWAY_CFG            = 1196 //获取门禁参数配置
	NET_DVR_SET_GATEWAY_CFG            = 1197 //设置门禁参数配置

	NET_DVR_GET_ALARMDIALMODECFG       = 1198 //获取报警主机拨号参数
	NET_DVR_SET_ALARMDIALMODECFG       = 1199 //设置报警主机拨号参数
	NET_DVR_SET_ALARMIN_PARAM_V50      = 1200 // 设置防区参数V50
	NET_DVR_GET_ALARMIN_PARAM_V50      = 1201 // 获取防区参数V50
	NET_DVR_SET_WINCFG                 = 1202 //窗口参数设置
	NET_DVR_GET_ALARMHOSTDIALSETUPMODE = 1204 //获取报警主机拨号启用方式
	NET_DVR_SET_ALARMHOSTDIALSETUPMODE = 1205 //设置报警主机拨号启用方式

	//视频报警主机海外版命令(视频报警主机 V1.3)
	NET_DVR_SET_SUBSYSTEM_ALARM       = 1210 //设置子系统布/撤防
	NET_DVR_GET_SUBSYSTEM_ALARM       = 1211 //获取子系统布/撤防
	NET_DVR_GET_WHITELIST_ALARM       = 1215 //获取白名单参数
	NET_DVR_SET_WHITELIST_ALARM       = 1216 //设置白名单参数
	NET_DVR_GET_ALARMHOST_MODULE_LIST = 1222 //获取所有模块
	NET_DVR_SET_PRIOR_ALARM           = 1223 //设置子系统布/撤防
	NET_DVR_GET_PRIOR_ALARM           = 1224 //获取子系统布/撤防
	NET_DVR_SET_TAMPER_ALARMIN_PARAM  = 1225 // 设置防区防拆参数
	NET_DVR_GET_TAMPER_ALARMIN_PARAM  = 1226 // 获取防区防拆参数

	NET_DVR_GET_HOLIDAY_PARAM_CFG = 1240 // 获取节假日参数
	NET_DVR_SET_HOLIDAY_PARAM_CFG = 1241 // 设置节假日参数

	NET_DVR_GET_MOTION_HOLIDAY_HANDLE = 1242 // 获取移动侦测假日报警处理方式
	NET_DVR_SET_MOTION_HOLIDAY_HANDLE = 1243 // 获取移动侦测假日报警处理方式

	NET_DVR_GET_VILOST_HOLIDAY_HANDLE = 1244 // 获取视频信号丢失假日报警处理方式
	NET_DVR_SET_VILOST_HOLIDAY_HANDLE = 1245 // 获取视频信号丢失假日报警处理方式

	NET_DVR_GET_HIDE_HOLIDAY_HANDLE = 1246 // 获取遮盖假日报警处理方式
	NET_DVR_SET_HIDE_HOLIDAY_HANDLE = 1247 // 设置遮盖假日报警处理方式

	NET_DVR_GET_ALARMIN_HOLIDAY_HANDLE  = 1248 // 获取报警输入假日报警处理方式
	NET_DVR_SET_ALARMIN_HOLIDAY_HANDLE  = 1249 // 设置报警输入假日报警处理方式
	NET_DVR_GET_ALARMOUT_HOLIDAY_HANDLE = 1250 // 获取报警输出假日报警处理方式
	NET_DVR_SET_ALARMOUT_HOLIDAY_HANDLE = 1251 // 设置报警输出假日报警处理方式
	NET_DVR_GET_HOLIDAY_RECORD          = 1252 // 获取假日录像参数
	NET_DVR_SET_HOLIDAY_RECORD          = 1253 // 设置假日录像参数
	NET_DVR_GET_NETWORK_BONDING         = 1254 // 获取BONDING网络参数
	NET_DVR_SET_NETWORK_BONDING         = 1255 // 设置BONDING网络参数
	NET_DVR_GET_LINK_STATUS             = 1256 // 获取通道IP工作状态
	NET_DVR_GET_DISK_QUOTA_CFG          = 1278 // 获取磁盘配额信息
	NET_DVR_SET_DISK_QUOTA_CFG          = 1279 // 设置磁盘配额信息
	NET_DVR_GET_JPEG_CAPTURE_CFG        = 1280 // 获取DVR抓图配置
	NET_DVR_SET_JPEG_CAPTURE_CFG        = 1281 // 设置DVR抓图配置
	NET_DVR_GET_SCHED_CAPTURECFG        = 1282 // 获取抓图计划
	NET_DVR_SET_SCHED_CAPTURECFG        = 1283 // 设置抓图计划
	NET_DVR_GET_VGA_PREVIEWCFG          = 1284 // 获取VGA预览配置
	NET_DVR_SET_VGA_PREVIEWCFG          = 1285 // 设置VGA预览配置
	NET_DVR_GET_VIDEO_INPUT_EFFECT      = 1286 // 获取通道视频输入图像参数
	NET_DVR_SET_VIDEO_INPUT_EFFECT      = 1287 // 设置通道视频输入图像参数

	NET_DVR_GET_STORAGE_SERVER_SWITCH = 1290 //获取存储服务器开关状态
	NET_DVR_SET_STORAGE_SERVER_SWITCH = 1291 //设置存储服务器开关状态

	NET_DVR_GET_DISK_QUOTA_CFG_V60 = 1292 //获取磁盘配额信息V60
	NET_DVR_SET_DISK_QUOTA_CFG_V60 = 1293 //设置磁盘配额信息V60

	NET_DVR_GET_OPTICAL_CHANNEL  = 1300 //获取光端子系统通道关联信息
	NET_DVR_SET_OPTICAL_CHANNEL  = 1301 //设置光端子系统通道关联信息
	NET_DVR_GET_FIBER_CASCADE    = 1302 //获取光纤级联模式
	NET_DVR_SET_FIBER_CASCADE    = 1303 //设置光纤级联模式
	NET_DVR_GET_SPARTAN_STATUS   = 1304 //获取畅显状态
	NET_DVR_SET_SPARTAN_STATUS   = 1305 //设置畅显状态
	NET_DVR_GET_ETHERNET_CHANNEL = 1306 //获取端口聚合参数
	NET_DVR_SET_ETHERMET_CHANNEL = 1307 //设置端口聚合参数
	NET_DVR_OPTICAL_REBOOT       = 1320 //光端机重启
	NET_DVR_SET_AUDIOCHAN_CFG    = 1321 //设置音频切换参数
	NET_DVR_GET_AUDIOCHAN_CFG    = 1322 //获取音频切换参数
	//SDI矩阵1.0
	NET_DVR_SET_MATRIX_BASE_CFG    = 1332 //设置矩阵基本参数
	NET_DVR_GET_MATRIX_BASE_CFG    = 1333 //获取矩阵基本参数
	NET_DVR_SWITCH_MATRIX_IO       = 1334 //矩阵输入输出切换
	NET_DVR_GET_MATRIX_IO_RELATION = 1335 //获取矩阵输入输入关联关系

	NET_DVR_V6PSUBSYSTEMARAM_GET = 1501 //获取V6子系统配置
	NET_DVR_V6PSUBSYSTEMARAM_SET = 1502 //设置V6子系统配置
	NET_DVR_GET_ALLWINCFG        = 1503 //窗口参数获取

	NET_DVR_BIGSCREENASSOCIATECFG_GET = 1511 //获取大屏关联配置
	NET_DVR_BIGSCREENASSOCIATECFG_SET = 1512 //设置大屏关联配置

	//1200起
	NET_DVR_GETSCREENINFO               = 1601 //获取大屏信息配置
	NET_DVR_SETSCREENINFO               = 1602 //设置大屏信息配置
	NET_DVR_GET_SCREEN_WINCFG           = 1603 //单个窗口参数获取
	NET_DVR_LAYOUTLIST_GET              = 1605 //获取布局列表
	NET_DVR_SET_LAYOUTCFG               = 1606 //布局设置
	NET_DVR_LAYOUTCTRL                  = 1607 //布局控制，1-open，2-close
	NET_DVR_INPUTLIST_GET               = 1608 //获取输入信号源列表
	NET_DVR_SET_INPUTSTREAMCFG          = 1609 //输入信号源设置
	NET_DVR_OUTPUT_SET                  = 1610 //输出参数设置
	NET_DVR_OUTPUT_GET                  = 1611 //输出参数获取
	NET_DVR_SET_OSDCFG                  = 1612 //OSD参数设置
	NET_DVR_GET_OSDCFG                  = 1613 //OSD参数获取
	NET_DVR_BIGSCREEN_GETSERIAL         = 1614 //获取大屏串口信息
	NET_DVR_GET_PLANLIST                = 1615 //获取预案列表
	NET_DVR_SET_PLAN                    = 1616 //设置预案
	NET_DVR_CTRL_PLAN                   = 1617 //控制预案
	NET_DVR_GET_DEVICE_RUN_STATUS       = 1618 //获取设备运行状态
	NET_DVR_GET_EXTERNAL_MATRIX_CFG     = 1619 //获取矩阵信息
	NET_DVR_SET_EXTERNAL_MATRIX_CFG     = 1620 //设置矩阵信息
	NET_DVR_GET_OUTPUT_SCREEN_RELATION  = 1621 //获取输出和屏幕的绑定关系
	NET_DVR_SET_OUTPUT_SCREEN_RELATION  = 1622 //设置输出和屏幕的绑定关系
	NET_DVR_GET_VCS_USER_CFG            = 1623 //获取用户信息配置
	NET_DVR_SET_VCS_USER_CFG            = 1624 //设置用户信息配置
	NET_DVR_CONTROL_SCREEN              = 1625 //屏幕控制
	NET_DVR_GET_EXTERNAL_MATRIX_CFG_V50 = 1626 //获取矩阵信息
	NET_DVR_SET_EXTERNAL_MATRIX_CFG_V50 = 1627 //设置矩阵信息

	NET_DVR_GET_DEV_BASEINFO       = 1650 //获取单个设备信息
	NET_DVR_SET_DEV_BASEINFO       = 1651 //设置单个设备信息
	NET_DVR_GET_DEV_NETINFO        = 1652 //获取设备的网络信息
	NET_DVR_SET_DEV_NETINFO        = 1653 //设置设备的网络信息
	NET_DVR_GET_SIGNAL_SOURCE_INFO = 1654 //获取信号源信息
	NET_DVR_SET_SIGNAL_SOURCE_INFO = 1655 //设置信号源信息
	NET_DVR_ADJUST_PIC_V40         = 1656 //图像微调
	NET_DVR_RESTORE_V40            = 1657 //恢复默认参数
	NET_DVR_SET_NET_SIGNAL         = 1658 //设置网络信号源
	NET_DVR_REBOOT_V40             = 1659 //重启
	NET_DVR_CONTROL_PICTURE_V41    = 1660 //图片控制V41

	NET_DVR_GET_AUTO_REBOOT_CFG = 1710 //获取自动重启参数
	NET_DVR_SET_AUTO_REBOOT_CFG = 1711 //设置自动重启参数
	NET_DVR_GET_TRUNK_USE_STATE = 1713 //获取指定干线使用状态
	NET_DVR_SET_PTZ_CTRL_INFO   = 1714 //设置PTZ控制参数
	NET_DVR_GET_PTZ_CTRL_INFO   = 1715 //获取PTZ控制参数
	NET_DVR_GET_PTZ_STATUS      = 1716 //获取PTZ状态
	NET_DVR_GET_DISP_ROUTE_LIST = 1717 //获取显示路径列表

	NET_DVR_GET_DEC_RESOURCE_LIST = 1720 //获取可用解码资源列表
	NET_DVR_SET_DEC_RESOURCE_LIST = 1721 //预分配解码资源
	NET_DVR_GET_DEC_YUV           = 1722 //获取解码通道关联YUV输出参数
	NET_DVR_SET_DEC_YUV           = 1723 //设置解码通道关联YUV输出参数
	NET_DVR_GET_DEC_RESOUCE       = 1724 //向视频综合平台申请解码资源
	NET_DVR_FREE_DEC_RESOURCE     = 1725 //释放解码资源

	NET_DVR_SET_VIDEOWALLDISPLAYMODE       = 1730 //设置电视墙拼接模式
	NET_DVR_GET_VIDEOWALLDISPLAYMODE       = 1731 //获取电视墙拼接模式
	NET_DVR_GET_VIDEOWALLDISPLAYNO         = 1732 //获取设备显示输出号
	NET_DVR_SET_VIDEOWALLDISPLAYPOSITION   = 1733 //设置显示输出位置参数
	NET_DVR_GET_VIDEOWALLDISPLAYPOSITION   = 1734 //获取显示输出位置参数
	NET_DVR_GET_VIDEOWALLWINDOWPOSITION    = 1735 //获取电视墙窗口参数
	NET_DVR_SET_VIDEOWALLWINDOWPOSITION    = 1736 //设置电视墙窗口参数
	NET_DVR_VIDEOWALLWINDOW_CLOSEALL       = 1737 //电视墙关闭所有窗口
	NET_DVR_SET_VIRTUALLED                 = 1738 //虚拟LED设置
	NET_DVR_GET_VIRTUALLED                 = 1739 //虚拟LED获取
	NET_DVR_GET_IMAGE_CUT_MODE             = 1740 //获取图像切割模式
	NET_DVR_SET_IMAGE_CUT_MODE             = 1741 //设置图像切割模式
	NET_DVR_GET_USING_SERIALPORT           = 1742 //获取当前使用串口
	NET_DVR_SET_USING_SERIALPORT           = 1743 //设置当前使用串口
	NET_DVR_SCENE_CONTROL                  = 1744 //场景控制
	NET_DVR_GET_CURRENT_SCENE              = 1745 //获取当前场景号
	NET_DVR_GET_VW_SCENE_PARAM             = 1746 //获取电视墙场景模式参数
	NET_DVR_SET_VW_SCENE_PARAM             = 1747 //设置电视墙场景模式参数
	NET_DVR_DISPLAY_CHANNO_CONTROL         = 1748 //电视墙显示编号控制
	NET_DVR_GET_WIN_DEC_INFO               = 1749 //获取窗口解码信息（批量）
	NET_DVR_RESET_VIDEOWALLDISPLAYPOSITION = 1750 //解除电视墙输出接口绑定
	NET_DVR_SET_VW_AUDIO_CFG               = 1752 //设置音频切换参数
	NET_DVR_GET_VW_AUDIO_CFG               = 1753 //获取音频切换参数
	NET_DVR_GET_GBT28181_DECCHANINFO_CFG   = 1754 //获取GBT28181协议接入设备的解码通道信息
	NET_DVR_SET_GBT28181_DECCHANINFO_CFG   = 1755 //设置GBT28181协议接入设备的解码通道信息
	NET_DVR_SET_MAINBOARD_SERIAL           = 1756 //设置主控板串口参数
	NET_DVR_GET_MAINBOARD_SERIAL           = 1757 //获取主控板串口参数
	NET_DVR_GET_SUBBOARD_INFO              = 1758 //获取子板信息
	NET_DVR_GET_SUBBOARD_EXCEPTION         = 1759 //获取异常子板异常信息

	NET_DVR_GET_CAMERACHAN_SERIALCFG = 1760 //获取Camera通道绑定串口配置
	NET_DVR_SET_CAMERACHAN_SERIALCFG = 1761 //设置Camera通道绑定串口配置
	NET_DVR_GET_MATRIX_STATUS        = 1762 //获取视频综合平台状态
	NET_SET_MULTIFUNCTION_SERIALCFG  = 1763 //设置多功能串口配置
	NET_GET_MULTIFUNCTION_SERIALCFG  = 1764 //获取多功能串口配置
	NET_DVR_PTZ_3D_SPEED             = 1765 // 3维带速度的云台控制

	NET_DVR_GET_SIGNAL_JOINT       = 1766 //获取信号源绑定配置
	NET_DVR_SET_SIGNAL_JOINT       = 1767 //设置信号源绑定配置
	NET_DVR_SIGNAL_CUT             = 1768 //信号源裁剪
	NET_DVR_DYNAMIC_DECODE_BATCH   = 1769 //批量动态解码
	NET_DVR_DECSWITCH_SET_BATCH    = 1770 //批量设置解码通道开关
	NET_DVR_DECSWITCH_GET_BATCH    = 1771 //批量获取解码通道开关
	NET_DVR_GET_ALL_SIGNAL_JOINT   = 1772 //获取所有信号源绑定配置
	NET_DVR_GET_PLAYING_PLAN       = 1773 //获取正在执行预案
	NET_DVR_WALL_RELATION_GET      = 1774 //获取设备墙与物理墙的关联
	NET_DVR_WALL_RELATION_SET      = 1775 //设置设备墙与物理墙的关联
	NET_DVR_SET_INPUTSTREAMCFG_V40 = 1776 //输入信号源设置
	NET_DVR_PTZCFG_INPUTSTREAM_GET = 1777 //获取输入源反向云台控制配置
	NET_DVR_PTZCFG_INPUTSTREAM_SET = 1778 //设置输入源反向云台控制配置
	NET_DVR_SIGNAL_CUTPARAM_GET    = 1779 //获取信号源裁剪参数

	NET_DVR_GET_SUBSYSTEM_NETCFG = 1780 //获取子系统网卡参数
	NET_DVR_SET_SUBSYSTEM_NETCFG = 1781 //设置子系统网卡参数
	NET_DVR_DEL_SIGNAL_JOINT     = 1782 //删除拼接信号源

	NET_DVR_GET_PICTURE_INFO          = 1783 //获取图片信息
	NET_DVR_SET_PICTURE_INFO          = 1784 //设置图片信息
	NET_DVR_GET_VIDEO_INFO            = 1785 //获取视频信息
	NET_DVR_SET_VIDEO_INFO            = 1786 //设置视频信息
	NET_DVR_SET_PLAYLIST              = 1787 //设置播放列表
	NET_DVR_GET_PLAYLIST              = 1788 //获取播放列表
	NET_DVR_GET_ALL_PLAYLIST          = 1789 //获取所有播放列表
	NET_DVR_PLAYITEM_CONTROL          = 1790 //播放项操作
	NET_DVR_SET_PLAYPLAN_TEMPLATE     = 1791 //设置播放计划模板
	NET_DVR_GET_PLAYPLAN_TEMPLATE     = 1792 //获取播放计划
	NET_DVR_GET_ALL_PLAYPLAN_TEMPLATE = 1793 //获取所有播放计划
	NET_DVR_SET_WINDOW_PLAYPLAN       = 1794 //设置窗口播放计划
	NET_DVR_GET_WINDOW_PLAYPLAN       = 1795 //获取窗口播放计划
	NET_DVR_TOPLAY_ITEM               = 1796 //指定播放
	NET_DVR_DEVICE_PLAY_CONTROL       = 1797 //设备播放控制
	NET_DVR_GET_PLAY_INFO             = 1798 //获取当前播放信息
	NET_DVR_GET_ALL_PICTURE_INFO      = 1799 //获取图片信息
	NET_DVR_GET_ALL_VIDEO_INFO        = 1800 //获取视频信息
	NET_DVR_DELETE_VIDEO_FILE         = 1801 //删除视频

	NET_DVR_GET_ALARMHOSTSUBSYSTEM_CFG          = 2001 //报警主机获取子系统参数
	NET_DVR_SET_ALARMHOSTSUBSYSTEM_CFG          = 2002 //报警主机设置子系统参数
	NET_DVR_GETEXTENDALARMININFO                = 2003 //获取防区编号信息
	NET_DVR_MODIFYALARMINNO                     = 2004 //修改防区编号信息
	NET_DVR_GET_ALARMHOST_WIRELESS_NETWORK_CFG  = 2005 //获取无线网络参数配置
	NET_DVR_SET_ALARMHOST_WIRELESS_NETWORK_CFG  = 2006 //设置无线网络参数配置
	NET_DVR_GET_ALARMHOST_NETCFG                = 2007 //获取网络参数配置
	NET_DVR_SET_ALARMHOST_NETCFG                = 2008 //设置网络参数配置
	NET_DVR_GET_LED_SCREEN_CFG                  = 2009 // 获取LED屏幕参数
	NET_DVR_SET_LED_SCREEN_CFG                  = 2010 // 设置LED屏幕参数
	NET_DVR_GET_LED_CONTENT_CFG                 = 2011 // 获取LED屏显内容
	NET_DVR_SET_LED_CONTENT_CFG                 = 2012 // 设置LED屏显内容
	NET_DVR_TURNON_LED                          = 2013 // 打开LED屏
	NET_DVR_TURNOFF_LED                         = 2014 // 关闭LED屏
	NET_DVR_GET_LED_TIMER_SWITCH                = 2015 // 获取LED屏定时开关参数
	NET_DVR_SET_LED_TIMER_SWITCH                = 2016 // 设置LED屏定时开关参数
	NET_DVR_SET_LED_BRIGHTNESS                  = 2017 // 手动设置LED屏亮度
	NET_DVR_GET_LED_TIMER_BRIGHTNESS            = 2018 //设置分时LED屏亮度
	NET_DVR_SET_LED_TIMER_BRIGHTNESS            = 2019 //获取分时LED屏亮度
	NET_DVR_LED_CHECKTIME                       = 2020 //LED校时
	NET_DVR_GET_ALARMHOST_AUDIO_ASSOCIATE_ALARM = 2021 //获取音频跟随报警事件
	NET_DVR_SET_ALARMHOST_AUDIO_ASSOCIATE_ALARM = 2022 //设置音频跟随报警事件
	NET_DVR_GET_LED_STATUS                      = 2023 //获取LED屏状态
	NET_DVR_CLOSE_SUBSYSTEM_FAULT_ALARM         = 2027 //关闭子系统故障提示音
	NET_DVR_SET_SUBSYSTEM_BYPASS                = 2028 //子系统旁路
	NET_DVR_CANCEL_SUBSYSTEM_BYPASS             = 2029 //子系统旁路恢复

	NET_DVR_GET_ALARMHOST_SUBSYSTEM_CFG_EX        = 2030 //获取子系统扩展参数
	NET_DVR_SET_ALARMHOST_SUBSYSTEM_CFG_EX        = 2031 //设置子系统扩展参数
	NET_DVR_GET_ALARMHOST_PRINTER_CFG             = 2032 //获取打印机打印使能
	NET_DVR_SET_ALARMHOST_PRINTER_CFG             = 2033 //设置打印机打印使能
	NET_DVR_GET_ALARMHOST_ZONE_LIST_IN_SUBSYSTEM  = 2034 //获取指定子系统内的所有防区
	NET_DVR_GET_ALARMHOST_TRIGGER_LIST            = 2035 //获取所有触发器
	NET_DVR_ARM_ALARMHOST_SUBSYSTEM               = 2036 //按布防类型对子系统布防
	NET_DVR_GET_ALARMHOST_EVENT_TRIG_ALARMOUT_CFG = 2037 // 获取事件触发报警输出配置
	NET_DVR_SET_ALARMHOST_EVENT_TRIG_ALARMOUT_CFG = 2038 // 设置事件触发报警输出配置
	NET_DVR_GET_ALARMHOST_FAULT_CFG               = 2039 // 获取故障处理配置
	NET_DVR_SET_ALARMHOST_FAULT_CFG               = 2040 // 设置故障处理配置
	NET_DVR_SEARCH_ARMHOST_EXTERNAL_MODULE        = 2041 //自动搜索
	NET_DVR_REGISTER_ALARMHOST_EXTERNAL_MODULE    = 2042 //自动注册
	NET_DVR_CLOSE_ALARMHOST_OVERALL_FAULT_ALARM   = 2043 //关闭全局故障提示音

	NET_DVR_GET_SAFETYCABIN_WORK_MODE         = 2044 //获取防护舱工作模式参数
	NET_DVR_SET_SAFETYCABIN_WORK_MODE         = 2045 //设置防护舱工作模式参数
	NET_DVR_GET_SAFETYCABIN_PERSON_SIGNAL_CFG = 2046 //获取防护舱人信号探测参数
	NET_DVR_SET_SAFETYCABIN_PERSON_SIGNAL_CFG = 2047 //设置防护舱人信号探测参数

	NET_DVR_GET_ALARMHOST_MODULE_CFG = 2048 //获取模块信息
	//     NET_DVR_SET_ALARMHOST_MODULE_CFG            2049//设置模块信息(预留)

	NET_DVR_GET_ALARMHOST_EXTERNAL_DEVICE_STATE       = 2050 //获取485外接设备状态
	NET_DVR_SET_ALARMHOST_EXTERNAL_DEVICE_LIMIT_VALUE = 2051 //设置外接设备报警限值
	NET_DVR_GET_ALARMHOST_EXTERNAL_DEVICE_LIMIT_VALUE = 2052 //获取外接设备报警限值
	NET_DVR_GET_ALARMHOST_SENSOR_JOINT_CFG            = 2053 // 获取模拟量关联配置
	NET_DVR_SET_ALARMHOST_SENSOR_JOINT_CFG            = 2054 // 设置模拟量关联配置
	NET_DVR_SET_ALARMHOST_RS485_SLOT_CFG              = 2055 // 设置报警主机485槽位参数
	NET_DVR_GET_ALARMHOST_RS485_SLOT_CFG              = 2056 // 获取报警主机485槽位参数

	NET_DVR_GET_ALL_VARIABLE_INFO      = 2057 // 获取所有变量元素信息
	NET_DVR_GET_ALARM_POINT_CFG        = 2058 // 获取点号信息
	NET_DVR_SET_ALARM_POINT_CFG        = 2059 // 设置点号信息
	NET_DVR_GET_HISTORY_VALUE          = 2060 // 获取历史数据
	NET_DVR_GET_ALARMHOST_ALARM_MODE   = 2061 // 获取数据上传方式
	NET_DVR_SET_ALARMHOST_ALARM_MODE   = 2062 // 设置数据上传方式
	NET_DVR_GET_ALARMHOST_SENSOR_VALUE = 2063 // 获取模拟量实时数据

	NET_DVR_GET_ALARMHOST_REPORT_CENTER_V40 = 2064 // 获取数据上传方式
	NET_DVR_SET_ALARMHOST_REPORT_CENTER_V40 = 2065 // 设置数据上传方式
	NET_DVR_GET_OUTPUT_SCHEDULE_RULECFG     = 2068 // 获取时控输出参数
	NET_DVR_SET_OUTPUT_SCHEDULE_RULECFG     = 2069 // 设置时控输出参数
	NET_DVR_GET_CMS_CFG                     = 2070
	NET_DVR_SET_CMS_CFG                     = 2071

	NET_DVR_GET_PASSTHROUGH_CAP = 2073 //获取设备透传能力集

	NET_DVR_GET_ALARMHOST_MAIN_STATUS_V40 = 2072 // 获取主要状态V40
	NET_DVR_GET_ALARMHOST_MAIN_STATUS_V51 = 2083 // 获取主要状态V51

	/*************************************视频报警主机1.3 begin*************************************/
	NET_DVR_GET_ALARM_CAPTRUE_CFG           = 2074 //获取报警抓图参数配置
	NET_DVR_SET_ALARM_CAPTRUE_CFG           = 2075 //设置报警抓图参数配置
	NET_DVR_GET_ONE_OUTPUT_SCH_RULECFG_V40  = 2078 // 获取单个时控输出参数V40
	NET_DVR_SET_ONE_OUTPUT_SCH_RULECFG_V40  = 2079 // 设置单个时控输出参数V40
	NET_DVR_GET_OUTPUT_SCHEDULE_RULECFG_V40 = 2080 // 获取时控输出参数V40
	NET_DVR_SET_OUTPUT_SCHEDULE_RULECFG_V40 = 2081 // 设置时控输出参数V40
	NET_DVR_ALARMHOST_CLOSE_SUBSYSTEM       = 2082 //对子系统撤防操作
	/*************************************视频报警主机1.3 end**************************************/

	NET_DVR_GET_WEEK_PLAN_CFG              = 2100 //获取门状态周计划参数
	NET_DVR_SET_WEEK_PLAN_CFG              = 2101 //设置门状态周计划参数
	NET_DVR_GET_DOOR_STATUS_HOLIDAY_PLAN   = 2102 //获取门状态假日计划参数
	NET_DVR_SET_DOOR_STATUS_HOLIDAY_PLAN   = 2103 //设置门状态假日计划参数
	NET_DVR_GET_DOOR_STATUS_HOLIDAY_GROUP  = 2104 //获取门状态假日组参数
	NET_DVR_SET_DOOR_STATUS_HOLIDAY_GROUP  = 2105 //设置门状态假日组参数
	NET_DVR_GET_DOOR_STATUS_PLAN_TEMPLATE  = 2106 //获取门状态计划模板参数
	NET_DVR_SET_DOOR_STATUS_PLAN_TEMPLATE  = 2107 //设置门状态计划模板参数
	NET_DVR_GET_DOOR_CFG                   = 2108 //获取门参数
	NET_DVR_SET_DOOR_CFG                   = 2109 //设置门参数
	NET_DVR_GET_DOOR_STATUS_PLAN           = 2110 //获取门状态计划参数
	NET_DVR_SET_DOOR_STATUS_PLAN           = 2111 //设置门状态计划参数
	NET_DVR_GET_GROUP_CFG                  = 2112 //获取群组参数
	NET_DVR_SET_GROUP_CFG                  = 2113 //设置群组参数
	NET_DVR_GET_MULTI_CARD_CFG             = 2114 //获取多重卡参数
	NET_DVR_SET_MULTI_CARD_CFG             = 2115 //设置多重卡参数
	NET_DVR_GET_CARD_CFG                   = 2116 //获取卡参数
	NET_DVR_SET_CARD_CFG                   = 2117 //设置卡参数
	NET_DVR_CLEAR_ACS_PARAM                = 2118 //清空门禁主机参数
	NET_DVR_GET_SNEAK_CFG                  = 2119 //获取反潜回参数
	NET_DVR_SET_SNEAK_CFG                  = 2120 //设置反潜回参数
	NET_DVR_GET_MULTI_DOOR_INTERLOCK_CFG   = 2121 //获取多门互锁参数
	NET_DVR_SET_MULTI_DOOR_INTERLOCK_CFG   = 2122 //设置多门互锁参数
	NET_DVR_GET_ACS_WORK_STATUS            = 2123 //获取门禁主机工作状态
	NET_DVR_GET_VERIFY_WEEK_PLAN           = 2124 //获取读卡器验证方式周计划参数
	NET_DVR_SET_VERIFY_WEEK_PLAN           = 2125 //设置读卡器验证方式周计划参数
	NET_DVR_GET_CARD_RIGHT_WEEK_PLAN       = 2126 //获取卡权限周计划参数
	NET_DVR_SET_CARD_RIGHT_WEEK_PLAN       = 2127 //设置卡权限周计划参数
	NET_DVR_GET_VERIFY_HOLIDAY_PLAN        = 2128 //获取读卡器验证方式假日计划参数
	NET_DVR_SET_VERIFY_HOLIDAY_PLAN        = 2129 //设置读卡器验证方式假日计划参数
	NET_DVR_GET_CARD_RIGHT_HOLIDAY_PLAN    = 2130 //获取卡权限假日计划参数
	NET_DVR_SET_CARD_RIGHT_HOLIDAY_PLAN    = 2131 //设置卡权限假日计划参数
	NET_DVR_GET_VERIFY_HOLIDAY_GROUP       = 2132 //获取读卡器验证方式假日组参数
	NET_DVR_SET_VERIFY_HOLIDAY_GROUP       = 2133 //设置读卡器验证方式假日组参数
	NET_DVR_GET_CARD_RIGHT_HOLIDAY_GROUP   = 2134 //获取卡权限假日组参数
	NET_DVR_SET_CARD_RIGHT_HOLIDAY_GROUP   = 2135 //设置卡权限假日组参数
	NET_DVR_GET_VERIFY_PLAN_TEMPLATE       = 2136 //获取读卡器验证方式计划模板参数
	NET_DVR_SET_VERIFY_PLAN_TEMPLATE       = 2137 //设置读卡器验证方式计划模板参数
	NET_DVR_GET_CARD_RIGHT_PLAN_TEMPLATE   = 2138 //获取卡权限计划模板参数
	NET_DVR_SET_CARD_RIGHT_PLAN_TEMPLATE   = 2139 //设置卡权限计划模板参数
	NET_DVR_GET_CARD_READER_CFG            = 2140 //获取读卡器参数
	NET_DVR_SET_CARD_READER_CFG            = 2141 //设置读卡器参数
	NET_DVR_GET_CARD_READER_PLAN           = 2142 //获取读卡器验证计划参数
	NET_DVR_SET_CARD_READER_PLAN           = 2143 //设置读卡器验证计划参数
	NET_DVR_GET_CASE_SENSOR_CFG            = 2144 //获取事件触发器参数
	NET_DVR_SET_CASE_SENSOR_CFG            = 2145 //设置事件触发器参数
	NET_DVR_GET_CARD_READER_ANTI_SNEAK_CFG = 2146 //获取读卡器反潜回参数
	NET_DVR_SET_CARD_READER_ANTI_SNEAK_CFG = 2147 //设置读卡器反潜回参数
	NET_DVR_GET_PHONE_DOOR_RIGHT_CFG       = 2148 //获取手机关联门权限参数
	NET_DVR_SET_PHONE_DOOR_RIGHT_CFG       = 2149 //获取手机关联门权限参数
	NET_DVR_GET_FINGERPRINT_CFG            = 2150 //获取指纹参数
	NET_DVR_SET_FINGERPRINT_CFG            = 2151 //设置指纹参数
	NET_DVR_DEL_FINGERPRINT_CFG            = 2152 //删除指纹参数
	NET_DVR_GET_EVENT_CARD_LINKAGE_CFG     = 2153 //获取事件卡号联动配置参数
	NET_DVR_SET_EVENT_CARD_LINKAGE_CFG     = 2154 //设置事件卡号联动配置参数
	NET_DVR_GET_ANTI_SNEAK_HOST_CFG        = 2155 //获取主机组反潜回参数
	NET_DVR_SET_ANTI_SNEAK_HOST_CFG        = 2156 //设置主机组反潜回参数
	NET_DVR_GET_READER_ANTI_SNEAK_HOST_CFG = 2157 //获取主机组读卡器反潜回参数
	NET_DVR_SET_READER_ANTI_SNEAK_HOST_CFG = 2158 //设置主机组读卡器反潜回参数
	NET_DVR_GET_ACS_CFG                    = 2159 //获取门禁主机参数
	NET_DVR_SET_ACS_CFG                    = 2160 //设置门禁主机参数
	NET_DVR_GET_CARD_PASSWD_CFG            = 2161 //获取卡密码开门使能配置
	NET_DVR_SET_CARD_PASSWD_CFG            = 2162 //设置卡密码开门使能配置
	NET_DVR_GET_CARD_USERINFO_CFG          = 2163 //获取卡号关联用户信息参数
	NET_DVR_SET_CARD_USERINFO_CFG          = 2164 //设置卡号关联用户信息参数

	NET_DVR_GET_ACS_EXTERNAL_DEV_CFG       = 2165 //获取门禁主机串口外设参数
	NET_DVR_SET_ACS_EXTERNAL_DEV_CFG       = 2166 //设置门禁主机串口外设参数
	NET_DVR_GET_PERSONNEL_CHANNEL_CFG      = 2167 //获取人员通道参数
	NET_DVR_SET_PERSONNEL_CHANNEL_CFG      = 2168 //设置人员通道参数
	NET_DVR_SET_PLATFORM_VERIFY_CFG        = 2169 //下发平台认证结果
	NET_DVR_GET_PERSON_STATISTICS_CFG      = 2170 //获取人数统计参数
	NET_DVR_SET_PERSON_STATISTICS_CFG      = 2171 //设置人数统计参数
	NET_DVR_GET_ACS_SCREEN_DISPLAY_CFG     = 2172 //获取屏幕字符串显示参数
	NET_DVR_SET_ACS_SCREEN_DISPLAY_CFG     = 2173 //设置屏幕字符串显示参数
	NET_DVR_GET_GATE_TIME_CFG              = 2174 //获取人员通道闸门时间参数
	NET_DVR_SET_GATE_TIME_CFG              = 2175 //设置人员通道闸门时间参数
	NET_DVR_GET_LOCAL_CONTROLLER_STATUS    = 2176 //获取就地控制器状态
	NET_DVR_GET_ONLINE_LOCAL_CONTROLLER    = 2177 //搜索在线就地控制器
	NET_DVR_GET_CARD_CFG_V50               = 2178 //获取新卡参数(V50)
	NET_DVR_SET_CARD_CFG_V50               = 2179 //设置新卡参数(V50)
	NET_DVR_GET_ACS_WORK_STATUS_V50        = 2180 //获取门禁主机工作状态(V50)
	NET_DVR_GET_EVENT_CARD_LINKAGE_CFG_V50 = 2181 //获取事件卡号联动配置参数(V50)
	NET_DVR_SET_EVENT_CARD_LINKAGE_CFG_V50 = 2182 //设置事件卡号联动配置参数(V50)
	NET_DVR_GET_FINGERPRINT_CFG_V50        = 2183 //获取指纹参数V50
	NET_DVR_SET_FINGERPRINT_CFG_V50        = 2184 //设置指纹参数V50

	NET_DVR_GET_SAFETYCABIN_STATE = 2197 //获取防护舱状态
	NET_DVR_GET_RS485_CASCADE_CFG = 2198 //获取Rs485级联设备配置
	NET_DVR_SET_RS485_CASCADE_CFG = 2199 //设置Rs485级联设备配置

	/*************************************视频报警主机2.0 begin*************************************/
	NET_DVR_GET_REMOTECONTROLLER_PERMISION_CFG    = 2200 //获取遥控器权限参数
	NET_DVR_SET_REMOTECONTROLLER_PERMISION_CFG    = 2201 //设置遥控器权限参数
	NET_DVR_GET_KEYBOARD_CFG                      = 2202 //获取键盘参数配置
	NET_DVR_SET_KEYBOARD_CFG                      = 2203 //设置键盘参数配置
	NET_DVR_GET_ALARMHOST_WIRELESS_BUSINNESS_INFO = 2204 //无线业务查询
	NET_DVR_GET_ALL_REMOTECONTROLLER_LIST         = 2205 //获取所有遥控器
	NET_DVR_GET_PREVIEW_DELAY_CFG                 = 2206 //获取延迟预览参数配置
	NET_DVR_SET_PREVIEW_DELAY_CFG                 = 2207 //设置延迟预览参数配置
	NET_DVR_GET_ZONE_CHANNEL_LINKAGE_CFG          = 2208 //获取防区联动视频通道配置
	NET_DVR_SET_ZONE_CHANNEL_LINKAGE_CFG          = 2209 //设置防区联动视频通道配置
	NET_DVR_GET_CENTER_SERVER_CFG                 = 2210 //获取报警中心服务器
	NET_DVR_SET_CENTER_SERVER_CFG                 = 2211 //设置报警中心服务器
	/*************************************视频报警主机2.0 end**************************************/

	/********************************一键式紧急报警产品V1.0.0 begin********************************/
	NET_DVR_GET_EMERGENCE_ALARM_PRODUCT_CAP  = 2212 //获取一键式紧急报警产品能力
	NET_DVR_GET_CALL_WAITTING_CFG_CAP        = 2213 //获取呼叫等待参数配置能力
	NET_DVR_GET_CALL_WAITTING_CFG            = 2214 //获取呼叫等待参数配置
	NET_DVR_SET_CALL_WAITTING_CFG            = 2215 //设置呼叫等待参数配置
	NET_DVR_GET_ALARM_LAMP_CFG_CAP           = 2216 //获取警灯参数配置能力
	NET_DVR_GET_ALARM_LAMP_CFG               = 2217 //获取警灯参数配置
	NET_DVR_SET_ALARM_LAMP_CFG               = 2218 //设置警灯参数配置
	NET_DVR_GET_VOICE_PROMPTION_CFG_CAP      = 2219 //获取语音提示配置能力
	NET_DVR_GET_VOICE_PROMPTION_CFG          = 2220 //获取语音提示配置
	NET_DVR_SET_VOICE_PROMPTION_CFG          = 2221 //设置语音提示配置
	NET_DVR_GET_EMERGENCE_ALARM_RESPONSE_CAP = 2222 //获取紧急报警处理能力
	NET_DVR_EMERGENCE_ALARM_RESPONSE_CTRL    = 2223 //紧急报警处理控制
	/********************************一键式紧急报警产品V1.0.0 end**********************************/

	//网络报警主机 V2.2
	NET_DVR_GET_ALARMHOST_NETCFG_V50 = 2224 //获取报警主机网络参数配置V50
	NET_DVR_SET_ALARMHOST_NETCFG_V50 = 2225 //设置报警主机网络参数配置V50
	NET_DVR_REGISTER_ALARM_RS485     = 2226 //RS485重新注册
	/**********************************动环报警主机V3.0****************************************/

	NET_DVR_GET_ALARMIN_PARAM_LIST = 2227 //获取防区参数列表
	//无线报警主机1.0.0
	NET_DVR_GET_ALARMHOST_OTHER_STATUS_V50   = 2228 //获取报警主机其他状态v50
	NET_DVR_GET_ALARMHOST_OTHER_STATUS_V51   = 2236 //获取报警主机其他状态V51
	NET_DVR_GET_ALARMIN_ASSOCIATED_CHAN_LIST = 2229 //获取防区防区联动视频通道参数列表
	NET_DVR_GET_ALARMIN_TRIGGER              = 2230 //获取报警主机防区联动配置
	NET_DVR_SET_ALARMIN_TRIGGER              = 2231 //设置报警主机防区联动配置
	NET_DVR_GET_EMERGENCY_CALL_HELP_TRIGGER  = 2232 //获取报警主机紧急求助联动配置
	NET_DVR_SET_EMERGENCY_CALL_HELP_TRIGGER  = 2233 //设置报警主机紧急求助联动配置
	NET_DVR_GET_CONSULT_TRIGGER              = 2234 //获取报警主机业务咨询联动配置
	NET_DVR_SET_CONSULT_TRIGGER              = 2235 //设置报警主机业务咨询联动配置
	NET_DVR_GET_ALARMIN_PARAM_LIST_V50       = 2237 //获取防区参数列表V50

	NET_DVR_GET_CARD_RIGHT_WEEK_PLAN_V50 = 2304 //获取卡权限周计划参数V50
	NET_DVR_SET_CARD_RIGHT_WEEK_PLAN_V50 = 2305 //设置卡权限周计划参数V50

	NET_DVR_GET_CARD_RIGHT_HOLIDAY_PLAN_V50 = 2310 //获取卡权限假日计划参数V50
	NET_DVR_SET_CARD_RIGHT_HOLIDAY_PLAN_V50 = 2311 //设置卡权限假日计划参数V50

	NET_DVR_GET_CARD_RIGHT_HOLIDAY_GROUP_V50 = 2316 //获卡权限假日组参数V50
	NET_DVR_SET_CARD_RIGHT_HOLIDAY_GROUP_V50 = 2317 //设置卡权限假日组参数V50

	NET_DVR_GET_CARD_RIGHT_PLAN_TEMPLATE_V50 = 2322 //获取卡权限计划模板参数V50
	NET_DVR_SET_CARD_RIGHT_PLAN_TEMPLATE_V50 = 2323 //设置卡权限计划模板参数V50

	/**********************************经济型指纹门禁产品V1.0 设备不做****************************************/
	NET_DVR_GET_SCHEDULE_INFO           = 2500 //获取排班信息
	NET_DVR_GET_ATTENDANCE_SUMMARY_INFO = 2501 //获取考勤汇总信息
	NET_DVR_GET_ATTENDANCE_RECORD_INFO  = 2502 //获取考勤记录信息
	NET_DVR_GET_ABNORMAL_INFO           = 2503 //获取异常统计信息
	/**********************************经济型指纹门禁产品V1.0****************************************/

	/*************************************视频门禁一体机1.0 begin**************************************/
	NET_DVR_CAPTURE_FINGERPRINT_INFO = 2504 //采集指纹信息
	/*************************************视频门禁一体机1.0 end**************************************/

	/*************************************嵌入式智能终端1.0 begin**************************************/
	NET_DVR_BULK_UPLOAD_BLACK_LIST_PICTURE = 2520 //批量上传黑名单图片
	NET_DVR_BULK_UPLOAD_ID_BLACK_LIST      = 2521 //批量上传身份证黑名单
	NET_DVR_GET_FAILED_FACE_INFO           = 2522 //获取设备升级建模失败的人脸记录
	NET_DVR_GET_FACE_AND_TEMPLATE          = 2523 //获取人脸及模板数据
	NET_DVR_SET_FACE_AND_TEMPLATE          = 2524 //设置人脸及模板数据
	/*************************************嵌入式智能终端1.0 end**************************************/

	/*************************************人脸识别门禁一体机1.0 begin**************************************/
	NET_DVR_GET_CARD_READER_CFG_V50 = 2505 //获取读卡器参数(V50)
	NET_DVR_SET_CARD_READER_CFG_V50 = 2506 //设置读卡器参数(V50)
	NET_DVR_GET_FACE_PARAM_CFG      = 2507 //获取人脸参数
	NET_DVR_SET_FACE_PARAM_CFG      = 2508 //设置人脸参数
	NET_DVR_DEL_FACE_PARAM_CFG      = 2509 //删除人脸参数
	NET_DVR_CAPTURE_FACE_INFO       = 2510 //采集人脸信息
	/*************************************人脸识别门禁一体机1.0 end**************************************/
	NET_DVR_GET_REGISTER_INFO = 2511 //登记信息获取

	NET_DVR_GET_SMSRELATIVEPARA_V50        = 2512 //获取短信相关参数
	NET_DVR_SET_SMSRELATIVEPARA_V50        = 2513 //设置短信相关参数
	NET_DVR_GET_ACS_EVENT                  = 2514 //设备事件获取
	NET_DVR_GET_MULTI_CARD_CFG_V50         = 2515 //获取多重卡参数V50
	NET_DVR_SET_MULTI_CARD_CFG_V50         = 2516 //设置多重卡参数V50
	NET_DVR_DEL_FINGERPRINT_CFG_V50        = 2517 //删除指纹参数V50
	NET_DVR_GET_EVENT_CARD_LINKAGE_CFG_V51 = 2518 //获取事件卡号联动配置参数(V51)
	NET_DVR_SET_EVENT_CARD_LINKAGE_CFG_V51 = 2519 //设置事件卡号联动配置参数(V51)

	NET_DVR_SET_EXAM_INFO              = 2530 //考试信息下发
	NET_DVR_SET_EXAMINEE_INFO          = 2531 //考生信息下发
	NET_DVR_SEARCH_EXAM_COMPARE_RESULT = 2532 //考试比对结果查询
	NET_DVR_BULK_CHECK_FACE_PICTURE    = 2533 //批量校验人脸图片
	NET_DVR_JSON_CONFIG                = 2550 //JSON透传数据
	NET_DVR_FACE_DATA_RECORD           = 2551 //添加人脸数据到人脸库
	NET_DVR_FACE_DATA_SEARCH           = 2552 //查询人脸库中的人脸数据
	NET_DVR_FACE_DATA_MODIFY           = 2553 //修改人脸库中的人脸数据
	NET_DVR_CAPTURE_DATA_SEARCH        = 2554 //查询离线采集数据集中数据

	NET_DVR_GET_CARD        = 2560
	NET_DVR_SET_CARD        = 2561
	NET_DVR_DEL_CARD        = 2562
	NET_DVR_GET_FINGERPRINT = 2563
	NET_DVR_SET_FINGERPRINT = 2564
	NET_DVR_DEL_FINGERPRINT = 2565
	NET_DVR_GET_FACE        = 2566
	NET_DVR_SET_FACE        = 2567

	NET_DVR_GET_ALL_ALARM_RS485CFG           = 2705 // 获取485参数
	NET_DVR_GET_ALL_ALARMHOST_RS485_SLOT_CFG = 2706 // 获取所有报警主机485槽位参数
	NET_DVR_GET_DEVICE_SELF_CHECK_STATE      = 2707 //获取设备自检功能
	NET_DVR_GET_ALL_ALARM_POINT_CFG          = 2708 // 获取所有点号参数
	NET_DVR_GET_ALL_ALARM_SENSOR_CFG         = 2709 // 获取所有模拟量参数
	NET_DVR_GET_ALL_ALARM_SENSOR_JOINT       = 2710 //获取所有模拟量联动参数
	NET_DVR_GET_AIR_CONDITION_PARAM          = 2711 //获取空调参数
	NET_DVR_GET_OUT_SCALE_CFG                = 2712 //获取主辅口输出配置
	NET_DVR_SET_OUT_SCALE_CFG                = 2713 //设置主辅口输出配置
	NET_DVR_GET_ALARM_CHAN_ABLITITY          = 2714 //获取报警相关通道参
	/**********************************动环报警主机V3.0****************************************/

	//动环报警主机D2000 V1.0
	NET_DVR_GET_ALARMCENTER_NETCFG = 2715 //获取报警中心网络参数配置
	NET_DVR_SET_ALARMCENTER_NETCFG = 2716 //获取报警中心网络参数配置

	NET_ITC_GET_TRIGGERCFG      = 3003 //获取触发参数
	NET_ITC_SET_TRIGGERCFG      = 3004 //设置触发参数
	NET_ITC_GET_IOOUT_PARAM_CFG = 3005 //获取IO输出参数（3.1含之后版本）
	NET_ITC_SET_IOOUT_PARAM_CFG = 3006 //设置IO输出参数（3.1含之后版本）

	NET_DVR_GET_CAMERA_SETUPCFG = 3007 //获取相机架设参数
	NET_DVR_SET_CAMERA_SETUPCFG = 3008 //设置相机架设参数

	NET_ITC_GET_TRIGGER_DEFAULTCFG = 3013 //获取触发模式推荐参数
	NET_DVR_GET_STATUS_DETECTCFG   = 3015 //获取状态检测使能参数
	NET_DVR_SET_STATUS_DETECTCFG   = 3016 //设置状态检测使能参数
	NET_ITC_GET_VIDEO_TRIGGERCFG   = 3017 //获取视频电警触发参数
	NET_ITC_SET_VIDEO_TRIGGERCFG   = 3018 //设置视频电警触发参数
	NET_DVR_GET_TPS_ALARMCFG       = 3019 //获取交通统计报警参数
	NET_DVR_SET_TPS_ALARMCFG       = 3020 //设置交通统计报警参数

	NET_DVR_GET_REDAREACFG               = 3100 //获取红绿灯区域参数
	NET_DVR_SET_REDAREACFG               = 3101 //设置红绿灯区域参数
	NET_DVR_GET_TEST_SPOT                = 3102 //获取SPOT口测试总步数和当前第几步
	NET_DVR_SET_TEST_SPOT                = 3103 //设置SPOT口测试总步数和当前第几步
	NET_DVR_GET_CABINETCFG               = 3104 //机柜参数配置获取
	NET_DVR_SET_CABINETCFG               = 3105 //机柜参数配置设置
	NET_DVR_VEHICLE_CHECK_START          = 3106 //黑名单稽查数据回传
	NET_DVR_SET_CAPTUREPIC_CFG           = 3107 //车载抓图配置设置命令
	NET_DVR_GET_CAPTUREPIC_CFG           = 3108 //车载抓图配置获取命令
	NET_DVR_SET_MOBILEPLATE_RECOG_CFG    = 3109 //车载车牌识别配置设置命令
	NET_DVR_GET_MOBILEPLATE_RECOG_CFG    = 3110 //车载车牌识别配置获取命令
	NET_DVR_SET_MOBILE_RADAR_CFG         = 3111 //车载雷达配置设置命令
	NET_DVR_GET_MOBILE_RADAR_CFG         = 3112 //车载雷达配置获取命令
	NET_DVR_SET_MOBILE_LOCALPLATECHK_CFG = 3113 //车载黑名单本地对比配置设置命令
	NET_DVR_GET_MOBILE_LOCALPLATECHK_CFG = 3114 //车载黑名单本地对比配置获取命令

	NET_ITC_GET_ICRCFG           = 3115 //获取ICR配置切换
	NET_ITC_SET_ICRCFG           = 3116 //设置ICR配置切换
	NET_ITC_GET_RS485_ACCESSINFO = 3117 //获取Rs485关联接入设备的信息
	NET_ITC_SET_RS485_ACCESSINFO = 3118 //设置Rs485关联接入设备的信息
	NET_ITC_GET_EXCEPTIONCFG     = 3119 //获取异常参数
	NET_ITC_SET_EXCEPTIONCFG     = 3120 //设置异常参数
	NET_ITC_GET_FTPCFG           = 3121 //获取ITC  FTP设置参数
	NET_ITC_SET_FTPCFG           = 3122 //设置ITC FTP设置参数

	NET_DVR_VEHICLE_CONTROL_LIST_START   = 3123 //设置车辆黑白名单信息
	NET_DVR_GET_ALL_VEHICLE_CONTROL_LIST = 3124 //获取所有车辆黑白名单信息
	NET_DVR_VEHICLE_DELINFO_CTRL         = 3125 //删除设备内黑名单数据库信息
	NET_DVR_GET_ENTRANCE_PARAMCFG        = 3126 //获取出入口控制参数
	NET_DVR_SET_ENTRANCE_PARAMCFG        = 3127 //设置出入口控制参数
	NET_DVR_BARRIERGATE_CTRL             = 3128 //远程控制道闸
	NET_DVR_GATELAMP_CTRL                = 3129 //常亮灯功能
	NET_DVR_GET_CURTRIGGERMODE           = 3130 //获取设备当前触发模式
	NET_DVR_GET_GPSDATACFG               = 3131 //获取GPS参数
	NET_DVR_SET_GPSDATACFG               = 3132 //设置GPS参数
	NET_DVR_VEHICLELIST_CTRL_START       = 3133 //设置车辆黑白名单信息

	NET_DVR_GET_GUARDCFG         = 3134 //获取车牌识别检测计划
	NET_DVR_SET_GUARDCFG         = 3135 //设置车牌识别检测计划
	NET_DVR_GET_SNAPINFO_CFG     = 3136 //获取抓拍图片参数
	NET_DVR_SET_SNAPINFO_CFG     = 3137 //设置抓拍图片参数
	NET_DVR_GET_SNAPINFO_CFG_V40 = 3138 //获取抓拍图片参数扩展
	NET_DVR_SET_SNAPINFO_CFG_V40 = 3139 //设置抓拍图片参数扩展
	NET_DVR_SET_CURTRIGGERMODE   = 3140 //设置设备当前触发模式(仅IPC/D支持)
	NET_DVR_GET_TRAFFIC_DATA     = 3141 //长连接获取交通数据
	NET_DVR_GET_TRAFFIC_FLOW     = 3142 //长连接获取交通流量
	NET_DVR_PARKING_VEHICLE_SEND = 3143 //停车车辆信息下发
	NET_DVR_PARKING_CARD_SEND    = 3144 //停车卡下发
	NET_DVR_PARKING_CARD_CTRL    = 3145 //停车场停车卡控制接口

	NET_DVR_GET_ALARMCTRL_CAPABILITIES = 3146 //获取报警控制能力
	NET_DVR_SET_ALARMCTRL_CFG          = 3147 //设置报警控制
	NET_DVR_GET_ALARMCTRL_CFG          = 3148 //获取报警控制

	NET_DVR_GET_AUDIO_INPUT           = 3201 //获取音频输入参数
	NET_DVR_SET_AUDIO_INPUT           = 3202 //设置音频输入参数
	NET_DVR_GET_CAMERA_DEHAZE_CFG     = 3203 //获取透雾参数配置
	NET_DVR_SET_CAMERA_DEHAZE_CFG     = 3204 //设置透雾参数配置
	NET_DVR_REMOTECONTROL_ALARM       = 3205 //远程控制遥控器布防
	NET_DVR_REMOTECONTROL_DISALARM    = 3206 //远程控制遥控器撤防
	NET_DVR_REMOTECONTROL_STUDY       = 3207 //远程控制遥控器学习
	NET_DVR_WIRELESS_ALARM_STUDY      = 3208 //远程控制无线报警学习
	NET_IPC_GET_AUX_ALARMCFG          = 3209 //获取辅助报警参数配置
	NET_IPC_SET_AUX_ALARMCFG          = 3210 //设置辅助报警参数配置
	NET_DVR_GET_PREVIEW_DISPLAYCFG    = 3211 //获取预览显示参数
	NET_DVR_SET_PREVIEW_DISPLAYCFG    = 3212 //设置预览显示参数
	NET_DVR_REMOTECONTROL_PTZ         = 3213 //远程控制PTZ
	NET_DVR_REMOTECONTROL_PRESETPOINT = 3214 //远程控制预置点
	NET_DVR_REMOTECONTROL_CRUISE      = 3215 //远程控制巡航

	NET_DVR_GET_MULTI_STREAM_COMPRESSIONCFG = 3216 //远程获取多码流压缩参数
	NET_DVR_SET_MULTI_STREAM_COMPRESSIONCFG = 3217 //远程设置多码流压缩参数

	NET_DVR_GET_WPSCFG        = 3218 //获取WPS参数
	NET_DVR_SET_WPSCFG        = 3219 //设置WPS参数
	NET_DVR_WPS_CONNECT       = 3220 //远程启用WPS连接
	NET_DVR_GET_DEVICE_PIN    = 3221 //获取设备PIN码
	NET_DVR_UPDATE_PIN        = 3223 //更新设备PIN码
	NET_DVR_GET_PRESETCFG     = 3224 //获取预置点参数
	NET_DVR_GET_PTZCRUISECFG  = 3225 //获取巡航路径参数
	NET_DVR_GET_PRESET_NUM    = 3226 //获取预置点个数
	NET_DVR_GET_PTZCRUISE_NUM = 3227 //获取巡航路径个数

	NET_DVR_GET_MOTION_TRACK_CFG = 3228 //获取跟踪参数
	NET_DVR_SET_MOTION_TRACK_CFG = 3229 //设置跟踪参数
	NET_DVR_CLEAR_IPC_PARAM      = 3230 //清空前端参数
	NET_DVR_GET_IPADDR_FILTERCFG = 3232 //获取IP地址过滤参数
	NET_DVR_SET_IPADDR_FILTERCFG = 3233 //设置IP地址过滤参数

	NET_DVR_GET_LOGO_OVERLAYCFG = 3234 //获取LOGO图片叠加参数
	NET_DVR_SET_LOGO_OVERLAYCFG = 3235 //设置LOGO图片叠加参数

	NET_DVR_GET_IPV6_LIST       = 3236 //获取网卡的全部IPV6地址信息
	NET_DVR_GET_AUDIOOUT_VOLUME = 3237 //获取输出音频大小
	NET_DVR_SET_AUDIOOUT_VOLUME = 3238 //设置输出音频大小
	NET_DVR_GET_FUZZY_UPGRADE   = 3239 //获取模糊升级匹配信息

	NET_DVR_GET_BV_CORRECT_PARAM  = 3240 //获取相机校正参数
	NET_DVR_SET_BV_CORRECT_PARAM  = 3241 //设置相机校正参数
	NET_DVR_GET_OUTPUT_VIDEO_TYPE = 3242 //获取输出视频类型
	NET_DVR_SET_OUTPUT_VIDEO_TYPE = 3243 //设置输出视频类型

	NET_DVR_FISHEYE_CFG               = 3244 //鱼眼长连接配置
	NET_DVR_GET_PTZ_POINT             = 3245 //获取PTZ点坐标
	NET_DVR_SET_PTZ_POINT             = 3246 //设置PTZ点坐标
	NET_DVR_REMOTECONTROL_DEV_PARAM   = 3247 //设置设备登录客户端参数
	NET_DVR_GET_FISHEYE_STREAM_STATUS = 3248 //获取鱼眼码流状态

	NET_DVR_GET_GBT28181_ACCESS_CFG   = 3249 //获取GBT28181协议接入配置
	NET_DVR_SET_GBT28181_ACCESS_CFG   = 3250 //设置GBT28181协议接入配置
	NET_DVR_GET_GBT28181_CHANINFO_CFG = 3251 //获取GBT28181协议接入设备的通道信息
	NET_DVR_SET_GBT28181_CHANINFO_CFG = 3252 //设置GBT28181协议接入设备的通道信息
	NET_DVR_GET_GBT28181_ALARMINCFG   = 3253 //获取GBT28181协议接入设备的报警信息
	NET_DVR_SET_GBT28181_ALARMINCFG   = 3254 //设置GBT28181协议接入设备的报警信息
	NET_DVR_GET_ISP_CAMERAPARAMCFG    = 3255 //获取ISP前端参数配置
	NET_DVR_SET_ISP_CAMERAPARAMCFG    = 3256 //设置ISP前端参数配置
	NET_DVR_GET_DEVSERVER_CFG         = 3257 //获取模块服务配置
	NET_DVR_SET_DEVSERVER_CFG         = 3258 //设置模块服务配置

	//2013-11-25
	NET_DVR_GET_WIPERINFO_CFG                 = 3259 //雨刷配置获取
	NET_DVR_SET_WIPERINFO_CFG                 = 3260 //雨刷配置设置
	NET_DVR_GET_TRACK_DEV_PARAM               = 3261 //获取跟踪设备参数
	NET_DVR_SET_TRACK_DEV_PARAM               = 3262 //设置跟踪设备参数
	NET_DVR_GET_PTZ_TRACK_PARAM               = 3263 //获取PTZ跟踪参数
	NET_DVR_SET_PTZ_TRACK_PARAM               = 3264 //设置PTZ跟踪参数
	NET_DVR_GET_CENTER_POINT_CFG              = 3265 //获取中心点参数
	NET_DVR_SET_CENTER_POINT_CFG              = 3266 //设置中心点参数
	NET_DVR_GET_CENTER_POINT_CFG_CAPABILITIES = 3267 //获取中心点参数能力
	NET_DVR_GET_FISHEYE_CAPABILITIES          = 3268 //获取鱼眼能力

	NET_DVR_GET_BASICPARAMCFG  = 3270 //获取PTZ配置基本参数信息
	NET_DVR_SET_BASICPARAMCFG  = 3271 //设置PTZ配置基本参数信息
	NET_DVR_GET_PTZOSDCFG      = 3272 //获取PTZ OSD配置参数信息
	NET_DVR_SET_PTZOSDCFG      = 3273 //设置PTZ OSD配置参数信息
	NET_DVR_GET_POWEROFFMEMCFG = 3274 //获取掉电记忆模式参数信息
	NET_DVR_SET_POWEROFFMEMCFG = 3275 //设置掉电记忆模式参数信息
	NET_DVR_GET_LIMITCFG       = 3276 //获取限位参数配置信息
	NET_DVR_SET_LIMITCFG       = 3277 //设置限位参数配置信息
	NET_DVR_PTZLIMIT_CTRL      = 3278 //清除限位参数控制

	NET_DVR_PTZ_CLEARCTRL               = 3279 //清除配置信息控制接口
	NET_DVR_GET_PRIORITIZECFG           = 3281 //获取云台优先配置信息
	NET_DVR_SET_PRIORITIZECFG           = 3282 //设置云台优先配置信息
	NET_DVR_PTZ_INITIALPOSITIONCTRL     = 3283 //零方位角控制
	NET_DVR_GET_PRIVACY_MASKSCFG        = 3285 //获取隐私遮蔽参数
	NET_DVR_SET_PRIVACY_MASKSCFG        = 3286 //设置隐私遮蔽参数
	NET_DVR_GET_PTZLOCKCFG              = 3287 //获取云台锁定信息
	NET_DVR_SET_PTZLOCKCFG              = 3288 //设置云台锁定信息
	NET_DVR_PTZ_ZOOMRATIOCTRL           = 3289 //设置跟踪倍率
	NET_DVR_GET_PTZLOCKINFO             = 3290 //获取云台锁定剩余秒数
	NET_DVR_GET_PRIVACY_MASKS_ENABLECFG = 3291 //获取全局使能
	NET_DVR_SET_PRIVACY_MASKS_ENABLECFG = 3292 //设置全局使能
	NET_DVR_GET_SMARTTRACKCFG           = 3293 //获取智能运动跟踪配置信息
	NET_DVR_SET_SMARTTRACKCFG           = 3294 //设置智能运动跟踪配置信息
	NET_DVR_GET_EPTZ_CFG                = 3295 //获取EPTZ参数
	NET_DVR_SET_EPTZ_CFG                = 3296 //设置EPTZ参数
	NET_DVR_GET_EPTZ_CFG_CAPABILITIES   = 3297 //获取EPTZ参数能力

	NET_DVR_GET_LOW_LIGHTCFG             = 3303 //获取快球低照度设置信息
	NET_DVR_SET_LOW_LIGHTCFG             = 3304 //设置快球低照度设置信息
	NET_DVR_GET_FOCUSMODECFG             = 3305 //获取快球聚焦模式信息
	NET_DVR_SET_FOCUSMODECFG             = 3306 //设置快球聚焦模式信息
	NET_DVR_GET_INFRARECFG               = 3307 //获取快球红外设置信息
	NET_DVR_SET_INFRARECFG               = 3308 //设置快球红外设置信息
	NET_DVR_GET_AEMODECFG                = 3309 //获取快球其他设置信息
	NET_DVR_SET_AEMODECFG                = 3310 //设置快球其他设置信息
	NET_DVR_CONTROL_RESTORE_SUPPORT      = 3311 //恢复前端默认参数(参数能力中有的前端参数配置相关的都恢复)
	NET_DVR_CONTROL_RESTART_SUPPORT      = 3312 //快球机芯重启
	NET_DVR_CONTROL_PTZ_PATTERN          = 3313 //云台花样扫描
	NET_DVR_GET_PTZ_PARKACTION_CFG       = 3314 //获取云台守望参数
	NET_DVR_SET_PTZ_PARKACTION_CFG       = 3315 //设置云台守望参数
	NET_DVR_CONTROL_PTZ_MANUALTRACE      = 3316 //手动跟踪定位
	NET_DVR_GET_ROI_DETECT_NUM           = 3349 //获取ROI检测区域编号数目
	NET_DVR_GET_ROI_DETECT               = 3350 //获取ROI检测区域配置
	NET_DVR_SET_ROI_DETECT               = 3351 //设置ROI检测区域配置
	NET_DVR_GET_FACE_DETECT              = 3352 //获取人脸侦测配置
	NET_DVR_SET_FACE_DETECT              = 3353 //设置人脸侦测配置
	NET_DVR_GET_CORRIDOR_MODE            = 3354 //获取走廊模式功能配置
	NET_DVR_SET_CORRIDOR_MODE            = 3355 //设置走廊模式功能配置
	NET_DVR_GET_SCENECHANGE_DETECTIONCFG = 3356 //获取场景变更报警配置
	NET_DVR_SET_SCENECHANGE_DETECTIONCFG = 3357 //设置场景变更报警配置
	NET_DVR_GET_TRAVERSE_PLANE_DETECTION = 3360
	NET_DVR_SET_TRAVERSE_PLANE_DETECTION = 3361
	NET_DVR_GET_FIELD_DETECTION          = 3362 //获取区域侦测配置
	NET_DVR_SET_FIELD_DETECTION          = 3363 //设置区域侦测配置
	NET_DVR_GET_DEFOCUSPARAM             = 3364 //获取虚焦侦测参数配置
	NET_DVR_SET_DEFOCUSPARAM             = 3365 //设置虚焦侦测参数配置
	NET_DVR_GET_AUDIOEXCEPTIONPARAM      = 3366 //获取音频异常配置
	NET_DVR_SET_AUDIOEXCEPTIONPARAM      = 3367 //设置音频异常配置
	NET_DVR_GET_CCDPARAMCFG_EX           = 3368 //获取CCD参数配置
	NET_DVR_SET_CCDPARAMCFG_EX           = 3369 //设置CCD参数配置
	NET_DVR_START_GET_INPUTVOLUME        = 3370 //开始获取音量
	NET_DVR_SET_SCH_TASK                 = 3380 //设置球机定时任务
	NET_DVR_GET_SCH_TASK                 = 3381 //获取球机定时任务
	NET_DVR_SET_PRESET_NAME              = 3382 //设置预置点名称
	NET_DVR_GET_PRESET_NAME              = 3383 //获取预置点名称
	NET_DVR_SET_AUDIO_NAME               = 3384 //设置语音名称
	NET_DVR_GET_AUDIO_NAME               = 3385 //获取语音名称
	NET_DVR_RESUME_INITRACKPOS           = 3386 //恢复跟踪初始位
	NET_DVR_NTP_SERVER_TEST              = 3387 //NTP服务器测试
	NET_DVR_NAS_SERVER_TEST              = 3388 //NAS服务器测试
	NET_DVR_EMAIL_SERVER_TEST            = 3389 //Email服务器测试
	NET_DVR_FTP_SERVER_TEST              = 3390 //FTP服务器测试
	NET_DVR_IP_TEST                      = 3391 //IP测试
	NET_DVR_GET_NET_DISKCFG_V40          = 3392 //网络硬盘接入获取v40
	NET_DVR_SET_NET_DISKCFG_V40          = 3393 //网络硬盘接入设置v40
	NET_DVR_GET_IOOUT_CFG                = 3394 //获取补光灯参数
	NET_DVR_SET_IOOUT_CFG                = 3395 //设置补光灯参数
	NET_DVR_GET_SIGNAL_SYNC              = 3396 //获取信号灯同步配置参数
	NET_DVR_SET_SIGNAL_SYNC              = 3397 //设置信号灯同步配置参数

	NET_DVR_GET_EZVIZ_ACCESS_CFG        = 3398 //获取EZVIZ接入参数
	NET_DVR_SET_EZVIZ_ACCESS_CFG        = 3399 //设置EZVIZ接入参数
	NET_DVR_GET_SCHEDULE_AUTO_TRACK_CFG = 3400 //获取定时智能跟踪参数
	NET_DVR_SET_SCHEDULE_AUTO_TRACK_CFG = 3401 //设置定时智能跟踪参数
	NET_DVR_MAKE_I_FRAME                = 3402 //强制I帧
	NET_DVR_GET_ALARM_RELATE            = 3403 //获取报警联动通道功能参数
	NET_DVR_SET_ALARM_RELATE            = 3404 //设置报警联动通道功能参数
	NET_DVR_GET_PDC_RULECFG_V42         = 3405 //设置人流量统计规则(扩展)
	NET_DVR_SET_PDC_RULECFG_V42         = 3406 //获取人流量统计规则(扩展)
	NET_DVR_GET_HEATMAP_CFG             = 3407 //设置热度图参数配置
	NET_DVR_SET_HEATMAP_CFG             = 3408 //获取热度图参数配置
	NET_DVR_REMOTECONTROL_LINEARSCAN    = 3409 //设置左右边界参数 2014-03-15
	NET_DVR_DPC_CTRL                    = 3410 //坏点校正控制
	NET_DVR_FFC_MANUAL_CTRL             = 3411 //非均匀性校正(FFC)手动模式
	NET_DVR_FFC_BACKCOMP_CTRL           = 3412 //非均匀性校正(FFC)背景补偿
	NET_DVR_GET_FOCUSING_POSITION_STATE = 3413 //获取聚焦到位状态参数
	NET_DVR_GET_PRIVATE_PROTOCOL_CFG    = 3414 //获取 私有关键信息上传配置接口配置
	NET_DVR_SET_PRIVATE_PROTOCOL_CFG    = 3415 //设置 私有关键信息上传配置接口配置
	NET_DVR_COMPLETE_RESTORE_CTRL       = 3420 //设置完全恢复出厂值

	NET_DVR_CLOUDSTORAGE_SERVER_TEST  = 3421 //云存储服务器测试
	NET_DVR_PHONE_NUM_TEST            = 3422 //电话号码测试
	NET_DVR_GET_REMOTECONTROL_STATUS  = 3423 //获取无线布防状态
	NET_DVR_GET_MONITOR_LOCATION_INFO = 3424 //获取监测点信息
	NET_DVR_SET_MONITOR_LOCATION_INFO = 3425 //设置监测点信息

	NET_DVR_GET_SMART_CAPABILITIES                      = 3500 //获取Smart能力
	NET_DVR_GET_EVENT_TRIGGERS_CAPABILITIES             = 3501 //获取事件触发能力
	NET_DVR_GET_REGION_ENTRANCE_CAPABILITIES            = 3502 //获取进入区域侦测能力
	NET_DVR_GET_REGION_ENTR_DETECTION                   = 3503 //获取进入区域配置
	NET_DVR_SET_REGION_ENTR_DETECTION                   = 3504 //设置进入区域配置
	NET_DVR_GET_REGION_ENTR_REGION                      = 3505 //获取进入区域的单个区域配置
	NET_DVR_SET_REGION_ENTR_REGION                      = 3506 //设置进入区域的单个区域配置
	NET_DVR_GET_REGION_ENTR_TRIGGER                     = 3507 //获取进入区域联动配置
	NET_DVR_SET_REGION_ENTR_TRIGGER                     = 3508 //设置进入区域联动配置
	NET_DVR_GET_REGION_ENTR_SCHEDULE                    = 3509 //获取进入区域布防时间配置
	NET_DVR_SET_REGION_ENTR_SCHEDULE                    = 3510 //设置进入区域布防时间配置
	NET_DVR_GET_REGION_EXITINT_CAPABILITIES             = 3511 //获取离开区域侦测能力
	NET_DVR_GET_REGION_EXITING_DETECTION                = 3512 //获取离开区域配置
	NET_DVR_SET_REGION_EXITING_DETECTION                = 3513 //设置离开区域配置
	NET_DVR_GET_REGION_EXITING_REGION                   = 3514 //获取离开区域的单个区域配置
	NET_DVR_SET_REGION_EXITING_REGION                   = 3515 //设置离开区域的单个区域配置
	NET_DVR_GET_REGION_EXIT_TRIGGER                     = 3516 //获取离开区域联动配置
	NET_DVR_SET_REGION_EXIT_TRIGGER                     = 3517 //设置离开区域联动配置
	NET_DVR_GET_REGION_EXIT_SCHEDULE                    = 3518 //获取离开区域布防时间配置
	NET_DVR_SET_REGION_EXIT_SCHEDULE                    = 3519 //设置离开区域布防时间配置
	NET_DVR_GET_LOITERING_CAPABILITIES                  = 3520 //获取徘徊侦测能力
	NET_DVR_GET_LOITERING_DETECTION                     = 3521 //获取徘徊侦测配置
	NET_DVR_SET_LOITERING_DETECTION                     = 3522 //设置徘徊侦测配置
	NET_DVR_GET_LOITERING_REGION                        = 3523 //获取徘徊的单个区域配置
	NET_DVR_SET_LOITERING_REGION                        = 3524 //设置徘徊的单个区域配置
	NET_DVR_GET_LOITERING_TRIGGER                       = 3525 //获取徘徊联动配置
	NET_DVR_SET_LOITERING_TRIGGER                       = 3526 //设置徘徊联动配置
	NET_DVR_GET_LOITERING_SCHEDULE                      = 3527 //获取徘徊布防时间配置
	NET_DVR_SET_LOITERING_SCHEDULE                      = 3528 //设置徘徊布防时间配置
	NET_DVR_GET_GROUPDETECTION_CAPABILITIES             = 3529 //获取人员聚集侦测能力
	NET_DVR_GET_GROUP_DETECTION                         = 3530 //获取人员聚集侦测配置
	NET_DVR_SET_GROUP_DETECTION                         = 3531 //设置人员聚集侦测配置
	NET_DVR_GET_GROUPDETECTION_REGION                   = 3532 //获取人员聚集的单个区域配置
	NET_DVR_SET_GROUPDETECTION_REGION                   = 3533 //设置人员聚集的单个区域配置
	NET_DVR_GET_GROUPDETECTION_TRIGGER                  = 3534 //获取人员聚集联动配置
	NET_DVR_SET_GROUPDETECTION_TRIGGER                  = 3535 //设置人员聚集联动配置
	NET_DVR_GET_GROUPDETECTION_SCHEDULE                 = 3536 //获取人员聚集布防时间配置
	NET_DVR_SET_GROUPDETECTION_SCHEDULE                 = 3537 //设置人员聚集布防时间配置
	NET_DVR_GET_RAPIDMOVE_CAPABILITIES                  = 3538 //获取快速运动侦测能力
	NET_DVR_GET_RAPIDMOVE_DETECTION                     = 3539 //获取快速运动侦测配置
	NET_DVR_SET_RAPIDMOVE_DETECTION                     = 3540 //设置快速运动侦测配置
	NET_DVR_GET_RAPIDMOVE_REGION                        = 3541 //获取快速运动的单个区域配置
	NET_DVR_SET_RAPIDMOVE_REGION                        = 3542 //设置快速运动的单个区域配置
	NET_DVR_GET_RAPIDMOVE_TRIGGER                       = 3543 //获取快速运动联动配置
	NET_DVR_SET_RAPIDMOVE_TRIGGER                       = 3544 //设置快速运动联动配置
	NET_DVR_GET_RAPIDMOVE_SCHEDULE                      = 3545 //获取快速运动的布防时间配置
	NET_DVR_SET_RAPIDMOVE_SCHEDULE                      = 3546 //设置快速运动的布防时间配置
	NET_DVR_GET_PATKING_CAPABILITIES                    = 3547 //获取停车侦测能力
	NET_DVR_GET_PARKING_DETECTION                       = 3548 //获取停车侦测配置
	NET_DVR_SET_PARKING_DETECTION                       = 3549 //设置停车侦测配置
	NET_DVR_GET_PARKING_REGION                          = 3550 //获取停车侦测的单个区域配置
	NET_DVR_SET_PARKING_REGION                          = 3551 //设置停车侦测的单个区域配置
	NET_DVR_GET_PARKING_TRIGGER                         = 3552 //获取停车侦测联动配置
	NET_DVR_SET_PARKING_TRIGGER                         = 3553 //设置停车侦测联动配置
	NET_DVR_GET_PARKING_SCHEDULE                        = 3554 //获取停车侦测的布防时间配置
	NET_DVR_SET_PARKING_SCHEDULE                        = 3555 //设置停车侦测的布防时间配置
	NET_DVR_GET_UNATTENDED_BAGGAGE_CAPABILITIES         = 3556 //获取物品遗留侦测能力
	NET_DVR_GET_UNATTENDED_BAGGAGE_DETECTION            = 3557 //获取物品遗留侦测配置
	NET_DVR_SET_UNATTENDED_BAGGAGE_DETECTION            = 3558 //设置物品遗留侦测配置
	NET_DVR_GET_UNATTENDED_BAGGAGE_REGION               = 3559 //获取物品遗留侦测的单个区域配置
	NET_DVR_SET_UNATTENDED_BAGGAGE_REGION               = 3560 //设置物品遗留侦测的单个区域配置
	NET_DVR_GET_UNATTENDED_BAGGAGE_TRIGGER              = 3561 //获取物品遗留侦测联动配置
	NET_DVR_SET_UNATTENDED_BAGGAGE_TRIGGER              = 3562 //设置物品遗留侦测联动配置
	NET_DVR_GET_UNATTENDED_BAGGAGE_SCHEDULE             = 3563 //获取物品遗留侦测的布防时间配置
	NET_DVR_SET_UNATTENDED_BAGGAGE_SCHEDULE             = 3564 //设置物品遗留侦测的布防时间配置
	NET_DVR_GET_ATTENDEDBAGGAGE_CAPABILITIES            = 3565 //获取物品拿取侦测能力
	NET_DVR_GET_ATTENDEDBAGGAGE_DETECTION               = 3566 //获取物品拿取侦测配置
	NET_DVR_SET_ATTENDEDBAGGAGE_DETECTION               = 3567 //设置物品拿取侦测配置
	NET_DVR_GET_ATTENDEDBAGGAGE_REGION                  = 3568 //获取物品拿取侦测的单个区域配置
	NET_DVR_SET_ATTENDEDBAGGAGE_REGION                  = 3569 //设置物品拿取侦测的单个区域配置
	NET_DVR_GET_ATTENDEDBAGGAGE_TRIGGER                 = 3570 //获取物品拿取侦测联动配置
	NET_DVR_SET_ATTENDEDBAGGAGE_TRIGGER                 = 3571 //设置物品拿取侦测联动配置
	NET_DVR_GET_ATTENDEDBAGGAGE_SCHEDULE                = 3572 //获取物品遗留侦测的布防时间配置
	NET_DVR_SET_ATTENDEDBAGGAGE_SCHEDULE                = 3573 //设置物品拿取侦测的布防时间配置
	NET_DVR_GET_REGIONCLIP_CAPABILITIES                 = 3574 //获取区域裁剪能力
	NET_DVR_GET_REGION_CLIP                             = 3575 //获取区域裁剪配置
	NET_DVR_SET_REGION_CLIP                             = 3576 //设置区域裁剪配置
	NET_DVR_GET_NETWORK_CAPABILITIES                    = 3577 //获取网络能力
	NET_DVR_GET_WIRELESS_DIAL                           = 3578 //获取无线参数配置
	NET_DVR_SET_WIRELESS_DIAL                           = 3579 //设置无线参数配置
	NET_DVR_GET_WIRELESSDIAL_CAPABILITIES               = 3580 //获取无线拨号参数能力
	NET_DVR_GET_WIRELESSDIAL_SCHEDULE                   = 3581 //获取拨号计划配置
	NET_DVR_SET_WIRELESSDIAL_SCHEDULE                   = 3582 //设置拨号计划配置
	NET_DVR_GET_WIRELESSDIAL_STATUS                     = 3583 //获取拨号状态
	NET_DVR_GET_REGION_ENTRANCE_SCHEDULE_CAPABILITIES   = 3584 //获取进入区域侦测布防时间能力
	NET_DVR_GET_REGION_EXITING_SCHEDULE_CAPABILITIES    = 3585 //获取离开区域侦测布防时间能力
	NET_DVR_GET_LOITERING_SCHEDULE_CAPABILITIES         = 3586 //获取徘徊侦测布防时间能力
	NET_DVR_GET_GROUP_SCHEDULE_CAPABILITIES             = 3587 //获取人员聚集侦测布防时间能力
	NET_DVR_GET_RAPIDMOVE_SCHEDULE_CAPABILITIES         = 3588 //获取快速运动侦测布防时间能力
	NET_DVR_GET_PARKING_SCHEDULE_CAPABILITIES           = 3589 //获取停车侦测布防时间能力
	NET_DVR_GET_UNATTENDEDBAGGAGE_SCHEDULE_CAPABILITIES = 3590 //获取物品遗留侦测布防时间能力
	NET_DVR_GET_ATTENDEDBAGGAGE_SCHEDULE_CAPABILITIES   = 3591 //获取物品拿取侦测布防时间能力
	NET_DVR_GET_WIRELESSDIAL_SCHEDULE_CAPABILITIES      = 3592 //获取拨号计划能力
	NET_DVR_WIRELESSDIAL_CONNECT                        = 3593 //控制无线网络连网断网

	NET_DVR_GET_LITESTORAGE              = 3594 //获取轻存储配置
	NET_DVR_SET_LITESTORAGE              = 3595 //设置轻存储配置
	NET_DVR_GET_LITESTORAGE_CAPABILITIES = 3596 //获取轻存储能力

	NET_DVR_GET_VEHICLE_CAPABILITIES = 3597 //获取车俩检测标定能力
	NET_DVR_GET_VEHICLE_CALIBRATION  = 3598 //获取车辆检测标定

	NET_DVR_GET_SLAVECAMERA_CAPABILITIES         = 3599 //获取从摄像机IP信息配置能力
	NET_DVR_GET_SLAVECAMERA                      = 3600 //获取从摄像机IP信息配置
	NET_DVR_SET_SLAVECAMERA                      = 3601 //设置从摄像机IP信息配置
	NET_DVR_GET_SLAVECAMERA_STATUS               = 3602 //获取从摄像机连接状态
	NET_DVR_GET_SLAVECAMERA_CALIB_CAPABILITIES   = 3603 //获取从摄像机配置&&标定能力
	NET_DVR_GET_SLAVECAMERA_CALIB                = 3604 //获取从摄像机标定配置
	NET_DVR_SET_SLAVECAMERA_CALIB                = 3605 //设置从摄像机标定配置
	NET_DVR_GET_PHY_RATIO                        = 3606 //获取物理倍率坐标信息
	NET_DVR_SET_PHY_RATIO                        = 3607 //设置物理倍率坐标信息
	NET_DVR_GET_MASTERSLAVETRACKING_CAPABILITIES = 3608 //获取主从跟踪能力
	NET_DVR_SET_TRACKINGRATIO                    = 3610 //设置从摄像机跟踪倍率
	NET_DVR_GET_TRACKING                         = 3611 //获取主从跟踪功能相机跟踪配置
	NET_DVR_SET_TRACKING                         = 3612 //设置主从跟踪功能相机跟踪配置
	NET_DVR_GET_TRACKING_CAPABILITIES            = 3613 //获取主从跟踪功能相机跟踪配置能力

	NET_DVR_GET_SLAVECAMERA_CALIB_V50 = 3614 //获取从摄像机标定配置V50
	NET_DVR_SET_SLAVECAMERA_CALIB_V50 = 3615 //设置从摄像机标定配置V50
	NET_DVR_SET_TRACKINGRATIO_MANUAL  = 3616 //设置从摄像机手动跟踪倍率
	NET_DVR_GET_TRACKINGRATIO_MANUAL  = 3617 //获取从摄像机手动跟踪倍率
	NET_DVR_SET_TRACK_INITPOSTION     = 3618 //设置从摄像机初始跟踪位置
	NET_DVR_GET_PTZ_CAPABILITIES      = 3619 //获取ptz球机控制能力

	NET_DVR_GET_THERMOMETRY_BASICPARAM_CAPABILITIES = 3620 //获取测温配置能力
	NET_DVR_GET_THERMOMETRY_BASICPARAM              = 3621 //获取测温配置参数
	NET_DVR_SET_THERMOMETRY_BASICPARAM              = 3622 //设置测温配置参数
	NET_DVR_GET_THERMOMETRY_SCENE_CAPABILITIES      = 3623 //获取测温预置点关联配置能力
	NET_DVR_GET_THERMOMETRY_PRESETINFO              = 3624 //获取测温预置点关联配置参数
	NET_DVR_SET_THERMOMETRY_PRESETINFO              = 3625 //设置测温预置点关联配置参数
	NET_DVR_GET_THERMOMETRY_ALARMRULE_CAPABILITIES  = 3626 //获取测温报警方式配置能力
	NET_DVR_GET_THERMOMETRY_ALARMRULE               = 3627 //获取测温预置点报警规则配置参数
	NET_DVR_SET_THERMOMETRY_ALARMRULE               = 3628 //设置测温预置点报警规则配置参数
	NET_DVR_GET_REALTIME_THERMOMETRY                = 3629 //实时温度检测
	NET_DVR_GET_THERMOMETRY_DIFFCOMPARISON          = 3630 //获取测温预置点温差规则配置参数
	NET_DVR_SET_THERMOMETRY_DIFFCOMPARISON          = 3631 //设置测温预置点温差规则配置参数
	NET_DVR_GET_THERMOMETRY_TRIGGER                 = 3632 //获取测温联动配置
	NET_DVR_SET_THERMOMETRY_TRIGGER                 = 3633 //设置测温联动配置

	NET_DVR_GET_THERMAL_CAPABILITIES       = 3634 //获取热成像（Thermal）能力
	NET_DVR_GET_FIREDETECTION_CAPABILITIES = 3635 //获取火点检测配置能力
	NET_DVR_GET_FIREDETECTION              = 3636 //获取火点检测参数
	NET_DVR_SET_FIREDETECTION              = 3637 //设置火点检测参数
	NET_DVR_GET_FIREDETECTION_TRIGGER      = 3638 //获取火点检测联动配置
	NET_DVR_SET_FIREDETECTION_TRIGGER      = 3639 //设置火点检测联动配置

	NET_DVR_GET_OIS_CAPABILITIES               = 3640 //获取光学防抖参数配置能力
	NET_DVR_GET_OIS_CFG                        = 3641 //获取光学防抖配置
	NET_DVR_SET_OIS_CFG                        = 3642 //设置光学防抖配置
	NET_DVR_GET_MACFILTER_CAPABILITIES         = 3643 //获取MAC地址过滤配置能力
	NET_DVR_GET_MACFILTER_CFG                  = 3644 //获取MAC地址过滤配置
	NET_DVR_SET_MACFILTER_CFG                  = 3645 //设置MAC地址过滤配置
	NET_DVR_GET_EAGLEFOCUS_CALCFG_CAPABILITIES = 3646 //鹰视聚焦标定配置能力
	NET_DVR_GET_EAGLEFOCUSING_CALCFG           = 3647 //获取鹰视聚焦标定配置
	NET_DVR_SET_EAGLEFOCUSING_CALCFG           = 3648 //设置鹰视聚焦标定配置
	NET_DVR_GET_EAGLEFOCUSING_CFG_CAPABILITIES = 3649 //获取鹰视聚焦配置能力
	NET_DVR_GET_EAGLEFOCUSING_CTRL             = 3650 //获取鹰视聚焦配置
	NET_DVR_SET_EAGLEFOCUSING_CTRL             = 3651 //设置鹰视聚焦配置

	NET_DVR_GET_PXOFFLINE_CAPABILITIES            = 3652 //获取停车场票箱脱机下参数配置 能力
	NET_DVR_SET_PXOFFLINE_CFG                     = 3653 //设置停车场票箱脱机下参数配置信息
	NET_DVR_GET_PXOFFLINE_CFG                     = 3654 //获取停车场票箱脱机下参数配置信息
	NET_DVR_GET_PAPERCHARGEINFO_CAPABILITIES      = 3655 //获取停车场出入口纸票信息下发 能力
	NET_DVR_SET_PAPERCHARGEINFO                   = 3656 //设置停车场出入口纸票信息下发
	NET_DVR_GET_PARKINGSAPCE_CAPABILITIES         = 3657 //获取停车场出入口停车位信息下发 能力
	NET_DVR_SET_PARKINGSAPCE_INFO                 = 3658 //设置停车场出入口停车位信息下发
	NET_DVR_GET_PXMULTICTRL_CAPABILITIES          = 3659 //获取停车场票箱从属设备多角度参数配置 能力
	NET_DVR_GET_CHARGEACCOUNT_CAPABILITIES        = 3661 //获取停车场票箱参数配置能力
	NET_DVR_SET_CHARGE_ACCOUNTINFO                = 3662 //设置缴费金额信息
	NET_DVR_SET_PXMULTICTRL_CFG                   = 3663 //设置停车场票箱从属设备多角度参数配置信息
	NET_DVR_GET_PXMULTICTRL_CFG                   = 3664 //获取停车场票箱从属设备多角度参数配置信息
	NET_DVR_GET_TME_CHARGERULE                    = 3665 //获取停车场出入口车卡收费规则规则
	NET_DVR_SET_TME_CHARGERULE                    = 3666 //设置停车场出入口车卡收费规则规则
	NET_DVR_GET_TME_CHARGERULE_CAPABILITIES       = 3667 //获取停车场出入口 车卡收费信息配置能力
	NET_DVR_GET_ILLEGALCARDFILTERING_CAPABILITIES = 3668 //获取停车场票箱参数配置能力
	NET_DVR_GET_ILLEGALCARDFILTERING_CFG          = 3669 //获取停车场票箱参数配置
	NET_DVR_SET_ILLEGALCARDFILTERING_CFG          = 3670 //设置停车场票箱参数配置
	NET_DVR_GET_LEDDISPLAY_CAPABILITIES           = 3671 //获取LED屏幕显示参数配置参数能力
	NET_DVR_SET_LEDDISPLAY_CFG                    = 3672 //设置LED屏幕显示参数
	NET_DVR_GET_LEDDISPLAY_CFG                    = 3673 //获取LED屏幕显示参数
	NET_DVR_GET_VOICEBROADCAST_CAPABILITIES       = 3674 //获取语音播报控制参数配置参数能力
	NET_DVR_SET_VOICEBROADCAST_CFG                = 3675 //设置语音播报控制参数
	NET_DVR_GET_PAPERPRINTFORMAT_CAPABILITIES     = 3676 //获取纸票打印格式配置能力
	NET_DVR_GET_PAPERPRINTFORMAT_CFG              = 3677 //获取纸票打印格式参数配置
	NET_DVR_SET_PAPERPRINTFORMAT_CFG              = 3678 //设置纸票打印格式参数配置
	NET_DVR_GET_LOCkGATE_CAPABILITIES             = 3679 //获取智能锁闸配置能力
	NET_DVR_GET_LOCKGATE_CFG                      = 3680 //获取智能锁闸参数配置
	NET_DVR_SET_LOCKGATE_CFG                      = 3681 //设置智能锁闸参数配置
	NET_DVR_GET_PARKING_DATASTATE                 = 3682 //获取数据同步状态
	NET_DVR_SET_PARKING_DATASTATE                 = 3683 //设置数据同步状态
	NET_DVR_GET_TME_CAPABILITIES                  = 3684 //获取停车场出入口设备 能力

	NET_DVR_GET_TMEVOICE_CAPABILITIES        = 3686 //获取语音配置信息能力
	NET_DVR_SET_TMEVOICE_CFG                 = 3687 //设置语音参数配置
	NET_DVR_GET_TMEVOICE_CFG                 = 3688 //获取语音参数配置
	NET_DVR_DEL_TMEVOICE_CFG                 = 3689 //删除语音参数配置
	NET_DVR_GET_POSITION                     = 3698 // 获取方位矫正配置参数
	NET_DVR_SET_POSITION                     = 3699 // 设置方位矫正配置参数
	NET_DVR_GET_CENTRALIZEDCTRL_CAPABILITIES = 3700 //获取集中布控能力
	NET_DVR_GET_CENTRALIZEDCTRL              = 3701 //获取集中布控参数配置
	NET_DVR_SET_CENTRALIZEDCTRL              = 3702 //设置集中布控参数配置
	NET_DVR_GET_COMPASS_CAPABILITIES         = 3703 //获取电子罗盘能力
	NET_DVR_GET_VANDALPROOFALARM             = 3704 //获取防破坏报警参数配置
	NET_DVR_SET_VANDALPROOFALARM             = 3705 //设置防破坏报警参数配置
	NET_DVR_COMPASS_CALIBRATE_CTRL           = 3706 //电子罗盘矫正控制接口
	NET_DVR_COMPASS_NORTH_CTRL               = 3707 //电子罗盘指向正北控制接口
	NET_DVR_GET_AZIMUTHINFO                  = 3708 //获取方位角度参数配置

	NET_DVR_GET_SATELLITETIME                      = 3709 //获取卫星定位参数配置
	NET_DVR_SET_SATELLITETIME                      = 3710 //设置卫星定位参数配置
	NET_DVR_GET_GISINFO                            = 3711 //获取当前球机的GIS信息数据
	NET_DVR_GET_STREAMING_CAPABILITIES             = 3712 //获取视频流的能力
	NET_DVR_GET_REFRESHFRAME_CAPABILITIES          = 3713 //获取刷新帧的能力
	NET_DVR_STREAMING_REFRESH_FRAME                = 3714 //取流预览的强制刷新帧
	NET_DVR_FACECAPTURE_STATISTICS                 = 3715 //长连接人员统计
	NET_DVR_GET_WIRELESSSERVER_CAPABILITIES        = 3716 //获取热点功能配置协议的能力
	NET_DVR_GET_WIRELESSSERVER                     = 3717 //获取热点功能配置协议
	NET_DVR_SET_WIRELESSSERVER                     = 3718 //设置热点功能配置协议
	NET_DVR_GET_CONNECT_LIST_CAPABILITIES          = 3719 //获取连接设备列表信息的能力
	NET_DVR_GET_THSCREEN_CAPABILITIES              = 3720 //获取温湿度配置协议的能力
	NET_DVR_GET_THSCREEN                           = 3721 //获取温湿度配置协议
	NET_DVR_GET_EXTERNALDEVICE_CAPABILITIES        = 3722 //获取外设配置协议的能力
	NET_DVR_GET_EXTERNALDEVICE                     = 3723 //获取外设配置协议
	NET_DVR_SET_EXTERNALDEVICE                     = 3724 //设置外设配置协议
	NET_DVR_GET_LEDDISPLAYINFO_CAPABILITIES        = 3725 //获取LED显示信息的能力
	NET_DVR_SET_LEDDISPLAYINFO                     = 3726 //设置LED显示信息
	NET_DVR_GET_SUPPLEMENTLIGHT_CAPABILITIES       = 3727 //获取内置补光灯配置协议的能力 (球机支持，软件实现，补光灯是设计在设备内部的)
	NET_DVR_GET_SUPPLEMENTLIGHT                    = 3728 //获取内置补光灯配置协议
	NET_DVR_SET_SUPPLEMENTLIGHT                    = 3729 //设置内置补光灯配置协议
	NET_DVR_SET_THSCREEN                           = 3730 //设置温湿度配置协议
	NET_DVR_GET_LOWPOWER_CAPABILITIES              = 3731 //获取低功耗配置协议的能力
	NET_DVR_GET_LOWPOWER                           = 3732 //获取低功耗配置协议
	NET_DVR_SET_LOWPOWER                           = 3733 //设置低功耗配置协议
	NET_DVR_GET_ZOOMLINKAGE_CAPABILITIES           = 3734 //获取变倍联动配置协议的能力
	NET_DVR_GET_ZOOMLINKAGE                        = 3735 //获取变倍联动配置协议
	NET_DVR_SET_ZOOMLINKAGE                        = 3736 //设置变倍联动配置协议
	NET_DVR_THSCREEN_TIMING                        = 3737 //温湿度
	NET_DVR_GET_OSD_BATTERY_POWER_CFG              = 3741 //获取OSD电池电量显示参数
	NET_DVR_SET_OSD_BATTERY_POWER_CFG              = 3742 //设置OSD电池电量显示参数
	NET_DVR_GET_OSD_BATTERY_POWER_CFG_CAPABILITIES = 3743 //OSD电池电量显示参数的能力
	NET_DVR_GET_VANDALPROOFALARM_TRIGGER           = 3744 //获取防破坏报警联动配置
	NET_DVR_SET_VANDALPROOFALARM_TRIGGER           = 3745 //设置防破坏报警联动配置
	NET_DVR_GET_PANORAMAIMAGE_CAPABILITIES         = 3746 //获取全景图像的能力
	NET_DVR_GET_PANORAMAIMAGE                      = 3747 //获取全景图像参数的协议
	NET_DVR_SET_PANORAMAIMAGE                      = 3748 //设置全景图像参数的协议
	NET_DVR_GET_STREAMENCRYPTION                   = 3749 //获取码流加密配置
	NET_DVR_SET_STREAMENCRYPTION                   = 3750 //设置码流加密配置
	NET_DVR_GET_STREAMENCRYPTION_CAPABILITIES      = 3751 //获取码流加密能力
	NET_DVR_GET_REVISE_GPS_CAPABILITIES            = 3752 //获取校准GPS经纬度能力
	NET_DVR_GET_REVISE_GPS                         = 3753 //获取校准GPS经纬度能力
	NET_DVR_SET_REVISE_GPS                         = 3754 //设置校准GPS经纬度能力
	NET_DVR_GET_PDC_RECOMMEND                      = 3755 //获取客流统计表示推荐值
	NET_DVR_REMOVE_FLASHSTORAGE                    = 3756 //客流数据清除操作
	NET_DVR_GET_COUNTING_CAPABILITIES              = 3757 //获取客流量统计能力
	NET_DVR_SET_SENSOR_ADJUSTMENT                  = 3758 //设置Sensor 调节参数的协议
	NET_DVR_GET_SENSOR_ADJUSTMENT_CAPABILITIES     = 3759 //获取Sensor 调节参数的协议的能力
	NET_DVR_GET_WIRELESSSERVER_FULLVERSION_CFG     = 3760 //获取wifi热点参数配置(完整版)
	NET_DVR_SET_WIRELESSSERVER_FULLVERSION_CFG     = 3761 //设置wifi热点参数配置(完整版)
	NET_DVR_GET_ONLINEUSER_INFO                    = 3762 //长连接获取用户在线信息
	NET_DVR_GET_SENSOR_ADJUSTMENT_INFO             = 3763 //获取指定sensor调节参数
	NET_DVR_SENSOR_RESET_CTRL                      = 3764 //Sensor 调节复位

	NET_DVR_GET_POSTRADAR_CAPABILITIES   = 3765 //获取雷达测速配置能力
	NET_DVR_GET_POSTRADARSPEED_CFG       = 3766 //获取雷达测速配置
	NET_DVR_SET_POSTRADARSPEED_CFG       = 3767 //设置雷达测速配置
	NET_DVR_GET_POSTRADARSPEED_RECOM_CFG = 3768 //获取雷达测速推荐值
	NET_DVR_GET_POSTRADARPARAM_CFG       = 3769 //获取雷达参数配置
	NET_DVR_SET_POSTRADARPARAM_CFG       = 3770 //设置雷达参数配置
	NET_DVR_GET_POSTRADARPARAM_RECOM_CFG = 3771 //获取雷达参数推荐值

	NET_DVR_GET_ENCRYPT_DEVICE_INFO = 3772 //获取加密设备信息
	NET_DVR_GET_ANR_ARMING_HOST     = 3773 //获取断网续传的主机信息
	NET_DVR_GET_FIRMWARE_VERSION    = 3776 //GET firmware version
	/********************************IPC基线FF车牌****************************/
	NET_DVR_GET_FTP_CAPABILITIES              = 3782 //获取ftp能力
	NET_DVR_GET_FTPUPLOAD_CFG                 = 3783 //获取ftp上传信息规整参数
	NET_DVR_SET_FTPUPLOAD_CFG                 = 3784 //设置ftp上传信息规整参数
	NET_DVR_GET_VEHICLE_INFORMATION           = 3785 //获取车辆信息
	NET_DVR_GET_DDNS_COUNTRY_ABILITY          = 3800 //获取设备支持的DDNS国家能力列表
	NET_DVR_GET_DEVICECFG_V50                 = 3801 //获取设备参数
	NET_DVR_SET_DEVICECFG_V50                 = 3802 //设置设备参数
	NET_DVR_SET_VEHICLE_RECOG_TASK_V50        = 3851 //车辆二次识别任务提交V50扩展
	NET_DVR_GET_SMARTCALIBRATION_CAPABILITIES = 3900 // Smart行为标定过滤尺寸功能能力
	NET_DVR_GET_TEMPERATURE_TRIGGER           = 3903 //获取测温差联动配置
	NET_DVR_SET_TEMPERATURE_TRIGGER           = 3904 //设置测温差联动配置

	NET_DVR_GET_SMARTCALIBRATION_CFG                = 3910 //获取Smart行为标定过滤尺寸功能
	NET_DVR_SET_SMARTCALIBRATION_CFG                = 3911 //设置Smart行为标定过滤尺寸功能
	NET_DVR_POST_SETUP_CALIB                        = 3912 //架设标定
	NET_DVR_SET_POS_INFO_OVERLAY                    = 3913 //设置Pos信息码流叠加控制
	NET_DVR_GET_POS_INFO_OVERLAY                    = 3914 //获取Pos信息码流叠加控制
	NET_DVR_GET_CAMERA_WORK_MODE                    = 3915 //设置相机工作模式参数
	NET_DVR_SET_CAMERA_WORK_MODE                    = 3916 //获取相机工作模式参数
	NET_DVR_GET_RESOLUTION_SWITCH_CAPABILITIES      = 3917 //获取分辨率模式切换能力
	NET_DVR_GET_RESOLUTION_SWITCH                   = 3918 //获取分辨率模式切换配置
	NET_DVR_SET_RESOLUTION_SWITCH                   = 3919 //设置分辨率模式切换配置
	NET_DVR_GET_CONFIRM_MECHANISM_CAPABILITIES      = 3920 //报警上传确认机制控制能力
	NET_DVR_CONFIRM_MECHANISM_CTRL                  = 3921 //报警上传确认机制控制
	NET_DVR_GET_VEHICLLE_RESULT_CAPABILITIES        = 3951 //获取获取车辆信息结果能力
	NET_DVR_GET_CALIB_CAPABILITIES                  = 3952 //获取架设标定能力
	NET_DVR_GET_POSINFO_OVERLAY_CAPABILITIES        = 3953 //获取获取Pos叠加能力
	NET_SDK_FINDMEDICALFILE                         = 3954 //慧影科技智慧医疗查找录像文件
	NET_SDK_FINDMEDICALPICTURE                      = 3955 //慧影科技智慧医疗查找图片文件
	NET_DVR_SET_POSINFO_OVERLAY                     = 3960 //设置Pos叠加
	NET_DVR_GET_POSINFO_OVERLAY                     = 3961 //获取Pos叠加
	NET_DVR_GET_FACELIB_TRIGGER                     = 3962 //获取人脸比对库的联动配置
	NET_DVR_SET_FACELIB_TRIGGER                     = 3963 //设置人脸比对库的联动配置
	NET_DVR_GET_FACECONTRAST_TRIGGER                = 3965 //获取人脸比对联动配置
	NET_DVR_SET_FACECONTRAST_TRIGGER                = 3966 //设置人脸比对联动配置
	NET_DVR_GET_FACECONTRAST_SCHEDULE_CAPABILITIES  = 3967 //获取人脸比对布防时间能力
	NET_DVR_GET_FACECONTRAST_SCHEDULE               = 3968 //获取人脸比对布防时间配置
	NET_DVR_SET_FACECONTRAST_SCHEDULE               = 3969 //设置人脸比对布防时间配置
	NET_DVR_GET_FACELIB_SCHEDULE_CAPABILITIES       = 3970 //获取人脸比对库的布防时间能力
	NET_DVR_GET_VCA_VERSION_LIST                    = 3973 //获取算法库版本
	NET_DVR_GET_SETUP_CALIB                         = 3974 //获取架设标定
	NET_DVR_GET_PANORAMA_LINKAGE                    = 3975 //获取联动抓图上传使能配置
	NET_DVR_SET_PANORAMA_LINKAGE                    = 3976 //设置联动抓图上传使能配置
	NET_DVR_GET_FACELIB_SCHEDULE                    = 3977 //获取人脸比对库的布防时间配置
	NET_DVR_SET_FACELIB_SCHEDULE                    = 3978 //设置人脸比对库的布防时间配置
	NET_DVR_GET_SOFTWARE_SERVICE_CAPABILITIES       = 3980 //获取软件服务能力
	NET_DVR_GET_SOFTWARE_SERVICE                    = 3981 //获取软件服务配置
	NET_DVR_SET_SOFTWARE_SERVICE                    = 3982 //设置软件服务配置
	NET_DVR_GET_PREVIEW_MODE_CAPABILITIES           = 3983 //获取预览模式配置能力
	NET_DVR_SET_EAGLE_FOCUS_GOTOSCENE               = 3984 //鹰式聚焦设置摄像机转向指定的场景ID
	NET_DVR_EAGLE_FOCUS_SCENE_DEL                   = 3985 //删除鹰式聚焦标定的场景
	NET_DVR_GET_SAFETY_HELMET_TRIGGER               = 3986 //获取安全帽检测联动配置
	NET_DVR_SET_SAFETY_HELMET_TRIGGER               = 3987 //设置安全帽检测联动配置
	NET_DVR_GET_SAFETY_HELMET_SCHEDULE_CAPABILITIES = 3988 //获取安全帽检测布防时间能力
	NET_DVR_GET_SAFETY_HELMET_SCHEDULE              = 3989 //获取安全帽检测布防时间配置
	NET_DVR_SET_SAFETY_HELMET_SCHEDULE              = 3990 //设置安全帽检测布防时间配置

	NET_DVR_GET_SIGN_ABNORMAL_TRIGGER = 4150 //获取体征异常联动配置
	NET_DVR_SET_SIGN_ABNORMAL_TRIGGER = 4151 //设置体征异常联动配置

	NET_DVR_ONE_KEY_CONFIG_SAN_V50             = 4152 //一键配置SAN(V50)
	NET_DVR_GET_HDCFG_V50                      = 4153 //获取硬盘信息参数V50
	NET_DVR_SET_HDCFG_V50                      = 4154 //设置硬盘信息参数V50
	NET_DVR_GET_HDVOLUME_CFG                   = 4155 //获取硬盘卷信息
	NET_DVR_SET_HDVOLUME_CFG                   = 4156 //设置硬盘卷信息
	NET_DVR_GET_POWER_SUPPLY_CABINET_TRIGGER   = 4157 //获取机柜供电检测的联动配置
	NET_DVR_SET_POWER_SUPPLY_CABINET_TRIGGER   = 4158 //设置机柜供电检测的联动配置
	NET_DVR_GET_SENSOR_TRIGGER                 = 4159 //获取传感器检测的联动配置
	NET_DVR_SET_SENSOR_TRIGGER                 = 4160 //设置传感器检测的联动配置
	NET_DVR_GET_FACESNAP_TRIGGER               = 4161 //获取人脸抓拍联动配置
	NET_DVR_SET_FACESNAP_TRIGGER               = 4162 //设置人脸抓拍联动配置
	NET_DVR_GET_FACESNAP_SCHEDULE_CAPABILITIES = 4163 //获取人脸抓拍布防时间能力
	NET_DVR_GET_FACESNAP_SCHEDULE              = 4164 //获取人脸抓拍布防时间配置
	NET_DVR_SET_FACESNAP_SCHEDULE              = 4165 //设置人脸抓拍布防时间配置

	NET_DVR_SET_SCREEN_SWITCH                = 4171 //画面切换控制
	NET_DVR_GET_BV_CALIB_PIC                 = 4172 //获取设备抓取图片和附加信息
	NET_DVR_GET_BV_CALIB_RESULT              = 4173 //获取双目外参标定结果
	NET_DVR_GET_BV_HCORRECTION               = 4174 //获取双目高度矫正数据
	NET_DVR_DEL_BV_CALIB_PIC                 = 4175 //删除样本图片
	NET_DVR_GET_TV_SCREEN_CFG                = 4176 //获取导播画面停留时间配置
	NET_DVR_SET_TV_SCREEN_CFG                = 4177 //设置导播画面停留时间配置
	NET_DVR_ADJUST_BV_CALIB                  = 4178 //双目标定微调
	NET_DVR_GET_HUMAN_CALIB                  = 4179 //获取人体坐标标定配置
	NET_DVR_SET_HUMAN_CALIB                  = 4180 //设置人体坐标标定配置
	NET_DVR_GET_USERCFG_V51                  = 4181 //获取用户参数
	NET_DVR_SET_USERCFG_V51                  = 4182 //设置用户参数
	NET_DVR_GET_SOFTIO_TRIGGER               = 4183 //获取SoftIO联动配置
	NET_DVR_SET_SOFTIO_TRIGGER               = 4184 //设置SoftIO联动配置
	NET_DVR_GET_SOFTIO_SCHEDULE_CAPABILITIES = 4185 //获取SoftIO布防时间能力
	NET_DVR_GET_SOFTIO_SCHEDULE              = 4186 //获取SoftIO布防时间配置
	NET_DVR_SET_SOFTIO_SCHEDULE              = 4187 //设置SoftIO布防时间配置
	NET_DVR_GET_HFPD_TRIGGER                 = 4188 //获取高频人员侦测联动配置
	NET_DVR_SET_HFPD_TRIGGER                 = 4189 //设置高频人员侦测联动配置
	NET_DVR_GET_HFPD_SCHEDULE_CAPABILITIES   = 4190 //获取高频人员侦测布防时间能力
	NET_DVR_GET_HFPD_SCHEDULE                = 4191 //获取高频人员侦测布防时间配置
	NET_DVR_SET_HFPD_SCHEDULE                = 4192 //设置高频人员侦测布防时间配置
	NET_DVR_GET_ALARM_INFO                   = 4193 //获取报警事件信息
	NET_DVR_GET_USERCFG_V52                  = 4194 //获取用户参数
	NET_DVR_SET_USERCFG_V52                  = 4195 //设置用户参数

	/********************************NVR_后端产品线****************************/
	NET_DVR_GET_MUTEX_FUNCTION = 4353 //获取功能互斥信息

	NET_DVR_GET_SINGLE_CHANNELINFO           = 4360 //获取单个通道属性数据
	NET_DVR_GET_CHANNELINFO                  = 4361 //获取通道属性数据
	NET_DVR_CHECK_LOGIN_PASSWORDCFG          = 4362 //用户登录密码校验
	NET_DVR_GET_SINGLE_SECURITY_QUESTION_CFG = 4363 //获取单个设备安全问题
	NET_DVR_SET_SINGLE_SECURITY_QUESTION_CFG = 4364 //设置单个设备安全问题
	NET_DVR_GET_SECURITY_QUESTION_CFG        = 4365 //获取设备安全问题
	NET_DVR_SET_SECURITY_QUESTION_CFG        = 4366 //设置设备安全问题
	NET_DVR_GET_ONLINEUSERLIST_SC            = 4367 //远程获取登陆用户信息（短连接）

	NET_DVR_GET_BLACKLIST_FACECONTRAST_TRIGGER               = 4368 //获取黑名单人脸比对联动配置
	NET_DVR_SET_BLACKLIST_FACECONTRAST_TRIGGER               = 4369 //设置黑名单人脸比对联动配置
	NET_DVR_GET_WHITELIST_FACECONTRAST_TRIGGER               = 4370 //获取白名单人脸比对联动配置
	NET_DVR_SET_WHITELIST_FACECONTRAST_TRIGGER               = 4371 //设置白名单人脸比对联动配置
	NET_DVR_GET_BLACKLIST_FACECONTRAST_SCHEDULE_CAPABILITIES = 4372 //获取黑名单人脸比对布防时间能力
	NET_DVR_GET_BLACKLIST_FACECONTRAST_SCHEDULE              = 4373 //获取黑名单人脸比对布防时间配置
	NET_DVR_SET_BLACKLIST_FACECONTRAST_SCHEDULE              = 4374 //设置黑名单人脸比对布防时间配置
	NET_DVR_GET_WHITELIST_FACECONTRAST_SCHEDULE_CAPABILITIES = 4375 //获取白名单人脸比对布防时间能力
	NET_DVR_GET_WHITELIST_FACECONTRAST_SCHEDULE              = 4376 //获取白名单人脸比对布防时间配置
	NET_DVR_SET_WHITELIST_FACECONTRAST_SCHEDULE              = 4377 //设置白名单人脸比对布防时间配置

	NET_DVR_GET_HUMAN_RECOGNITION_SCHEDULE_CAPABILITIES = 4378 //获取人体识别布防时间能力
	NET_DVR_GET_HUMAN_RECOGNITION_SCHEDULE              = 4379 //获取人体识别布防时间配置
	NET_DVR_SET_HUMAN_RECOGNITION_SCHEDULE              = 4380 //设置人体识别布防时间配置
	NET_DVR_GET_HUMAN_RECOGNITION_TRIGGER               = 4381 //获取人体识别联动配置
	NET_DVR_SET_HUMAN_RECOGNITION_TRIGGER               = 4382 //设置人体识别联动配置
	NET_DVR_GET_GBT28181_AUDIO_OUTPUT_CFG               = 4383 //获取GBT28181协议接入设备的语音对讲信息
	NET_DVR_SET_GBT28181_AUDIO_OUTPUT_CFG               = 4384 //设置GBT28181协议接入设备的语音对讲信息

	NET_DVR_GET_STUDENTS_STOODUP_TRIGGER                     = 4386 //获取学生起立检测联动配置
	NET_DVR_SET_STUDENTS_STOODUP_TRIGGER                     = 4387 //设置学生起立检测联动配置
	NET_DVR_GET_FRAMES_PEOPLE_COUNTING_SCHEDULE_CAPABILITIES = 4388 //获取区域人数统计布防时间能力
	NET_DVR_GET_FRAMES_PEOPLE_COUNTING_SCHEDULE              = 4389 //获取区域人数统计布防时间配置
	NET_DVR_SET_FRAMES_PEOPLE_COUNTING_SCHEDULE              = 4390 //设置区域人数统计布防时间配置
	NET_DVR_GET_FRAMES_PEOPLE_COUNTING_TRIGGER               = 4391 //获取区域人数统计联动配置
	NET_DVR_SET_FRAMES_PEOPLE_COUNTING_TRIGGER               = 4392 //设置区域人数统计联动配置

	NET_DVR_GET_PERSONDENSITY_TRIGGER               = 4393 //获取人员密度检测的联动配置
	NET_DVR_SET_PERSONDENSITY_TRIGGER               = 4394 //设置人员密度检测的联动配置
	NET_DVR_GET_PERSONDENSITY_SCHEDULE_CAPABILITIES = 4395 //获取人员密度检测的布防时间能力
	NET_DVR_GET_PERSONDENSITY_SCHEDULE              = 4396 //获取人员密度检测的布防时间配置
	NET_DVR_SET_PERSONDENSITY_SCHEDULE              = 4397 //设置人员密度检测的布防时间配置

	NET_DVR_GET_STUDENTS_STOODUP_SCHEDULE_CAPABILITIES = 4398 //获取学生起立检测布防时间能力
	NET_DVR_GET_STUDENTS_STOODUP_SCHEDULE              = 4399 //获取学生起立检测布防时间配置
	NET_DVR_SET_STUDENTS_STOODUP_SCHEDULE              = 4400 //设置学生起立检测布防时间配置

	NET_DVR_SET_FACE_THERMOMETRY_TRIGGER               = 4401 //设置人脸测温联动配置
	NET_DVR_GET_FACE_THERMOMETRY_SCHEDULE_CAPABILITIES = 4402 //获取人脸测温布防时间能力
	NET_DVR_GET_FACE_THERMOMETRY_SCHEDULE              = 4403 //获取人脸测温布防时间配置
	NET_DVR_SET_FACE_THERMOMETRY_SCHEDULE              = 4404 //设置人脸测温布防时间配置
	NET_DVR_GET_FACE_THERMOMETRY_TRIGGER               = 4405 //获取人脸测温联动配置
	NET_DVR_GET_PERSONQUEUE_TRIGGER                    = 4406 //获取人员排队检测的联动配置
	NET_DVR_SET_PERSONQUEUE_TRIGGER                    = 4407 //设置人员排队检测的联动配置
	NET_DVR_GET_PERSONQUEUE_SCHEDULE_CAPABILITIES      = 4408 //获取人员排队检测的布防时间能力
	NET_DVR_GET_PERSONQUEUE_SCHEDULE                   = 4409 //获取人员排队检测的布防时间配置
	NET_DVR_SET_PERSONQUEUE_SCHEDULE                   = 4410 //设置人员排队检测的布防时间配置

	/********************************智能人脸识别****************************/
	NET_DVR_GET_FACESNAPCFG            = 5001 //获取人脸抓拍参数
	NET_DVR_SET_FACESNAPCFG            = 5002 //设置人脸抓拍参数
	NET_DVR_GET_DEVACCESS_CFG          = 5005 //获取接入设备参数
	NET_DVR_SET_DEVACCESS_CFG          = 5006 //设置接入设备参数
	NET_DVR_GET_SAVE_PATH_CFG          = 5007 //获取存储信息参数
	NET_DVR_SET_SAVE_PATH_CFG          = 5008 //设置存储信息参数
	NET_VCA_GET_RULECFG_V41            = 5011 //获取行为分析参数(扩展)
	NET_VCA_SET_RULECFG_V41            = 5012 //设置行为分析参数(扩展)
	NET_DVR_GET_AID_RULECFG_V41        = 5013 //获取交通事件规则参数
	NET_DVR_SET_AID_RULECFG_V41        = 5014 //设置交通事件规则参数
	NET_DVR_GET_TPS_RULECFG_V41        = 5015 //获取交通统计规则参数(扩展)
	NET_DVR_SET_TPS_RULECFG_V41        = 5016 //设置交通统计规则参数(扩展)
	NET_VCA_GET_FACEDETECT_RULECFG_V41 = 5017 //获取ATM人脸检测规则(扩展)
	NET_VCA_SET_FACEDETECT_RULECFG_V41 = 5018 //设置ATM人脸检测规则(扩展)
	NET_DVR_GET_PDC_RULECFG_V41        = 5019 //设置人流量统计规则(扩展)
	NET_DVR_SET_PDC_RULECFG_V41        = 5020 //获取人流量统计规则(扩展)
	NET_DVR_GET_TRIAL_VERSION_CFG      = 5021 //获取试用版信息
	NET_DVR_GET_VCA_CTRLINFO_CFG       = 5022 //批量获取智能控制参数
	NET_DVR_SET_VCA_CTRLINFO_CFG       = 5023 //批量设置智能控制参数
	NET_DVR_SYN_CHANNEL_NAME           = 5024 //同步通道名
	NET_DVR_GET_RESET_COUNTER          = 5025 //获取统计数据清零参数（人流量、交通统计）
	NET_DVR_SET_RESET_COUNTER          = 5026 //设置统计数据清零参数（人流量、交通统计）
	NET_DVR_GET_OBJECT_COLOR           = 5027 //获取物体颜色属性
	NET_DVR_SET_OBJECT_COLOR           = 5028 //设置物体颜色属性
	NET_DVR_GET_AUX_AREA               = 5029 //获取辅助区域
	NET_DVR_SET_AUX_AREA               = 5030 //设置辅助区域
	NET_DVR_GET_CHAN_WORKMODE          = 5031 //获取通道工作模式
	NET_DVR_SET_CHAN_WORKMODE          = 5032 //设置通道工作模式
	NET_DVR_GET_SLAVE_CHANNEL          = 5033 //获取从通道参数
	NET_DVR_SET_SLAVE_CHANNEL          = 5034 //设置从通道参数
	NET_DVR_GET_VQD_EVENT_RULE         = 5035 //获取视频质量诊断事件规则
	NET_DVR_SET_VQD_EVENT_RULE         = 5036 //设置视频质量诊断事件规则
	NET_DVR_GET_BASELINE_SCENE         = 5037 //获取基准场景参数
	NET_DVR_SET_BASELINE_SCENE         = 5038 //设置基准场景参数
	NET_DVR_CONTROL_BASELINE_SCENE     = 5039 //基准场景操作
	NET_DVR_SET_VCA_DETION_CFG         = 5040 //设置智能移动参数配置
	NET_DVR_GET_VCA_DETION_CFG         = 5041 //获取智能移动参数配置
	NET_DVR_GET_STREAM_ATTACHINFO_CFG  = 5042 //获取码流附加信息配置
	NET_DVR_SET_STREAM_ATTACHINFO_CFG  = 5043 //设置码流附加信息配置

	NET_DVR_GET_BV_CALIB_TYPE       = 5044 //获取双目标定类型
	NET_DVR_CONTROL_BV_SAMPLE_CALIB = 5045 //双目样本标定
	NET_DVR_GET_BV_SAMPLE_CALIB_CFG = 5046 //获取双目标定参数
	NET_DVR_GET_RULECFG_V42         = 5049 //获取行为分析参数(支持16条规则扩展)
	NET_DVR_SET_RULECFG_V42         = 5050 //设置行为分析参数(支持16条规则扩展)
	NET_DVR_SET_VCA_DETION_CFG_V40  = 5051 //设置智能移动参数配置
	NET_DVR_GET_VCA_DETION_CFG_V40  = 5052 //获取智能移动参数配置
	NET_DVR_SET_FLASH_CFG           = 5110 //写入数据到Flash 测试使用
	/********************************智能人脸识别 end****************************/

	//2014-12-03
	NET_DVR_GET_T1TEST_CFG = 5053 //产线测试配置接口（获取）
	NET_DVR_SET_T1TEST_CFG = 5054 ////产线测试配置接口（设置）

	/********************************ITS****************************/
	NET_ITS_GET_OVERLAP_CFG_V50 = 5055 //获取字符叠加参数配置扩展
	NET_ITS_SET_OVERLAP_CFG_V50 = 5056 //设置字符叠加参数配置扩展

	NET_DVR_GET_PARKLAMP_STATE   = 5057 //获取停车场信号灯状态信息
	NET_DVR_GET_CLOUDSTORAGE_CFG = 5058 //获取云存储配置参数
	NET_DVR_SET_CLOUDSTORAGE_CFG = 5059 //设置云存储配置参数

	NET_ITS_GET_BASE_INFO                = 5060 //获取终端基本信息
	NET_DVR_GET_SENSOR_INFO              = 5061 //传感器信息查询
	NET_DVR_SET_SENSOR_SWITCH            = 5062 //传感器远程控制
	NET_ITS_GET_IMGMERGE_CFG             = 5063 //获取图片合成配置参数
	NET_ITS_SET_IMGMERGE_CFG             = 5064 //设置图片合成配置参数
	NET_ITS_GET_UPLOAD_CFG               = 5065 //获取数据上传配置
	NET_ITS_SET_UPLOAD_CFG               = 5066 //设置数据上传配置
	NET_DVR_GET_SENSOR_PORT_CAPABILITIES = 5067 //获取传感器能力
	NET_ITS_GET_WORKSTATE                = 5069 //获取终端工作状态
	NET_ITS_GET_IPC_CHAN_CFG             = 5070 //获取通道IPC信息
	NET_ITS_SET_IPC_CHAN_CFG             = 5071 //设置通道IPC信息
	NET_ITS_GET_OVERLAP_CFG              = 5072 //获取字符叠加参数配置
	NET_ITS_SET_OVERLAP_CFG              = 5073 //设置字符叠加参数配置
	NET_DVR_GET_TRIGGEREX_CFG            = 5074 //获取ITC扩展配置
	NET_DVR_SET_TRIGGEREX_CFG            = 5075 //设置ITC扩展配置
	NET_ITS_GET_ROAD_INFO                = 5076 //获取路口信息

	NET_ITS_REMOTE_DEVICE_CONTROL       = 5077 //设置远程设备控制
	NET_ITS_GET_GATEIPC_CHAN_CFG        = 5078 //获取出入口参数
	NET_ITS_SET_GATEIPC_CHAN_CFG        = 5079 //设置出入口参数
	NET_ITS_TRANSCHAN_START             = 5080 //同步数据服务器建立连接
	NET_ITS_GET_ECTWORKSTATE            = 5081 //获取出入口终端工作状态
	NET_ITS_GET_ECT_CHAN_INFO           = 5082 //获取出入口终端通道状态
	NET_DVR_GET_HEATMAP_RESULT          = 5083 //热度图数据查找
	NET_DVR_SET_ITS_EXDEVCFG            = 5084 //设置ITS外接设备信息
	NET_DVR_GET_ITS_EXDEVCFG            = 5085 //获取ITS外接设备信息
	NET_DVR_GET_ITS_EXDEVSTATUS         = 5086 //获取ITS所有外接设备信息
	NET_DVR_SET_ITS_ENDEVCMD            = 5087 //设置ITS终端出入口控制命令
	NET_DVR_SET_ENISSUED_DATADEL        = 5088 //设置ITS终端出入口控制清除
	NET_DVR_GET_PDC_RESULT              = 5089 //客流量数据查询 2014-03-21
	NET_ITS_GET_LAMP_CTRLCFG            = 5090 //获取内外置灯参数
	NET_ITS_SET_LAMP_CTRLCFG            = 5091 //设置内外置灯参数
	NET_ITS_GET_PARKSPACE_ATTRIBUTE_CFG = 5092 //获取特殊车位参数
	NET_ITS_SET_PARKSPACE_ATTRIBUTE_CFG = 5093 //设置特殊车位参数
	NET_ITS_SET_LAMP_EXTERNAL_CFG       = 5095 //设置外控配置参数
	NET_ITS_SET_COMPEL_CAPTURE          = 5096 //设置车位强制抓图
	NET_DVR_SET_TIMESIGN_CFG            = 5097 //设置扩展校时自定义标记
	NET_DVR_GET_TIMESIGN_CFG            = 5098 //获取扩展校时自定义标记
	NET_DVR_GET_SIGNALLAMP_STATUS       = 5099 //信号灯检测
	/********************************ITS end****************************/

	NET_DVR_GET_MONITOR_PLAN_VQD    = 5100 //长连接获取诊断服务器计划
	NET_DVR_GET_MONITORID_VQD       = 5101 //长连接获取对应计划内的监控点信息
	NET_DVR_SET_MONITOR_INFO        = 5102 //批量设置计划内的监控点信息
	NET_DVR_DEL_MONITOR_PLAN_VQD    = 5103 //删除计划
	NET_DVR_GET_MONITOR_VQD_STATUS  = 5104 //平台查询诊断服务器的状态
	NET_DVR_GET_RECORD_INFO         = 5105 //获取资源图片查询
	NET_DVR_GET_MONITOR_VQDCFG      = 5106 //获取服务器的监控点信息
	NET_DVR_SET_MONITOR_VQDCFG      = 5107 //设置服务器的监控点信息
	NET_DVR_SET_MONITOR_PLAN_VQDCFG = 5108 //设置管理计划(单独的计划)

	NET_DVR_SCENE_CHANGE_UPDATE = 5109 //场景变更数据更新

	NET_DVR_GET_CALIBRATE_POINT            = 5153 //归一化坐标转换（枪球联动设备 外部交互命令码 基线代码不实现，防止冲突，提交基线）/*************************智能多场景********************************/
	NET_DVR_GET_SCENE_CFG                  = 5201 //获取场景信息
	NET_DVR_SET_SCENE_CFG                  = 5202 //设置场景信息
	NET_DVR_GET_SCENE_REFERENCE_REGION     = 5203 //获取参考区域
	NET_DVR_SET_SCENE_REFERENCE_REGION     = 5204 //设置参考区域
	NET_DVR_GET_SCENE_CALIBRATION          = 5205 //获取标定信息
	NET_DVR_SET_SCENE_CALIBRATION          = 5206 //设置标定信息
	NET_DVR_GET_SCENE_MASK_REGION          = 5207 //获取屏蔽区域
	NET_DVR_SET_SCENE_MASK_REGION          = 5208 //设置屏蔽区域
	NET_DVR_GET_SCENE_LANECFG              = 5209 //获取车道规则
	NET_DVR_SET_SCENE_LANECFG              = 5210 //设置车道规则
	NET_DVR_GET_SCENE_AID_RULECFG          = 5211 //获取交通事件规则参数
	NET_DVR_SET_SCENE_AID_RULECFG          = 5212 //设置交通事件规则参数
	NET_DVR_GET_SCENE_TPS_RULECFG          = 5213 //获取交通统计规则参数
	NET_DVR_SET_SCENE_TPS_RULECFG          = 5214 //设置交通统计规则参数
	NET_DVR_GET_SCENE_TIME_CFG             = 5215 //获取通道的场景时间段配置
	NET_DVR_SET_SCENE_TIME_CFG             = 5216 //设置通道的场景时间段配置
	NET_DVR_GET_FORENSICS_MODE             = 5217 //获取取证方式参数
	NET_DVR_SET_FORENSICS_MODE             = 5218 //设置取证方式参数
	NET_DVR_FORCESTOP_FORENSICS_CTRL       = 5219 //强制停止取证
	NET_DVR_GET_ALARM_PROCESS_CFG          = 5220 //获取报警处理参数
	NET_DVR_SET_ALARM_PROCESS_CFG          = 5221 //设置报警处理参数
	NET_DVR_GET_BLACKLIST_ALARM_INFO       = 5222 //获取黑白名单报警轨迹
	NET_DVR_GET_STORAGE_RESOURCE_CFG       = 5225 //获取存储资源参数
	NET_DVR_SET_STORAGE_RESOURCE_CFG       = 5226 //设置存储资源参数
	NET_DVR_DEL_BLACKLIST_ALARM_RECORD     = 5227 //远程删除名单报警记录
	NET_DVR_SET_BLACKLIST_GROUP_INFO       = 5229 //远程分组列表参数配置
	NET_DVR_DEL_BLACKLIST_GROUP_INFO       = 5230 //远程删除分组列表
	NET_DVR_GET_BLACKLIST_GROUP_INFO       = 5231 //远程获取全部分组列表
	NET_DVR_SET_BLACKLIST_GROUP_RECORD_CFG = 5232 //分组记录参数配置
	NET_DVR_GET_BLACKLIST_GROUP_RECORD_CFG = 5234 //远程获取分组记录参数
	NET_DVR_DEL_BLACKLIST_GROUP_RECORD_CFG = 5235 //远程删除分组记录参数
	NET_DVR_GET_AREA_MONITOR_CFG           = 5236 //获取区域监控点参数
	NET_DVR_SET_AREA_MONITOR_CFG           = 5237 //设置区域监控点参数
	NET_DVR_DEL_AREA_MONITOR_CFG           = 5238 //删除区域监控点
	NET_DVR_RETRIEVAL_SNAP_RECORD          = 5240 //抓拍库检索
	NET_DVR_GET_ALARMLIST                  = 5241 //获取名单报警列表
	NET_DVR_DETECT_IMAGE                   = 5242 //单张图片检测
	NET_DVR_GET_SNAP_RECORD                = 5243 //获取抓拍记录
	NET_DVR_DEL_SNAP_RECORD                = 5244 //删除抓拍记录
	NET_DVR_GET_FACE_RECORD                = 5245 //远程获取人脸记录列表
	NET_DVR_SET_FACE_RECORD                = 5246 //添加人脸记录
	NET_DVR_DEL_FACE_RECORD                = 5247 //删除人脸记录
	NET_DVR_GET_FACE_DATABASE              = 5248 //获取人脸库配置参数
	NET_DVR_SET_FACE_DATABASE              = 5249 //设置人脸库配置参数
	NET_DVR_DEL_FACE_DATABASE              = 5250 //删除人脸库
	NET_DVR_RETRIEVAL_FACE_DATABASE        = 5251 //人脸库检索
	NET_DVR_SET_BLACKLIST_REL_DEV_CFG      = 5252 //设备关联名单分组关联
	NET_DVR_DEL_BLACKLIST_REL_DEV          = 5253 //删除 设备关联名单分组信息
	/*************************智能多场景end*****************************/

	NET_DVR_GET_DISK_RAID_INFO = 6001 //获取磁盘Raid信息
	NET_DVR_SET_DISK_RAID_INFO = 6002 //设置磁盘Raid信息

	NET_DVR_GET_DVR_SYNCHRONOUS_IPC = 6005 //获取：是否为前端IPC同步设备参数
	NET_DVR_SET_DVR_SYNCHRONOUS_IPC = 6006 //设置：是否为前端IPC同步设备参数

	NET_DVR_SET_DVR_IPC_PASSWD        = 6008 //设置：IPC用户名密码
	NET_DVR_GET_DEVICE_NET_USING_INFO = 6009 //获取：当前设备网络资源使用情况

	NET_DVR_SET_DVR_IPC_NET = 6012 //设置：设置前端IPC的网络地址

	NET_DVR_GET_RECORD_CHANNEL_INFO = 6013 //获取：录像通道信息
	NET_DVR_SET_RECORD_CHANNEL_INFO = 6014 //设置：录像通道信息

	NET_DVR_MOUNT_DISK   = 6015 // 加载磁盘
	NET_DVR_UNMOUNT_DISK = 6016 // 卸载磁盘

	// CVR
	NET_DVR_GET_STREAM_SRC_INFO         = 6017 //获取：流来源信息
	NET_DVR_SET_STREAM_SRC_INFO         = 6018 //设置：流来源信息
	NET_DVR_GET_STREAM_RECORD_INFO      = 6019 //获取：流录像信息
	NET_DVR_SET_STREAM_RECORD_INFO      = 6020 //设置：流录像信息
	NET_DVR_GET_STREAM_RECORD_STATUS    = 6021 //获取：流录像状态
	NET_DVR_SET_STREAM_RECORD_STATUS    = 6022 //设置：流录像状态
	NET_DVR_GET_STREAM_INFO             = 6023 //获取已添加的流ID信息
	NET_DVR_GET_STREAM_SRC_INFO_V40     = 6024 //获取：流来源信息
	NET_DVR_SET_STREAM_SRC_INFO_V40     = 6025 //设置：流来源信息
	NET_DVR_GET_RELOCATE_INFO           = 6026 //获取N+0模式下重定向信息
	NET_DVR_START_GOP_INFO_PASSBACK     = 6032 //智能信息回填
	NET_DVR_GET_CHANS_RECORD_STATUS_CFG = 6035 //获取通道录像状态信息
	NET_DVR_SET_CHANS_RECORD_STATUS_CFG = 6036 //设置通道录像状态信息
	//NVR：96xx
	NET_DVR_GET_IP_ALARM_GROUP_NUM = 6100 //获取：IP通道报警输入输出组数
	NET_DVR_GET_IP_ALARM_IN        = 6101 //获取：IP通道报警输入信息
	NET_DVR_GET_IP_ALARM_OUT       = 6102 //获取：IP通道报警输出信息

	//9000 v2.2
	NET_DVR_GET_FTPCFG_SECOND = 6103 //获取图片上传FTP参数
	NET_DVR_SET_FTPCFG_SECOND = 6104 //设置图片上传FTP参数

	NET_DVR_GET_DEFAULT_VIDEO_EFFECT = 6105 // 获取视频输入效果参数默认值
	NET_DVR_SET_VIDEO_EFFECT         = 6106 // 设置通道视频输入图像参数
	NET_DVR_DEL_INVALID_DISK         = 6107 // 删除无效磁盘

	NET_DVR_GET_DRAWFRAME_DISK_QUOTA_CFG = 6109 //获取抽帧通道磁盘配额
	NET_DVR_SET_DRAWFRAME_DISK_QUOTA_CFG = 6110 //设置抽帧通道磁盘配额

	NET_DVR_GET_NAT_CFG                 = 6111 //获取NAT映射参数
	NET_DVR_SET_NAT_CFG                 = 6112 //设置NAT映射参数
	NET_DVR_GET_AES_KEY                 = 6113 //获取设备AES加密密钥
	NET_DVR_GET_POE_CFG                 = 6114 //获取POE参数
	NET_DVR_SET_POE_CFG                 = 6115 //设置POE参数
	NET_DVR_GET_CUSTOM_PRO_CFG          = 6116 //获取自定义协议参数
	NET_DVR_SET_CUSTOM_PRO_CFG          = 6117 //设置自定义协议参数
	NET_DVR_GET_STREAM_CABAC            = 6118 //获取码流压缩性能选项
	NET_DVR_SET_STREAM_CABAC            = 6119 //设置码流压缩性能选项
	NET_DVR_GET_ESATA_MINISAS_USAGE_CFG = 6120 //获取eSATA和miniSAS用途
	NET_DVR_SET_ESATA_MINISAS_USAGE_CFG = 6121 //设置eSATA和miniSAS用途

	NET_DVR_GET_HDCFG_V40             = 6122 //获取硬盘信息参数
	NET_DVR_SET_HDCFG_V40             = 6123 //设置硬盘信息参数
	NET_DVR_GET_POE_CHANNEL_ADD_MODE  = 6124 //获取POE通道添加方式
	NET_DVR_SET_POE_CHANNEL_ADD_MODE  = 6125 //设置POE通道添加方式
	NET_DVR_GET_DIGITAL_CHANNEL_STATE = 6126 //获取设备数字通道状态
	NET_DVR_GET_BONJOUR_CFG           = 6127 // 获取Bonjour信息
	NET_DVR_SET_BONJOUR_CFG           = 6128 // 设置Bonjour信息

	NET_DVR_GET_SOCKS_CFG = 6130 //获取SOCKS信息
	NET_DVR_SET_SOCKS_CFG = 6131 //设置SOCKS信息

	NET_DVR_GET_QOS_CFG = 6132 //获取QoS信息
	NET_DVR_SET_QOS_CFG = 6133 //设置QoS信息

	NET_DVR_GET_HTTPS_CFG = 6134 //获取HTTPS信息
	NET_DVR_SET_HTTPS_CFG = 6135 //设置HTTPS信息

	NET_DVR_GET_WD1_CFG = 6136 //远程获取WD1使能开关
	NET_DVR_SET_WD1_CFG = 6137 //远程设置WD1使能开关

	NET_DVR_CREATE_CERT = 6138 //创建证书
	NET_DVR_DELETE_CERT = 6139 //删除证书

	NET_DVR_GET_RECORD_LOCK_PERCENTAGE = 6140 //获取录像段锁定比例
	NET_DVR_SET_RECORD_LOCK_PERCENTAGE = 6141 //设置录像段锁定比例

	NET_DVR_CMD_TRIGGER_PERIOD_RECORD = 6144 //外部命令触发指定时间录像
	NET_DVR_UPLOAD_CERT               = 6145 //上传证书
	NET_DVR_DOWNLOAD_CERT             = 6146 //下载证书
	NET_DVR_GET_CERT                  = 6147 //获取证书

	NET_DVR_GET_POS_FILTER_CFG  = 6148 //获取POS过滤规则
	NET_DVR_SET_POS_FILTER_CFG  = 6149 //设置POS过滤规则
	NET_DVR_GET_CONNECT_POS_CFG = 6150 //获取DVR与POS连接方式
	NET_DVR_SET_CONNECT_POS_CFG = 6151 //设置DVR与POS连接方式
	NET_DVR_GET_CHAN_FILTER_CFG = 6152 //获取规则与通道关联信息
	NET_DVR_SET_CHAN_FILTER_CFG = 6153 //设置规则与通道关联信息

	NET_DVR_GET_FTPCFG_V40 = 6162 //获取FTP信息
	NET_DVR_SET_FTPCFG_V40 = 6163 //设置FTP信息

	NET_DVR_GET_MONTHLY_RECORD_DISTRIBUTION = 6164 //获取月历录像分布
	NET_DVR_GET_ACCESS_DEVICE_CHANNEL_INFO  = 6165 //获取待接入设备通道信息
	NET_DVR_GET_PREVIEW_SWITCH_CFG          = 6166 //获取设备本地预览切换参数
	NET_DVR_SET_PREVIEW_SWITCH_CFG          = 6167 //设置设备本地预览切换参数

	//Netra3.0.0
	NET_DVR_GET_N_PLUS_ONE_WORK_MODE = 6168 //获取N+1工作模式
	NET_DVR_SET_N_PLUS_ONE_WORK_MODE = 6169 //设置N+1工作模式

	NET_DVR_GET_HD_STATUS = 6170 //获取硬盘状态
	NET_DVR_SET_HD_STATUS = 6171 //设置硬盘状态

	NET_DVR_IMPORT_IPC_CFG_FILE = 6172 //导入IPC配置文件
	NET_DVR_EXPORT_IPC_CFG_FILE = 6173 //导出IPC配置文件
	NET_DVR_UPGRADE_IPC         = 6174 //升级IP通道

	NET_DVR_GET_RAID_BACKGROUND_TASK_SPEED = 6175 //获取RAID后台任务速度
	NET_DVR_SET_RAID_BACKGROUND_TASK_SPEED = 6176 //设置RAID后台任务速度

	//marvell 256路NVR
	NET_DVR_GET_EXCEPTIONCFG_V40       = 6177 //获取异常参数配置
	NET_DVR_SET_EXCEPTIONCFG_V40       = 6178 //设置异常参数配置
	NET_DVR_GET_PICCFG_V40             = 6179 //获取图象参数 支持变长    NetSDK_
	NET_DVR_SET_PICCFG_V40             = 6180 //设置图象参数， 支持变长
	NET_DVR_GET_ALARMINCFG_V40         = 6181 //获取报警输入参数，支持变长
	NET_DVR_SET_ALARMINCFG_V40         = 6182 //获取报警输入参数，支持变长
	NET_DVR_GET_IPALARMINCFG_V40       = 6183 //获取IP报警输入接入配置信息
	NET_DVR_GET_IPALARMOUTCFG_V40      = 6185 //获取IP报警输出接入配置信息
	NET_DVR_GET_USERCFG_V40            = 6187 //获取用户参数
	NET_DVR_SET_USERCFG_V40            = 6188 //设置用户参数
	NET_DVR_GET_WORK_STATUS            = 6189 //获取设备工作状态
	NET_DVR_GET_JPEG_CAPTURE_CFG_V40   = 6190 //获取DVR抓图配置
	NET_DVR_SET_JPEG_CAPTURE_CFG_V40   = 6191 //设置DVR抓图配置
	NET_DVR_GET_HDGROUP_CFG_V40        = 6192 //获取盘组管理配置参数
	NET_DVR_SET_HDGROUP_CFG_V40        = 6193 //设置盘组管理配置参数
	NET_DVR_GET_SMD_HOLIDAY_HANDLE     = 6194 //获取简易智能假日计划
	NET_DVR_SET_SMD_HOLIDAY_HANDLE     = 6195 //设置简易智能假日计划
	NET_DVR_GET_PIC_MODEL_CFG          = 6196 //获取图片建模配置参数
	NET_DVR_SET_PIC_MODEL_CFG          = 6197 //设置图片建模配置参数
	NET_DVR_START_LOCAL_MOUSE_EVENT    = 6198 //开启设备本地鼠标事件记录
	NET_DVR_START_SIMULARE_MOUSE_EVENT = 6199 //远程模拟鼠标事件
	NET_DVR_GET_WORK_STATUS_V50        = 6200 //获取设备工作状态V50

	//91系列HD-SDI高清DVR
	NET_DVR_GET_ACCESS_CAMERA_INFO = 6201 // 获取前端相机信息
	NET_DVR_SET_ACCESS_CAMERA_INFO = 6202 // 设置前端相机信息
	NET_DVR_PULL_DISK              = 6203 // 安全拔盘
	NET_DVR_SCAN_RAID              = 6204 // 扫描阵列
	// CVR 2.0.X
	NET_DVR_GET_USER_RIGHT_CFG = 6210 // 获取用户权限
	NET_DVR_SET_USER_RIGHT_CFG = 6211 // 设置用户权限

	NET_DVR_ONE_KEY_CONFIG  = 6212 // 一键配置CVR
	NET_DVR_RESTART_SERVICE = 6213 // 重启CVR服务

	NET_DVR_GET_MAX_MACHINE_NUM_CFG = 6214 // 获取备机最大个数
	NET_DVR_SET_MAX_MACHINE_NUM_CFG = 6215 // 设置备机最大个数

	NET_DVR_ADD_DEVICE = 6216 //N+1模式添加设备
	NET_DVR_DEL_DEVICE = 6217 //N+1模式删除设备

	NET_DVR_GET_DATA_CALLBACK_CFG = 6218 // 获取数据回迁状态
	NET_DVR_SET_DATA_CALLBACK_CFG = 6219 // 设置数据回迁状态

	NET_DVR_CLONE_LUN  = 6220 //克隆LUN卷
	NET_DVR_EXPAND_LUN = 6221 //扩展和重命名LUN卷

	NET_DVR_GET_N_PLUS_ONE_DEVICE_INFO = 6222 //获取N+1设备信息
	NET_DVR_MODIFY_DVR_NET_DISK        = 6223 //修改DVR网盘
	//    NET_DVR_DEL_DVR_NET_DISK                6224    //删除DVR网盘

	NET_DVR_CREATE_NAS = 6225 //创建NAS
	NET_DVR_DELETE_NAS = 6226 //删除NAS

	NET_DVR_OPEN_ISCSI  = 6227 //开启iSCSI
	NET_DVR_CLOSE_ISCSI = 6228 //关闭iSCSI

	NET_DVR_GET_FC             = 6229 //获取光纤信息
	NET_DVR_OPEN_FC            = 6230 //开启FC
	NET_DVR_CLOSE_FC           = 6231 //关闭FC
	NET_DVR_ONE_KEY_CONFIG_SAN = 6232 // 一键配置SAN, 与一键配置CVR逻辑一样

	//CVR2.3.2
	NET_DVR_RECORD_CHECK                        = 6233 //录像完整性检测
	NET_DVR_ADD_RECORD_PASSBACK_TASK_MANUAL     = 6234 //手动添加录像回传任务
	NET_DVR_GET_ALL_RECORD_PASSBACK_TASK_MANUAL = 6235 //获取所有手动添加录像回传任务
	NET_DVR_RECORD_PASSBACK_TASK_MANUAL_CTRL    = 6236 //控制手动录像回传任务
	NET_DVR_DEL_RECORD_PASSBACK_TASK_MANUAL     = 6237 //删除手动录像回传任务
	NET_DVR_GET_RECORD_PASSBACK_PLAN_CFG        = 6238 //获取录像回传计划配置
	NET_DVR_SET_RECORD_PASSBACK_PLAN_CFG        = 6239 //设置录像回传计划配置
	NET_DVR_GET_DEV_STORAGE_CFG                 = 6240 //获取设备存储信息
	NET_DVR_GET_ONLINE_USER_CFG                 = 6241 //获取在线用户参数
	NET_DVR_GET_RECORD_SEGMENT_CFG              = 6242 //获取录像段总量

	NET_DVR_GET_REC_PASSBACK_TASK_EXECUTABLE = 6243 //查询手动录像回传任务可执行性
	NET_DVR_GET_STREAM_MEDIA_CFG             = 6244 //获取流媒体回传录像参数配置（流ID方式）
	NET_DVR_SET_STREAM_MEDIA_CFG             = 6245 //设置流媒体回传录像参数配置（流ID方式）
	NET_DVR_GET_USERCFG_V50                  = 6246 //获取用户参数V50
	NET_DVR_SET_USERCFG_V50                  = 6247 //设置用户参数V50

	NET_DVR_GET_RECORD_PASSBACK_BASIC_CFG_CAP        = 6248 //获取CVR回传功能基础配置能力
	NET_DVR_GET_RECORD_PASSBACK_BASIC_CFG            = 6249 //获取CVR回传功能基础配置
	NET_DVR_SET_RECORD_PASSBACK_BASIC_CFG            = 6250 //设置CVR回传功能基础配置
	NET_DVR_ONE_KEY_CONFIG_V50                       = 6251 // 一键配置CVR(V50)
	NET_DVR_GET_RACM_CAP                             = 6252 //获取存储总能力（RACM能力）
	NET_DVR_GET_THUMBNAILS                           = 6253 //获取缩略图（默认是录像的缩略图）(支持流ID)
	NET_DVR_ADD_RECORD_PASSBACK_TASK_MANUAL_V50      = 6254 //手动添加录像回传任务V50（返回任务ID）
	NET_DVR_GET_RECORD_PASSBACK_HISTORY_PLAN_CFG_CAP = 6255 //获取CVR回传历史录像计划能力
	NET_DVR_GET_RECORD_PASSBACK_HISTORY_PLAN_CFG     = 6256 //获取CVR回传历史录像计划配置
	NET_DVR_SET_RECORD_PASSBACK_HISTORY_PLAN_CFG     = 6257 //设置CVR回传历史录像计划配置
	NET_DVR_ONE_KEY_CONFIG_V51                       = 6258 // 一键配置CVR(V51)

	NET_DVR_GET_RECORD_PACK = 6301 //获取录像打包参数
	NET_DVR_SET_RECORD_PACK = 6302 //设置录像打包参数

	NET_DVR_GET_CLOUD_STORAGE_CFG = 6303 //获取设备当前工作模式
	NET_DVR_SET_CLOUD_STORAGE_CFG = 6304 //设置设备当前工作模式
	NET_DVR_GET_GOP_INFO          = 6305 //获取GOP信息
	NET_DVR_GET_PHY_DISK_INFO     = 6306 //获取物理磁盘信息
	//录播主机外部命令
	NET_DVR_GET_RECORDING_AUTO_TRACK_CFG = 6307 //获取SDI自动跟踪配置信息
	NET_DVR_SET_RECORDING_AUTO_TRACK_CFG = 6308 //设置SDI自动跟踪配置信息

	NET_DVR_GET_RECORDING_PUBLISH_CFG = 6309 //获取一键发布信息
	NET_DVR_SET_RECORDING_PUBLISH_CFG = 6310 //设置一键发布信息

	NET_DVR_RECORDING_ONEKEY_CONTROL = 6311 //录播主机控制

	NET_DVR_GET_RECORDING_END_TIME = 6312 //获取录播剩余时间

	NET_DVR_RECORDING_PUBLISH = 6313 //一键发布录像

	NET_DVR_GET_CURRICULUM_CFG = 6314 //获取课表配置信息
	NET_DVR_SET_CURRICULUM_CFG = 6315 //设置课表配置信息

	NET_DVR_GET_COURSE_INDEX_CFG = 6316 //获取课程信息索引
	NET_DVR_SET_COURSE_INDEX_CFG = 6317 //设置课程信息索引

	NET_DVR_GET_PPT_CHANNEL    = 6318 //获取PPT支持通道号
	NET_DVR_GET_PPT_DETECT_CFG = 6319 //获取PPT检测参数
	NET_DVR_SET_PPT_DETECT_CFG = 6320 //设置PPT检测参数

	NET_DVR_GET_RECORDINGHOST_CFG = 6321 //获取录播主机配置信息
	NET_DVR_SET_RECORDINGHOST_CFG = 6322 //设置录播主机配置信息
	NET_DVR_GET_BACKUP_RECORD_CFG = 6323 //获取一键备份配置信息
	NET_DVR_SET_BACKUP_RECORD_CFG = 6324 //设置一键备份配置信息

	//庭审主机
	NET_DVR_GET_AUDIO_ACTIVATION_CFG = 6326 //获取语音激励配置参数
	NET_DVR_SET_AUDIO_ACTIVATION_CFG = 6327 //设置语音激励配置参数
	NET_DVR_GET_DECODERCFG_V40       = 6328 //获取解码器参数信息
	NET_DVR_SET_DECODERCFG_V40       = 6329 //设置解码器参数信息

	NET_DVR_INFRARED_OUTPUT_CONTROL   = 6330 //红外输出控制
	NET_DVR_GET_INFRARED_CMD_NAME_CFG = 6331 //获取红外命令名称参数配置
	NET_DVR_SET_INFRARED_CMD_NAME_CFG = 6332 //设置红外命令名称参数配置
	NET_DVR_START_INFRARED_LEARN      = 6333 //远程红外学码

	NET_DVR_GET_TRIAL_SYSTEM_CFG        = 6334 //获取庭审主机系统信息
	NET_DVR_SET_CASE_INFO               = 6335 //案件信息录入
	NET_DVR_GET_TRIAL_MICROPHONE_STATUS = 6336 //获取麦克风状态信息
	NET_DVR_SET_TRIAL_MICROPHONE_STATUS = 6337 //获取麦克风状态信息
	NET_DVR_GET_TRIAL_HOST_STATUS       = 6338 //获取庭审主机状态信息
	NET_DVR_GET_LAMP_OUT                = 6339 //获取LAMP输出口信息
	NET_DVR_SET_LAMP_OUT                = 6340 //设置LAMP输出口信息
	NET_DVR_LAMP_REMOTE_CONTROL         = 6341 // LAMP控制
	NET_DVR_REMOTE_CONTROL_PLAY         = 6342 //远程控制本地回放
	NET_DVR_GET_LOCAL_INPUT_CFG         = 6343 //获取庭审主机状态信息庭审主机本地输入信息
	NET_DVR_SET_LOCAL_INPUT_CFG         = 6344 //设置庭审主机本地输入信息
	NET_DVR_GET_CASE_INFO               = 6345 //获取当前案件信息

	//审讯机外部命令
	NET_DVR_INQUEST_GET_CDW_STATUS            = 6350 //获取审讯机刻录状态-长连接
	NET_DVR_GET_MIX_AUDIOIN_CFG               = 6351 //获取混音输入口参数配置
	NET_DVR_SET_MIX_AUDIOIN_CFG               = 6352 //设置混音输入口参数配置
	NET_DVR_GET_MIX_AUDIOOUT_CFG              = 6353 //获取混音输出口参数配置
	NET_DVR_SET_MIX_AUDIOOUT_CFG              = 6354 //设置混音输出口参数配置
	NET_DVR_GET_AUDIOIN_VOLUME_CFG            = 6355 //获取音频输入口音量调节参数配置
	NET_DVR_SET_AUDIOIN_VOLUME_CFG            = 6356 //设置音频输入口音量调节参数配置
	NET_DVR_GET_AREA_MASK_CFG                 = 6357 //获取马赛克区域配置
	NET_DVR_SET_AREA_MASK_CFG                 = 6358 //设置马赛克区域配置
	NET_DVR_GET_AUDIO_DIACRITICAL_CFG         = 6359 //获取音频变音配置
	NET_DVR_SET_AUDIO_DIACRITICAL_CFG         = 6360 //设置音频变音配置
	NET_DVR_GET_WIFI_DHCP_ADDR_CFG            = 6361 //获WIFI DHCP 地址范围参数配置
	NET_DVR_SET_WIFI_DHCP_ADDR_CFG            = 6362 //设WIFI DHCP 地址范围参数配置
	NET_DVR_GET_WIFI_CLIENT_LIST_INFO         = 6363 //获取wifi热点下连接的设备信息
	NET_DVR_REMOTECONTROL_POWER_ON            = 6364 //远程开机
	NET_DVR_GET_MULTISTREAM_RELATION_CHAN_CFG = 6365 //获取多码流关联通道参数配置
	NET_DVR_SET_MULTISTREAM_RELATION_CHAN_CFG = 6366 //设置多码流关联通道参数配置
	NET_DVR_GET_VIDEOOUT_RESOLUTION_CFG       = 6367 //获取设备本地视频输出口分辨率
	NET_DVR_SET_VIDEOOUT_RESOLUTION_CFG       = 6368 //设置设备本地视频输出口分辨率
	NET_DVR_GET_AUDIOOUT_VOLUME_CFG           = 6369 //获取音频输出口音量调节参数配置
	NET_DVR_SET_AUDIOOUT_VOLUME_CFG           = 6370 //设置音频输出口音量调节参数配置
	NET_DVR_INQUEST_PAUSE_CDW                 = 6371 //暂停刻录
	NET_DVR_INQUEST_RESUME_CDW                = 6372 //恢复刻录
	NET_DVR_GET_INPUT_CHAN_CFG                = 6373 //获取输入通道配置
	NET_DVR_SET_INPUT_CHAN_CFG                = 6374 //设置输入通道配置
	NET_DVR_GET_INQUEST_MIX_AUDIOIN_CFG       = 6375 //获取审讯机音频输入混音配置
	NET_DVR_SET_INQUEST_MIX_AUDIOIN_CFG       = 6376 //设置审讯机音频输入混音配置
	NET_DVR_CASE_INFO_CTRL                    = 6377 //案件信息显示控制
	NET_DVR_GET_INQUEST_USER_RIGHT            = 6378 //获取审讯机用户权限
	NET_DVR_SET_INQUEST_USER_RIGHT            = 6379 //设置审讯机用户权限
	NET_DVR_GET_INQUEST_CASE_INFO             = 6380 //获取审讯案件信息配置
	NET_DVR_SET_INQUEST_CASE_INFO             = 6381 //设置审讯案件信息配置

	NET_DVR_GET_FILM_MODE_CFG     = 6387 //获取电影模式
	NET_DVR_SET_FILM_MODE_CFG     = 6388 //设置电影模式
	NET_DVR_GET_FILM_MODE_CFG_CAP = 6389 //获取电影模式配置能力

	NET_DVR_GET_DIRECTED_STRATEGY_CFG     = 6390 //获取导播策略类型
	NET_DVR_SET_DIRECTED_STRATEGY_CFG     = 6391 //设置导播策略类型
	NET_DVR_GET_DIRECTED_STRATEGY_CFG_CAP = 6392 //获取电影模式配置能力
	NET_DVR_GET_FRAME_CFG                 = 6393 //获取画面边框
	NET_DVR_SET_FRAME_CFG                 = 6394 //设置画面边框
	NET_DVR_GET_FRAME_CFG_CAP             = 6395 //获取画面边框配置能力
	NET_DVR_GET_AUDIO_EFFECTIVE_CFG       = 6396 //获取音频优化参数
	NET_DVR_SET_AUDIO_EFFECTIVE_CFG       = 6397 //设置音频效果参数
	NET_DVR_GET_AUDIO_EFFECTIVE_CFG_CAP   = 6398 //获取音频效果优化配置能力
	NET_DVR_GET_RECORD_VIDEO_CFG          = 6399 //获取录制视频参数
	NET_DVR_SET_RECORD_VIDEO_CFG          = 6400 //设置录制视频参数

	NET_DVR_GET_OUTPUT_CFG      = 6401 //获取显示输出参数
	NET_DVR_SET_OUTPUT_CFG      = 6402 //设置显示输出参数
	NET_DVR_CODER_DISPLAY_START = 6403 //开始输出
	NET_DVR_CODER_DISPLAY_STOP  = 6404 //停止输出
	NET_DVR_GET_WINDOW_STATUS   = 6405 //获取显示窗口状态

	//VQD功能接口
	NET_DVR_GET_VQD_LOOP_DIAGNOSE_CFG = 6406 //获取VQD循环诊断配置参数
	NET_DVR_SET_VQD_LOOP_DIAGNOSE_CFG = 6407 //设置VQD循环诊断配置参数
	NET_DVR_GET_VQD_DIAGNOSE_INFO     = 6408 //手动获取VQD诊断信息

	NET_DVR_RECORDING_PUBLISH_FILE              = 6421 //文件发布
	NET_DVR_GET_RECORDING_PUBLISH_FILE_CAP      = 6422 //获取文件发布能力
	NET_DVR_GET_PUBLISH_PROGRESS                = 6423 //获取发布进度
	NET_DVR_GET_RECORD_VIDEO_CFG_CAP            = 6424 //获取录制视频配置能力
	NET_DVR_GET_RTMP_CFG                        = 6425 //获取RTMP参数
	NET_DVR_SET_RTMP_CFG                        = 6426 //设置RTMP参数
	NET_DVR_GET_RTMP_CFG_CAP                    = 6427 //获取RTMP配置能力
	NET_DVR_DEL_BACKGROUND_PIC                  = 6428 //删除背景图片文件
	NET_DVR_GET_BACKGROUND_PIC_CFG              = 6429 //查询背景图片文件
	NET_DVR_GET_BACKGROUND_PIC_INFO             = 6430 //获取哪张图片作为背景图片
	NET_DVR_SET_BACKGROUND_PIC_INFO             = 6431 //设置哪张图片作为背景图片
	NET_DVR_GET_BACKGROUND_PIC_INFO_CAP         = 6432 //获取哪张图片作为背景图片配置能力
	NET_DVR_GET_RECORD_HOST_CAP                 = 6433 //获取录播主机总能力
	NET_DVR_GET_COURSE_LIST                     = 6434 //获取课程列表
	NET_DVR_GET_RECORD_STATUS                   = 6435 //查询录播主机当前状态
	NET_DVR_MANUAL_CURRICULUM_CONTROL           = 6436 //手动课表控制
	NET_DVR_GET_IMAGE_DIFF_DETECTION_CFG        = 6437 //获取图像差分检测参数
	NET_DVR_SET_IMAGE_DIFF_DETECTION_CFG        = 6438 //设置图像差分检测参数
	NET_DVR_GET_IMAGE_DIFF_DETECTION_CFG_CAP    = 6439 //获取图像差分检测配置能力
	NET_DVR_GET_RECORDING_PUBLISH_FILE_INFO     = 6440 //获取发布文件信息参数
	NET_DVR_SET_RECORDING_PUBLISH_FILE_INFO     = 6441 //设置发布文件信息参数
	NET_DVR_GET_RECORDING_PUBLISH_FILE_INFO_CAP = 6442 //获取发布文件信息配置能力
	NET_DVR_MANUAL_CURRICULUM_CONTROL_CAP       = 6443 //获取手动课程录像的能力
	NET_DVR_GET_STATISTIC_DATA_LIST             = 6444 //获取统计数据列表

	NET_DVR_GET_DEVICE_LAN_ENCODE             = 6501 //获取设备的语言编码
	NET_DVR_GET_GBT28181_SERVICE_CFG          = 6503 //获取GB28181服务器参数
	NET_DVR_SET_GBT28181_SERVICE_CFG          = 6504 //设置GB28181服务器参数
	NET_DVR_GET_GBT28181_SERVICE_CAPABILITIES = 6505 //获取GB28181服务器能力

	NET_DVR_GET_CLOUD_URL                       = 6506 //获取云存储URL
	NET_DVR_GET_CLOUD_URL_CAP                   = 6507 //获取云存储URL-能力集
	NET_DVR_GET_CLOUD_CFG                       = 6508 //获取云存储配置参数
	NET_DVR_SET_CLOUD_CFG                       = 6509 //设置云存储配置参数
	NET_DVR_GET_CLOUD_CFG_CAP                   = 6510 //获取云存储配置-能力集
	NET_DVR_GET_CLOUD_UPLOADSTRATEGY            = 6511 //获取云存储上传策略
	NET_DVR_SET_CLOUD_UPLOADSTRATEGY            = 6512 //设置云存储上传策略
	NET_DVR_GET_CLOUDSTORAGE_UPLOADSTRATEGY_CAP = 6513 //云存储上传策略配置-能力集

	NET_DVR_GET_VIDEO_IMAGE_DB_CFG     = 6601 //获取视图库信息
	NET_DVR_SET_VIDEO_IMAGE_DB_CFG     = 6602 //设置视图库信息
	NET_DVR_GET_VIDEO_IMAGE_DB_CFG_CAP = 6603 //获取视图库相关能力
	NET_DVR_GET_FILE_INFO_BY_ID        = 6604 //根据文件ID获取视图库中文件信息
	NET_DVR_QUERY_FILE_INFO_CAP        = 6605 //根据文件名查询文件信息能力
	NET_DVR_DEL_FILE_FROM_DB           = 6606 //从视图库中删除文件
	NET_DVR_GET_VIDEO_IMAGE_DB_CAP     = 6607 //获取视图库总能力

	NET_DVR_GET_FIGURE = 6610 //获取缩略图

	NET_DVR_SYNC_IPC_PASSWD               = 6621 //同步IPC密码与NVR一致
	NET_DVR_GET_VEHICLE_BLACKLST_SCHEDULE = 6622 //获取黑名单布防时间配置
	NET_DVR_SET_VEHICLE_BLACKLST_SCHEDULE = 6623 //设置黑名单布防时间配置

	NET_DVR_GET_VEHICLE_WHITELST_SCHEDULE = 6624 //获取白名单布防时间配置
	NET_DVR_SET_VEHICLE_WHITELST_SCHEDULE = 6625 //设置白名单布防时间配置

	NET_DVR_GET_VEHICLE_BLACKLIST_EVENT_TRIGGER = 6626 //获取黑名单布防联动配置
	NET_DVR_SET_VEHICLE_BLACKLIST_EVENT_TRIGGER = 6627 //设置黑名单布防联动配置

	NET_DVR_GET_VEHICLE_WHITELIST_EVENT_TRIGGER = 6628 //获取白名单布防联动配置
	NET_DVR_SET_VEHICLE_WHITELIST_EVENT_TRIGGER = 6629 //设置白名单布防联动配置

	NET_DVR_GET_TRAFFIC_CAP                     = 6630 //获取抓拍相关能力集
	NET_DVR_GET_VEHICLE_ALLLIST_EVENT_TRIGGER   = 6631 //获取全部车辆检测布防联动配置
	NET_DVR_SET_VEHICLE_ALLLIST_EVENT_TRIGGER   = 6632 //设置全部车辆检测布防联动配置
	NET_DVR_GET_VEHICLE_OTHERLIST_EVENT_TRIGGER = 6633 //获取其他单布防联动配置
	NET_DVR_SET_VEHICLE_OTHERLIST_EVENT_TRIGGER = 6634 //设置其他单布防联动配置

	NET_DVR_GET_STORAGEDETECTION_EVENT_TRIGGER         = 6635 //获取存储健康检测联动配置
	NET_DVR_SET_STORAGEDETECTION_EVENT_TRIGGER         = 6636 //设置存储健康检测联动配置
	NET_DVR_GET_STORAGEDETECTION_SCHEDULE_CAPABILITIES = 6637 //获取存储健康检测布防时间能力
	NET_DVR_GET_STORAGEDETECTION_SCHEDULE              = 6638 //获取存储健康布防时间配置
	NET_DVR_SET_STORAGEDETECTION_SCHEDULE              = 6639 //设置存储健康布防时间配置
	NET_DVR_GET_STORAGEDETECTION_STATE                 = 6640 //获取存储健康状态

	NET_DVR_GET_STORAGEDETECTION_RWLOCK              = 6646 //获取存储侦测的读写锁配置
	NET_DVR_GET_STORAGEDETECTION_RWLOCK_CAPABILITIES = 6647 //获取存储侦测的读写锁配置能力
	NET_DVR_SET_STORAGEDETECTION_RWLOCK              = 6648 //设置存储侦测的读写锁配置
	NET_DVR_GET_PTZTRACKSTATUS                       = 6649 //获取球机联动跟踪状态

	NET_DVR_SET_STORAGEDETECTION_UNLOCK              = 6653 //设置存储侦测的解锁配置
	NET_DVR_GET_STORAGEDETECTION_UNLOCK_CAPABILITIES = 6654 //获取存储侦测的解锁配置能力

	NET_DVR_SET_SHIPSDETECTION_CFG          = 6655 //设置船只检测参数配置
	NET_DVR_GET_SHIPSDETECTION_CFG          = 6656 //获取船只检测参数配置
	NET_DVR_GET_SHIPSDETECTION_CAPABILITIES = 6657 //获取船只检测参数配置能力
	NET_DVR_GET_SHIPSDETECTION_COUNT        = 6658 //获取船只计数信息
	NET_DVR_SHIPSCOUNT_DELETE_CTRL          = 6659 //清空船只计数信息

	NET_DVR_GET_BAREDATAOVERLAY_CAPABILITIES         = 6660 //获取裸数据叠加能力
	NET_DVR_SET_BAREDATAOVERLAY_CFG                  = 6661 //设置裸数据叠加
	NET_DVR_GET_BAREDATAOVERLAY_CFG                  = 6662 //获取裸数据叠加
	NET_DVR_GET_SHIPSDETECTION_SCHEDULE              = 6663 //获取船只检测布防时间配置
	NET_DVR_SET_SHIPSDETECTION_SCHEDULE              = 6664 //设置船只检测布防时间配置
	NET_DVR_GET_SHIPSDETECTION_EVENT_TRIGGER         = 6665 //获取船只检测联动配置
	NET_DVR_SET_SHIPSDETECTION_EVENT_TRIGGER         = 6666 //设置船只检测联动配置
	NET_DVR_GET_SHIPSDETECTION_SCHEDULE_CAPABILITIES = 6667 //获取船只检测布防时间能力

	NET_DVR_FIRE_FOCUSZOOM_CTRL = 6670 //火点可见光镜头聚焦变倍

	NET_DVR_GET_FIREDETECTION_SCHEDULE_CAPABILITIES = 6671 //获取火点检测布防时间能力
	NET_DVR_GET_FIREDETECTION_SCHEDULE              = 6672 //获取火点检测布防时间配置
	NET_DVR_SET_FIREDETECTION_SCHEDULE              = 6673 //设置火点检测布防时间配置
	NET_DVR_GET_MANUALRANGING_CAPABILITIES          = 6675 //获取手动测距配置能力
	NET_DVR_SET_MANUALRANGING                       = 6677 //设置手动测距参数
	NET_DVR_GET_MANUALDEICING_CAPABILITIES          = 6678 //获取手动除冰配置能力
	NET_DVR_SET_MANUALDEICING                       = 6679 //设置手动除冰
	NET_DVR_GET_MANUALDEICING                       = 6680 //获取手动除冰

	NET_DVR_GET_THERMALPOWER_CAPABILITIES  = 6689 //获取相机电源配置能力
	NET_DVR_GET_THERMALPOWER               = 6690 //获取相机电源配置参数
	NET_DVR_SET_THERMALPOWER               = 6691 //设置相机电源配置参数
	NET_DVR_GET_PTZABSOLUTEEX_CAPABILITIES = 6695 //获取高精度PTZ绝对位置配置扩展能力
	NET_DVR_GET_PTZABSOLUTEEX              = 6696 //获取高精度PTZ绝对位置配置扩展
	NET_DVR_SET_PTZABSOLUTEEX              = 6697 //设置高精度PTZ绝对位置配置扩展

	NET_DVR_GET_CRUISE_CAPABILITIES    = 6698 //获取设备巡航模式配置能力
	NET_DVR_GET_CRUISE_INFO            = 6699 //获取设备巡航模式
	NET_DVR_GET_TEMP_HUMI_CAPABILITIES = 6700 //温湿度实时能力获取
	NET_DVR_GET_TEMP_HUMI_INFO         = 6701 //温湿度实时获取

	NET_DVR_GET_MANUALTHERM_INFO         = 6706 //手动测温实时获取
	NET_DVR_GET_MANUALTHERM_CAPABILITIES = 6707 //获取手动测温实时数据能力
	NET_DVR_SET_MANUALTHERM              = 6708 //设置手动测温数据设置

	//DVR96000
	NET_DVR_GET_ACCESSORY_CARD_INFO_CAPABILITIES = 6709 //获取配件板信息能力
	NET_DVR_GET_ACCESSORY_CARD_INFO              = 6710 //获取配件板信息

	NET_DVR_GET_THERMINTELL_CAPABILITIES       = 6711 //获取热成像智能互斥能力
	NET_DVR_GET_THERMINTELL                    = 6712 //获取热成像智能互斥配置参数
	NET_DVR_SET_THERMINTELL                    = 6713 //设置热成像智能互斥配置参数
	NET_GET_CRUISEPOINT_V50                    = 6714 //获取巡航路径配置扩展
	NET_DVR_GET_MANUALTHERM_BASIC_CAPABILITIES = 6715 //获取手动测温基本参数配置能力
	NET_DVR_SET_MANUALTHERM_BASICPARAM         = 6716 //设置手动测温基本参数
	NET_DVR_GET_MANUALTHERM_BASICPARAM         = 6717 //获取手动测温基本参数

	NET_DVR_GET_FIRESHIELDMASK_CAPABILITIES = 6718 //获取火点区域屏蔽能力

	NET_DVR_GET_HIDDEN_INFORMATION_CAPABILITIES = 6720 //隐藏信息配置能力
	NET_DVR_GET_HIDDEN_INFORMATION              = 6721 //获取隐藏信息参数
	NET_DVR_SET_HIDDEN_INFORMATION              = 6722 //设置隐藏信息参数

	NET_DVR_SET_FIRESHIELDMASK_CFG = 6723 //设置火点区域屏蔽参数
	NET_DVR_GET_FIRESHIELDMASK_CFG = 6724 //获取火点区域屏蔽参数

	NET_DVR_GET_SMOKESHIELDMASK_CAPABILITIES = 6725 //获取烟雾区域屏蔽能力
	NET_DVR_SET_SMOKESHIELDMASK_CFG          = 6726 //设置烟雾区域屏蔽参数
	NET_DVR_GET_SMOKESHIELDMASK_CFG          = 6727 //获取烟雾区域屏蔽参数

	NET_DVR_GET_AREASCAN_CAPABILITIES = 6728 //获取区域扫描能力
	NET_DVR_GET_AREASCAN_CFG          = 6730 //获取区域扫描参数

	NET_DVR_DEL_AREASCAN_CFG         = 6731 //扫描区域删除
	NET_DVR_AREASCAN_INIT_CTRL       = 6732 //进入区域扫描设置
	NET_DVR_AREASCAN_CONFIRM_CTRL    = 6733 //保存区域扫描设置
	NET_DVR_AREASCAN_STOP_CTRL       = 6734 //停止区域扫描设置
	NET_DVR_SAVE_SCANZOOM_CTRL       = 6735 //设置扫描倍率；保存当前光学倍率为扫描倍率
	NET_DVR_GET_SCANZOOM_CTRL        = 6736 //获取扫描倍率；将预览界面中的光学倍率返回到当前扫描倍率。
	NET_DVR_DEL_FIRESHIELDMASK_CTRL  = 6737 //删除火点屏蔽区域
	NET_DVR_DEL_SMOKESHIELDMASK_CTRL = 6738 //删除烟雾屏蔽区域

	NET_DVR_GET_DENSEFOG_EVENT_TRIGGER         = 6740 //获取大雾检测联动配置
	NET_DVR_SET_DENSEFOG_EVENT_TRIGGER         = 6741 //设置大雾检测联动配置
	NET_DVR_SET_DENSEFOGDETECTION_CFG          = 6742 //设置大雾检测参数配置
	NET_DVR_GET_DENSEFOGDETECTION_CFG          = 6743 //获取大雾检测参数配置
	NET_DVR_GET_DENSEFOGDETECTION_CAPABILITIES = 6744 //获取大雾检测参数配置能力

	NET_DVR_GET_THERMOMETRY_SCHEDULE_CAPABILITIES = 6750 //获取测温检测布防时间能力
	NET_DVR_GET_THERMOMETRY_SCHEDULE              = 6751 //获取测温检测布防时间配置
	NET_DVR_SET_THERMOMETRY_SCHEDULE              = 6752 //设置测温检测布防时间配置
	NET_DVR_GET_TEMPERTURE_SCHEDULE_CAPABILITIES  = 6753 //获取温差布防时间能力
	NET_DVR_GET_TEMPERTURE_SCHEDULE               = 6754 //获取温差布防时间配置
	NET_DVR_SET_TEMPERTURE_SCHEDULE               = 6755 //设置温差布防时间配置
	NET_DVR_GET_SEARCH_LOG_CAPABILITIES           = 6756 //日志类型支持能力
	NET_DVR_GET_VEHICLEFLOW                       = 6758 //获取车流量数据
	NET_DVR_GET_IPADDR_FILTERCFG_V50              = 6759 //获取IP地址过滤参数扩展
	NET_DVR_SET_IPADDR_FILTERCFG_V50              = 6760 //设置IP地址过滤参数扩展
	NET_DVR_GET_TEMPHUMSENSOR_CAPABILITIES        = 6761 //获取温湿度传感器的能力
	NET_DVR_GET_TEMPHUMSENSOR                     = 6762 //获取温湿度传感器配置协议
	NET_DVR_SET_TEMPHUMSENSOR                     = 6763 //设置温湿度传感器配置协议

	NET_DVR_GET_THERMOMETRY_MODE_CAPABILITIES = 6764 //获取测温模式能力
	NET_DVR_GET_THERMOMETRY_MODE              = 6765 //获取测温模式参数
	NET_DVR_SET_THERMOMETRY_MODE              = 6766 //设置测温模式参数

	NET_DVR_GET_THERMAL_PIP_CAPABILITIES              = 6767 //获取热成像画中画配置能力
	NET_DVR_GET_THERMAL_PIP                           = 6768 //获取热成像画中画配置参数
	NET_DVR_SET_THERMAL_PIP                           = 6769 //设置热成像画中画配置参数
	NET_DVR_GET_THERMAL_INTELRULEDISPLAY_CAPABILITIES = 6770 //获取热成像智能规则显示能力
	NET_DVR_GET_THERMAL_INTELRULE_DISPLAY             = 6771 //获取热成像智能规则显示参数
	NET_DVR_SET_THERMAL_INTELRULE_DISPLAY             = 6772 //设置热成像智能规则显示参数
	NET_DVR_GET_THERMAL_ALGVERSION                    = 6773 //获取热成像相关算法库版本
	NET_DVR_GET_CURRENT_LOCK_CAPABILITIES             = 6774 //获取电流锁定配置能力
	NET_DVR_GET_CURRENT_LOCK                          = 6775 //获取电流锁定配置参数
	NET_DVR_SET_CURRENT_LOCK                          = 6776 //设置电流锁定配置参数

	NET_DVR_DEL_MANUALTHERM_RULE = 6778 //删除手动测温规则

	NET_DVR_GET_UPGRADE_INFO = 6779 //获取升级信息

	NET_DVR_SWITCH_TRANSFER = 7000

	NET_DVR_GET_MB_POWERCTRLPARA   = 8000 //获取启动控制参数
	NET_DVR_SET_MB_POWERCTRLPARA   = 8001 //设置启动控制参数
	NET_DVR_GET_AUTOBACKUPPARA     = 8002 //获取自动备份参数
	NET_DVR_SET_AUTOBACKUPPARA     = 8003 //设置自动备份参数
	NET_DVR_GET_MB_GPSPARA         = 8004 //获取GPS参数
	NET_DVR_SET_MB_GPSPARA         = 8005 //设置GPS参数
	NET_DVR_GET_MB_SENSORINPARA    = 8006 //获取SENSOR参数
	NET_DVR_SET_MB_SENSORINPARA    = 8007 //设置SENSOR参数
	NET_DVR_GET_GSENSORPARA        = 8008 //获取GSENSOR参数
	NET_DVR_SET_GSENSORPARA        = 8009 //设置GSENSOR参数
	NET_DVR_GET_MB_DOWNLOADSVRPARA = 8010 //获取下载服务器参数
	NET_DVR_SET_MB_DOWNLOADSVRPARA = 8011 //设置下载服务器参数
	NET_DVR_GET_PLATERECOG_PARA    = 8012 //获取车牌识别参数
	NET_DVR_SET_PLATERECOG_PARA    = 8013 //设置车牌识别参数
	NET_DVR_GET_ENFORCESYS_PARA    = 8014 //获取车辆稽查参数
	NET_DVR_SET_ENFORCESYS_PARA    = 8015 //设置车辆稽查参数
	NET_DVR_GET_GPS_DATA           = 8016 //获取GPS数据
	NET_DVR_GET_ANALOG_ALARMINCFG  = 8017 //获取模拟报警输入参数
	NET_DVR_SET_ANALOG_ALARMINCFG  = 8018 //设置模拟报警输入参数

	NET_DVR_GET_SYSTEM_CAPABILITIES   = 8100 //获取设备的系统能力
	NET_DVR_GET_EAGLEEYE_CAPABILITIES = 8101 //获取设备鹰眼能力
	NET_DVR_GET_SLAVECAMERA_CALIB_V51 = 8102 //获取从摄像机标定配置V51
	NET_DVR_SET_SLAVECAMERA_CALIB_V51 = 8103 //设置从摄像机标定配置V51
	NET_DVR_SET_GOTOSCENE             = 8105 //设置主摄像机转向指定的场景ID

	//I、K、E系列NVR产品升级
	NET_DVR_GET_PTZ_NOTIFICATION = 8201 //获取多通道事件联动PTZ
	NET_DVR_SET_PTZ_NOTIFICATION = 8202 //设置多通道事件联动PTZ
	/*****************************电视墙 start****************************/
	NET_DVR_MATRIX_WALL_SET          = 9001 //设置电视墙中屏幕参数
	NET_DVR_MATRIX_WALL_GET          = 9002 //获取电视墙中屏幕参数
	NET_DVR_WALLWIN_GET              = 9003 //电视墙中获取窗口参数
	NET_DVR_WALLWIN_SET              = 9004 //电视墙中设置窗口参数
	NET_DVR_WALLWINPARAM_SET         = 9005 //设置电视墙窗口相关参数
	NET_DVR_WALLWINPARAM_GET         = 9006 //获取电视墙窗口相关参数
	NET_DVR_WALLSCENEPARAM_GET       = 9007 //设置场景模式参数
	NET_DVR_WALLSCENEPARAM_SET       = 9008 //获取场景模式参数
	NET_DVR_MATRIX_GETWINSTATUS      = 9009 //获取窗口解码状态
	NET_DVR_GET_WINASSOCIATEDDEVINFO = 9010 //电视墙中获取对应资源信息
	NET_DVR_WALLOUTPUT_GET           = 9011 //电视墙中获取显示输出参数
	NET_DVR_WALLOUTPUT_SET           = 9012 //电视墙中设置显示输出参数
	NET_DVR_GET_UNITEDMATRIXSYSTEM   = 9013 //电视墙中获取对应资源
	NET_DVR_GET_WALL_CFG             = 9014 //获取电视墙全局参数
	NET_DVR_SET_WALL_CFG             = 9015 //设置电视墙全局参数
	NET_DVR_CLOSE_ALL_WND            = 9016 //关闭所有窗口
	NET_DVR_SWITCH_WIN_TOP           = 9017 //窗口置顶
	NET_DVR_SWITCH_WIN_BOTTOM        = 9018 //窗口置底

	NET_DVR_CLOSE_ALL_WND_V41        = 9019 //电视墙关闭所有窗口v41（有多个电视墙）
	NET_DVR_GET_WALL_WINDOW_V41      = 9020 //获取电视墙中的窗口v41
	NET_DVR_SET_WALL_WINDOW_V41      = 9021 //设置电视墙中的窗口v41
	NET_DVR_GET_CURRENT_SCENE_V41    = 9022 //获取当前电视墙中正在使用的场景v41
	NET_DVR_GET_WALL_SCENE_PARAM_V41 = 9023 //获取当前电视墙中正在使用的场景v41
	NET_DVR_SET_WALL_SCENE_PARAM_V41 = 9024 //设置当前电视墙中正在使用的场景v41
	NET_DVR_GET_MATRIX_LOGO_CFG      = 9025 //获取logo参数
	NET_DVR_SET_MATRIX_LOGO_CFG      = 9026 //设置logo参数
	NET_DVR_GET_WIN_LOGO_CFG         = 9027 //获取窗口logo参数
	NET_DVR_SET_WIN_LOGO_CFG         = 9028 //设置窗口logo参数
	NET_DVR_DELETE_LOGO              = 9029 //删除logo
	NET_DVR_SET_DISPLAY_EFFECT_CFG   = 9030 //设置显示输出效果参数v41
	NET_DVR_GET_DISPLAY_EFFECT_CFG   = 9031 //获取显示输出效果参数v41
	NET_DVR_DEC_PLAY_REMOTE_FILE     = 9032 //解码播放远程文件
	NET_DVR_DEC_PLAY_REMOTE_FILE_V50 = 9314 //解码播放远程文件V50
	NET_DVR_GET_WIN_ZOOM_STATUS      = 9033 //获取窗口电子放大状态
	NET_DVR_GET_ALL_MATRIX_LOGOCFG   = 9034 //获取所有logo参数

	/*****************************电视墙 end******************************/

	/*******************************LCD拼接屏 begin******************************************/
	NET_DVR_SIMULATE_REMOTE_CONTROL  = 9035 //模拟遥控按键 2013-09-05
	NET_DVR_SET_SCREEN_SIGNAL_CFG    = 9036 //设置屏幕信号源参数
	NET_DVR_GET_SCREEN_SIGNAL_CFG    = 9037 //获取屏幕信号源参数
	NET_DVR_SET_SCREEN_SPLICE_CFG    = 9038 //设置屏幕拼接
	NET_DVR_GET_SCREEN_SPLICE_CFG    = 9039 //获取屏幕拼接
	NET_DVR_GET_SCREEN_FAN_WORK_MODE = 9040 //获取风扇工作方式
	NET_DVR_SET_SCREEN_FAN_WORK_MODE = 9041 //设置风扇工作方式
	NET_DVR_SHOW_SCREEN_WORK_STATUS  = 9044 //显示屏幕状态
	NET_DVR_GET_VGA_CFG              = 9045 //获取VGA信号配置
	NET_DVR_SET_VGA_CFG              = 9046 //设置VGA信号配置
	NET_DVR_GET_SCREEN_MENU_CFG      = 9048 //获取屏幕菜单配置
	NET_DVR_SET_SCREEN_MENU_CFG      = 9049 //设置屏幕菜单配置
	NET_DVR_SET_SCREEN_DISPLAY_CFG   = 9050 //设置显示参数 2013-08-28
	NET_DVR_GET_SCREEN_DISPLAY_CFG   = 9051 //获取显示参数 2013-08-28

	NET_DVR_SET_FUSION_CFG = 9052 //设置图像融合参数
	NET_DVR_GET_FUSION_CFG = 9053 //获取图像融合参数

	NET_DVR_SET_PIP_CFG             = 9060 //设置画中画参数
	NET_DVR_GET_PIP_CFG             = 9061 //获取画中画参数
	NET_DVR_SET_DEFOG_LCD           = 9073 //设置透雾参数
	NET_DVR_GET_DEFOG_LCD           = 9074 //获取透雾参数
	NET_DVR_SHOW_IP                 = 9075 //显示IP
	NET_DVR_SCREEN_MAINTENANCE_WALL = 9076 //屏幕维墙
	NET_DVR_SET_SCREEN_POS          = 9077 //设置屏幕位置参数
	NET_DVR_GET_SCREEN_POS          = 9078 //获取屏幕位置参数
	/*******************************LCD拼接屏 end******************************************/

	/*******************************LCD拼接屏V1.2 begin******************************************/
	NET_DVR_SCREEN_INDEX_SET             = 9079 //屏幕索引相关参数设置
	NET_DVR_SCREEN_INDEX_GET             = 9080 //屏幕索引相关参数获取
	NET_DVR_SCREEN_SPLICE_SET            = 9081 //设置屏幕拼接参数
	NET_DVR_SCREEN_SPLICE_GET            = 9082 //获取屏幕拼接参数
	NET_DVR_SET_SCREEN_PARAM             = 9083 //设置屏幕相关参数
	NET_DVR_GET_SCREEN_PARAM             = 9084 //获取屏幕相关参数
	NET_DVR_SET_SWITCH_CFG               = 9085 //设置定时开关机参数
	NET_DVR_GET_SWITCH_CFG               = 9086 //获取定时开关机参数
	NET_DVR_SET_POWERON_DELAY_CFG        = 9087 //设置延时开机参数
	NET_DVR_GET_POWERON_DELAY_CFG        = 9088 //获取延时开机参数
	NET_DVR_SET_SCREEN_POSITION          = 9089 //设置屏幕位置参数
	NET_DVR_GET_SCREEN_POSITION          = 9090 //获取屏幕位置参数
	NET_DVR_SCREEN_SCENE_CONTROL         = 9091 //屏幕场景控制
	NET_DVR_GET_CURRENT_SCREEN_SCENE     = 9092 //获取当前屏幕场景号
	NET_DVR_GET_SCREEN_SCENE_PARAM       = 9093 //获取屏幕场景模式参数
	NET_DVR_SET_SCREEN_SCENE_PARAM       = 9094 //设置屏幕场景模式参数
	NET_DVR_GET_EXTERNAL_MATRIX_RELATION = 9095 //获取外接矩阵输入输出关联关系
	NET_DVR_GET_LCD_AUDIO_CFG            = 9096 //获取LCD屏幕音频参数
	NET_DVR_SET_LCD_AUDIO_CFG            = 9097 //设置LCD屏幕音频参数
	NET_DVR_GET_LCD_WORK_STATE           = 9098 //获取LCD屏幕工作状态
	NET_DVR_GET_BOOT_LOGO_CFG            = 9099 //获取LCD屏幕开机logo显示参数
	NET_DVR_SET_BOOT_LOGO_CFG            = 9100 //设置LCD屏幕开机logo显示参数

	/*******************************LCD拼接屏V1.2 end ******************************************/
	NET_DVR_GET_STREAM_DST_COMPRESSIONINFO = 9101 //获取目标压缩参数
	NET_DVR_SET_STREAM_DST_COMPRESSIONINFO = 9102 //设置目标压缩参数
	NET_DVR_GET_STREAM_TRANS_STATUS        = 9103 //获取流状态
	NET_DVR_GET_DEVICE_TRANS_STATUS        = 9104 //获取设备转码状态
	NET_DVR_GET_ALLSTREAM_SRC_INFO         = 9105 //获取所有流信息
	NET_DVR_GET_BIG_SCREEN_AUDIO           = 9106 //获取大屏音频信息
	NET_DVR_SET_BIG_SCREEN_AUDIO           = 9107 //设置大屏音频信息
	NET_DVR_GET_DEV_WORK_MODE              = 9108 //获取转码设备工作模式
	NET_DVR_SET_DEV_WORK_MODE              = 9109 //设置转码设备工作模式
	NET_DVR_APPLY_TRANS_CHAN               = 9110 //按流ID申请转码通道
	NET_DVR_GET_DISPCHAN_CFG               = 9111 //批量获取显示通道参数
	NET_DVR_SET_DISPCHAN_CFG               = 9112 //批量设置显示通道参数

	NET_DVR_GET_DEC_CHAN_STATUS   = 9113 //获取解码通道解码状态
	NET_DVR_GET_DISP_CHAN_STATUS  = 9114 //获取显示通道状态
	NET_DVR_GET_ALARMIN_STATUS    = 9115 //获取报警输入状态
	NET_DVR_GET_ALARMOUT_STATUS   = 9116 //获取报警输出状态
	NET_DVR_GET_AUDIO_CHAN_STATUS = 9117 //获取语音对讲状态

	NET_DVR_GET_VIDEO_AUDIOIN_CFG = 9118 //获取视频的音频输入参数
	NET_DVR_SET_VIDEO_AUDIOIN_CFG = 9119 //设置视频的音频输入参数

	NET_DVR_SET_BASEMAP_CFG         = 9120 //设置底图参数
	NET_DVR_GET_BASEMAP_CFG         = 9121 //获取底图参数
	NET_DVR_GET_VIRTUAL_SCREEN_CFG  = 9122 //获取超高清输入子系统参数
	NET_DVR_SET_VIRTUAL_SCREEN_CFG  = 9123 //设置超高清输入子系统参数
	NET_DVR_GET_BASEMAP_WIN_CFG     = 9124 //获取底图窗口参数
	NET_DVR_SET_BASEMAP_WIN_CFG     = 9125 //设置底图窗口参数
	NET_DVR_DELETE_PICTURE          = 9126 //删除底图
	NET_DVR_GET_BASEMAP_PIC_INFO    = 9127 //获取底图图片信息
	NET_DVR_SET_BASEMAP_WIN_CFG_V40 = 9128 //设置底图窗口参数V40
	NET_DVR_GET_BASEMAP_WIN_CFG_V40 = 9129 //获取底图窗口参数V40

	NET_DVR_GET_DEC_VCA_CFG = 9130 //获取解码器智能报警参数
	NET_DVR_SET_DEC_VCA_CFG = 9131 //设置解码器智能报警参数

	NET_DVR_SET_VS_INPUT_CHAN_INIT_ALL = 9132 //初始化虚拟屏子板的所有输入通道
	NET_DVR_GET_VS_INPUT_CHAN_INIT_ALL = 9133 //获取虚拟屏子板的所有输入通道的初始化参数
	NET_DVR_GET_VS_INPUT_CHAN_INIT     = 9134 //获取虚拟屏输入通道的初始化参数
	NET_DVR_GET_VS_INPUT_CHAN_CFG      = 9135 //获取虚拟屏输入通道配置参数

	NET_DVR_GET_TERMINAL_CONFERENCE_STATUS = 9136 //获取终端会议状态
	NET_DVR_GET_TERMINAL_INPUT_CFG_CAP     = 9137 //获取终端输入参数能力
	NET_DVR_GET_TERMINAL_INPUT_CFG         = 9138 //获取终端视频会议输入参数
	NET_DVR_SET_TERMINAL_INPUT_CFG         = 9139 //设置终端视频会议输入参数

	NET_DVR_GET_CONFERENCE_REGION_CAP = 9140 //获取终端会议区域能力
	NET_DVR_GET_CONFERENCE_REGION     = 9141 //获取终端会议区域参数
	NET_DVR_SET_CONFERENCE_REGION     = 9142 //设置终端会议区域参数
	NET_DVR_GET_TERMINAL_CALL_CFG_CAP = 9143 //获取终端呼叫配置能力
	NET_DVR_GET_TERMINAL_CALL_CFG     = 9144 //获取终端呼叫参数
	NET_DVR_SET_TERMINAL_CALL_CFG     = 9145 //设置终端呼叫参数
	NET_DVR_GET_TERMINAL_CTRL_CAP     = 9146 //获取终端呼叫控制能力
	NET_DVR_TERMINAL_CTRL             = 9147 //终端呼叫控制
	NET_DVR_GET_CALL_QUERY_CAP        = 9148 //获取会议查找能力
	NET_DVR_GET_CALLINFO_BY_COND      = 9149 //按条件查询呼叫记录

	NET_DVR_SET_FUSION_SCALE = 9150 //设置图像融合规模
	NET_DVR_GET_FUSION_SCALE = 9151 //获取图像融合规模

	NET_DVR_GET_VCS_CAP = 9152 //获取MCU能力集

	NET_DVR_GET_TERMINAL_GK_CFG_CAP      = 9153 //获取终端注册GK能力
	NET_DVR_GET_TERMINAL_GK_CFG          = 9154 //获取终端注册GK参数
	NET_DVR_SET_TERMINAL_GK_CFG          = 9155 //设置终端注册GK参数
	NET_DVR_GET_MCU_CONFERENCESEARCH_CAP = 9156 //获取MCU设备的能力
	NET_DVR_SET_VS_INPUT_CHAN_CFG        = 9157 //设置虚拟屏输入通道配置参数
	NET_DVR_GET_VS_NETSRC_CFG            = 9158 //设置虚拟屏网络源参数
	NET_DVR_SET_VS_NETSRC_CFG            = 9159 //设置虚拟屏网络源参数

	NET_DVR_GET_LLDP_CFG                 = 9160 //获取LLDP参数
	NET_DVR_SET_LLDP_CFG                 = 9161 //设置LLDP参数
	NET_DVR_GET_LLDP_CAP                 = 9162 //获取LLDP能力集
	NET_DVR_GET_FIBER_CONVERT_BASIC_INFO = 9163 //获取光纤收发器基本信息
	NET_DVR_GET_FIBER_CONVERT_WORK_STATE = 9164 //获取光纤收发器工作状
	NET_DVR_GET_FIBER_CONVERT_TOPOLOGY   = 9165 //获取光纤收发器拓扑信息
	NET_DVR_GET_FC_PORT_REMARKS          = 9166 //获取光纤收发器端口注释参数
	NET_DVR_SET_FC_PORT_REMARKS          = 9167 //设置光纤收发器端口注释参数
	NET_DVR_GET_PORT_REMARKS_CAP         = 9168 //获取光纤收发器端口注释能力集

	NET_DVR_GET_MCU_CONFERENCECONTROL_CAP = 9169 //获取会议控制能力
	NET_DVR_GET_MCU_TERMINALCONTROL_CAP   = 9170 //获取终端控制能力
	NET_DVR_GET_MCU_TERIMINALGROUP_CAP    = 9171 //获取终端分组能力
	NET_DVR_GET_MCU_TERMINAL_CAP          = 9174 //获取终端管理能力
	NET_DVR_GET_MCU_CONFERENCE_CAP        = 9175 //获取会议能力
	NET_DVR_GET_MCU_GK_CFG_CAP            = 9176 //获取MCUGK配置能力
	NET_DVR_GET_MCU_GK_SERVER_CAP         = 9177 //获取MCUGK服务能力

	NET_DVR_GET_EDID_CFG_FILE_INFO      = 9178 //获取EDID配置文件信息
	NET_DVR_GET_EDID_CFG_FILE_INFO_LIST = 9179 //获取所有EDID配置文件信息
	NET_DVR_SET_EDID_CFG_FILE_INFO      = 9180 //设置EDID配置文件信息
	NET_DVR_DEL_EDID_CFG_FILE_INFO      = 9181 //删除EDID配置文件信息（包括文件）
	NET_DVR_GET_EDID_CFG_FILE_INFO_CAP  = 9182 //EDID配置文件信息能力集

	NET_DVR_GET_SUBWND_DECODE_OSD     = 9183 //获取子窗口解码OSD信息
	NET_DVR_GET_SUBWND_DECODE_OSD_ALL = 9184 //获取所有子窗口解码OSD信息
	NET_DVR_SET_SUBWND_DECODE_OSD     = 9185 //设置子窗口解码OSD信息
	NET_DVR_GET_SUBWND_DECODE_OSD_CAP = 9186 //获取子窗口解码OSD信息能力集
	NET_DVR_GET_DECODE_CHANNEL_OSD    = 9187 //获取解码通道OSD信息
	NET_DVR_SET_DECODE_CHANNEL_OSD    = 9188 //设置解码通道OSD信息

	NET_DVR_GET_OUTPUT_PIC_INFO         = 9200 //获取输出口图片参数
	NET_DVR_SET_OUTPUT_PIC_INFO         = 9201 //设置输出口图片参数
	NET_DVR_GET_OUTPUT_PIC_WIN_CFG      = 9202 //获取输出口图片窗口参数
	NET_DVR_SET_OUTPUT_PIC_WIN_CFG      = 9203 //设置输出口图片窗口参数
	NET_DVR_GET_OUTPUT_ALL_PIC_WIN_CFG  = 9204 //获取输出口所有图片窗口参数
	NET_DVR_DELETE_OUPUT_PIC            = 9205 //删除输出口图片
	NET_DVR_GET_OUTPUT_OSD_CFG          = 9206 //获取输出口OSD参数
	NET_DVR_SET_OUTPUT_OSD_CFG          = 9207 //设置输出口OSD参数
	NET_DVR_GET_OUTPUT_ALL_OSD_CFG      = 9208 //获取输出口所有OSD参数
	NET_DVR_GET_CHAN_RELATION           = 9209 //获取编码通道关联资源参数
	NET_DVR_SET_CHAN_RELATION           = 9210 //设置编码通道关联资源参数
	NET_DVR_GET_ALL_CHAN_RELATION       = 9211 //获取所有编码通道关联资源参数
	NET_DVR_GET_NS_RING_CFG             = 9212 //获取光纤板环网配置
	NET_DVR_SET_NS_RING_CFG             = 9213 //设置光纤板环网配置
	NET_DVR_GET_NS_RING_STATUS          = 9214 //获取光纤板环网状态
	NET_DVR_GET_OPTICAL_PORT_INFO       = 9220 //获取光口信息
	NET_DVR_SET_OPTICAL_PORT_INFO       = 9221 //设置光口信息
	NET_DVR_GET_OPTICAL_CHAN_RELATE_CFG = 9222 //获取编码通道关联光口输入源参数
	NET_DVR_SET_OPTICAL_CHAN_RELATE_CFG = 9223 //设置编码通道关联光口输入源参数
	NET_DVR_GET_WIN_ROAM_SWITCH_CFG     = 9224 //获取解码器窗口漫游开关参数
	NET_DVR_SET_WIN_ROAM_SWITCH_CFG     = 9225 //设置解码器窗口漫游开关参数
	NET_DVR_START_SCREEN_CRTL           = 9226 //开始屏幕控制
	NET_DVR_GET_SCREEN_FLIE_LIST        = 9227 //获取屏幕文件列表
	NET_DVR_GET_SCREEN_FILEINFO         = 9228 //获取屏幕文件信息参数
	NET_DVR_SET_SCREEN_FILEINFO         = 9229 //设置屏幕文件信息参数

	/*******************************小间距LED显示屏 begin***************************************/
	NET_DVR_GET_LED_OUTPUT_CFG         = 9230 //获取发送卡输出参数
	NET_DVR_SET_LED_OUTPUT_CFG         = 9231 //设置发送卡输出参数
	NET_DVR_GET_LED_OUTPUT_PORT_CFG    = 9232 //获取LED发送卡输出端口参数
	NET_DVR_SET_LED_OUTPUT_PORT_CFG    = 9233 //设置LED发送卡输出端口参数
	NET_DVR_GET_LED_DISPLAY_AREA_CFG   = 9234 //获取LED发送卡显示区域
	NET_DVR_SET_LED_DISPLAY_AREA_CFG   = 9235 //设置LED发送卡显示区域
	NET_DVR_GET_LED_PORT_CFG           = 9236 //获取LED发送卡端口参数
	NET_DVR_SET_LED_PORT_CFG           = 9237 //设置LED发送卡端口参数
	NET_DVR_GET_LED_DISPLAY_CFG        = 9238 //获取LED发送卡显示参数
	NET_DVR_SET_LED_DISPLAY_CFG        = 9239 //设置LED发送卡显示参数
	NET_DVR_GET_ALL_LED_PORT_CFG       = 9240 //获取LED发送卡某个输出对应
	NET_DVR_SAVE_LED_CONFIGURATION     = 9241 //参数固化
	NET_DVR_GET_LED_TEST_SIGNAL_CFG    = 9242 //获取LED屏测试信号参数
	NET_DVR_SET_LED_TEST_SIGNAL_CFG    = 9243 //设置LED屏测试信号参数
	NET_DVR_GET_LED_NOSIGNAL_CFG       = 9244 //获取LED屏无信号显示模式参数
	NET_DVR_SET_LED_NOSIGNAL_CFG       = 9245 //设置LED屏无信号显示模式参数
	NET_DVR_GET_LED_INPUT_CFG          = 9246 //获取LED发送卡输入参数
	NET_DVR_SET_LED_INPUT_CFG          = 9247 //设置LED发送卡输入参数
	NET_DVR_GET_LED_RECV_GAMMA_CFG     = 9248 //获取接收卡GAMMA表参数
	NET_DVR_SET_LED_RECV_GAMMA_CFG     = 9249 //设置接收卡GAMMA表参数
	NET_DVR_GET_LED_RECV_CFG           = 9250 //获取接收卡基本参数
	NET_DVR_SET_LED_RECV_CFG           = 9251 //设置接收卡基本参数
	NET_DVR_GET_LED_RECV_ADVANCED_CFG  = 9252 //获取接收卡高级参数
	NET_DVR_SET_LED_RECV_ADVANCED_CFG  = 9253 //设置接收卡高级参数
	NET_DVR_GET_LED_SCREEN_DISPLAY_CFG = 9254 //获取LED屏显示参数
	NET_DVR_SET_LED_SCREEN_DISPLAY_CFG = 9255 //设置LED屏显示参数
	/*******************************小间距LED显示屏 end*****************************************/

	NET_DVR_GET_INSERTPLAY_PROGRESS = 9273 //获取插播进度

	NET_DVR_GET_SCREEN_CONFIG     = 9260 //获取屏幕服务器参数
	NET_DVR_SET_SCREEN_CONFIG     = 9261 //设置屏幕服务器参数
	NET_DVR_GET_SCREEN_CONFIG_CAP = 9262 //获取屏幕服务器参数能力集

	NET_DVR_GET_SCHEDULE_PUBLISH_PROGRESS = 9271 //获取日程发布进度
	NET_DVR_GET_PUBLISH_UPGRADE_PROGRESS  = 9272 //获取信息发布终端升级进度

	NET_DVR_GET_INPUT_BOARD_CFG      = 9281 //获取输入板配置信息
	NET_DVR_GET_INPUT_BOARD_CFG_LIST = 9282 //获取输入板配置信息列表
	NET_DVR_SET_INPUT_BOARD_CFG      = 9283 //设置输入板配置信息

	NET_DVR_GET_INPUT_SOURCE_TEXT_CAP            = 9284 //获取输入源字符叠加能力
	NET_DVR_GET_INPUT_SOURCE_TEXT_CFG            = 9285 //获取输入源字符叠加参数
	NET_DVR_GET_INPUT_SOURCE_TEXT_CFG_LSIT       = 9286 //获取输入源字符叠加参数列表
	NET_DVR_SET_INPUT_SOURCE_TEXT_CFG            = 9287 //设置输入源字符叠加参数
	NET_DVR_SET_INPUT_SOURCE_TEXT_CFG_LIST       = 9288 //设置输入源字符叠加参数列表
	NET_DVR_GET_INPUT_SOURCE_RESOLUTION_CAP      = 9289 //获取输入源自定义分辨率能力
	NET_DVR_GET_INPUT_SOURCE_RESOLUTION_CFG      = 9290 //获取输入源自定义分辨率参数
	NET_DVR_GET_INPUT_SOURCE_RESOLUTION_CFG_LIST = 9291 //获取输入源自定义分辨率列表
	NET_DVR_SET_INPUT_SOURCE_RESOLUTION_CFG      = 9292 //设置输入源自定义分辨率参数
	NET_DVR_SET_INPUT_SOURCE_RESOLUTION_CFG_LIST = 9293 //设置输入源自定义分辨率参数
	NET_DVR_GET_LED_AREA_INFO_LIST               = 9295 //获取LED区域列表

	NET_DVR_GET_DISPINPUT_CFG      = 9296 //获取显示输入参数
	NET_DVR_GET_DISPINPUT_CFG_LIST = 9297 //获取所有显示输入参数
	NET_DVR_SET_DISPINPUT_CFG      = 9298 //设置显示输入参数
	NET_DVR_GET_DISPINPUT_CFG_CAP  = 9299 //获取显示输入参数能力集
	NET_DVR_GET_CURRENT_VALID_PORT = 9300 //获取当前有效的,可以连接的端口

	NET_DVR_SET_ONLINE_UPGRADE         = 9301 //允许在线升级
	NET_DVR_GET_ONLINEUPGRADE_PROGRESS = 9302 //获取在线升级进度
	NET_DVR_GET_FIRMWARECODE           = 9303 //获取识别码
	NET_DVR_GET_ONLINEUPGRADE_SERVER   = 9304 //获取升级服务器状态
	NET_DVR_GET_ONLINEUPGRADE_VERSION  = 9305 //获取新版本信息
	NET_DVR_GET_RECOMMEN_VERSION       = 9306 //检测是否推荐升级到此版本
	NET_DVR_GET_ONLINEUPGRADE_ABILITY  = 9309 //获取在线升级能力集

	NET_DVR_GET_FIBER_CONVERT_BASIC_INFO_V50 = 9310 //获取远端网管收发器基本信息V50
	NET_DVR_GET_FIBER_CONVERT_WORK_STATE_V50 = 9311 //获取远端网管收发器工作状态

	NET_SDK_LED_SCREEN_CHECK        = 9312 //LED屏幕校正
	NET_SDK_GENERATE_OUTPUT_CONTROL = 9315 //通用扩展输出口模块控制
	NET_SDK_GET_MATRIX_STATUS_V51   = 9313 /*获取视频综合平台状态*/

	//DS-19D2000-S V2.0升级 报警联动配置参数命令码
	NET_DVR_GET_ALARM_LINKAGE_CFG = 9316 //获取动环报警联动配置参数
	NET_DVR_SET_ALARM_LINKAGE_CFG = 9317 //设置动环报警联动配置参数

	NET_DVR_GET_RS485_WORK_MODE         = 10001 //获取RS485串口工作模式
	NET_DVR_SET_RS485_WORK_MODE         = 10002 //设置RS485串口工作模式
	NET_DVR_GET_SPLITTER_TRANS_CHAN_CFG = 10003 //获取码分器透明通道参数
	NET_DVR_SET_SPLITTER_TRANS_CHAN_CFG = 10004 //设置码分器透明通道参数

	NET_DVR_GET_RS485_PROTOCOL_VERSION  = 10301 //获取RS485协议库版本信息
	NET_DVR_ALARMHOST_REGISTER_DETECTOR = 10302 //自动注册探测器

	NET_DVR_GET_SIP_CFG               = 11001 //IP可视化机获取SIP参数
	NET_DVR_SET_SIP_CFG               = 11002 //IP可视化机设置SIP参数
	NET_DVR_GET_IP_VIEW_DEVCFG        = 11003 //获取IP对讲分机配置
	NET_DVR_SET_IP_VIEW_DEVCFG        = 11004 //设置IP对讲分机配置
	NET_DVR_GET_IP_VIEW_AUDIO_CFG     = 11005 //获取IP对讲分机音频参数
	NET_DVR_SET_IP_VIEW_AUDIO_CFG     = 11006 //设置IP对讲分机音频参数
	NET_DVR_GET_IP_VIEW_CALL_CFG      = 11007 //获取IP对讲分机呼叫参数
	NET_DVR_SET_IP_VIEW_CALL_CFG      = 11008 //设置IP对讲分机呼叫参数
	NET_DVR_GET_AUDIO_LIMIT_ALARM_CFG = 11009 //获取声音超限配置参数
	NET_DVR_SET_AUDIO_LIMIT_ALARM_CFG = 11010 //设置声音超限配置参数
	NET_DVR_GET_BUTTON_DOWN_ALARM_CFG = 11011 //获取按钮按下告警配置参数
	NET_DVR_SET_BUTTON_DOWN_ALARM_CFG = 11012 //设置按钮按下告警配置参数

	NET_DVR_GET_ISCSI_CFG = 11070 // 获取ISCSI存储配置协议
	NET_DVR_SET_ISCSI_CFG = 11071 // 获取ISCSI存储配置协议

	NET_DVR_GET_SECURITYMODE = 12004 //获取当前安全模式
	//2013-11-21 获取设备当前的温度和湿度
	NET_DVR_GET_TEMP_HUMI = 12005

	//2014-02-15 民用IPC自动化测试项目
	NET_DVR_SET_ALARMSOUNDMODE = 12006 //设置报警声音模式
	NET_DVR_GET_ALARMSOUNDMODE = 12007 //获取报警声音模式

	NET_DVR_SET_IPDEVICE_ACTIVATED           = 13000 //通过NVR激活前端设备
	NET_DVR_GET_DIGITAL_CHAN_SECURITY_STATUS = 13001 //获取数字通道对应设备安全状态
	NET_DVR_GET_ACTIVATE_IPC_ABILITY         = 13003 //获取NVR激活IPC能力集

	/*******************************楼宇可视对讲机 start***********************************/
	NET_DVR_GET_VIDEO_INTERCOM_DEVICEID_CFG  = 16001 //获取可视对讲设备编号
	NET_DVR_SET_VIDEO_INTERCOM_DEVICEID_CFG  = 16002 //设置可视对讲设备编号
	NET_DVR_SET_PRIVILEGE_PASSWORD           = 16003 //设置权限密码配置信息
	NET_DVR_GET_OPERATION_TIME_CFG           = 16004 //获取操作时间配置
	NET_DVR_SET_OPERATION_TIME_CFG           = 16005 //设置操作时间配置
	NET_DVR_GET_VIDEO_INTERCOM_RELATEDEV_CFG = 16006 //获取关联网络设备参数
	NET_DVR_SET_VIDEO_INTERCOM_RELATEDEV_CFG = 16007 //设置关联网络设备参数
	NET_DVR_REMOTECONTROL_NOTICE_DATA        = 16008 //公告信息下发
	NET_DVR_REMOTECONTROL_GATEWAY            = 16009 //远程开锁
	NET_DVR_REMOTECONTROL_OPERATION_AUTH     = 16010 //操作权限验证

	NET_DVR_GET_VIDEO_INTERCOM_IOIN_CFG  = 16016 //获取IO输入参数
	NET_DVR_SET_VIDEO_INTERCOM_IOIN_CFG  = 16017 //设置IO输入参数
	NET_DVR_GET_VIDEO_INTERCOM_IOOUT_CFG = 16018 //获取IO输出参数
	NET_DVR_SET_VIDEO_INTERCOM_IOOUT_CFG = 16019 //设置IO输出参数
	NET_DVR_GET_ELEVATORCONTROL_CFG      = 16020 //获取梯控器参数
	NET_DVR_SET_ELEVATORCONTROL_CFG      = 16021 //设置梯控器参数
	NET_DVR_GET_VIDEOINTERCOM_STREAM     = 16022 //获取可视对讲流通道参数
	NET_DVR_SET_VIDEOINTERCOM_STREAM     = 16023 //设置可视对讲流通道参数
	NET_DVR_GET_WDR_CFG                  = 16024 //获取宽动态参数配置
	NET_DVR_SET_WDR_CFG                  = 16025 //设置宽动态参数配置
	NET_DVR_GET_VIS_DEVINFO              = 16026 //获取可设备编号信息
	NET_DVR_GET_VIS_REGISTER_INFO        = 16027 //获取可设备注册的设备信息
	NET_DVR_GET_ELEVATORCONTROL_CFG_V40  = 16028 //获取梯控器参数-扩展
	NET_DVR_SET_ELEVATORCONTROL_CFG_V40  = 16029 //设置梯控器参数-扩展
	NET_DVR_GET_CALL_ROOM_CFG            = 16030 //获取按键呼叫住户配置
	NET_DVR_SET_CALL_ROOM_CFG            = 16031 //设置按键呼叫住户配置
	NET_DVR_VIDEO_CALL_SIGNAL_PROCESS    = 16032 //可视话对讲信令处理
	NET_DVR_GET_CALLER_INFO              = 16033 //获取主叫长号信息
	NET_DVR_GET_CALL_STATUS              = 16034 //获取通话状态
	NET_DVR_GET_SERVER_DEVICE_INFO       = 16035 //获取设备列表
	NET_DVR_SET_CALL_SIGNAL              = 16036 //可视对讲手机端发送信令
	NET_DVR_GET_VIDEO_INTERCOM_ALARM_CFG = 16037 //获取可视对讲报警事件参数
	NET_DVR_SET_VIDEO_INTERCOM_ALARM_CFG = 16038 //设置可视对讲报警事件参数
	NET_DVR_GET_RING_LIST                = 16039 //查询铃音参数列表

	NET_DVR_GET_ROOM_CUSTOM_CFG         = 16040 //房间自定义获取
	NET_DVR_SET_ROOM_CUSTOM_CFG         = 16041 //房间自定义设置
	NET_DVR_GET_ELEVATORCONTROL_CFG_V50 = 16042 //获取梯控器参数V50
	NET_DVR_SET_ELEVATORCONTROL_CFG_V50 = 16043 //设置梯控器参数V50
	NET_DVR_GET_SIP_CFG_V50             = 16044 //获取SIP参数V50
	NET_DVR_SET_SIP_CFG_V50             = 16045 //设置SIP参数V50
	NET_DVR_GET_NOTICE_VIDEO_DATA       = 16050 //公告视频获取

	/*******************************楼宇可视对讲机 end***********************************/

	NET_DVR_DEBUGINFO_START = 18000 //网传设备调试信息启动命令
	NET_DVR_AUTO_TEST_START = 18001 //自动测试长连接获取

	NET_DVR_GET_SELFCHECK_RESULT = 20000 //获取设备自检结果
	NET_DVR_SET_TEST_COMMAND     = 20001 //设置测试控制命令
	NET_DVR_SET_TEST_DEVMODULE   = 20002 //设置测试硬件模块控制命令
	NET_DVR_GET_TEST_DEVMODULE   = 20003 //获取测试硬件模块控制命令

	NET_DVR_SET_AUTOFOCUS_TEST          = 20004 //保存自动对焦参数 2013-10-26
	NET_DVR_CHECK_USER_STATUS           = 20005 //检测用户是否在线
	NET_DVR_GET_TEST_COMMAND            = 20010 //获取测试控制命令
	NET_DVR_GET_DIAL_SWITCH_CFG         = 20200 //获取拨码开关信息
	NET_DVR_SET_AGING_TRICK_SCAN        = 20201 //设置老化前后工具参数
	NET_DVR_GET_ECCENTRIC_CORRECT_STATE = 20202 //获取获取偏心校正状态

	NET_DVR_GET_THERMOMETRYRULE_TEMPERATURE_INFO = 23001 //手动获取测温规则温度信息

	NET_DVR_T1_TEST_CMD = 131073 //当测试命令来用，通过数据区域的文本内容区分具体做什么.数据长度不得大于1024
	//数据区格式为：<T1TestCmd type="0"/>//恢复设备默认参数并关机。

	// 美分定制菜单输出模式外部命令
	NET_DVR_GET_MEMU_OUTPUT_MODE = 155649 // 获取菜单输出模式
	NET_DVR_SET_MEMU_OUTPUT_MODE = 155650 // 设置菜单输出模式

	/***************************DS9000新增命令(_V30) end *****************************/

	NET_DVR_GET_DEV_LOGIN_RET_INFO = 16777200 //设备登陆返回参数

	NET_DVR_GET_TEST_VERSION_HEAD    = 268435441 //获取测试版本头
	NET_DVR_SET_TEST_VERSION_HEAD    = 268435442 //设置测试版本头
	NET_DVR_GET_TEST_VERSION_HEAD_V1 = 268435443 //获取测试版本头-第二版
	NET_DVR_SET_TEST_VERSION_HEAD_V1 = 268435444 //设置测试版本头-第二版
	NET_DVR_GET_TEST_VERSION_HEAD_V2 = 268435445 //获取测试版本头-第三版
	NET_DVR_SET_TEST_VERSION_HEAD_V2 = 268435446 //设置测试版本头-第三版

	NET_DVR_GET_TEST_VERSION_HEAD_ONLY_0 = 268435447 //获取测试版本头,当前仅有一个版本
	NET_DVR_SET_TEST_VERSION_HEAD_ONLY_0 = 268435448 //设置测试版本头,当前仅有一个版本

	MAX_LOCAL_ADDR_LEN   = 96 //SOCKS最大本地网段个数
	MAX_COUNTRY_NAME_LEN = 4  //国家简写名称长度

	/************************DVR日志 begin***************************/

	/* 报警 */
	//主类型
	MAJOR_ALARM = 0x1
	//次类型
	MINOR_ALARM_IN         = 0x1  /* 报警输入 */
	MINOR_ALARM_OUT        = 0x2  /* 报警输出 */
	MINOR_MOTDET_START     = 0x3  /* 移动侦测报警开始 */
	MINOR_MOTDET_STOP      = 0x4  /* 移动侦测报警结束 */
	MINOR_HIDE_ALARM_START = 0x5  /* 遮挡报警开始 */
	MINOR_HIDE_ALARM_STOP  = 0x6  /* 遮挡报警结束 */
	MINOR_VCA_ALARM_START  = 0x7  /*智能报警开始*/
	MINOR_VCA_ALARM_STOP   = 0x8  /*智能报警停止*/
	MINOR_ITS_ALARM_START  = 0x09 // 交通事件报警开始
	MINOR_ITS_ALARM_STOP   = 0x0A // 交通事件报警结束
	//2010-11-10 网络报警日志
	MINOR_NETALARM_START = 0x0b /*网络报警开始*/
	MINOR_NETALARM_STOP  = 0x0c /*网络报警结束*/
	//2010-12-16 报警板日志，与"MINOR_ALARM_IN"配对使用
	MINOR_NETALARM_RESUME = 0x0d /*网络报警恢复*/
	//2012-4-5 IPC PIR、无线、呼救报警
	MINOR_WIRELESS_ALARM_START      = 0x0e /* 无线报警开始 */
	MINOR_WIRELESS_ALARM_STOP       = 0x0f /* 无线报警结束 */
	MINOR_PIR_ALARM_START           = 0x10 /* 人体感应报警开始 */
	MINOR_PIR_ALARM_STOP            = 0x11 /* 人体感应报警结束 */
	MINOR_CALLHELP_ALARM_START      = 0x12 /* 呼救报警开始 */
	MINOR_CALLHELP_ALARM_STOP       = 0x13 /* 呼救报警结束 */
	MINOR_IPCHANNEL_ALARMIN_START   = 0x14 //数字通道报警输入开始：PCNVR在接收到数字通道的MINOR_ALARM_IN产生“数字通道报警输入开始”，10s，再收不到MINOR_ALARM_IN，产生“数字通道报警输入结束”
	MINOR_IPCHANNEL_ALARMIN_STOP    = 0x15 //数字通道报警输入开始：同上
	MINOR_DETECTFACE_ALARM_START    = 0x16 /* 人脸侦测报警开始 */
	MINOR_DETECTFACE_ALARM_STOP     = 0x17 /* 人脸侦测报警结束 */
	MINOR_VQD_ALARM_START           = 0x18 //VQD报警
	MINOR_VQD_ALARM_STOP            = 0x19 //VQD报警结束
	MINOR_VCA_SECNECHANGE_DETECTION = 0x1a //场景侦测报警 2013-07-16

	MINOR_SMART_REGION_EXITING_BEGIN = 0x1b //离开区域侦测开始
	MINOR_SMART_REGION_EXITING_END   = 0x1c //离开区域侦测结束
	MINOR_SMART_LOITERING_BEGIN      = 0x1d //徘徊侦测开始
	MINOR_SMART_LOITERING_END        = 0x1e //徘徊侦测结束

	MINOR_VCA_ALARM_LINE_DETECTION_BEGIN = 0x20
	MINOR_VCA_ALARM_LINE_DETECTION_END   = 0x21
	MINOR_VCA_ALARM_INTRUDE_BEGIN        = 0x22 //区域侦测开始
	MINOR_VCA_ALARM_INTRUDE_END          = 0x23 //区域侦测结束
	MINOR_VCA_ALARM_AUDIOINPUT           = 0x24 //音频异常输入
	MINOR_VCA_ALARM_AUDIOABNORMAL        = 0x25 //声强突变
	MINOR_VCA_DEFOCUS_DETECTION_BEGIN    = 0x26 //虚焦侦测开始
	MINOR_VCA_DEFOCUS_DETECTION_END      = 0x27 //虚焦侦测结束

	//民用NVR
	MINOR_EXT_ALARM                    = 0x28 /*IPC外部报警*/
	MINOR_VCA_FACE_ALARM_BEGIN         = 0x29 /*人脸侦测开始*/
	MINOR_SMART_REGION_ENTRANCE_BEGIN  = 0x2a //进入区域侦测开始
	MINOR_SMART_REGION_ENTRANCE_END    = 0x2b //进入区域侦测结束
	MINOR_SMART_PEOPLE_GATHERING_BEGIN = 0x2c //人员聚集侦测开始
	MINOR_SMART_PEOPLE_GATHERING_END   = 0x2d //人员聚集侦测结束
	MINOR_SMART_FAST_MOVING_BEGIN      = 0x2e //快速运动侦测开始
	MINOR_SMART_FAST_MOVING_END        = 0x2f //快速运动侦测结束

	MINOR_VCA_FACE_ALARM_END            = 0x30 /*人脸侦测结束*/
	MINOR_VCA_SCENE_CHANGE_ALARM_BEGIN  = 0x31 /*场景变更侦测开始*/
	MINOR_VCA_SCENE_CHANGE_ALARM_END    = 0x32 /*场景变更侦测结束*/
	MINOR_VCA_ALARM_AUDIOINPUT_BEGIN    = 0x33 /*音频异常输入开始*/
	MINOR_VCA_ALARM_AUDIOINPUT_END      = 0x34 /*音频异常输入结束*/
	MINOR_VCA_ALARM_AUDIOABNORMAL_BEGIN = 0x35 /*声强突变侦测开始*/
	MINOR_VCA_ALARM_AUDIOABNORMAL_END   = 0x36 /*声强突变侦测结束*/

	MINOR_VCA_LECTURE_DETECTION_BEGIN = 0x37 //授课侦测开始报警
	MINOR_VCA_LECTURE_DETECTION_END   = 0x38 //授课侦测结束报警
	MINOR_VCA_ALARM_AUDIOSTEEPDROP    = 0x39 //声强陡降 2014-03-21
	MINOR_VCA_ANSWER_DETECTION_BEGIN  = 0x3a //回答问题侦测开始报警
	MINOR_VCA_ANSWER_DETECTION_END    = 0x3b //回答问题侦测结束报警

	MINOR_SMART_PARKING_BEGIN            = 0x3c //停车侦测开始
	MINOR_SMART_PARKING_END              = 0x3d //停车侦测结束
	MINOR_SMART_UNATTENDED_BAGGAGE_BEGIN = 0x3e //物品遗留侦测开始
	MINOR_SMART_UNATTENDED_BAGGAGE_END   = 0x3f //物品遗留侦测结束
	MINOR_SMART_OBJECT_REMOVAL_BEGIN     = 0x40 //物品拿取侦测开始
	MINOR_SMART_OBJECT_REMOVAL_END       = 0x41 //物品拿取侦测结束
	MINOR_SMART_VEHICLE_ALARM_START      = 0x46 //车牌检测开始
	MINOR_SMART_VEHICLE_ALARM_STOP       = 0x47 //车牌检测结束
	MINOR_THERMAL_FIREDETECTION          = 0x48 //热成像火点检测侦测开始
	MINOR_THERMAL_FIREDETECTION_END      = 0x49 //热成像火点检测侦测结束
	MINOR_SMART_VANDALPROOF_BEGIN        = 0x50 //防破坏检测开始
	MINOR_SMART_VANDALPROOF_END          = 0x51 //防破坏检测结束

	MINOR_FACESNAP_MATCH_ALARM_START           = 0x55 /*人脸比对报警开始*/
	MINOR_FACESNAP_MATCH_ALARM_STOP            = 0x56 /*人脸比对报警结束*/
	MINOR_WHITELIST_FACESNAP_MATCH_ALARM_START = 0x57 /*白名单人脸比对（陌生人）报警开始*/
	MINOR_WHITELIST_FACESNAP_MATCH_ALARM_STOP  = 0x58 /*白名单人脸比对（陌生人）报警结束*/

	MINOR_THERMAL_SHIPSDETECTION                 = 0x5a //热成像船只检测侦测
	MINOR_THERMAL_THERMOMETRY_EARLYWARNING_BEGIN = 0x5b //热成像测温预警开始
	MINOR_THERMAL_THERMOMETRY_EARLYWARNING_END   = 0x5c //热成像测温预警结束
	MINOR_THERMAL_THERMOMETRY_ALARM_BEGIN        = 0x5d //热成像测温报警开始
	MINOR_THERMAL_THERMOMETRY_ALARM_END          = 0x5e //热成像测温报警结束
	MINOR_THERMAL_THERMOMETRY_DIFF_ALARM_BEGIN   = 0x5f //热成像温差报警开始
	MINOR_THERMAL_THERMOMETRY_DIFF_ALARM_END     = 0x60 //热成像温差报警结束
	MINOR_FACE_THERMOMETRY_ALARM                 = 0x63 //人脸测温报警
	MINOR_SMART_DENSEFOGDETECTION_BEGIN          = 0x6e //大雾侦测开始
	MINOR_SMART_DENSEFOGDETECTION_END            = 0x6f //大雾侦测结束
	MINOR_RUNNING_ALARM                          = 0x70 //奔跑检测
	MINOR_RETENTION_ALARM                        = 0x71 //滞留检测
	MINOR_SAFETY_HELMET_ALARM_START              = 0x72 /*未佩戴安全帽检测报警开始*/
	MINOR_SAFETY_HELMET_ALARM_STOP               = 0x73 /*未佩戴安全帽检测报警结束*/
	MINOR_HFPD_ALARM_START                       = 0x74 /*高频人员检测报警开始*/
	MINOR_HFPD_ALARM_STOP                        = 0x75 /*高频人员检测报警结束*/
	MINOR_MIXED_TARGET_ALARM_START               = 0x76 /*混合目标检测报警开始*/
	MINOR_MIXED_TARGET_ALARM_STOP                = 0x77 /*混合目标检测报警结束*/
	MINOR_VCA_PLAY_CELLPHONE_ALARM_BEGIN         = 0x78 //玩手机检测报警开始
	MINOR_VCA_PLAY_CELLPHONE_ALARM_END           = 0x79 //玩手机检测报警结束
	MINOR_VCA_GET_UP_ALARM_BEGIN                 = 0x80 //起床检测报警开始
	MINOR_VCA_GET_UP_ALARM_END                   = 0x81 //起床检测报警结束
	MINOR_VCA_ADV_REACH_HEIGHT_ALARM_BEGIN       = 0x82 //折线攀高报警开始
	MINOR_VCA_ADV_REACH_HEIGHT_ALARM_END         = 0x83 //折线攀高报警结束
	MINOR_VCA_TOILET_TARRY_ALARM_BEGIN           = 0x84 //如厕超时报警开始
	MINOR_VCA_TOILET_TARRY_ALARM_END             = 0x85 //如厕超时报警结束
	MINOR_HUMAN_RECOGNITION_ALARM_BEGIN          = 0x86 //人体识别报警开始
	MINOR_HUMAN_RECOGNITION_ALARM_END            = 0x87 //人体识别报警结束
	MINOR_STUDENTS_STOODUP_ALARM_BEGIN           = 0x88 //学生起立报警开始
	MINOR_STUDENTS_STOODUP_ALARM_END             = 0x89 //学生起立报警结束
	MINOR_FRAMES_PEOPLE_COUNTING_ALARM           = 0x8a //区域人数统计报警
	MINOR_FACE_SNAP_ALARM_BEGIN                  = 0x8b //人脸抓拍报警开始
	MINOR_FACE_SNAP_ALARM_END                    = 0x8c //人脸抓拍报警结束
	MINOR_TEACHER_BEHAVIOR_DETECT_ALARM_BEGIN    = 0x8d //教师行为检测报警开始
	MINOR_TEACHER_BEHAVIOR_DETECT_ALARM_END      = 0x8e //教师行为检测报警结束
	MINOR_PERIMETER_CAPTURE_ALARM_BEGIN          = 0x8f //周界抓拍报警开始
	MINOR_PERIMETER_CAPTURE_ALARM_END            = 0x90 //周界抓拍报警结束
	MINOR_UNREGISTERED_STREET_VENDOR_ALARM       = 0x91 //非法摆摊报警

	MINOR_PERSON_QUEUE_TIME_ALARM_BEGIN       = 0x92 //排队时长检测报警开始
	MINOR_PERSON_QUEUE_TIME_ALARM_END         = 0x93 //排队时长检测报警结束
	MINOR_PERSON_QUEUE_COUNTING_ALARM_BEGIN   = 0x94 //排队人数检测报警开始
	MINOR_PERSON_QUEUE_COUNTING_ALARM_END     = 0x95 //排队人数检测报警结束
	MINOR_FACE_SNAP_MATCH_FAILURE_ALARM_START = 0x96 //人脸比对失败报警开始
	MINOR_FACE_SNAP_MATCH_FAILURE_ALARM_END   = 0x97 //人脸比对失败报警结束

	MINOR_ACCESS_CONTROLLER_EVENT      = 0x100 //门禁主机事件
	MINOR_VIDEO_INTERCOM_EVENT         = 0x101 //可视对讲设备事件
	MINOR_GJD_EVENT                    = 0x102 //GJD报警主机事件
	MINOR_LUMINITE_EVENT               = 0x103 // LUMINITE报警主机事件
	MINOR_OPTEX_EVENT                  = 0x104 // OPTEX报警主机事件
	MINOR_CAMERA_DETECTOR_EVENT        = 0x105 // 传感器事件
	MINOR_SECURITY_CONTROL_PANEL_EVENT = 0x106 //海康报警主机事件

	MINOR_VCA_SPACE_CHANGE_START       = 0x10c //间距异常检测开始
	MINOR_VCA_SPACE_CHANGE_STOP        = 0x10d //间距异常检测结束
	MINOR_MANUAL_ALARM                 = 0x10e //手动报警
	MINOR_DETECTOR_ALARM               = 0x10f //探测器报警
	MINOR_LINKAGE_ALARM                = 0x110 //联动报警
	MINOR_VCA_SITUATION_ANALYSIS_START = 0x111 //态势分析检测开始
	MINOR_VCA_SITUATION_ANALYSIS_STOP  = 0x112 //态势分析检测结束
	MINOR_FIRE_ALARM                   = 0x113 //火警报警
	MINOR_SUPERVISE_ALARM              = 0x114 //监管报警
	MINOR_SHIELD_ALARM                 = 0x115 //屏蔽报警
	MINOR_ABNORMAL_ALARM               = 0x116 //故障报警
	MINOR_RESIDUAL_CURRENT_ALARM       = 0x117 //剩余电流报警
	MINOR_TEMPERATURE_ALARM            = 0x118 //温度报警
	MINOR_ARC_ALARM                    = 0x119 //电弧报警

	MINOR_VCA_YARD_TARRY_ALARM_BEGIN        = 0x11a //放风场滞留报警开始
	MINOR_VCA_YARD_TARRY_ALARM_END          = 0x11b //放风场滞留报警结束
	MINOR_VCA_KEY_PERSON_GET_UP_ALARM_BEGIN = 0x11c //重点人员起身报警开始
	MINOR_VCA_KEY_PERSON_GET_UP_ALARM_END   = 0x11d //重点人员起身报警结束
	MINOR_VCA_SIT_QUIETLY_ALARM_BEGIN       = 0x11e //静坐报警开始
	MINOR_VCA_SIT_QUIETLY_ALARM_END         = 0x11f //静坐报警结束
	MINOR_VCA_STAND_UP_ALARM_BEGIN          = 0x120 //站立报警开始
	MINOR_VCA_STAND_UP_ALARM_END            = 0x121 //站立报警结束
	MINOR_VCA_REACH_HIGHT_ALARM_BEGIN       = 0x122 //攀高报警开始
	MINOR_VCA_REACH_HIGHT_ALARM_END         = 0x123 //攀高报警结束

	MINOR_LFPD_ALARM_START = 0x124 /*低频人员检测报警开始*/
	MINOR_LFPD_ALARM_STOP  = 0x125 /*低频人员检测报警结束*/

	MINOR_DREDGERDETECTION_ALARM          = 0x126 // 挖沙船检测报警
	MINOR_STUDENT_BEHAVIOR_ALARM_BEGIN    = 0x127 //课堂学生行为报警开始
	MINOR_STUDENT_BEHAVIOR_ALARM_END      = 0x128 //课堂学生行为报警结束
	MINOR_VCA_ALARM_VEHICLEMONITOR        = 0x129 //车辆布控报警(用于车辆布控第一次上来车辆检测（带车辆布控信息）的日志)
	MINOR_WASTEGASDETECTION_ALARM         = 0x130 // 废气排放监测事件上报
	MINOR_GREYSCALE_ALARM                 = 0x131 // 灰度报警
	MINOR_VIBRATION_DETECTION_ALARM_BEGIN = 0x132 //振动侦测报警开始
	MINOR_VIBRATION_DETECTION_ALARM_END   = 0x133 //振动侦测报警结束

	//0x400-0x1000 门禁报警
	MINOR_ALARMIN_SHORT_CIRCUIT            = 0x400 //防区短路报警
	MINOR_ALARMIN_BROKEN_CIRCUIT           = 0x401 //防区断路报警
	MINOR_ALARMIN_EXCEPTION                = 0x402 //防区异常报警
	MINOR_ALARMIN_RESUME                   = 0x403 //防区报警恢复
	MINOR_HOST_DESMANTLE_ALARM             = 0x404 //设备防拆报警
	MINOR_HOST_DESMANTLE_RESUME            = 0x405 //设备防拆恢复
	MINOR_CARD_READER_DESMANTLE_ALARM      = 0x406 //读卡器防拆报警
	MINOR_CARD_READER_DESMANTLE_RESUME     = 0x407 //读卡器防拆恢复
	MINOR_CASE_SENSOR_ALARM                = 0x408 //事件输入报警
	MINOR_CASE_SENSOR_RESUME               = 0x409 //事件输入恢复
	MINOR_STRESS_ALARM                     = 0x40a //胁迫报警
	MINOR_OFFLINE_ECENT_NEARLY_FULL        = 0x40b //离线事件满90%报警
	MINOR_CARD_MAX_AUTHENTICATE_FAIL       = 0x40c //卡号认证失败超次报警
	MINOR_SD_CARD_FULL                     = 0x40d //SD卡存储满报警
	MINOR_LINKAGE_CAPTURE_PIC              = 0x40e //联动抓拍事件报警
	MINOR_SECURITY_MODULE_DESMANTLE_ALARM  = 0x40f //门控安全模块防拆报警
	MINOR_SECURITY_MODULE_DESMANTLE_RESUME = 0x410 //门控安全模块防拆恢复

	MINOR_POS_START_ALARM            = 0x411 //POS开启
	MINOR_POS_END_ALARM              = 0x412 //POS结束
	MINOR_FACE_IMAGE_QUALITY_LOW     = 0x413 //人脸图像画质低
	MINOR_FINGE_RPRINT_QUALITY_LOW   = 0x414 //指纹图像画质低
	MINOR_FIRE_IMPORT_SHORT_CIRCUIT  = 0x415 //消防输入短路报警
	MINOR_FIRE_IMPORT_BROKEN_CIRCUIT = 0x416 //消防输入断路报警
	MINOR_FIRE_IMPORT_RESUME         = 0x417 //消防输入恢复
	MINOR_FIRE_BUTTON_TRIGGER        = 0x418 //消防按钮触发
	MINOR_FIRE_BUTTON_RESUME         = 0x419 //消防按钮恢复
	MINOR_MAINTENANCE_BUTTON_TRIGGER = 0x41a //维护按钮触发
	MINOR_MAINTENANCE_BUTTON_RESUME  = 0x41b //维护按钮恢复
	MINOR_EMERGENCY_BUTTON_TRIGGER   = 0x41c //紧急按钮触发
	MINOR_EMERGENCY_BUTTON_RESUME    = 0x41d //紧急按钮恢复
	MINOR_DISTRACT_CONTROLLER_ALARM  = 0x41e //分控器防拆报警
	MINOR_DISTRACT_CONTROLLER_RESUME = 0x41f //分控器防拆报警恢复

	MINOR_PERSON_DENSITY_DETECTION_START = 0x420 //人员密度超阈值报警开始
	MINOR_PERSON_DENSITY_DETECTION_END   = 0x421 //人员密度超阈值报警结束

	MINOR_CHANNEL_CONTROLLER_DESMANTLE_ALARM    = 0x422 //通道控制器防拆报警
	MINOR_CHANNEL_CONTROLLER_DESMANTLE_RESUME   = 0x423 //通道控制器防拆报警恢复
	MINOR_CHANNEL_CONTROLLER_FIRE_IMPORT_ALARM  = 0x424 //通道控制器消防输入报警
	MINOR_CHANNEL_CONTROLLER_FIRE_IMPORT_RESUME = 0x425 //通道控制器消防输入报警恢复

	MINOR_HEART_RATE_ABNORMAL_BEGIN               = 0x426 //心率异常报警开始
	MINOR_HEART_RATE_ABNORMAL_END                 = 0x427 //心率异常报警结束
	MINOR_BLOOD_OXYGEN_ABNORMAL_BEGIN             = 0x428 //血氧异常报警开始
	MINOR_BLOOD_OXYGEN_ABNORMAL_END               = 0x429 //血氧异常报警结束
	MINOR_SYSTOLIC_BLOOD_PRESSURE_ABNORMAL_BEGIN  = 0x42a //血压收缩压异常报警开始
	MINOR_SYSTOLIC_BLOOD_PRESSURE_ABNORMAL_END    = 0x42b //血压收缩压异常报警结束
	MINOR_DIASTOLIC_BLOOD_PRESSURE_ABNORMAL_BEGIN = 0x42c //血压舒张压异常报警开始
	MINOR_DIASTOLIC_BLOOD_PRESSURE_ABNORMAL_END   = 0x42d //血压舒张压异常报警结束
	MINOR_VCA_LEAVE_POSITION_START                = 0x42e //离岗检测开始
	MINOR_VCA_LEAVE_POSITION_STOP                 = 0x42f //离岗检测结束
	MINOR_VCA_STOOODUP_START                      = 0x430 //起立检测开始
	MINOR_VCA_STOOODUP_STOP                       = 0x431 //起立检测结束
	MINOR_VCA_PEOPLENUM_CHANGE_START              = 0x434 //人数变化开始
	MINOR_VCA_PEOPLENUM_CHANGE_STOP               = 0x435 //人数变化结束
	MINOR_VCA_RUNNING_START                       = 0x438 //人员奔跑开始
	MINOR_VCA_RUNNING_STOP                        = 0x439 //人员奔跑结束
	MINOR_VCA_VIOLENT_MOTION_START                = 0x43a //剧烈运动开始
	MINOR_VCA_VIOLENT_MOTION_STOP                 = 0x43b //剧烈运动结束
	MINOR_VCA_FAIL_DOWN_START                     = 0x43c //人员倒地开始
	MINOR_VCA_FAIL_DOWN_STOP                      = 0x43d //人员倒地结束
	MINOR_VCA_RETENTION_START                     = 0x43e //人员滞留开始
	MINOR_VCA_RETENTION_STOP                      = 0x43f //人员滞留结束

	MINOR_PRINTER_OUT_OF_PAPER    = 0x440 //打印机缺纸报警
	MINOR_LEGAL_EVENT_NEARLY_FULL = 0x442 //离线合法事件满90%报警

	MINOR_ALARM_CUSTOM1               = 0x900 //门禁自定义报警1
	MINOR_ALARM_CUSTOM2               = 0x901 //门禁自定义报警2
	MINOR_ALARM_CUSTOM3               = 0x902 //门禁自定义报警3
	MINOR_ALARM_CUSTOM4               = 0x903 //门禁自定义报警4
	MINOR_ALARM_CUSTOM5               = 0x904 //门禁自定义报警5
	MINOR_ALARM_CUSTOM6               = 0x905 //门禁自定义报警6
	MINOR_ALARM_CUSTOM7               = 0x906 //门禁自定义报警7
	MINOR_ALARM_CUSTOM8               = 0x907 //门禁自定义报警8
	MINOR_ALARM_CUSTOM9               = 0x908 //门禁自定义报警9
	MINOR_ALARM_CUSTOM10              = 0x909 //门禁自定义报警10
	MINOR_ALARM_CUSTOM11              = 0x90a //门禁自定义报警11
	MINOR_ALARM_CUSTOM12              = 0x90b //门禁自定义报警12
	MINOR_ALARM_CUSTOM13              = 0x90c //门禁自定义报警13
	MINOR_ALARM_CUSTOM14              = 0x90d //门禁自定义报警14
	MINOR_ALARM_CUSTOM15              = 0x90e //门禁自定义报警15
	MINOR_ALARM_CUSTOM16              = 0x90f //门禁自定义报警16
	MINOR_ALARM_CUSTOM17              = 0x910 //门禁自定义报警17
	MINOR_ALARM_CUSTOM18              = 0x911 //门禁自定义报警18
	MINOR_ALARM_CUSTOM19              = 0x912 //门禁自定义报警19
	MINOR_ALARM_CUSTOM20              = 0x913 //门禁自定义报警20
	MINOR_ALARM_CUSTOM21              = 0x914 //门禁自定义报警21
	MINOR_ALARM_CUSTOM22              = 0x915 //门禁自定义报警22
	MINOR_ALARM_CUSTOM23              = 0x916 //门禁自定义报警23
	MINOR_ALARM_CUSTOM24              = 0x917 //门禁自定义报警24
	MINOR_ALARM_CUSTOM25              = 0x918 //门禁自定义报警25
	MINOR_ALARM_CUSTOM26              = 0x919 //门禁自定义报警26
	MINOR_ALARM_CUSTOM27              = 0x91a //门禁自定义报警27
	MINOR_ALARM_CUSTOM28              = 0x91b //门禁自定义报警28
	MINOR_ALARM_CUSTOM29              = 0x91c //门禁自定义报警29
	MINOR_ALARM_CUSTOM30              = 0x91d //门禁自定义报警30
	MINOR_ALARM_CUSTOM31              = 0x91e //门禁自定义报警31
	MINOR_ALARM_CUSTOM32              = 0x91f //门禁自定义报警32
	MINOR_ALARM_CUSTOM33              = 0x920 //门禁自定义报警33
	MINOR_ALARM_CUSTOM34              = 0x921 //门禁自定义报警34
	MINOR_ALARM_CUSTOM35              = 0x922 //门禁自定义报警35
	MINOR_ALARM_CUSTOM36              = 0x923 //门禁自定义报警36
	MINOR_ALARM_CUSTOM37              = 0x924 //门禁自定义报警37
	MINOR_ALARM_CUSTOM38              = 0x925 //门禁自定义报警38
	MINOR_ALARM_CUSTOM39              = 0x926 //门禁自定义报警39
	MINOR_ALARM_CUSTOM40              = 0x927 //门禁自定义报警40
	MINOR_ALARM_CUSTOM41              = 0x928 //门禁自定义报警41
	MINOR_ALARM_CUSTOM42              = 0x929 //门禁自定义报警42
	MINOR_ALARM_CUSTOM43              = 0x92a //门禁自定义报警43
	MINOR_ALARM_CUSTOM44              = 0x92b //门禁自定义报警44
	MINOR_ALARM_CUSTOM45              = 0x92c //门禁自定义报警45
	MINOR_ALARM_CUSTOM46              = 0x92d //门禁自定义报警46
	MINOR_ALARM_CUSTOM47              = 0x92e //门禁自定义报警47
	MINOR_ALARM_CUSTOM48              = 0x92f //门禁自定义报警48
	MINOR_ALARM_CUSTOM49              = 0x930 //门禁自定义报警49
	MINOR_ALARM_CUSTOM50              = 0x931 //门禁自定义报警50
	MINOR_ALARM_CUSTOM51              = 0x932 //门禁自定义报警51
	MINOR_ALARM_CUSTOM52              = 0x933 //门禁自定义报警52
	MINOR_ALARM_CUSTOM53              = 0x934 //门禁自定义报警53
	MINOR_ALARM_CUSTOM54              = 0x935 //门禁自定义报警54
	MINOR_ALARM_CUSTOM55              = 0x936 //门禁自定义报警55
	MINOR_ALARM_CUSTOM56              = 0x937 //门禁自定义报警56
	MINOR_ALARM_CUSTOM57              = 0x938 //门禁自定义报警57
	MINOR_ALARM_CUSTOM58              = 0x939 //门禁自定义报警58
	MINOR_ALARM_CUSTOM59              = 0x93a //门禁自定义报警59
	MINOR_ALARM_CUSTOM60              = 0x93b //门禁自定义报警60
	MINOR_ALARM_CUSTOM61              = 0x93c //门禁自定义报警61
	MINOR_ALARM_CUSTOM62              = 0x93d //门禁自定义报警62
	MINOR_ALARM_CUSTOM63              = 0x93e //门禁自定义报警63
	MINOR_ALARM_CUSTOM64              = 0x93f //门禁自定义报警64
	MINOR_LOCK_HIJIACK_FINGER_ALARM   = 0x950 //智能锁防劫持指纹报警
	MINOR_LOCK_HIJIACK_PASSWORD_ALARM = 0x951 //智能锁防劫持密码报警
	MINOR_LOCK_PRY_DOOR_ALARM         = 0x952 //智能锁撬门报警
	MINOR_LOCK_LOCKED_ALARM           = 0x953 //智能锁锁定报警
	MINOR_LOCK_BATTERLOW_ALARM        = 0x954 //智能锁低电压报警
	MINOR_LOCK_BLACKLIST_DOOR_ALARM   = 0x955 //智能锁黑名单报警
	MINOR_LOCK_OFFLINE_ALARM          = 0x956 //智能锁离线报警
	MINOR_LOCK_UNCLOSED_ALARM         = 0x957 //智能锁虚掩报警
	MINOR_LOCK_NO_HOME_ALARM          = 0x958 //智能锁用户未回家报警
	MINOR_LOCK_MAGNETOMETER_ALARM     = 0x959 //门磁探测器报警
	MINOR_LOCK_IR_DETECTOR_ALARM      = 0x95a //红外探测器报警
	MINOR_LOCK_FP_LOCKED_ALARM        = 0x95b //指纹锁定报警
	MINOR_LOCK_PASSWORD_LOCKED_ALARM  = 0x95c //密码锁定报警
	MINOR_LOCK_HIJIACK_ALARM          = 0x95d //智能锁防劫持报警

	//2018-04-23 通用物联网关报警日志类型
	MINOR_ALARMHOST_SHORT_CIRCUIT         = 0x1001 //短路报警
	MINOR_ALARMHOST_BROKEN_CIRCUIT        = 0x1002 //断路报警
	MINOR_ALARMHOST_ALARM_RESET           = 0x1003 //报警复位
	MINOR_ALARMHOST_ALARM_NORMAL          = 0x1004 //报警恢复正常
	MINOR_ALARMHOST_PASSWORD_ERROR        = 0x1005 //密码错误（连续3次输入密码错误）
	MINOR_ALARMHOST_ID_CARD_ILLEGALLY     = 0x1006 //非法感应卡ID
	MINOR_ALARMHOST_KEYPAD_REMOVE         = 0x1007 //键盘防拆
	MINOR_ALARMHOST_KEYPAD_REMOVE_RESTORE = 0x1008 //键盘防拆复位

	MINOR_ALARMHOST_BELOW_ALARM_LIMIT1 = 0x1011 //模拟量低于报警限1
	MINOR_ALARMHOST_BELOW_ALARM_LIMIT2 = 0x1012 //模拟量低于报警限2
	MINOR_ALARMHOST_BELOW_ALARM_LIMIT3 = 0x1013 //模拟量低于报警限3
	MINOR_ALARMHOST_BELOW_ALARM_LIMIT4 = 0x1014 //模拟量低于报警限4
	MINOR_ALARMHOST_ABOVE_ALARM_LIMIT1 = 0x1015 //模拟量高于报警限1
	MINOR_ALARMHOST_ABOVE_ALARM_LIMIT2 = 0x1016 //模拟量高于报警限2
	MINOR_ALARMHOST_ABOVE_ALARM_LIMIT3 = 0x1017 //模拟量高于报警限3
	MINOR_ALARMHOST_ABOVE_ALARM_LIMIT4 = 0x1018 //模拟量高于报警限4

	MINOR_ALARMHOST_VIRTUAL_DEFENCE_BANDIT = 0x1021 //软防区匪警
	MINOR_ALARMHOST_VIRTUAL_DEFENCE_FIRE   = 0x1022 //软防区火警
	MINOR_ALARMHOST_VIRTUAL_DEFENCE_URGENT = 0x1023 //软防区紧急

	MINOR_UPS_ALARM                            = 0x1028 //UPS报警
	MINOR_ELECTRICITY_METER_ALARM              = 0x1029 //智能电表报警
	MINOR_SWITCH_POWER_ALARM                   = 0x1030 //开关电源报警
	MINOR_GAS_DETECT_SYS_ALARM                 = 0x1031 //气体检测系统报警
	MINOR_TRANSFORMER_TEMPRATURE_ALARM         = 0x1032 //变电器温显表报警
	MINOR_TEMP_HUMI_ALARM                      = 0x1033 //温湿度传感器报警
	MINOR_UPS_ALARM_RESTORE                    = 0x1034 //UPS报警恢复
	MINOR_ELECTRICITY_METER_ALARM_RESTORE      = 0x1035 //智能电表报警恢复
	MINOR_SWITCH_POWER_ALARM_RESTORE           = 0x1036 //开关电源报警恢复
	MINOR_GAS_DETECT_SYS_ALARM_RESTORE         = 0x1037 //气体检测系统报警恢复
	MINOR_TRANSFORMER_TEMPRATURE_ALARM_RESTORE = 0x1038 //变电器温显表报警恢复
	MINOR_TEMP_HUMI_ALARM_RESTORE              = 0x1039 //温湿度传感器报警恢复
	MINOR_WATER_LEVEL_SENSOR_ALARM             = 0x1040 //水位传感器报警
	MINOR_WATER_LEVEL_SENSOR_ALARM_RESTORE     = 0x1041 //水位传感器报警恢复
	MINOR_DUST_NOISE_ALARM                     = 0x1042 //扬尘噪声传感器报警
	MINOR_DUST_NOISE_ALARM_RESTORE             = 0x1043 //扬尘噪声传感器报警恢复
	MINOR_ENVIRONMENTAL_LOGGER_ALARM           = 0x1044 //环境采集仪报警
	MINOR_ENVIRONMENTAL_LOGGER_ALARM_RESTORE   = 0x1045 //环境采集仪报警恢复

	MINOR_TRIGGER_TAMPER                    = 0x1046 //探测器防拆
	MINOR_TRIGGER_TAMPER_RESTORE            = 0x1047 //探测器防拆恢复
	MINOR_EMERGENCY_CALL_HELP_ALARM         = 0x1048 //紧急呼叫求助报警
	MINOR_EMERGENCY_CALL_HELP_ALARM_RESTORE = 0x1049 //紧急呼叫求助报警恢复
	MINOR_CONSULTING_ALARM                  = 0x1050 //业务咨询报警
	MINOR_CONSULTING_ALARM_RESTORE          = 0x1051 //业务咨询报警恢复
	MINOR_ALARMHOST_ZONE_MODULE_REMOVE      = 0x1052 //防区模块防拆
	MINOR_ALARMHOST_ZONE_MODULE_RESET       = 0x1053 //防区模块防拆复位

	MINOR_ALARMHOST_ALARM_WIND_SPEED_ALARM          = 0x1054 //风速传感器告警
	MINOR_ALARMHOST_ALARM_WIND_SPEED_ALARM_RESTORE  = 0x1055 //风速传感器告警恢复
	MINOR_ALARMHOST_ALARM_GENERATE_OUTPUT_ALARM     = 0x1056 //通用扩展输出模块告警
	MINOR_ALARMHOST_ALARM_GENERATE_OUTPUT_RESTORE   = 0x1057 //通用扩展输出模块告警恢复
	MINOR_ALARMHOST_ALARM_SOAK_ALARM                = 0x1058 //水浸传感器告警
	MINOR_ALARMHOST_ALARM_SOAK_ALARM_RESTORE        = 0x1059 //水浸传感器告警恢复
	MINOR_ALARMHOST_ALARM_SOLAR_POWER_ALARM         = 0x1060 //太阳能传感器告警
	MINOR_ALARMHOST_ALARM_SOLAR_POWER_ALARM_RESTORE = 0x1061 //太阳能传感器告警恢复
	MINOR_ALARMHOST_ALARM_SF6_ALARM                 = 0x1062 //SF6报警主机告警
	MINOR_ALARMHOST_ALARM_SF6_ALARM_RESTORE         = 0x1063 //SF6报警主机告警恢复
	MINOR_ALARMHOST_ALARM_WEIGHT_ALARM              = 0x1064 //称重仪告警
	MINOR_ALARMHOST_ALARM_WEIGHT_ALARM_RESTORE      = 0x1065 //称重仪告警恢复
	MINOR_ALARMHOST_ALARM_WEATHER_ALARM             = 0x1066 //气象采集系统告警
	MINOR_ALARMHOST_ALARM_WEATHER_ALARM_RESTORE     = 0x1067 //气象采集系统告警恢复
	MINOR_ALARMHOST_ALARM_FUEL_GAS_ALARM            = 0x1068 //燃气监测系统告警
	MINOR_ALARMHOST_ALARM_FUEL_GAS_ALARM_RESTORE    = 0x1069 //燃气监测系统告警恢
	MINOR_ALARMHOST_ALARM_FIRE_ALARM                = 0x1070 //火灾报警系统告警
	MINOR_ALARMHOST_ALARM_FIRE_ALARM_RESTORE        = 0x1071 //火灾报警系统告警恢复
	MINOR_ALARMHOST_WIRELESS_OUTPUT_MODULE_REMOVE   = 0x1072 //无线输出模块防拆
	MINOR_ALARMHOST_WIRELESS_OUTPUT_MODULE_RESET    = 0x1073 //无线输出模块防拆复位
	MINOR_ALARMHOST_WIRELESS_REPEATER_MODULE_REMOVE = 0x1074 //无线中继器防拆

	MINOR_ALARMHOST_WIRELESS_SIREN_MODULE_REMOVE = 0x1075 //无线警号防拆
	MINOR_ALARMHOST_WIRELESS_SIREN_MODULE_RESET  = 0x1076 //无线警号防拆复位

	MINOR_RS485_DEV_ALARM              = 0x1077 //RS485外接设备报警（针对设备类型未知的设备）
	MINOR_RS485_DEV_RESTORE            = 0x1078 //RS485外接设备报警恢复（针对设备类型未知的设备）
	MINOR_ALARMHOST_ALARM_HOST_ALARM   = 0x1079 //消防主机报警
	MINOR_ALARMHOST_ALARM_HOST_RESTORE = 0x107a //消防主机报警恢复

	MINOR_AIR_CONDITION_DEV_ALARM   = 0x107b //空调控制器报警
	MINOR_AIR_CONDITION_DEV_RESTORE = 0x107c //空调控制器报警恢复

	MINOR_ALARMHOST_WIRELESS_REPEATER_MODULE_RESET = 0x107d //无线中继器防拆复位

	MINOR_ALARM_ELEVATOR_BREAKDOWN     = 0x107e //电梯故障
	MINOR_WATER_PRESSURE_SENSOR_ALARM  = 0x107f //水压传感器报警
	MINOR_FLOW_SENSOR_ALARM            = 0x1080 //流量传感器报警
	MINOR_SENSOR_LINKAGE_ALARM         = 0x1081 //传感器联动报警
	MINOR_SENSOR_LINKAGE_ALARM_RESTORE = 0x1082 //传感器联动报警恢复

	//LED报警次类型 0x1201 ~ 0x1300
	MINOR_SYSTEM_CHECK_ALARM = 0x1201 //系统检测报警

	/* 异常 */
	//主类型
	MAJOR_EXCEPTION = 0x2
	//次类型
	MINOR_SUBSYSTEM_ERROR = 0x0a /* 子系统异常 */
	MINOR_RAID_ERROR      = 0x20 /* 阵列异常 */
	MINOR_VI_LOST         = 0x21 /* 视频信号丢失 */
	MINOR_ILLEGAL_ACCESS  = 0x22 /* 非法访问 */
	MINOR_HD_FULL         = 0x23 /* 硬盘满 */
	MINOR_HD_ERROR        = 0x24 /* 硬盘错误 */
	MINOR_DCD_LOST        = 0x25 /* MODEM 掉线(保留不使用) */
	MINOR_IP_CONFLICT     = 0x26 /* IP地址冲突 */
	MINOR_NET_BROKEN      = 0x27 /* 网络断开*/
	MINOR_REC_ERROR       = 0x28 /* 录像出错 */
	MINOR_IPC_NO_LINK     = 0x29 /* IPC连接异常 */
	MINOR_VI_EXCEPTION    = 0x2a /* 视频输入异常(只针对模拟通道) */
	MINOR_IPC_IP_CONFLICT = 0x2b /*ipc ip 地址 冲突*/
	MINOR_SENCE_EXCEPTION = 0x2c // 场景异常

	MINOR_PIC_REC_ERROR       = 0x2d /* 抓图出错--获取图片文件失败*/
	MINOR_VI_MISMATCH         = 0x2e /* 视频制式不匹配*/
	MINOR_RESOLUTION_MISMATCH = 0x2f /*前端/录像分辨率不匹配  */

	//2009-12-16 增加视频综合平台日志类型
	MINOR_FANABNORMAL              = 0x31 /* 视频综合平台：风扇状态异常 */
	MINOR_FANRESUME                = 0x32 /* 视频综合平台：风扇状态恢复正常 */
	MINOR_SUBSYSTEM_ABNORMALREBOOT = 0x33 /* 视频综合平台：6467异常重启 */
	MINOR_MATRIX_STARTBUZZER       = 0x34 /* 视频综合平台：dm6467异常，启动蜂鸣器 */

	//2010-01-22 增加视频综合平台异常日志次类型
	MINOR_NET_ABNORMAL          = 0x35 /*网络状态异常*/
	MINOR_MEM_ABNORMAL          = 0x36 /*内存状态异常*/
	MINOR_FILE_ABNORMAL         = 0x37 /*文件状态异常*/
	MINOR_PANEL_ABNORMAL        = 0x38 /*前面板连接异常*/
	MINOR_PANEL_RESUME          = 0x39 /*前面板恢复正常*/
	MINOR_RS485_DEVICE_ABNORMAL = 0x3a /*RS485连接状态异常*/
	MINOR_RS485_DEVICE_REVERT   = 0x3b /*RS485连接状态异常恢复*/

	//2012-2-18 增加大屏控制器异常日志次类型
	MINOR_SCREEN_SUBSYSTEM_ABNORMALREBOOT  = 0x3c //子板异常启动
	MINOR_SCREEN_SUBSYSTEM_ABNORMALINSERT  = 0x3d //子板插入
	MINOR_SCREEN_SUBSYSTEM_ABNORMALPULLOUT = 0x3e //子板拔出
	MINOR_SCREEN_ABNARMALTEMPERATURE       = 0x3f //温度异常
	//2012-07-26 视频综合平台v2.1
	MINOR_HIGH_TEMPERATURE_PROTECT = 0x40 //子板过热保护

	//Netra 2.2.2
	MINOR_RECORD_OVERFLOW = 0x41 /*缓冲区溢出*/
	MINOR_DSP_ABNORMAL    = 0x42 //DSP异常

	//Netra 3.0.0
	MINOR_ANR_RECORD_FAIED         = 0x43 /*ANR录像失败*/
	MINOR_SPARE_WORK_DEVICE_EXCEPT = 0x44 /*热备设备工作机异常*/
	MINOR_START_IPC_MAS_FAILED     = 0x45 /*开启IPC MAS失败*/
	//高性能 256路NVR
	MINOR_IPCM_CRASH                       = 0x46 /*IPCM异常重启*/
	MINOR_POE_POWER_EXCEPTION              = 0x47 /*POE 供电异常*/
	MINOR_UPLOAD_DATA_CS_EXCEPTION         = 0x48 //云存储数据上传失败/
	MINOR_DIAL_EXCEPTION                   = 0x49 /*拨号异常*/
	MINOR_DEV_EXCEPTION_OFFLINE            = 0x50 //设备异常下线
	MINOR_UPGRADEFAIL                      = 0x51 //远程升级设备失败
	MINOR_AI_LOST                          = 0x52 /* 音频信号丢失 */
	MINOR_SYNC_IPC_PASSWD                  = 0x53 /* 同步IPC密码异常 */
	MINOR_EZVIZ_OFFLINE                    = 0x54 /* 萤石下线异常*/
	MINOR_VQD_ABNORMAL                     = 0x55 //VQD异常
	MINOR_ACCESSORIES_PLATE                = 0x57 //配件板异常
	MINOR_KMS_EXPAMSION_DISK_LOST          = 0x58 // KMS扩容盘丢失
	MINOR_ABNORMAL_PORT                    = 0x59 // 端口异常
	MINOR_CAMERA_ANGLE_ANOMALY             = 0x60 //  相机视角异常
	MINOR_DATA_DISK_ERROE                  = 0x61 //  数据盘错误
	MINOR_INTELLIGENT_SYSTEM_RUNNING_ERROR = 0x62 //  智能系统运行异常
	MINOR_FACESNAP_RESOLUTION_OVERFLOW     = 0x63 //  人脸抓拍码流分辨率超限
	MINOR_SMD_RESOLUTION_OVERFLOW          = 0x64 //  SMD码流分辨率超限
	MINOR_AUDIO_LOSS_EXCEPTION             = 0x65 //  音频丢失异常
	MINOR_SAFETY_HELMET_EXCEPTION          = 0x66 //未佩戴安全帽检测异常
	MINOR_VCA_PIC_LENGTH_OVERFLOW          = 0x67 // VCA图片长度过长（例如超过2M大小的图片）
	MINOR_FACE_MODEL_EXCEPTION             = 0x68 //  人脸库模型同步异常
	MINOR_SSD_EXCEPTION                    = 0x69 // SSD异常
	//NVR集群
	MINOR_CLUSTER_DEVICE_OFFLINE            = 0x70 // 集群内设备下线
	MINOR_CLUSTER_CONFIG_FAILED             = 0x71 // 集群内设备配置失败
	MINOR_CLUSTER_DISASTER_TOLERANCE_EXCEPT = 0x72 // 集群容灾异常:集群CM选举失败,集群存储周期不足,集群带宽不足,集群通道资源不足,集群设备不足等
	MINOR_CLUSTER_STORFULL_EXCEPTION        = 0x73 //集群硬盘满
	MINOR_CLUSTER_VERSION_EXCEPTION         = 0x74 //集群版本异常
	MINOR_CLUSTER_OFFLINENODE_EXCEPTION     = 0x75 //集群掉线数超限
	MINOR_CLUSTER_RECORDCYCLE_EXCEPTION     = 0x76 //集群录像周期不足
	MINOR_CLUSTER_IPCTRANSFER_EXCEPTION     = 0x77 //集群IPC迁移失败
	MINOR_CLUSTER_IPCONFLICT_EXCEPTION      = 0x78 // 集群IP冲突，记录CM的IP地址

	MINOR_GET_SUB_STREAM_FAILURE   = 0x79 //子码流取流失败
	MINOR_HDD_SHM_DETECT_EXCEPTION = 0x7a //硬盘SHM检测异常
	MINOR_DEVICE_FORTIFY_FAILURE   = 0x7b //前端设备报警布防失败
	MINOR_EVENT_UPLOAD_EXCEPTION   = 0x7c //事件发送异常（设备上传事件失败或者丢弃了）

	MINOR_LORA_EXCEPTION    = 0x7d //LoRa异常
	MINOR_AK_OR_SK_IS_EMPTY = 0x7e //云存储密码或加密密码为空

	MINOR_HIGH_HD_TEMPERATURE    = 0x80 /*硬盘温度过高*/
	MINOR_LOW_HD_TEMPERATURE     = 0x81 /*硬盘温度过低*/
	MINOR_HD_IMPACT              = 0x82 /*硬盘受到冲击*/
	MINOR_HD_BAD_BLOCK           = 0x83 /*硬盘出现坏块*/
	MINOR_SEVERE_HD_FAILURE      = 0x84 /*硬盘严重故障*/
	MINOR_RELEASE_FAILED         = 0x85 //信息发布失败
	MINOR_PORT_CONFLICT          = 0x86 //端口冲突
	MINOR_MODULE_STARTUP_FAILED  = 0x87 //模块启动失败
	MINIOR_VCA_RUNNING_EXCEPTION = 0x88 //智能板运行异常

	//0x400-0x1000 门禁异常类型
	MINOR_DEV_POWER_ON                         = 0x400 //设备上电启动
	MINOR_DEV_POWER_OFF                        = 0x401 //设备掉电关闭
	MINOR_WATCH_DOG_RESET                      = 0x402 //看门狗复位
	MINOR_LOW_BATTERY                          = 0x403 //蓄电池电压低
	MINOR_BATTERY_RESUME                       = 0x404 //蓄电池电压恢复正常
	MINOR_AC_OFF                               = 0x405 //交流电断电
	MINOR_AC_RESUME                            = 0x406 //交流电恢复
	MINOR_NET_RESUME                           = 0x407 //网络恢复
	MINOR_FLASH_ABNORMAL                       = 0x408 //FLASH读写异常
	MINOR_CARD_READER_OFFLINE                  = 0x409 //读卡器掉线
	MINOR_CARD_READER_RESUME                   = 0x40a //读卡器掉线恢复
	MINOR_INDICATOR_LIGHT_OFF                  = 0x40b //指示灯关闭
	MINOR_INDICATOR_LIGHT_RESUME               = 0x40c //指示灯恢复
	MINOR_CHANNEL_CONTROLLER_OFF               = 0x40d //通道控制器掉线
	MINOR_CHANNEL_CONTROLLER_RESUME            = 0x40e //通道控制器恢复
	MINOR_SECURITY_MODULE_OFF                  = 0x40f //门控安全模块掉线
	MINOR_SECURITY_MODULE_RESUME               = 0x410 //门控安全模块在线
	MINOR_BATTERY_ELECTRIC_LOW                 = 0x411 //电池电压低(仅人脸设备使用)
	MINOR_BATTERY_ELECTRIC_RESUME              = 0x412 //电池电压恢复正常(仅人脸设备使用)
	MINOR_LOCAL_CONTROL_NET_BROKEN             = 0x413 //就地控制器网络断开
	MINOR_LOCAL_CONTROL_NET_RSUME              = 0x414 //就地控制器网络恢复
	MINOR_MASTER_RS485_LOOPNODE_BROKEN         = 0x415 //主控RS485环路节点断开
	MINOR_MASTER_RS485_LOOPNODE_RESUME         = 0x416 //主控RS485环路节点恢复
	MINOR_LOCAL_CONTROL_OFFLINE                = 0x417 //就地控制器掉线
	MINOR_LOCAL_CONTROL_RESUME                 = 0x418 //就地控制器掉线恢复
	MINOR_LOCAL_DOWNSIDE_RS485_LOOPNODE_BROKEN = 0x419 //就地下行RS485环路断开
	MINOR_LOCAL_DOWNSIDE_RS485_LOOPNODE_RESUME = 0x41a //就地下行RS485环路恢复
	MINOR_DISTRACT_CONTROLLER_ONLINE           = 0x41b //分控器在线
	MINOR_DISTRACT_CONTROLLER_OFFLINE          = 0x41c //分控器离线
	MINOR_ID_CARD_READER_NOT_CONNECT           = 0x41d //身份证阅读器未连接（智能专用）
	MINOR_ID_CARD_READER_RESUME                = 0x41e //身份证阅读器连接恢复（智能专用）
	MINOR_FINGER_PRINT_MODULE_NOT_CONNECT      = 0x41f //指纹模组未连接（智能专用）
	MINOR_FINGER_PRINT_MODULE_RESUME           = 0x420 //指纹模组连接恢复（智能专用）
	MINOR_CAMERA_NOT_CONNECT                   = 0x421 //摄像头未连接
	MINOR_CAMERA_RESUME                        = 0x422 //摄像头连接恢复
	MINOR_COM_NOT_CONNECT                      = 0x423 //COM口未连接
	MINOR_COM_RESUME                           = 0x424 //COM口连接恢复
	MINOR_DEVICE_NOT_AUTHORIZE                 = 0x425 //设备未授权
	MINOR_PEOPLE_AND_ID_CARD_DEVICE_ONLINE     = 0x426 //人证设备在线
	MINOR_PEOPLE_AND_ID_CARD_DEVICE_OFFLINE    = 0x427 //人证设备离线
	MINOR_LOCAL_LOGIN_LOCK                     = 0x428 //本地登录锁定
	MINOR_LOCAL_LOGIN_UNLOCK                   = 0x429 //本地登录解锁
	MINOR_SUBMARINEBACK_COMM_BREAK             = 0x42a //与反潜回服务器通信断开
	MINOR_SUBMARINEBACK_COMM_RESUME            = 0x42b //与反潜回服务器通信恢复
	MINOR_MOTOR_SENSOR_EXCEPTION               = 0x42c //电机或传感器异常
	MINOR_CAN_BUS_EXCEPTION                    = 0x42d //CAN总线异常
	MINOR_CAN_BUS_RESUME                       = 0x42e //CAN总线恢复
	MINOR_GATE_TEMPERATURE_OVERRUN             = 0x42f //闸机腔体温度超限
	MINOR_IR_EMITTER_EXCEPTION                 = 0x430 //红外对射异常
	MINOR_IR_EMITTER_RESUME                    = 0x431 //红外对射恢复
	MINOR_LAMP_BOARD_COMM_EXCEPTION            = 0x432 //灯板通信异常
	MINOR_LAMP_BOARD_COMM_RESUME               = 0x433 //灯板通信恢复
	MINOR_IR_ADAPTOR_COMM_EXCEPTION            = 0x434 //红外转接板通信异常
	MINOR_IR_ADAPTOR_COMM_RESUME               = 0x435 //红外转接板通信恢复
	MINOR_PRINTER_ONLINE                       = 0x436 //打印机在线
	MINOR_PRINTER_OFFLINE                      = 0x437 //打印机离线
	MINOR_4G_MOUDLE_ONLINE                     = 0x438 //4G模块在线
	MINOR_4G_MOUDLE_OFFLINE                    = 0x439 //4G模块离线
	MINOR_DSP_START_FAILED                     = 0x43a //DSP启动失败
	MINOR_SMART_REGULATION_NOT_ALLOWED         = 0x43b //智能规则不支持
	MINOR_AUXILIARY_BOARD_OFFLINE              = 0x43c //辅助板掉线
	MINOR_AUXILIARY_BOARD_RESUME               = 0x43d //辅助板掉线恢复
	MINOR_IDCARD_SECURITY_MOUDLE_EXCEPTION     = 0x43e //身份证安全模块异常
	MINOR_IDCARD_SECURITY_MOUDLE_RESUME        = 0x43f //身份证安全模块恢复
	MINOR_FP_PERIPHERAL_EXCEPTION              = 0x440 //指纹采集外设异常
	MINOR_FP_PERIPHERAL_RESUME                 = 0x441 //指纹采集外设恢复

	MINOR_EXCEPTION_CUSTOM1       = 0x900 //门禁自定义异常1
	MINOR_EXCEPTION_CUSTOM2       = 0x901 //门禁自定义异常2
	MINOR_EXCEPTION_CUSTOM3       = 0x902 //门禁自定义异常3
	MINOR_EXCEPTION_CUSTOM4       = 0x903 //门禁自定义异常4
	MINOR_EXCEPTION_CUSTOM5       = 0x904 //门禁自定义异常5
	MINOR_EXCEPTION_CUSTOM6       = 0x905 //门禁自定义异常6
	MINOR_EXCEPTION_CUSTOM7       = 0x906 //门禁自定义异常7
	MINOR_EXCEPTION_CUSTOM8       = 0x907 //门禁自定义异常8
	MINOR_EXCEPTION_CUSTOM9       = 0x908 //门禁自定义异常9
	MINOR_EXCEPTION_CUSTOM10      = 0x909 //门禁自定义异常10
	MINOR_EXCEPTION_CUSTOM11      = 0x90a //门禁自定义异常11
	MINOR_EXCEPTION_CUSTOM12      = 0x90b //门禁自定义异常12
	MINOR_EXCEPTION_CUSTOM13      = 0x90c //门禁自定义异常13
	MINOR_EXCEPTION_CUSTOM14      = 0x90d //门禁自定义异常14
	MINOR_EXCEPTION_CUSTOM15      = 0x90e //门禁自定义异常15
	MINOR_EXCEPTION_CUSTOM16      = 0x90f //门禁自定义异常16
	MINOR_EXCEPTION_CUSTOM17      = 0x910 //门禁自定义异常17
	MINOR_EXCEPTION_CUSTOM18      = 0x911 //门禁自定义异常18
	MINOR_EXCEPTION_CUSTOM19      = 0x912 //门禁自定义异常19
	MINOR_EXCEPTION_CUSTOM20      = 0x913 //门禁自定义异常20
	MINOR_EXCEPTION_CUSTOM21      = 0x914 //门禁自定义异常21
	MINOR_EXCEPTION_CUSTOM22      = 0x915 //门禁自定义异常22
	MINOR_EXCEPTION_CUSTOM23      = 0x916 //门禁自定义异常23
	MINOR_EXCEPTION_CUSTOM24      = 0x917 //门禁自定义异常24
	MINOR_EXCEPTION_CUSTOM25      = 0x918 //门禁自定义异常25
	MINOR_EXCEPTION_CUSTOM26      = 0x919 //门禁自定义异常26
	MINOR_EXCEPTION_CUSTOM27      = 0x91a //门禁自定义异常27
	MINOR_EXCEPTION_CUSTOM28      = 0x91b //门禁自定义异常28
	MINOR_EXCEPTION_CUSTOM29      = 0x91c //门禁自定义异常29
	MINOR_EXCEPTION_CUSTOM30      = 0x91d //门禁自定义异常30
	MINOR_EXCEPTION_CUSTOM31      = 0x91e //门禁自定义异常31
	MINOR_EXCEPTION_CUSTOM32      = 0x91f //门禁自定义异常32
	MINOR_EXCEPTION_CUSTOM33      = 0x920 //门禁自定义异常33
	MINOR_EXCEPTION_CUSTOM34      = 0x921 //门禁自定义异常34
	MINOR_EXCEPTION_CUSTOM35      = 0x922 //门禁自定义异常35
	MINOR_EXCEPTION_CUSTOM36      = 0x923 //门禁自定义异常36
	MINOR_EXCEPTION_CUSTOM37      = 0x924 //门禁自定义异常37
	MINOR_EXCEPTION_CUSTOM38      = 0x925 //门禁自定义异常38
	MINOR_EXCEPTION_CUSTOM39      = 0x926 //门禁自定义异常39
	MINOR_EXCEPTION_CUSTOM40      = 0x927 //门禁自定义异常40
	MINOR_EXCEPTION_CUSTOM41      = 0x928 //门禁自定义异常41
	MINOR_EXCEPTION_CUSTOM42      = 0x929 //门禁自定义异常42
	MINOR_EXCEPTION_CUSTOM43      = 0x92a //门禁自定义异常43
	MINOR_EXCEPTION_CUSTOM44      = 0x92b //门禁自定义异常44
	MINOR_EXCEPTION_CUSTOM45      = 0x92c //门禁自定义异常45
	MINOR_EXCEPTION_CUSTOM46      = 0x92d //门禁自定义异常46
	MINOR_EXCEPTION_CUSTOM47      = 0x92e //门禁自定义异常47
	MINOR_EXCEPTION_CUSTOM48      = 0x92f //门禁自定义异常48
	MINOR_EXCEPTION_CUSTOM49      = 0x930 //门禁自定义异常49
	MINOR_EXCEPTION_CUSTOM50      = 0x931 //门禁自定义异常50
	MINOR_EXCEPTION_CUSTOM51      = 0x932 //门禁自定义异常51
	MINOR_EXCEPTION_CUSTOM52      = 0x933 //门禁自定义异常52
	MINOR_EXCEPTION_CUSTOM53      = 0x934 //门禁自定义异常53
	MINOR_EXCEPTION_CUSTOM54      = 0x935 //门禁自定义异常54
	MINOR_EXCEPTION_CUSTOM55      = 0x936 //门禁自定义异常55
	MINOR_EXCEPTION_CUSTOM56      = 0x937 //门禁自定义异常56
	MINOR_EXCEPTION_CUSTOM57      = 0x938 //门禁自定义异常57
	MINOR_EXCEPTION_CUSTOM58      = 0x939 //门禁自定义异常58
	MINOR_EXCEPTION_CUSTOM59      = 0x93a //门禁自定义异常59
	MINOR_EXCEPTION_CUSTOM60      = 0x93b //门禁自定义异常60
	MINOR_EXCEPTION_CUSTOM61      = 0x93c //门禁自定义异常61
	MINOR_EXCEPTION_CUSTOM62      = 0x93d //门禁自定义异常62
	MINOR_EXCEPTION_CUSTOM63      = 0x93e //门禁自定义异常63
	MINOR_EXCEPTION_CUSTOM64      = 0x93f //门禁自定义异常64
	MINOR_SWITCH_WIRED_NETWORK    = 0x950 //切换有线网络
	MINOR_SWITCH_WIRELESS_NETWORK = 0x951 //切换无线网络
	MINOR_LOCK_ONLINE_RESUME      = 0x952 //智能锁恢复上线

	//2018-04-23 通用物联网关异常日志类型
	MINOR_ALARMHOST_WDT_RESET     = 0x1003 //WDT 复位
	MINOR_ALARMHOST_RTC_EXCEPTION = 0x1007 //RTC实时时钟异常

	MINOR_ALARMHOST_TEL_LINE_CONNECT_FAILURE = 0x100a //电话线连接断
	MINOR_ALARMHOST_TEL_LINE_CONNECT_RESTORE = 0x100b //电话线连接恢复
	MINOR_ALARMHOST_EXPANDER_BUS_LOSS        = 0x100c //扩展总线模块掉线
	MINOR_ALARMHOST_EXPANDER_BUS_RESTORE     = 0x100d //扩展总线模块掉线恢复
	MINOR_ALARMHOST_KEYPAD_BUS_LOSS          = 0x100e //键盘总线模块掉线
	MINOR_ALARMHOST_KEYPAD_BUS_RESTORE       = 0x100f //键盘总线模块掉线恢复
	MINOR_ALARMHOST_SENSOR_FAILURE           = 0x1010 //模拟量传感器故障
	MINOR_ALARMHOST_SENSOR_RESTORE           = 0x1011 //模拟量传感器恢复
	MINOR_ALARMHOST_RS485_CONNECT_FAILURE    = 0x1012 //RS485通道连接断
	MINOR_ALARMHOST_RS485_CONNECT_RESTORE    = 0x1013 //RS485通道连接断恢复

	//“有线网络异常”和“有线网络异常恢复”这两个日志跟“网络连接断”“网络连接恢复”这两个日志时一样的，且没有设备支持“有线网络异常”和“有线网络异常恢复”这两种类型。
	MINOR_ALARMHOST_WIRED_NETWORK_ABNORMAL = 0x1015 //有线网络异常
	MINOR_ALARMHOST_WIRED_NETWORK_RESTORE  = 0x1016 //有线网络恢复正常
	MINOR_ALARMHOST_GPRS_ABNORMAL          = 0x1017 //GPRS通信异常
	MINOR_ALARMHOST_GPRS_RESTORE           = 0x1018 //GPRS恢复正常
	MINOR_ALARMHOST_3G_ABNORMAL            = 0x1019 //3G通信异常
	MINOR_ALARMHOST_3G_RESTORE             = 0x101a //3G恢复正常
	MINOR_ALARMHOST_SIM_CARD_ABNORMAL      = 0x101b //SIM卡异常
	MINOR_ALARMHOST_SIM_CARD_RESTORE       = 0x101c //SIM卡恢复正常

	MINOR_FORMAT_HDD_ERROR                        = 0x1026 //远程格式化硬盘失败
	MINOR_USB_ERROR                               = 0x1027 //USB通信故障
	MINOR_USB_RESTORE                             = 0x1028 //USB通信故障恢复
	MINOR_PRINT_ERROR                             = 0x1029 //打印机故障
	MINOR_PRINT_RESTORE                           = 0x1030 //打印机故障恢复
	MINOR_ALARMHOST_SUBSYSTEM_COMMUNICATION_ERROR = 0x1031 //子板通讯错误

	MINOR_MCU_RESTART                      = 0x1035 //MCU重启
	MINOR_GPRS_MODULE_FAULT                = 0x1036 //GPRS模块故障
	MINOR_TELEPHONE_MODULE_FAULT           = 0x1037 //电话模块故障
	MINOR_WIFI_ABNORMAL                    = 0x1038 //WIFI通信异常
	MINOR_WIFI_RESTORE                     = 0x1039 //WIFI恢复正常
	MINOR_RF_ABNORMAL                      = 0x103a //RF信号异常
	MINOR_RF_RESTORE                       = 0x103b //RF信号恢复正常
	MINOR_DETECTOR_ONLINE                  = 0x103c //探测器在线
	MINOR_DETECTOR_OFFLINE                 = 0x103d //探测器离线
	MINOR_DETECTOR_BATTERY_NORMAL          = 0x103e //探测器电量正常
	MINOR_DETECTOR_BATTERY_LOW             = 0x103f //探测器电量欠压
	MINOR_DATA_TRAFFIC_OVERFLOW            = 0x1040 //流量超额
	MINOR_ALARMHOST_ZONE_MODULE_LOSS       = 0x1041 //防区模块掉线
	MINOR_ALARMHOST_ZONE_MODULE_RESTORE    = 0x1042 //防区模块掉线恢复
	MINOR_WIRELESS_OUTPUT_LOSS             = 0x1043 //无线输出模块离线
	MINOR_WIRELESS_OUTPUT_RESTORE          = 0x1044 //无线输出模块恢复在线
	MINOR_WIRELESS_REPEATER_LOSS           = 0x1045 //无线中继器离线
	MINOR_WIRELESS_REPEATER_RESTORE        = 0x1046 //无线中继器恢复在线
	MINOR_ALARMHOST_TRIGGER_MODULE_LOSS    = 0x1047 //触发器模块掉线
	MINOR_ALARMHOST_TRIGGER_MODULE_RESTORE = 0x1048 //触发器模块掉线恢复
	MINOR_ALARMHOST_WIRELESS_SIREN_LOSS    = 0x1049 //无线警号离线
	MINOR_ALARMHOST_WIRELESS_SIREN_RESTORE = 0x104a //无线警号恢复在线
	MINOR_TX1_SUB_SYSTEM_EXCEPTION         = 0x1050 /*TX1子系统异常*/
	MINOR_TX1_REBOOT_EXCEPTION             = 0x1051 /*TX1系统异常重启*/
	MINOR_TX1_SUB_SYSTEM_LOSS              = 0x1052 /*智能子系统异常离线*/
	MINOR_TX1_SUB_SYSTEM_RESTORE           = 0x1053 /*智能子系统离线恢复*/
	MINOR_WIRELESS_SPEED_EXCEPTION         = 0x1054 //无线传输速率异常

	//LED 异常次类型 0x1201~0x1300
	MINOR_LED_SYSTEM_EXCEPTION      = 0x1201 //LED系统异常
	MINOR_FLASH_NOTENOUGH_EXCEPTION = 0x1202 //FLASH空间不足

	MINOR_LOG_EXCEPTION = 0x1301 //日志盘异常

	//[add]by silujie 2013-3-22 14:16
	//0x2000~0x3fff 为设备报警日志
	//0x4000~0x5000 为设备异常日志
	MINOR_SUBSYSTEM_IP_CONFLICT          = 0x4000 //子板IP冲突
	MINOR_SUBSYSTEM_NET_BROKEN           = 0x4001 //子板断网
	MINOR_FAN_ABNORMAL                   = 0x4002 //风扇异常
	MINOR_BACKPANEL_TEMPERATURE_ABNORMAL = 0x4003 //背板温度异常

	MINOR_SDCARD_ABNORMAL               = 0x4004 //SD卡不健康
	MINOR_SDCARD_DAMAGE                 = 0x4005 //SD卡损坏
	MINOR_POC_ABNORMAL                  = 0x4006 //设备POC模块异常
	MINOR_MAIN_POWER_FAULT              = 0x4007 //主电故障
	MINOR_BACK_UP_POWER_FAULT           = 0x4008 //备电故障
	MINOR_TAMPER_FAULT                  = 0x4009 //防拆故障
	MINOR_RS232_FAULT                   = 0x400a //232总线故障
	MINOR_RS485_FAULT                   = 0x400b //485总线故障
	MINOR_LAN_STATUS_FAULT              = 0x400c //LAN网线接入状态故障
	MINOR_LAN_LINK1_FAULT               = 0x400d //LAN链路1故障
	MINOR_LAN_LINK2_FAULT               = 0x400e //LAN链路2故障
	MINOR_SIM_CARD_STATUS_FAULT         = 0x400f //4G-SIM卡状态故障
	MINOR_4G_LINK1_FAULT                = 0x4010 //4G链路1故障
	MINOR_4G_LINK2_FAULT                = 0x4011 //4G链路2故障
	MINOR_OTHER_FAULT                   = 0x4012 //其他故障
	MINOR_FIRE_CONTROL_CONNECT_FAULT    = 0x4013 //与消控主机连接故障
	MINOR_SENSOR_SHORT_CIRCUIT          = 0x4014 //传感器短路
	MINOR_SENSOR_OPEN_CIRCUIT           = 0x4015 //传感器断路
	MINOR_SENSOR_MIS_CONNECT            = 0x4016 //传感器错接
	MINOR_SENSOR_FAULT_RESTORE          = 0x4017 //传感器故障恢复
	MINOR_DEVICE_FAULT                  = 0x4018 //设备故障
	MINOR_OVERVOLTAGE                   = 0x4019 //电源电压过高
	MINOR_UNDERVOLTAGE                  = 0x401a //电源电压过低
	MINOR_PANLOCKING                    = 0x401b //云台水平堵转
	MINOR_TILTLOCKING                   = 0x401c //云台垂直堵转
	MINOR_SUBBOARD_TEMPERATURE_ABNORMAL = 0x401d //子板温度异常
	MINOR_EZVIZ_UPGRADE_EXCEPTION       = 0x401e //萤石升级异常

	//萤石相关操作异常日志
	MINOR_EZVIZ_OPERATION_ABNORMAL = 0x4020 //萤石操作异常

	/* 操作 */
	//主类型
	MAJOR_OPERATION = 0x3

	//次类型
	MINOR_VCA_MOTIONEXCEPTION = 0x29 //智能侦测异常
	MINOR_START_DVR           = 0x41 /* 开机 */
	MINOR_STOP_DVR            = 0x42 /* 关机 */
	MINOR_STOP_ABNORMAL       = 0x43 /* 异常关机 */
	MINOR_REBOOT_DVR          = 0x44 /*本地重启设备*/

	MINOR_LOCAL_LOGIN               = 0x50 /* 本地登陆 */
	MINOR_LOCAL_LOGOUT              = 0x51 /* 本地注销登陆 */
	MINOR_LOCAL_CFG_PARM            = 0x52 /* 本地配置参数 */
	MINOR_LOCAL_PLAYBYFILE          = 0x53 /* 本地按文件回放或下载 */
	MINOR_LOCAL_PLAYBYTIME          = 0x54 /* 本地按时间回放或下载*/
	MINOR_LOCAL_START_REC           = 0x55 /* 本地开始录像 */
	MINOR_LOCAL_STOP_REC            = 0x56 /* 本地停止录像 */
	MINOR_LOCAL_PTZCTRL             = 0x57 /* 本地云台控制 */
	MINOR_LOCAL_PREVIEW             = 0x58 /* 本地预览 (保留不使用)*/
	MINOR_LOCAL_MODIFY_TIME         = 0x59 /* 本地修改时间(保留不使用) */
	MINOR_LOCAL_UPGRADE             = 0x5a /* 本地升级 */
	MINOR_LOCAL_RECFILE_OUTPUT      = 0x5b /* 本地备份录象文件 */
	MINOR_LOCAL_FORMAT_HDD          = 0x5c /* 本地初始化硬盘 */
	MINOR_LOCAL_CFGFILE_OUTPUT      = 0x5d /* 导出本地配置文件 */
	MINOR_LOCAL_CFGFILE_INPUT       = 0x5e /* 导入本地配置文件 */
	MINOR_LOCAL_COPYFILE            = 0x5f /* 本地备份文件 */
	MINOR_LOCAL_LOCKFILE            = 0x60 /* 本地锁定录像文件 */
	MINOR_LOCAL_UNLOCKFILE          = 0x61 /* 本地解锁录像文件 */
	MINOR_LOCAL_DVR_ALARM           = 0x62 /* 本地手动清除和触发报警*/
	MINOR_IPC_ADD                   = 0x63 /* 本地添加IPC */
	MINOR_IPC_DEL                   = 0x64 /* 本地删除IPC */
	MINOR_IPC_SET                   = 0x65 /* 本地设置IPC */
	MINOR_LOCAL_START_BACKUP        = 0x66 /* 本地开始备份 */
	MINOR_LOCAL_STOP_BACKUP         = 0x67 /* 本地停止备份*/
	MINOR_LOCAL_COPYFILE_START_TIME = 0x68 /* 本地备份开始时间*/
	MINOR_LOCAL_COPYFILE_END_TIME   = 0x69 /* 本地备份结束时间*/
	MINOR_LOCAL_ADD_NAS             = 0x6a /*本地添加网络硬盘 （nfs、iscsi）*/
	MINOR_LOCAL_DEL_NAS             = 0x6b /* 本地删除nas盘 （nfs、iscsi）*/
	MINOR_LOCAL_SET_NAS             = 0x6c /* 本地设置nas盘 （nfs、iscsi）*/
	MINOR_LOCAL_RESET_PASSWD        = 0x6d /* 本地恢复管理员默认密码*/

	MINOR_REMOTE_LOGIN              = 0x70  /* 远程登录 */
	MINOR_REMOTE_LOGOUT             = 0x71  /* 远程注销登陆 */
	MINOR_REMOTE_START_REC          = 0x72  /* 远程开始录像 */
	MINOR_REMOTE_STOP_REC           = 0x73  /* 远程停止录像 */
	MINOR_START_TRANS_CHAN          = 0x74  /* 开始透明传输 */
	MINOR_STOP_TRANS_CHAN           = 0x75  /* 停止透明传输 */
	MINOR_REMOTE_GET_PARM           = 0x76  /* 远程获取参数 */
	MINOR_REMOTE_CFG_PARM           = 0x77  /* 远程配置参数 */
	MINOR_REMOTE_GET_STATUS         = 0x78  /* 远程获取状态 */
	MINOR_REMOTE_ARM                = 0x79  /* 远程布防 */
	MINOR_REMOTE_DISARM             = 0x7a  /* 远程撤防 */
	MINOR_REMOTE_REBOOT             = 0x7b  /* 远程重启 */
	MINOR_START_VT                  = 0x7c  /* 开始语音对讲 */
	MINOR_STOP_VT                   = 0x7d  /* 停止语音对讲 */
	MINOR_REMOTE_UPGRADE            = 0x7e  /* 远程升级 */
	MINOR_REMOTE_PLAYBYFILE         = 0x7f  /* 远程按文件回放 */
	MINOR_REMOTE_PLAYBYTIME         = 0x80  /* 远程按时间回放 */
	MINOR_REMOTE_PTZCTRL            = 0x81  /* 远程云台控制 */
	MINOR_REMOTE_FORMAT_HDD         = 0x82  /* 远程格式化硬盘 */
	MINOR_REMOTE_STOP               = 0x83  /* 远程关机 */
	MINOR_REMOTE_LOCKFILE           = 0x84  /* 远程锁定文件 */
	MINOR_REMOTE_UNLOCKFILE         = 0x85  /* 远程解锁文件 */
	MINOR_REMOTE_CFGFILE_OUTPUT     = 0x86  /* 远程导出配置文件 */
	MINOR_REMOTE_CFGFILE_INTPUT     = 0x87  /* 远程导入配置文件 */
	MINOR_REMOTE_RECFILE_OUTPUT     = 0x88  /* 远程导出录象文件 */
	MINOR_REMOTE_DVR_ALARM          = 0x89  /* 远程手动清除和触发报警*/
	MINOR_REMOTE_IPC_ADD            = 0x8a  /* 远程添加IPC */
	MINOR_REMOTE_IPC_DEL            = 0x8b  /* 远程删除IPC */
	MINOR_REMOTE_IPC_SET            = 0x8c  /* 远程设置IPC */
	MINOR_REBOOT_VCA_LIB            = 0x8d  /*重启智能库*/
	MINOR_REMOTE_ADD_NAS            = 0x8e  /* 远程添加nas盘 （nfs、iscsi）*/
	MINOR_REMOTE_DEL_NAS            = 0x8f  /* 远程删除nas盘 （nfs、iscsi）*/
	MINOR_REMOTE_SET_NAS            = 0x90  /* 远程设置nas盘 （nfs、iscsi）*/
	MINOR_LOCAL_OPERATE_LOCK        = 0x9d  /* 本地操作锁定             */
	MINOR_LOCAL_OPERATE_UNLOCK      = 0x9e  /* 本地操作解除锁定         */
	MINOR_REMOTE_DELETE_HDISK       = 0x9a  /* 远程删除异常不存在的硬盘 */
	MINOR_REMOTE_LOAD_HDISK         = 0x9b  /* 远程加载硬盘             */
	MINOR_REMOTE_UNLOAD_HDISK       = 0x9c  /* 远程卸载硬盘   */
	MINOR_SCHEDULE_ANGLECALIBRATION = 0x139 /*定期倾角校准*/
	MINOR_OTHER_OPERATE             = 0x200 /* 其他操作 */

	//2010-05-26 增加审讯DVR日志类型
	MINOR_LOCAL_START_REC_CDRW  = 0x91 /* 本地开始讯问 */
	MINOR_LOCAL_STOP_REC_CDRW   = 0x92 /* 本地停止讯问 */
	MINOR_REMOTE_START_REC_CDRW = 0x93 /* 远程开始讯问 */
	MINOR_REMOTE_STOP_REC_CDRW  = 0x94 /* 远程停止讯问 */

	MINOR_LOCAL_PIC_OUTPUT  = 0x95 /* 本地备份图片文件 */
	MINOR_REMOTE_PIC_OUTPUT = 0x96 /* 远程备份图片文件 */

	//2011-07-26 增加81审讯DVR日志类型
	MINOR_LOCAL_INQUEST_RESUME  = 0x97 /* 本地恢复审讯事件*/
	MINOR_REMOTE_INQUEST_RESUME = 0x98 /* 远程恢复审讯事件*/

	//2013-01-23 增加86高清审讯NVR操作日志
	MINOR_LOCAL_ADD_FILE          = 0x99  /*本地导入文件*/
	MINOR_LOCAL_DEL_FILE          = 0x9f  /*本地删除审讯*/
	MINOR_REMOTE_INQUEST_ADD_FILE = 0x100 /*远程导入文件*/

	//2009-12-16 增加视频综合平台日志类型
	MINOR_SUBSYSTEMREBOOT           = 0xa0 /*视频综合平台：dm6467 正常重启*/
	MINOR_MATRIX_STARTTRANSFERVIDEO = 0xa1 /*视频综合平台：矩阵切换开始传输图像*/
	MINOR_MATRIX_STOPTRANSFERVIDEO  = 0xa2 /*视频综合平台：矩阵切换停止传输图像*/
	MINOR_REMOTE_SET_ALLSUBSYSTEM   = 0xa3 /*视频综合平台：设置所有6467子系统信息*/
	MINOR_REMOTE_GET_ALLSUBSYSTEM   = 0xa4 /*视频综合平台：获取所有6467子系统信息*/
	MINOR_REMOTE_SET_PLANARRAY      = 0xa5 /*视频综合平台：设置计划轮巡组*/
	MINOR_REMOTE_GET_PLANARRAY      = 0xa6 /*视频综合平台：获取计划轮巡组*/
	MINOR_MATRIX_STARTTRANSFERAUDIO = 0xa7 /*视频综合平台：矩阵切换开始传输音频*/
	MINOR_MATRIX_STOPRANSFERAUDIO   = 0xa8 /*视频综合平台：矩阵切换停止传输音频*/
	MINOR_LOGON_CODESPITTER         = 0xa9 /*视频综合平台：登陆码分器*/
	MINOR_LOGOFF_CODESPITTER        = 0xaa /*视频综合平台：退出码分器*/

	//2010-01-22 增加视频综合平台中解码器操作日志
	MINOR_START_DYNAMIC_DECODE = 0xb0 /*开始动态解码*/
	MINOR_STOP_DYNAMIC_DECODE  = 0xb1 /*停止动态解码*/
	MINOR_GET_CYC_CFG          = 0xb2 /*获取解码器通道轮巡配置*/
	MINOR_SET_CYC_CFG          = 0xb3 /*设置解码通道轮巡配置*/
	MINOR_START_CYC_DECODE     = 0xb4 /*开始轮巡解码*/
	MINOR_STOP_CYC_DECODE      = 0xb5 /*停止轮巡解码*/
	MINOR_GET_DECCHAN_STATUS   = 0xb6 /*获取解码通道状态*/
	MINOR_GET_DECCHAN_INFO     = 0xb7 /*获取解码通道当前信息*/
	MINOR_START_PASSIVE_DEC    = 0xb8 /*开始被动解码*/
	MINOR_STOP_PASSIVE_DEC     = 0xb9 /*停止被动解码*/
	MINOR_CTRL_PASSIVE_DEC     = 0xba /*控制被动解码*/
	MINOR_RECON_PASSIVE_DEC    = 0xbb /*被动解码重连*/
	MINOR_GET_DEC_CHAN_SW      = 0xbc /*获取解码通道总开关*/
	MINOR_SET_DEC_CHAN_SW      = 0xbd /*设置解码通道总开关*/
	MINOR_CTRL_DEC_CHAN_SCALE  = 0xbe /*解码通道缩放控制*/
	MINOR_SET_REMOTE_REPLAY    = 0xbf /*设置远程回放*/
	MINOR_GET_REMOTE_REPLAY    = 0xc0 /*获取远程回放状态*/
	MINOR_CTRL_REMOTE_REPLAY   = 0xc1 /*远程回放控制*/
	MINOR_SET_DISP_CFG         = 0xc2 /*设置显示通道*/
	MINOR_GET_DISP_CFG         = 0xc3 /*获取显示通道设置*/
	MINOR_SET_PLANTABLE        = 0xc4 /*设置计划轮巡表*/
	MINOR_GET_PLANTABLE        = 0xc5 /*获取计划轮巡表*/
	MINOR_START_PPPPOE         = 0xc6 /*开始PPPoE连接*/
	MINOR_STOP_PPPPOE          = 0xc7 /*结束PPPoE连接*/
	MINOR_UPLOAD_LOGO          = 0xc8 /*上传LOGO*/
	//推模式操作日志
	MINOR_LOCAL_PIN   = 0xc9 /* 本地PIN功能操作 */
	MINOR_LOCAL_DIAL  = 0xca /* 本地手动启动断开拨号 */
	MINOR_SMS_CONTROL = 0xcb /* 短信控制上下线 */
	MINOR_CALL_ONLINE = 0xcc /* 呼叫控制上线 */
	MINOR_REMOTE_PIN  = 0xcd /* 远程PIN功能操作 */

	//2010-12-16 报警板日志
	MINOR_REMOTE_BYPASS             = 0xd0 /* 远程旁路*/
	MINOR_REMOTE_UNBYPASS           = 0xd1 /* 远程旁路恢复*/
	MINOR_REMOTE_SET_ALARMIN_CFG    = 0xd2 /* 远程设置报警输入参数*/
	MINOR_REMOTE_GET_ALARMIN_CFG    = 0xd3 /* 远程获取报警输入参数*/
	MINOR_REMOTE_SET_ALARMOUT_CFG   = 0xd4 /* 远程设置报警输出参数*/
	MINOR_REMOTE_GET_ALARMOUT_CFG   = 0xd5 /* 远程获取报警输出参数*/
	MINOR_REMOTE_ALARMOUT_OPEN_MAN  = 0xd6 /* 远程手动开启报警输出*/
	MINOR_REMOTE_ALARMOUT_CLOSE_MAN = 0xd7 /* 远程手动关闭报警输出*/
	MINOR_REMOTE_ALARM_ENABLE_CFG   = 0xd8 /* 远程设置报警主机的RS485串口使能状态*/
	MINOR_DBDATA_OUTPUT             = 0xd9 /* 导出数据库记录 */
	MINOR_DBDATA_INPUT              = 0xda /* 导入数据库记录 */
	MINOR_MU_SWITCH                 = 0xdb /* 级联切换 */
	MINOR_MU_PTZ                    = 0xdc /* 级联PTZ控制 */
	MINOR_DELETE_LOGO               = 0xdd /* 删除logo */
	MINOR_REMOTE_INQUEST_DEL_FILE   = 0xde /*远程删除文件*/

	MINOR_LOCAL_CONF_REB_RAID    = 0x101 /*本地配置自动重建*/
	MINOR_LOCAL_CONF_SPARE       = 0x102 /*本地配置热备*/
	MINOR_LOCAL_ADD_RAID         = 0x103 /*本地创建阵列*/
	MINOR_LOCAL_DEL_RAID         = 0x104 /*本地删除阵列*/
	MINOR_LOCAL_MIG_RAID         = 0x105 /*本地迁移阵列*/
	MINOR_LOCAL_REB_RAID         = 0x106 /* 本地手动重建阵列*/
	MINOR_LOCAL_QUICK_CONF_RAID  = 0x107 /*本地一键配置*/
	MINOR_LOCAL_ADD_VD           = 0x108 /*本地创建虚拟磁盘*/
	MINOR_LOCAL_DEL_VD           = 0x109 /*本地删除虚拟磁盘*/
	MINOR_LOCAL_RP_VD            = 0x10a /*本地修复虚拟磁盘*/
	MINOR_LOCAL_FORMAT_EXPANDVD  = 0x10b /*本地扩展虚拟磁盘扩容*/
	MINOR_LOCAL_RAID_UPGRADE     = 0x10c /*本地raid卡升级*/
	MINOR_LOCAL_STOP_RAID        = 0x10d /*本地暂停RAID操作(即安全拔盘)*/
	MINOR_REMOTE_CONF_REB_RAID   = 0x111 /*远程配置自动重建*/
	MINOR_REMOTE_CONF_SPARE      = 0x112 /*远程配置热备*/
	MINOR_REMOTE_ADD_RAID        = 0x113 /*远程创建阵列*/
	MINOR_REMOTE_DEL_RAID        = 0x114 /*远程删除阵列*/
	MINOR_REMOTE_MIG_RAID        = 0x115 /*远程迁移阵列*/
	MINOR_REMOTE_REB_RAID        = 0x116 /* 远程手动重建阵列*/
	MINOR_REMOTE_QUICK_CONF_RAID = 0x117 /*远程一键配置*/
	MINOR_REMOTE_ADD_VD          = 0x118 /*远程创建虚拟磁盘*/
	MINOR_REMOTE_DEL_VD          = 0x119 /*远程删除虚拟磁盘*/
	MINOR_REMOTE_RP_VD           = 0x11a /*远程修复虚拟磁盘*/
	MINOR_REMOTE_FORMAT_EXPANDVD = 0x11b /*远程虚拟磁盘扩容*/
	MINOR_REMOTE_RAID_UPGRADE    = 0x11c /*远程raid卡升级*/
	MINOR_REMOTE_STOP_RAID       = 0x11d /*远程暂停RAID操作(即安全拔盘)*/
	MINOR_LOCAL_START_PIC_REC    = 0x121 /*本地开始抓图*/
	MINOR_LOCAL_STOP_PIC_REC     = 0x122 /*本地停止抓图*/
	MINOR_LOCAL_SET_SNMP         = 0x125 /*本地配置SNMP*/
	MINOR_LOCAL_TAG_OPT          = 0x126 /*本地标签操作*/
	MINOR_REMOTE_START_PIC_REC   = 0x131 /*远程开始抓图*/
	MINOR_REMOTE_STOP_PIC_REC    = 0x132 /*远程停止抓图*/
	MINOR_REMOTE_SET_SNMP        = 0x135 /*远程配置SNMP*/
	MINOR_REMOTE_TAG_OPT         = 0x136 /*远程标签操作*/
	MINOR_REMOTE_LOGIN_LOCK      = 0x137 //远程登录锁定
	MINOR_REMOTE_LOGIN_UNLOCK    = 0x138 //远程登录解锁
	// 9000 v2.2.0
	MINOR_LOCAL_VOUT_SWITCH = 0x140 /* 本地输出口切换操作*/
	MINOR_STREAM_CABAC      = 0x141 /* 码流压缩性能选项配置操作*/

	//Netra 3.0.0
	MINOR_LOCAL_SPARE_OPT          = 0x142 /*本地N+1 热备相关操作*/
	MINOR_REMOTE_SPARE_OPT         = 0x143 /*远程N+1 热备相关操作*/
	MINOR_LOCAL_IPCCFGFILE_OUTPUT  = 0x144 /* 本地导出ipc配置文件*/
	MINOR_LOCAL_IPCCFGFILE_INPUT   = 0x145 /* 本地导入ipc配置文件 */
	MINOR_LOCAL_IPC_UPGRADE        = 0x146 /* 本地升级IPC */
	MINOR_REMOTE_IPCCFGFILE_OUTPUT = 0x147 /* 远程导出ipc配置文件*/
	MINOR_REMOTE_IPCCFGFILE_INPUT  = 0x148 /* 远程导入ipc配置文件*/
	MINOR_REMOTE_IPC_UPGRADE       = 0x149 /* 远程升级IPC */

	MINOR_LOCAL_UNLOAD_HDISK              = 0x150 /*本地卸载硬盘*/
	MINOR_LOCAL_AUDIO_MIX                 = 0x151 /*本地配置音频混音参数*/
	MINOR_REMOTE_AUDIO_MIX                = 0x152 /*远程配置音频混音参数*/
	MINOR_LOCAL_TRIAL_PAUSE               = 0x153 /*本地暂停讯问*/
	MINOR_LOCAL_TRIAL_RESUME              = 0x154 /*本地继续讯问*/
	MINOR_REMOTE_TRIAL_PAUSE              = 0x155 /*远程暂停讯问*/
	MINOR_REMOTE_TRIAL_RESUME             = 0x156 /*远程继续讯问*/
	MINOR_REMOTE_MODIFY_VERIFICATION_CODE = 0x157 /*修改平台的验证码*/

	MINOR_LOCAL_MAKECALL   = 0x180 /*本地呼叫*/
	MINOR_LOCAL_REJECTCALL = 0x181 /*本地拒接*/
	MINOR_LOCAL_ANSWERCALL = 0x182 /*本地接听*/
	MINOR_LOCAL_HANGUPCALL = 0x183 /*本地挂断*/

	MINOR_REMOTE_MAKECALL   = 0x188 /*远程呼叫*/
	MINOR_REMOTE_REJECTCALL = 0x189 /*远程拒接*/
	MINOR_REMOTE_ANSWERCALL = 0x18a /*远程接听*/
	MINOR_REMOTE_HANGUPCALL = 0x18b /*远程挂断*/

	MINOR_LOCAL_CHANNEL_ORDERED = 0x19b /*本地通道排序*/

	MINOR_SET_MULTI_MASTER    = 0x201 /*设置大屏主屏*/
	MINOR_SET_MULTI_SLAVE     = 0x202 /*设置大屏子屏*/
	MINOR_CANCEL_MULTI_MASTER = 0x203 /*取消大屏主屏*/
	MINOR_CANCEL_MULTI_SLAVE  = 0x204 /*取消大屏子屏*/

	MINOR_DISPLAY_LOGO              = 0x205 /*显示LOGO*/
	MINOR_HIDE_LOGO                 = 0x206 /*隐藏LOGO*/
	MINOR_SET_DEC_DELAY_LEVEL       = 0x207 /*解码通道延时级别设置*/
	MINOR_SET_BIGSCREEN_DIPLAY_AREA = 0x208 /*设置大屏显示区域*/
	MINOR_CUT_VIDEO_SOURCE          = 0x209 /*大屏视频源切割设置*/
	MINOR_SET_BASEMAP_AREA          = 0x210 /*大屏底图区域设置*/
	MINOR_DOWNLOAD_BASEMAP          = 0x211 /*下载大屏底图*/
	MINOR_CUT_BASEMAP               = 0x212 /*底图切割配置*/
	MINOR_CONTROL_ELEC_ENLARGE      = 0x213 /*电子放大操作(放大或还原)*/
	MINOR_SET_OUTPUT_RESOLUTION     = 0x214 /*显示输出分辨率设置*/
	MINOR_SET_TRANCSPARENCY         = 0x215 /*图层透明度设置*/
	MINOR_SET_OSD                   = 0x216 /*显示OSD设置*/
	MINOR_RESTORE_DEC_STATUS        = 0x217 /*恢复初始状态(场景切换时，解码恢复初始状态)*/

	//2011-11-11 增加大屏控制器操作日志次类型
	MINOR_SCREEN_OPEN_SCREEN            = 0x218 //打开屏幕
	MINOR_SCREEN_CLOSE_SCREEN           = 0x219 //关闭屏幕
	MINOR_SCREEN_SWITCH_SIGNAL          = 0x21a //信号源切换
	MINOR_SCREEN_MODIFY_NETWORK         = 0x21b //配置网络参数
	MINOR_SCREEN_MODIFY_LEDRES          = 0x21c //配置输出口LED分辨率
	MINOR_SCREEN_SHOW_NORMAL            = 0x21d //配置窗口普通显示模式
	MINOR_SCREEN_SHOW_TILE              = 0x21e //配置窗口平铺显示模式
	MINOR_SCREEN_DEC_NORMAL             = 0x21f //配置普通解码模式
	MINOR_SCREEN_DEC_LOWLATENCY         = 0x220 //配置低延时解码模式
	MINOR_SCREEN_MODIFY_SELFRES         = 0x221 //配置信号源自定义分辨率
	MINOR_SCREEN_OUTPUT_POSITION        = 0x222 //输出口关联屏幕
	MINOR_SCREEN_IMAGE_ENHANCE          = 0x223 //图像增强
	MINOR_SCREEN_JOIN_SIGNAL            = 0x224 //信号源拼接
	MINOR_SCREEN_SIGNAL_OSD             = 0x225 //信号源字符叠加
	MINOR_SCREEN_ASSOCIATED_INTERACTION = 0x226 //信号源关联多屏互动服务器
	MINOR_SCREEN_MODIFY_MATRIX          = 0x227 //配置矩阵参数
	MINOR_SCREEN_WND_TOP_KEEP           = 0x228 //窗口置顶保持
	MINOR_SCREEN_WND_OPEN_KEEP          = 0x229 //窗口打开保持
	MINOR_SCREEN_WALL_MIRROR            = 0x22a //电视墙区域镜像
	MINOR_SCREEN_UPLOAD_BASEMAP         = 0x22b //上传底图
	MINOR_SCREEN_SHOW_BASEMAP           = 0x22c //显示底图
	MINOR_SCREEN_HIDE_BASEMAP           = 0x22d //隐藏底图
	MINOR_SCREEN_MODIFY_SERIAL          = 0x22e //配置串口参数

	MINOR_SCREEN_SET_INPUT      = 0x251 /*修改输入源*/
	MINOR_SCREEN_SET_OUTPUT     = 0x252 /*修改输出通道*/
	MINOR_SCREEN_SET_OSD        = 0x253 /*修改虚拟LED*/
	MINOR_SCREEN_SET_LOGO       = 0x254 /*修改LOGO*/
	MINOR_SCREEN_SET_LAYOUT     = 0x255 /*设置布局*/
	MINOR_SCREEN_PICTUREPREVIEW = 0x256 /*回显操作*/

	//2012-06-14 CVCS2.0, 窗口设置等操作在V1.0， V1.1中已经有了，当时在设备日志中没有定义
	MINOR_SCREEN_GET_OSD             = 0x257 /*获取虚拟LED*/
	MINOR_SCREEN_GET_LAYOUT          = 0x258 /*获取布局*/
	MINOR_SCREEN_LAYOUT_CTRL         = 0x259 /*布局控制*/
	MINOR_GET_ALL_VALID_WND          = 0x260 /*获取所有有效窗口*/
	MINOR_GET_SIGNAL_WND             = 0x261 /*获取单个窗口信息*/
	MINOR_WINDOW_CTRL                = 0x262 /*窗口控制*/
	MINOR_GET_LAYOUT_LIST            = 0x263 /*获取布局列表*/
	MINOR_LAYOUT_CTRL                = 0x264 /*布局控制*/
	MINOR_SET_LAYOUT                 = 0x265 /*设置布局*/
	MINOR_GET_SIGNAL_LIST            = 0x266 /*获取输入信号源列表*/
	MINOR_GET_PLAN_LIST              = 0x267 /*获取预案列表*/
	MINOR_SET_PLAN                   = 0x268 /*修改预案*/
	MINOR_CTRL_PLAN                  = 0x269 /*控制预案*/
	MINOR_CTRL_SCREEN                = 0x270 /*屏幕控制*/
	MINOR_ADD_NETSIG                 = 0x271 /*添加信号源*/
	MINOR_SET_NETSIG                 = 0x272 /*修改信号源*/
	MINOR_SET_DECBDCFG               = 0x273 /*设置解码板参数*/
	MINOR_GET_DECBDCFG               = 0x274 /*获取解码板参数*/
	MINOR_GET_DEVICE_STATUS          = 0x275 /*获取设备信息*/
	MINOR_UPLOAD_PICTURE             = 0x276 /*底图上传*/
	MINOR_SET_USERPWD                = 0x277 /*设置用户密码*/
	MINOR_ADD_LAYOUT                 = 0x278 /*添加布局*/
	MINOR_DEL_LAYOUT                 = 0x279 /*删除布局*/
	MINOR_DEL_NETSIG                 = 0x280 /*删除信号源*/
	MINOR_ADD_PLAN                   = 0x281 /*添加预案*/
	MINOR_DEL_PLAN                   = 0x282 /*删除预案*/
	MINOR_GET_EXTERNAL_MATRIX_CFG    = 0x283 //获取外接矩阵配置
	MINOR_SET_EXTERNAL_MATRIX_CFG    = 0x284 //设置外接矩阵配置
	MINOR_GET_USER_CFG               = 0x285 //获取用户配置
	MINOR_SET_USER_CFG               = 0x286 //设置用户配置
	MINOR_GET_DISPLAY_PANEL_LINK_CFG = 0x287 //获取显示墙连接配置
	MINOR_SET_DISPLAY_PANEL_LINK_CFG = 0x288 //设置显示墙连接配置

	MINOR_GET_WALLSCENE_PARAM   = 0x289 //获取电视墙场景
	MINOR_SET_WALLSCENE_PARAM   = 0x28a //设置电视墙场景
	MINOR_GET_CURRENT_WALLSCENE = 0x28b //获取当前使用场景
	MINOR_SWITCH_WALLSCENE      = 0x28c //场景切换
	MINOR_SIP_LOGIN             = 0x28d //SIP注册成功
	MINOR_VOIP_START            = 0x28e //VOIP对讲开始
	MINOR_VOIP_STOP             = 0x28f //VOIP对讲停止
	MINOR_WIN_TOP               = 0x290 //电视墙窗口置顶
	MINOR_WIN_BOTTOM            = 0x291 //电视墙窗口置底
	MINOR_SET_USER_ADD_CFG      = 0x292 //增加用户
	MINOR_SET_USER_MODF_CFG     = 0x293 //修改用户
	MINOR_SET_USER_DEL_CFG      = 0x294 //删除用户

	// Netra 2.2.2
	MINOR_LOCAL_LOAD_HDISK   = 0x300 //本地加载硬盘
	MINOR_LOCAL_DELETE_HDISK = 0x301 //本地删除异常不存在的硬盘

	//KY2013 3.0.0
	MINOR_LOCAL_MAIN_AUXILIARY_PORT_SWITCH = 0x302 //本地主辅口切换
	MINOR_LOCAL_HARD_DISK_CHECK            = 0x303 //本地物理硬盘自检

	//Netra3.1.0
	MINOR_LOCAL_CFG_DEVICE_TYPE      = 0x310 //本地配置设备类型
	MINOR_REMOTE_CFG_DEVICE_TYPE     = 0x311 //远程配置设备类型
	MINOR_LOCAL_CFG_WORK_HOT_SERVER  = 0x312 //本地配置工作机热备服务器
	MINOR_REMOTE_CFG_WORK_HOT_SERVER = 0x313 //远程配置工作机热备服务器
	MINOR_LOCAL_DELETE_WORK          = 0x314 //本地删除工作机
	MINOR_REMOTE_DELETE_WORK         = 0x315 //远程删除工作机
	MINOR_LOCAL_ADD_WORK             = 0x316 //本地添加工作机
	MINOR_REMOTE_ADD_WORK            = 0x317 //远程添加工作机
	MINOR_LOCAL_IPCHEATMAP_OUTPUT    = 0x318 /* 本地导出热度图文件      */
	MINOR_LOCAL_IPCHEATFLOW_OUTPUT   = 0x319 /* 本地导出热度流量文件      */
	MINOR_REMOTE_SMS_SEND            = 0x350 /*远程发送短信*/
	MINOR_LOCAL_SMS_SEND             = 0x351 /*本地发送短信*/
	MINOR_ALARM_SMS_SEND             = 0x352 /*发送短信报警*/
	MINOR_SMS_RECV                   = 0x353 /*接收短信*/
	//（备注：0x350、0x351是指人工在GUI或IE控件上编辑并发送短信）
	MINOR_LOCAL_SMS_SEARCH                   = 0x354 /*本地搜索短信*/
	MINOR_REMOTE_SMS_SEARCH                  = 0x355 /*远程搜索短信*/
	MINOR_LOCAL_SMS_READ                     = 0x356 /*本地查看短信*/
	MINOR_REMOTE_SMS_READ                    = 0x357 /*远程查看短信*/
	MINOR_REMOTE_DIAL_CONNECT                = 0x358 /*远程开启手动拨号*/
	MINOR_REMOTE_DIAL_DISCONN                = 0x359 /*远程停止手动拨号*/
	MINOR_LOCAL_WHITELIST_SET                = 0x35A /*本地配置白名单*/
	MINOR_REMOTE_WHITELIST_SET               = 0x35B /*远程配置白名单*/
	MINOR_LOCAL_DIAL_PARA_SET                = 0x35C /*本地配置拨号参数*/
	MINOR_REMOTE_DIAL_PARA_SET               = 0x35D /*远程配置拨号参数*/
	MINOR_LOCAL_DIAL_SCHEDULE_SET            = 0x35E /*本地配置拨号计划*/
	MINOR_REMOTE_DIAL_SCHEDULE_SET           = 0x35F /*远程配置拨号计划*/
	MINOR_PLAT_OPER                          = 0x360 /* 平台操作*/
	MINOR_REMOTE_CFG_POE_WORK_MODE           = 0x361 //远程设置POE工作模式
	MINOR_LOCAL_CFG_POE_WORK_MODE            = 0x362 //本地设置POE工作模式
	MINOR_REMOTE_CFG_FACE_CONTRAST           = 0x363 //远程设置人脸比对配置
	MINOR_LOCAL_CFG_FACE_CONTRAST            = 0x364 //本地设置人脸比对配置
	MINOR_REMOTE_CFG_WHITELIST_FACE_CONTRAST = 0x365 //远程设置白名单人脸比对配置
	MINOR_LOCAL_CFG_WHITELIST_FACE_CONTRAST  = 0x366 //本地设置白名单人脸比对配置
	MINOR_LOCAL_CHECK_TIME                   = 0x367 //本地手动校时
	MINOR_VCA_ONEKEY_EXPORT_PICTURE          = 0x368 //一键导出图片
	MINOR_VCA_ONEKEY_DELETE_PICTURE          = 0x369 //一键删除图片
	MINOR_VCA_ONEKEY_EXPORT_VIDEO            = 0x36a //一键导出录像
	MINOR_VCA_ONEKEY_DELETE_VIDEO            = 0x36b //一键删除录像
	MINOR_REMOTE_CFG_WIRELESS_DIALPARAM      = 0x36c /*远程配置无线拨号参数*/
	MINOR_LOCAL_CFG_WIRELESS_DIALPARAM       = 0x36d /*本地配置无线拨号参数*/
	MINOR_REMOTE_CFG_WIRELESS_SMSPARAM       = 0x36e /*远程配置无线短信配置参数*/
	MINOR_LOCAL_CFG_WIRELESS_SMSPARAM        = 0x36f /*本地配置无线短信配置参数*/
	MINOR_REMOTE_CFG_WIRELESS_SMSSElFHELP    = 0x370 /*远程配置无线短信自助配置参数*/
	MINOR_LOCAL_CFG_WIRELESS_SMSSElFHELP     = 0x371 /*本地配置无线短信自助配置参数*/
	MINOR_REMOTE_CFG_WIRELESS_NETFLOWPARAM   = 0x372 /*远程配置无线流量配置参数*/
	MINOR_LOCAL_CFG_WIRELESS_NETFLOWPARAM    = 0x373 /*本地配置无线流量配置参数*/

	//0x400-0x1000 门禁操作类型
	MINOR_REMOTE_OPEN_DOOR                    = 0x400 //远程开门
	MINOR_REMOTE_CLOSE_DOOR                   = 0x401 //远程关门(受控)
	MINOR_REMOTE_ALWAYS_OPEN                  = 0x402 //远程常开(自由)
	MINOR_REMOTE_ALWAYS_CLOSE                 = 0x403 //远程常关(禁用)
	MINOR_REMOTE_CHECK_TIME                   = 0x404 //远程手动校时
	MINOR_NTP_CHECK_TIME                      = 0x405 //NTP自动校时
	MINOR_REMOTE_CLEAR_CARD                   = 0x406 //远程清空卡号
	MINOR_REMOTE_RESTORE_CFG                  = 0x407 //远程恢复默认参数
	MINOR_ALARMIN_ARM                         = 0x408 //防区布防
	MINOR_ALARMIN_DISARM                      = 0x409 //防区撤防
	MINOR_LOCAL_RESTORE_CFG                   = 0x40a //本地恢复默认参数
	MINOR_REMOTE_CAPTURE_PIC                  = 0x40b //远程抓拍
	MINOR_MOD_NET_REPORT_CFG                  = 0x40c //修改网络中心参数配置
	MINOR_MOD_GPRS_REPORT_PARAM               = 0x40d //修改GPRS中心参数配置
	MINOR_MOD_REPORT_GROUP_PARAM              = 0x40e //修改中心组参数配置
	MINOR_UNLOCK_PASSWORD_OPEN_DOOR           = 0x40f //解除码输入
	MINOR_AUTO_RENUMBER                       = 0x410 //自动重新编号
	MINOR_AUTO_COMPLEMENT_NUMBER              = 0x411 //自动补充编号
	MINOR_NORMAL_CFGFILE_INPUT                = 0x412 //导入普通配置文件
	MINOR_NORMAL_CFGFILE_OUTTPUT              = 0x413 //导出普通配置文件
	MINOR_CARD_RIGHT_INPUT                    = 0x414 //导入卡权限参数
	MINOR_CARD_RIGHT_OUTTPUT                  = 0x415 //导出卡权限参数
	MINOR_LOCAL_USB_UPGRADE                   = 0x416 //本地U盘升级
	MINOR_REMOTE_VISITOR_CALL_LADDER          = 0x417 //访客呼梯
	MINOR_REMOTE_HOUSEHOLD_CALL_LADDER        = 0x418 //住户呼梯
	MINOR_REMOTE_ACTUAL_GUARD                 = 0x419 //远程实时布防
	MINOR_REMOTE_ACTUAL_UNGUARD               = 0x41a //远程实时撤防
	MINOR_REMOTE_CONTROL_NOT_CODE_OPER_FAILED = 0x41b //遥控器未对码操作失败
	MINOR_REMOTE_CONTROL_CLOSE_DOOR           = 0x41c //遥控器关门
	MINOR_REMOTE_CONTROL_OPEN_DOOR            = 0x41d //遥控器开门
	MINOR_REMOTE_CONTROL_ALWAYS_OPEN_DOOR     = 0x41e //遥控器常开门
	MINOR_M1_CARD_ENCRYPT_VERIFY_OPEN         = 0x41f /*M1卡加密验证功能开启*/
	MINOR_M1_CARD_ENCRYPT_VERIFY_CLOSE        = 0x420 /*M1卡加密验证功能关闭*/
	MINOR_NFC_FUNCTION_OPEN                   = 0x421 /*NFC开门功能开启*/
	MINOR_NFC_FUNCTION_CLOSE                  = 0x422 /*NFC开门功能关闭*/
	MINOR_OFFLINE_DATA_OUTPUT                 = 0x423 //离线采集数据导出
	MINOR_CREATE_SSH_LINK                     = 0x42d //建立SSH连接
	MINOR_CLOSE_SSH_LINK                      = 0x42e //断开SSH连接

	MINOR_OPERATION_CUSTOM1  = 0x900 //门禁自定义操作1
	MINOR_OPERATION_CUSTOM2  = 0x901 //门禁自定义操作2
	MINOR_OPERATION_CUSTOM3  = 0x902 //门禁自定义操作3
	MINOR_OPERATION_CUSTOM4  = 0x903 //门禁自定义操作4
	MINOR_OPERATION_CUSTOM5  = 0x904 //门禁自定义操作5
	MINOR_OPERATION_CUSTOM6  = 0x905 //门禁自定义操作6
	MINOR_OPERATION_CUSTOM7  = 0x906 //门禁自定义操作7
	MINOR_OPERATION_CUSTOM8  = 0x907 //门禁自定义操作8
	MINOR_OPERATION_CUSTOM9  = 0x908 //门禁自定义操作9
	MINOR_OPERATION_CUSTOM10 = 0x909 //门禁自定义操作10
	MINOR_OPERATION_CUSTOM11 = 0x90a //门禁自定义操作11
	MINOR_OPERATION_CUSTOM12 = 0x90b //门禁自定义操作12
	MINOR_OPERATION_CUSTOM13 = 0x90c //门禁自定义操作13
	MINOR_OPERATION_CUSTOM14 = 0x90d //门禁自定义操作14
	MINOR_OPERATION_CUSTOM15 = 0x90e //门禁自定义操作15
	MINOR_OPERATION_CUSTOM16 = 0x90f //门禁自定义操作16
	MINOR_OPERATION_CUSTOM17 = 0x910 //门禁自定义操作17
	MINOR_OPERATION_CUSTOM18 = 0x911 //门禁自定义操作18
	MINOR_OPERATION_CUSTOM19 = 0x912 //门禁自定义操作19
	MINOR_OPERATION_CUSTOM20 = 0x913 //门禁自定义操作20
	MINOR_OPERATION_CUSTOM21 = 0x914 //门禁自定义操作21
	MINOR_OPERATION_CUSTOM22 = 0x915 //门禁自定义操作22
	MINOR_OPERATION_CUSTOM23 = 0x916 //门禁自定义操作23
	MINOR_OPERATION_CUSTOM24 = 0x917 //门禁自定义操作24
	MINOR_OPERATION_CUSTOM25 = 0x918 //门禁自定义操作25
	MINOR_OPERATION_CUSTOM26 = 0x919 //门禁自定义操作26
	MINOR_OPERATION_CUSTOM27 = 0x91a //门禁自定义操作27
	MINOR_OPERATION_CUSTOM28 = 0x91b //门禁自定义操作28
	MINOR_OPERATION_CUSTOM29 = 0x91c //门禁自定义操作29
	MINOR_OPERATION_CUSTOM30 = 0x91d //门禁自定义操作30
	MINOR_OPERATION_CUSTOM31 = 0x91e //门禁自定义操作31
	MINOR_OPERATION_CUSTOM32 = 0x91f //门禁自定义操作32
	MINOR_OPERATION_CUSTOM33 = 0x920 //门禁自定义操作33
	MINOR_OPERATION_CUSTOM34 = 0x921 //门禁自定义操作34
	MINOR_OPERATION_CUSTOM35 = 0x922 //门禁自定义操作35
	MINOR_OPERATION_CUSTOM36 = 0x923 //门禁自定义操作36
	MINOR_OPERATION_CUSTOM37 = 0x924 //门禁自定义操作37
	MINOR_OPERATION_CUSTOM38 = 0x925 //门禁自定义操作38
	MINOR_OPERATION_CUSTOM39 = 0x926 //门禁自定义操作39
	MINOR_OPERATION_CUSTOM40 = 0x927 //门禁自定义操作40
	MINOR_OPERATION_CUSTOM41 = 0x928 //门禁自定义操作41
	MINOR_OPERATION_CUSTOM42 = 0x929 //门禁自定义操作42
	MINOR_OPERATION_CUSTOM43 = 0x92a //门禁自定义操作43
	MINOR_OPERATION_CUSTOM44 = 0x92b //门禁自定义操作44
	MINOR_OPERATION_CUSTOM45 = 0x92c //门禁自定义操作45
	MINOR_OPERATION_CUSTOM46 = 0x92d //门禁自定义操作46
	MINOR_OPERATION_CUSTOM47 = 0x92e //门禁自定义操作47
	MINOR_OPERATION_CUSTOM48 = 0x92f //门禁自定义操作48
	MINOR_OPERATION_CUSTOM49 = 0x930 //门禁自定义操作49
	MINOR_OPERATION_CUSTOM50 = 0x931 //门禁自定义操作50
	MINOR_OPERATION_CUSTOM51 = 0x932 //门禁自定义操作51
	MINOR_OPERATION_CUSTOM52 = 0x933 //门禁自定义操作52
	MINOR_OPERATION_CUSTOM53 = 0x934 //门禁自定义操作53
	MINOR_OPERATION_CUSTOM54 = 0x935 //门禁自定义操作54
	MINOR_OPERATION_CUSTOM55 = 0x936 //门禁自定义操作55
	MINOR_OPERATION_CUSTOM56 = 0x937 //门禁自定义操作56
	MINOR_OPERATION_CUSTOM57 = 0x938 //门禁自定义操作57
	MINOR_OPERATION_CUSTOM58 = 0x939 //门禁自定义操作58
	MINOR_OPERATION_CUSTOM59 = 0x93a //门禁自定义操作59
	MINOR_OPERATION_CUSTOM60 = 0x93b //门禁自定义操作60
	MINOR_OPERATION_CUSTOM61 = 0x93c //门禁自定义操作61
	MINOR_OPERATION_CUSTOM62 = 0x93d //门禁自定义操作62
	MINOR_OPERATION_CUSTOM63 = 0x93e //门禁自定义操作63
	MINOR_OPERATION_CUSTOM64 = 0x93f //门禁自定义操作64

	MINOR_SET_WIFI_PARAMETER              = 0x950 //设置WIFI配置参数
	MINOR_EZVIZ_LOGIN                     = 0x951 //萤石云登陆
	MINOR_EZVIZ_LOGINOUT                  = 0x952 //萤石云登出
	MINOR_LOCK_ADD                        = 0x953 //智能锁添加
	MINOR_LOCK_DELETE                     = 0x954 //智能锁删除
	MINOR_LOCK_GET_STATUS                 = 0x955 //智能锁状态获取
	MINOR_LOCK_SET_TMP_PASSWORD           = 0x956 //智能锁临时密码下发
	MINOR_LOCK_SET_SILENT_MODE            = 0x957 //智能锁静音设置
	MINOR_LOCK_SET_LATE_WARNING           = 0x958 //智能锁晚归提醒设置
	MINOR_LOCK_IPC_ADD                    = 0x959 //智能锁IPC关联
	MINOR_LOCK_IPC_REMOVE                 = 0x95a //智能锁IPC解除关联
	MINOR_LOCK_DETECTOR_ADD               = 0x95b //智能锁探测器关联
	MINOR_LOCK_DETECTOR_REMOVE            = 0x95c //智能锁探测器解除关联
	MINOR_LOCK_MESSAGE_REMINDING_OPEN     = 0x95d //智能锁消息提醒打开
	MINOR_LOCK_MESSAGE_REMINDING_CLOSE    = 0x95e //智能锁消息提醒关闭
	MINOR_LOCK_SET_HEART_BEAT             = 0x95f //智能锁心跳设置
	MINOR_LOCK_REBOOT                     = 0x960 //智能锁重启
	MINOR_LOCK_CLEAR_USER                 = 0x961 //智能锁清空用户
	MINOR_LOCK_FORMAT                     = 0x962 //智能锁格式化
	MINOR_LOCK_FINGER_CHANGE              = 0x963 //智能锁指纹改动
	MINOR_LOCK_PASSWORD_CHANGE            = 0x964 //智能锁密码改动
	MINOR_LOCK_CARD_CHANGE                = 0x965 //智能锁卡信息改动
	MINOR_LOCK_USER_CHANGE                = 0x966 //智能锁用户信息改动
	MINOR_LOCK_SYSTEM_CHANGE              = 0x967 //智能锁系统信息改动
	MINOR_LOCK_CHANGE_ADD_UESR            = 0x968 //智能锁新增用户
	MINOR_LOCK_CHANGE_DEL_UESR            = 0x969 //智能锁删除用户
	MINOR_LOCK_CHANGE_CUSTOM_USER_NAME    = 0x96a //智能锁自定义用户用户名改动
	MINOR_LOCK_CHANGE_REMOTE_DEVICE       = 0x96b //智能锁遥控器信息改动
	MINOR_LOCK_CHANGE_ADD_FP              = 0x96c //智能锁新增指纹
	MINOR_LOCK_CHANGE_DEL_FP              = 0x96d //智能锁删除指纹
	MINOR_LOCK_CHANGE_ADD_PASSWORD        = 0x96e //智能锁新增密码
	MINOR_LOCK_CHANGE_DEL_PASSWORD        = 0x96f //智能锁删除密码
	MINOR_LOCK_CHANGE_ADD_CARD            = 0x970 //智能锁新增卡片
	MINOR_LOCK_CHANGE_DEL_CARD            = 0x971 //智能锁删除卡片
	MINOR_LOCK_NETWORK_SWITCH             = 0x972 //智能锁网络功能开关改动
	MINOR_LOCK_CLEAR_NETWORK_DATA         = 0x973 //智能锁网络数据清空
	MINOR_LOCK_CLEAR_HOST_USER            = 0x974 //智能锁清空主人用户
	MINOR_LOCK_CLEAR_GUEST_USER           = 0x975 //智能锁清空客人用户
	MINOR_LOCK_CLEAN_ALL_REMOTE_DEVICE    = 0x976 //遥控器用户信息清空
	MINOR_LOCK_CLEAN_NORMAL_USER_FINGRT   = 0x977 //智能锁清空普通用户指纹
	MINOR_LOCK_CLEAN_ALL_CARD             = 0x978 //智能锁清空所有卡片
	MINOR_LOCK_CLEAN_ALL_PASSWORD         = 0x979 //智能锁清空所有密码
	MINOR_START_WIRELESSSERVER            = 0x97a //开启设备热点
	MINOR_STOP_WIRELESSSERVER             = 0x97b //关闭设备热点
	MINOR_EMERGENCY_CARD_AUTH_NORMAL_CARD = 0x97c //应急管理卡授权普通卡
	MINOR_CHANGE_ALWAYS_OPEN_RIGHT        = 0x97d //通道模式改动
	MINOR_LOCK_DOOR_BELL_EVENT            = 0x97e //门铃事件（操作锁触发）

	//传显信息发布操作日志
	MINOR_BACKUP_DATA                         = 0xc41 //数据备份
	MINOR_TRANSFER_DATA                       = 0xc42 //数据迁移
	MINOR_RESTORE_DATA                        = 0xc43 //数据还原
	MINOR_SET_INPUT_PLAN                      = 0xc44 //设置终端定时输入切换计划
	MINOR_TERMINAL_ADB                        = 0xc45 //终端ADB配置
	MINOR_TERMINAL_VOLUME                     = 0xc46 //终端音量配置
	MINOR_TERMINAL_LOGO                       = 0xc47 //终端LOGO配置
	MINOR_TERMINAL_DEFAULT_SCHEDULE           = 0xc48 //垫片日程使能
	MINOR_TERMINAL_PASSWORD                   = 0xc49 //设置终端密码
	MINOR_TERMINAL_IP                         = 0xc4a //终端IP配置
	MINOR_TERMINAL_RELATE_IPC                 = 0xc4b //终端关联IPC
	MINOR_TERMINAL_SERVER                     = 0xc4c //终端关联服务器配置
	MINOR_TERMINAL_SADP                       = 0xc4d //终端SADP开关配置
	MINOR_TERMINAL_TIMEZONE                   = 0xc4e //终端时区配置
	MINOR_TERMINAL_TEMP_PROTECT               = 0xc4f //终端温度保护配置
	MINOR_ADD_ORGANIZATION                    = 0xc50 //添加组织
	MINOR_DELETE_ORGANIZATION                 = 0xc51 //删除组织
	MINOR_MODIFY_ORGANIZATION                 = 0xc52 //修改组织
	MINOR_WEATHER_FACTORY                     = 0xc53 //天气厂商配置
	MINOR_SADP_ENABLE                         = 0xc54 //sadp开关配置
	MINOR_SSH_ENABLE                          = 0xc55 //SSH开关配置
	MINOR_MODIFY_MATERIAL                     = 0xc56 //素材参数修改
	MINOR_INSERT_CHARACTER                    = 0xc57 //插播文字消息
	MINOR_TERMINAL_BACKLIGHT                  = 0xc58 //终端背光配置
	MINOR_DOWNLOAD_MATERIAL_THUMBNAIL         = 0xc59 //下载素材缩略图
	MINOR_UPLOAD_PROGRAM_THUMBNAIL            = 0xc5a //上传节目缩略图
	MINOR_TDOWNLOAD_PROGRAM_THUMBNAIL         = 0xc5b //下载节目缩略图
	MINOR_BATCH_DELETE_SCHEDULE_PLAN          = 0xc5c //批量删除发布计划
	MINOR_REPUBLISH                           = 0xc5d //重新发布
	MINOR_CLEAR_TERMINAL_PLAY_INFO            = 0xc5e //清空终端播放信息
	MINOR_GET_TERMINAL_RESOLUTION             = 0xc5f //获取终端分辨率
	MINOR_SET_TERMINAL_RESOLUTION             = 0xc60 //设置终端分辨率
	MINOR_GET_BATCH_TERMINAL_UPGRATE_PROGRESS = 0xc61 //批量获取终端升级进度
	MINOR_GET_BATCH_PROGRESS                  = 0xc62 //批量获取终端发布进度
	MINOR_GET_TEMPLATE                        = 0xc64 //获取模板
	MINOR_INIT_TEMPLATE                       = 0xc65 //初始化模板
	MINOR_GET_TERMINAL_NTP_SERVERS            = 0xc66 //获取终端NTP服务
	MINOR_SET_TERMINAL_NTP_SERVERS            = 0xc67 //设置终端NTP服务
	MINOR_GET_RELEASE_DETAILS                 = 0xc68 //获取发布详细信息
	MINOR_UPLOAD_TEMPLATE_THUMBNAIL           = 0xc69 //上传模板缩略图
	MINOR_DOWNLOAD_TEMPLATE_THUMBNAIL         = 0xc6a //下载模板缩略图
	MINOR_ADD_TEMPLATE                        = 0xc6b //添加模板
	MINOR_DELETE_TEMPLATE                     = 0xc6c //删除模板
	MINOR_MODIFY_TEMPLATE                     = 0xc6d //修改模板
	MINOR_ADD_SCHEDULE_PLAN                   = 0xc6e //添加发布计划
	MINOR_MODIFY_SCHEDULE_PLAN                = 0xc6f //修改发布计划
	MINOR_CANCEL_SCHEDULE_RELEASE             = 0xc70 //取消日程发布
	MINOR_GET_SCHEDULE                        = 0xc71 //获取日程
	MINOR_ADD_INSERT                          = 0xc72 //新建插播
	MINOR_CANCEL_INSERT                       = 0xc73 //取消插播
	MINOR_SWITCH_LANGUAGE                     = 0xc74 //切换语言
	MINOR_SET_ADMIN_INITIAL_PASSWORD          = 0xc75 //设置admin初始密码
	MINOR_MODIFY_PORT                         = 0xc76 //修改端口
	MINOR_MODIFY_STORAGE_PATH                 = 0xc77 //修改存储路径
	MINOR_EXIT_PROGRAM                        = 0xc78 //退出程序
	MINOR_MODULE_STARTUP_SUCCESS              = 0xc79 //模块启动成功
	MINOR_APPROVE_SCHEDULE                    = 0xc80 //日程审核
	MINOR_GENERAL_DATA_SEND                   = 0xc81 //第三方数据下发
	MINOR_SET_SIGN_INTERFACE                  = 0xc82 //配置签到界面参数
	MINOR_GET_SIGN_INTERFACE                  = 0xc83 //获取签到界面参数
	MINOR_SET_SHOW_MODE                       = 0xc84 //配置显示模式参数
	MINOR_GET_SHOW_MODE                       = 0xc85 //获取显示模式参数
	MINOR_SET_SCREEN_DIRECTION                = 0xc86 //配置屏幕方向参数
	MINOR_GET_SCREEN_DIRECTION                = 0xc87 //获取屏幕方向参数
	MINOR_SET_LOCK_SCREEN                     = 0xc88 //配置锁屏参数
	MINOR_GET_LOCK_SCREEN                     = 0xc89 //获取锁屏参数
	MINOR_SET_FACE_DATA_LIB                   = 0xc8a //配置人脸库参数
	MINOR_DELETE_FACE_DATA_LIB                = 0xc8b //删除人脸库
	MINOR_SET_SPEC_FACE_DATA_LIB              = 0xc8c //配置指定人脸库参数
	MINOR_DELETE_SPEC_FACE_DATA_LIB           = 0xc8d //删除指定人脸库参数
	MINOR_ADD_FACE_DATA                       = 0xc8e //添加人脸数据
	MINOR_SEARCH_FACE_DATA                    = 0xc8f //查询人脸数据
	MINOR_MODIFY_FACE_DATA                    = 0xc90 //修改人脸数据
	MINOR_DELETE_FACE_DATA                    = 0xc91 //删除人脸数据
	MINOR_DELETE_USERINFO_DETAIL              = 0xc92 //人员信息及权限删除
	MINOR_ADD_USERINFO                        = 0xc93 //添加人员信息
	MINOR_MODIFY_USERINFO                     = 0xc94 //修改人员信息
	MINOR_DELETE_USERINFO                     = 0xc95 //删除人员信息
	MINOR_ADD_CARD_INFO                       = 0xc96 //添加卡信息
	MINOR_MODIFY_CARD_INFO                    = 0xc97 //修改卡信息
	MINOR_DELETE_CARD_INFO                    = 0xc98 //删除卡信息
	MINOR_SET_USER_RIGHT_WEEK                 = 0xc99 //人员权限周计划设置
	MINOR_SET_USER_RIGHT_HOLIDAY              = 0xc9a //人员权限节日计划设置
	MINOR_SET_USER_RIGHT_HOLIDAYGROUP         = 0xc9b //人员权限假日组计划设置
	MINOR_SET_USER_RIGHT_TEMPLATE             = 0xc9c //人员权限计划模板设置

	//2012-03-05 ITC操作日志类型
	MINOR_SET_TRIGGERMODE_CFG       = 0x1001 /*设置触发模式参数*/
	MINOR_GET_TRIGGERMODE_CFG       = 0x1002 /*获取触发模式参数*/
	MINOR_SET_IOOUT_CFG             = 0x1003 /*设置IO输出参数*/
	MINOR_GET_IOOUT_CFG             = 0x1004 /*获取IO输出参数*/
	MINOR_GET_TRIGGERMODE_DEFAULT   = 0x1005 /*获取触发模式推荐参数*/
	MINOR_GET_ITCSTATUS             = 0x1006 /*获取状态检测参数*/
	MINOR_SET_STATUS_DETECT_CFG     = 0x1007 /*设置状态检测参数*/
	MINOR_GET_STATUS_DETECT_CFG     = 0x1008 /*获取状态检测参数*/
	MINOR_SET_VIDEO_TRIGGERMODE_CFG = 0x1009 /*设置视频触发模式参数*/
	MINOR_GET_VIDEO_TRIGGERMODE_CFG = 0x100a /*获取视频触发模式参数*/

	//2018-04-23 通用物联网关操作日志类型
	MINOR_ALARMHOST_GUARD         = 0x1010 //普通布防(外出布防)
	MINOR_ALARMHOST_UNGUARD       = 0x1011 //普通撤防
	MINOR_ALARMHOST_BYPASS        = 0x1012 //旁路
	MINOR_ALARMHOST_DURESS_ACCESS = 0x1013 //挟持

	MINOR_ALARMHOST_RS485_PARAM         = 0x1018 //修改485配置参数
	MINOR_ALARMHOST_ALARM_OUTPUT        = 0x1019 //控制触发器
	MINOR_ALARMHOST_ACCESS_OPEN         = 0x101a //控制门禁开
	MINOR_ALARMHOST_ACCESS_CLOSE        = 0x101b //控制门禁关
	MINOR_ALARMHOST_SIREN_OPEN          = 0x101c //控制警号开
	MINOR_ALARMHOST_SIREN_CLOSE         = 0x101d //控制警号关
	MINOR_ALARMHOST_MOD_ZONE_CONFIG     = 0x101e //修改防区参数
	MINOR_ALARMHOST_MOD_ALARMOUT_CONIFG = 0x101f //修改触发器参数
	MINOR_ALARMHOST_MOD_ANALOG_CONFIG   = 0x1020 //修改模拟量配置
	MINOR_ALARMHOST_RS485_CONFIG        = 0x1021 //修改485通道配置
	MINOR_ALARMHOST_PHONE_CONFIG        = 0x1022 //修改拨号配置
	MINOR_ALARMHOST_ADD_ADMIN           = 0x1023 //增加管理员
	MINOR_ALARMHOST_MOD_ADMIN_PARAM     = 0x1024 //修改管理员参数
	MINOR_ALARMHOST_DEL_ADMIN           = 0x1025 //删除管理员
	MINOR_ALARMHOST_ADD_NETUSER         = 0x1026 //增加后端操作员
	MINOR_ALARMHOST_MOD_NETUSER_PARAM   = 0x1027 //修改后端操作员参数
	MINOR_ALARMHOST_DEL_NETUSER         = 0x1028 //删除后端操作员
	MINOR_ALARMHOST_ADD_OPERATORUSER    = 0x1029 //增加前端操作员
	MINOR_ALARMHOST_MOD_OPERATORUSER_PW = 0x102a //修改前端操作员密码
	MINOR_ALARMHOST_DEL_OPERATORUSER    = 0x102b //删除前端操作员
	MINOR_ALARMHOST_ADD_KEYPADUSER      = 0x102c //增加键盘/读卡器用户
	MINOR_ALARMHOST_DEL_KEYPADUSER      = 0x102d //删除键盘/读卡器用户

	MINOR_ALARMHOST_MOD_HOST_CONFIG = 0x1032 //修改主机配置
	MINOR_ALARMHOST_RESTORE_BYPASS  = 0x1033 //旁路恢复

	MINOR_ALARMHOST_ALARMOUT_OPEN        = 0x1034 //触发器开启
	MINOR_ALARMHOST_ALARMOUT_CLOSE       = 0x1035 //触发器关闭
	MINOR_ALARMHOST_MOD_SUBSYSTEM_PARAM  = 0x1036 //修改子系统参数配置
	MINOR_ALARMHOST_GROUP_BYPASS         = 0x1037 //组旁路
	MINOR_ALARMHOST_RESTORE_GROUP_BYPASS = 0x1038 //组旁路恢复
	MINOR_ALARMHOST_MOD_GRPS_PARAM       = 0x1039 //修改GPRS参数

	MINOR_ALARMHOST_MOD_REPORT_MOD    = 0x103b //修改上传方式配置
	MINOR_ALARMHOST_MOD_GATEWAY_PARAM = 0x103c //修改门禁参数配置

	MINOR_STAY_ARM                       = 0x104c //留守布防
	MINOR_QUICK_ARM                      = 0x104d //即时布防
	MINOR_AUTOMATIC_ARM                  = 0x104e //自动布防
	MINOR_AUTOMATIC_DISARM               = 0x104f //自动撤防
	MINOR_KEYSWITCH_ARM                  = 0x1050 //钥匙布撤防防区布防
	MINOR_KEYSWITCH_DISARM               = 0x1051 //钥匙布撤防防区撤防
	MINOR_CLEAR_ALARM                    = 0x1052 //消警
	MINOR_MOD_FAULT_CFG                  = 0x1053 //修改系统故障配置
	MINOR_MOD_EVENT_TRIGGER_ALARMOUT_CFG = 0x1054 //修改事件触发触发器配置
	MINOR_SEARCH_EXTERNAL_MODULE         = 0x1055 //搜索外接模块
	MINOR_REGISTER_EXTERNAL_MODULE       = 0x1056 //重新注册外接模块
	MINOR_CLOSE_KEYBOARD_ALARM           = 0x1057 //关闭键盘报警提示音
	MINOR_MOD_3G_PARAM                   = 0x1058 //修改3G参数
	MINOR_MOD_PRINT_PARAM                = 0x1059 //修改打印机参数
	MINOR_ALARMHOST_SD_CARD_FORMAT       = 0x1060 //SD卡格式化
	MINOR_ALARMHOST_SUBSYSTEM_UPGRADE    = 0x1061 //子板固件升级

	MINOR_PLAN_ARM_CFG      = 0x1062 //计划布撤防参数配置
	MINOR_PHONE_ARM         = 0x1063 //手机布防
	MINOR_PHONE_STAY_ARM    = 0x1064 //手机留守布防
	MINOR_PHONE_QUICK_ARM   = 0x1065 //手机即时布防
	MINOR_PHONE_DISARM      = 0x1066 //手机撤防
	MINOR_PHONE_CLEAR_ALARM = 0x1067 //手机消警
	MINOR_WHITELIST_CFG     = 0x1068 //白名单配置
	MINOR_TIME_TRIGGER_CFG  = 0x1069 //定时开关触发器配置
	MINOR_CAPTRUE_CFG       = 0x106a //抓图参数配置
	MINOR_TAMPER_CFG        = 0x106b //防区防拆参数配置

	MINOR_REMOTE_KEYPAD_UPGRADE                = 0x106c //远程升级键盘
	MINOR_ONETOUCH_AWAY_ARMING                 = 0x106d //一键外出布防
	MINOR_ONETOUCH_STAY_ARMING                 = 0x106e //一键留守布防
	MINOR_SINGLE_PARTITION_ARMING_OR_DISARMING = 0x106f //单防区布撤防
	MINOR_CARD_CONFIGURATION                   = 0x1070 //卡参数配置
	MINOR_CARD_ARMING_OR_DISARMING             = 0x1071 //刷卡布撤防
	MINOR_EXPENDING_NETCENTER_CONFIGURATION    = 0x1072 //扩展网络中心配置
	MINOR_NETCARD_CONFIGURATION                = 0x1073 //网卡配置
	MINOR_DDNS_CONFIGURATION                   = 0x1074 //DDNS配置
	MINOR_RS485BUS_CONFIGURATION               = 0x1075 // 485总线参数配置
	MINOR_RS485BUS_RE_REGISTRATION             = 0x1076 //485总线重新注册

	MINOR_REMOTE_OPEN_ELECTRIC_LOCK  = 0x1077 //远程打开电锁
	MINOR_REMOTE_CLOSE_ELECTRIC_LOCK = 0x1078 //远程关闭电锁
	MINOR_LOCAL_OPEN_ELECTRIC_LOCK   = 0x1079 //本地打开电锁
	MINOR_LOCAL_CLOSE_ELECTRIC_LOCK  = 0x107a //本地关闭电锁
	MINOR_OPEN_ALARM_LAMP            = 0x107b //打开警灯(远程)
	MINOR_CLOSE_ALARM_LAMP           = 0x107c //关闭警灯(远程)

	MINOR_TEMPORARY_PASSWORD = 0x107d //临时密码操作记录

	MINOR_HIDDNS_CONFIG                = 0x1082 // HIDDNS配置
	MINOR_REMOTE_KEYBOARD_UPDATA       = 0x1083 //远程键盘升级日志
	MINOR_ZONE_ADD_DETECTOR            = 0x1084 //防区添加探测器
	MINOR_ZONE_DELETE_DETECTOR         = 0x1085 //防区删除探测器
	MINOR_QUERY_DETECTOR_SIGNAL        = 0x1086 //主机查询探测器信号强度
	MINOR_QUERY_DETECTOR_BATTERY       = 0x1087 //主机查询探测器电量
	MINOR_SET_DETECTOR_GUARD           = 0x1088 //探测器布防
	MINOR_SET_DETECTOR_UNGUARD         = 0x1089 //探测器撤防
	MINOR_WIRELESS_CONFIGURATION       = 0x108a //无线参数配置
	MINOR_OPEN_VOICE                   = 0x108b //打开语音
	MINOR_CLOSE_VOICE                  = 0x108c //关闭语音
	MINOR_ENABLE_FUNCTION_KEY          = 0x108d //启用功能键
	MINOR_DISABLE_FUNCTION_KEY         = 0x108e //关闭功能键
	MINOR_READ_CARD                    = 0x108f //巡更刷卡
	MINOR_START_BROADCAST              = 0x1090 //打开语音广播
	MINOR_STOP_BROADCAST               = 0x1091 //关闭语音广播
	MINOR_REMOTE_ZONE_MODULE_UPGRADE   = 0x1092 //远程升级防区模块
	MINOR_NETWORK_MODULE_EXTEND        = 0x1093 //网络模块参数配置
	MINOR_ADD_CONTROLLER               = 0x1094 //添加遥控器用户
	MINOR_DELETE_CONTORLLER            = 0x1095 //删除遥控器用户
	MINOR_REMOTE_NETWORKMODULE_UPGRADE = 0x1096 //远程升级网络模块
	MINOR_WIRELESS_OUTPUT_ADD          = 0x1097 //注册无线输出模块
	MINOR_WIRELESS_OUTPUT_DEL          = 0x1098 //删除无线输出模块
	MINOR_WIRELESS_REPEATER_ADD        = 0x1099 //注册无线中继器
	MINOR_WIRELESS_REPEATER_DEL        = 0x109a //删除无线中继器
	MINOR_PHONELIST_CFG                = 0x109b //电话名单参数配置
	MINOR_RF_SIGNAL_CHECK              = 0x109c // RF信号查询
	MINOR_USB_UPGRADE                  = 0x109d // USB升级
	MINOR_DOOR_TIME_REMINDER_CFG       = 0x109f //门磁定时提醒参数配置
	MINOR_WIRELESS_SIREN_ADD           = 0x1100 //注册无线警号
	MINOR_WIRELESS_SIREN_DEL           = 0x1101 //删除无线警号
	MINOR_OUT_SCALE_OPEN               = 0x1102 //辅电开启
	MINOR_OUT_SCALE_CLOSE              = 0x1103 //辅电关闭

	MINOR_ALARMHOST_4G_MODULS_START = 0x1108 //4G模块启用
	MINOR_ALARMHOST_4G_MODULS_STOP  = 0x1109 //4G模块停用

	MINOR_EZVIZ_CLOUD_START = 0x1110 //萤石云启用
	MINOR_EZVIZ_CLOUD_STOP  = 0x1111 //萤石云停用
	MINOR_SIPUA_GRID_START  = 0x1112 //国网B启用
	MINOR_SIPUA_GRID_STOP   = 0x1113 //国网B停用

	MINOR_MODBUS_FILE_DOWNLOAD = 0x1114 //导出modbus协议配置文件
	MINOR_MODBUS_FILE_UPLOAD   = 0x1115 //导入modbus协议配置文件

	MINOR_RS485_DLL_FILE_DOWNLOAD = 0x1116 //导出485协议库文件
	MINOR_RS485_DLL_FLIE_UPLOAD   = 0x1117 //导入485协议库文件
	MINOR_TX1_REBOOT              = 0x1118 //TX1系统正常重启

	MINOR_LORA_PARAM               = 0x1119 //LoRa参数
	MINOR_GB28181_PLATE_CFG_PARAM  = 0x111a //国标平台接入参数配置
	MINOR_GB28181_SERVER_START     = 0x111b //国标服务启用
	MINOR_GB28181_SERVER_STOP      = 0x111c //国标服务停用
	MINOR_WEB_AUTHENTICATION       = 0x111d //web认证方式配置
	MINOR_SADP_ENABLED             = 0x111e //SADP开关配置
	MINOR_HTTPS_ENABLED            = 0x111f //HTTPS开关配置
	MINOR_EZVIZ_PARAM_CFG          = 0x1120 //萤石云配置
	MINOR_SET_MOTION_DETECTION_CFG = 0x1121 //设置移动侦测参数配置
	MINOR_GET_MOTION_DETECTION_CFG = 0x1122 //获取移动侦测参数配置
	MINOR_SET_SHELTER_ALARM_CFG    = 0x1123 //设置遮挡报警参数配置
	MINOR_GET_SHELTER_ALARM_CFG    = 0x1124 //获取遮挡报警参数配置
	MINOR_SET_VIDEO_LOSS_CFG       = 0x1125 //设置视频丢失参数配置
	MINOR_GET_VIDEO_LOSS_CFG       = 0x1126 //获取视频丢失参数配置
	MINOR_SET_ABNORMAL_CFG         = 0x1127 //设置异常参数配置
	MINOR_GET_ABNORMAL_CFG         = 0x1128 //获取异常参数配置
	MINOR_SET_ALARM_LINKAGE_CFG    = 0x1129 //设置报警联动配置
	MINOR_GET_ALARM_LINKAGE_CFG    = 0x112a //获取报警联动配置
	MINOR_SET_NETWORK_CFG          = 0x112b //设置网络参数配置
	MINOR_GET_NETWORK_CFG          = 0x112c //获取网络参数配置
	MINOR_SET_VIDEO_MASK_CFG       = 0x112d //设置视频遮盖参数配置
	MINOR_GET_VIDEO_MASK_CFG       = 0x112e //获取视频遮盖参数配置

	MINOR_BASIC_OPERATION_CFG       = 0x112f //基本操作
	MINOR_DISPLAY_EFFECT_ADJUST_CFG = 0x1130 //显示效果调节
	MINOR_DISPLAY_PROPERTY_CFG      = 0x1131 //显示屏属性配置
	MINOR_SIGNAL_CABLE_CFG          = 0x1132 //信号线缆配置
	MINOR_BASIC_CFG                 = 0x1133 //基础配置
	MINOR_IMAGE_ADJUST_CFG          = 0x1134 //图像调整配置
	MINOR_IMAGE_ENHANCE_CFG         = 0x1135 //图像增强配置
	MINOR_NOSIGN_SCREEN_SAVER_CFG   = 0x1136 //无信号屏保
	MINOR_ADVANCED_OPERATION_CFG    = 0x1137 //高级操作
	MINOR_RECEIVE_CARD_CFG          = 0x1138 //接收卡配置
	MINOR_INPUT_SUPPORT_CFG         = 0x1139 //输入支持管理
	MINOR_SYSTEM_MAINTAIN_CFG       = 0x113a //系统维护配置
	MINOR_SYSTEM_TEST_CFG           = 0x113b //系统检测配置
	MINOR_JOINT_CFG                 = 0x113c //拼接配置
	MINOR_SHOW_MODE_CFG             = 0x113d //显示模式配置
	MINOR_ADVANCED_IMAGE_CFG        = 0x113e //高级图像配置

	//2013-04-19 ITS操作日志类型
	MINOR_LOCAL_ADD_CAR_INFO              = 0x2001 /*本地添加车辆信息*/
	MINOR_LOCAL_MOD_CAR_INFO              = 0x2002 /*本地修改车辆信息*/
	MINOR_LOCAL_DEL_CAR_INFO              = 0x2003 /*本地删除车辆信息*/
	MINOR_LOCAL_FIND_CAR_INFO             = 0x2004 /*本地查找车辆信息*/
	MINOR_LOCAL_ADD_MONITOR_INFO          = 0x2005 /*本地添加布控信息*/
	MINOR_LOCAL_MOD_MONITOR_INFO          = 0x2006 /*本地修改布控信息*/
	MINOR_LOCAL_DEL_MONITOR_INFO          = 0x2007 /*本地删除布控信息*/
	MINOR_LOCAL_FIND_MONITOR_INFO         = 0x2008 /*本地查询布控信息*/
	MINOR_LOCAL_FIND_NORMAL_PASS_INFO     = 0x2009 /*本地查询正常通行信息*/
	MINOR_LOCAL_FIND_ABNORMAL_PASS_INFO   = 0x200a /*本地查询异常通行信息*/
	MINOR_LOCAL_FIND_PEDESTRIAN_PASS_INFO = 0x200b /*本地查询正常通行信息*/
	MINOR_LOCAL_PIC_PREVIEW               = 0x200c /*本地图片预览*/
	MINOR_LOCAL_SET_GATE_PARM_CFG         = 0x200d /*设置本地配置出入口参数*/
	MINOR_LOCAL_GET_GATE_PARM_CFG         = 0x200e /*获取本地配置出入口参数*/
	MINOR_LOCAL_SET_DATAUPLOAD_PARM_CFG   = 0x200f /*设置本地配置数据上传参数*/
	MINOR_LOCAL_GET_DATAUPLOAD_PARM_CFG   = 0x2010 /*获取本地配置数据上传参数*/

	//2013-11-19新增日志类型
	MINOR_LOCAL_DEVICE_CONTROL                       = 0x2011 /*本地设备控制(本地开关闸)*/
	MINOR_LOCAL_ADD_EXTERNAL_DEVICE_INFO             = 0x2012 /*本地添加外接设备信息 */
	MINOR_LOCAL_MOD_EXTERNAL_DEVICE_INFO             = 0x2013 /*本地修改外接设备信息 */
	MINOR_LOCAL_DEL_EXTERNAL_DEVICE_INFO             = 0x2014 /*本地删除外接设备信息 */
	MINOR_LOCAL_FIND_EXTERNAL_DEVICE_INFO            = 0x2015 /*本地查询外接设备信息 */
	MINOR_LOCAL_ADD_CHARGE_RULE                      = 0x2016 /*本地添加收费规则 */
	MINOR_LOCAL_MOD_CHARGE_RULE                      = 0x2017 /*本地修改收费规则 */
	MINOR_LOCAL_DEL_CHARGE_RULE                      = 0x2018 /*本地删除收费规则 */
	MINOR_LOCAL_FIND_CHARGE_RULE                     = 0x2019 /*本地查询收费规则 */
	MINOR_LOCAL_COUNT_NORMAL_CURRENTINFO             = 0x2020 /*本地统计正常通行信息 */
	MINOR_LOCAL_EXPORT_NORMAL_CURRENTINFO_REPORT     = 0x2021 /*本地导出正常通行信息统计报表 */
	MINOR_LOCAL_COUNT_ABNORMAL_CURRENTINFO           = 0x2022 /*本地统计异常通行信息 */
	MINOR_LOCAL_EXPORT_ABNORMAL_CURRENTINFO_REPORT   = 0x2023 /*本地导出异常通行信息统计报表 */
	MINOR_LOCAL_COUNT_PEDESTRIAN_CURRENTINFO         = 0x2024 /*本地统计行人通行信息 */
	MINOR_LOCAL_EXPORT_PEDESTRIAN_CURRENTINFO_REPORT = 0x2025 /*本地导出行人通行信息统计报表 */
	MINOR_LOCAL_FIND_CAR_CHARGEINFO                  = 0x2026 /*本地查询过车收费信息 */
	MINOR_LOCAL_COUNT_CAR_CHARGEINFO                 = 0x2027 /*本地统计过车收费信息 */
	MINOR_LOCAL_EXPORT_CAR_CHARGEINFO_REPORT         = 0x2028 /*本地导出过车收费信息统计报表 */
	MINOR_LOCAL_FIND_SHIFTINFO                       = 0x2029 /*本地查询交接班信息 */
	MINOR_LOCAL_FIND_CARDINFO                        = 0x2030 /*本地查询卡片信息 */
	MINOR_LOCAL_ADD_RELIEF_RULE                      = 0x2031 /*本地添加减免规则 */
	MINOR_LOCAL_MOD_RELIEF_RULE                      = 0x2032 /*本地修改减免规则 */
	MINOR_LOCAL_DEL_RELIEF_RULE                      = 0x2033 /*本地删除减免规则 */
	MINOR_LOCAL_FIND_RELIEF_RULE                     = 0x2034 /*本地查询减免规则 */
	MINOR_LOCAL_GET_ENDETCFG                         = 0x2035 /*本地获取出入口控制机离线检测配置 */
	MINOR_LOCAL_SET_ENDETCFG                         = 0x2036 /*本地设置出入口控制机离线检测配置*/
	MINOR_LOCAL_SET_ENDEV_ISSUEDDATA                 = 0x2037 /*本地设置出入口控制机下发卡片信息 */
	MINOR_LOCAL_DEL_ENDEV_ISSUEDDATA                 = 0x2038 /*本地清空出入口控制机下发卡片信息 */

	MINOR_REMOTE_DEVICE_CONTROL          = 0x2101 /*远程设备控制*/
	MINOR_REMOTE_SET_GATE_PARM_CFG       = 0x2102 /*设置远程配置出入口参数*/
	MINOR_REMOTE_GET_GATE_PARM_CFG       = 0x2103 /*获取远程配置出入口参数*/
	MINOR_REMOTE_SET_DATAUPLOAD_PARM_CFG = 0x2104 /*设置远程配置数据上传参数*/
	MINOR_REMOTE_GET_DATAUPLOAD_PARM_CFG = 0x2105 /*获取远程配置数据上传参数*/
	MINOR_REMOTE_GET_BASE_INFO           = 0x2106 /*远程获取终端基本信息*/
	MINOR_REMOTE_GET_OVERLAP_CFG         = 0x2107 /*远程获取字符叠加参数配置*/
	MINOR_REMOTE_SET_OVERLAP_CFG         = 0x2108 /*远程设置字符叠加参数配置*/
	MINOR_REMOTE_GET_ROAD_INFO           = 0x2109 /*远程获取路口信息*/
	MINOR_REMOTE_START_TRANSCHAN         = 0x210a /*远程建立同步数据服务器*/
	MINOR_REMOTE_GET_ECTWORKSTATE        = 0x210b /*远程获取出入口终端工作状态*/
	MINOR_REMOTE_GET_ECTCHANINFO         = 0x210c /*远程获取出入口终端通道状态*/

	//远程控制 2013-11-19
	MINOR_REMOTE_ADD_EXTERNAL_DEVICE_INFO = 0x210d /*远程添加外接设备信息 */
	MINOR_REMOTE_MOD_EXTERNAL_DEVICE_INFO = 0x210e /*远程修改外接设备信息 */
	MINOR_REMOTE_GET_ENDETCFG             = 0x210f /*远程获取出入口控制机离线检测配置 */
	MINOR_REMOTE_SET_ENDETCFG             = 0x2110 /*远程设置出入口控制机离线检测配置*/
	MINOR_REMOTE_ENDEV_ISSUEDDATA         = 0x2111 /*远程设置出入口控制机下发卡片信息 */
	MINOR_REMOTE_DEL_ENDEV_ISSUEDDATA     = 0x2112 /*远程清空出入口控制机下发卡片信息 */

	//ITS 0x2115~0x2120 停车场车位项目
	MINOR_REMOTE_ON_CTRL_LAMP  = 0x2115 /*开启远程控制车位指示灯*/
	MINOR_REMOTE_OFF_CTRL_LAMP = 0x2116 /*关闭远程控制车位指示灯*/
	//Netra3.1.0
	MINOR_SET_VOICE_LEVEL_PARAM    = 0x2117 /*设置音量大小 */
	MINOR_SET_VOICE_INTERCOM_PARAM = 0x2118 /*设置音量录音 */
	MINOR_SET_INTELLIGENT_PARAM    = 0x2119 /*智能配置*/
	MINOR_LOCAL_SET_RAID_SPEED     = 0x211a /*本地设置raid速度*/
	MINOR_REMOTE_SET_RAID_SPEED    = 0x211b /*远程设置raid速度*/
	//Nerta3.1.2
	MINOR_REMOTE_CREATE_STORAGE_POOL = 0x211c //远程添加存储池
	MINOR_REMOTE_DEL_STORAGE_POOL    = 0x211d //远程删除存储池

	MINOR_REMOTE_DEL_PIC                   = 0x2120 //远程删除图片数据
	MINOR_REMOTE_DEL_RECORD                = 0x2121 //远程删除录像数据
	MINOR_REMOTE_CLOUD_ENABLE              = 0x2123 //远程设置云系统启用
	MINOR_REMOTE_CLOUD_DISABLE             = 0x2124 //远程设置云系统禁用
	MINOR_REMOTE_CLOUD_MODIFY_PARAM        = 0x2125 //远程修改存储池参数
	MINOR_REMOTE_CLOUD_MODIFY_VOLUME       = 0x2126 //远程修改存储池容量
	MINOR_REMOTE_GET_GB28181_SERVICE_PARAM = 0x2127 //远程获取GB28181服务参数
	MINOR_REMOTE_SET_GB28181_SERVICE_PARAM = 0x2128 //远程设置GB28181服务参数
	MINOR_LOCAL_GET_GB28181_SERVICE_PARAM  = 0x2129 //本地获取GB28181服务参数
	MINOR_LOCAL_SET_GB28181_SERVICE_PARAM  = 0x212a //本地配置B28181服务参数
	MINOR_REMOTE_SET_SIP_SERVER            = 0x212b //远程配置SIP SERVER
	MINOR_LOCAL_SET_SIP_SERVER             = 0x212c //本地配置SIP SERVER
	MINOR_LOCAL_BLACKWHITEFILE_OUTPUT      = 0x212d //本地黑白名单导出
	MINOR_LOCAL_BLACKWHITEFILE_INPUT       = 0x212e //本地黑白名单导入
	MINOR_REMOTE_BALCKWHITECFGFILE_OUTPUT  = 0x212f //远程黑白名单导出
	MINOR_REMOTE_BALCKWHITECFGFILE_INPUT   = 0x2130 //远程黑白名单导入

	MINOR_REMOTE_CREATE_MOD_VIEWLIB_SPACE = 0x2200 /*远程创建/修改视图库空间*/
	MINOR_REMOTE_DELETE_VIEWLIB_FILE      = 0x2201 /*远程删除视图库文件*/
	MINOR_REMOTE_DOWNLOAD_VIEWLIB_FILE    = 0x2202 /*远程下载视图库文件*/
	MINOR_REMOTE_UPLOAD_VIEWLIB_FILE      = 0x2203 /*远程上传视图库文件*/
	MINOR_LOCAL_CREATE_MOD_VIEWLIB_SPACE  = 0x2204 /*本地创建/修改视图库空间*/

	MINOR_LOCAL_SET_DEVICE_ACTIVE     = 0x3000 //本地激活设备
	MINOR_REMOTE_SET_DEVICE_ACTIVE    = 0x3001 //远程激活设备
	MINOR_LOCAL_PARA_FACTORY_DEFAULT  = 0x3002 //本地回复出厂设置
	MINOR_REMOTE_PARA_FACTORY_DEFAULT = 0x3003 //远程恢复出厂设置

	/*信息发布服务器操作日志*/
	MINOR_UPLAOD_STATIC_MATERIAL      = 0x2401 //静态素材上传
	MINOR_UPLOAD_DYNAMIC_MATERIAL     = 0x2402 //动态素材上传
	MINOR_DELETE_MATERIAL             = 0x2403 //删除素材
	MINOR_DOWNLOAD_STATIC_MATERIAL    = 0x2404 //静态素材下载
	MINOR_COVER_STATIC_MATERIAL       = 0x2405 //静态素材覆盖
	MINOR_APPROVE_MATERIAL            = 0x2406 //素材审核
	MINOR_UPLAOD_PROGRAM              = 0x2407 //上传节目
	MINOR_DOWNLOAD_PROGRAM            = 0x2408 //下载节目
	MINOR_DELETE_PROGRAM              = 0x2409 //删除节目
	MINOR_MODIFY_PROGRAM              = 0x240a //节目属性修改
	MINOR_APPROVE_PRAGRAM             = 0x240b //节目审核
	MINOR_UPLAOD_SCHEDULE             = 0x240c //上传日程
	MINOR_DOWNLOAD_SCHEDULE           = 0x240d //下载日程
	MINOR_DELETE_SCHEDULE             = 0x240e //删除日程
	MINOR_MODIFY_SCHEDULE             = 0x240f //修改日程属性
	MINOR_RELEASE_SCHEDULE            = 0x2410 //发布日程
	MINOR_ADD_TERMINAL                = 0x2411 //添加终端
	MINOR_DELETE_TERMINAL             = 0x2412 //删除终端
	MINOR_MODIFY_TERMIANL_PARAM       = 0x2413 //修改终端参数
	MINOR_MODIFY_TERMIANL_PLAY_PARAM  = 0x2414 //配置终端播放参数
	MINOR_ADD_TERMIANL_GROUP          = 0x2415 //添加终端组
	MINOR_MODIFY_TERMINAL_GROUP_PARAM = 0x2416 //修改终端组参数
	MINOR_DELETE_TERMIANL_GROUP       = 0x2417 //删除终端组
	MINOR_TERMINAL_PLAY_CONTROL       = 0x2418 //终端播放控制
	MINOR_TERMINAL_ON_OFF_LINE        = 0x2419 //终端上下线
	MINOR_SET_SWITCH_PLAN             = 0x241a //设置终端定时开关机计划
	MINOR_SET_VOLUME_PLAN             = 0x241b //设置终端定时音量计划
	MINOR_TERMINAL_SCREENSHOT         = 0x241c //终端截屏
	MINOR_SYSTEM_TIME_CFG             = 0x241d //系统校时
	MINOR_ADD_USER_CFG                = 0x241e //添加用户配置
	MINOR_DEL_USER_CFG                = 0x241f //删除用户配置
	MINOR_REMOTE_MANAGE_HDD           = 0x2420 //远程编辑硬盘
	MINOR_TERMINAL_UPDATE_START       = 0x2421 //终端升级
	MINOR_SVR_RESTORE_DEFAULT_PARAM   = 0x2422 //服务器远程恢复默认
	MINOR_SVR_REMOTE_RESTORE_FACTORY  = 0x2423 //服务器远程恢复出厂设置
	MINOR_SVR_REMOTE_REBOOT           = 0x2424 //服务器远程重启
	MINOR_SVR_MODIFY_NETWORK_PARAM    = 0x2425 //服务器网络参数修改
	MINOR_SVR_SOFTWARE_UPGRADE        = 0x2426 //服务器软件升级

	MINOR_REMOTE_CONFERENCE_CONFIG = 0x2501 //MCU会议配置
	MINOR_REMOTE_TERMINAL_CONFIG   = 0x2502 //MCU终端配置
	MINOR_REMOTE_GROUP_CONFIG      = 0x2503 //MCU分组配置
	MINOR_REMOTE_CONFERENCE_CTRL   = 0x2504 //MCU会议控制
	MINOR_REMOTE_TERMINAL_CTRL     = 0x2505 //MCU终端控制

	//NVR后端
	MINOR_LOCAL_RESET_LOGIN_PASSWORD  = 0x2600 /* 本地重置admin登陆密码*/
	MINOR_REMOTE_RESET_LOGIN_PASSWORD = 0x2601 /* 远程重置admin登录密码 */
	MINOR_LOCAL_FACE_BASE_CREATE      = 0x2602 /* 本地人脸对比库创建*/
	MINOR_REMOTE_FACE_BASE_CREATE     = 0x2603 /* 远程人脸对比库创建*/
	MINOR_LOCAL_FACE_BASE_MODIFY      = 0x2604 /* 本地人脸对比库修改*/
	MINOR_REMOTE_FACE_BASE_MODIFY     = 0x2605 /* 远程人脸对比库修改*/
	MINOR_LOCAL_FACE_BASE_DELETE      = 0x2606 /* 本地人脸对比库删除*/
	MINOR_REMOTE_FACE_BASE_DELETE     = 0x2607 /* 远程人脸对比库删除*/
	MINOR_LOCAL_FACE_DATA_APPEND      = 0x2608 /* 本地录入人脸数据*/
	MINOR_REMOTE_FACE_DATA_APPEND     = 0x2609 /* 远程录入人脸数据*/
	MINOR_LOCAL_FACE_DATA_SEARCH      = 0x2610 /* 本地人脸比对数据查找*/
	MINOR_REMOTE_FACE_DATA_SEARCH     = 0x2611 /* 远程人脸比对数据查找*/
	MINOR_LOCAL_FACE_DATA_ANALYSIS    = 0x2612 /* 本地图片分析操作*/
	MINOR_REMOTE_FACE_DATA_ANALYSIS   = 0x2613 /* 远程图片分析操作*/
	MINOR_LOCAL_FACE_DATA_EDIT        = 0x2614 /* 本地人脸数据修改*/
	MINOR_REMOTE_FACE_DATA_EDIT       = 0x2615 /* 远程人脸数据修改*/

	MINOR_LOCAL_FACE_DATA_DELETE = 0x2616 /* 本地人脸数据删除*/
	MINOR_REMOTE_FACE_DATA_DELET = 0x2617 /* 远程人脸数据删除*/

	MINOR_LOCAL_VCA_ANALYSIS_CFG  = 0x2618 /* 本地智能分析配置*/
	MINOR_REMOTE_VCA_ANALYSIS_CFG = 0x2619 /* 远程智能分析配置*/

	MINOR_LOCAL_FACE_BASE_IMPORT = 0x261a /* 本地导入人脸库*/
	MINOR_LOCAL_FACE_BASE_EXPORT = 0x261b /* 本地导出人脸库*/
	//NVR集群
	MINOR_REMOTE_CLUSTER_MODE_CONFIG    = 0x261c /* 远程集群模式配置操作*/
	MINOR_LOCAL_CLUSTER_MODE_CONFIG     = 0x261d /* 本地集群模式配置操作*/
	MINOR_REMOTE_CLUSTER_NETWORK_CONFIG = 0x261e /* 远程集群组网配置操作*/
	MINOR_LOCAL_CLUSTER_NETWORK_CONFIG  = 0x261f /* 本地集群组网配置操作*/
	MINOR_REMOTE_CLUSTER_ADD_DEVICE     = 0x2620 /* 远程集群添加设备操作*/
	MINOR_LOCAL_CLUSTER_ADD_DEVICE      = 0x2621 /* 本地集群添加设备操作*/
	MINOR_REMOTE_CLUSTER_DEL_DEVICE     = 0x2622 /* 远程集群删除设备操作*/
	MINOR_LOCAL_CLUSTER_DEL_DEVICE      = 0x2623 /* 本地集群删除设备操作*/
	MINOR_REMOTE_HFPD_CFG               = 0x2624 /* 远程高频人员检测配置*/
	MINOR_REMOTE_FACE_CONTRAST_TASK     = 0x2625 /* 远程人脸比对任务配置 */
	MINOR_REMOTE_LFPD_CFG               = 0x2626 /* 远程低频人员检测配置*/
	MINOR_REMOTE_IOTCFGFILE_INPUT       = 0x2627 //远程导入IOT配置文件
	MINOR_REMOTE_IOTCFGFILE_OUTPUT      = 0x2628 //远程导出IOT配置文件
	MINOR_LOCAL_IOT_ADD                 = 0x2629 //本地添加IOT通道
	MINOR_REMOTE_IOT_ADD                = 0x262a //远程添加IOT通道
	MINOR_LOCAL_IOT_DEL                 = 0x262b //本地删除IOT通道
	MINOR_REMOTE_IOT_DEL                = 0x262c //远程删除IOT通道
	MINOR_LOCAL_IOT_SET                 = 0x262d //本地配置IOT通道
	MINOR_REMOTE_IOT_SET                = 0x262e //远程配置IOT通道
	MINOR_LOCAL_IOTCFGFILE_INPUT        = 0x262f //本地导入IOT配置文件
	MINOR_LOCAL_IOTCFGFILE_OUTPUT       = 0x2630 //本地导出IOT配置文件
	MINOR_LOCAL_VAD_CFG                 = 0x2631 /* 本地语音活动检测配置*/
	MINOR_REMOTE_VAD_CFG                = 0x2632 /* 远程语音活动检测配置*/
	MINOR_LOCAL_ADDRESS_FILTER_CONFIG   = 0x2633 /* 本地地址过滤配置*/
	MINOR_REMOTE_ADDRESS_FILTER_CONFIG  = 0x2634 /* 远程地址过滤配置*/
	MINOR_LOCAL_POE_CFG                 = 0x2635 /* 本地POE配置*/
	MINOR_REMOTE_POE_CFG                = 0x2636 /* 远程POE配置*/
	MINOR_LOCAL_RESET_CHANNEL_PASSWORD  = 0x2637 /* 本地重置通道密码*/
	MINOR_REMOTE_RESET_CHANNEL_PASSWORD = 0x2638 /* 远程重置通道密码*/
	MINOR_LOCAL_SSD_UPGRADE_START       = 0x2639 /* 本地SSD文件系统升级开始*/
	MINOR_LOCAL_SSD_UPGRADE_STOP        = 0x2640 /* 本地SSD文件系统升级结束*/
	MINOR_REMOTE_SSD_UPGRADE_START      = 0x2641 /* 远程SSD文件系统升级开始*/
	MINOR_REMOTE_SSD_UPGRADE_STOP       = 0x2642 /* 远程SSD文件系统升级结束*/
	MINOR_LOCAL_SSD_FORMAT_START        = 0x2643 /*本地SSD文件系统格式化开始*/
	MINOR_LOCAL_SSD_FORMAT_STOP         = 0x2644 /*本地SSD文件系统格式化结束*/
	MINOR_REMOTE_SSD_FORMAT_START       = 0x2645 /*远程SSD文件系统格式化开始*/
	MINOR_REMOTE_SSD_FORMAT_STOP        = 0x2646 /*远程SSD文件系统格式化结束*/
	MINOR_LOCAL_AUTO_SWITCH_CONFIG      = 0x2647 /* 本地自动开关机配置*/
	MINOR_REMOTE_AUTO_SWITCH_CONFIG     = 0x2648 /* 远程自动开关机配置*/

	MINOR_LOCAL_SSD_INITIALIZATION_START  = 0x264a /* 本地SSD初始化开始*/
	MINOR_LOCAL_SSD_INITIALIZATION_END    = 0x264b /* 本地SSD初始化结束*/
	MINOR_REMOTE_SSD_INITIALIZATION_START = 0x264c /* 远程SSD初始化开始*/
	MINOR_REMOTE_SSD_INITIALIZATION_END   = 0x264d /* 远程SSD初始化结束*/

	//定义AI开放平台的操作日志
	MINOR_REMOTE_AI_MODEL_ADD                   = 0x2650 //模型包添加
	MINOR_REMOTE_AI_MODEL_QUERY                 = 0x2651 //模型包查询
	MINOR_REMOTE_AI_MODEL_DELETE                = 0x2652 //模型包删除
	MINOR_REMOTE_AI_MODEL_UPDATE                = 0x2653 //模型包更新
	MINOR_REMOTE_AI_PICTURE_POLLING_TASK_ADD    = 0x2654 //图片轮询任务增加
	MINOR_REMOTE_AI_PICTURE_POLLING_TASK_QUERY  = 0x2655 //图片轮询任务查询
	MINOR_REMOTE_AI_PICTURE_POLLING_TASK_DELETE = 0x2656 //图片轮询任务删除
	MINOR_REMOTE_AI_PICTURE_POLLING_TASK_MODIFY = 0x2657 //图片轮询任务修改
	MINOR_REMOTE_AI_VIDEO_POLLING_TASK_ADD      = 0x2658 //视频轮询任务增加
	MINOR_REMOTE_AI_VIDEO_POLLING_TASK_QUERY    = 0x2659 //视频轮询任务查询
	MINOR_REMOTE_AI_VIDEO_POLLING_TASK_DELETE   = 0x265A //视频轮询任务删除
	MINOR_REMOTE_AI_VIDEO_POLLING_TASK_MODIFY   = 0x265B //视频轮询任务修改
	MINOR_REMOTE_AI_PICTURE_TASK_ADD            = 0x265C //图片任务增加
	MINOR_REMOTE_AI_PICTURE_TASK_QUERY          = 0x265D //图片任务查询
	MINOR_REMOTE_AI_PICTURE_TASK_DELETE         = 0x265E //图片任务删除
	MINOR_REMOTE_AI_PICTURE_TASK_MODIFY         = 0x265F //图片任务修改
	MINOR_REMOTE_AI_VIDEO_TASK_ADD              = 0x2660 //视频任务增加
	MINOR_REMOTE_AI_VIDEO_TASK_QUERY            = 0x2661 //视频任务查询
	MINOR_REMOTE_AI_VIDEO_TASK_DELETE           = 0x2662 //视频任务删除
	MINOR_REMOTE_AI_VIDEO_TASK_MODIFY           = 0x2663 //视频任务修改
	MINOR_REMOTE_AI_RULE_CONFIG                 = 0x2664 //AI规则配置

	MINOR_REMOTE_LOG_STORAGE_CONFIG = 0x2665 //日志存储配置
	MINOR_REMOTE_LOG_SERVER_CONFIG  = 0x2666 //日志服务器参数配置

	MINOR_REMOTE_RESET_IPC_PASSWORD = 0x2670 //NVR重置IPC密码日志

	//定义萤石平台操作日志
	MINOR_LOCAL_EZVIZ_OPERATION  = 0x2671 //本地萤石操作(包括萤石参数配置和升级)
	MINOR_REMOTE_EZVIZ_OPERATION = 0x2672 //远程萤石操作(包括萤石参数配置和升级)

	MINOR_EZVIZ_BITSTREAM_PARAMATERS_CONFIG = 0x2673 /* 萤石码流参数配置*/
	MINOR_EZVIZ_ALARM_PARAMATERS_CONFIG     = 0x2674 /* 萤石报警参数配置*/
	MINOR_EZVIZ_UPGRADE                     = 0x2675 /* 萤石升级*/
	MINOR_EZVIZ_REGISTER                    = 0x2676 /* 萤石注册*/
	MINOR_EZVIZ_LOCAL_PARAMATERS_CONFIG     = 0x2677 /* 萤石本地参数配置*/
	MINOR_EZVIZ_REMOTE_PARAMATERS_CONFIG    = 0x2678 /* 萤石远程参数配置*/

	//消防操作日志
	MINOR_STOP_SOUND               = 0x2700 /*消音*/
	MINOR_SELF_CHECK               = 0x2701 /*自检*/
	MINOR_DUTY_CHECK               = 0x2702 /*查岗*/
	MINOR_SWITCH_SIMPLE_WORKMODE   = 0x2703 /*切换至简易模式*/
	MINOR_SWITCH_NORMAL_WORKMODE   = 0x2704 /*切换至标准模式*/
	MINOR_LOCAL_SSD_OPERATE_START  = 0x2705 /* 本地SSD操作开始*/
	MINOR_LOCAL_SSD_OPERATE_STOP   = 0x2706 /* 本地SSD操作结束*/
	MINOR_REMOTE_SSD_OPERATE_START = 0x2707 /* 远程SSD操作开始*/
	MINOR_REMOTE_SSD_OPERATE_STOP  = 0x2708 /* 远程SSD操作结束*/

	/*日志附加信息*/
	//主类型
	MAJOR_INFORMATION = 0x4 /*附加信息*/
	//次类型
	MINOR_HDD_INFO        = 0xa1 /*硬盘信息*/
	MINOR_SMART_INFO      = 0xa2 /*SMART信息*/
	MINOR_REC_START       = 0xa3 /*开始录像*/
	MINOR_REC_STOP        = 0xa4 /*停止录像*/
	MINOR_REC_OVERDUE     = 0xa5 /*过期录像删除*/
	MINOR_LINK_START      = 0xa6 //连接前端设备
	MINOR_LINK_STOP       = 0xa7 //断开前端设备
	MINOR_NET_DISK_INFO   = 0xa8 //网络硬盘信息
	MINOR_RAID_INFO       = 0xa9 //raid相关信息
	MINOR_RUN_STATUS_INFO = 0xaa /*系统运行状态信息*/

	//Netra3.0.0
	MINOR_SPARE_START_BACKUP   = 0xab /*热备系统开始备份指定工作机*/
	MINOR_SPARE_STOP_BACKUP    = 0xac /*热备系统停止备份指定工作机*/
	MINOR_SPARE_CLIENT_INFO    = 0xad /*热备客户机信息*/
	MINOR_ANR_RECORD_START     = 0xae /*ANR录像开始*/
	MINOR_ANR_RECORD_END       = 0xaf /*ANR录像结束*/
	MINOR_ANR_ADD_TIME_QUANTUM = 0xb0 /*ANR添加时间段*/
	MINOR_ANR_DEL_TIME_QUANTUM = 0xb1 /*ANR删除时间段*/

	MINOR_PIC_REC_START   = 0xb3 /* 开始抓图*/
	MINOR_PIC_REC_STOP    = 0xb4 /* 停止抓图*/
	MINOR_PIC_REC_OVERDUE = 0xb5 /* 过期图片文件删除 */
	//Netra3.1.0
	MINOR_CLIENT_LOGIN           = 0xb6 /*登录服务器成功*/
	MINOR_CLIENT_RELOGIN         = 0xb7 /*重新登录服务器*/
	MINOR_CLIENT_LOGOUT          = 0xb8 /*退出服务器成功*/
	MINOR_CLIENT_SYNC_START      = 0xb9 /*录像同步开始*/
	MINOR_CLIENT_SYNC_STOP       = 0xba /*录像同步终止*/
	MINOR_CLIENT_SYNC_SUCC       = 0xbb /*录像同步成功*/
	MINOR_CLIENT_SYNC_EXCP       = 0xbc /*录像同步异常*/
	MINOR_GLOBAL_RECORD_ERR_INFO = 0xbd /*全局错误记录信息*/
	MINOR_BUFFER_STATE           = 0xbe /*缓冲区状态日志记录*/
	MINOR_DISK_ERRORINFO_V2      = 0xbf /*硬盘错误详细信息V2*/
	MINOR_CS_DATA_EXPIRED        = 0xc0 //云存储数据过期
	MINOR_PLAT_INFO              = 0xc1 //平台操作信息
	MINOR_DIAL_STAT              = 0xc2 /*拨号状态*/

	MINOR_UNLOCK_RECORD             = 0xc3 //开锁记录
	MINOR_VIS_ALARM                 = 0xc4 //防区报警
	MINOR_TALK_RECORD               = 0xc5 //通话记录
	MINOR_ACCESSORIES_MESSAGE       = 0xc6 //配件板信息
	MINOR_KMS_EXPAMSION_DISK_INSERT = 0xc7 // KMS扩容盘插入
	MINOR_IPC_CONNECT               = 0xc8 //  IPC连接信息
	MINOR_INTELLIGENT_BOARD_STATUS  = 0xc9 //  智能板状态
	MINOR_IPC_CONNECT_STATUS        = 0xca //  IPC连接状态
	MINOR_AUTO_TIMING               = 0xcb //自动校时
	MINOR_EZVIZ_OPERATION           = 0xcc //萤石运行状态
	//NVR集群
	MINOR_CLUSTER_DEVICE_ONLINE       = 0xcd //集群设备上线
	MINOR_CLUSTER_MGR_SERVICE_STARTUP = 0xce //集群管理服务启动
	MINOR_CLUSTER_BUSINESS_TRANSFER   = 0xcf //集群业务迁移
	MINOR_CLUSTER_STATUS              = 0xd0 //集群状态信息
	MINOR_CLUSTER_CS_STATUS           = 0xd1 //集群CS向CM发送设备状态失败，记录CS和CM的IP地址
	MINOR_CLUSTER_CM_STATUS           = 0xd2 //CM状态切换，记录CM转变的角色，如leader、follower、candidate
	MINOR_VOICE_START_DETECTED        = 0xd3 /*检测到语音开始*/
	MINOR_VOICE_END_DETECTED          = 0xd4 /*检测到语音结束*/
	MINOR_DOUBLE_VERIFICATION_PASS    = 0xd5 /*二次认证通过*/
	MINOR_WIRELESS_RUNNING_STATUS     = 0xd6 /*无线运行状态*/
	MINOR_SYSTEM_DATA_SYNCHRONIZATION = 0xd7 /*系统数据同步*/
	MINOR_HD_FORMAT_START             = 0xd8 /*硬盘格式化开始*/
	MINOR_HD_FORMAT_STOP              = 0xd9 /*硬盘格式化结束*/

	//0x400-0x4ff 门禁附件信息日志类型
	MINOR_LIVE_DETECT_OPEN          = 0x400 //真人检测开启
	MINOR_LIVE_DETECT_CLOSE         = 0x401 //真人检测关闭
	MINOR_CLEAR_DATA_COLLECTION     = 0x402 //采集数据清空
	MINOR_DELETE_DATA_COLLECTION    = 0x403 //采集数据删除
	MINOR_EXPORT_DATA_COLLECTION    = 0x404 //采集数据导出
	MINOR_CARD_LEN_CONFIG           = 0x405 //卡长度配置
	MINOR_DATA_BASE_INIT_FAILED     = 0x406 //数据库初始化失败
	MINOR_DATA_BASE_PATCH_UPDATE    = 0x407 //数据库补丁升级
	MINOR_PSAM_CARD_INSERT          = 0x408 //Psam卡插入
	MINOR_PSAM_CARD_REMOVE          = 0x409 //Psam卡拔出
	MINOR_HARD_FAULT_REBOOT         = 0x40a //硬件异常（hardfault）重启
	MINOR_PSAM_CARD_OCP             = 0x40b //Psam卡过流保护
	MINOR_STACK_OVERFLOW            = 0x40c //堆栈溢出
	MINOR_PARM_CFG                  = 0x40d //参数配置
	MINOR_CLR_USER                  = 0x40e //清空所有用户
	MINOR_CLR_CARD                  = 0x40f //清空所有卡
	MINOR_CLR_FINGER_BY_READER      = 0x410 //清空所有指纹(按读卡器)
	MINOR_CLR_FINGER_BY_CARD        = 0x411 //清空所有指纹(按卡号)
	MINOR_CLR_FINGER_BY_EMPLOYEE_ON = 0x412 //清空所有指纹(按工号)
	MINOR_DEL_FINGER                = 0x413 //删除一个指纹
	MINOR_CLR_WEEK_PLAN             = 0x414 //清除权限周计划
	MINOR_SET_WEEK_PLAN             = 0x415 //设置权限周计划
	MINOR_SET_HOLIDAY_PLAN          = 0x416 //设置权限假日计划
	MINOR_CLR_HOLIDAY_PLAN          = 0x417 //清除权限假日计划
	MINOR_SET_HOLIDAY_GROUP         = 0x418 //设置权限假日组计划
	MINOR_CLR_HOLIDAY_GROUP         = 0x419 //清除权限假日组计划
	MINOR_CLR_TEMPLATE_PLAN         = 0x41a //清除权限计划
	MINOR_SET_TEMPLATE_PLAN         = 0x41b //设置权限计划
	MINOR_ADD_CARD                  = 0x41c //新增卡
	MINOR_MOD_CARD                  = 0x41d //修改卡
	MINOR_ADD_FINGER_BY_CARD        = 0x41e //新增指纹(按卡号)
	MINOR_ADD_FINGER_BY_EMPLOYEE_NO = 0x41f //新增指纹(按工号)
	MINOR_MOD_FINGER_BY_CARD        = 0x420 //修改指纹(按卡号)
	MINOR_MOD_FINGER_BY_EMPLOYEE_NO = 0x421 //修改指纹(按工号)
	MINOR_IMPORT_USER_LIST          = 0x422 //用户列表导入（离线采集）

	//802.1x认证操作日志
	MINOR_802_1X_AUTH_SUCC = 0x320 /*802.1x认证成功*/
	MINOR_802_1X_AUTH_FAIL = 0x321 /*802.1x认证失败*/

	/*事件*/
	//主类型
	MAJOR_EVENT = 0x5 /*事件*/
	//次类型
	MINOR_LEGAL_CARD_PASS                         = 0x01 //合法卡认证通过
	MINOR_CARD_AND_PSW_PASS                       = 0x02 //刷卡加密码认证通过
	MINOR_CARD_AND_PSW_FAIL                       = 0x03 //刷卡加密码认证失败
	MINOR_CARD_AND_PSW_TIMEOUT                    = 0x04 //数卡加密码认证超时
	MINOR_CARD_AND_PSW_OVER_TIME                  = 0x05 //刷卡加密码超次
	MINOR_CARD_NO_RIGHT                           = 0x06 //未分配权限
	MINOR_CARD_INVALID_PERIOD                     = 0x07 //无效时段
	MINOR_CARD_OUT_OF_DATE                        = 0x08 //卡号过期
	MINOR_INVALID_CARD                            = 0x09 //无此卡号
	MINOR_ANTI_SNEAK_FAIL                         = 0x0a //反潜回认证失败
	MINOR_INTERLOCK_DOOR_NOT_CLOSE                = 0x0b //互锁门未关闭
	MINOR_NOT_BELONG_MULTI_GROUP                  = 0x0c //卡不属于多重认证群组
	MINOR_INVALID_MULTI_VERIFY_PERIOD             = 0x0d //卡不在多重认证时间段内
	MINOR_MULTI_VERIFY_SUPER_RIGHT_FAIL           = 0x0e //多重认证模式超级权限认证失败
	MINOR_MULTI_VERIFY_REMOTE_RIGHT_FAIL          = 0x0f //多重认证模式远程认证失败
	MINOR_MULTI_VERIFY_SUCCESS                    = 0x10 //多重认证成功
	MINOR_LEADER_CARD_OPEN_BEGIN                  = 0x11 //首卡开门开始
	MINOR_LEADER_CARD_OPEN_END                    = 0x12 //首卡开门结束
	MINOR_ALWAYS_OPEN_BEGIN                       = 0x13 //常开状态开始
	MINOR_ALWAYS_OPEN_END                         = 0x14 //常开状态结束
	MINOR_LOCK_OPEN                               = 0x15 //门锁打开
	MINOR_LOCK_CLOSE                              = 0x16 //门锁关闭
	MINOR_DOOR_BUTTON_PRESS                       = 0x17 //开门按钮打开
	MINOR_DOOR_BUTTON_RELEASE                     = 0x18 //开门按钮放开
	MINOR_DOOR_OPEN_NORMAL                        = 0x19 //正常开门（门磁）
	MINOR_DOOR_CLOSE_NORMAL                       = 0x1a //正常关门（门磁）
	MINOR_DOOR_OPEN_ABNORMAL                      = 0x1b //门异常打开（门磁）
	MINOR_DOOR_OPEN_TIMEOUT                       = 0x1c //门打开超时（门磁）
	MINOR_ALARMOUT_ON                             = 0x1d //报警输出打开
	MINOR_ALARMOUT_OFF                            = 0x1e //报警输出关闭
	MINOR_ALWAYS_CLOSE_BEGIN                      = 0x1f //常关状态开始
	MINOR_ALWAYS_CLOSE_END                        = 0x20 //常关状态结束
	MINOR_MULTI_VERIFY_NEED_REMOTE_OPEN           = 0x21 //多重多重认证需要远程开门
	MINOR_MULTI_VERIFY_SUPERPASSWD_VERIFY_SUCCESS = 0x22 //多重认证超级密码认证成功事件
	MINOR_MULTI_VERIFY_REPEAT_VERIFY              = 0x23 //多重认证重复认证事件
	MINOR_MULTI_VERIFY_TIMEOUT                    = 0x24 //多重认证重复认证事件
	MINOR_DOORBELL_RINGING                        = 0x25 //门铃响
	MINOR_FINGERPRINT_COMPARE_PASS                = 0x26 //指纹比对通过
	MINOR_FINGERPRINT_COMPARE_FAIL                = 0x27 //指纹比对失败
	MINOR_CARD_FINGERPRINT_VERIFY_PASS            = 0x28 //刷卡加指纹认证通过
	MINOR_CARD_FINGERPRINT_VERIFY_FAIL            = 0x29 //刷卡加指纹认证失败
	MINOR_CARD_FINGERPRINT_VERIFY_TIMEOUT         = 0x2a //刷卡加指纹认证超时
	MINOR_CARD_FINGERPRINT_PASSWD_VERIFY_PASS     = 0x2b //刷卡加指纹加密码认证通过
	MINOR_CARD_FINGERPRINT_PASSWD_VERIFY_FAIL     = 0x2c //刷卡加指纹加密码认证失败
	MINOR_CARD_FINGERPRINT_PASSWD_VERIFY_TIMEOUT  = 0x2d //刷卡加指纹加密码认证超时
	MINOR_FINGERPRINT_PASSWD_VERIFY_PASS          = 0x2e //指纹加密码认证通过
	MINOR_FINGERPRINT_PASSWD_VERIFY_FAIL          = 0x2f //指纹加密码认证失败
	MINOR_FINGERPRINT_PASSWD_VERIFY_TIMEOUT       = 0x30 //指纹加密码认证超时
	MINOR_FINGERPRINT_INEXISTENCE                 = 0x31 //指纹不存在
	MINOR_CARD_PLATFORM_VERIFY                    = 0x32 //刷卡平台认证
	MINOR_CALL_CENTER                             = 0x33 //呼叫中心事件
	MINOR_FIRE_RELAY_TURN_ON_DOOR_ALWAYS_OPEN     = 0x34 //消防继电器导通触发门常开
	MINOR_FIRE_RELAY_RECOVER_DOOR_RECOVER_NORMAL  = 0x35 //消防继电器恢复门恢复正常
	MINOR_FACE_AND_FP_VERIFY_PASS                 = 0x36 //人脸加指纹认证通过
	MINOR_FACE_AND_FP_VERIFY_FAIL                 = 0x37 //人脸加指纹认证失败
	MINOR_FACE_AND_FP_VERIFY_TIMEOUT              = 0x38 //人脸加指纹认证超时
	MINOR_FACE_AND_PW_VERIFY_PASS                 = 0x39 //人脸加密码认证通过
	MINOR_FACE_AND_PW_VERIFY_FAIL                 = 0x3a //人脸加密码认证失败
	MINOR_FACE_AND_PW_VERIFY_TIMEOUT              = 0x3b //人脸加密码认证超时
	MINOR_FACE_AND_CARD_VERIFY_PASS               = 0x3c //人脸加刷卡认证通过
	MINOR_FACE_AND_CARD_VERIFY_FAIL               = 0x3d //人脸加刷卡认证失败
	MINOR_FACE_AND_CARD_VERIFY_TIMEOUT            = 0x3e //人脸加刷卡认证超时
	MINOR_FACE_AND_PW_AND_FP_VERIFY_PASS          = 0x3f //人脸加密码加指纹认证通过
	MINOR_FACE_AND_PW_AND_FP_VERIFY_FAIL          = 0x40 //人脸加密码加指纹认证失败
	MINOR_FACE_AND_PW_AND_FP_VERIFY_TIMEOUT       = 0x41 //人脸加密码加指纹认证超时
	MINOR_FACE_CARD_AND_FP_VERIFY_PASS            = 0x42 //人脸加刷卡加指纹认证通过
	MINOR_FACE_CARD_AND_FP_VERIFY_FAIL            = 0x43 //人脸加刷卡加指纹认证失败
	MINOR_FACE_CARD_AND_FP_VERIFY_TIMEOUT         = 0x44 //人脸加刷卡加指纹认证超时
	MINOR_EMPLOYEENO_AND_FP_VERIFY_PASS           = 0x45 //工号加指纹认证通过
	MINOR_EMPLOYEENO_AND_FP_VERIFY_FAIL           = 0x46 //工号加指纹认证失败
	MINOR_EMPLOYEENO_AND_FP_VERIFY_TIMEOUT        = 0x47 //工号加指纹认证超时
	MINOR_EMPLOYEENO_AND_FP_AND_PW_VERIFY_PASS    = 0x48 //工号加指纹加密码认证通过
	MINOR_EMPLOYEENO_AND_FP_AND_PW_VERIFY_FAIL    = 0x49 //工号加指纹加密码认证失败
	MINOR_EMPLOYEENO_AND_FP_AND_PW_VERIFY_TIMEOUT = 0x4a //工号加指纹加密码认证超时
	MINOR_FACE_VERIFY_PASS                        = 0x4b //人脸认证通过
	MINOR_FACE_VERIFY_FAIL                        = 0x4c //人脸认证失败
	MINOR_EMPLOYEENO_AND_FACE_VERIFY_PASS         = 0x4d //工号加人脸认证通过
	MINOR_EMPLOYEENO_AND_FACE_VERIFY_FAIL         = 0x4e //工号加人脸认证失败
	MINOR_EMPLOYEENO_AND_FACE_VERIFY_TIMEOUT      = 0x4f //工号加人脸认证超时
	MINOR_FACE_RECOGNIZE_FAIL                     = 0x50 //人脸识别失败
	MINOR_FIRSTCARD_AUTHORIZE_BEGIN               = 0x51 //首卡授权开始
	MINOR_FIRSTCARD_AUTHORIZE_END                 = 0x52 //首卡授权结束
	MINOR_DOORLOCK_INPUT_SHORT_CIRCUIT            = 0x53 //门锁输入短路报警
	MINOR_DOORLOCK_INPUT_BROKEN_CIRCUIT           = 0x54 //门锁输入断路报警
	MINOR_DOORLOCK_INPUT_EXCEPTION                = 0x55 //门锁输入异常报警
	MINOR_DOORCONTACT_INPUT_SHORT_CIRCUIT         = 0x56 //门磁输入短路报警
	MINOR_DOORCONTACT_INPUT_BROKEN_CIRCUIT        = 0x57 //门磁输入断路报警
	MINOR_DOORCONTACT_INPUT_EXCEPTION             = 0x58 //门磁输入异常报警
	MINOR_OPENBUTTON_INPUT_SHORT_CIRCUIT          = 0x59 //开门按钮输入短路报警
	MINOR_OPENBUTTON_INPUT_BROKEN_CIRCUIT         = 0x5a //开门按钮输入断路报警
	MINOR_OPENBUTTON_INPUT_EXCEPTION              = 0x5b //开门按钮输入异常报警
	MINOR_DOORLOCK_OPEN_EXCEPTION                 = 0x5c //门锁异常打开
	MINOR_DOORLOCK_OPEN_TIMEOUT                   = 0x5d //门锁打开超时
	MINOR_FIRSTCARD_OPEN_WITHOUT_AUTHORIZE        = 0x5e //首卡未授权开门失败
	MINOR_CALL_LADDER_RELAY_BREAK                 = 0x5f //呼梯继电器断开
	MINOR_CALL_LADDER_RELAY_CLOSE                 = 0x60 //呼梯继电器闭合
	MINOR_AUTO_KEY_RELAY_BREAK                    = 0x61 //自动按键继电器断开
	MINOR_AUTO_KEY_RELAY_CLOSE                    = 0x62 //自动按键继电器闭合
	MINOR_KEY_CONTROL_RELAY_BREAK                 = 0x63 //按键梯控继电器断开
	MINOR_KEY_CONTROL_RELAY_CLOSE                 = 0x64 //按键梯控继电器闭合
	MINOR_EMPLOYEENO_AND_PW_PASS                  = 0x65 //工号加密码认证通过
	MINOR_EMPLOYEENO_AND_PW_FAIL                  = 0x66 //工号加密码认证失败
	MINOR_EMPLOYEENO_AND_PW_TIMEOUT               = 0x67 //工号加密码认证超时
	MINOR_HUMAN_DETECT_FAIL                       = 0x68 //真人检测失败
	MINOR_PEOPLE_AND_ID_CARD_COMPARE_PASS         = 0x69 //人证比对通过
	MINOR_PEOPLE_AND_ID_CARD_COMPARE_FAIL         = 0x70 //人证比对失败
	MINOR_CERTIFICATE_BLACK_LIST                  = 0x71 //黑名单事件
	MINOR_LEGAL_MESSAGE                           = 0x72 //合法短信
	MINOR_ILLEGAL_MESSAGE                         = 0x73 //非法短信
	MINOR_DOOR_OPEN_OR_DORMANT_FAIL               = 0x75 //门状态常闭或休眠状态认证失败
	MINOR_AUTH_PLAN_DORMANT_FAIL                  = 0x76 //认证计划休眠模式认证失败
	MINOR_CARD_ENCRYPT_VERIFY_FAIL                = 0x77 //卡加密校验失败
	MINOR_SUBMARINEBACK_REPLY_FAIL                = 0x78 //反潜回服务器应答失败
	MINOR_DOOR_OPEN_OR_DORMANT_OPEN_FAIL          = 0x82 //门常闭或休眠时开门按钮开门失败
	MINOR_HEART_BEAT                              = 0x83 //心跳事件
	MINOR_DOOR_OPEN_OR_DORMANT_LINKAGE_OPEN_FAIL  = 0x84 //门常闭或休眠时开门联动开门失败
	MINOR_TRAILING                                = 0x85 //尾随通行
	MINOR_REVERSE_ACCESS                          = 0x86 //反向闯入
	MINOR_FORCE_ACCESS                            = 0x87 //外力冲撞
	MINOR_CLIMBING_OVER_GATE                      = 0x88 //翻越
	MINOR_PASSING_TIMEOUT                         = 0x89 //通行超时
	MINOR_INTRUSION_ALARM                         = 0x8a //误闯报警
	MINOR_FREE_GATE_PASS_NOT_AUTH                 = 0x8b //闸机自由通行时未认证通过
	MINOR_DROP_ARM_BLOCK                          = 0x8c //摆臂被阻挡
	MINOR_DROP_ARM_BLOCK_RESUME                   = 0x8d //摆臂阻挡消除
	MINOR_LOCAL_FACE_MODELING_FAIL                = 0x8e //设备升级本地人脸建模失败
	MINOR_STAY_EVENT                              = 0x8f //逗留事件
	MINOR_PASSWORD_MISMATCH                       = 0x97 //密码不匹配
	MINOR_EMPLOYEE_NO_NOT_EXIST                   = 0x98 //工号不存在
	MINOR_COMBINED_VERIFY_PASS                    = 0x99 //组合认证通过
	MINOR_COMBINED_VERIFY_TIMEOUT                 = 0x9a //组合认证超时
	MINOR_VERIFY_MODE_MISMATCH                    = 0x9b //认证方式不匹配

	MINOR_PASSPORT_VERIFY_FAIL             = 0xa1 //护照信息校验失败
	MINOR_INFORMAL_MIFARE_CARD_VERIFY_FAIL = 0xa2 //非正规Mifare卡认证失败
	MINOR_CPU_CARD_ENCRYPT_VERIFY_FAIL     = 0xa3 //CPU卡加密校验失败
	MINOR_NFC_DISABLE_VERIFY_FAIL          = 0xa4 //NFC功能关闭验证失败

	MINOR_LORA_MODULE_ONLINE  = 0xa5 //LoRa模块上线
	MINOR_LORA_MODULE_OFFLINE = 0xa6 //LoRa模块下线
	MINOR_MQTT_STATUS         = 0xa7 //Mqtt连接状态

	MINOR_EM_CARD_RECOGNIZE_NOT_ENABLED           = 0xa8 //EM卡识别未启用
	MINOR_M1_CARD_RECOGNIZE_NOT_ENABLED           = 0xa9 //M1卡识别未启用
	MINOR_CPU_CARD_RECOGNIZE_NOT_ENABLED          = 0xaa //CPU卡识别未启用
	MINOR_ID_CARD_RECOGNIZE_NOT_ENABLED           = 0xab //身份证识别未启用
	MINOR_CARD_SET_SECRET_KEY_FAIL                = 0xac //卡灌装密钥失败
	MINOR_LOCAL_UPGRADE_FAIL                      = 0xad /* 本地升级失败 */
	MINOR_REMOTE_UPGRADE_FAIL                     = 0xae /* 远程升级失败 */
	MINOR_REMOTE_EXTEND_MODULE_UPGRADE_SUCC       = 0xaf /*远程扩展模块升级成功*/
	MINOR_REMOTE_EXTEND_MODULE_UPGRADE_FAIL       = 0xb0 /*远程扩展模块升级失败*/
	MINOR_REMOTE_FINGER_PRINT_MODULE_UPGRADE_SUCC = 0xb1 /*远程指纹模组升级成功*/
	MINOR_REMOTE_FINGER_PRINT_MODULE_UPGRADE_FAIL = 0xb2 /*远程指纹模组升级失败*/
	MINOR_PASSWD_VERIFY_PASS                      = 0xb5 //密码认证通过

	MINOR_EVENT_CUSTOM1  = 0x500 //门禁自定义事件1
	MINOR_EVENT_CUSTOM2  = 0x501 //门禁自定义事件2
	MINOR_EVENT_CUSTOM3  = 0x502 //门禁自定义事件3
	MINOR_EVENT_CUSTOM4  = 0x503 //门禁自定义事件4
	MINOR_EVENT_CUSTOM5  = 0x504 //门禁自定义事件5
	MINOR_EVENT_CUSTOM6  = 0x505 //门禁自定义事件6
	MINOR_EVENT_CUSTOM7  = 0x506 //门禁自定义事件7
	MINOR_EVENT_CUSTOM8  = 0x507 //门禁自定义事件8
	MINOR_EVENT_CUSTOM9  = 0x508 //门禁自定义事件9
	MINOR_EVENT_CUSTOM10 = 0x509 //门禁自定义事件10
	MINOR_EVENT_CUSTOM11 = 0x50a //门禁自定义事件11
	MINOR_EVENT_CUSTOM12 = 0x50b //门禁自定义事件12
	MINOR_EVENT_CUSTOM13 = 0x50c //门禁自定义事件13
	MINOR_EVENT_CUSTOM14 = 0x50d //门禁自定义事件14
	MINOR_EVENT_CUSTOM15 = 0x50e //门禁自定义事件15
	MINOR_EVENT_CUSTOM16 = 0x50f //门禁自定义事件16
	MINOR_EVENT_CUSTOM17 = 0x510 //门禁自定义事件17
	MINOR_EVENT_CUSTOM18 = 0x511 //门禁自定义事件18
	MINOR_EVENT_CUSTOM19 = 0x512 //门禁自定义事件19
	MINOR_EVENT_CUSTOM20 = 0x513 //门禁自定义事件20
	MINOR_EVENT_CUSTOM21 = 0x514 //门禁自定义事件21
	MINOR_EVENT_CUSTOM22 = 0x515 //门禁自定义事件22
	MINOR_EVENT_CUSTOM23 = 0x516 //门禁自定义事件23
	MINOR_EVENT_CUSTOM24 = 0x517 //门禁自定义事件24
	MINOR_EVENT_CUSTOM25 = 0x518 //门禁自定义事件25
	MINOR_EVENT_CUSTOM26 = 0x519 //门禁自定义事件26
	MINOR_EVENT_CUSTOM27 = 0x51a //门禁自定义事件27
	MINOR_EVENT_CUSTOM28 = 0x51b //门禁自定义事件28
	MINOR_EVENT_CUSTOM29 = 0x51c //门禁自定义事件29
	MINOR_EVENT_CUSTOM30 = 0x51d //门禁自定义事件30
	MINOR_EVENT_CUSTOM31 = 0x51e //门禁自定义事件31
	MINOR_EVENT_CUSTOM32 = 0x51f //门禁自定义事件32
	MINOR_EVENT_CUSTOM33 = 0x520 //门禁自定义事件33
	MINOR_EVENT_CUSTOM34 = 0x521 //门禁自定义事件34
	MINOR_EVENT_CUSTOM35 = 0x522 //门禁自定义事件35
	MINOR_EVENT_CUSTOM36 = 0x523 //门禁自定义事件36
	MINOR_EVENT_CUSTOM37 = 0x524 //门禁自定义事件37
	MINOR_EVENT_CUSTOM38 = 0x525 //门禁自定义事件38
	MINOR_EVENT_CUSTOM39 = 0x526 //门禁自定义事件39
	MINOR_EVENT_CUSTOM40 = 0x527 //门禁自定义事件40
	MINOR_EVENT_CUSTOM41 = 0x528 //门禁自定义事件41
	MINOR_EVENT_CUSTOM42 = 0x529 //门禁自定义事件42
	MINOR_EVENT_CUSTOM43 = 0x52a //门禁自定义事件43
	MINOR_EVENT_CUSTOM44 = 0x52b //门禁自定义事件44
	MINOR_EVENT_CUSTOM45 = 0x52c //门禁自定义事件45
	MINOR_EVENT_CUSTOM46 = 0x52d //门禁自定义事件46
	MINOR_EVENT_CUSTOM47 = 0x52e //门禁自定义事件47
	MINOR_EVENT_CUSTOM48 = 0x52f //门禁自定义事件48
	MINOR_EVENT_CUSTOM49 = 0x530 //门禁自定义事件49
	MINOR_EVENT_CUSTOM50 = 0x531 //门禁自定义事件50
	MINOR_EVENT_CUSTOM51 = 0x532 //门禁自定义事件51
	MINOR_EVENT_CUSTOM52 = 0x533 //门禁自定义事件52
	MINOR_EVENT_CUSTOM53 = 0x534 //门禁自定义事件53
	MINOR_EVENT_CUSTOM54 = 0x535 //门禁自定义事件54
	MINOR_EVENT_CUSTOM55 = 0x536 //门禁自定义事件55
	MINOR_EVENT_CUSTOM56 = 0x537 //门禁自定义事件56
	MINOR_EVENT_CUSTOM57 = 0x538 //门禁自定义事件57
	MINOR_EVENT_CUSTOM58 = 0x539 //门禁自定义事件58
	MINOR_EVENT_CUSTOM59 = 0x53a //门禁自定义事件59
	MINOR_EVENT_CUSTOM60 = 0x53b //门禁自定义事件60
	MINOR_EVENT_CUSTOM61 = 0x53c //门禁自定义事件61
	MINOR_EVENT_CUSTOM62 = 0x53d //门禁自定义事件62
	MINOR_EVENT_CUSTOM63 = 0x53e //门禁自定义事件63
	MINOR_EVENT_CUSTOM64 = 0x53f //门禁自定义事件64

	MINOR_LOCK_FINGER_OPEN_DOOR        = 0x600 //智能锁指纹开门
	MINOR_LOCK_PASSWORD_OPEN_DOOR      = 0x601 //智能锁密码开门
	MINOR_LOCK_CARD_OPEN_DOOR          = 0x602 //智能锁刷卡开门
	MINOR_LOCK_CENTER_OPEN_DOOR        = 0x603 //智能锁中心开门
	MINOR_LOCK_APP_OPEN_DOOR           = 0x604 //智能锁APP开门
	MINOR_LOCK_KEY_OPEN_DOOR           = 0x605 //智能锁钥匙开门
	MINOR_LOCK_REMOTE_DEVICE_OPEN_DOOR = 0x606 //智能锁遥控器开门
	MINOR_LOCK_TMP_PASSWORD_OPEN_DOOR  = 0x607 //智能锁临时密码开门
	MINOR_LOCK_BLUETOOTH_OPEN_DOOR     = 0x608 //智能锁蓝牙开门
	MINOR_LOCK_MULTI_OPEN_DOOR         = 0x609 //智能锁多重开门

	//2018-04-23 通用物联网关事件日志类型
	MINOR_ALARMHOST_SCHOOLTIME_IRGI_B         = 0x1001 //B码校时
	MINOR_ALARMHOST_SCHOOLTIME_SDK            = 0x1002 //SDK校时
	MINOR_ALARMHOST_SCHOOLTIME_SELFTEST       = 0x1003 //定制自检校时
	MINOR_ALARMHOST_SUBSYSTEM_ABNORMALINSERT  = 0x1004 //子板插入
	MINOR_ALARMHOST_SUBSYSTEM_ABNORMALPULLOUT = 0x1005 //子板拔出

	MINOR_ALARMHOST_AUTO_ARM              = 0x1006 //自动布防
	MINOR_ALARMHOST_AUTO_DISARM           = 0x1007 //自动撤防
	MINOR_ALARMHOST_TIME_TIGGER_ON        = 0x1008 //定时开启触发器
	MINOR_ALARMHOST_TIME_TIGGER_OFF       = 0x1009 //定时关闭触发器
	MINOR_ALARMHOST_AUTO_ARM_FAILD        = 0x100a //自动布防失败
	MINOR_ALARMHOST_AUTO_DISARM_FAILD     = 0x100b //自动撤防失败
	MINOR_ALARMHOST_TIME_TIGGER_ON_FAILD  = 0x100c //定时开启触发器失败
	MINOR_ALARMHOST_TIME_TIGGER_OFF_FAILD = 0x100d //定时关闭触发器失败
	MINOR_ALARMHOST_MANDATORY_ALARM       = 0x100e //强制布防
	MINOR_ALARMHOST_KEYPAD_LOCKED         = 0x100f //键盘锁定
	MINOR_ALARMHOST_USB_INSERT            = 0x1010 //USB插入
	MINOR_ALARMHOST_USB_PULLOUT           = 0x1011 //USB拔出
	MINOR_ALARMHOST_4G_MODULS_ONLINE      = 0x1012 //4G模块上线
	MINOR_ALARMHOST_4G_MODULS_OFFLINE     = 0x1013 //4G模块下线

	MINOR_EZVIZ_CLOUD_ONLINE  = 0x1014 //萤石云上线
	MINOR_EZVIZ_CLOUD_OFFLINE = 0x1015 //萤石云下线

	MINOR_SIPUA_GRID_ONLINE  = 0x1016 //国网B上线
	MINOR_SIPUA_GRID_OFFLINE = 0x1017 //国网B下线

	MINOR_INTERNET_ACCESS_CONNECTED = 0x1018 //网口连接
	MINOR_INTERNET_ACCESS_BREAK     = 0x1019 //网口断开

	MINOR_WIRELESS_CONNECTED      = 0x101a //无线连接
	MINOR_WIRELESS_BREAK          = 0x101b //无线断开
	MINOR_PORT_LINK_DOWN          = 0x101c //端口网络down
	MINOR_PORT_LINK_UP            = 0x101d //端口网络up
	MINOR_POE_PORT_POWER_ON       = 0x101e //POE端口power on
	MINOR_POE_PORT_POWER_OFF      = 0x101f //POE端口power off
	MINOR_POE_TOTAL_POWER_MAX     = 0x1020 //POE总功率达到poe-max
	MINNOR_POE_TOTAL_POWER_RESUME = 0x1021 //POE总功率恢复正常

	//当日志的主类型为MAJOR_OPERATION=03，次类型为MINOR_LOCAL_CFG_PARM=0x52或者MINOR_REMOTE_GET_PARM=0x76
	//或者MINOR_REMOTE_CFG_PARM=0x77时，dwParaType:参数类型有效，其含义如下：
	PARA_VIDEOOUT  = 0x1
	PARA_IMAGE     = 0x2
	PARA_ENCODE    = 0x4
	PARA_NETWORK   = 0x8
	PARA_ALARM     = 0x10
	PARA_EXCEPTION = 0x20
	PARA_DECODER   = 0x40 /*解码器*/
	PARA_RS232     = 0x80
	PARA_PREVIEW   = 0x100
	PARA_SECURITY  = 0x200
	PARA_DATETIME  = 0x400
	PARA_FRAMETYPE = 0x800  /*帧格式*/
	PARA_DETECTION = 0x1000 //侦测配置
	PARA_VCA_RULE  = 0x1001 //行为规则
	PARA_VCA_CTRL  = 0x1002 //配置智能控制信息
	PARA_VCA_PLATE = 0x1003 // 车牌识别

	PARA_CODESPLITTER = 0x2000 /*码分器参数*/
	//2010-01-22 增加视频综合平台日志信息次类型
	PARA_RS485    = 0x2001 /* RS485配置信息*/
	PARA_DEVICE   = 0x2002 /* 设备配置信息*/
	PARA_HARDDISK = 0x2003 /* 硬盘配置信息 */
	PARA_AUTOBOOT = 0x2004 /* 自动重启配置信息*/
	PARA_HOLIDAY  = 0x2005 /* 节假日配置信息*/
	PARA_IPC      = 0x2006 /* IP通道配置 */
	/*************************参数配置命令 end*******************************/

	/*******************查找文件和日志函数返回值*************************/
	NET_DVR_FILE_SUCCESS   = 1000 //获得文件信息
	NET_DVR_FILE_NOFIND    = 1001 //没有文件
	NET_DVR_ISFINDING      = 1002 //正在查找文件
	NET_DVR_NOMOREFILE     = 1003 //查找文件时没有更多的文件
	NET_DVR_FILE_EXCEPTION = 1004 //查找文件时异常

	/*********************回调函数类型 begin************************/

	//报警回调命令
	COMM_ALARM = 0x1100 //8000报警信息主动上传

	//对应NET_VCA_RULE_ALARM
	COMM_ALARM_RULE          = 0x1102 //行为分析报警信息
	COMM_ALARM_PDC           = 0x1103 //人数统计报警信息
	COMM_ALARM_VIDEOPLATFORM = 0x1104 //视频综合平台报警
	COMM_ALARM_ALARMHOST     = 0x1105 //网络报警主机报警
	COMM_ALARM_FACE          = 0x1106 //人脸检测识别报警信息
	COMM_RULE_INFO_UPLOAD    = 0x1107 // 事件数据信息上传
	COMM_ALARM_AID           = 0x1110 //交通事件报警信息
	COMM_ALARM_TPS           = 0x1111 //交通参数统计报警信息
	//智能人脸抓拍结果上传
	COMM_UPLOAD_FACESNAP_RESULT          = 0x1112 //人脸识别结果上传
	COMM_ALARM_TFS                       = 0x1113 //交通取证报警信息
	COMM_ALARM_TPS_V41                   = 0x1114 //交通参数统计报警信息扩展
	COMM_ALARM_AID_V41                   = 0x1115 //交通事件报警信息扩展
	COMM_ALARM_VQD_EX                    = 0x1116 //视频质量诊断报警
	COMM_ALARM_NOTIFICATION_REPORT       = 0x1117 //通知事件上报
	COMM_SENSOR_VALUE_UPLOAD             = 0x1120 //模拟量数据实时上传
	COMM_SENSOR_ALARM                    = 0x1121 //模拟量报警上传
	COMM_SWITCH_ALARM                    = 0x1122 //开关量报警
	COMM_ALARMHOST_EXCEPTION             = 0x1123 //报警主机故障报警
	COMM_ALARMHOST_OPERATEEVENT_ALARM    = 0x1124 //操作事件报警上传
	COMM_ALARMHOST_SAFETYCABINSTATE      = 0x1125 //防护舱状态
	COMM_ALARMHOST_ALARMOUTSTATUS        = 0x1126 //报警输出口/警号状态
	COMM_ALARMHOST_CID_ALARM             = 0x1127 //报告报警上传
	COMM_ALARMHOST_EXTERNAL_DEVICE_ALARM = 0x1128 //报警主机外接设备报警上传
	COMM_ALARMHOST_DATA_UPLOAD           = 0x1129 //报警数据上传
	COMM_FACECAPTURE_STATISTICS_RESULT   = 0x112a //人脸抓拍统计上传
	COMM_ALARM_WIRELESS_INFO             = 0x122b // 无线网络信息上传
	COMM_SCENECHANGE_DETECTION_UPLOAD    = 0x1130 //场景变更报警上传(布防)2013-7-16
	COMM_CROSSLINE_ALARM                 = 0x1131 //压线报警(监听) 2013-09-27
	COMM_UPLOAD_VIDEO_INTERCOM_EVENT     = 0x1132 //可视对讲事件记录上传
	COMM_ALARM_VIDEO_INTERCOM            = 0x1133 //可视对讲报警上传
	COMM_UPLOAD_NOTICE_DATA              = 0x1134 //可视对讲公告信息上传
	COMM_ALARM_AUDIOEXCEPTION            = 0x1150 //声音报警信息
	COMM_ALARM_DEFOCUS                   = 0x1151 //虚焦报警信息
	COMM_ALARM_BUTTON_DOWN_EXCEPTION     = 0x1152 //按钮按下报警信息
	COMM_ALARM_ALARMGPS                  = 0x1202 //GPS报警信息上传
	COMM_TRADEINFO                       = 0x1500 //ATMDVR主动上传交易信息
	COMM_UPLOAD_PLATE_RESULT             = 0x2800 //上传车牌信息
	COMM_ITC_STATUS_DETECT_RESULT        = 0x2810 //实时状态检测结果上传(智能高清IPC)
	COMM_IPC_AUXALARM_RESULT             = 0x2820 //PIR报警、无线报警、呼救报警上传
	COMM_UPLOAD_PICTUREINFO              = 0x2900 //上传图片信息
	COMM_SNAP_MATCH_ALARM                = 0x2902 //黑名单比对结果上传
	COMM_ITS_PLATE_RESULT                = 0x3050 //终端图片上传
	COMM_ITS_TRAFFIC_COLLECT             = 0x3051 //终端统计数据上传
	COMM_ITS_GATE_VEHICLE                = 0x3052 //出入口车辆抓拍数据上传
	COMM_ITS_GATE_FACE                   = 0x3053 //出入口人脸抓拍数据上传
	COMM_ITS_GATE_COSTITEM               = 0x3054 //出入口过车收费明细 2013-11-19
	COMM_ITS_GATE_HANDOVER               = 0x3055 //出入口交接班数据 2013-11-19
	COMM_ITS_PARK_VEHICLE                = 0x3056 //停车场数据上传
	COMM_ITS_BLACKLIST_ALARM             = 0x3057 //黑名单报警上传

	COMM_VEHICLE_CONTROL_LIST_DSALARM = 0x3058 //黑白名单数据需要同步报警2013-11-04
	COMM_VEHICLE_CONTROL_ALARM        = 0x3059 //车辆报警2013-11-04
	COMM_FIRE_ALARM                   = 0x3060 //消防报警2013-11-04

	COMM_ITS_GATE_ALARMINFO = 0x3061 //出入口控制机数据上传

	COMM_VEHICLE_RECOG_RESULT = 0x3062 //车辆二次识别结果上传 2014-11-12
	COMM_PLATE_RESULT_V50     = 0x3063 //车牌上传 V50

	COMM_GATE_CHARGEINFO_UPLOAD      = 0x3064 //出入口付费信息上传
	COMM_TME_VEHICLE_INDENTIFICATION = 0x3065 //TME车辆抓图上传
	COMM_GATE_CARDINFO_UPLOAD        = 0x3066 //出入口卡片信息上传
	COMM_LOADING_DOCK_OPERATEINFO    = 0x3067 //月台作业上传

	COMM_ALARM_SENSORINFO_UPLOAD = 0x3077 //传感器上传信息
	COMM_ALARM_CAPTURE_UPLOAD    = 0x3078 //抓拍图片上传

	COMM_ITS_RADARINFO = 0x3079 //雷达报警上传

	COMM_SIGNAL_LAMP_ABNORMAL = 0x3080 //信号灯异常检测上传

	COMM_ALARM_TPS_REAL_TIME  = 0x3081 //TPS实时过车数据上传
	COMM_ALARM_TPS_STATISTICS = 0x3082 //TPS统计过车数据上传

	COMM_ALARM_V30       = 0x4000 //9000报警信息主动上传
	COMM_IPCCFG          = 0x4001 //9000设备IPC接入配置改变报警信息主动上传
	COMM_IPCCFG_V31      = 0x4002 //9000设备IPC接入配置改变报警信息主动上传扩展 9000_1.1
	COMM_IPCCFG_V40      = 0x4003 // IVMS 2000 编码服务器 NVR IPC接入配置改变时报警信息上传
	COMM_ALARM_DEVICE    = 0x4004 //设备报警内容，由于通道值大于256而扩展
	COMM_ALARM_CVR       = 0x4005 //CVR 2.0.X外部报警类型
	COMM_ALARM_HOT_SPARE = 0x4006 //热备异常报警（N+1模式异常报警）
	COMM_ALARM_V40       = 0x4007 //移动侦测，视频丢失，遮挡，IO信号量等报警信息主动上传，报警数据为可变长

	COMM_UPLOAD_HEATMAP_RESULT              = 0x4008 //热度图报警上传 2014-03-21
	COMM_ALARM_DEVICE_V40                   = 0x4009 //设备报警内容扩展
	COMM_ALARM_FACE_DETECTION               = 0x4010 //人脸侦测报警
	COMM_ALARM_TARGET_LEFT_REGION           = 0x4011 //检测目标离开检测区域报警(教师走向学生报警(用于联动切换录播主机控制检测学生的球机))
	COMM_GISINFO_UPLOAD                     = 0x4012 //GIS信息上传
	COMM_VANDALPROOF_ALARM                  = 0x4013 //上传防破坏报警信息
	COMM_PEOPLE_DETECTION_UPLOAD            = 0x4014 //人员侦测信息上传
	COMM_ALARM_STORAGE_DETECTION            = 0x4015 //存储智能检测报警上传
	COMM_MVM_REGISTER                       = 0x4016 //地磁管理器（Magnetic Vehicle Manager）注册
	COMM_MVM_STATUS_INFO                    = 0x4017 //地磁管理器（Magnetic Vehicle Manager）状态上报
	COMM_UPLOAD_HEATMAP_RESULT_PDC          = 0x4018 //热度图按人数统计数据上传事件
	COMM_UPLOAD_HEATMAP_RESULT_DURATION     = 0x4019 //热度图按人员停留时间统计数据上传事件
	COMM_UPLOAD_HEATMAP_RESULT_INTERSECTION = 0x4020 //路口分析热度值结果上传
	COMM_UPLOAD_AIOP_VIDEO                  = 0x4021 //设备支持AI开放平台接入，上传视频检测数据
	COMM_UPLOAD_AIOP_PICTURE                = 0x4022 //设备支持AI开放平台接入，上传图片检测数据
	COMM_UPLOAD_AIOP_POLLING_SNAP           = 0x4023 //设备支持AI开放平台接入，上传轮巡抓图图片检测数据 对应的结构体(NET_AIOP_POLLING_PICTURE_HEAD)
	COMM_UPLOAD_AIOP_POLLING_VIDEO          = 0x4024 //设备支持AI开放平台接入，上传轮巡视频检测数据 对应的结构体(NET_AIOP_POLLING_VIDEO_HEAD)

	COMM_ITS_ROAD_EXCEPTION         = 0x4500 //路口设备异常报警
	COMM_ITS_EXTERNAL_CONTROL_ALARM = 0x4520 //外控报警
	COMM_ALARM_SHIPSDETECTION       = 0x4521 // 船只检测报警信息

	COMM_VCA_DBD_ALARM     = 0x4550 //驾驶行为报警信息
	COMM_VCA_ADAS_ALARM    = 0x4551 //高级辅助驾驶报警信息
	COMM_VEH_REALTIME_INFO = 0x4552 //行车实时数据信息
	COMM_VCA_ATTEND_ALARM  = 0x4553 //考勤事件报警信息

	COMM_FIREDETECTION_ALARM     = 0x4991 //火点检测报警
	COMM_ALARM_DENSEFOGDETECTION = 0x4992 //大雾检测报警信息
	COMM_VCA_ALARM               = 0x4993 //智能检测报警
	COMM_FACE_THERMOMETRY_ALARM  = 0x4994 //人脸测温报警

	COMM_TAPE_ARCHIVE_ALARM = 0x4996 //磁带库归档报警

	COMM_SCREEN_ALARM          = 0x5000 //多屏控制器报警类型
	COMM_DVCS_STATE_ALARM      = 0x5001 //分布式大屏控制器报警上传
	COMM_ALARM_ACS             = 0x5002 //门禁主机报警
	COMM_ALARM_FIBER_CONVERT   = 0x5003 //光纤收发器报警
	COMM_ALARM_SWITCH_CONVERT  = 0x5004 //交换机报警
	COMM_ALARM_DEC_VCA         = 0x5010 //智能解码报警
	COMM_ALARM_LCD             = 0x5011 //屏幕报警
	COMM_CONFERENCE_CALL_ALARM = 0x5012 //会议呼叫告警

	COMM_ALARM_WALL_CONFERNECE = 0x5015 //MCU单个已开会的会议信息报警

	COMM_DIAGNOSIS_UPLOAD    = 0x5100 //诊断服务器VQD报警上传
	COMM_HIGH_DENSITY_UPLOAD = 0x5101 //人员聚集密度输出报警上传

	COMM_ID_INFO_ALARM      = 0x5200 //身份证信息上传
	COMM_PASSNUM_INFO_ALARM = 0x5201 //通行人数上报
	COMM_PASSPORT_ALARM     = 0x5202 //护照信息上传

	COMM_THERMOMETRY_DIFF_ALARM   = 0x5211 //温差报警上传
	COMM_THERMOMETRY_ALARM        = 0x5212 //温度报警上传
	COMM_PANORAMIC_LINKAGE_ALARM  = 0x5213 //全景联动到位上传
	COMM_TAG_INFO_ALARM           = 0x5215 // 标签信息上传
	COMM_ALARM_VQD                = 0x6000 //VQD主动报警上传
	COMM_PUSH_UPDATE_RECORD_INFO  = 0x6001 //推模式录像信息上传
	COMM_SWITCH_LAMP_ALARM        = 0x6002 //开关灯检测
	COMM_INQUEST_ALARM            = 0x6005 // 审讯主机报警上传
	COMM_VIDEO_PARKING_POLE_ALARM = 0x6006 //视频桩报警
	COMM_GPS_STATUS_ALARM         = 0x6010 // GPS状态上传
	COMM_BASE_STATION_INFO_ALARM  = 0x6011 //基站信息上传
	COMM_ALARM_SUBSCRIBE_EVENT    = 0x6012 //订阅结果上报

	COMM_FACESNAP_RAWDATA_ALARM = 0x6015 //人脸比对报警（数据透传方式）
	COMM_CLUSTER_ALARM          = 0x6020 //集群报警上传

	COMM_ISAPI_ALARM = 0x6009

	//PJ01C20170209084超脑录播NVS软件功能开发定制项目专用
	COMM_FRAMES_PEOPLE_COUNTING_ALARM = 0x6069 //单帧画面人数统计结果上传

	COMM_SIGN_ABNORMAL_ALARM = 0x6120 //体征异常报警
	COMM_HFPD_ALARM          = 0x6121 //高频人员检测报警

	COMM_HCU_ALARM = 0x6150 //车载智能盒子报警（本报警SDK库不实现，用作设备和DVR通信协议，占位防止冲突）

	COMM_DEV_STATUS_CHANGED = 0x7000 //设备状态改变报警上传

	/*************操作异常类型(消息方式, 回调方式(保留))****************/
	EXCEPTION_EXCHANGE                 = 0x8000 //用户交互时异常
	EXCEPTION_AUDIOEXCHANGE            = 0x8001 //语音对讲异常
	EXCEPTION_ALARM                    = 0x8002 //报警异常
	EXCEPTION_PREVIEW                  = 0x8003 //网络预览异常
	EXCEPTION_SERIAL                   = 0x8004 //透明通道异常
	EXCEPTION_RECONNECT                = 0x8005 //预览时重连
	EXCEPTION_ALARMRECONNECT           = 0x8006 //报警时重连
	EXCEPTION_SERIALRECONNECT          = 0x8007 //透明通道重连
	SERIAL_RECONNECTSUCCESS            = 0x8008 //透明通道重连成功
	EXCEPTION_PLAYBACK                 = 0x8010 //回放异常
	EXCEPTION_DISKFMT                  = 0x8011 //硬盘格式化
	EXCEPTION_PASSIVEDECODE            = 0x8012 //被动解码异常
	EXCEPTION_EMAILTEST                = 0x8013 //邮件测试异常
	EXCEPTION_BACKUP                   = 0x8014 //备份异常
	PREVIEW_RECONNECTSUCCESS           = 0x8015 //预览时重连成功
	ALARM_RECONNECTSUCCESS             = 0x8016 //报警时重连成功
	RESUME_EXCHANGE                    = 0x8017 //用户交互恢复
	NETWORK_FLOWTEST_EXCEPTION         = 0x8018 //网络流量检测异常
	EXCEPTION_PICPREVIEWRECONNECT      = 0x8019 //图片预览重连
	PICPREVIEW_RECONNECTSUCCESS        = 0x8020 //图片预览重连成功
	EXCEPTION_PICPREVIEW               = 0x8021 //图片预览异常
	EXCEPTION_MAX_ALARM_INFO           = 0x8022 //报警信息缓存已达上限
	EXCEPTION_LOST_ALARM               = 0x8023 //报警丢失
	EXCEPTION_PASSIVETRANSRECONNECT    = 0x8024 //被动转码重连
	PASSIVETRANS_RECONNECTSUCCESS      = 0x8025 //被动转码重连成功
	EXCEPTION_PASSIVETRANS             = 0x8026 //被动转码异常
	SUCCESS_PUSHDEVLOGON               = 0x8030 //推模式设备注册成功
	EXCEPTION_RELOGIN                  = 0x8040 //用户重登陆
	RELOGIN_SUCCESS                    = 0x8041 //用户重登陆成功
	EXCEPTION_PASSIVEDECODE_RECONNNECT = 0x8042 //被动解码重连
	EXCEPTION_CLUSTER_CS_ARMFAILED     = 0x8043 //集群报警异常

	EXCEPTION_RELOGIN_FAILED                  = 0x8044 //重登陆失败，停止重登陆
	EXCEPTION_PREVIEW_RECONNECT_CLOSED        = 0x8045 //关闭预览重连功能
	EXCEPTION_ALARM_RECONNECT_CLOSED          = 0x8046 //关闭报警重连功能
	EXCEPTION_SERIAL_RECONNECT_CLOSED         = 0x8047 //关闭透明通道重连功能
	EXCEPTION_PIC_RECONNECT_CLOSED            = 0x8048 //关闭回显重连功能
	EXCEPTION_PASSIVE_DECODE_RECONNECT_CLOSED = 0x8049 //关闭被动解码重连功能
	EXCEPTION_PASSIVE_TRANS_RECONNECT_CLOSED  = 0x804a //关闭被动转码重连功能
	EXCEPTION_VIDEO_DOWNLOAD                  = 0x804b // [add] by yangzheng 2019/11/09 录像下载异常

	/********************预览回调函数*********************/
	NET_DVR_SYSHEAD             = 1   //系统头数据
	NET_DVR_STREAMDATA          = 2   //视频流数据（包括复合流和音视频分开的视频流数据）
	NET_DVR_AUDIOSTREAMDATA     = 3   //音频流数据
	NET_DVR_STD_VIDEODATA       = 4   //标准视频流数据
	NET_DVR_STD_AUDIODATA       = 5   //标准音频流数据
	NET_DVR_SDP                 = 6   //SDP信息(Rstp传输时有效)
	NET_DVR_CHANGE_FORWARD      = 10  //码流改变为正放
	NET_DVR_CHANGE_REVERSE      = 11  //码流改变为倒放
	NET_DVR_PLAYBACK_ALLFILEEND = 12  //回放文件结束标记
	NET_DVR_VOD_DRAW_FRAME      = 13  //回放抽帧码流
	NET_DVR_VOD_DRAW_DATA       = 14  //拖动平滑码流
	NET_DVR_PRIVATE_DATA        = 112 //私有数据,包括智能信息

	//设备型号(DVR类型)
	/* 设备类型 */
	DVR              = 1   /*对尚未定义的dvr类型返回DVR*/
	ATMDVR           = 2   /*atm dvr*/
	DVS              = 3   /*DVS*/
	DEC              = 4   /* 6001D */
	ENC_DEC          = 5   /* 6001F */
	DVR_HC           = 6   /*8000HC*/
	DVR_HT           = 7   /*8000HT*/
	DVR_HF           = 8   /*8000HF*/
	DVR_HS           = 9   /* 8000HS DVR(no audio) */
	DVR_HTS          = 10  /* 8016HTS DVR(no audio) */
	DVR_HB           = 11  /* HB DVR(SATA HD) */
	DVR_HCS          = 12  /* 8000HCS DVR */
	DVS_A            = 13  /* 带ATA硬盘的DVS */
	DVR_HC_S         = 14  /* 8000HC-S */
	DVR_HT_S         = 15  /* 8000HT-S */
	DVR_HF_S         = 16  /* 8000HF-S */
	DVR_HS_S         = 17  /* 8000HS-S */
	ATMDVR_S         = 18  /* ATM-S */
	DVR_7000H        = 19  /*7000H系列*/
	DEC_MAT          = 20  /*多路解码器*/
	DVR_MOBILE       = 21  /* mobile DVR */
	DVR_HD_S         = 22  /* 8000HD-S */
	DVR_HD_SL        = 23  /* 8000HD-SL */
	DVR_HC_SL        = 24  /* 8000HC-SL */
	DVR_HS_ST        = 25  /* 8000HS_ST */
	DVS_HW           = 26  /* 6000HW */
	DS630X_D         = 27  /* 多路解码器 */
	DS640X_HD        = 28  /*640X高清解码器*/
	DS610X_D         = 29  /*610X解码器*/
	IPCAM            = 30  /*IP 摄像机*/
	MEGA_IPCAM       = 31  /*高清IP摄像机*/
	IPCAM_X62MF      = 32  /*862MF可以接入9000设备*/
	ITCCAM           = 35  /*智能高清网络摄像机*/
	IVS_IPCAM        = 36  /*智能分析高清网络摄像机*/
	ZOOMCAM          = 38  /*一体机*/
	IPDOME           = 40  /*IP 标清球机*/
	IPDOME_MEGA200   = 41  /*IP 200万高清球机*/
	IPDOME_MEGA130   = 42  /*IP 130万高清球机*/
	IPDOME_AI        = 43  /*IP 高清智能快球*/
	TII_IPCAM        = 44  /*红外热成像摄像机*/
	IPTC_DOME        = 45  /*红外热成像双目球机*/
	DS_2DP_Z         = 46  /*球型鹰眼（大）*/
	DS_2DP           = 47  /*非球型鹰眼（小）*/
	ITS_WMS          = 48  /*称重数据管理服务器*/
	IPMOD            = 50  /*IP 模块*/
	TRAFFIC_YTDOME   = 51  //交通智能云台（不带雷达测速）
	TRAFFIC_RDDOME   = 52  //交通智能云台（带雷达测速）
	IDS6501_HF_P     = 60  // 6501 车牌
	IDS6101_HF_A     = 61  //智能ATM
	IDS6002_HF_B     = 62  //双机跟踪：DS6002-HF/B
	IDS6101_HF_B     = 63  //行为分析：DS6101-HF/B DS6101-HF/B_SATA
	IDS52XX          = 64  //智能分析仪IVMS
	IDS90XX          = 65  // 9000智能
	IDS8104_AHL_S_HX = 66  // 海鑫人脸识别 ATM
	IDS8104_AHL_S_H  = 67  // 私有人脸识别 ATM
	IDS91XX          = 68  // 9100智能
	IIP_CAM_B        = 69  // 智能行为IP摄像机
	IIP_CAM_F        = 70  //智能人脸IP摄像机
	DS71XX_H         = 71  /* DS71XXH_S */
	DS72XX_H_S       = 72  /* DS72XXH_S */
	DS73XX_H_S       = 73  /* DS73XXH_S */
	DS72XX_HF_S      = 74  //DS72XX_HF_S
	DS73XX_HFI_S     = 75  //DS73XX_HFI_S
	DS76XX_H_S       = 76  /* DVR,e.g. DS7604_HI_S */
	DS76XX_N_S       = 77  /* NVR,e.g. DS7604_NI_S */
	DS_TP3200_EC     = 78  /*机柜智能检测仪*/
	DS81XX_HS_S      = 81  /* DS81XX_HS_S */
	DS81XX_HL_S      = 82  /* DS81XX_HL_S */
	DS81XX_HC_S      = 83  /* DS81XX_HC_S */
	DS81XX_HD_S      = 84  /* DS81XX_HD_S */
	DS81XX_HE_S      = 85  /* DS81XX_HE_S */
	DS81XX_HF_S      = 86  /* DS81XX_HF_S */
	DS81XX_AH_S      = 87  /* DS81XX_AH_S */
	DS81XX_AHF_S     = 88  /* DS81XX_AHF_S */
	DS90XX_HF_S      = 90  /*DS90XX_HF_S*/
	DS91XX_HF_S      = 91  /*DS91XX_HF_S*/
	DS91XX_HD_S      = 92  /*91XXHD-S(MD)*/
	IDS90XX_A        = 93  // 9000智能 ATM
	IDS91XX_A        = 94  // 9100智能 ATM
	DS95XX_N_S       = 95  /*DS95XX_N_S NVR 不带任何输出*/
	DS96XX_N_SH      = 96  /*DS96XX_N_SH NVR*/
	DS90XX_HF_SH     = 97  /*DS90XX_HF_SH */
	DS91XX_HF_SH     = 98  /*DS91XX_HF_SH */
	DS_B10_XY        = 100 /*视频综合平台设备型号(X:编码板片数，Y:解码板片数)*/
	DS_6504HF_B10    = 101 /*视频综合平台内部编码器*/
	DS_6504D_B10     = 102 /*视频综合平台内部解码器*/
	DS_1832_B10      = 103 /*视频综合平台内部码分器*/
	DS_6401HFH_B10   = 104 /*视频综合平台内部光纤板*/
	DS_65XXHC        = 105 //65XXHC DVS
	DS_65XXHC_S      = 106 //65XXHC-SATA DVS
	DS_65XXHF        = 107 //65XXHF DVS
	DS_65XXHF_S      = 108 //65XXHF-SATA DVS
	DS_6500HF_B      = 109 //65 rack DVS
	IVMS_6200_C      = 110 // iVMS-6200(/C)
	IVMS_6200_B      = 111 // iVMS-6200(/B)
	DS_72XXHV_ST15   = 112 //72XXHV_ST  海思3515平台 DVR
	DS_72XXHV_ST20   = 113 //72XXHV_ST  海思3520平台 DVR
	IVMS_6200_T      = 114 // IVMS-6200(/T)
	IVMS_6200_BP     = 115 // IVMS-6200(/BP)
	DS_81XXHC_ST     = 116 //DS_81XXHC_ST
	DS_81XXHS_ST     = 117 //DS_81XXHS_ST
	DS_81XXAH_ST     = 118 //DS_81XXAH_ST
	DS_81XXAHF_ST    = 119 //DS_81XXAHF_ST
	DS_66XXDVS       = 120 //66XX DVS

	DS_1964_B10          = 121 /*视频综合平台内部报警器*/
	DS_B10N04_IN         = 122 /*视频综合平台内部级联输入*/
	DS_B10N04_OUT        = 123 /*视频综合平台内部级联输出*/
	DS_B10N04_INTEL      = 124 /*视频综合平台内部智能*/
	DS_6408HFH_B10E_RM   = 125 //V6高清
	DS_B10N64F1_RTM      = 126 //V6级联不带DSP
	DS_B10N64F1D_RTM     = 127 //V6级联带DSP
	DS_B10_SDS           = 128 //视频综合平台子域控制器
	DS_B10_DS            = 129 //视频综合平台域控制器
	DS_6401HFH_B10V      = 130 //VGA高清编码器
	DS_6504D_B10B        = 131 /*视频综合平台内部标清解码器*/
	DS_6504D_B10H        = 132 /*视频综合平台内部高清解码器*/
	DS_6504D_B10V        = 133 /*视频综合平台内部VGA解码器*/
	DS_6408HFH_B10S      = 134 //视频综合平台SDI子板
	DS_18XX_N            = 135 /* 矩阵接入网关*/
	DS_6504HF_B10F_CLASS = 136 //光端机SD
	DS_18XX_PTZ          = 141 /*网络码分类产品*/
	DS_19AXX             = 142 /*通用报警主机类产品*/
	DS_19BXX             = 143 /*家用报警主机*/
	DS_19CXX             = 144 /*自助银行报警主机*/
	DS_19DXX             = 145 /*动环监控报警主机*/
	DS_19XX              = 146 /*1900系列报警主机*/
	DS_19SXX             = 147 /*视频报警主机*/
	DS_1HXX              = 148 /*CS类产品*/ //防护舱
	DS_PEAXX             = 149 /*一键式紧急报警产品*/
	DS_PWXX              = 150 /*无线报警主机产品*/
	DS_PMXX              = 151 /*4G网络模块*/
	DS_19DXX_S           = 152 /*视频动环监控主机*/
	DS_PWAXX             = 153 /* Axiom Hub无线报警主机 */
	DS_PHAXX             = 154 /* Axiom Hybrid混合报警主机 */

	//2011-11-30
	DS_C10H        = 161 /*多屏控制器*/
	DS_C10N_BI     = 162 //BNC处理器
	DS_C10N_DI     = 163 //rbg处理器
	DS_C10N_SI     = 164 //码流处理器
	DS_C10N_DO     = 165 //显示处理器
	DS_C10N_SERVER = 166 //分布式服务器

	IDS_8104_AHFL_S_H = 171 // 8104ATM
	IDS_65XX_HF_A     = 172 // 65 ATM
	IDS90XX_HF_RH     = 173 // 9000 智能RH
	IDS91XX_HF_RH     = 174 // 9100 智能RH设备
	IDS_65XX_HF_B     = 175 // 65 行为分析
	IDS_65XX_HF_P     = 176 // 65 车牌识别
	IVMS_6200_F       = 177 // IVMS-6200(/F)
	IVMS_6200_A       = 178 //iVMS-6200(/A)
	IVMS_6200_F_S     = 179 // IVMS-6200(/F_S)人脸后检索分析仪

	DS90XX_HF_RH = 181 // 9000 RH    648
	DS91XX_HF_RH = 182 // 9100 RH设备 648
	DS78XX_S     = 183 // 78系列设备 6446
	DS81XXHW_S   = 185 // 81 Resolution 960 KY2011
	DS81XXHW_ST  = 186 // DS81XXHW_ST  KY2011
	DS91XXHW_ST  = 187 // DS91XXHW_ST  KY2011
	DS91XX_ST    = 188 // DS91XX_ST netra
	DS81XX_ST    = 189 // DS81XX_ST netra
	DS81XXHX_ST  = 190 // DS81XXHDI_ST,DS81XXHE_ST ky2012
	DS73XXHX_ST  = 191 // DS73XXHI_ST ky2012
	DS81XX_SH    = 192 // 审讯81SH,81SHF
	DS81XX_SN    = 193 // 审讯81SNL

	DS96XXN_ST  = 194 //NVR:DS96xxN_ST
	DS86XXN_ST  = 195 //NVR:DS86xxN_ST
	DS80XXHF_ST = 196 //DVR:DS80xxHF_ST
	DS90XXHF_ST = 197 //DVR:DS90xxHF_ST
	DS76XXN_ST  = 198 //NVR:DS76xxN_ST

	DS_9664N_RX       = 199 //NVR:DS_9664N_RX
	ENCODER_SERVER    = 200 // 编码卡服务器
	DECODER_SERVER    = 201 // 解码卡服务器
	PCNVR_SERVER      = 202 // PCNVR存储服务器
	CVR_SERVER        = 203 // 邦诺CVR，他给自己定的类型为DVR_S-1
	DS_91XXHFH_ST     = 204 // 91系列HD-SDI高清DVR
	DS_66XXHFH        = 205 // 66高清编码器
	TRAFFIC_TS_SERVER = 210 //终端服务器
	TRAFFIC_VAR       = 211 //视频分析记录仪
	IPCALL            = 212 //IP可视对讲分机
	SAN_SERVER        = 213 //与CVR_SERVER相同的程序，只是模式不同

	DS_B11_M_CLASS        = 301 /*视频综合平台设备型号*/
	DS_B12_M_CLASS        = 302 /*视频综合平台设备型号*/
	DS_6504HF_B11_CLASS   = 303 /*视频综合平台内部编码器*/
	DS_6504HF_B12_CLASS   = 304 /*视频综合平台内部编码器*/
	DS_6401HFH_B11V_CLASS = 305 //VGA高清
	DS_6401HFH_B12V_CLASS = 306 //VGA高清
	DS_6408HFH_B11S_CLASS = 307 //SDI
	DS_6408HFH_B12S_CLASS = 308 //SDI
	DS_6504D_B11H_CLASS   = 309 /*视频综合平台内部高清解码器*/
	DS_6504D_B11B_CLASS   = 310 /*视频综合平台内部标清解码器*/
	DS_6504D_B12B_CLASS   = 311 /*视频综合平台内部标清解码器*/
	DS_6504D_B11V_CLASS   = 312 /*视频综合平台内部VGA解码器*/
	DS_6504D_B12V_CLASS   = 313 /*视频综合平台内部VGA解码器*/
	//B10新增
	DS_6401HFH_B10R_CLASS = 314 //B10 RGB高清
	DS_6401HFH_B10D_CLASS = 315 //B10 DVI高清
	DS_6401HFH_B10H_CLASS = 316 //B10 HDMI高清
	//B11新增
	DS_6401HFH_B11R_CLASS = 317 //B11 RGB高清
	DS_6401HFH_B11D_CLASS = 318 //B11 DVI高清
	DS_6401HFH_B11H_CLASS = 319 //B11 HDMI高清
	//B12新增
	DS_6401HFH_B12R_CLASS = 320 //B12 RGB高清
	DS_6401HFH_B12D_CLASS = 321 //B12 DVI高清
	DS_6401HFH_B12H_CLASS = 322 //B12 HDMI高清
	DS_65XXD_B10Ex_CLASS  = 323 //netra高清解码

	//B10 V2.1新增
	DS_6516HW_B10_CLASS      = 324 //netra高线编码
	DS_6401HFH_B10F_RX_CLASS = 326 //高清光端机接入（支持1/2路光端机接入）
	DS_6502HW_B10F_RX_CLASS  = 327 //960H光端机接入（支持1/4/8路光端机接入）
	//2012-5-16新增
	DS_6504D_B11Ex_CLASS = 328 //netra高清解码
	DS_6504D_B12Ex_CLASS = 329 //netra高清解码
	DS_6512_B11_CLASS    = 330 //netra高线编码
	DS_6512_B12_CLASS    = 331 //netra高线编码
	DS_6504D_B10H_CLASS  = 332 //视频综合平台内部高清解码器

	DS_65XXT_B10_CLASS       = 333 //视频综合平台转码子系统
	DS_65XXD_B10_CLASS       = 335 //视频综合平台万能解码板
	DS_IVMSE_B10X_CLASS      = 336 //X86服务器子系统
	DS_6532D_B10ES_CLASS     = 337 //增强型解码板_SDI(B10)
	DS_6508HFH_B10ES_CLASS   = 338 //SDI输入编码子系统
	DS_82NCG_CLASS           = 340 //联网网关中的子系统
	DS_82VAG_CLASS           = 341 //联网网关中的子系统
	DS_1802XXF_B10_CLASS     = 342 //光口交换子系统
	iDS_6504_B10EVAC_CLASS   = 343 //智能子系统
	iDS_6504_B10EDEC_CLASS   = 344 //智能子系统
	DS_6402HFH_B10EV_CLASS   = 345 //netra编码(VGA)
	DS_6402HFH_B10ED_CLASS   = 346 //netra编码(DVI)
	DS_6402HFH_B10EH_CLASS   = 347 //netra编码(HDMI)
	DS_6404HFH_B10T_RX_CLASS = 348 //光纤接入编码
	DS_6504D_AIO_CLASS       = 349 //netra高清解码
	DS_IVMST_B10_CLASS       = 350 //X86转码子系统
	DS_6402_AIO_CLASS        = 351 //netra编码
	DS_iVMSE_AIO_CLASS       = 352 //x86服务器子系统
	DS_AIO_M_CLASS           = 353 //一体机

	DS_6508HF_B10E_CLASS     = 355 //BNC输入编码子系统
	DS_6404HFH_B10ES_CLASS   = 356 //SDI输入编码子系统
	DS_6402HFH_B10ER_CLASS   = 358 //RGB输入编码子系统
	DS_6404HFH_B10T_RM_CLASS = 361 //光纤输入编码子系统
	DS_6516D_B10EB_CLASS     = 362 //BNC输出解码子系统
	DS_6516D_B10ES_CLASS     = 363 //SDI输出解码子系统

	//DVI/HDMI/VGA畅显解码公用一个类型
	DS_6508D_B10FH_CLASS = 364
	DS_6508D_B10FD_CLASS = 364
	DS_6508D_B10FV_CLASS = 364

	DS_6508_B11E_CLASS   = 365 //BNC输入编码子系统
	DS_6402_B11ES_CLASS  = 366 //SDI输入编码子系统
	DS_6402_B11EV_CLASS  = 367 //VGA输入编码子系统
	DS_6402_B11ER_CLASS  = 368 //RGB输入编码子系统
	DS_6402_B11ED_CLASS  = 369 //DVI输入编码子系统
	DS_6402_B11EH_CLASS  = 370 //HDMI输入编码子系统
	DS_6516D_B11EB_CLASS = 371 //BNC输出解码子系统
	DS_6516D_B11ES_CLASS = 372 //SDI输出解码子系统

	DS_6508_B12E_CLASS   = 373 //BNC输入编码子系统
	DS_6402_B12ES_CLASS  = 375 //SDI输入编码子系统
	DS_6402_B12EV_CLASS  = 376 //VGA输入编码子系统
	DS_6402_B12ER_CLASS  = 377 //RGB输入编码子系统
	DS_6402_B12ED_CLASS  = 378 //DVI输入编码子系统
	DS_6402_B12EH_CLASS  = 379 //HDMI输入编码子系统
	DS_6516D_B12EB_CLASS = 380 //BNC输出解码子系统

	DS_iVMSE_AIO_8100x_CLASS = 381 //金融行业一体机X86子系统
	DS_iVMSE_AIO_87x_CLASS   = 382 //智能楼宇一体机X86子系统
	DS_6532D_B11ES_CLASS     = 384 //增强型解码板_SDI(B11)
	DS_6532D_B12ES_CLASS     = 385 //增强型解码板_SDI(B12)
	//B20新增
	DS_B20_MSU_NP              = 400 //B20主控板
	DS_6416HFH_B20S            = 401 //SDI输入编码
	DS_6416HFH_B20_RM          = 402 //光纤输入高清编码
	DS_6564D_B20D              = 403 //DVI解码
	DS_6564D_B20H              = 404 //HDMI解码
	DS_6564D_B20V              = 405 //VGA解码
	DS_B20_6516D_DEV_CLASS     = 406 //B20解码子系统
	DS_6408HFH_B20V            = 407 //VGA编码板
	DS_MMC_B20_CLASS           = 408 //B20主控
	DS_CARD_CHIP_B20_CLASS     = 409 //B20主控子板
	DS_6564D_B20B_DEV_CLASS    = 410 //BNC解码子系统
	DS_6564D_B20S_DEV_CLASS    = 411 //SDI解码子系统
	DS_6532HF_B20B_DEV_CLASS   = 412 //BNC编码子系统
	DS_6408HFH_B20D_DEV_CLASS  = 413 //DVI编码子系统
	DS_6408HFH_B20H_DEV_CLASS  = 414 //HDMI编码子系统
	DS_IVMSE_B20_CLASS         = 415 //X86服务器子系统
	DS_6402HFH_B20Y_DEV_CLASS  = 416 //YUV编码子系统
	DS_6508HW_B20_DEV_CLASS    = 417 //HW编码子系统
	DS_B20N128Fx_B20_DEV_CLASS = 418 //DS_B20N128Fx_M级联板
	DS_AIO_MCU_NP_DEV_CLASS    = 419 //IO主控板
	DS_6402_AIO_EV_DEV_CLASS   = 420 //VGA编码
	DS_6508D_AIO_EV_DEV_CLASS  = 421 //VGA解码
	DS_6508D_AIO_ED_DEV_CLASS  = 422 //DVI解码
	DS_6508D_AIO_EH_DEV_CLASS  = 423 //HDMI解码
	DS_6508HD_B20F_DEV_CLASS   = 424 //视频增强板
	DS_6402HFH_B20ES_DEV_CLASS = 425 //3G SID编码
	DS_6532D_B20_DEV_CLASS     = 426 //B20解码子系统
	DS_IVMST_B20_DEV_CLASS     = 427 //X86服务器子系统
	DS_6416HFH_B20DD_DEV_CLASS = 428 //DVI双链路
	DS_6441VS_B20_DEV_CLASS    = 429 //相机拼接类型
	DS_6404HFH_B20T_CLASS      = 431 //TVI
	DS_FS22_B20_DEV_CLASS      = 432 //交换机子系统
	DS_IVMSE_B20UH_DEV_CLASS   = 433 //超高清X86输入
	IDS_6524J_B20_DEV_CLASS    = 434
	IDS_6532B_B20_DEV_CLASS    = 435
	DS_6404HFH_B20Fx_DEV_CLASS = 436 //光端机接入子系统
	DS_N128x_B20Fy_CLASS       = 437 //级联子系统
	DS_181600F_B20_CLASS       = 438 //网络光纤子系统
	DS_6904UD_B20H_CLASS       = 439 //超高清解码子系统

	DS_B21_MCU_NP_CLASS   = 440 //B21主控
	DS_B21_S10_x_CLASS    = 441 //B21机箱 x = A/S/D
	DS_6402HFH_B21D_CLASS = 442 //B21编码子系统
	DS_6508HD_B21D_CLASS  = 443 //B21解码子系统
	DS_iVMSE_B21HW_CLASS  = 444 //B21 X86子系统

	DS_C10S = 501 //C10S 集中式大屏控制器

	DS_C10N_SDI       = 551 //SDI处理器
	DS_C10N_BIW       = 552 //8路BNC处理器
	DS_C10N_DON       = 553 //显示处理器
	DS_C10N_TVI       = 554 //TVI输入板
	DS_C10N_DI2       = 555 //DVI 2路输入板
	DS_C10N_AUDIO_OUT = 556 //DVI，HDMI，VGA输出板带音频
	DS_C10N_AUDIO_IN  = 557 //DVI，HDMI，VGA输入板带音频

	//软服务器版本的分布式大屏控制器
	DS_C20N     = 570 //分布式大屏控制器
	DS_C20N_BNC = 571 //BNC输入设备
	DS_C20N_DVI = 572 //DVI输入设备
	DS_C20N_DP  = 573 //DP输入设备
	DS_C20N_OUT = 574 //输出设备

	//硬服务器版本的分布式大屏控制器
	DS_C20N_VWMS = 5351 //C20N服务器
	DS_C20N_DI2  = 5352 //两路DVI输入设备
	DS_C20N_DO2  = 5353 //两路DVI输出设备
	DS_C20N_UHDI = 5354 //DP/HDMI 1路4K超高清输入设备
	DS_C20N_YI2  = 5355 //两路YPbPr输入板

	DS_C12N_VWMS = 5356 //C12N服务器
	DS_C12N_DO   = 5357 //1路DVI输出
	DS_C12N_HOE  = 5358 //1路HDMI 4K输出

	DS_C20N_HI2   = 5359 // HDMI输入节点(2路)
	DS_C20N_VI2   = 5360 //VGA输入节点(2路)
	DS_C20N_SI2   = 5361 //SDI输入节点(2路)
	DS_C20N_HO2   = 5362 //HDMI输出节点(2路)
	DS_C20N_S24G  = 5363 //C20N专用千兆交换机(24路)
	DS_C20N_S24X  = 5364 //C20N专用万兆交换机(24路)
	DS_C12A_0104H = 5365 //创意拼接控制器

	//全息显示设备
	DS_D1HXX = 5591 //全息显示设备

	//无线传输设备
	ELEVATO_BRIDGE  = 5751 //无线电梯网桥
	DS_3WF01S_5NG_M = 5752 //5.8G室外1公里无线工地网
	DS_3WF0EC_2N_D  = 5753 //2.4G室外500米网桥
	DS_3WF0EC_5ACT  = 5754 //5.8G 11ac电梯网桥
	DS_3WF03S_5AC   = 5755 //5.8G 3公里经济型网桥
	DS_3WF0ES_5AC_H = 5756 //5.8G室外500米高穿透网桥
	DS_3WF05S_5AC_H = 5765 //5.8G 5KM千兆网口拨码网桥
	DS_3WSXXX       = 5766 //网关路由器系列（包括AC控制器）
	DS_3WAXXX       = 5767 //AP系列
	DS_3WRXXX       = 5768 //普通无线路由器系列
	DS_3WMRXXX      = 5769 //MESH路由器系列

	//报警设备
	DS_19M00_MN  = 601 //报警主机百兆网络模块
	DS_KH8302_A  = 602 //室内机
	DS_KD8101_2  = 603 //门口机
	DS_KM8301    = 604 //管理机
	DS_KVXXXX_XX = 605 //别墅门口机
	DS_LHPM1D_C  = 606 //酒店锁
	DS_KAM03     = 607 //半数字转接模块

	DS64XXHD_T  = 701  //64-T高清解码器
	DS_65XXD    = 703  //65万能解码器
	DS63XXD_T   = 704  //63-T标清解码器
	SCE_SERVER  = 705  //抓屏服务器
	DS_64XXHD_S = 706  //64XXHD-S高清解码器
	DS_68XXT    = 707  //多功能视音频转码器
	DS_65XXD_T  = 708  //65D-T万能解码器
	DS_65XXHD_T = 709  //65HD-T万能解码器
	DS_69XXUD   = 710  //69XXUD超高清解码器
	DS_65XXUD   = 711  //65XXUD解码器
	DS_65XXUD_L = 712  //65XXUD_L解码器
	DS_65XXUD_T = 713  //65XXUD_T解码器
	DS_69XXHD   = 5001 //69XXHD解码器

	DS_D20XX = 750 //LCD屏幕 解码卡
	//SDI矩阵
	DS_C50S  = 751 //SDI矩阵
	DS_D50XX = 752 //LCD屏幕 解码卡

	DS_D40 = 760 //LED屏发送卡

	DS_65VMXX       = 770  //视频会议服务器
	DS_65VTXX       = 771  //视频会议终端
	DS_65VTA        = 772  //视频会议一体式终端
	DS_65VT_RX      = 773  //互动教学终端
	DS_65VM_MCU_NP  = 774  //大容量MCU
	DS_65VT0010     = 5571 //一体式终端
	DS_65VM_MCU     = 5572 //高密度MCU主控板
	DS_65VM_MPC     = 5573 //高密度MCU媒体资源板
	DS_65VT2XXYR_SF = 5574 //司法提讯终端
	DS_65VT0XXY     = 5575 //视频会议大终端
	DS_65VT0010Z    = 5576 //智能一体式终端
	DS_65VT0050Z    = 5577 //智能分体式终端
	DS_65VS0XXXS    = 5878 //opensips服务器设备
	DS_65VS0XXXM    = 5579 //视频会议会控平台设备
	DS_65VM00XX_X   = 5580 //视频会议X86系统MCU
	DS_65VA800_BOX  = 5581 //视频会议麦克风

	DS_CS_SERVER  = 800 //虚拟屏服务器
	DS_68GAP_MCU  = 831 //视频网闸
	DS_K260X      = 850 //门禁主机
	DS_K1TXXX     = 851 //指纹一体机
	DS_K280X      = 852 //经济型门禁主机
	DS_K1T80X     = 854 //经济型门禁一体机
	RAC_6X00      = 856 //汉军指纹门禁一体机
	DS_K2602_AX   = 857 //人员通道主机
	DS_K1T803F    = 858 //经济型指纹门禁产品
	DS_K2700      = 859 //分布式三层架构门禁主机
	DS_K270X      = 860 //分布式三层架构就地控制器
	DS_K1T500S    = 861 //视屏门禁一体机
	DS_K1A801F    = 862 //经济型指纹门禁产品
	DS_K1T600X    = 863 //人脸识别门禁一体机
	DS_K22X       = 864 //梯控主控制器
	DS_K2M0016AX  = 865 //梯控分控制器
	DS_K2602S_AX  = 867 //人员通道产品
	DS_K560XX     = 870 //立式/台式智能身份识别终端
	DS_K260X_E    = 873 //低端门禁主机
	DS_K5603X     = 874 //台式/壁挂式嵌入式人证设备
	DS_K3M200X    = 875 //权限控制器
	DS_K3M100X    = 879 //通道控制器
	DS_K1T8101XT  = 881 //经销型人脸识别门禁一体机
	DS_K1T8101XX  = 882 //经销型人脸识别门禁一体机
	DS_K5604      = 883 //5604人证设备
	DS_K1T607XXXX = 884 //中端人脸门禁一体机DS_K1T607（M/MF/E/EF/MG/MFG?MW/MFW/TM/TMF/TE/TEF/TMG/TMFG/TMW/TMFW）
	DS_K1A850X    = 885 //低端经销指纹考勤机
	DS_K1T800X    = 886 //低端经销指纹门禁一体机
	DS_K1T610XXXX = 887 //中端人脸门禁一体机DS_K1T610（M/MF/E/EF/MG/MFG?MW/MFW/TM/TMF/TE/TEF/TMG/TMFG/TMW/TMFW）
	DS_K1T8115X   = 888 //DS-K1T8115、DS-K1T8115M、DS-K1T8115M-Z 经销型人脸识别一体机
	DS_K1T815LC_M = 889 //经销型人脸识别一体机
	DS_K1T606M_Z  = 890 //经销型人脸识别一体机
	DS_K5607_XXX  = 891 //K5607轻薄款通道人脸组件
	DS_K1T950MX   = 893 //经销款人脸指纹门禁一体机
	DS_K1T331XX   = 894 //低端人脸门禁考勤一体机
	DS_K1T671T    = 895 //人脸门禁一体机
	DS_K1T671     = 896 //人脸门禁一体机
	DS_K5671      = 897 //轻薄款通道人脸组件
	DS_K1T640     = 898 //4.3寸人脸门禁一体机
	DS_K1A802A    = 899 //经济型指纹考勤机

	DS_6800M           = 900 //68M合码器
	DS_68NTH           = 901 //信息发布主机
	DS_D60S            = 902 //信息发布服务器
	DS_D60W            = 903 //信息发布Windows终端
	DS_D10             = 931 //背投显示设备
	DS_3K0X_NM         = 951 //光纤收发器
	DS_3E2328          = 952 //百兆交换机
	DS_3E1528          = 953 //千兆交换机
	SCREEN_LINK_SERVER = 971 //屏幕服务器
	DS_D51OPSXX        = 972 //OPS电脑盒
	//一体化云台
	IP_PTSYS_MEGA200 = 1001 //IP 200万一体化云台
	IPCAM_FISHEYE    = 1002 //鱼眼摄像机
	IPCAM_FISHEYE_E  = 1003 //经济型鱼眼摄像机

	//68xx系列双目3D相机
	IPCAM_BINOCULAR = 1004 //双目摄像机

	IPCAM_365 = 1010 //支持365的平台的IPC CAM
	IPCAM_R0  = 1011 //支持A5S的平台的IPC CAM
	IPCAM_R1  = 1012 //支持385的平台的IPC CAM
	IPCAM_R2  = 1013 //支持R2的平台的IPC CAM
	IPCAM_R3  = 1014 //支持8127的平台的IPC CAM
	IPCAM_R4  = 1015 //支持S2的平台的IPC CAM

	IPDOME_365 = 1110 //支持365的平台的IPD CAM
	IPDOME_R0  = 1111 //支持A5S的平台的IPD CAM
	IPDOME_R1  = 1112 //支持385的平台的IPD CAM
	IPDOME_R2  = 1113 //支持R2的平台的IPD CAM
	IPDOME_R3  = 1114 //支持8127的平台的IPD CAM
	IPDOME_R4  = 1115 //支持S2的平台的IPD CAM
	ITCCAM_R3  = 1211 //支持8127的平台的ITCCAM

	//无人机业务设备（1300~1350）
	UAV_S = 1300 //无人机基站设备 （Ummanned Aerial Vehicle – Station）

	//新增设备类型 2013-11-19
	TRAFFIC_ECT            = 1400 //ECT设备类型
	TRAFFIC_PARKING_SERVER = 1401 //停车场服务器
	TRAFFIC_TME            = 1402 //出入口控制机
	// DVR
	DS90XXHW_ST   = 2001 // DS90XXHW_ST混合DVR
	DS72XXHX_SH   = 2002 // DS-72xxHV_SH, DS-72xxHF-SH
	DS_92XX_HF_ST = 2003 // DS-92XX-HF-ST
	DS_91XX_HF_XT = 2004 // 9100DVR_HF_XT
	DS_90XX_HF_XT = 2005 // 9000DVR_HF_XT
	DS_73XXHX_SH  = 2006 // 7300DVR_HX_SH
	DS_72XXHFH_ST = 2007 // 7200DVR_HFH_ST
	//DS_67系列

	DS_67XXHF_SATA = 2008 // DS-67XXHF-SATA
	DS_67XXHW      = 2009 // DS-67XXHW
	DS_67XXHW_SATA = 2010 // DS-67XXHW-SATA
	DS_67XXHF      = 2011 // DS-67XXHF

	//Netra2.3
	DS_72XXHF_SV = 2012 // DS-72xxHF-SV
	DS_72XXHW_SV = 2013 // DS-72xxHW-SV

	DS_81XXHX_SH = 2014 // 8100DVR_HX_SH

	DS_71XXHX_SL = 2015 //小型DVR

	DS_76XXH_ST = 2016 //DS_76XXH_ST

	DS_73XXHFH_ST = 2017 //73HFH系列
	DS_81XXHFH_ST = 2018 //81HFH系列 ST

	DS_72XXHFH_SL  = 2019 //hi3521
	DS_FDXXCGA_FLT = 2020 //2盘位ATM

	IDS_91XX_HF_ST_A = 2100 //iDS-9100HF-ST/A
	IDS_91XX_HF_ST_B = 2101 //iDS-9100HF-ST/B
	IDS_90XX_HF_ST_A = 2102 //iDS-9000HF-ST/A
	IDS_90XX_HF_ST_B = 2103 //iDS-9000HF-ST/B
	IDS_81XX_HF_ST_A = 2104 //iDS-8100HF-ST/A
	IDS_81XX_HF_ST_B = 2105 //iDS-8100HF-ST/B
	IDS_80XX_HF_ST_A = 2106 //iDS-8000HF-ST/A
	IDS_80XX_HF_ST_B = 2107 //iDS-8000HF-ST/B
	IDS_8104_AHFL_ST = 2108 //智能混合ATM机
	IDS_2CD6812F_C   = 2109 //垂直双目相机

	// NVR
	DS_77XXN_ST  = 2201 //  NVR DS-77XXHF-ST
	DS_95XX_N_ST = 2202 //  95XXN_ST NVR
	DS_85XX_N_ST = 2203 //  85XXN_ST NVR
	DS_96XX_N_XT = 2204 // 9600NVR_N_XT
	DS_76XX_N_SE = 2205 // 7600NVR_N_SE

	//高清审讯机
	DS_86XXSN_SX = 2206 // 8608NVR_SX，包括4中类型DS-8608SNL-SP、DS-8608SNL-ST、DS-8608SN-SP、DS-8608SN-ST，L表示带LCD，P表POE

	//    DS_96XX_N_RX                2207  //DS-96XX-N-RX
	DS_71XXN_SL = 2208 //DS-71XXN-SL 民用产品
	CS_N1_1XX   = 2209 //CS_N1_1XX，民用事业部所用

	DS_71XXN_SN   = 2210 //71XX_N_SN  经济型民用产品
	CS_N1_2XX     = 2211 //N1_2XX        民用事业部所用
	DS_76XX_N_SHT = 2212 //76XX_N_SHT  后端基线产品
	DS_96XXX_N_E  = 2213 //高新性能NVR(256)

	DS_76XXN_EX      = 2214 /* 76 78系列NVR，注：包括 4 8 16路的E1一盘位， 8 16 32路 E2两盘位； /N /P设备*/
	DS_77XXN_E4      = 2215 /* 77系列NVR，   注：包括8 16 32路， /N /P设备*/
	DS_86XXN_E8      = 2216 /* 86系列NVR，   注：包括8 16 32路， /N /P设备*/
	DS_9616N_H8      = 2217 //DS_9616N_H8
	DS_72XXHX_EX_GXY = 2218 //表示72系列无线DVR产品，其中72后面的xx表示通道数，H后的X目前仅有HW,为了后续扩展保留，E后面的X表示的是盘位数，G XY表示的是无线类型
	DS_76XXN_EX_GXY  = 2219 //表示是76系列无线NVR产品，其中76后面的xx表示通道数，E后面的X表示的是盘位数，G XY表示的是无线类型
	DS_72XXHXH_SH_21 = 2220 //  72XXHXH_SH_21
	DS_72XXHXH_SH_31 = 2221 //  72XXHXH_SH_31
	DS_73XXHXH_SH_31 = 2222 // 73XXHXH_SH_31
	DS_81XXHXH_SH_31 = 2223 //  81XXHXH_SH_31
	DS_71XXHXH_SH_21 = 2224 // 71XXHXH_SH_21
	DS_71XXHXH_SH_31 = 2225 // 71XXHXH_SH_31
	DS_NBXX_N_E1_X   = 2226 //便携式主机
	DS_96XXN_FX      = 2230
	DS_86XXN_FX      = 2231
	DS_96XXXN_HX     = 2232 //96系列高性能设备

	DS_86XXN_I   = 2233 //DS_86XXN_I
	DS_77XX_N_I  = 2234 //DS_77XX_N_I
	DS_76XX_N_I  = 2235 //DS_76XX_N_I
	DS_78XX_N_I  = 2236 //DS_78XX_N_I
	DS_96XXXN_IX = 2237 //DS-96XXX_N_I（DS-96128N-I16、DS-96128N-I24、DS-96256N-I16、DS-96256N-I24）

	DS_90XXHXH_XT = 2238 //DS_90XXHXH_XT（DS-9016HQH-XT）

	//PCNVR
	PCNVR_IVMS_4200 = 2301 // PCNVR_IVMS_4200

	//智能分析仪
	IVMS_6200_TP    = 2401 //IVMS-6200 交通诱导分析仪
	IVMS_6200_TF    = 2402 //IVMS-6200 交通取证分析仪
	IVMS_6200_D     = 2403 //iVMS-6200(/D)
	IDS_81XXAHW_ST  = 2405 //iDS-81xxAHW-ST
	IDS_81XXAHW_SP  = 2406 //iDS-81xxAHW-SP
	IDS_81XXAHWL_ST = 2407 //iDS-81xxAHWL-ST
	IDS_81XXAHWL_SP = 2408 //iDS-81xxAHWL-SP
	IDS_9616N_H8    = 2409 //iDS_9616N_H8
	IVMS_6200_SMD   = 2500 //IVMS_6200_SMD
	//HISI3531平台混合ATM DVR， 其中L表示带液晶屏，P表示带poe
	DS_81XXAHW_ST  = 2501
	DS_81XXAHW_SP  = 2502
	DS_81XXAHWL_ST = 2503
	DS_81XXAHWL_SP = 2504

	//TVI ATM
	DS_81XXAHGH_E4 = 2601 //DS_81XXAHGH_E4

	/************************************************************************/
	/* 传输与显示产品设备类型start（不包含老的设备类型）                    */
	/************************************************************************/

	/** B20系列（4001-4900） */

	DS_6904UD_AIOE_H_CLASS = 4002 //B20一体机
	DS_6402HFH_B21B_CLASS  = 4005 //B21 编码
	DS_6902UD_B21H_CLASS   = 4006 //B21 HDMI解码
	DS_6902UD_B21D_CLASS   = 4007 //B21 DVI解码
	DS_6902UD_B21V_CLASS   = 4008 //B21 VGA解码
	DS_6904UD_B20D_CLASS   = 4010 //B20 DVI解码
	DS_6904UD_B20V_CLASS   = 4011 //B20 VGA解码

	//B21一体机设备型号
	DS_AIOH_MCU_NP    = 4042 //主控板
	DS_6404HFH_AIOH_D = 4043 //4路DVI输入板
	DS_6908UD_AIOH_D  = 4044 //4路DVI输出板
	DS_6908UD_AIOH_H  = 4045 //4路HDMI输出板

	//解码一体机设备型号
	DS_69XXUD_B  = 4046 //解码一体机
	DS_6404HFH_I = 4047 //4路DVI输入板
	DS_6904UD_O  = 4048 //8路HDMI输出板

	//集中式拼控器
	DS_C12L_0204H = 4060 //经济型小型拼控器

	//转码器    5401 - 5450  （50）
	DS_68VTG = 5401 //综合转码网关

	//视频云模方（Y10系列）
	DS_Y10_MCU        = 5771 //主控板
	DS_Y10_SW1        = 5772 //交换板
	DS_6516UD_Y10D    = 5773 //8路DVI显示接口解码板
	DS_6532UD_Y10D    = 5774 //16路DVI显示接口解码板
	DS_6E2724_Y10     = 5775 //24个万兆接口编码接入板
	DS_68SAA_AA_Y10   = 5776 //双路2620、内存32G业务板
	DS_68SBA_AA_Y10   = 5777 //双路2630、内存32G业务板
	DS_68SCA_AA_Y10   = 5778 //双路2650、内存32G业务板
	DS_68GPU_A1_Y10   = 5779 //P4显卡版
	DS_CPU_SERVER     = 5780 //CPU服务板
	DS_GPU_SERVER     = 5781 //GPU服务板
	DS_BIGDATA_SERVER = 5782 //大数据服务板

	//安监一体机系列
	DS_B80_MCU   = 5821 //主控板
	DS_B80_SW    = 5822 //交换板
	DS_B80_SR_01 = 5823 //服务板
	DS_B80_AI04  = 5824 //智能分析板，支持4路分析能力
	DS_B80_ST    = 5825 //存储板，支持4个2.5寸2T硬盘
	DS_B80_D04   = 5826 //数据采集业务板
	DS_B80_BNC08 = 5827 //标清编码板
	DS_B80_SDI08 = 5828 //高清编码板
	DS_B80_VP    = 5829 //转码/合码板
	DS_B80_VO02  = 5830 //视频输出板

	//网关类
	DS_3LGCX    = 5841 //通用物联网关
	DS_3LGRX    = 5842 //LoRa网关
	DS_3LGT4    = 5843 //电梯网关设备
	DS_3LGT40_N = 5844 //NB-IoT电梯网关
	DS_3LGS_I   = 5845 //智能网关

	//交换机
	DS_3E11XX  = 5681 //百兆交换机
	DS_3E13XX  = 5682 //千兆上行交换机
	DS_3E15XX  = 5683 //千兆交换机
	DS_3E11XXP = 5684 //百兆POE交换机
	DS_3E13XXP = 5685 //千兆上行POE交换机
	DS_3E15XXP = 5686 //千兆POE交换机
	DS_3E1520U = 5687 //USB充电柜
	DS_3EODM_H = 5688 //华三交换机
	DS_3EODM_W = 5689 //恒茂交换机
	DS_3T1XXX  = 5690 //DS-3T1XXX系列交换机

	//光纤收发器
	DS_3D01R_NM = 5691 //收发器接收卡
	DS_3D01T_NM = 5692 //收发器发送机
	DS_3K02_RNM = 5693 //收发器网管卡

	/************************************************************************/
	/* 传输与显示产品设备类型end（不包含老的设备类型）                      */
	/************************************************************************/

	IDS_90XX_HXH_XX_S  = 6001 //超脑人体DVR产品
	IDS_90XX_HXH_XX_FA = 6002 //超脑人脸DVR产品

	DS_81XX_SHXL_K4 = 6101 //KY2017平台审讯机DS-8104SHFH（L）-K4/4P

	DS_8116THFHL_F4 = 6201 //标准庭审主机DS-8116THFHL-F4

	DS_81XXAHQ_E4  = 6301 //DS_81XXAHQ_E4(TVI ATM)
	IDS_81XXAHQ_E4 = 6302 //IDS_81XXAHQ_E4(智能TVI ATM)

	IDS_86XX_NX_A    = 7501 //超脑智能ATM NVR(iDS-8632NX-I8/A)
	IDS_96XX_NX_S    = 7502 //人体超脑智能NVR
	IDS_96XX_NX_V    = 7503 //超脑加油站NVR
	IDS_96XX_NX_FA   = 7504 //人脸超脑智能NVR iDS-9632NX-I8/FA
	IDS_86XX_NX_IX_B = 7505 //86系列安全帽检测NVR产品
	IDS_67XX_NX_S_X  = 7506 //人脸超脑智能安检产品NVR(IDS_67XX_NX/S_X)
	NP_ST204_X_      = 7507 //第二代智能安检分析仪（NP_ST204_S:NP_ST204_D:NP_ST204_D_4G）

	IDS_ECDXX_HE = 8001 //学生三目相机
	IDS_ECDXX_HT = 8002 //教师三目相机

	IDS_ECDXX_E = 8101 //4K半球

	IDS_EPTXX_HF = 8201 //二代人脸抓拍机

	DS_2CD69XXX_I = 8351 //3200W全景拼接IPC型号
	DS_TRIXX      = 8471 //超高频固定式读卡器DS_TRI900

	DS_K1F600_D6EXX = 10501 //多功能采集仪
	DS_K1T341       = 10502 //经销人脸门禁
	DS_K1T641XXX    = 10503 //中端通用人脸门禁
	DS_K1T642XXX    = 10504 //中端通用人脸门禁
	DS_K1T601       = 10505 //酒店人脸产品
	DS_K1T672XXX    = 10506 //室内7寸超薄款人脸门禁产品
	DS_K56A0X       = 10507 //安卓中端人证
	DS_K56Q_F70     = 10508 //低端人脸组件（经销型)
	DS_K1T6Q_F70M   = 10509 //F70系列人脸识别一体机（经销型)
	DS_K1T6Q_F40X   = 10510 //F40系列人脸识别一体机（经销型)
	DS_K5604A_XXX   = 10511 //中低端人脸组件
	DS_K1A330X      = 10512 //超低端人脸考勤机（经销型)
	DS_K1T804A      = 10513 //经济型指纹门禁一体机
	DS_K1T202       = 10514 //经济型指纹门禁一体机
	DS_K1T343MX     = 10515 //人脸门禁一体机（经销型)
	DS_K1T331W_D    = 10516 //低端人脸门禁考勤一体机（DS-K1T331W(D)）
	DS_K1T671WX_D   = 10517 //人脸门禁一体机（DS-K1T671M(D) DS-K1T671MW(D) DS-K1T671MG(D)）
	DS_K1T680X      = 10518 //8.0寸室内外人脸门禁（DS-K1T680M DS-K1T680D）

	DS_IEXX_E_J = 11501 //监所智能服务器

	IDS_67XX_NX_A = 12501 //67系列NVS产品(iDS-6704NX/A)
	IDS_67XX_NX_L = 12502 //67系列录播NVS产品
	IDS_ENIXX_XHE = 12503 //录播NVS行业专业产品
	IDS_67XX_NX_V = 12504 //超脑加油站NVS
	IDS_67XX_NX_B = 12505 //67系列安全帽检测NVS产品

	//智能中心类设备
	DS_IE63XX_E_FA = 13001 //脸谱单机
	DS_DS_GPKIA    = 13002 //猎鹰服务器
	DS_DS_PURE     = 13003 //脸谱纯分析
	DS_DS_FS       = 13004 //人脸静态数据服务器
	DS_DS_FD       = 13005 //抓拍检测服务器
	DS_DS_BLADE    = 13006 //刀锋
	DS_DS_HMCP     = 13007 //模型对比服务器

	//智能锁设备
	DS_LNX_RF = 13501 //智能锁盒子

	//雷达设备    13551-14000（500）
	DS_PA_RADAR        = 13551 //PA雷达
	DS_PERIMETER_RADAR = 13552 //周界雷达
	DS_SECURITY_RADAR  = 13553 //120米安防雷达
	//消防设备    14001-14500（500）
	DS_N1104X = 14001 //消防网关
	DS_N1103X = 14002 //用户信息传输设备
	NP_FSC201 = 14003 //用水设备
	NP_FDC240 = 14004 //组合式电气火灾探测器
	DS_N1107  = 14005 //物联网网关
	NP_FAXXX  = 14006 //消防分析仪

	//安检设备 14501-15000（500）
	NP_ST204_X   = 14501 //第二代智能安检分析仪
	ISD_SG2XXL_X = 14502 //安检门(ISD-SG206L ISD-SG218L ISD-SG218L-F)
	/**********************设备类型 end***********************/
	/**********************设备大类 begin**********************/

	/* dvr相关 1-50 */
	DEV_CLASS_DVR           = 1 //普通dvr类型
	DEV_CLASS_INTERROGATION = 2 //审讯机
	DEV_CLASS_SIMPLE_TRAIL  = 3 //简易庭审主机
	DEV_CLASS_TRAIL         = 4 //标准庭审主机
	DEV_CLASS_RECORD_PLAY   = 5 //录播主机
	DEV_CLASS_ATM           = 6 //ATM机

	/* dvs相关 51-100 */
	DEV_CLASS_DVS = 51 //普通dvs

	/* nvr相关 101-150 */
	DEV_CLASS_NVR = 101 //普通nvr

	/* ipc相关 151-200 */
	DEV_CLASS_GUN          = 151 //ipc枪机
	DEV_CLASS_BALL         = 152 //ipc球机
	DEV_CLASS_SNAP         = 153 //抓拍机
	DEV_CLASS_INTELLI_TILT = 154 //智能云台
	DEV_CLASS_FISH_EYE     = 155 //鱼眼
	DEV_CLASS_2DP_Z        = 156 //大鹰眼
	DEV_CLASS_2DP          = 157 //小鹰眼
	DEV_CLASS_PT           = 158 //全景细节相机
	DEV_CLASS_TRI          = 159 //超高频固定式读卡器

	/* CVR相关 201 - 250*/
	DEV_CLASS_CVR = 201 //CVR

	/* 传显相关 251 - 300*/
	DEV_CLASS_B20                           = 251 //传显B20系列
	DEV_CLASS_B10                           = 252 //传显B10系列
	DEV_CLASS_DECODER                       = 253 //解码器
	DEV_CLASS_MATRIXMANAGEDEVICE            = 254 //矩阵接入网关
	DEV_CLASS_OTICAL                        = 255 //光端机
	DEV_CLASS_CODESPITTER                   = 256 //码分器
	DEV_CLASS_ALARMHOST                     = 257 //行业报警主机
	DEV_CLASS_MOVING_RING                   = 258 //动环设备
	DEV_CLASS_CVCS                          = 259 //集中式多屏控制器
	DEV_CLASS_DVCS                          = 260 //分布式多屏控制器
	DEV_CLASS_TRANSCODER                    = 261 //转码器
	DEV_CLASS_LCD_SCREEN                    = 262 //LCD屏幕
	DEV_CLASS_LED_SCREEN                    = 263 //LED屏幕
	DEV_CLASS_MATRIX                        = 264 //矩阵
	DEV_CLASS_CONFERENCE_SYSTEM             = 265 //视频会议设备
	DEV_CLASS_INFORMATION_RELEASE_EQUIPMENT = 266 //信息发布设备
	DEV_CLASS_NET_GAP                       = 267 //网闸
	DEV_CLASS_MERGE                         = 268 //合码器
	DEV_CLASS_REAR_PROJECTION               = 269 //背投显示设备
	DEV_CLASS_SWITCH                        = 270 //交换机
	DEV_CLASS_FIBER_CONVERTER               = 271 //光纤收发器
	DEV_CLASS_SCREEN_SERVER                 = 272 //屏幕服务器
	DEV_CLASS_SCE_SERVER                    = 273 //抓屏服务器
	DEV_CLASS_WIRELESS_TRANS                = 274 //无线传输设备
	DEV_CLASS_Y10_SERIES                    = 275 //Y10系列
	DEV_CLASS_SAFETY_MAVHINE                = 276 //安监一体机
	DEV_CLASS_IOTGATEWAY                    = 277 //物联网关类
	/* 报警主机相关 301 - 350*/
	DEV_CLASS_VIDEO_ALARM_HOST    = 301 //视频报警主机
	DEV_CLASS_NET_ALARM_HOST      = 302 //网络报警主机
	DEV_CLASS_ONE_KEY_ALARM       = 303 //一键式报警产品
	DEV_CLASS_WIRELESS_ALARM_HOST = 304 //无线报警主机
	DEV_CLASS_ALARM_MODULE        = 305 //报警模块
	DEV_CLASS_HOME_ALARM_HOST     = 306 //家用报警主机
	DEV_CLASS_HYBRID_ALARM_HOST   = 307 //混合报警主机

	/* 门禁相关 351 - 400*/
	DEV_CLASS_ACCESS_CONTROL = 351 //门禁产品

	/* 可视对讲 401 - 450*/
	DEV_CLASS_VIDEO_INTERCOM = 401 //可视对讲

	/* 无人机 451 - 500*/
	DEV_CLASS_UMMANNED_AERIAL_VEHICLE = 451 //无人机产品

	/* 移动产品: 501-550*/
	DEV_CLASS_MOBILE = 501 //移动产品

	/* 移动车载设备: 551-600*/
	DEV_CLASS_MOBILE_VEHICLE = 551 //移动车载设备

	//智能分析仪：601-650
	DEV_CLASS_INTELLIGENT_ANALYZER = 601 //智能分析仪

	//智能交通服务器：651-700
	DEV_CLASS_INTELLIGENT_TRAFFIC_SERVER = 651 //智能交通服务器
	DS_TP2200_EC                         = 652 //经济型机柜监测仪

	/* nvs相关 701-750 */
	DEV_CLASS_NVS = 701 //普通nvs

	/*有源RFID系列 751-800*/
	DS_TRI21A_1_P = 751 //有源RFID读取器

	/* 数据中心设备801-850 */
	DS_CLASS_FA    = 801 //脸谱单机
	DS_CLASS_PURE  = 802 //脸谱纯分析
	DS_CLASS_FS    = 803 //人脸静态数据服务器
	DS_CLASS_FD    = 804 //抓拍检测服务器
	DS_CLASS_HAWK  = 805 //猎鹰服务器
	DS_CLASS_BLADE = 806 //刀锋
	DS_CLASS_HMCP  = 807 //模型对比服务器

	/* 智能锁相关 851 - 900*/
	DEV_CLASS_SMART_LOCK = 851 //智能锁盒子

	/* 雷达相关 901 - 950*/
	DEV_CLASS_RADAR = 901 //雷达产品

	/* 智慧消防相关 951 - 1000*/
	DEV_CLASS_FIRE_CONTROL = 951 //消防产品

	/* 安检相关 1001 - 1050*/
	DEV_CLASS_SECURITY_CHECK = 1001 //安检产品

	/*全景细节相机： 8451-8470*/
	iDS_PT = 8451 //全景细节相机

	/* 其他设备类型 65534 */
	DEV_CLASS_DEFAULT = 65534 //默认设备类型
	/**********************设备大类 end**********************/

	/******************************能力集获取*********************************/
	//能力获取命令
	DEVICE_SOFTHARDWARE_ABILITY         = 0x001 //设备软硬件能力
	DEVICE_NETWORK_ABILITY              = 0x002 //设备网络能力
	DEVICE_ENCODE_ALL_ABILITY           = 0x003 //设备所有编码能力
	DEVICE_ENCODE_CURRENT               = 0x004 //设备当前编码能力
	IPC_FRONT_PARAMETER                 = 0x005 //ipc前端参数1.0
	IPC_UPGRADE_DESCRIPTION             = 0x006 //ipc升级信息
	DEVICE_RAID_ABILITY                 = 0x007 //RAID能力
	DEVICE_ENCODE_ALL_ABILITY_V20       = 0x008 //设备所有编码能力2.0
	IPC_FRONT_PARAMETER_V20             = 0x009 //ipc前端参数2.0
	DEVICE_ALARM_ABILITY                = 0x00a //辅助报警能力
	DEVICE_DYNCHAN_ABILITY              = 0x00b //设备数字通道能力
	DEVICE_USER_ABILITY                 = 0x00c //设备用户管理参数能力
	DEVICE_NETAPP_ABILITY               = 0x00d //设备网络应用参数能力
	DEVICE_VIDEOPIC_ABILITY             = 0x00e //设备图像参数能力
	DEVICE_JPEG_CAP_ABILITY             = 0x00f //设备JPEG抓图能力
	DEVICE_SERIAL_ABILITY               = 0x010 //RS232和RS485串口能力
	DEVICE_ABILITY_INFO                 = 0x011 //设备通用能力类型，具体能力根据发送的能力节点来区分
	STREAM_ABILITY                      = 0x012 //流能力
	SYSTEM_MANAGEMENT_ABILITY           = 0x013 //设备系统管理能力
	IP_VIEW_DEV_ABILITY                 = 0x014 //IP可视对讲分机能力
	VCA_DEV_ABILITY                     = 0x100 //设备智能分析的总能力
	VCA_CHAN_ABILITY                    = 0x110 //行为分析能力
	TRANSFER_ABILITY                    = 0x120
	MATRIXDECODER_ABILITY               = 0x200 //多路解码器显示、解码能力
	VIDEOPLATFORM_ABILITY               = 0x210 //视频综合平台能力集
	VIDEOPLATFORM_SBUCODESYSTEM_ABILITY = 0x211 //视频综合平台编码子系统能力集
	WALL_ABILITY                        = 0x212 //电视墙能力集
	MATRIX_ABILITY                      = 0x213 //SDI矩阵能力
	DECODECARD_ABILITY                  = 0x220 //解码卡服务器能力集
	VIDEOPLATFORM_ABILITY_V40           = 0x230 //视频综合平台能力集
	MATRIXMANAGEDEVICE_ABILITY          = 0x240 //矩阵接入网关能力集
	MATRIXDECODER_ABILITY_V41           = 0x260 //解码器能力集
	DECODER_ABILITY                     = 0x261 //解码器xml能力集
	DECODECARD_ABILITY_V41              = 0x270 //解码卡服务器能力集V41
	CODECARD_ABILITY                    = 0x271 //编码卡能力集
	SNAPCAMERA_ABILITY                  = 0x300 //抓拍机能力集
	ITC_TRIGGER_MODE_ABILITY            = 0x301 //智能IPC设备的触发模式能力
	COMPRESSIONCFG_ABILITY              = 0x400 //获取压缩参数能力集合
	COMPRESSION_LIMIT                   = 0x401 //获取主子码流压缩参数能力限制
	PIC_CAPTURE_ABILITY                 = 0x402 //获图片分辨率能力集合
	ALARMHOST_ABILITY                   = 0x500 //网络报警主机能力集
	IT_DEVICE_ABILITY                   = 0x501 //智能交通能力集
	SCREENCONTROL_ABILITY               = 0x600 //大屏控制器能力集
	SCREENSERVER_ABILITY                = 0x610 //大屏服务器能力集
	FISHEYE_ABILITY                     = 0x700 //鱼眼能力集
	LCD_SCREEN_ABILITY                  = 0x800 //LCD屏幕能力 2013-10-12
	ACS_ABILITY                         = 0x801 //门禁能力
	MERGEDEV_ABILITY                    = 0x802 //合码器能力集
	CAM_FUSION_ABILITY                  = 0x803 //相机拼接能力
	OPTICAL_DEV_ACCESS_ABILITY          = 0x805 //光端机接入能力
	NET_RING_ABILITY                    = 0x806 //环网能力集
	LED_ABILITY                         = 0x807 //LED屏能力集
	PUBLISHDEV_ABILITY                  = 0x80a //信息发布能力
	SCREEN_EXCHANGE_ABILITY             = 0x80b //屏幕互动能力
	REMOTE_NETMGR_FOT_ABILITY           = 0x80e //远端网管收发器能力
	/*************************************************
	  参数配置结构、参数(其中_V30为9000新增)
	  **************************************************/

	//子板异常信息
	//主类型
	EXCEPTION_MAJOR_MAINBOARD_BOOT = 0x1 //主板启动类型
	//次类型
	EXCEPTION_MINOR_PCIE_SCAN         = 0x1 // pcie链路扫描异常
	EXCEPTION_MINOR_DOWNLOAD_SUBBOARD = 0xa //下载子板完成异常

	//主类型
	EXCEPTION_MAJOR_SUBBOARD_BOOT = 0x2 //子板启动类型
	//次类型
	EXCEPTION_MINOR_INEXISTENCE          = 0x1  //PCI-E扫不到或当前子板不存在
	EXCEPTION_MINOR_UBOOT_DOWNLOAD       = 0xa  // uboot下载异常
	EXCEPTION_MINOR_UBOOT_INIT           = 0xe  //uboot初始化异常
	EXCEPTION_MINOR_ROOTFS_DOWNLOAD      = 0x14 //rootfs.img下载异常
	EXCEPTION_MINOR_UIMAGE_DOWNLOAD      = 0x19 //uImage下载异常
	EXCEPTION_MINOR_UBOOT_SETBOOTFLAG    = 0x1e // uboot启动标志位置位异常
	EXCEPTION_MINOR_ROOTFS_BOOT_SUBBOARD = 0x23 // rootfs启动异常
	EXCEPTION_MINOR_NEED_FILE_FINISH     = 0x28 //子板所需文件传输异常

	//主类型
	EXCEPTION_MAJOR_SUBBOARD_HARDWARE = 0x3 //子板硬件类型
	//次类型
	EXCEPTION_MINOR_AD    = 0x1 //AD异常
	EXCEPTION_MINOR_DA    = 0xa // DA异常
	EXCEPTION_MINOR_TIMER = 0xb //时钟异常

	//主类型
	EXCEPTION_MAJOR_FPGA = 0x4 //FPGA类型
	//次类型
	EXCEPTION_MINOR_IDLE             = 0x1  //无法IDLE
	EXCEPTION_MINOR_LANE             = 0xa  // LANE OK失败
	EXCEPTION_MINOR_REGISTER_ALL_F   = 0xe  //FPGA寄存器全F
	EXCEPTION_MINOR_MEMORY_INIT_FAIL = 0x14 //FPGA内存初始化失败
	//主类型
	EXCEPTION_MAJOR_DSP = 0x5 //DSP类型

	//主类型
	EXCEPTION_MAJOR_ARM = 0x6 //ARM类型

	//主类型
	EXCEPTION_MAJOR_BACKBOARD = 0x7 //背板类型
	//次类型
	EXCEPTION_MINOR_BLACKBOARD_TYPE = 0x1  //获取背板类型异常
	EXCEPTION_MINOR_SERDES          = 0xa  //视频交换芯片
	EXCEPTION_MINOR_CLOCK           = 0xe  //时钟故障
	EXCEPTION_MINOR_SYNCH           = 0x14 //同步信号故障

	//主类型
	EXCEPTION_MAJOR_SUBBOARD_NET = 0x8 //子板网络
	//次类型
	EXCEPTION_MINOR_IP_CONFLICT = 0x1  //IP冲突
	EXCEPTION_MINOR_DISCONNECT  = 0x14 // 断网

	//[add]by zengxiaole 2017-09-27 DS-19D2000-S v2.0
	MAX_FIRE_ALARM_ZONE       = 12 //最大消防主机报警分区个数
	MAX_FIRE_ALARM_POINT_ZONE = 32 //最大消防主机报警火点个数

	MAX_TIMESIGN_LEN = 32 //自定义校时标记信息长度

	/*设备报警和异常处理方式*/
	/*设备报警和异常处理方式*/
	NOACTION        = 0x0   /*无响应*/
	WARNONMONITOR   = 0x1   /*监视器上警告*/
	WARNONAUDIOOUT  = 0x2   /*声音警告*/
	UPTOCENTER      = 0x4   /*上传中心*/
	TRIGGERALARMOUT = 0x8   /*触发报警输出*/
	TRIGGERCATPIC   = 0x10  /*触发抓图并上传E-mail*/
	SEND_PIC_FTP    = 0x200 /*抓图并上传ftp*/

	DEV_ID_LEN = 32 //设备ID长度

	URL_LEN = 240 //URL长度

	// 邦诺CVR
	MAX_ID_COUNT        = 256
	MAX_STREAM_ID_COUNT = 1024
	STREAM_ID_LEN       = 32
	PLAN_ID_LEN         = 32
	DEVICE_NO_LEN       = 24
	MAX_VOLUMENAME_LEN  = 32 //录像卷名称
	MAX_VAG_CHANNO_LEN  = 32 //VAG协议取流时通道号编码长度

	MAX_STREAM_ID_NUM = 30 //最大流ID数目

	//ATM专用
	/****************************ATM(begin)***************************/
	NCR            = 0
	DIEBOLD        = 1
	WINCOR_NIXDORF = 2
	SIEMENS        = 3
	OLIVETTI       = 4
	FUJITSU        = 5
	HITACHI        = 6
	SMI            = 7
	IBM            = 8
	BULL           = 9
	YiHua          = 10
	LiDe           = 11
	GDYT           = 12
	Mini_Banl      = 13
	GuangLi        = 14
	DongXin        = 15
	ChenTong       = 16
	NanTian        = 17
	XiaoXing       = 18
	GZYY           = 19
	QHTLT          = 20
	DRS918         = 21
	KALATEL        = 22
	NCR_2          = 23
	NXS            = 24

	/*解码设备控制码定义*/
	NET_DEC_STARTDEC      = 1
	NET_DEC_STOPDEC       = 2
	NET_DEC_STOPCYCLE     = 3
	NET_DEC_CONTINUECYCLE = 4
	/*连接的通道配置*/

	MAX_RESOLUTIONNUM = 64 //支持的最大分辨率数目

	/*编码类型*/
	NET_DVR_ENCODER_UNKOWN   = 0 /*未知编码格式*/
	NET_DVR_ENCODER_H264     = 1 /*私有 264*/
	NET_DVR_ENCODER_S264     = 2 /*Standard H264*/
	NET_DVR_ENCODER_MPEG4    = 3 /*MPEG4*/
	NET_DVR_ORIGINALSTREAM   = 4 /*Original Stream*/
	NET_DVR_PICTURE          = 5 /*Picture*/
	NET_DVR_ENCODER_MJPEG    = 6
	NET_DVR_ENCODER_MPEG2    = 7
	NET_DVR_ENCODER_H265     = 8
	NET_DVR_ENCODER_SVAC     = 9
	NET_DVR_ENCODER_SMART264 = 10 /*Smart 264*/
	NET_DVR_ENCODER_SMART265 = 11 /*Smart 265*/

	/* 打包格式 */
	NET_DVR_STREAM_TYPE_UNKOWN = 0  /*未知打包格式*/
	NET_DVR_STREAM_TYPE_PRIVT  = 1  /*私有格式*/
	NET_DVR_STREAM_TYPE_TS     = 7  /* TS打包 */
	NET_DVR_STREAM_TYPE_PS     = 8  /* PS打包 */
	NET_DVR_STREAM_TYPE_RTP    = 9  /* RTP打包 */
	NET_DVR_STREAM_TYPE_ORIGIN = 10 //未打包(视频综合平台解码子系统用)

	//低帧率定义
	LOW_DEC_FPS_1_2  = 51
	LOW_DEC_FPS_1_4  = 52
	LOW_DEC_FPS_1_8  = 53
	LOW_DEC_FPS_1_16 = 54

	/*显示通道状态*/
	NET_DVR_MAX_DISPREGION = 16 /*每个显示通道最多可以显示的窗口*/

	MAX_DECODECHANNUM = 32 //多路解码器最大解码通道数
	MAX_DISPCHANNUM   = 24 //多路解码器最大显示通道数

	PASSIVE_DEC_PAUSE       = 1 /*被动解码暂停(仅文件流有效)*/
	PASSIVE_DEC_RESUME      = 2 /*恢复被动解码(仅文件流有效)*/
	PASSIVE_DEC_FAST        = 3 /*快速被动解码(仅文件流有效)*/
	PASSIVE_DEC_SLOW        = 4 /*慢速被动解码(仅文件流有效)*/
	PASSIVE_DEC_NORMAL      = 5 /*正常被动解码(仅文件流有效)*/
	PASSIVE_DEC_ONEBYONE    = 6 /*被动解码单帧播放(保留)*/
	PASSIVE_DEC_AUDIO_ON    = 7 /*音频开启*/
	PASSIVE_DEC_AUDIO_OFF   = 8 /*音频关闭*/
	PASSIVE_DEC_RESETBUFFER = 9 /*清空缓冲区*/

	/************************************多路解码器(end)***************************************/
	//2009-8-19 视频综合平台接口函数
	/************************************视频综合平台(begin)***************************************/
	MAX_SUBSYSTEM_NUM  = 80 //一个矩阵系统中最多子系统数量
	MAX_SERIALLEN      = 36 //最大序列号长度
	MAX_LOOPPLANNUM    = 16 //最大计划切换组
	DECODE_TIMESEGMENT = 4  //计划解码每天时间段数

	NET_DVR_DEV_ADDRESS_MAX_LEN    = 129
	NET_DVR_LOGIN_USERNAME_MAX_LEN = 64
	NET_DVR_LOGIN_PASSWD_MAX_LEN   = 64

	WIFI_WEP_MAX_KEY_COUNT         = 4
	WIFI_WEP_MAX_KEY_LENGTH        = 33
	WIFI_WPA_PSK_MAX_KEY_LENGTH    = 63
	WIFI_WPA_PSK_MIN_KEY_LENGTH    = 8
	WIFI_MAX_AP_COUNT              = 20
	WIFI_WPA_PSK_MAX_HEXKEY_LENGTH = 68 //WPA16进制密钥最大长度

	//结构参数宏定义
	VCA_MAX_POLYGON_POINT_NUM      = 10   //检测区域最多支持10个点的多边形
	MAX_RULE_NUM                   = 8    //最多规则条数
	MAX_RULE_NUM_V42               = 16   //最多规则条数扩展
	MAX_TARGET_NUM                 = 30   //最多目标个数
	MAX_CALIB_PT                   = 6    //最大标定点个数
	MIN_CALIB_PT                   = 4    //最小标定点个数
	MAX_TIMESEGMENT_2              = 2    //最大时间段数
	DATA_INDEX_LEN                 = 64   //数据流水号
	MAX_DEV_DATAINDEX_LEN          = 64   //设备数据流水号
	MAX_TRAFFIC_PICTURE_NUM        = 8    //交通图片数量
	MAX_LICENSE_LEN                = 16   //车牌号最大长度
	MAX_LICENSE_LEN_EX             = 32   //车牌号最大长度
	MAX_CARDNO_LEN                 = 48   //卡号最大长度 2013-11-04
	MAX_OPERATE_INDEX_LEN          = 32   //操作数最大长度2014-03-03
	MAX_PLATE_NUM                  = 3    //车牌个数
	MAX_MASK_REGION_NUM            = 4    //最多四个屏蔽区域
	MAX_SEGMENT_NUM                = 6    //摄像机标定最大样本线数目
	MIN_SEGMENT_NUM                = 3    //摄像机标定最小样本线数目
	MAX_REL_SNAPCHAN_NUM           = 3    //最大关联抓图通道数
	MAX_PIC_SWITCH_STORAGE_SERVER  = 64   //云存储服务器存储的最大图片类型数
	MAX_INFO_SWITCH_STORAGE_SERVER = 64   //云存储服务器存储的最大附加信息类型数
	RTMP_URL_LEN                   = 128  //RTMP URL 长度
	MAX_ID_LEN_128                 = 128  //发布文件ID长度
	MAX_DEBUGCMD_LEN               = 1024 //设备调试命令最大长度
	MAX_DEBUGINFO_LEN              = 1400 //设备调试信息最大长度
	MAX_VEHICLE_ID_LEN             = 32   //最大车辆标识长度
	LEN_PROPERTY                   = 128

	//智能控制信息
	MAX_VCA_CHAN = 16 //最大智能通道数

	VCA_ATTEND_MAX_PIC_NUM      = 3  //考勤事件最大图片张数
	VCA_ATTEND_DRIVER_NAME_LEN  = 64 //分组信息司机名字长度
	VCA_ATTEND_CARD_ID_LEN      = 32 //分组信息司机证件号码长度
	VCA_ATTEND_MAX_ALARM_ID_LEN = 32 //报警事件唯一编号的最大长度

	MAX_NET_DISK = 16 //最大网络硬盘个数

	INQUEST_START_INFO   = 0x1001 /*讯问开始信息*/
	INQUEST_STOP_INFO    = 0x1002 /*讯问停止信息*/
	INQUEST_TAG_INFO     = 0x1003 /*重点标记信息*/
	INQUEST_SEGMENT_INFO = 0x1004 /*审讯片断状态信息*/
	INQUEST_CASE_INFO    = 0x1005 // 案件信息类型

	SEARCH_EVENT_INFO_LEN  = 300 //事件信息长度
	CASE_NO_LEN            = 64
	CASE_NAME_LEN          = 128
	LITIGANT_LEN           = 32
	CHIEF_JUDGE_LEN        = 32
	SEARCH_CASE_NO_LEN     = 56
	SEARCH_CASE_NAME_LEN   = 100
	SEARCH_LITIGANT_LEN    = 32
	SEARCH_CHIEF_JUDGE_LEN = 32
	CASE_NO_RET_LEN        = 52
	CASE_NAME_RET_LEN      = 64
	LITIGANT_RET_LEN       = 24
	CHIEF_JUDGE_RET_LEN    = 24
	NET_SDK_CASETYPE_LEN   = 32

	MAX_POS_KEYWORDS_NUM      = 3   //支持关键字查找条数
	MAX_POS_KEYWORD_LEN       = 128 //每条关键字长度
	SEARCH_EVENT_INFO_LEN_V40 = 800

	NET_SDK_MAX_TAPE_INDEX_LEN = 32  //磁带编号最大长度
	NET_SDK_MAX_FILE_LEN       = 256 //文件名最大长度

	MAX_RECT_NUM = 6

	MAX_LINE_SEG_NUM = 8

	MAX_SAMPLE_NUM    = 5 //直接标定点最大个数
	MAX_SAMPLE_NUM_EX = 7 //样本标定点个数扩展

	CALIB_PT_NUM = 4

	NET_SDK_MAX_RELATED_CHAN_NUM = 4 //最大关联通道数

	MAX_POSITION_NUM = 10

	/********************************智能交通事件 begin****************************************/
	MAX_REGION_NUM = 8 // 区域列表最大数目
	MAX_TPS_RULE   = 8 // 最大参数规则数目
	MAX_AID_RULE   = 8 // 最大事件规则数目
	MAX_LANE_NUM   = 8 // 最大车道数目

	XXX_MAJOR_VERSION = 2

	XXX_SUB_VERSION = 3

	XXX_REVISION_VERSION = 4

	/*可用来命名图片的相关元素 */
	PICNAME_ITEM_DEV_NAME    = 1  /*设备名*/
	PICNAME_ITEM_DEV_NO      = 2  /*设备号*/
	PICNAME_ITEM_DEV_IP      = 3  /*设备IP*/
	PICNAME_ITEM_CHAN_NAME   = 4  /*通道名*/
	PICNAME_ITEM_CHAN_NO     = 5  /*通道号*/
	PICNAME_ITEM_TIME        = 6  /*时间*/
	PICNAME_ITEM_CARDNO      = 7  /*卡号*/
	PICNAME_ITEM_PLATE_NO    = 8  /*车牌号码*/
	PICNAME_ITEM_PLATE_COLOR = 9  /*车牌颜色*/
	PICNAME_ITEM_CAR_CHAN    = 10 /*车道号*/
	PICNAME_ITEM_CAR_SPEED   = 11 /*车辆速度*/
	PICNAME_ITEM_CARCHAN     = 12 /*监测点*/
	PICNAME_ITEM_PIC_NUMBER  = 13 //图片序号
	PICNAME_ITEM_CAR_NUMBER  = 14 //车辆序号
	PICNAME_MAXITEM          = 15

	PICNAME_ITEM_SPEED_LIMIT_VALUES = 15 //限速值
	PICNAME_ITEM_ILLEGAL_CODE       = 16 //国标违法代码
	PICNAME_ITEM_CROSS_NUMBER       = 17 //路口编号
	PICNAME_ITEM_DIRECTION_NUMBER   = 18 //方向编号

	//(3.7Ver)
	PICNAME_ITEM_CAR_COLOR        = 19  //车身颜色
	PICNAME_ITEM_PLATE_COORDINATE = 20  //车牌坐标
	PICNAME_ITEM_CAR_TYPE         = 21  //车辆类型
	PICNAME_ITEM_VIOLATION_TYPE   = 22  //违规类型
	PICNAME_ITEM_CUSTOM           = 255 //自定义

	//命名规则：2013-09-27
	PICNAME_ITEM_PARK_DEV_IP     = 1  /*设备IP*/
	PICNAME_ITEM_PARK_PLATE_NO   = 2  /*车牌号码*/
	PICNAME_ITEM_PARK_TIME       = 3  /*时间*/
	PICNAME_ITEM_PARK_INDEX      = 4  /*车位编号*/
	PICNAME_ITEM_PARK_STATUS     = 5  /*车位状态*/
	PICNAME_ITEM_BUILDING_NUMBER = 6  /*栋号单元号*/
	PICNAME_ITEM_OUTDOOR_UNIT_ID = 7  /*门口机编号*/
	PICNAME_ITEM_UNLOCK_TYPE     = 8  /*开锁方式*/
	PICNAME_ITEM_DEVICE_NAME     = 9  //设备名称
	PICNAME_ITEM_PERIOD_NO       = 10 /*期号*/
	PICNAME_ITEM_DEV_INDEX       = 11 /*设备编号*/
	PICNAME_PREFIX               = 32 /*图片名自定义前缀长度*/

	MAX_ALERTLINE_NUM = 8 //最大警戒线条数

	MAX_INTRUSIONREGION_NUM = 8 //最大区域数数

	MAX_ZEROCHAN_NUM = 16

	DESC_LEN_64 = 64

	PROCESSING                   = 0   //正在处理
	PROCESS_SUCCESS              = 100 //过程完成
	PROCESS_EXCEPTION            = 400 //过程异常
	PROCESS_FAILED               = 500 //过程失败
	PROCESS_QUICK_SETUP_PD_COUNT = 501 //一键配置至少3块硬盘

	SOFTWARE_VERSION_LEN     = 48
	NET_SDK_DEVICE_MODEL_LEN = 24 //设备型号长度

	MAX_SADP_NUM = 256 // 搜索到设备最大数目

	/*******************************备份接口 begin********************************/
	//获取备份设备信息接口定义
	DESC_LEN_32  = 32  //描述字长度
	MAX_NODE_NUM = 256 //节点个数

	//备份进度列表
	BACKUP_SUCCESS       = 100 //备份完成
	BACKUP_CHANGE_DEVICE = 101 //备份设备已满，更换设备继续备份

	BACKUP_SEARCH_DEVICE   = 300 //正在搜索备份设备
	BACKUP_SEARCH_FILE     = 301 //正在搜索录像文件
	BACKUP_SEARCH_LOG_FILE = 302 //正在搜索日志文件
	BACKUP_CHANGE_DISK     = 303 //正在更换光盘

	BACKUP_EXCEPTION = 400 //备份异常
	BACKUP_FAIL      = 500 //备份失败

	BACKUP_TIME_SEG_NO_FILE  = 501 //时间段内无录像文件
	BACKUP_NO_RESOURCE       = 502 //申请不到资源
	BACKUP_DEVICE_LOW_SPACE  = 503 //备份设备容量不足
	BACKUP_DISK_FINALIZED    = 504 //刻录光盘封盘
	BACKUP_DISK_EXCEPTION    = 505 //刻录光盘异常
	BACKUP_DEVICE_NOT_EXIST  = 506 //备份设备不存在
	BACKUP_OTHER_BACKUP_WORK = 507 //有其他备份操作在进行
	BACKUP_USER_NO_RIGHT     = 508 //用户没有操作权限
	BACKUP_OPERATE_FAIL      = 509 //操作失败
	BACKUP_NO_LOG_FILE       = 510 //硬盘中无日志

	MAX_ABILITYTYPE_NUM = 12 //最大能力项

	/********************************9000RH begin****************************************/
	SUPPORT_PD_NUM         = 16
	SUPPORT_ARRAY_NUM      = 8
	SUPPORT_VD_NUM         = 128
	SUPPORT_PD_NUM_        = 16
	SUPPORT_PD_NUM_PARTTWO = 8

	INIT_QUICK          = 0 /*快速初始化*/
	INIT_FULLFOREGROUND = 1 /*完全初始化(前台)*/
	INIT_FULLBACKGROUND = 2 /*完全初始化(后台)*/

	MATRIX_MAXDECSUBSYSTEMCHAN = 4 //视频综合平台解码子系统通道号

	/***************************审讯DVR begin *****************************/
	MAX_INQUEST_PIP_NUM    = 3  //单通道显示的画中画数目
	MAX_INQUEST_CDRW_NUM   = 4  //最大刻录机数目
	MAX_INQUEST_PIP_NUM_EX = 16 //审讯机画中画最大个数

	/********************************接口参数结构(end)*********************************/
	MAX_BIGSCREENNUM = 100 //最多大屏拼接数

	DECODEPIC_LEFTADJUST                = 1  /*图像左移*/
	DECODEPIC_RIGHTADJUST               = 2  /*图像右移*/
	DECODEPIC_UPADJUST                  = 3  /*图像上移*/
	DECODEPIC_DOWNADJUST                = 4  /*图像下移*/
	DECODEPIC_REDUCEADJUST              = 5  /*图像缩小*/
	DECODEPIC_FULL_SCREEN_ADJUST        = 6  /*图像全屏*/
	DECODEPIC_CANCEL_FULL_SCREEN_ADJUST = 7  /*图像取消全屏显示*/
	DECODEPIC_AUTOADJUST                = 8  /*  图像自动调整 */
	DECODEPIC_HEIGHTADJUST              = 9  /* 图像高度调整 */
	DECODEPIC_WIDTHADJUST               = 10 /* 图像宽度调整 */

	MAX_UNITEDMATRIX_NUM = 8 //级联中最多视频综合平台数量
	MAX_SUBDOMAIN_NUM    = 4 //级联中最多子域数量

	///拨号功能
	DIALPASSWD_LEN = 32 //拨号密码长度

	MAX_PHONE_NUM = 32 //最长号码长度

	//短信功能
	MAX_WHITELIST_NUM            = 8  //最大白名单数
	NET_SDK_MAX_WHITELIST_NUM_32 = 32 //最大白名单数

	HARDDISKFULL_EXCEPTION   = 0x0 /*硬盘满*/
	HARDDISKERROR_EXCEPTION  = 0x1 /*硬盘错*/
	ETHERNETBROKEN_EXCEPTION = 0x2 /*网线断*/
	IPADDRCONFLICT_EXCEPTION = 0x3 /*IP地址冲突*/
	ILLEGALACCESS_EXCEPTION  = 0x4 /*非法访问*/
	VI_EXCEPTION             = 0x5 /*视频信号异常*/
	VS_MISMATCH_EXCEPTION    = 0x6 /*输入/输出视频制式不匹配 */
	VIDEOCABLELOSE_EXCEPTION = 0x7 /*视频无信号*/
	AUDIOCABLELOSE_EXCEPTION = 0x8 /*音频无信号*/
	ALARMIN_EXCEPTION        = 0x9 /*报警输入*/
	MASKALARM_EXCEPTION      = 0xa /*遮挡报警*/
	MOTDET_EXCEPTION         = 0xb /*移动侦测*/
	RECORDING_EXCEPTION      = 0xc /*录像异常*/
	WIRELESS_EXCEPTION       = 0xd /*PIR报警*/
	PIR_EXCEPTION            = 0xe /*无线报警*/
	CALLHELP_EXCEPTION       = 0xf /*呼救报警*/

	AUDIO_DETECTION_EXCEPTION          = 0x10 /*音频异常侦测报警*/
	SCENECHANGE_DETECTION_EXCEPTION    = 0x11 /*场景侦测报警*/
	DEFOCUS_DETECTION_EXCEPTION        = 0x12 /*虚焦侦测报警*/
	FACE_DETECTION_ENTRANCE_EXCEPTION  = 0x13 /*人脸侦测报警*/
	LINE_DETECTION_ENTRANCE_EXCEPTION  = 0x14 /*越界侦测报警*/
	FIELD_DETECTION_ENTRANCE_EXCEPTION = 0x15 /*区域入侵侦测报警*/
	REGION_EXITING_EXCEPTION           = 0x16 /*离开区域侦测报警*/
	REGION_ENTRANCE_EXCEPTION          = 0x17 /*进入区域报警*/
	LOITERING_EXCEPTION                = 0x18 /*人员徘徊侦测报警*/
	GROUP_EXCEPTION                    = 0x19 /*人员聚集侦测报警*/
	RAPIDMOVE_EXCEPTION                = 0x1a /*快速移动侦测报警*/
	PARKING_EXCEPTION                  = 0x1b /*停车侦测报警*/
	UNATTENDEDBAGGAGE_EXCEPTION        = 0x1c /*物品遗留侦测报警*/
	ATTENDEDBAGGAGE_EXCEPTION          = 0x1d /*物品拿取侦测报警*/
	DATATRAFFIC_EXCESS                 = 0x1e /*流量超额*/
	VOLTAGEINSTABLE_EXCEPTION          = 0x1f /*电源电压异常报警*/
	ALL_EXCEPTION                      = 0xff /*所有事件，根据事件联动方式判断是否开启短信联动*/

	PHONECFG_RECEIVE_SMS         = 0x0 /* 支持接收报警短信 */
	PHONECFG_SMS_CONTROL         = 0x1 /* 支持短信控制上下线 */
	PHONECFG_CALL_CONTROL        = 0x2 /* 支持呼叫控制上线 */
	PHONECFG_SMS_REBOOT          = 0x3 /*支持短信重启*/
	PHONECFG_DOOR_CONTROL        = 0x4 /* 支持门操作控制*/
	PHONECFG_SMS_GET_DAIL_STATUS = 0x5 /* 支持短信获取拨号状态*/

	NET_SDK_MONITOR_ID_LEN = 64 //监控点ID长度

	MAX_SMSCONTENT_LEN = 140 //短信内容长度

	MAX_PIN_LEN = 12 //PIN码最大长度

	COM_PUSHALARM     = 0x1200 //设备基本报警信息上传，推模式设备使用
	COM_PUSHALARM_V30 = 0x1201 //设备基本报警信息上传v30，推模式设备使用

	SENSOR_IN_NUMBER = 8

	MAX_ALARMHOST_ALARMIN_NUM  = 512 //网络报警主机最大报警输入口数
	MAX_ALARMHOST_ALARMOUT_NUM = 512 //网络报警主机最大报警输出口数

	ALARMHOST_MAX_AUDIOOUT_NUM        = 32  //网络报警主机最大语音输出数
	ALARMHOST_MAX_ELECTROLOCK_NUM     = 32  //网络报警主机最大电锁数
	ALARMHOST_MAX_MOBILEGATE_NUM      = 32  //网络报警主机最大移动门数
	ALARMHOST_MAX_SIREN_NUM           = 8   // 最大警号数目
	MAX_ALARMHOST_SUBSYSTEM           = 32  //报警主机最大子系统数
	ALARMHOST_DETECTOR_SERIAL_LEN     = 9   //报警主机关联探测器序列号长度
	ALARMHOST_DETECTOR_SERIAL_LEN_V50 = 16  //报警主机关联探测器序列号V50长度
	MAX_DETECTOR_NUM                  = 128 //最大关联探测器数
	MAX_DETECTOR_NUM_V51              = 256 //最大关联探测器数
	MAX_REPEATER_NUM                  = 16  //最大中继器数
	MAX_OUTPUT_MODULE_NUM             = 64  //最大输出模块数
	MAX_ELECTRIC_LOCK_NUM             = 64  //最大电锁数量

	MAX_MAX_ALARMIN_NUM = 64 /* 批量获取最大防区数*/

	MAX_485CHAN = 256 //485通道号
	MAX_485SLOT = 256 //485槽位号

	MAX_DEVICE_PROTO_NUM = 256
	MAX_DEVICE_TYPE_NUM  = 256

	MAX_ALARMHOST_VIDEO_CHAN = 64

	PROTOCOL_VERTION_LEN = 32

	//自助行拨号参数配置及启用方式配置
	MAX_CENTERNUM = 4 //G1,G2 G3 G4或者N1，N2，N3，N4或者T1，T2，T3，T4

	MAX_PU_CHAN_NUM = 512

	MAX_ALARM_CAM_NUM = 32 // 报警触发CAM最大个数

	MAX_GATEWAY_NUM = 8 // 最大门禁个数

	//LED屏显内容
	MAX_CONTENT_LEN = 512

	//LED定时开关机
	LED_TIMER_NUM      = 3  // LED开机、关机时间组数
	TIME_SEGMENT_A_DAY = 48 // 时间段个数，一天24小时，半小时一个段

	//2010-12-28 高清解码卡能力集 begin
	//新的解码卡服务器能力集
	MAX_DECODE_CARD_NUM = 6 //最多高清解码卡数

	MAX_SUBMATRIX_NUM = 8 //级联中子最多从系统数量

	MAX_GATEWAYTRUNKNUM = 1024 //级联视频综合平台中最大路由干线数

	MAX_SUBSYSTEM_NUM_V40 = 120

	MAX_OPTICALFIBER_NUM = 16

	MAX_HOLIDAY_NUM = 32

	MAX_LINK_V30 = 128

	MAX_BOND_NUM = 2

	MAX_PIC_EVENT_NUM   = 32
	MAX_ALARMIN_CAPTURE = 16

	//  录像标签
	LABEL_NAME_LEN = 40

	LABEL_IDENTIFY_LEN = 64

	MAX_DEL_LABEL_IDENTIFY = 20 // 删除的最大标签标识个数

	CARDNUM_LEN_V30 = 40

	PICTURE_NAME_LEN      = 64
	PICTURE_INFO_MAX_SIZE = 640 * 960 * 1.5

	MAX_RECORD_PICTURE_NUM = 50 //  最大备份图片张数

	STEP_READY     = 0   //准备升级
	STEP_RECV_DATA = 1   //接收升级包数据
	STEP_UPGRADE   = 2   //升级系统
	STEP_BACKUP    = 3   //备份系统
	STEP_SEARCH    = 255 //搜索升级文件

	MAX_REDAREA_NUM = 6 //最大红绿灯区域个数

	INQUEST_MESSAGE_LEN  = 44 //审讯重点标记信息长度
	INQUEST_MAX_ROOM_NUM = 2  //最大审讯室个数
	MAX_RESUME_SEGMENT   = 2  //支持同时恢复的片段数目

	UPLOAD_CERTIFICATE = 1 //上传证书

	MATRIX_PROTOCOL_NUM   = 20 //支持的最大矩阵协议数
	KEYBOARD_PROTOCOL_NUM = 20 //支持的最大键盘协议数

	MAX_HUMAN_PICTURE_NUM   = 10 //最大照片数
	MAX_HUMAN_BIRTHDATE_LEN = 10 //最大出生年月长度

	MAX_FACE_PIC_LEN = 6144 //最大人脸图片数据长度

	//显示通道画面分割模式
	MAX_WINDOWS_NUM = 12 //画面分割模式的数量
	MAX_SUPPORT_RES = 32
	MAX_DISPNUM_V41 = 32
	MAX_SDI_RES     = 16 //SDI显示通道最大支持分辨率数

	//显示通道配置
	MAX_WINDOWS     = 16
	MAX_WINDOWS_V41 = 36

	STARTDISPCHAN_VGA  = 1
	STARTDISPCHAN_BNC  = 9
	STARTDISPCHAN_HDMI = 25
	STARTDISPCHAN_DVI  = 29

	MAX_BIGSCREENNUM_SCENE = 100

	/*******************************窗口设置*******************************/
	MAX_WIN_COUNT = 224 //支持的最大开窗数

	MAX_LAYOUT_COUNT = 16 //最大布局数

	MAX_CAM_COUNT     = 224
	MAX_CAM_COUNT_V50 = 512

	/*******************************能力集*******************************/
	SCREEN_PROTOCOL_NUM = 20 //支持的最大大屏控制器协议数

	/*******************************OSD*******************************/
	MAX_OSDCHAR_NUM = 256

	/*******************************预案管理*******************************/
	MAX_PLAN_ACTION_NUM = 32 //预案动作个数
	DAYS_A_WEEK         = 7  //一周7天
	MAX_PLAN_COUNT      = 16 //预案个数

	// 安全拔盘状态
	PULL_DISK_SUCCESS     = 1 // 安全拔盘成功
	PULL_DISK_FAIL        = 2 // 安全拔盘失败
	PULL_DISK_PROCESSING  = 3 // 正在停止阵列
	PULL_DISK_NO_ARRAY    = 4 // 阵列不存在
	PULL_DISK_NOT_SUPPORT = 5 // 不支持安全拔盘

	// 扫描阵列状态
	SCAN_RAID_SUC         = 1 // 扫描阵列成功
	SCAN_RAID_FAIL        = 2 // 扫描阵列失败
	SCAN_RAID_PROCESSING  = 3 // 正在扫描阵列
	SCAN_RAID_NOT_SUPPORT = 4 // 不支持阵列扫描

	// 设置前端相机类型状态
	SET_CAMERA_TYPE_SUCCESS    = 1 // 成功
	SET_CAMERA_TYPE_FAIL       = 2 // 失败
	SET_CAMERA_TYPE_PROCESSING = 3 // 正在处理

	MAX_PRO_PATH = 256 //最大协议路径长度

	//////////子系统配置/////////////
	MAX_ALARMHOSTKEYBOARD = 64 //网络报警主机最大键盘数

	MAX_SUBSYSTEM_ID_LEN = 16 //子系统ID最大长度

	//////////GPRS参数配置/////////////
	ACCOUNTNUM_LEN       = 6
	ACCOUNTNUM_LEN_32    = 32
	ACCOUNTNUM_LEN_V40   = 9
	APN_NAME_LEN         = 32
	APN_USERNAME_LEN     = 24
	APN_USERPASSWORD_LEN = 16

	//子系统参数配置扩展
	MAX_KEYBOARD_USER_NUM = 256

	//////////积木上传方式/////////////
	MAX_REPORTCHAN_NUM  = 4
	MAX_CENTERGROUP_NUM = 16

	MAX_EVENT_NUM = 32 //网络报警主机最大事件数

	CID_CODE_LEN   = 4
	DEV_SERIAL_LEN = 9

	MAX_DECODE_CARD_SUPPORTDISPNUMS = 8 //每个解码卡最多支持的显示通道数

	CHAN_NO_LEN = 24

	MAX_IOSPEED_GROUP_NUM      = 4  //IO测速组个数
	MAX_IOOUT_NUM              = 4  //最大IO输出口个数
	MAX_IOIN_NUM               = 8  //最大IO输入口个数
	MAX_RELAY_NUM              = 12 //继电器控制设备最大数 2013-11-04
	MAX_VEHICLE_TYPE_NUM       = 8  //车辆信息管控最大数2013-11-04
	MAX_IOIN_NUMEX             = 10 //最大IO输入口个数(扩展)
	MAX_ITC_LANE_NUM           = 6  //最大车道个数
	MAX_LANEAREA_NUM           = 2  //单车道最大区域个数
	ITC_MAX_POLYGON_POINT_NUM  = 20 //检测区域最多支持20个点的多边形
	MAX_ITC_SERIALCHECK_NUM    = 8  //串口校验类型个数
	MAX_LIGHT_NUM              = 6  //最大交通灯数
	MAX_VIDEO_INTERVAL_NUM     = 2  //最大抓拍间隔数
	MAX_VIDEO_DETECT_LIGHT_NUM = 12 //视频检测的最大检测区域
	MAX_CALIB_RECOG_NUM        = 2  //标定区域个数
	MAX_RS485_NUM              = 12 //485口最大支持数
	MAX_MOBILE_POLYGON_NUM     = 3  //移动布控支持最大牌识区域个数
	MAX_MOBILE_DETECTLINE_NUM  = 3  //移动布控支持最大违规检测线个数
	MAX_IOOUT_K_NUM            = 8  //K系列最大IO输出口个数

	MAX_AUX_ALARM_NUM      = 8 //最大辅助报警个数
	MAX_WIRELESS_ALARM_NUM = 8 //最大无线报警个数

	DVCS_DEVICEID_LEN = 16

	DEVICEID_LEN = 32

	MAX_ESATA_NUM   = 16
	MAX_MINISAS_NUM = 96

	MAX_DISK_NUM = 128

	/*************************************ITS****************************/
	VERSION_LEN          = 32 //版本长度
	MAX_OVERLAP_ITEM_NUM = 50 //最大字符叠加种数
	ITS_MAX_DEVICE_NUM   = 32 //最大设备个数

	MAX_PTZCRUISE_POINT_NUM = 32 //最大支持32个巡航点

	NET_DVR_GPS_FINDING   = 0 //正在查找
	NET_DVR_GPS_RECV      = 1 //接收数据
	NET_DVR_GPS_OVER      = 2 //查找结束
	NET_DVR_GPS_EXCEPTION = 3 //接收异常

	IP_ADDR_LEN = 16

	NET_DVR_GPS_STATUS      = 0
	NET_DVR_GSENSOR_STATUS  = 1
	NET_DVR_WIFI_STATUS     = 2
	NET_DVR_PLATFORM_STATUS = 3

	NET_SDK_MAX_CARD_LEN = 32 //最大卡号长度

	MAX_CERTIFICATE_ISSUER_LEN   = 64  //证书颁发者长度
	MAX_CERTIFICATE_VALIDITY_LEN = 128 //证书有效时间长度
	MAX_CERTIFICATE_SUBJECT_LEN  = 64  //证书持有者长度

	//端口聚合
	MAX_ETHERNET_PORT_NUM = 8 //每条链路最大端口数

	//一键配置通用状态
	NET_SDK_OKC_STATUS_SUCCESS = 1000 //一键配置成功
	NET_SDK_OKC_STATUS_FAILED  = 1002 //一键配置失败

	//一键配置CVR状态
	NET_SDK_OKC_STATUS_START_CONFIG        = 1003 //开始配置
	NET_SDK_OKC_CHECK_HD                   = 1004 //检测磁盘
	NET_SDK_OKC_INIT_HD                    = 1005 //初始化磁盘
	NET_SDK_OKC_CREATE_RAID_OR_SINGLE_DISK = 1006 //创建阵列或者单盘模式
	NET_SDK_OKC_INIT_CVR_SERVICE           = 1007 //初始化CVR服务
	NET_SDK_OKC_CREATE_RECORD_VOLUME       = 1008 //创建录像卷

	//以下为一键配置失败的状态码——part1
	NET_SDK_OKC_ERR_LOAD_CONF_FAILED           = 1009 //加载配置文件失败
	NET_SDK_OKC_ERR_NOT_SUPPORT_RAID_LEVLE     = 1010 //不支持此种类型的raid
	NET_SDK_OKC_ERR_CONFIGURATION_CONFLICT     = 1011 //系统已经存在raid或存储池
	NET_SDK_OKC_ERR_GET_DISK_INFO_FAILED       = 1012 //获取磁盘信息失败
	NET_SDK_OKC_ERR_CHECK_DISK_FAILED          = 1013 //检测磁盘失败
	NET_SDK_OKC_ERR_INIT_DISK_FAILED           = 1014 //初始化磁盘失败
	NET_SDK_OKC_ERR_DISK_CAPACITY_SMALL        = 1015 //磁盘总容量不足
	NET_SDK_OKC_ERR_BOTH_SV_NS                 = 1016 //同时存在SV盘和NS盘
	NET_SDK_OKC_ERR_CREATE_RAID_FAILED         = 1017 //创建raid失败
	NET_SDK_OKC_ERR_GET_RAID_FAILED            = 1018 //获取raid失败
	NET_SDK_OKC_ERR_CREATE_SPARE_FAILED        = 1019 //创建热备盘失败
	NET_SDK_OKC_ERR_CREATE_STORAGE_POOL_FAILED = 1020 //创建存储池失败
	NET_SDK_OKC_ERR_GET_POOL_INFO_FAILED       = 1021 //获取存储池信息失败
	NET_SDK_OKC_ERR_CREATE_LUN_FAILED          = 1022 //创建lun卷失败
	NET_SDK_OKC_ERR_GET_LUN_INFO_FAILED        = 1023 //获取lun信息失败
	NET_SDK_OKC_ERR_CREATE_BACKUP_FAILED       = 1024 //创建预留卷失败
	NET_SDK_OKC_ERR_GET_BACKUP_FAILED          = 1025 //获取预留卷失败
	NET_SDK_OKC_ERR_CREATE_PRIVATE_LUN_FAILED  = 1026 //创建私有卷失败
	NET_SDK_OKC_ERR_CREATE_RV_FAILED           = 1027 //创建录像卷失败
	NET_SDK_OKC_ERR_CREATE_ARCH_RV_FAILED      = 1028 //创建存档卷失败
	NET_SDK_OKC_ERR_START_CVR_SERVICE_FAILED   = 1029 //开启CVR服务失败

	//一键配置SAN状态
	NET_SDK_OKC_CREATING_ARRAY           = 1030 //创建阵列阶段
	NET_SDK_OKC_CREATING_STORAGE_POOL    = 1031 //创建存储池阶段
	NET_SDK_OKC_CREATING_LUN_VOL         = 1032 //创建逻辑卷阶段
	NET_SDK_OKC_CREATING_ISCSI           = 1033 //创建ISCSI阶段
	NET_SDK_OKC_ERR_HOT_SPARE_CONFICT    = 1034 //已存在热备盘
	NET_SDK_OKC_ERR_STORAGE_POOL_CONFICT = 1035 //已存在存储池
	NET_SDK_OKC_ERR_RAID_CONFLICT        = 1036 //系统已经存在阵列
	NET_SDK_OKC_ERR_OPEN_ISCSI_FAILED    = 1037 //开启ISCSI失败
	NET_SDK_OKC_ERR_DEVICE_NOSUPPORT_SAN = 1038 //设备不支持san

	//以下为一键配置失败的状态码——part2
	NET_SDK_OKC_ERR_SAPRE_NUM_EXCEED         = 1101 //热备盘个数过多
	NET_SDK_OKC_ERR_CREATE_PIC_VOLUME_FAILED = 1102 //创建图片卷失败

	MAX_CODE_CARD_SUPPORTDISPNUMS = 8 //每个编码卡最多支持的显示通道数

	MAX_CODE_CARD_NUM = 8 //最多高清编码卡数

	NET_DVR_FIND_NAS_DIRECTORY = 6161 //查找NAS目录
	NET_DVR_NAS_FINDING        = 0    //正在查找
	NET_DVR_NAS_RECV           = 1    //接收数据
	NET_DVR_NAS_OVER           = 2    //查找结束
	NET_DVR_NAS_EXCEPTION      = 3    //接收异常

	NET_SDK_ACCESS_KEY_LEN = 64 //访问密码长度
	NET_SDK_SECRET_KEY_LEN = 64 //加密密码长度

	MAX_INPUTNUMS  = 1024
	MAX_OUTPUTNUMS = 256

	//显示路径
	MAX_MATRIX_CASCADE = 32

	ALARMHOST_ALARMOUT_NUM = 64 //触发器个数
	MAX_OSD_UNIT_LEN       = 8  //OSD单位长度

	UPLOAD_POS_INFO = 1001 //上传Pos信息

	MAX_IGNORE_STRING_NUM   = 4
	FILTERRULE_NUM          = 4
	MAX_POS_FILTER_DATA_LEN = 128

	SERIAL_NUM_LEN = 8

	MAX_LAN_ENCODE_LEN = 32 //语言编码格式最大长度

	IPC_PARAM_AGING_TRICK_SCAN = 0x00000001 //清除花样扫描参数设置

	MAX_DEVMODULE_NUM = 8

	//SDI矩阵1.0
	MATRIX_MAX_OUTPUT_NUM = 256 //矩阵最大输出通道个数

	MAX_HEATMAPREGION_NUM = 8

	MAX_CAMERAID_LEN = 64

	MAX_DISPLAY_NUM    = 512 //最大显示输出个数
	MAX_LEDCONTENT_NUM = 512 //虚拟LED字符串最大长度
	MAX_PPT_CHAN       = 128 //PPT长度

	MAX_SUBBOARD_NUM               = 42 //集中式大屏设备板数目
	MAX_SINGLE_BOARD_EXCEPTION_NUM = 16 //单板最大并发异常数

	MAX_VARIABLE_DATA_NUM = 65535 //最大可变数据个数

	ACS_PARAM_DOOR_STATUS_WEEK_PLAN     = 0x00000001 //门状态周计划参数
	ACS_PARAM_VERIFY_WEEK_PALN          = 0x00000002 //读卡器周计划参数
	ACS_PARAM_CARD_RIGHT_WEEK_PLAN      = 0x00000004 //卡权限周计划参数
	ACS_PARAM_DOOR_STATUS_HOLIDAY_PLAN  = 0x00000008 //门状态假日计划参数
	ACS_PARAM_VERIFY_HOLIDAY_PALN       = 0x00000010 //读卡器假日计划参数
	ACS_PARAM_CARD_RIGHT_HOLIDAY_PLAN   = 0x00000020 //卡权限假日计划参数
	ACS_PARAM_DOOR_STATUS_HOLIDAY_GROUP = 0x00000040 //门状态假日组参数
	ACS_PARAM_VERIFY_HOLIDAY_GROUP      = 0x00000080 //读卡器验证方式假日组参数
	ACS_PARAM_CARD_RIGHT_HOLIDAY_GROUP  = 0x00000100 //卡权限假日组参数
	ACS_PARAM_DOOR_STATUS_PLAN_TEMPLATE = 0x00000200 //门状态计划模板参数
	ACS_PARAM_VERIFY_PALN_TEMPLATE      = 0x00000400 //读卡器验证方式计划模板参数
	ACS_PARAM_CARD_RIGHT_PALN_TEMPLATE  = 0x00000800 //卡权限计划模板参数
	ACS_PARAM_CARD                      = 0x00001000 //卡参数
	ACS_PARAM_GROUP                     = 0x00002000 //群组参数
	ACS_PARAM_ANTI_SNEAK_CFG            = 0x00004000 //反潜回参数
	ACS_PAPAM_EVENT_CARD_LINKAGE        = 0x00008000 //事件及卡号联动参数
	ACS_PAPAM_CARD_PASSWD_CFG           = 0x00010000 //密码开门使能参数
	ACS_PARAM_PERSON_STATISTICS_CFG     = 0x00020000 //人数统计参数
	ACS_PARAM_BLACK_LIST_PICTURE        = 0x00040000 //黑名单图片参数
	ACS_PARAM_ID_BLACK_LIST             = 0x00080000 //身份证黑名单参数
	ACS_PARAM_EXAM_INFO                 = 0x00100000 //考试信息参数
	ACS_PARAM_EXAMINEE_INFO             = 0x00200000 //考生信息参数
	ACS_PARAM_FAILED_FACE_INFO          = 0x00400000 //升级设备人脸建模失败记录

	UNLOCK_PASSWORD_LEN       = 8  //解除密码长度
	LOCAL_CONTROLLER_NAME_LEN = 32 //就地控制器名称长度

	MAX_SERVER_DEVICE_NUMBER = 16 //最大设备数量

	MAX_SCREEN_ADDRESS_LEN = 16 //特征码最大长度
	MAX_DAY_TIME_POINT     = 8  //每天最大时间点个数
	MAX_TIME_POINT         = 16 //每年最大时间点个数

	MAX_CHECK_485CHAN = 64 //485自检设备通道号个数

	MAX_UPLOADFILE_URL_LEN = 240

	MAX_DIALNUM_LENGTH = 32

	MAX_LENGTH_32 = 32

	MAX_FILE_PATH_LEN = 256 //文件路径长度

	CARD_PARAM_CARD_VALID    = 0x00000001 //卡是否有效参数
	CARD_PARAM_VALID         = 0x00000002 //有效期参数
	CARD_PARAM_CARD_TYPE     = 0x00000004 //卡类型参数
	CARD_PARAM_DOOR_RIGHT    = 0x00000008 //门权限参数
	CARD_PARAM_LEADER_CARD   = 0x00000010 //首卡参数
	CARD_PARAM_SWIPE_NUM     = 0x00000020 //最大刷卡次数参数
	CARD_PARAM_GROUP         = 0x00000040 //所属群组参数
	CARD_PARAM_PASSWORD      = 0x00000080 //卡密码参数
	CARD_PARAM_RIGHT_PLAN    = 0x00000100 //卡权限计划参数
	CARD_PARAM_SWIPED_NUM    = 0x00000200 //已刷卡次数
	CARD_PARAM_EMPLOYEE_NO   = 0x00000400 //工号
	CARD_PARAM_NAME          = 0x00000800 //姓名
	CARD_PARAM_DEPARTMENT_NO = 0x00001000 //部门编号
	CARD_SCHEDULE_PLAN_NO    = 0x00002000 //排班计划编号
	CARD_SCHEDULE_PLAN_TYPE  = 0x00004000 //排班计划类型
	CARD_ROOM_NUMBER         = 0x00008000 //房间号
	CARD_SIM_NO              = 0x00010000 //SIM卡号（手机号）
	CARD_FLOOR_NUMBER        = 0x00020000 //楼层号
	CARD_USER_TYPE           = 0x00040000 //用户类型

	JUDGE_MAX_VIDEOOUT_NUM = 9

	MAX_IR_CMD_NAME_LEN = 32 //红外输出命令名称长度
	MAX_IR_CMD_NUM      = 32 //红外命令个数

	MAX_VIDEOIN_TYPE_NUM = 10 //最大支持的视频输入源类型

	MICROPHONE_NUM  = 16
	FAN_NUM         = 8
	FPGA_NUM        = 8
	MAIN_BOARD      = 8
	LOCAL_INPUT_NUM = 24
	LAMP_STATE_NAME = 32
	LAMP_NAME       = 32
	FILE_NAME_LEN   = 32

	DPC_CORRECT             = 1  //校正
	DPC_CORRECT_CANCEL      = 2  //取消校正
	DPC_CROSS_DISPALY_OPEN  = 3  //坏点检测十字叉显示开启
	DPC_CROSS_DISPALY_CLOSE = 4  //坏点检测十字叉显示关闭
	DPC_POINT               = 5  //坏点校正坐标
	DPC_UP                  = 6  //坏点校正坐标点向上偏移
	DPC_DOWN                = 7  //坏点校正坐标点向下偏移
	DPC_RIGHT               = 8  //坏点校正坐标点向右偏移
	DPC_LEFT                = 9  //坏点校正坐标点向左偏移
	DPC_ALL_CORRECT         = 10 //所有坏点校正
	DPC_SAVE                = 11 //坏点保存

	MAX_OSD_LEN = 64 //输出口OSD长度

	MAX_AUDIOOUT_PRO_TYPE = 8 //音频输出处理方式

	MAX_CENTERNUM_V40 = 6 //报警中心地址个数

	MAX_FINGER_PRINT_LEN = 768 //最大指纹长度

	MAX_ZONE_LINKAGE_CHAN_NUM = 4 /* 防区关联最大通道数*/

	MAX_MASK_AREA_NUM = 8 //马赛克区域个数

	MAX_PLAYLIST_NUM = 50 //最大播放列表数目
	MAX_PLAYPLAN_NUM = 50 //最大播放计划数目

	MAX_LEN_256          = 256
	MAX_GROUP_RECORD_NUM = 10 //最大记录个数

	MAX_MATRIX_SUBBOARD_NUM     = 16 //综合平台最大子板数
	MAX_MATRIX_SUBBOARD_NUM_V51 = 32 //综合平台最大子板数
	MAX_BOARD_SUBSYSTEM_NUM     = 12 //每个子板最大系统数

	MAX_FIRESHIELDMASK_REGION = 24

	MAX_SMOKESHIELDMASK_REGION = 24

	MAX_SIGNAL_JOINT_NUM = 64 //最大的拼接规模

	MAX_WHITELIST_PHONE_NUM = 16

	PRIOR_SCHEDTIME = 30

	MAX_SSID_LEN      = 32 //SSID号长度
	MAX_WS_PASSWD_LEN = 64 //密码长度

	NET_SDK_MAX_FILE_NAME = 100 //最大文件名称

	CLIENT_VERSION_LEN = 64

	MAX_OPTICAL_DEV_NODE = 32 //最多节点光端机数

	MAX_ID_NUM_LEN               = 32  //最大身份证号长度
	MAX_ID_NAME_LEN              = 128 //最大姓名长度
	MAX_ID_ADDR_LEN              = 280 //最大住址长度
	MAX_ID_ISSUING_AUTHORITY_LEN = 128 //最大签发机关长度

	MAX_PROXY_COUNT = 32

	COURSE_NAME_LEN_128    = 128 //课程名称
	INSTRUCTOR_NAME_LEN_64 = 64  //授课教师
	MAX_UUID_LEN           = 64  //手动课程UUID

	/******************************安全防范视频监控联网信息安全 end******************************************/
	MAX_RING_NAME_LEN_128 = 128 //铃音名称长度

	NET_SDK_MAX_INDENTITY_KEY_LEN = 64 //交互操作口令长度

	NET_SDK_MAX_EXAM_ROUND_NO    = 64 //考试场次编号最大长度
	NET_SDK_MAX_EXAM_NO          = 64 //考试编号最大长度
	NET_SDK_MAX_EXAM_SUBJECT     = 64 //考试科目最大长度
	NET_SDK_MAX_TEACHER_NO       = 64 //监考老师编号最大长度
	NET_SDK_MAX_TEACHER_NAME     = 64 //监考老师姓名最大长度
	NET_SDK_MAX_EXAMINEE_NO      = 64 //考生编号最大长度
	NET_SDK_MAX_ADMISSION_TICKET = 64 //准考证号最大长度

	MAX_FILE_NAME_LEN = 100 //最大文件名长

	AUDIO_FILE_NAME_LEN = 32 //音频文件名称长度

	EZVIZ_CLASSSESSION_LEN = 64
	EZVIZ_DEVICEID_LEN     = 32

	EZVIZ_REQURL_LEN      = 64
	EZVIZ_ACCESSTOKEN_LEN = 128
	EZVIZ_CLIENTTYPE_LEN  = 32
	EZVIZ_FEATURECODE_LEN = 64
	EZVIZ_OSVERSION_LEN   = 32
	EZVIZ_NETTYPE_LEN     = 32
	EZVIZ_SDKVERSION_LEN  = 32
	EZVIZ_APPID_LEN       = 64

	NET_DVR_SHOWLOGO = 1 /*显示LOGO*/
	NET_DVR_HIDELOGO = 2 /*隐藏LOGO*/

	DISP_CMD_ENLARGE_WINDOW = 1 /*显示通道放大某个窗口*/
	DISP_CMD_RENEW_WINDOW   = 2 /*显示通道窗口还原*/
	DISP_CMD_OPENAUDIO      = 3 /*显示通道打开音频*/
	DISP_CMD_CLOSEAUDIO     = 4 /*显示通道关闭音频*/

	// 标定校验
	NET_DVR_PDC_VERIFY_CALIBRATION      = 1 //当值为1是为PDC标定校验 pdc传入值为NET_VCA_POINT   传出值为 NET_VCA_RECT结构
	NET_DVR_VERIFY_BEHAVIOR_CALIBRATION = 2 // 行为分析标定线校验
	NET_DVR_VERIFY_ITS_CALIBRATION      = 3 // 智能交通标定校验
	NET_DVR_VERIFY_BV_CALIBRATION       = 5 //双目标定校验
)
