package types

// ConvReq 回传入参
type ConvReq struct {
	BaseParams     *BaseConv     `json:"base_params" structs:"base_params"`
	OcanParams     *OcanConv     `json:"ocan_params" structs:"ocan_params"`
	KuaishouParams *KuaiShouConv `json:"kuaishou_params" structs:"kuaishou_params"`
	HuaweiParams   *HuaWeiConv   `json:"huawei_params" structs:"huawei_params"`
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
	OAID                string `json:"oaid"`            // 设备标识符，明文，没有传空字符
	ConversionType      string `json:"conversion_type"` // 转化事件的类型，详细枚举值见附录3
	ContentID           string `json:"content_id"`      // 素材id，与该条转化行为匹配的、广告主接收到素材id
	Callback            string `json:"callback"`        // 与该条转化行为数据的、广告主接收到的事件中的callback参数，该参数是经过URL编码的
	CampaignID          string `json:"campaign_id"`     // 与该条转化行为匹配的、广告主接收到的事件中的计划id
	Timestamp           string `json:"timestamp"`       // 本请求发起的时间戳，Unix时间戳，单位毫秒
	ConversionTime      string `json:"conversion_time"`
	ConversionSecretKey string `json:"conversion_secret_key"` // 秘钥
}
