package test

import (
	"context"
	"fmt"
	"github.com/braid2048/sugar/conv"
	"github.com/braid2048/sugar/conv/types"
	"testing"
)

func TestAlipay(t *testing.T) {
	convReq := &types.ConvReq{
		BaseParams: &types.BaseConv{ // ---- 基础参数都是必传
			PID:     "test_pid_01",
			AdID:    "test_adid_01",
			Channel: "honor",
			Brand:   "huawei",
			Ip:      "127.0.0.1",
		},
		AlipayParams: &types.AlipayConv{
			AppID:           "2021003193608516", // 支付宝分配给开发者的应用ID 必填
			PrivateKey:      "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCOr31ktaIOe4/t+gy2pbd3/StpMjjCdz7QUZqkP0DASbfYTRHnE40rAwzBkVfFR5sDhG98bDJtGwEcR5Qv5BPMe8FqXzfrhUxY12p9aH7Y/R5pzojLNgdLcdvJ8sRKaQiM0OA05OyKAKNGUELvn6d7GPHRoNQSTEFooEfgz4pmyZodqP/5xjFdx2WnklMOenza2Bd+kZKhhg/STxTT4hrv8+QK6sZayNsSXF1HNm23d29J+sjIgKlmkwgyRsweGODH0zV0OnuhC0mVYbYeqOy1XUN0KbTxRrAObJ2QHo4IsKjxPqh5HT4nnCQu7y+T2csU0XoXzm1DVDE/UFpEjpbBAgMBAAECggEAd2uTVG3ck6iA/xlP+LJcCvX+lk3tSX1KOkqCVkOGaymH0GY4vKEVftFPpNaDkl7q85etQ4K+9agrVsLl2OtYUsWlTOUixWFMU/L/crS5rdyzzrSIsyOmGVPTM4OXG/wqpsjPThXnj10XSms0ip+iKhnnkp67wBDIkcvMPPEXEg0bLouVDJPlAyCp9avqIed1SzcJqBxlq+rmZ6eiVSpxYmn7yuqEpkmXyAQ3PWeej8y4MuqadVxJFc07S0TF6552cIJqCO1GtBT/EJS2uRe8zbE7iO95SBJDvopMYqEGbNHgJDq8jWAX3Sl7X+VvZBgWAyI4cBXrjd8lesKBcjPwjQKBgQDJZhm3siKlHtN6CpgTYdtoOQLm3A5XnzRnRWb6Sa5J8UAUMrYIPAs4XnkkK+ytDGjjnt8zB5OrsSiEVSYoc7T/vjUHKeNzXp9ihf/JKWTRaj7hSWh5JTSExazV8D3MdNJaYADoqL+RXvz4I8pzLiJEtxDRRbwEBpAhHa5HAeYRTwKBgQC1XnILVvArxmaciRKTNWBrATNYsU6vArVWtpzLQSvY0CoETxPwYvlUHk3fL4VhepnKirMgegXjb8e2uEU4SSXA+pEfyH4d9dDOeHvm2rLAjUtqhnEUCOohT7GXhSgIgW+uQhwNnY5OUY40TX0NQ2RkdBHqzv1qQq0JwteyE60y7wKBgEgNzk+pXLnEoJZ+KdBtZ0kPdJlRy5PPsrjr3J7ZM5DizjErv3KMwNrm+eJWELQMx44ALgycvllj5YIK0L8SPoRs8Z3hf62sAcwG8u9ubtJ5d4u0brtA81w5OM/FxuZaOVP6GvkiPn9tA4Znj6vLqpj5AfxNPaoXCwO+Ebb31+8DAoGANPL+vnrCITWJ7XdDGgECRhsSn3kcLJHQ2SD1A43iPCkBq0Je4tYyTjGOsHLSMNQ3I998TiXxpCWVG64OX1FTmVRNnAbPcvW54R1hptMATqCxfMcFKkE0AUef5El2l40aSrh4Mi1mR00eA/z/XijnnUAZnwCRET2oAOqGSDHgZbcCgYEAvg32ZNoWYu1ywlcOoET0WEIaTGGgdYDC8jAB0qfLGeDwSR+jl+yqYhdysLGVT0Ewq+S52aLg4goA9tbXRkw4s36JwRn/hplHRibYrJ6WIhZFb73pY0+qPtY6S1FRjNpvZXUcCqva9GAp1sUFyVdRYY84ZccPdmcnzvUZK1TMI4c=",
			BizToken:        "5add1fe98201421b9ff152821af2c7a0",                                                                                                 // 业务令牌，访问灯火平台的token，必填
			PrincipalTag:    "2f054aa06dbb468fa4c3f87aeb6123f2",                                                                                                 // 商家标签，必填
			BizNo:           "test_biz_no",                                                                                                                      // 转化流水号，由商家自主定义作为转化数据唯一标识 , 必填
			ConversionType:  "140",                                                                                                                              // 转化事件类型，必填
			ConversionTime:  1730270543,                                                                                                                         // 转化时间，秒级时间戳
			CallbackExtInfo: "hAXLRVRoQHTgh2CQLhjfkX_0HR80F7ktEabHmKgD_XSSXPF2Pwv-a_pEAQVW2FS-6FCLShsXMMU4Ctx0X61swqcK4Qme3YdxzrjynZ_gYjQ8Iv4MNcM0d8oHKFiDWyOx", // callback , 监测获取，必填
		},
	}

	fmt.Println("--私钥长度", len(convReq.AlipayParams.PrivateKey))

	// step2. 获取回传工厂的实例
	convH, err := conv.NewChannelHandler("alipay")
	if err != nil {
		panic(err)
	}
	// step3. 调用回传
	convRes, err := convH.DoConv(context.Background(), convReq)
	if err != nil {
		panic(err)
	}

	fmt.Println(convRes)
}
