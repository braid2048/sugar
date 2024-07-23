package vivo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/braid2048/sugar/conv/types"
	"github.com/braid2048/sugar/conv/utils"
	"github.com/fatih/structs"
	"golang.org/x/exp/slices"
	"net/http"
	"strings"
	"time"
)

const (
	convURL     = "https://marketing-api.vivo.com.cn/openapi/v1/advertiser/behavior/upload"
	nonceLength = 32
)

type Handler struct{}

type HandlerReq struct {
	Req         *ConvParams `json:"req" structs:"req"`
	URLWithSign string      `json:"url_with_sign" structs:"url_with_sign"`
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) DoConv(ctx context.Context, req *types.ConvReq) (*types.ConvRes, error) {
	// step1. 检验参数
	if err := h.Validate(req); err != nil {
		return nil, fmt.Errorf("vivo-conv -- Validate err: %w", err)
	}
	// step2. 构造参数
	convReq, err := h.MakeReq(req)
	if err != nil {
		return nil, fmt.Errorf("vivo-conv -- MakeReq err: %w", err)
	}
	// step3. 解析body
	body, err := json.Marshal(convReq.Req)
	if err != nil {
		return nil, fmt.Errorf("vivo-conv -- body json err: %w", err)
	}
	// step4. 回传
	respCode, respBody, err := utils.SendPostRequest(ctx, convReq.URLWithSign, map[string]string{"Content-Type": "application/json"}, body)
	if err != nil {
		return nil, fmt.Errorf("vivo-conv -- send request err: %w", err)
	}

	return h.MakeRes(respCode, respBody, req, convReq)
}

// Validate 检验参数
func (h *Handler) Validate(req *types.ConvReq) error {
	if err := req.Validate(); err != nil {
		return err
	}
	// NOTE: 必填字段
	if req.VivoParams == nil {
		return fmt.Errorf("conv params is nil")
	}
	// NOTE: srcType
	if req.VivoParams.SrcType == "" {
		return fmt.Errorf("srcType is nil")
	}
	// NOTE: srcId
	if req.VivoParams.SrcID == "" {
		return fmt.Errorf("srcId is nil")
	}
	// NOTE: cvType
	if req.VivoParams.CvType == "" {
		return fmt.Errorf("cvType is nil")
	}
	// NOTE: requestId
	if req.VivoParams.RequestID == "" {
		return fmt.Errorf("requestId is nil")
	}
	// NOTE: creativeId
	if req.VivoParams.CreativeID == "" {
		return fmt.Errorf("creativeId is nil")
	}
	// NOTE: srcType 为 APP/Quickapp 时，pkgName 必传
	if slices.Contains([]string{"app", "quickapp"}, strings.ToLower(req.VivoParams.SrcType)) && req.VivoParams.PackageName == "" {
		return fmt.Errorf("pkgName is nil")
	}
	// NOTE: srcType 为 Web 时 , pageUrl 必传
	if strings.ToLower(req.VivoParams.SrcType) == "web" && req.VivoParams.PageURL == "" {
		return fmt.Errorf("pageUrl is nil")
	}
	// NOTE: srcType 为 App 时 ,userIdType 和 userId 必传
	if strings.ToLower(req.VivoParams.SrcType) == "app" && (req.VivoParams.UserIDType == "" || req.VivoParams.UserID == "") {
		return fmt.Errorf("userIdType or userId is nil")
	}
	// NOTE: accessToken
	if req.VivoParams.AccessToken == "" {
		return fmt.Errorf("accessToken is nil")
	}
	// NOTE: advertiserId
	if req.VivoParams.AdvertiserID == "" {
		return fmt.Errorf("advertiserId is nil")
	}

	return nil
}

// MakeReq 构造请求参数
func (h *Handler) MakeReq(req *types.ConvReq) (*HandlerReq, error) {
	// step1. 构造请求参数
	convDateList := &DataItem{
		CreativeID: req.VivoParams.CreativeID,
		RequestID:  req.VivoParams.RequestID,
		CvType:     req.VivoParams.CvType,
		CvTime:     time.Now().UnixMilli(),
	}

	convReq := &ConvParams{
		DataList:    []*DataItem{convDateList},
		SrcID:       req.VivoParams.SrcID,
		SrcType:     req.VivoParams.SrcType,
		PackageName: req.VivoParams.PackageName,
		PageURL:     req.VivoParams.PageURL,
	}
	// step2. 构造加密url
	urlWithSign := fmt.Sprintf("%v?access_token=%v&timestamp=%v&nonce=%v&advertiser_id=%v", convURL, req.VivoParams.AccessToken, time.Now().UnixMilli(), h.RandString(nonceLength), req.VivoParams.AdvertiserID)

	return &HandlerReq{Req: convReq, URLWithSign: urlWithSign}, nil
}

// MakeRes 解析响应
func (h *Handler) MakeRes(respCode int, resBody []byte, req *types.ConvReq, HandlerReq *HandlerReq) (*types.ConvRes, error) {
	// step1. 初始化响应体
	res := &types.ConvRes{
		Channel: req.BaseParams.Channel,
		Request: &types.ChannelRequestData{
			ReqType: types.RequestTypeHttp,
			ReqData: structs.Map(HandlerReq),
		},
		Response: &types.ChannelResponseData{
			StatusCode: respCode,
			ResData:    string(resBody),
		},
	}
	// step2. 判断响应码
	if respCode >= http.StatusBadRequest {
		res.IsSuccess = false

		return res, nil
	}
	// step3. 解析响应
	result := &ConvResp{}

	if err := json.Unmarshal(resBody, result); err != nil {
		return nil, fmt.Errorf("MakeRes json err: %w", err)
	}
	// step4. 判断成功
	res.IsSuccess = result.Code == 0
	// NOTE: 错误消息
	if !res.IsSuccess {
		res.Err = fmt.Errorf(result.Message)
	}

	return res, nil
}
