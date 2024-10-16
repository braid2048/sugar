package iqiyi

import (
	"context"
	"fmt"
	"github.com/braid2048/sugar/conv/types"
	"github.com/braid2048/sugar/conv/utils"
	"net/http"
	"net/url"
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
		return nil, fmt.Errorf("iqiyi-conv -- Validate err: %w", err)
	}
	// step2. 构造参数
	convReq, err := h.MakeReq(req)
	if err != nil {
		return nil, fmt.Errorf("iqiyi-conv -- MakeReq err: %w", err)
	}

	// step3. 回传
	respCode, respBody, err := utils.SendGetRequest(ctx, convReq.Req, nil)
	if err != nil {
		return nil, fmt.Errorf("iqiyi-conv -- SendGetRequest err: %w", err)
	}

	// step4. 整理
	res := &types.ConvRes{
		IsSuccess: respCode < http.StatusBadRequest,
		Channel:   req.BaseParams.Channel,
		Request: &types.ChannelRequestData{
			ReqType: types.RequestTypeHttp,
			ReqData: convReq.Req,
		},
		Response: &types.ChannelResponseData{
			StatusCode: respCode,
			ResData:    string(respBody),
		},
		Err: nil,
	}

	return res, nil
}

// Validate 检验参数
func (h *Handler) Validate(req *types.ConvReq) error {
	if err := req.Validate(); err != nil {
		return err
	}
	// NOTE: 必填字段
	if req.IQiYiParams == nil {
		return fmt.Errorf("conv params is nil")
	}
	// NOTE: callback
	if req.IQiYiParams.Callback == "" {
		return fmt.Errorf("callback is nil")
	}
	// NOTE: event_type
	if req.IQiYiParams.EventType == "" {
		return fmt.Errorf("event_type is nil")
	}

	return nil
}

// MakeReq 构造请求参数
func (h *Handler) MakeReq(req *types.ConvReq) (*HandlerReq, error) {
	// step1. url解码
	callbackURL, err := url.QueryUnescape(req.IQiYiParams.Callback)
	if err != nil {
		return nil, fmt.Errorf("failed to unescape callback, err: %w", err)
	}
	// step2. 拼接事件
	callback := fmt.Sprintf("%s&event_type=%v", callbackURL, req.IQiYiParams.EventType)

	return &HandlerReq{Req: callback}, nil
}
