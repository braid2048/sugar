package baidu

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/braid2048/sugar/conv/types"
	"github.com/braid2048/sugar/conv/utils"
	"github.com/fatih/structs"
	"net/http"
	"strings"
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
		return nil, fmt.Errorf("baidu-conv -- Validate err: %w", err)
	}
	// step2. 构造参数
	convReq, err := h.MakeReq(req)
	if err != nil {
		return nil, fmt.Errorf("baidu-conv -- MakeReq err: %w", err)
	}
	// step3. 回传
	respCode, respBody, err := utils.SendGetRequest(ctx, convReq.Req, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, fmt.Errorf("baidu-conv -- send request err: %w", err)
	}

	return h.MakeRes(respCode, respBody, req, convReq)
}

// Validate 检验参数
func (h *Handler) Validate(req *types.ConvReq) error {
	if err := req.Validate(); err != nil {
		return err
	}
	// NOTE: 必填字段
	if req.BaiDuParams == nil {
		return fmt.Errorf("conv params is nil")
	}
	// NOTE: callback
	if req.BaiDuParams.CallBack == "" {
		return fmt.Errorf("callback is nil")
	}
	// NOTE: aType
	if req.BaiDuParams.AType == "" {
		return fmt.Errorf("a_type is nil")
	}
	// NOTE: aKey
	if req.BaiDuParams.Akey == "" {
		return fmt.Errorf("aKey is nil")
	}
	// NOTE: join_type
	if req.BaiDuParams.JoinType == "" {
		return fmt.Errorf("join_type is nil")
	}

	return nil
}

// MakeReq 构造请求参数
func (h *Handler) MakeReq(req *types.ConvReq) (*HandlerReq, error) {
	// step1. 构造callback
	convURL := strings.Replace(strings.Replace(req.BaiDuParams.CallBack, "{{ATYPE}}", req.BaiDuParams.AType, 1), "{{AVALUE}}", req.BaiDuParams.AValue, 1)
	// step2. 获取签名并拼接
	res := &HandlerReq{Req: fmt.Sprintf("%s&sign=%s&join_type=%v&oaid=%v&android_id=%v&bd_vid=%v", convURL, h.GetSign(convURL, req.BaiDuParams.Akey), req.BaiDuParams.JoinType, req.BaiDuParams.OaID, req.BaiDuParams.AndroidID, req.BaiDuParams.BdVID)}

	return res, nil
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
	res.IsSuccess = result.ErrorCode == 0
	// NOTE: 错误消息
	if !res.IsSuccess {
		res.Err = fmt.Errorf(result.ErrorMsg)
	}

	return res, nil
}
