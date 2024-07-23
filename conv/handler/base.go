package handler

import (
	"context"
	"github.com/braid2048/sugar/conv/types"
)

type IChannelHandler interface {
	// DoConv 执行回传
	DoConv(ctx context.Context, req *types.ConvReq) (*types.ConvRes, error)
}
