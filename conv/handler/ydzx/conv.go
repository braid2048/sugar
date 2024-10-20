package ydzx

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/braid2048/sugar/conv/types"
	"github.com/braid2048/sugar/conv/utils"
	"github.com/fatih/structs"
	"net/http"
)

type Handler struct{}

type HandlerReq struct {
	Req string `json:"req" structs:"req"`
}

func New() *Handler {
	return &Handler{}
}

type ConvResp struct {
	Time string `json:"time"`
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func (h *Handler) DoConv(ctx context.Context, req *types.ConvReq) (*types.ConvRes, error) {
	// step1. 检验参数
	if err := h.Validate(req); err != nil {
		return nil, fmt.Errorf("ydzx-conv -- Validate err: %w", err)
	}
	// step2. 构造参数
	convReq, err := h.MakeReq(req)
	if err != nil {
		return nil, fmt.Errorf("ydzx-conv -- MakeReq err: %w", err)
	}

	// step3. 回传
	respCode, respBody, err := utils.SendGetRequest(ctx, convReq.Req, nil)
	if err != nil {
		return nil, fmt.Errorf("ydzx-conv -- SendGetRequest err: %w", err)
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
	if req.YdzxParams == nil {
		return fmt.Errorf("conv params is nil")
	}
	// NOTE: conv_ext
	if req.YdzxParams.ConvExt == "" {
		return fmt.Errorf("conv_ext is nil")
	}
	// NOTE: conv_action
	if req.YdzxParams.ConvAction == "" {
		return fmt.Errorf("conv_action is nil")
	}

	return nil
}

// MakeReq 构造请求参数
func (h *Handler) MakeReq(req *types.ConvReq) (*HandlerReq, error) {
	callback := fmt.Sprintf("%s?conv_ext=%v&conv_action=%v", "http://conv.youdao.com/api/track", req.YdzxParams.ConvExt, req.YdzxParams.ConvAction)

	return &HandlerReq{Req: callback}, nil
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
	res.IsSuccess = result.Code == "success"
	// NOTE: 错误消息
	if !res.IsSuccess {
		res.Err = fmt.Errorf(result.Msg)
	}

	return res, nil
}
