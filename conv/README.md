### 渠道回传包：

##### 目录结构：

- entry         ------   入口文件

- types        ------   常量与结构体

- handler     ------   各渠道执行回传的工厂

##### 使用方法：

1. 先构造通用请求体： types.ConvReq
2. 根据渠道创建实例：NewChannelHandler（channel）
3. 执行实例的 DoConv（&types.ConvReq）进行回传

##### 目前支持的渠道：

- 巨量
- 快手

##### 依赖包：

- github.com/bububa/oceanengine  ： 巨量sdk
- github.com/fatih/structs ：结构体快捷转化工具