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
            EventType: "1",  // 事件编号
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

- ###### oppohap

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
          OppoHapParams: &types.OppoHapConv{
              Imei: "" ,                     // imei原值，与oaid二选1
              OUID: "xxxxx"  ,             // oaid原值，与imei二选1
              RequestID:"click_id",     // 非必传--请求id,这里可填写click_id
              ClientIp:"127.0.0.1",      // 非必传--ip
              Pkg:"xxx" ,       // 必传--快应用id
              DataType:1,      // 必传--转化事件枚举
              Channel:1 ,     // 必传--渠道枚举，0：其他 1：oppo 2:一加
              Type :2 ,       // 必传--加密类型，0：无加密，1:imeiMD5 2:oaidMD5
              AppType:3 ,      // 必传--应用类别，0：其他，1：应用，2：游戏，3：快应用
              AscribeType: 0 , // 必传--归因类型，0：oppo归因，1：广告主归因，2：助攻归因
              AdID: 1111,            // 必传--adid
              PayID:"xxx",          // 非必传--付费交易id
              CustomType: 0,  // 非必传--自定义目标类型
              PayAmount: 20  // 非必传--付费金额,单位分
          },
      }
  // step2. 获取回传工厂的实例
  convH, err := conv.NewChannelHandler("oppohap")
  // step3. 调用回传
  convRes, err := convH.DoConv(ctx, convReq)
  ```

- ###### 趣头条

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
          QTTParams: &types.QTTConv{	
              CallBack :"xxx", // 回传地址--clickid
              OP2:"xxx",     // 回传事件
              Arpu: 1000  // arpu值，单位是厘
          },
      }
  // step2. 获取回传工厂的实例
  convH, err := conv.NewChannelHandler("qutoutiao")
  // step3. 调用回传
  convRes, err := convH.DoConv(ctx, convReq)
  ```

- ###### vivo

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
          VivoParams: &types.VivoConv{	
              PageURL:"xxx",       // 落地页 URL | 当事件源类型为Web时,该字段必传
              SrcID:"xxx",       // 数据源id，营销平台转化管理工具中新建，每个产品在每个账号下仅可新建一个
              SrcType:"Quickapp",  // 事件源类型，枚举值：APP/Web/Quickapp/offline(不区分大小写)
              PackageName:"xxx",      // 应用包名, 当事件源类型为 APP/Quickapp 时,该字段必传
              UserIDType:"OAID",    // 用户标识类型, 枚举值IMEI/IMEI_MD5/OAID/OAID_MD5/OTHER/OPENID, 当事件源类型为APP时,该字段必传。
              UserID:"xxx",      // 标识的值，如IMEI号等,当事件源类型为APP时,该字段必传。IMEI：15-17位，明文 IMEI_MD5 ：32位，加密 OAID：64位，明文OAID_MD5：32位   加密OTHER：不限OPENID：不限
              CreativeID:"xx",   // vivo回传点击数据时,透传给广告主的creativeId,使用点击归因的广告主需要回传
              RequestID:"click_id",   // vivo回传点击数据时,透传给广告主的RequestID,使用点击归因的广告主需要回传。
              CvType:"xxx",    // 事件类型
              AccessToken:"xxx", // token
              AdvertiserID:"xxx",  // 投放账户id
          },
      }
  // step2. 获取回传工厂的实例
  convH, err := conv.NewChannelHandler("vivo")
  // step3. 调用回传
  convRes, err := convH.DoConv(ctx, convReq)
  ```

- ###### 腾讯

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
          TencentParams: &types.TencentConv{	
              OuterActionID:"xxx", //选填，若上报可能有重复请填写该id，系统会根据该ID进行去重
              ActionType:"xxx",  // 行为类型
              CallBack:"xxx",    // 回传地址
              HashIMEI:"xxx", // 监测的参数，选填
              HashOAID:"xxx", // 监测的参数，选填  
              HashAndroidID:"xxx",// 监测的参数，选填
          },
      }
  // step2. 获取回传工厂的实例
  convH, err := conv.NewChannelHandler("tencent")
  // step3. 调用回传
  convRes, err := convH.DoConv(ctx, convReq)
  ```

- ###### 百度

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
          BaiDuParams: &types.BaiduConv{	
             CallBack:"xxx", // 回传地址
             AType:"xx",    // 转化类型
             AValue:"0",  // 转化指标 ， 转化类型为付费时，该字段定为付费金额-单位(分) ，无转化金额回传时，数值填写为“0”即可
             Akey:"xx",        // 签名
          },
      }
  // step2. 获取回传工厂的实例
  convH, err := conv.NewChannelHandler("baidu")
  // step3. 调用回传
  convRes, err := convH.DoConv(ctx, convReq)
  ```

- ###### uc

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
          UCParams: &types.UcConv{	
             ConvURL:"xxx", // 回传链接
             ImeiSum:"xxx", // imei的md5
             OAID:"xxx",     // oaid原值
             Event:"xxx",    // 回传事件
          },
      }
  // step2. 获取回传工厂的实例
  convH, err := conv.NewChannelHandler("uc")
  // step3. 调用回传
  convRes, err := convH.DoConv(ctx, convReq)
  ```

