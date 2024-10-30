package alipay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/fatih/structs"
	"sort"
	"strings"
)

const (
	OpenAPIURL         = "https://openapi.alipay.com/gateway.do"
	OpenAPIMethod      = "alipay.data.dataservice.ad.conversion.upload"
	OpenAPIFormat      = "JSON"
	OpenAPICharset     = "utf-8"
	OpenAPISignType    = "RSA2"
	OpenAPIVersion     = "1.0"
	ConversionSource   = "COMMON_TARGET"
	ConversionUuIDType = "PID"
	ConversionUuID     = "2088UID"
)

type ConvResp struct {
	Sign    string `json:"sign" structs:"sign"`
	RespCon struct {
		Code string `json:"code" structs:"code"`
		Msg  string `json:"msg" structs:"msg"`
	} `json:"alipay_data_dataservice_ad_conversion_upload_response" structs:"resp_con"`
}

type OpenAPIReq struct {
	AppID      string `json:"app_id" structs:"app_id"`           // 应用id
	Method     string `json:"method" structs:"method"`           // 调用的方法名
	Format     string `json:"format" structs:"format"`           // 仅支持‘JSON’
	Charset    string `json:"charset" structs:"charset"`         // 字符编码 utf-8
	SignType   string `json:"sign_type" structs:"sign_type"`     // 签名类型,推荐RSA2
	Sign       string `json:"sign" structs:"-"`                  // 签名
	Timestamp  string `json:"timestamp" structs:"timestamp"`     // 发送时间：“2006-01-02 15:04:05”
	Version    string `json:"version" structs:"version"`         // 固定 1.0
	BizContent string `json:"biz_content" structs:"biz_content"` // 业务参数
}

type BizContentReq struct {
	BizToken           string            `json:"biz_token" structs:"biz_token"`                       // 灯火token
	ConversionDataList []*ConversionData `json:"conversion_data_list" structs:"conversion_data_list"` // 转化数据参数
}

type ConversionData struct {
	Source          string              `json:"source" structs:"source"`                       // 来源，固定为 “COMMON_TARGET”
	PrincipalTag    string              `json:"principal_tag" structs:"principal_tag"`         // 商家标签
	BizNo           string              `json:"biz_no" structs:"biz_no"`                       // 转化流水号，商家定义的唯一标识
	ConversionType  string              `json:"conversion_type" structs:"conversion_type"`     // 转化事件类型
	PropertyList    []*PropertyListItem `json:"property_list" structs:"property_list"`         // 归因字段
	ConversionTime  int64               `json:"conversion_time" structs:"conversion_time"`     // 转化规则时间戳，秒级
	UuIDType        string              `json:"uuid_type" structs:"uuid_type"`                 // 枚举值，固定“PID”
	UuID            string              `json:"uuid" structs:"uuid"`                           // 转化用户，固定“2088UID”
	CallbackExtInfo string              `json:"callback_ext_info" structs:"callback_ext_info"` // callback
}

type PropertyListItem struct {
	Key   string `json:"key" structs:"key"`
	Value string `json:"value" structs:"value"`
}

func (o *OpenAPIReq) GetSignOfRSA2(privateKey string) (string, error) {
	// step1. 转化map,剔除sign字段
	m := structs.Map(o)
	// 提取key并排序
	var (
		sortKey       []string
		waitSignSlice []string
		waitSignStr   string
	)

	for k := range m {
		sortKey = append(sortKey, k)
	}

	sort.Strings(sortKey)
	// step2. 按排序拼接值与&
	for k := range sortKey {
		waitSignSlice = append(waitSignSlice, fmt.Sprintf("%v=%v", sortKey[k], m[sortKey[k]]))
	}

	waitSignStr = strings.Join(waitSignSlice, "&")
	// step3. 创建私钥对象
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", fmt.Errorf("无效的私钥格式")
	}
	parsedKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("解析私钥失败: %v", err)
	}
	// step4. 初始化签名函数
	hash := crypto.SHA256
	h := sha256.New()
	h.Write([]byte(waitSignStr))
	hashed := h.Sum(nil)
	// step5. 私钥进行签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, parsedKey, hash, hashed)
	if err != nil {
		return "", fmt.Errorf("签名生成失败: %v", err)
	}
	// step6. base64编码
	return base64.StdEncoding.EncodeToString(signature), nil
}
