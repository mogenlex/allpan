package uc

// 内部
import (
	"errors"
	"github.com/imroc/req/v3"
	"net/url"
)

// 验证cookie 是否有效
func (c core) isCookie(cookie string) (info Info, uc UCloud, err error) {
	client := req.C()
	values := url.Values{}
	values.Add("fr", "pc")
	values.Add("platform", "pc")
	client.SetCommonHeader("Cookie", cookie)
	client.QueryParams = values
	resp, err := client.R().Get("https://drive.uc.cn/account/info")
	if err != nil {
		return
	}
	err = resp.Into(&info)
	if err != nil {
		return
	}

	if info.Code == "AUTH_ERROR" {
		err = errors.New("AUTH_ERROR")
		return
	} else {
		return info, UCloud{
				core: core{
					invoker: invoker{
						client: client}}},
			nil
	}

}

// 获取分享链接详情
func (c core) getShare(pwdId, passcode string) (shareInfo shareInfo, err error) {
	values := url.Values{}

	body := map[string]any{
		"banner_platform": "other", "fetch_banner": "1", "fetch_error_background": "1", "fetch_share": "1", "fetch_total": "1", "force": "0", "page": "1",
		"passcode": passcode,
		"pwd_id":   pwdId,
		"size":     "50", "sort": "file_type:asc,file_name:asc", "web_platform": "windows",
	}

	path := "/1/clouddrive/share/sharepage/v2/detail"

	err = c.invoker.Post("https://pc-api.uc.cn", path, values, body, &shareInfo)
	return
}

// 获取分享目录节点
func (c core) getShareNote(pwdId, stoken, pdirFid string) (info SharePageResp, err error) {
	values := url.Values{}

	values.Add("pwd_id", pwdId)
	values.Add("stoken", stoken)
	values.Add("pdir_fid", pdirFid)
	values.Add("pdir_fid", "0")
	values.Add("_page", "1")
	values.Add("_size", "50")
	values.Add("_fetch_banner", "0")
	values.Add("_fetch_share", "0")
	values.Add("_fetch_total", "1")
	values.Add("_sort", "file_type:asc,file_name:asc")

	path := "/1/clouddrive/share/sharepage/detail"

	err = c.invoker.Get("https://pc-api.uc.cn", path, values, &info)
	return
}

// 只获取文件夹
func (c core) getFolderNodes(pdirFid string) (resp FolderNodes, err error) {
	values := url.Values{}
	values.Set("sys", "win32")
	values.Set("pdir_fid", pdirFid)
	values.Set("_page", "1")
	values.Set("_size", "100")
	values.Set("_fetch_total", "false")
	values.Set("_fetch_sub_dirs", "1")
	values.Set("_sort", "")
	values.Set("recent_file_size", "0")
	values.Set("ve", "1.7.2")
	path := "/1/clouddrive/file/sort"
	err = c.invoker.Get("https://pc-api.uc.cn", path, values, &resp)
	return
}
