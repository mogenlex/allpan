package uc

import "github.com/imroc/req/v3"

type core struct {
	invoker invoker
}
type invoker struct {
	client *req.Client
}
type UCloud struct {
	core core
}
