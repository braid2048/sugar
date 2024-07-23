package qutoutiao

import (
	"context"
	"fmt"
	"github.com/braid2048/sugar/conv/types"
	"github.com/braid2048/sugar/conv/utils"
	"net/http"
	"time"
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
		return nil, fmt.Errorf("qtt-conv -- Validate err: %w", err)
	}
	// step2. 构造参数
	convReq, err := h.MakeReq(req)
	if err != nil {
		return nil, fmt.Errorf("qtt-conv -- MakeReq err: %w", err)
	}

	// step3. 回传
	respCode, respBody, err := utils.SendGetRequest(ctx, convReq.Req, nil)
	if err != nil {
		return nil, fmt.Errorf("qtt-conv -- SendGetRequest err: %w", err)
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
	if req.QTTParams == nil {
		return fmt.Errorf("conv params is nil")
	}
	// NOTE: clickID
	if req.QTTParams.CallBack == "" {
		return fmt.Errorf("callback is nil")
	}
	// NOTE: op2
	if req.QTTParams.OP2 == "" {
		return fmt.Errorf("op2 is nil")
	}

	return nil
}

// MakeReq 构造请求参数
func (h *Handler) MakeReq(req *types.ConvReq) (*HandlerReq, error) {
	reqURL := fmt.Sprintf("%v&op2=%v&opt_active_time=%v", req.QTTParams.CallBack, req.QTTParams.OP2, time.Now().Unix())

	if req.QTTParams.Arpu > 0 {
		reqURL = fmt.Sprintf("%v&arpu=%v", reqURL, req.QTTParams.Arpu)
	}

	return &HandlerReq{Req: reqURL}, nil
}
