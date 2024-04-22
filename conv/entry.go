package conv

import (
	"fmt"
	"github.com/halo2024/sugar/conv/handler"
	"github.com/halo2024/sugar/conv/handler/huawei"
	"github.com/halo2024/sugar/conv/handler/kuaishou"
	"github.com/halo2024/sugar/conv/handler/ocan"
	"github.com/halo2024/sugar/conv/types"
)

var ChannelHandlers = map[string]handler.IChannelHandler{
	types.ChannelOceanH5: ocan.New(),
	types.ChannelKWaiH5:  kuaishou.New(),
	types.ChannelHuawei:  huawei.New(),
}

func NewChannelHandler(channel string) (handler.IChannelHandler, error) {
	if _, ok := ChannelHandlers[channel]; !ok {
		return nil, fmt.Errorf("channel not supported")
	}

	return ChannelHandlers[channel], nil
}
