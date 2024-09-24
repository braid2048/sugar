package honor

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/braid2048/sugar/conv/types"
	"github.com/braid2048/sugar/conv/utils"
	"github.com/fatih/structs"
)

type Handler struct{}

type HandlerReq struct {
	Req string `json:"req" structs:"req"`
}

const (
	honorConvURL = "https://ads-drcn.platform.hihonorcloud.com/api/ad-tracking/v1/conversion"
)

func New() *Handler {
	return &Handler{}
}

func (h *Handler) DoConv(ctx context.Context, req *types.ConvReq) (*types.ConvRes, error) {
	// step1. 检验参数
	if err := h.Validate(req); err != nil {
		return nil, fmt.Errorf("honor-conv -- Validate err: %w", err)
	}
	// step2. 构造参数
	convReq, err := h.MakeReq(req)
	if err != nil {
		return nil, fmt.Errorf("honor-conv -- MakeReq err: %w", err)
	}

	// step3. 回传
	respCode, respBody, err := utils.SendGetRequest(ctx, convReq.Req, nil)
	if err != nil {
		return nil, fmt.Errorf("honor-conv -- SendGetRequest err: %w", err)
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
	if req.HonorParams == nil {
		return fmt.Errorf("conv params is nil")
	}
	// NOTE: clickID
	if req.HonorParams.TrackID == "" {
		return fmt.Errorf("track id is nil")
	}
	// NOTE: event_type
	if req.HonorParams.ConversionID == "" {
		return fmt.Errorf("ConversionID is nil")
	}

	if req.HonorParams.AdvertiserID == "" {
		return fmt.Errorf("advertiser id is nil")
	}

	return nil
}

// MakeReq 构造请求参数
func (h *Handler) MakeReq(req *types.ConvReq) (*HandlerReq, error) {
	// 如果回传时间戳没传，这个给
	if req.HonorParams.ConversionTime == 0 {
		req.HonorParams.ConversionTime = time.Now().UnixMilli()
	}

	mapParams := structs.Map(req.HonorParams)

	// 额外字段是json结构，需要转成字符串
	if req.HonorParams.Extra != nil {
		extra, err := json.Marshal(req.HonorParams.Extra)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal honor extra: %w", err)
		}

		mapParams["extra"] = string(extra)
	}

	u, err := url.Parse(honorConvURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	query := u.Query()
	for k, v := range mapParams {
		query.Add(k, fmt.Sprintf("%v", v))
	}

	u.RawQuery = query.Encode()

	return &HandlerReq{Req: u.String()}, nil
}
