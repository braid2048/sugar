### 渠道回传包：

##### 目录结构：

- entry         ------   入口文件

- types        ------   常量与结构体

- handler     ------   各渠道执行回传的工厂

##### 执行流程：

1. 先构造通用请求体： types.ConvReq

2. 根据渠道创建实例：NewChannelHandler（channel）

3. 执行实例的 DoConv（&types.ConvReq）进行回传

##### 目前支持的渠道：

- ###### 巨量

  ```
  // step1. 构造请求参数
  convReq := &types.ConvReq{
  		BaseParams: &types.BaseConv{		// ---- 基础参数都是必传
  			PID:     "test_pid_01",
  			AdID:    "test_adid_01",
  			Channel: "ocan",
  			Brand:   "huawei",
  			Ip:      "127.0.0.1",
  		},
  		OcanParams: &types.OcanConv{		// ---- 巨量回传参数，必传
  			CallBack:  "test_clickid",    // -------- 回传的callback ,对应clickid 
  			ConvEvent: "active",					// -------- 回传的事件，对应巨量提供的文档key
  		},
  	}
  // step2. 获取回传工厂的实例
  convH, err := conv.NewChannelHandler("ocan")
  // step3. 调用回传
  convRes, err := convH.DoConv(ctx, convReq)
  ```



- ###### 快手

  ```
  // step1. 构造请求参数
  convReq := &types.ConvReq{
  		BaseParams: &types.BaseConv{		// ---- 基础参数都是必传
  			PID:     "test_pid_01",
  			AdID:    "test_adid_01",
  			Channel: "kuaishou",
  			Brand:   "huawei",
  			Ip:      "127.0.0.1",
  		},
  		KuaishouParams: &types.KuaiShouConv{		// ---- 快手回传参数，必传
  			CallBack:  "test_clickid",    				// -------- 回传的callback ,对应clickid 
  		},
  	}
  // step2. 获取回传工厂的实例
  convH, err := conv.NewChannelHandler("kuaishou")
  // step3. 调用回传
  convRes, err := convH.DoConv(ctx, convReq)
  ```



- ###### 华为

  ```
  // step1. 构造请求参数
  convReq := &types.ConvReq{
  		BaseParams: &types.BaseConv{		// ---- 基础参数都是必传
  			PID:     "test_pid_01",
  			AdID:    "test_adid_01",
  			Channel: "huawei",
  			Brand:   "huawei",
  			Ip:      "127.0.0.1",
  		},
  		HuaweiParams: &types.HuaWeiConv{		// ---- 华为回传参数，必传
  			OAID:                "test_oaid",		// -------- oaid原值
  			ConversionType:      "custom_landingpage", // -------- 转化事件
  			ContentID:           "test_u_did", // --------- 素材id,对应u_did
  			Callback:            "test_clickid", // -------- 回传的callback ,对应clickid 
  			CampaignID:          "test_adid", // -------- 计划id，对应adid
  			ConversionSecretKey: "xxxxx",  // -------- 转化秘钥
  		},
  	}
  // step2. 获取回传工厂的实例
  convH, err := conv.NewChannelHandler("huawei")
  // step3. 调用回传
  convRes, err := convH.DoConv(ctx, convReq)
  ```



- ###### oppoH5

  ```
  // step1. 构造请求参数
  convReq := &types.ConvReq{
  		BaseParams: &types.BaseConv{		// ---- 基础参数都是必传
  			PID:     "test_pid_01",
  			AdID:    "test_adid_01",
  			Channel: "oppo",
  			Brand:   "huawei",
  			Ip:      "127.0.0.1",
  		},
  		OppoParams: &types.OppoConv{		// ---- oppoH5回传参数，必传
  			OwnerID:       12345, // -------- 广告主id
  			AppID:         "test_app_id", // -------- appid
  			AppKey:        "test_app_key",// -------- appkey
  			LbID:          "xxxx", // -------- 流量号
  			TransformType: 101, // -------- 事件类型枚举，目前我们的都是101
  			PageID：       123456, // -------- 落地页id
  			TID:           "xxxx", // -------- traceId
  		},
  	}
  // step2. 获取回传工厂的实例
  convH, err := conv.NewChannelHandler("oppo")
  // step3. 调用回传
  convRes, err := convH.DoConv(ctx, convReq)
  ```



##### 依赖包：

- github.com/bububa/oceanengine  ： 巨量SDK
- github.com/bububa/oppo-omni : oppoSDK
- github.com/fatih/structs ：结构体快捷转化工具