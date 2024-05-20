package tencent

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/halo2024/sugar/conv/types"
	"github.com/halo2024/sugar/conv/utils"
	"net/http"
	"time"
)

type Handler struct{}

type HandlerReq struct {
	Req *ConvParams `json:"req" structs:"req"`
	URL string      `json:"url" structs:"url"`
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) DoConv(ctx context.Context, req *types.ConvReq) (*types.ConvRes, error) {
	// step1. 检验参数
	if err := h.Validate(req); err != nil {
		return nil, fmt.Errorf("tencent-conv -- Validate err: %w", err)
	}
	// step2. 构造参数
	convReq, err := h.MakeReq(req)
	if err != nil {
		return nil, fmt.Errorf("tencent-conv -- MakeReq err: %w", err)
	}
	// step3. 解析body
	body, err := json.Marshal(convReq.Req)
	if err != nil {
		return nil, fmt.Errorf("tencent-conv -- body json err: %w", err)
	}
	// step4. 回传
	respCode, respBody, err := utils.SendPostRequest(ctx, convReq.URL, map[string]string{"Content-Type": "application/json"}, body)
	if err != nil {
		return nil, fmt.Errorf("tencent-conv -- send request err: %w", err)
	}
	// step4. 整理
	res := &types.ConvRes{
		IsSuccess: respCode < http.StatusBadRequest,
		Channel:   req.BaseParams.Channel,
		Request: &types.ChannelRequestData{
			ReqType: types.RequestTypeHttp,
			ReqData: structs.Map(convReq),
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
	if req.TencentParams == nil {
		return fmt.Errorf("conv params is nil")
	}
	// NOTE: action_type
	if req.TencentParams.ActionType == "" {
		return fmt.Errorf("action_type is nil")
	}
	// NOTE: callback
	if req.TencentParams.CallBack == "" {
		return fmt.Errorf("callback is nil")
	}
	// NOTE: hash_imei & hash_oaid & hash_android_id
	if req.TencentParams.HashOAID == "" && req.TencentParams.HashAndroidID == "" && req.TencentParams.HashIMEI == "" {
		return fmt.Errorf("hash_imei or hash_oaid or hash_android_id is nil")
	}

	return nil
}

// MakeReq 构造请求参数
func (h *Handler) MakeReq(req *types.ConvReq) (*HandlerReq, error) {
	// step1. 构造请求参数
	convUserID := &UserID{
		HashIMEI:      req.TencentParams.HashIMEI,
		HashOAID:      req.TencentParams.HashOAID,
		HashAndroidID: req.TencentParams.HashAndroidID,
	}

	convReq := &ConvParams{
		Actions: []*Action{
			{
				OuterActionID: req.TencentParams.OuterActionID,
				ActionTime:    time.Now().Unix(),
				UserID:        convUserID,
				ActionType:    req.TencentParams.ActionType,
			},
		},
	}
	// step2. 构造加密url
	URL := req.TencentParams.CallBack

	return &HandlerReq{Req: convReq, URL: URL}, nil
}
