package huawei

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

type ConvParams struct {
	OAID           string `json:"oaid" structs:"oaid" `                        // 设备标识符，明文，没有传空字符
	ConversionType string `json:"conversion_type"  structs:"conversion_type" ` // 转化事件的类型，详细枚举值见附录3
	ContentID      string `json:"content_id"  structs:"content_id" `           // 素材id，与该条转化行为匹配的、广告主接收到素材id
	Callback       string `json:"callback"  structs:"callback" `               // 与该条转化行为数据的、广告主接收到的事件中的callback参数，该参数是经过URL编码的
	CampaignID     string `json:"campaign_id"  structs:"campaign_id" `         // 与该条转化行为匹配的、广告主接收到的事件中的计划id
	Timestamp      string `json:"timestamp"  structs:"timestamp" `             // 本请求发起的时间戳，Unix时间戳，单位毫秒
	ConversionTime string `json:"conversion_time"  structs:"conversion_time" `
}

type ConvResp struct {
	ResultCode    int    `json:"resultCode" structs:"resultCode"`
	ResultMessage string `json:"resultMessage" structs:"resultMessage"`
}

func (h *Handler) GetSign(body []byte, secretKey, timeS string) string {
	// step1. 加密
	salt := hmac.New(sha256.New, []byte(secretKey))
	salt.Write(body)

	return fmt.Sprintf("Digest validTime=\"%v\", response=\"%s\"", timeS, fmt.Sprintf("%x", salt.Sum(nil)))
}
