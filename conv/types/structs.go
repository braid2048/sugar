package types

import oceanreq "github.com/bububa/oceanengine/marketing-api/model/conversion"

// ConvReq 回传入参
type ConvReq struct {
	BaseParams     *BaseConv     `json:"base_params" structs:"base_params"`
	OcanParams     *OcanConv     `json:"ocan_params" structs:"ocan_params"`
	KuaishouParams *KuaiShouConv `json:"kuaishou_params" structs:"kuaishou_params"`
	HuaweiParams   *HuaWeiConv   `json:"huawei_params" structs:"huawei_params"`
	OppoParams     *OppoConv     `json:"oppo_params" structs:"oppo_params"`
	OppoHapParams  *OppoHapConv  `json:"oppo_hap_params" structs:"oppo_hap_params"`
	VivoParams     *VivoConv     `json:"vivo_params" structs:"vivo_params"`
	QTTParams      *QTTConv      `json:"qtt_params" structs:"qtt_params"`
	TencentParams  *TencentConv  `json:"tencent_params" structs:"tencent_params"`
	BaiDuParams    *BaiduConv    `json:"baidu_params" structs:"baidu_params"`
	UCParams       *UcConv       `json:"uc_params" structs:"uc_params"`
	WeiBoParams    *WeiBoConv    `json:"weibo_params" structs:"weibo_params"`
	HonorParams    *HonorConv    `json:"honor_params" structs:"honor_params"`
	WifiParams     *WifiConv     `json:"wifi_params" structs:"wifi_params"`
	MagicParams    *MagicConv    `json:"magic_params" structs:"magic_params"`
	OctopusParams  *OctopusConv  `json:"octopus_params" structs:"octopus_params"`
	XmlyParams     *XmlyConv     `json:"xmly_params" structs:"xmly_params"`
	IQiYiParams    *IQiYiConv    `json:"iqiyi_params" structs:"iqiyi_params"`
}

// ChannelRequestData 渠道请求数据
type ChannelRequestData struct {
	ReqType string
	ReqData interface{}
}

// ChannelResponseData 渠道回传响应
type ChannelResponseData struct {
	StatusCode int
	ResData    interface{}
}

// ConvRes 回传结果
type ConvRes struct {
	IsSuccess bool
	Channel   string
	Request   *ChannelRequestData
	Response  *ChannelResponseData
	Err       error
}

// BaseConv 基础数据
type BaseConv struct {
	PID     string `structs:"pid" json:"pid"`         // 渠道号
	Brand   string `structs:"brand" json:"brand"`     // 厂商
	Channel string `structs:"channel" json:"channel"` // 渠道
	AdID    string `structs:"adid" json:"adid"`       // adid
	Ip      string `structs:"ip" json:"ip"`           // ip
}

// OcanConv 巨量
type OcanConv struct {
	CallBack   string               `structs:"callback" json:"callback"`     // 回传地址--clickid
	ConvEvent  string               `structs:"conv_event" json:"conv_event"` // 回传事件
	Properties *oceanreq.Properties `structs:"properties" json:"properties"` // 额外参数，选传
}

// KuaiShouConv 快手
type KuaiShouConv struct {
	CallBack  string `structs:"call_back" json:"call_back"`   // 回传地址--clickid
	EventType string `structs:"event_type" json:"event_type"` // 回传事件
}

// HuaWeiConv 华为
type HuaWeiConv struct {
	OAID                string `json:"oaid" structs:"oaid" `                        // 设备标识符，明文，没有传空字符
	ConversionType      string `json:"conversion_type"  structs:"conversion_type" ` // 转化事件的类型，详细枚举值见附录3
	ContentID           string `json:"content_id"  structs:"content_id" `           // 素材id，与该条转化行为匹配的、广告主接收到素材id
	Callback            string `json:"callback"  structs:"callback" `               // 与该条转化行为数据的、广告主接收到的事件中的callback参数，该参数是经过URL编码的
	CampaignID          string `json:"campaign_id"  structs:"campaign_id" `         // 与该条转化行为匹配的、广告主接收到的事件中的计划id
	Timestamp           string `json:"timestamp"  structs:"timestamp" `             // 本请求发起的时间戳，Unix时间戳，单位毫秒
	ConversionTime      string `json:"conversion_time"  structs:"conversion_time" `
	ConversionSecretKey string `json:"conversion_secret_key"  structs:"conversion_secret_key" ` // 秘钥
}

// OppoConv oppo
type OppoConv struct {
	OwnerID       uint64 `json:"ownerId" structs:"ownerId"`             // 广告主 ID
	AppID         string `json:"appId" structs:"appId"`                 // APP ID
	AppKey        string `json:"appKey" structs:"appKey"`               // APP Key
	PageID        uint64 `json:"pageId" structs:"pageId"`               // 落地页id
	TID           string `json:"tid" structs:"tid"`                     // traceId
	LbID          string `json:"lbid" structs:"lbid"`                   // 流量号
	TransformType int    `json:"transformType" structs:"transformType"` // 转化类型
}

