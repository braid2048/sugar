package weibo

import (
	"context"
	"fmt"
	"github.com/halo2024/sugar/conv/types"
	"github.com/halo2024/sugar/conv/utils"
	"net/http"
)

type Handler struct{}

type HandlerReq struct {
	Req string `json:"req" structs:"req"`
}

const (
	weiBoURL = "https://appmonitor.biz.weibo.com/sdkserver/active?IMP=%v"
)

func New() *Handler {
	return &Handler{}
}

func (h *Handler) DoConv(ctx context.Context, req *types.ConvReq) (*types.ConvRes, error) {
	// step1. 检验参数
	if err := h.Validate(req); err != nil {
		return nil, fmt.Errorf("weibo-conv -- Validate err: %w", err)
	}
	// step2. 构造参数
	convReq, err := h.MakeReq(req)
	if err != nil {
		return nil, fmt.Errorf("weibo-conv -- MakeReq err: %w", err)
	}

	// step3. 回传
	respCode, respBody, err := utils.SendGetRequest(ctx, convReq.Req, nil)
	if err != nil {
		return nil, fmt.Errorf("weibo-conv -- SendGetRequest err: %w", err)
	}

	// step4. 整理
	res := &types.ConvRes{
		IsSuccess: respCode == http.StatusOK,
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
	if req.WeiBoParams == nil {
		return fmt.Errorf("conv params is nil")
	}
	// NOTE: imp
	if req.WeiBoParams.IMP == "" {
		return fmt.Errorf("callback is nil")
	}

	return nil
}

// MakeReq 构造请求参数
func (h *Handler) MakeReq(req *types.ConvReq) (*HandlerReq, error) {
	convURL := fmt.Sprintf(weiBoURL, req.WeiBoParams.IMP)
	// 转化事件类型
	if req.WeiBoParams.ActionType != "" {
		convURL = fmt.Sprintf(convURL+"&action_type=%v", req.WeiBoParams.ActionType)
	}
	// 付费金额
	if req.WeiBoParams.Price > 0 {
		convURL = fmt.Sprintf(convURL+"&price=%v", req.WeiBoParams.Price)
	}
	// 行为时间
	if req.WeiBoParams.ActiveTime > 0 {
		convURL = fmt.Sprintf(convURL+"&active_time=%v", req.WeiBoParams.ActiveTime)
	}

	return &HandlerReq{Req: convURL}, nil
}
