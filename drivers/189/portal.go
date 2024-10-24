package _89

import "net/url"

// 只获取目录 目录id
func (c core) getObjectFolderNodes(id string) (resp []GetObjectFolderNodesResp, err error) {
	path := "/portal/getObjectFolderNodes.action"

	values := url.Values{}
	if id == "" {
		values.Set("id", "-11")
	} else {
		values.Set("id", id)
	}
	values.Set("orderBy", "1")
	values.Set("order", "ASC")
	err = c.invoker.Post(path, values, &resp)
	if err != nil {
		return
	}
	return
}
