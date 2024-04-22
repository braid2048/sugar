package types

// ConvReq 回传入参
type ConvReq struct {
	BaseParams     *BaseConv     `json:"base_params" structs:"base_params"`
	OcanParams     *OcanConv     `json:"ocan_params" structs:"ocan_params"`
	KuaishouParams *KuaiShouConv `json:"kuaishou_params" structs:"kuaishou_params"`
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
