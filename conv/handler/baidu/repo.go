package baidu

import (
	"crypto/md5"
	"encoding/hex"
)

type ConvResp struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

// GetSign 获取签名
func (h *Handler) GetSign(convURL, aKey string) string {
	sum := md5.Sum([]byte(convURL + aKey))

	return hex.EncodeToString(sum[:])
}