// OppoHapConv oppoHap
type OppoHapConv struct {
	Imei        string `json:"imei" structs:"imei"`               // imei原值，与oaid二选1
	OUID        string `json:"ouId" structs:"ouId"`               // oaid原值，与imei二选1
	RequestID   string `json:"requestId" structs:"requestId"`     // 非必传--请求id
	ClientIp    string `json:"clientIp" structs:"clientIp"`       // 非必传--ip
	Pkg         string `json:"pkg" structs:"pkg"`                 // 必传--快应用id
	DataType    int    `json:"dataType" structs:"dataType"`       // 必传--转化事件枚举
	Channel     int    `json:"channel" structs:"channel"`         // 必传--渠道枚举，0：其他 1：oppo 2:一加
	Type        int    `json:"type" structs:"type"`               // 必传--加密类型，0：无加密，1:imeiMD5 2:oaidMD5
	AppType     int    `json:"appType" structs:"appType"`         // 必传--应用类别，0：其他，1：应用，2：游戏，3：快应用
	AscribeType int    `json:"ascribeType" structs:"ascribeType"` // 必传--归因类型，0：oppo归因，1：广告主归因，2：助攻归因
	AdID        int64  `json:"adId" structs:"adId"`               // 必传--adid
	PayID       string `json:"payId" structs:"payId"`             // 非必传--付费交易id
	CustomType  int    `json:"customType" structs:"customType"`   // 非必传--自定义目标类型
	PayAmount   int64  `json:"payAmount" structs:"payAmount"`     // 非必传--付费金额,单位分
}

// VivoConv vivo
type VivoConv struct {
	PageURL      string `json:"pageUrl"  structs:"pageUrl"`           // 落地页 URL | 当事件源类型为Web时,该字段必传
	SrcID        string `json:"srcId"  structs:"srcId"`               // 数据源id，营销平台转化管理工具中新建，每个产品在每个账号下仅可新建一个
	SrcType      string `json:"srcType"  structs:"srcType"`           // 事件源类型，枚举值：APP/Web/Quickapp/offline(不区分大小写)
	PackageName  string `json:"pkgName"  structs:"pkgName"`           // 应用包名, 当事件源类型为 APP/Quickapp 时,该字段必传
	UserIDType   string `json:"userIdType"  structs:"userIdType"`     // 用户标识类型, 枚举值IMEI/IMEI_MD5/OAID/OAID_MD5/OTHER/OPENID, 当事件源类型为APP时,该字段必传。
	UserID       string `json:"userId"  structs:"userId"`             // 标识的值，如IMEI号等,当事件源类型为APP时,该字段必传。IMEI：15-17位，明文 IMEI_MD5 ：32位，加密 OAID：64位，明文OAID_MD5：32位   加密OTHER：不限OPENID：不限
	CreativeID   string `json:"creativeId"  structs:"creativeId"`     // vivo回传点击数据时,透传给广告主的creativeId,使用点击归因的广告主需要回传
	RequestID    string `json:"requestId"  structs:"requestId"`       // vivo回传点击数据时,透传给广告主的RequestID,使用点击归因的广告主需要回传。
	CvType       string `json:"cvType"  structs:"cvType"`             // 事件类型
	AccessToken  string `json:"accessToken"  structs:"accessToken"`   // token
	AdvertiserID string `json:"advertiserId"  structs:"advertiserId"` // 投放账户id
}

// QTTConv 趣头条
type QTTConv struct {
	CallBack string `structs:"callback" json:"callback"` // 回传地址--clickid
	OP2      string `structs:"op2" json:"op2"`           // 回传事件
	Arpu     int64  `structs:"arpu" json:"arpu"`         // arpu值，单位是厘
}

// TencentConv 腾讯
type TencentConv struct {
	OuterActionID string `json:"outer_action_id" structs:"outer_action_id"` //选填，若上报可能有重复请填写该id，系统会根据该ID进行去重
	ActionType    string `json:"action_type" structs:"action_type"`         // 行为类型
	CallBack      string `structs:"callback" json:"callback"`               // 回传地址
	HashIMEI      string `json:"hash_imei"  structs:"hash_imei"`
	HashOAID      string `json:"hash_oaid"  structs:"hash_oaid"`
	HashAndroidID string `json:"hash_android_id"  structs:"hash_android_id"`
}

// BaiduConv 百度
type BaiduConv struct {
	CallBack string `structs:"callback" json:"callback"` // 回传地址
	AType    string `json:"a_type"  structs:"a_type"`    // 转化类型
	AValue   string `json:"a_value"  structs:"a_value"`  // 转化指标 ， 转化类型为付费时，该字段定为付费金额-单位(分) ，无转化金额回传时，数值填写为“0”即可
	Akey     string `json:"akey"  structs:"akey"`        // 签名
}

