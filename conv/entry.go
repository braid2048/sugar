package conv

import (
	"fmt"
	"github.com/halo2024/sugar/conv/handler"
	"github.com/halo2024/sugar/conv/handler/baidu"
	"github.com/halo2024/sugar/conv/handler/huawei"
	"github.com/halo2024/sugar/conv/handler/kuaishou"
	"github.com/halo2024/sugar/conv/handler/ocan"
	"github.com/halo2024/sugar/conv/handler/oppo"
	"github.com/halo2024/sugar/conv/handler/oppoHap"
	"github.com/halo2024/sugar/conv/handler/qutoutiao"
	"github.com/halo2024/sugar/conv/handler/tencent"
	"github.com/halo2024/sugar/conv/handler/uc"
	"github.com/halo2024/sugar/conv/handler/vivo"
	"github.com/halo2024/sugar/conv/handler/weibo"
	"github.com/halo2024/sugar/conv/types"
)

var ChannelHandlers = map[string]handler.IChannelHandler{
	types.ChannelOceanH5:      ocan.New(),
	types.ChannelKWaiH5:       kuaishou.New(),
	types.ChannelHuawei:       huawei.New(),
	types.ChannelOppoH5:       oppo.New(),
	types.ChannelOppoH5InSite: oppo.New(),
	types.ChannelOppoHap:      oppoHap.New(),
	types.ChannelVivoH5:       vivo.New(),
	types.ChannelQTT:          qutoutiao.New(),
	types.ChannelTencent:      tencent.New(),
	types.ChannelBaidu:        baidu.New(),
	types.ChannelUc:           uc.New(),
	types.ChannelWeiBo:        weibo.New(),
}

func NewChannelHandler(channel string) (handler.IChannelHandler, error) {
	if _, ok := ChannelHandlers[channel]; !ok {
		return nil, fmt.Errorf("channel not supported")
	}

	return ChannelHandlers[channel], nil
}
