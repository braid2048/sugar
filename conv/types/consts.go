package types

/**
 * 常量集
 */

// 厂商
const (
	BrandHuaWei = "huawei"
	BrandOppo   = "oppo"
	BrandVivo   = "vivo"
	BrandXiaoMi = "xiaomi"
	BrandHonnr  = "honor"
)

// 渠道
const (
	ChannelOceanH5      = "ocan"         // 巨量，ocean 拼错了
	ChannelOppoH5       = "oppo"         // OPPO 营销联盟
	ChannelOppoH5InSite = "oppo_in_site" // OPPO 营销站内
	ChannelKWaiH5       = "kuaishou"     // 快手
	ChannelYouKu        = "youku"        // 优酷
	ChannelVivoH5       = "vivo"         // vivo 营销
	ChannelBaidu        = "baidu"        // 百度百青藤
	ChannelOppoHap      = "oppohap"      // Oppo Hap
	ChannelQTT          = "qutoutiao"    // 趣头条
	ChannelTencent      = "tencent"      // 腾讯优量汇
	ChannelHuawei       = "huawei"       // 华为 Ads
	ChannelWifi         = "wkanx"        // wifi万能钥匙
	ChannelUc           = "uc"           // uc 阿里汇川
	ChannelSigmob       = "sigmob"       // sigmob
	ChannelWeiBo        = "weibo"        // 微博
	ChannelHonor        = "honor"        // 荣耀
	ChannelMagic        = "magic"        // magic
	ChannelOctopus      = "octopus"      // 章鱼互动
)

// 请求类型
const (
	RequestTypeHttp = "HTTP"
	RequestTypeSDK  = "SDK"
)