- ###### 微博

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
          WeiBoParams: &types.WeiBoConv{
		            ConvType: 1, // 回传方式：1 快应用，2 落地页 必传;
		            QuickParams: &types.WeiBoQuickParams{ // 快应用回传时该参数必填
			            IMP:        "", // 监测callback;必传
			            ActionType: "", // 激活后的行为数据，3注册4付费;选填
			            Price:      0,  // 单位元，源文档是int型，貌似没有角分，付费事件的金额;选填
			            ActiveTime: 0,  // 行为时间，秒级时间戳，需要特别注意：有重传机制的广告主在重新回传需要保证active_time 完全一致，否则会被处理成多次激活;选填
		            },
		            LandParams: &types.WeiBoLandParams{ // 落地页回传时该参数必填
                      MarkID:   "", // mark_id允许为空，此时为自然量 需要urlencode一次，请不要多次urlencode
			            Behavior: "", // 行为码;必传
			            Time:     0,  // 转化时间,毫秒时间戳;必传
		            },
          },
  }
  // step2. 获取回传工厂的实例
  convH, err := conv.NewChannelHandler("weibo")
  // step3. 调用回传
  convRes, err := convH.DoConv(ctx, convReq)
  ```

- ###### wifi万能钥匙

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
          WifiParams: &types.WifiConv{
              Cid:       "xxx",  // 广告创意 ID ; 取值来⾃监测
              Sid:       "xxx",  // 广告检索 ID ; 必填 ; 取值来⾃监测
              STime:     "xxx",  // 广告检索时间 ; 取值来⾃监测
              Os:        "xxx",  // 移动端系统说明 ; 取值来⾃监测
              Idfa:      "xxx",  // iOS 设备的 IDFA 的 Md5 值 ; 取值来⾃监测
              Mac:       "xxx",  // 设备 MAC 的 Md5 值 ; 取值来⾃监测
              Imei:      "xxx",  // Android 设备 imei 的 Md5 值 ; 取值来⾃监测
              ClientId:  "xxx",  // 由 WiFi 万能钥匙分配
              EventType: "xxx",  // 1 激活 2 注册 3 付费 5 次留 6 关键行为 7 18天腊货 8 30天拉活 18 ⼩程序唤起（ 快应⽤类别回传到关键⾏为）
              Ts:        "xxx",  // 秒级时间戳
              SecretKey: "xxx",  // 签名的盐 ; 必填
          },
  }
  // step2. 获取回传工厂的实例
  convH, err := conv.NewChannelHandler("wkanx")
  // step3. 调用回传
  convRes, err := convH.DoConv(ctx, convReq)
  ```

- ###### honor

  ```
  // step1. 构造请求参数
  convReq := &types.ConvReq{
          BaseParams: &types.BaseConv{		// ---- 基础参数都是必传
              PID:     "test_pid_01",
              AdID:    "test_adid_01",
              Channel: "honor",
              Brand:   "huawei",
              Ip:      "127.0.0.1",
          },
          HonorParams: &types.HonorConv{
              TrackID:        "aa",  // 回传ID，取自监测，必传
		      ConversionID:   "bb",  // 转化事件类型id ， 必传
		      ConversionTime: 123,   // 转化时间，毫秒级时间戳，必传
		      AdvertiserID:   "cc",  // 广告主id，必传，取自监测
		      OaID:           "dd",  // oaid
		      Extra: &HonorExtra{    // 额外参数，注意目前只在事件是关键行为（900401）时必传
			      PkgName: "aa", // 包名
			      AppName: "bb", // 应用名
		      },
          },
  }
  // step2. 获取回传工厂的实例
  convH, err := conv.NewChannelHandler("honor")
  // step3. 调用回传
  convRes, err := convH.DoConv(ctx, convReq)
  ```

- ###### magic

  ```
  // step1. 构造请求参数
  convReq := &types.ConvReq{
          BaseParams: &types.BaseConv{		// ---- 基础参数都是必传
              PID:     "test_pid_01",
              AdID:    "test_adid_01",
              Channel: "honor",
              Brand:   "huawei",
              Ip:      "127.0.0.1",
          },
          MagicParams: &types.MagicConv{
              MgcCb:        "aa",  // 未解码的callback，客户端原值，必传
              Event:   "1",  // 转化事件，从pid配置中获取，必传
          },
  }
  // step2. 获取回传工厂的实例
  convH, err := conv.NewChannelHandler("magic")
  // step3. 调用回传
  convRes, err := convH.DoConv(ctx, convReq)
  ```