// UcConv uc阿里汇川
type UcConv struct {
	ConvURL string `json:"conv_url" structs:"conv_url"` // 回传链接
	ImeiSum string `json:"imei_md5" structs:"imei_md5"` // imei的md5
	OAID    string `json:"oaid" structs:"oaid"`         // oaid原值
	Event   string `json:"event" structs:"event"`       // 回传事件
}

// WeiBoConv 微博
type WeiBoConv struct {
	ConvType    int               `json:"conv_type" structs:"conv_type"`       // 回传方式，必传：1.快应用回传；2.落地页回传
	QuickParams *WeiBoQuickParams `json:"quick_params" structs:"quick_params"` // 快应用回传参数
	LandParams  *WeiBoLandParams  `json:"land_params" structs:"land_params"`   // 落地页回传参数
}

// WeiBoQuickParams 微博快应用回传参数
type WeiBoQuickParams struct {
	IMP        string `json:"imp" structs:"imp"`                 // 监测callback;必传
	ActionType string `json:"action_type" structs:"action_type"` // 激活后的行为数据，3注册4付费;选填
	Price      int64  `json:"price" structs:"price"`             // 单位元，源文档是int型，貌似没有角分，付费事件的金额;选填
	ActiveTime int64  `json:"active_time" structs:"active_time"` // 行为时间，秒级时间戳，需要特别注意：有重传机制的广告主在重新回传需要保证active_time 完全一致，否则会被处理成多次激活;选填
}

// WeiBoLandParams 微博落地页回传参数
type WeiBoLandParams struct {
	Time     int64  `json:"time" structs:"time"`         // 转化时间,毫秒时间戳;必传
	MarkID   string `json:"mark_id" structs:"mark_id"`   // mark_id允许为空，此时为自然量 需要urlencode一次，请不要多次urlencode
	Behavior string `json:"behavior" structs:"behavior"` // 行为码;必传
}

// HonorConv 荣耀
type HonorConv struct {
	TrackID        string      `json:"trackId" structs:"trackId"`               // 回传ID，取自监测
	ConversionID   string      `json:"conversionId" structs:"conversionId"`     // 转化类型id
	ConversionTime int64       `json:"conversionTime" structs:"conversionTime"` // 毫秒级时间戳
	AdvertiserID   string      `json:"advertiserId" structs:"advertiserId"`     // 广告主id，取自监测
	OaID           string      `json:"oaid" structs:"oaid"`
	Extra          *HonorExtra `json:"extra" structs:"extra"` // 额外参数，当事件是900401关键行为时，这个字段必填
}

// HonorExtra 荣耀回传额外字段
type HonorExtra struct {
	PkgName string `json:"pkgName" structs:"pkgName"` // 包名
	AppName string `json:"appName" structs:"appName"` // 应用名称
}

type WifiConv struct {
	Cid       string `json:"cid" structs:"cid"`               // 广告创意 ID
	Sid       string `json:"sid" structs:"sid"`               // 广告检索 ID ; 必填
	STime     string `json:"stime" structs:"stime"`           // 广告检索时间
	Os        string `json:"os" structs:"os"`                 // 0 安卓 1 ios
	Idfa      string `json:"idfa" structs:"idfa"`             // iOS 设备的 IDFA 的 Md5 值
	Mac       string `json:"mac" structs:"mac"`               // 设备 MAC 的 Md5 值
	Imei      string `json:"imei" structs:"imei"`             // Android 设备 imei 的 Md5 值
	ClientId  string `json:"clientid" structs:"clientid"`     // 由 WiFi 万能钥匙分配
	EventType string `json:"event_type" structs:"event_type"` // 1 激活 2 注册 3 付费 5 次留 6 关键行为 7 18天腊货 8 30天拉活 18 ⼩程序唤起（ 快应⽤类别回传到关键⾏为）
	Ts        string `json:"ts" structs:"ts"`                 // 时间戳
	SecretKey string `json:"secret_key" structs:"secret_key"` // 签名的盐
}

// MagicConv magic
type MagicConv struct {
	MgcCb string `json:"mgc_cb" structs:"mgc_cb"` // callback
	Event string `json:"event" structs:"event"`   // 转化事件
}

// OctopusConv 章鱼互动
type OctopusConv struct {
	Callback  string `json:"callback" structs:"callback"`
	EventType string `json:"event_type" structs:"event_type"` // 回传事件
	Timestamp int64  `json:"timestamp" structs:"timestamp"`   // 毫秒级
}

// XmlyConv 喜马拉雅
type XmlyConv struct {
	Callback string `json:"callback" structs:"callback"` // callback
	Type     string `json:"type" structs:"type"`         // 转化事件
}

// IQiYiConv 爱奇艺
type IQiYiConv struct {
	Callback  string `json:"callback" structs:"callback"`     // callback
	EventType string `json:"event_type" structs:"event_type"` // 转化事件
}
