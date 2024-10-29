package quark

import "net/url"

// 获取我的目录
func (c core) getMyNode(pdirFid string) (resp MyNodeResp, err error) {
	values := url.Values{}
	values.Add("uc_param_str", "")
	values.Add("pdir_fid", pdirFid)
	values.Add("_page", "1")
	values.Add("_size", "100")
	values.Add("_fetch_total", "false")
	values.Add("_fetch_sub_dirs", "1")
	values.Add("_sort", "")
	values.Add("__t", GetTimestamp())

	err = c.invoker.Get("https://drive-pc.quark.cn", "/1/clouddrive/file/sort", values, &resp)
	return
}