- ###### 章鱼互动

  ```
  // step1. 构造请求参数
  convReq := &types.ConvReq{
          BaseParams: &types.BaseConv{		// ---- 基础参数都是必传
              PID:     "test_pid_01",
              AdID:    "test_adid_01",
              Channel: "honor",
              Brand:   "huawei",
              Ip:      "127.0.0.1",
          },
          OctopusParams: &types.OctopusConv{
              Callback:        "aa",  // 回传地址，监测链接参数中获取，必传
              EventType:   "1001",  // 转化事件，从pid配置中获取，必传
              Timestamp: 1728546445000, // 毫秒级时间戳
          },
  }
  // step2. 获取回传工厂的实例
  convH, err := conv.NewChannelHandler("octopus")
  // step3. 调用回传
  convRes, err := convH.DoConv(ctx, convReq)
  ```
- ###### 喜马拉雅

  ```
  // step1. 构造请求参数
  convReq := &types.ConvReq{
          BaseParams: &types.BaseConv{		// ---- 基础参数都是必传
              PID:     "test_pid_01",
              AdID:    "test_adid_01",
              Channel: "honor",
              Brand:   "huawei",
              Ip:      "127.0.0.1",
          },
          XmlyParams: &types.XmlyConv{
              Callback:        "aa",  // 未解码的callback，监测链接获取原值，必传
              Type:   "act",  // 转化事件，从新pid回传配置中获取，必传
          },
  }
  // step2. 获取回传工厂的实例
  convH, err := conv.NewChannelHandler("xmly")
  // step3. 调用回传
  convRes, err := convH.DoConv(ctx, convReq)
  ```
- ###### 爱奇艺

  ```
  // step1. 构造请求参数
  convReq := &types.ConvReq{
          BaseParams: &types.BaseConv{		// ---- 基础参数都是必传
              PID:     "test_pid_01",
              AdID:    "test_adid_01",
              Channel: "honor",
              Brand:   "huawei",
              Ip:      "127.0.0.1",
          },
          IQiYiParams: &types.IQiYiConv{
              Callback:        "aa",  // 回传地址，监测链接参数中获取，必传
              EventType:   "1001",  // 转化事件，从pid配置中获取，必传
          },
  }
  // step2. 获取回传工厂的实例
  convH, err := conv.NewChannelHandler("iqiyi")
  // step3. 调用回传
  convRes, err := convH.DoConv(ctx, convReq)
  ```
  
- ###### 有道智选

  ```
  // step1. 构造请求参数
  convReq := &types.ConvReq{
          BaseParams: &types.BaseConv{		// ---- 基础参数都是必传
              PID:     "test_pid_01",
              AdID:    "test_adid_01",
              Channel: "honor",
              Brand:   "huawei",
              Ip:      "127.0.0.1",
          },
          YdzxParams: &types.YdzxConv{
              ConvExt:        "aa",  // 回传地址，监测链接参数中获取，必传
              ConvAction:   "xx",  // 转化事件，从pid配置中获取，必传
          },
  }
  // step2. 获取回传工厂的实例
  convH, err := conv.NewChannelHandler("ydzx")
  // step3. 调用回传
  convRes, err := convH.DoConv(ctx, convReq)
  ```

- ###### 必得

  ```
  // step1. 构造请求参数
  convReq := &types.ConvReq{
          BaseParams: &types.BaseConv{		// ---- 基础参数都是必传
              PID:     "test_pid_01",
              AdID:    "test_adid_01",
              Channel: "honor",
              Brand:   "huawei",
              Ip:      "127.0.0.1",
          },
          BideParams: &types.BideConv{
              Callback:        "aa",  // 回传地址，监测链接参数中获取，必传
              TransformType:   "xx",  // 转化事件，从pid配置中获取，必传
          },
  }
  // step2. 获取回传工厂的实例
  convH, err := conv.NewChannelHandler("bide")
  // step3. 调用回传
  convRes, err := convH.DoConv(ctx, convReq)
  ```
- ###### 支付宝（阿里灯火）

  ```
  // step1. 构造请求参数
  convReq := &types.ConvReq{
          BaseParams: &types.BaseConv{		// ---- 基础参数都是必传
              PID:     "test_pid_01",
              AdID:    "test_adid_01",
              Channel: "honor",
              Brand:   "huawei",
              Ip:      "127.0.0.1",
          },
          AlipayParams: &types.AlipayConv{
              AppID:        "aa",  // 支付宝分配给开发者的应用ID 必填
              PrivateKey:   "xx",  // 应用私钥 必填
              BizToken:     "xx",  // 业务令牌，访问灯火平台的token，必填
              PrincipalTag:     "xx", // 商家标签，必填
              BizNo:     "xx", // 转化流水号，由商家自主定义作为转化数据唯一标识 , 必填
              ConversionType:     "xx", // 转化事件类型，必填
              ConversionTime:     12345, // 转化时间，秒级时间戳
              CallbackExtInfo:  "xx", // callback , 监测获取，必填
          },
  }
  // step2. 获取回传工厂的实例
  convH, err := conv.NewChannelHandler("alipay")
  // step3. 调用回传
  convRes, err := convH.DoConv(ctx, convReq)
  ```


##### 依赖包：

- github.com/bububa/oceanengine  ： 巨量SDK
- github.com/bububa/oppo-omni : oppoSDK
- github.com/fatih/structs ：结构体快捷转化工具