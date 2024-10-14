package conv

import (
	"fmt"
	"github.com/braid2048/sugar/conv/handler"
	"github.com/braid2048/sugar/conv/handler/baidu"
	"github.com/braid2048/sugar/conv/handler/honor"
	"github.com/braid2048/sugar/conv/handler/huawei"
	"github.com/braid2048/sugar/conv/handler/kuaishou"
	"github.com/braid2048/sugar/conv/handler/magic"
	"github.com/braid2048/sugar/conv/handler/ocan"
	"github.com/braid2048/sugar/conv/handler/octopus"
	"github.com/braid2048/sugar/conv/handler/oppo"
	"github.com/braid2048/sugar/conv/handler/oppoHap"
	"github.com/braid2048/sugar/conv/handler/qutoutiao"
	"github.com/braid2048/sugar/conv/handler/tencent"
	"github.com/braid2048/sugar/conv/handler/uc"
	"github.com/braid2048/sugar/conv/handler/vivo"
	"github.com/braid2048/sugar/conv/handler/weibo"
	"github.com/braid2048/sugar/conv/handler/wifi"
	"github.com/braid2048/sugar/conv/handler/xmly"
	"github.com/braid2048/sugar/conv/types"
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
	types.ChannelHonor:        honor.New(),
	types.ChannelWifi:         wifi.New(),
	types.ChannelMagic:        magic.New(),
	types.ChannelOctopus:      octopus.New(),
	types.ChannelXmly:         xmly.New(),
}

func NewChannelHandler(channel string) (handler.IChannelHandler, error) {
	if _, ok := ChannelHandlers[channel]; !ok {
		return nil, fmt.Errorf("channel not supported")
	}

	return ChannelHandlers[channel], nil
}
