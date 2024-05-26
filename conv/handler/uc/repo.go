package uc

import (
	"fmt"
	"net/url"
)

type ConvParams struct {
	Callback string `json:"conv_url"` // 回传链接
	ImeiSum  string `json:"imei_md5"` // imei的md5
	OAID     string `json:"oaid"`     // oaid原值
}

type ConvResp struct {
	Status int `json:"status"`
}

func (h *Handler) GetURL(convParams *ConvParams, event string) (string, error) {
	// NOTE: decode回传链接
	callbackURL, err := url.QueryUnescape(convParams.Callback)
	if err != nil {
		return "", fmt.Errorf("uc callback failed to decode, err: %w", err)
	}
	// NOTE: 拼接请求参数
	callbackURL += "&type=" + event + "&imei_sum=" + convParams.ImeiSum + "&oaid=" + convParams.OAID

	return callbackURL, nil
}
