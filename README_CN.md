# xmhOpenApiSdk

xmhOpenApiSdk

```markdown
# xmhOpenApiSdk

## 功能模块

- 认证 (auth)：处理 API 认证相关功能
- 试算报价 (calc)：提供试算报价
- 商品同步 (item)：延保险必须要同步商品，邮包险可忽略
- 订单    (order)：电商平台订单同步，以及投保，取消投保等
- 发货    (ship)：发货数据同步
- 理赔    (claim)：理赔报案

## 安装

```go
go get github.com/cjay-shouhui/xmhOpenApiSdk
```

## 使用说明

### 1. 初始化 SDK

```go
import "github.com/cjay-shouhui/xmhOpenApiSdk"

xmhsdk.AppId = "{your appId}" //联系小棉花获取
xmhsdk.AppSecret = "{your appSecret}" //联系小棉花获取
xmhsdk.SetEnv(xmhsdk.EnvAlpha) // 使用开发环境

```

### 2. 配置日志（可选）

```go
    xmhsdk.SetLogger({your logger impl xmhsdk.LeveledLoggerInterface})

```

### 3. 配置token存储（可选）

```go 
默认使用本机内存存储，强烈建议使用redis存储token
xmhsdk.SetKvStorage(your kvStorage impl xmhsdk.Storage)

定时任务调用auth.New()刷新token


```

###商品同步
参考 [2.2 商品同步](https://iqir63fbchy.feishu.cn/docx/LAKxdsqfUoO0SyxQ8VYcpZhTncp#share-AonXdE2D0oWfyTxglSQcSH5fnte)

###试算报价
参考 [2.3 试算和风控](https://iqir63fbchy.feishu.cn/docx/LAKxdsqfUoO0SyxQ8VYcpZhTncp#share-LemGdmoVfoel7BxkAB2cft5QnFf)

###订单同步
参考 [2.4 订单同步](https://iqir63fbchy.feishu.cn/docx/LAKxdsqfUoO0SyxQ8VYcpZhTncp#share-B2VndFGLaoCI54xdy9acuUCgnkb)

###发货同步
参考 [2.5 发货同步](https://iqir63fbchy.feishu.cn/docx/LAKxdsqfUoO0SyxQ8VYcpZhTncp#share-JzPldZlhgoQzyDx4OO5cFatJnPb)

###取消投保
参考 [2.6 取消投保](https://iqir63fbchy.feishu.cn/docx/LAKxdsqfUoO0SyxQ8VYcpZhTncp#share-Bf1Ydeh56oJYiUxWL0ucywwgnmg)

###查询可理赔商品
参考 [2.6 查询可理赔商品](https://iqir63fbchy.feishu.cn/docx/LAKxdsqfUoO0SyxQ8VYcpZhTncp#share-XUZpdubwJo9g7axwonBcYOOPnHg)

###理赔报案
参考 [2.7 理赔报案](https://iqir63fbchy.feishu.cn/docx/LAKxdsqfUoO0SyxQ8VYcpZhTncp#share-RjGGdtREQo3mZ9xIxhRcDb0qnQc)




