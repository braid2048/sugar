package oppoHap

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/halo2024/sugar/conv/types"
	"github.com/halo2024/sugar/conv/utils"
	"github.com/jinzhu/copier"
	"golang.org/x/exp/slices"
	"net/http"
	"strconv"
	"time"
)

const (
	convURL      = "https://api.ads.heytapmobi.com/api/uploadActiveData"
	slat         = "e0u6fnlag06lc3pl"         // 接入时由 OPPO 提供
	aesBase64Key = "XGAXicVG5GMBsx5bueOe4w==" // 接入时由 OPPO 提供
)

type Handler struct{}

type HandlerReq struct {
	Req  *ConvParams `json:"req" structs:"req"`
	Sign string      `json:"sign" structs:"sign"`
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) DoConv(ctx context.Context, req *types.ConvReq) (*types.ConvRes, error) {
	// step1. 检验参数
	if err := h.Validate(req); err != nil {
		return nil, fmt.Errorf("oppoH5-conv -- Validate err: %w", err)
	}
	// step2. 构造参数
	convReq, err := h.MakeReq(req)
	if err != nil {
		return nil, fmt.Errorf("oppoH5-conv -- MakeReq err: %w", err)
	}
	// step3. 解析body
	body, err := json.Marshal(convReq.Req)
	if err != nil {
		return nil, fmt.Errorf("oppoH5-conv -- body json err: %w", err)
	}
	// step4. 加密
	sign := h.GetSign(body, convReq.Req.Timestamp, slat)
	convReq.Sign = sign

	// step3. 回传
	respCode, respBody, err := utils.SendPostRequest(ctx, convURL, map[string]string{"Content-Type": "application/json", "signature": convReq.Sign, "timestamp": strconv.FormatInt(convReq.Req.Timestamp, 10)}, body)
	if err != nil {
		return nil, fmt.Errorf("huawei-conv -- send request err: %w", err)
	}

	return h.MakeRes(respCode, respBody, req, convReq)
}

// Validate 检验参数
func (h *Handler) Validate(req *types.ConvReq) error {
	if err := req.Validate(); err != nil {
		return err
	}
	// NOTE: 必填字段
	if req.OppoHapParams == nil {
		return fmt.Errorf("conv params is nil")
	}
	// NOTE: oaid
	if req.OppoHapParams.OUID == "" && req.OppoHapParams.Imei == "" {
		return fmt.Errorf("oaid and imei is nil")
	}
	// NOTE: pkg
	if req.OppoHapParams.Pkg == "" {
		return fmt.Errorf("pkg is nil")
	}
	// NOTE: dataType
	if req.OppoHapParams.DataType == 0 {
		return fmt.Errorf("dataType is nil")
	}
	// NOTE: channel
	if !slices.Contains([]int{0, 1, 2}, req.OppoHapParams.Channel) {
		return fmt.Errorf("channel is abnormal")
	}
	// NOTE: type
	if !slices.Contains([]int{0, 1, 2}, req.OppoHapParams.Type) {
		return fmt.Errorf("type is abnormal")
	}
	// NOTE: appType
	if !slices.Contains([]int{0, 1, 2, 3}, req.OppoHapParams.AppType) {
		return fmt.Errorf("appType is abnormal")
	}
	// NOTE: ascribeType
	if !slices.Contains([]int{0, 1, 2}, req.OppoHapParams.AscribeType) {
		return fmt.Errorf("ascribeType is abnormal")
	}
	// NOTE: adid
	if req.OppoHapParams.AdID == 0 {
		return fmt.Errorf("adid is nil")
	}

	return nil
}

// MakeReq 构造请求参数
func (h *Handler) MakeReq(req *types.ConvReq) (*HandlerReq, error) {
	// step1. 构造SDK请求参数
	convReq := &ConvParams{}

	if err := copier.Copy(&convReq, req.OppoHapParams); err != nil {
		return nil, fmt.Errorf("copy params err: %w", err)
	}
	// step2. imei加密
	if req.OppoHapParams.Imei != "" {
		enImei, err := h.EncryptByAes([]byte(req.OppoHapParams.Imei))
		if err != nil {
			return nil, fmt.Errorf("encrypt imei err: %w", err)
		}

		convReq.IMEI = enImei
	}
	// step3. oaid加密
	if req.OppoHapParams.OUID != "" {
		enOaid, err := h.EncryptByAes([]byte(req.OppoHapParams.OUID))
		if err != nil {
			return nil, fmt.Errorf("encrypt oaid err: %w", err)
		}

		convReq.OAID = enOaid
	}
	// step4. 毫秒级时间戳
	convReq.Timestamp = time.Now().UnixMilli()

	return &HandlerReq{Req: convReq}, nil
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
	res.IsSuccess = result.Ret == 0
	// NOTE: 错误消息
	if !res.IsSuccess {
		res.Err = fmt.Errorf(result.Msg)
	}

	return res, nil
}
