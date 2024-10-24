package _89

import (
	"github.com/spf13/cast"
	"net/url"
)

// 获取链接信息
func (c core) getShareInfoByCodeV2(shareCode string) (resp ShareInfoResp, err error) {
	path := "/open/share/getShareInfoByCodeV2.action"
	values := url.Values{}
	values.Set("shareCode", shareCode)
	err = c.invoker.Get(path, values, &resp)
	if err != nil {
		return
	}
	return
}

// 带密码的shareId
func (c core) checkAccessCode(shareCode, accessCode string) (resp checkAccessCode, err error) {
	path := "/open/share/checkAccessCode.action"
	values := url.Values{}
	values.Set("shareCode", shareCode)
	values.Set("accessCode", accessCode)
	err = c.invoker.Get(path, values, &resp)
	if err != nil {
		return
	}
	return
}

// 获取链接文件信息
func (c core) listShareDir(req ShareInfoResp) (resp ListShareDirResp, err error) {
	path := "/open/share/listShareDir.action"
	values := url.Values{}
	values.Set("pageNum", "1")
	values.Set("pageSize", "60")
	values.Set("fileId", req.FileId)
	values.Set("shareDirFileId", req.FileId)
	values.Set("isFolder", "true")
	values.Set("shareId", cast.ToString(req.ShareId))
	values.Set("shareMode", cast.ToString(req.ShareMode))
	values.Set("iconOption", "5")
	values.Set("orderBy", "lastOpTime")
	values.Set("descending", "true")
	values.Set("accessCode", req.AccessCode)
	err = c.invoker.Get(path, values, &resp)
	if err != nil {
		return
	}
	return
}
