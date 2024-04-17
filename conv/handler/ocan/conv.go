package ocan

import (
	"context"
	"fmt"
	oceanconversion "github.com/bububa/oceanengine/marketing-api/api/conversion"
	"github.com/fatih/structs"
	"github.com/halo2024/sugar/conv/types"
	"time"

	oceancore "github.com/bububa/oceanengine/marketing-api/core"
	oceanenum "github.com/bububa/oceanengine/marketing-api/enum"
	oceanreq "github.com/bububa/oceanengine/marketing-api/model/conversion"
)

type Handler struct{}

type HandlerReq struct {
	Req    *oceanreq.Request    `json:"req" structs:"req"`
	Client *oceancore.SDKClient `json:"client" structs:"client"`
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

func (h Handler) DoConv(ctx context.Context, req *types.ConvReq) (*types.ConvRes, error) {
	// step1. 检验参数
	if err := h.Validate(req); err != nil {
		return nil, fmt.Errorf("ocan-conv -- Validate err: %w", err)
	}
	// step2. 构造参数
	convReq := h.MakeReq(req)

	// step3. 回传
	respCode, convErr := oceanconversion.Conversion(convReq.Client, convReq.Req)

	// step4. 整理
	res := &types.ConvRes{
		IsSuccess: convErr == nil,
		Channel:   req.Channel,
		Request: &types.ChannelRequestData{
			ReqType: types.RequestTypeSDK,
			ReqData: structs.Map(convReq.Req),
		},
		Response: &types.ChannelResponseData{
			StatusCode: respCode,
			ResData:    nil,
		},
		Err: convErr,
	}

	return res, nil
}

// Validate 检验参数
func (h Handler) Validate(req *types.ConvReq) error {
	// NOTE: clickID
	if req.ClickID == "" {
		return fmt.Errorf("clickid is nil")
	}

	return nil
}

// MakeReq 构造请求参数
func (h Handler) MakeReq(req *types.ConvReq) *HandlerReq {
	// step1. 事件转化
	var event oceanenum.ConversionEventType

	switch req.ConvEvent {
	case oceanenum.Conversion_ACTIVE:
		event = oceanenum.Conversion_ACTIVE
	case oceanenum.Conversion_CUSTOMER_EFFECTIVE:
		event = oceanenum.Conversion_CUSTOMER_EFFECTIVE
	case oceanenum.Conversion_PAGE_VIEW:
		event = oceanenum.Conversion_PAGE_VIEW
	case oceanenum.Conversion_CLUE_CONFIRM:
		event = oceanenum.Conversion_CLUE_CONFIRM
	case "submit_certification": // 提交认证 包里没有预定义宏，这里直接根据文档字段硬编码
		event = "submit_certification"
	case oceanenum.Conversion_ACTIVE_REGISTER:
		event = oceanenum.Conversion_ACTIVE_REGISTER
	case oceanenum.Conversion_GAME_ADDICTION:
		event = oceanenum.Conversion_GAME_ADDICTION
	default:
		event = oceanenum.Conversion_FORM
	}
	// step2. 构造SDK请求参数
	conversionReq := &oceanreq.Request{
		EventType: event,
		Context:   &oceanreq.Context{Ad: &oceanreq.ContextAd{Callback: req.ClickID}},
		Timestamp: time.Now().UnixMilli(),
	}

	conversionClt := oceancore.NewSDKClient(0, "")

	return &HandlerReq{Req: conversionReq, Client: conversionClt}
}
