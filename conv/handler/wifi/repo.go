package wifi

import (
	"crypto/md5"
	"fmt"
	"github.com/fatih/structs"
	"sort"
	"strings"
)

type ConvParams struct {
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
	Sign      string `json:"sign" structs:"-"`                // 验签，大写
}

func (c *ConvParams) GetSignAndURL(secretKey string) (string, string) {
	// 参数转为map
	paramsMap := structs.Map(c)
	// 提取key, 用于排序
	var (
		keys     []string
		paramNew []string
	)

	for k := range paramsMap {
		keys = append(keys, k)
	}
	// key排序
	sort.Strings(keys)
	// 将k-v整合成切片
	for k := range keys {
		paramNew = append(paramNew, keys[k]+"="+paramsMap[keys[k]].(string))
	}
	// kv切片用&拼接成url参数
	URLParam := strings.Join(paramNew, "&")
	// kv参数字符串最后加盐，拼接成待签名字符串
	waitSignStr := URLParam + secretKey
	// md5加密
	hash := md5.Sum([]byte(waitSignStr))
	// 签名需要大写
	signStr := strings.ToUpper(fmt.Sprintf("%x", hash))
	// 拼接请求URL
	URL := convURL + "?" + URLParam + "&sign=" + signStr

	return signStr, URL
}
