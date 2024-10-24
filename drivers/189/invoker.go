package cloud189

import (
	"errors"
	"github.com/imroc/req/v3"
	jsoniter "github.com/json-iterator/go"
	"net/url"
)

const (
	baseUrl = "https://cloud.189.cn/api"
)

type invoker struct {
	client *req.Client
}

// handle400Error 处理400错误
func handle400Error(i *invoker, path string, params url.Values, data interface{}) error {
	return nil
}

func (i *invoker) Get(path string, params url.Values, data interface{}) error {

	client := i.client.SetBaseURL(baseUrl).DevMode().R()

	client.QueryParams = params
	client.SetQueryParam("noCache", Random())
	client.SetHeader("Accept", "application/json;charset=UTF-8")

	err := i.do(client, "GET", path, &data)
	return err
}
func (i *invoker) do(client *req.Request, method string, path string, data interface{}) (err error) {
	resp, err := client.Send(method, path)

	if resp.StatusCode == 400 {
		resMessage := jsoniter.Get(resp.Bytes(), "res_message").ToString()
		return errors.New(resMessage)
	}

	if err != nil {
		return
	}
	err = resp.Into(&data)
	if err != nil {
		return
	}
	return
}

func (i *invoker) Post(path string, params url.Values, data interface{}) error {

	client := i.client.SetBaseURL(baseUrl).R()

	client.FormData = params
	client.SetQueryParam("noCache", Random())
	client.SetHeader("Accept", "application/json;charset=UTF-8")
	client.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	err := i.do(client, "POST", path, &data)
	return err
}
