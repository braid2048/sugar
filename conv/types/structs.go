package types

// ConvReq 回传入参
type ConvReq struct {
	BaseParams     *BaseConv     `json:"base_params" structs:"base_params"`
	OcanParams     *OcanConv     `json:"ocan_params" structs:"ocan_params"`
	KuaishouParams *KuaiShouConv `json:"kuaishou_params" structs:"kuaishou_params"`
	HuaweiParams   *HuaWeiConv   `json:"huawei_params" structs:"huawei_params"`
	OppoParams     *OppoConv     `json:"oppo_params" structs:"oppo_params"`
	OppoHapParams  *OppoHapConv  `json:"oppo_hap_params" structs:"oppo_hap_params"`
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
	CallBack  string `structs:"call_back" json:"call_back"`   // 回传地址--clickid
	ConvEvent string `structs:"conv_event" json:"conv_event"` // 回传事件
}

// KuaiShouConv 快手
type KuaiShouConv struct {
	CallBack string `structs:"call_back" json:"call_back"` // 回传地址--clickid
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
