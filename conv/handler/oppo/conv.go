package oppo

import (
	"context"
	"fmt"
	"github.com/braid2048/sugar/conv/types"
	oppoconversion "github.com/bububa/oppo-omni/api/clue"
	oppocore "github.com/bububa/oppo-omni/core"
	oppoenum "github.com/bububa/oppo-omni/enum"
	oppomodel "github.com/bububa/oppo-omni/model"
	opporeq "github.com/bububa/oppo-omni/model/clue"
	"github.com/fatih/structs"
)

type Handler struct{}

type HandlerReq struct {
	Req    *opporeq.SendDataRequest `json:"req" structs:"req"`
	Client *oppocore.SDKClient      `json:"client" structs:"client"`
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
		return nil, fmt.Errorf("oppoH5-conv -- Validate err: %w", err)
	}
	// step2. 构造参数
	convReq := h.MakeReq(req)

	// step3. 回传
	convErr := oppoconversion.SendData(convReq.Client, convReq.Req)

	// step4. 整理
	res := &types.ConvRes{
		IsSuccess: convErr == nil,
		Channel:   req.BaseParams.Channel,
		Request: &types.ChannelRequestData{
			ReqType: types.RequestTypeSDK,
			ReqData: structs.Map(convReq),
		},
		Response: &types.ChannelResponseData{
			StatusCode: 0,
			ResData:    nil,
		},
		Err: convErr,
	}

	return res, nil
}

// Validate 检验参数
func (h *Handler) Validate(req *types.ConvReq) error {
	if err := req.Validate(); err != nil {
		return err
	}
	// NOTE: 必填字段
	if req.OppoParams == nil {
		return fmt.Errorf("conv params is nil")
	}
	// NOTE: owner_id
	if req.OppoParams.OwnerID == 0 {
		return fmt.Errorf("ownerId is nil")
	}
	// NOTE: page_id
	if req.OppoParams.PageID == 0 {
		return fmt.Errorf("pageId is nil")
	}
	// NOTE: app_id
	if req.OppoParams.AppID == "" {
		return fmt.Errorf("appId is nil")
	}
	// NOTE: app_key
	if req.OppoParams.AppKey == "" {
		return fmt.Errorf("appKey is nil")
	}
	// NOTE: tid
	if req.OppoParams.TID == "" {
		return fmt.Errorf("tid is nil")
	}
	// NOTE: lbid
	if req.OppoParams.LbID == "" {
		return fmt.Errorf("lbid is nil")
	}
	// NOTE: transform_type
	if req.OppoParams.TransformType == 0 {
		return fmt.Errorf("transformType is nil")
	}

	return nil
}

// MakeReq 构造请求参数
func (h *Handler) MakeReq(req *types.ConvReq) *HandlerReq {
	// step1. 构造SDK请求参数
	conversionReq := &opporeq.SendDataRequest{
		BaseRequest:   oppomodel.BaseRequest{OwnerID: req.OppoParams.OwnerID},
		PageID:        req.OppoParams.PageID,
		Ip:            req.BaseParams.Ip,
		Tid:           req.OppoParams.TID,
		Lbid:          req.OppoParams.LbID,
		TransformType: oppoenum.ClueTransformType(req.OppoParams.TransformType),
	}

	conversionClt := oppocore.NewSDKClient(req.OppoParams.AppID, req.OppoParams.AppKey)

	return &HandlerReq{Req: conversionReq, Client: conversionClt}
}
