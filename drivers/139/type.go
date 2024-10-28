package cloud139

import "github.com/imroc/req/v3"

type core struct {
	invoker invoker
}
type invoker struct {
	client *req.Client
}
type cloud139 struct {
	core core
}
