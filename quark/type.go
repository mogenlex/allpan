package quark

import "github.com/imroc/req/v3"

// Quark 结构体表示一个 Quark 实例
type Quark struct {
	core core
}

// core 结构体包含一个 invoker 实例
type core struct {
	invoker invoker
}

type invoker struct {
	client *req.Client
}
