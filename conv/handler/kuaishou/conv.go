package kuaishou

import (
	"context"
	"fmt"
	"github.com/halo2024/sugar/conv/types"
	"github.com/halo2024/sugar/conv/utils"

	"net/http"
	"net/url"
	"strings"
	"time"
)

type Handler struct{}

type HandlerReq struct {
	Req string `json:"req" structs:"req"`
}

const (
	kWaiActivate    = 1
	kWaiActivateURL = "http://ad.partner.gifshow.com/track/activate?event_type=%v&event_time=%v&callback=%s"
)

func New() *Handler {
	return &Handler{}
}

func (h *Handler) DoConv(ctx context.Context, req *types.ConvReq) (*types.ConvRes, error) {
	// step1. 检验参数
	if err := h.Validate(req); err != nil {
		return nil, fmt.Errorf("kuaishou-conv -- Validate err: %w", err)
	}
	// step2. 构造参数
	convReq, err := h.MakeReq(req)
	if err != nil {
		return nil, fmt.Errorf("kuaishou-conv -- MakeReq err: %w", err)
	}

	// step3. 回传
	respCode, respBody, err := utils.SendGetRequest(ctx, convReq.Req, nil)
	if err != nil {
		return nil, fmt.Errorf("kuaishou-conv -- SendGetRequest err: %w", err)
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
	if req.KuaishouParams == nil {
		return fmt.Errorf("conv params is nil")
	}
	// NOTE: clickID
	if req.KuaishouParams.CallBack == "" {
		return fmt.Errorf("callback is nil")
	}
	// NOTE: event_type
	if req.KuaishouParams.EventType == "" {
		return fmt.Errorf("eventType is nil")
	}

	return nil
}

// MakeReq 构造请求参数
func (h *Handler) MakeReq(req *types.ConvReq) (*HandlerReq, error) {
	// step1. 加工click_id
	activateURL := strings.Replace(req.KuaishouParams.CallBack, "http=//", "http://", 1)

	now := time.Now()
	// NOTE: 少数情况快手宏会像巨量的 CALLBACK_PARAM 一样，没有拼接在 URL 后面
	if !strings.Contains(activateURL, "http://") {
		activateURL = fmt.Sprintf(kWaiActivateURL, req.KuaishouParams.EventType, now.UnixMilli(), activateURL)

		return &HandlerReq{Req: activateURL}, nil
	}
	// NOTE: 解码 URL
	decodedURL, err := url.QueryUnescape(activateURL)
	if err != nil {
		return nil, fmt.Errorf("failed to decode URL, err: %w", err)
	}
	// NOTE: 将 URL 解析为查询参数
	queryParams, err := url.ParseQuery(decodedURL[strings.IndexByte(decodedURL, '?')+1:])
	if err != nil {
		return nil, fmt.Errorf("failed to parse query parameters in URL, err: %w", err)
	}
	// NOTE: 获取 callback 值
	callback := queryParams.Get("callback")
	if callback == "" {
		return nil, fmt.Errorf("callback value not found in URL")
	}

	activateURL = fmt.Sprintf(kWaiActivateURL, req.KuaishouParams.EventType, now.UnixMilli(), callback)

	return &HandlerReq{Req: activateURL}, nil
}
