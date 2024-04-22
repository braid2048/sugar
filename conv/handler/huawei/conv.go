package huawei

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/halo2024/sugar/conv/types"
	"github.com/halo2024/sugar/conv/utils"
	"strconv"
	"time"
)

const (
	convURL = "https://ppscrowd-drcn.op.cloud.huawei.com/action-lib-track/hiad/v2/actionupload"
)

type Handler struct{}

type HandlerReq struct {
	Req       *ConvParams `json:"req" structs:"req"`
	Sign      string      `json:"sign" structs:"sign"`
	SecretKey string      `json:"secret_key" structs:"secret_key"`
}

func New() *Handler {
	return &Handler{}
}

/**
 * 执行回传
 * 内部流程三段式：
 * 1. 检验参数
 * 2. 参数转化
 * 3. 回传请求
 * 4. 构造响应
 */

func (h *Handler) DoConv(ctx context.Context, req *types.ConvReq) (*types.ConvRes, error) {
	// step1. 检验参数
	if err := h.Validate(req); err != nil {
		return nil, fmt.Errorf("huawei-conv -- Validate err: %w", err)
	}
	// step2. 构造参数
	convReq := h.MakeReq(req)
	// step3. 解析body
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("huawei-conv -- body json err: %w", err)
	}
	// step4. 加密
	sign := h.GetSign(body, convReq.SecretKey, convReq.Req.Timestamp)
	convReq.Sign = sign

	// step3. 回传
	respCode, respBody, err := utils.SendPostRequest(ctx, convURL, map[string]string{"Content-Type": "application/json", "Authorization": convReq.Sign}, body)
	if err != nil {
		return nil, fmt.Errorf("huawei-conv -- send request err: %w", err)
	}

	return h.MakeRes(respCode, respBody, req, convReq)
}

// Validate 检验参数
func (h *Handler) Validate(req *types.ConvReq) error {
	if err := req.Validate(); err != nil {
		return err
	}
	// NOTE: 必填字段
	if req.HuaweiParams == nil {
		return fmt.Errorf("conv params is nil")
	}
	// NOTE: oaid
	if req.HuaweiParams.OAID == "" {
		return fmt.Errorf("oaid is nil")
	}
	// NOTE: event
	if req.HuaweiParams.ConversionType == "" {
		return fmt.Errorf("conversion_type is nil")
	}
	// NOTE: secret_key
	if req.HuaweiParams.ConversionSecretKey == "" {
		return fmt.Errorf("conversion_secret_key is nil")
	}
	// NOTE: callback
	if req.HuaweiParams.Callback == "" {
		return fmt.Errorf("callback is nil")
	}
	// NOTE: campaign_id
	if req.HuaweiParams.CampaignID == "" {
		return fmt.Errorf("campaign_id is nil")
	}
	// NOTE: content_id
	if req.HuaweiParams.ContentID == "" {
		return fmt.Errorf("content_id is nil")
	}

	return nil
}

// MakeReq 构造请求参数
func (h *Handler) MakeReq(req *types.ConvReq) *HandlerReq {
	// step1. 构造SDK请求参数
	conversionReq := &ConvParams{
		OAID:           req.HuaweiParams.OAID,
		ConversionType: req.HuaweiParams.ConversionType,
		ContentID:      req.HuaweiParams.ContentID,
		Callback:       req.HuaweiParams.Callback,
		CampaignID:     req.HuaweiParams.CampaignID,
		Timestamp:      strconv.FormatInt(time.Now().UnixMilli(), 10),
		ConversionTime: strconv.FormatInt(time.Now().Unix(), 10),
	}

	return &HandlerReq{Req: conversionReq, SecretKey: req.HuaweiParams.ConversionSecretKey}
}

// MakeRes 解析响应
func (h *Handler) MakeRes(respCode int, resBody []byte, req *types.ConvReq, HandlerReq *HandlerReq) (*types.ConvRes, error) {
	result := &ConvResp{}
	// step1. 解析响应body
	if err := json.Unmarshal(resBody, result); err != nil {
		return nil, fmt.Errorf("MakeRes json err: %w", err)
	}
	// step2. 构造通用请求体
	res := &types.ConvRes{
		IsSuccess: result.ResultCode == 0,
		Channel:   req.BaseParams.Channel,
		Request: &types.ChannelRequestData{
			ReqType: types.RequestTypeHttp,
			ReqData: structs.Map(HandlerReq),
		},
		Response: &types.ChannelResponseData{
			StatusCode: respCode,
			ResData:    string(resBody),
		},
	}
	// NOTE: 错误消息
	if !res.IsSuccess {
		res.Err = fmt.Errorf(result.ResultMessage)
	}

	return res, nil
}
