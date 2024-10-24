package _89

import "net/url"

type Invoker struct {
}

// handle400Error 处理400错误
func handle400Error(i *Invoker, path string, params url.Values, data interface{}) error {
	return nil
}

func (i *Invoker) Get(path string, params url.Values, data interface{}) error {
	return nil
}
