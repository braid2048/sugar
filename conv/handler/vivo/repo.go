package vivo

import (
	"crypto/rand"
	"encoding/base64"
)

type ConvParams struct {
	DataList    []*DataItem `json:"dataList"`
	PageURL     string      `json:"pageUrl"` // 落地页 URL
	SrcID       string      `json:"srcId"`   // 数据源id，营销平台转化管理工具中新建，每个产品在每个账号下仅可新建一个
	SrcType     string      `json:"srcType"` // 数据源类型，一律使用 Web
	PackageName string      `json:"pkgName"` // 应用包名, 当事件源类型为 APP/Quickapp 时,该字段必传
}

type DataItem struct {
	CreativeID string `json:"creativeId"`
	RequestID  string `json:"requestId"`
	CvTime     int64  `json:"cvTime"` // 事件发生的系统时间戳，精确到毫秒，13位
	CvType     string `json:"cvType"` // 事件类型
}

type ConvResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (h *Handler) RandString(n int) string {
	rb := make([]byte, n)
	if _, err := rand.Read(rb); err != nil {
		panic(err)
	}

	return base64.URLEncoding.EncodeToString(rb)[:n]
}
