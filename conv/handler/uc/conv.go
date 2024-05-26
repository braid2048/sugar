package uc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/halo2024/sugar/conv/types"
	"github.com/halo2024/sugar/conv/utils"
	"net/http"
)

type Handler struct{}

type HandlerReq struct {
	Req string `json:"req" structs:"req"`
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) DoConv(ctx context.Context, req *types.ConvReq) (*types.ConvRes, error) {
	// step1. 检验参数
	if err := h.Validate(req); err != nil {
		return nil, fmt.Errorf("uc-conv -- Validate err: %w", err)
	}
	// step2. 构造参数
	convReq, err := h.MakeReq(req)
	if err != nil {
		return nil, fmt.Errorf("uc-conv -- MakeReq err: %w", err)
	}

	// step3. 回传
	respCode, respBody, err := utils.SendGetRequest(ctx, convReq.Req, nil)
	if err != nil {
		return nil, fmt.Errorf("uc-conv -- SendGetRequest err: %w", err)
	}

	// step4. 整理
	return h.MakeRes(respCode, respBody, req, convReq)
}

// Validate 检验参数
func (h *Handler) Validate(req *types.ConvReq) error {
	if err := req.Validate(); err != nil {
		return err
	}
	// NOTE: 必填字段
	if req.UCParams == nil {
		return fmt.Errorf("conv params is nil")
	}
	// NOTE: ConvURL
	if req.UCParams.ConvURL == "" {
		return fmt.Errorf("ConvURL is nil")
	}
	// NOTE: imei_MD5 && oaid
	if req.UCParams.ImeiSum == "" && req.UCParams.OAID == "" {
		return fmt.Errorf("ImeiSum and OAID is nil")
	}
	// NOTE: event
	if req.UCParams.Event == "" {
		return fmt.Errorf("event is nil")
	}

	return nil
}

// MakeReq 构造请求参数
func (h *Handler) MakeReq(req *types.ConvReq) (*HandlerReq, error) {
	// step1. 构造参数
	cParam := &ConvParams{
		Callback: req.UCParams.ConvURL,
		ImeiSum:  req.UCParams.ImeiSum,
		OAID:     req.UCParams.OAID,
	}
	// step2. 获取请求地址
	cURL, err := h.GetURL(cParam, req.UCParams.Event)
	if err != nil {
		return nil, fmt.Errorf("get callback url err , %w", err)
	}

	return &HandlerReq{Req: cURL}, nil
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
	res.IsSuccess = result.Status == 0

	return res, nil
}
