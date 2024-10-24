package _89

import "net/url"

type core struct {
	invoker invoker
}

// 创建文件夹
func (c core) createFolder(parentFolderId, folderName string) (resp CreateFolderResp, err error) {
	path := "/open/file/createFolder.action"
	values := url.Values{}
	values.Set("parentFolderId", parentFolderId)
	values.Set("folderName", folderName)
	err = c.invoker.Post(path, values, &resp)
	if err != nil {
		return
	}
	return
}
