package types

// ConvReq 回传入参
type ConvReq struct {
	ProjectID string `structs:"pid" json:"pid"`               // 渠道号
	Brand     string `structs:"brand" json:"brand"`           // 厂商
	Channel   string `structs:"channel" json:"channel"`       // 渠道
	AdID      string `structs:"adid" json:"adid"`             // 渠道
	ClickID   string `structs:"clickid" json:"clickid"`       // clickid
	ConvEvent string `structs:"conv_event" json:"conv_event"` // 回传事件
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
