package quark

import (
	"errors"
	"fmt"
	"github.com/imroc/req/v3"
	"net/url"
)

// Quark 结构体表示一个 Quark 实例
type Quark struct {
	core core
}

// core 结构体包含一个 invoker 实例
type core struct {
	invoker invoker
}

// NewQuark 函数用于创建一个新的 Quark 实例
func NewQuark(cookie string) (Quark, error) {
	// 初始化 URL 参数
	values := url.Values{}
	values.Set("fr", "pc")
	values.Set("platform", "pc")

	// 创建一个 HTTP 客户端
	client := req.C()
	client.QueryParams = values
	client.SetCommonHeader("Cookie", cookie)

	// 发送 GET 请求获取账户信息
	path := "/account/info"
	resp, err := client.R().Get(baseUrl + path) // 替换为实际的 baseUrl
	if err != nil {
		return Quark{}, fmt.Errorf("failed to get account info: %w", err)
	}

	// 解析响应
	var r infoResp
	err = resp.Into(&r)
	if err != nil {
		return Quark{}, fmt.Errorf("failed to parse response: %w", err)
	}

	// 检查用户是否登录成功
	if r.Data.Nickname == "" {
		return Quark{}, errors.New("login error")
	}

	// 创建并返回 Quark 实例
	return Quark{
		core: core{
			invoker: invoker{
				client: client,
			},
		},
	}, nil
}
