package wifi

import (
	"context"
	"fmt"
	"github.com/braid2048/sugar/conv/types"
	"github.com/braid2048/sugar/conv/utils"
	"github.com/jinzhu/copier"
	"net/http"
)

const (
	convURL = "http://c2.wkanx.com/tracking"
)

type Handler struct{}

type HandlerReq struct {
	Req  string `json:"req" structs:"req"`
	Sign string `json:"sign" structs:"sign"`
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) DoConv(ctx context.Context, req *types.ConvReq) (*types.ConvRes, error) {
	// step1. 检验参数
	if err := h.Validate(req); err != nil {
		return nil, fmt.Errorf("wifi-conv -- Validate err: %w", err)
	}
	// step2. 构造参数
	convReq, err := h.MakeReq(req)
	if err != nil {
		return nil, fmt.Errorf("wifi-conv -- MakeReq err: %w", err)
	}
	// step3. 回传
	respCode, respBody, err := utils.SendGetRequest(ctx, convReq.Req, nil)
	if err != nil {
		return nil, fmt.Errorf("wifi-conv -- SendGetRequest err: %w", err)
	}

	// step4. 整理
	res := &types.ConvRes{
		IsSuccess: respCode != http.StatusOK,
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
	if req.WifiParams == nil {
		return fmt.Errorf("conv params is nil")
	}
	// NOTE: sid
	if req.WifiParams.Sid == "" {
		return fmt.Errorf("sid is nil")
	}
	// NOTE: secret_key
	if req.WifiParams.SecretKey == "" {
		return fmt.Errorf("secret_key is nil")
	}

	return nil
}

// MakeReq 构造请求参数
func (h *Handler) MakeReq(req *types.ConvReq) (*HandlerReq, error) {
	// step1. 构造SDK请求参数
	convReq := &ConvParams{}

	if err := copier.Copy(&convReq, req.WifiParams); err != nil {
		return nil, fmt.Errorf("copy params err: %w", err)
	}

	// step2. 签名
	sign, URL := convReq.GetSignAndURL(req.WifiParams.SecretKey)

	return &HandlerReq{Req: URL, Sign: sign}, nil
}
