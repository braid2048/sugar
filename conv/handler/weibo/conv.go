package weibo

import (
	"context"
	"fmt"
	"github.com/braid2048/sugar/conv/types"
	"github.com/braid2048/sugar/conv/utils"
	"net/http"
)

type Handler struct{}

type HandlerReq struct {
	Req string `json:"req" structs:"req"`
}

const (
	weiBoQuickURL = "https://appmonitor.biz.weibo.com/sdkserver/active?IMP=%v"        // 快应用回传
	weiBoLandURL  = "https://api.biz.weibo.com/v4/track/activate?time=%v&behavior=%v" // 落地页回传
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
	// NOTE: 回传方式
	switch req.WeiBoParams.ConvType {
	case 1:
		// 快应用回传
		if req.WeiBoParams.QuickParams == nil {
			return fmt.Errorf("quick_params is nil")
		}
		// imp
		if req.WeiBoParams.QuickParams.IMP == "" {
			return fmt.Errorf("imp is nil")
		}
	case 2:
		// 落地页回传
		if req.WeiBoParams.LandParams == nil {
			return fmt.Errorf("land params is nil")
		}
		// time
		if req.WeiBoParams.LandParams.Time <= 0 {
			return fmt.Errorf("time is nil")
		}
		// behavior
		if req.WeiBoParams.LandParams.Behavior == "" {
			return fmt.Errorf("behavior is nil")
		}
	default:
		return fmt.Errorf("conv_type is invalid")
	}

	return nil
}

// MakeReq 构造请求参数
func (h *Handler) MakeReq(req *types.ConvReq) (*HandlerReq, error) {
	var convURL string
	// 回传方式判断
	switch req.WeiBoParams.ConvType {
	case 1: // 快应用回传
		convURL = fmt.Sprintf(weiBoQuickURL, req.WeiBoParams.QuickParams.IMP)
		// 转化事件类型
		if req.WeiBoParams.QuickParams.ActionType != "" {
			convURL = fmt.Sprintf(convURL+"&action_type=%v", req.WeiBoParams.QuickParams.ActionType)
		}
		// 付费金额
		if req.WeiBoParams.QuickParams.Price > 0 {
			convURL = fmt.Sprintf(convURL+"&price=%v", req.WeiBoParams.QuickParams.Price)
		}
		// 行为时间
		if req.WeiBoParams.QuickParams.ActiveTime > 0 {
			convURL = fmt.Sprintf(convURL+"&active_time=%v", req.WeiBoParams.QuickParams.ActiveTime)
		}
	case 2: // 落地页回传
		convURL = fmt.Sprintf(weiBoLandURL, req.WeiBoParams.LandParams.Time, req.WeiBoParams.LandParams.Behavior)
		// mark_id
		if req.WeiBoParams.LandParams.MarkID != "" {
			convURL = fmt.Sprintf(convURL+"&mark_id=%v", req.WeiBoParams.LandParams.MarkID)
		}
	default:
		return nil, fmt.Errorf("conv_type is invalid")
	}

	return &HandlerReq{Req: convURL}, nil
}
