package alipay

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/braid2048/sugar/conv/types"
	"github.com/braid2048/sugar/conv/utils"
	"github.com/fatih/structs"
	"net/http"
	"net/url"
	"time"
)

type Handler struct{}

type HandlerReq struct {
	Req *OpenAPIReq `json:"req" structs:"req"`
	URL string      `json:"url" structs:"url"`
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) DoConv(ctx context.Context, req *types.ConvReq) (*types.ConvRes, error) {
	// step1. 检验参数
	if err := h.Validate(req); err != nil {
		return nil, fmt.Errorf("alipay-conv -- Validate err: %w", err)
	}
	// step2. 构造参数
	convReq, err := h.MakeReq(req)
	if err != nil {
		return nil, fmt.Errorf("alipay-conv -- MakeReq err: %w", err)
	}
	// step3. 解析body
	body, err := json.Marshal(convReq.Req)
	if err != nil {
		return nil, fmt.Errorf("alipay-conv -- body json err: %w", err)
	}
	// step4. 回传
	respCode, respBody, err := utils.SendPostRequest(ctx, convReq.URL, map[string]string{"Content-Type": "application/json"}, body)
	if err != nil {
		return nil, fmt.Errorf("alipay-conv -- send request err: %w", err)
	}

	return h.MakeRes(respCode, respBody, req, convReq)
}

// Validate 检验参数
func (h *Handler) Validate(req *types.ConvReq) error {
	if err := req.Validate(); err != nil {
		return err
	}
	// NOTE: 必填字段
	if req.AlipayParams == nil {
		return fmt.Errorf("conv params is nil")
	}
	// NOTE: appid
	if req.AlipayParams.AppID == "" {
		return fmt.Errorf("app_id is nil")
	}
	// NOTE: private_key
	if req.AlipayParams.PrivateKey == "" {
		return fmt.Errorf("private_key is nil")
	}
	// NOTE: biz_token
	if req.AlipayParams.BizToken == "" {
		return fmt.Errorf("biz_token is nil")
	}
	// NOTE: principal_tag
	if req.AlipayParams.PrincipalTag == "" {
		return fmt.Errorf("principal_tag is nil")
	}
	// NOTE: biz_no
	if req.AlipayParams.BizNo == "" {
		return fmt.Errorf("biz_no is nil")
	}
	// NOTE: conversion_type
	if req.AlipayParams.ConversionType == "" {
		return fmt.Errorf("conversion_type is nil")
	}
	// NOTE: conversion_time
	if req.AlipayParams.ConversionTime <= 0 {
		return fmt.Errorf("conversion_time is nil")
	}
	// NOTE: callback_ext_info
	if req.AlipayParams.CallbackExtInfo == "" {
		return fmt.Errorf("callback_ext_info is nil")
	}

	return nil
}

// MakeReq 构造请求参数
func (h *Handler) MakeReq(req *types.ConvReq) (*HandlerReq, error) {
	// step1. 构造公共网关请求参数
	openAPIReq := &OpenAPIReq{
		AppID:     req.AlipayParams.AppID,
		Method:    OpenAPIMethod,
		Format:    OpenAPIFormat,
		Charset:   OpenAPICharset,
		SignType:  OpenAPISignType,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Version:   OpenAPIVersion,
	}
	// step2. 构造业务参数
	bizContentReq := &BizContentReq{
		BizToken: req.AlipayParams.BizToken,
		ConversionDataList: []*ConversionData{{
			BizNo:          req.AlipayParams.BizNo,
			ConversionTime: req.AlipayParams.ConversionTime,
			ConversionType: req.AlipayParams.ConversionType,
			PrincipalTag:   req.AlipayParams.PrincipalTag,
			PropertyList:   []*PropertyListItem{},
			Source:         ConversionSource,
			UuID:           ConversionUuID,
			UuIDType:       ConversionUuIDType,
		}},
	}
	// callback urlDecode
	callback, err := url.QueryUnescape(req.AlipayParams.CallbackExtInfo)
	if err != nil {
		return nil, fmt.Errorf("callback url decode err: %w", err)
	}

	bizContentReq.ConversionDataList[0].CallbackExtInfo = callback
	// step3. 将业务参数转为字符串并赋值在公共参数里，注意这里暂时将 PropertyList 参数置空，后期可能要写工厂
	bizContentStr, err := json.Marshal(bizContentReq)
	if err != nil {
		return nil, fmt.Errorf("json marshal err: %w", err)
	}

	openAPIReq.BizContent = string(bizContentStr)
	// step4. 加密
	openAPIReq.Sign, err = openAPIReq.GetSignOfRSA2(req.AlipayParams.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("get sign err: %w", err)
	}

	return &HandlerReq{Req: openAPIReq, URL: OpenAPIURL}, nil
}

// MakeRes 解析响应
func (h *Handler) MakeRes(respCode int, resBody []byte, req *types.ConvReq, HandlerReq *HandlerReq) (*types.ConvRes, error) {
	// step1. 初始化响应体
	res := &types.ConvRes{
		Channel: req.BaseParams.Channel,
		Request: &types.ChannelRequestData{
			ReqType: types.RequestTypeHttp,
			ReqData: structs.Map(HandlerReq),
		},
		Response: &types.ChannelResponseData{
			StatusCode: respCode,
			ResData:    string(resBody),
		},
	}
	// step2. 判断响应码
	if respCode >= http.StatusBadRequest {
		res.IsSuccess = false

		return res, nil
	}
	// step3. 解析响应
	result := &ConvResp{}

	if err := json.Unmarshal(resBody, result); err != nil {
		return nil, fmt.Errorf("MakeRes json err: %w", err)
	}
	// step4. 判断成功
	res.IsSuccess = result.RespCon.Code == "10000"
	// NOTE: 错误消息
	if !res.IsSuccess {
		res.Err = fmt.Errorf(result.RespCon.Msg)
	}

	return res, nil
}
