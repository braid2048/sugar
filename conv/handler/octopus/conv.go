package octopus

import (
	"context"
	"fmt"
	"github.com/braid2048/sugar/conv/types"
	"github.com/braid2048/sugar/conv/utils"
	"net/http"
	"net/url"
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
		return nil, fmt.Errorf("octopus-conv -- Validate err: %w", err)
	}
	// step2. 构造参数
	convReq, err := h.MakeReq(req)
	if err != nil {
		return nil, fmt.Errorf("octopus-conv -- MakeReq err: %w", err)
	}

	// step3. 回传
	respCode, respBody, err := utils.SendGetRequest(ctx, convReq.Req, nil)
	if err != nil {
		return nil, fmt.Errorf("octopus-conv -- SendGetRequest err: %w", err)
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
	if req.OctopusParams == nil {
		return fmt.Errorf("conv params is nil")
	}
	// NOTE: callback
	if req.OctopusParams.Callback == "" {
		return fmt.Errorf("callback is nil")
	}
	// NOTE: event_type
	if req.OctopusParams.EventType == "" {
		return fmt.Errorf("event_type is nil")
	}
	// NOTE: 毫秒级
	if req.OctopusParams.Timestamp == 0 {
		req.OctopusParams.Timestamp = time.Now().UnixMilli()
	}

	return nil
}

// MakeReq 构造请求参数
func (h *Handler) MakeReq(req *types.ConvReq) (*HandlerReq, error) {
	// step1. url解码
	callbackURL, err := url.QueryUnescape(req.OctopusParams.Callback)
	if err != nil {
		return nil, fmt.Errorf("failed to unescape callback, err: %w", err)
	}
	// step2. 拼接事件
	callback := fmt.Sprintf("%s&event_type=%v&timestamp=%v", callbackURL, req.OctopusParams.EventType, req.OctopusParams.Timestamp)

	return &HandlerReq{Req: callback}, nil
}
