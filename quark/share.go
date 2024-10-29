package quark

import (
	"net/url"
)

func (c core) sharePage(shareCode, passcode string) (stoken string, err error) {
	values := url.Values{}

	values.Set("uc_param_str", "")
	values.Set("__t", GetTimestamp())

	body := make(map[string]any)
	body["pwd_id"] = shareCode
	body["passcode"] = passcode

	var r sharePageResp
	err = c.invoker.Post("https://drive-h.quark.cn", "/1/clouddrive/share/sharepage/token", values, body, &r)
	if err != nil {
		return
	}
	return r.Data.Stoken, nil
}
func (c core) shareDetail(shareCode, pdirFid, stoken string) (resp ShareDetailResp, err error) {
	values := url.Values{}

	values.Add("uc_param_str", "")
	values.Add("pwd_id", shareCode)
	values.Add("stoken", stoken)
	values.Add("pdir_fid", pdirFid)
	values.Add("force", "0")
	values.Add("_page", "1")
	values.Add("_size", "50")
	values.Add("_fetch_banner", "1")
	values.Add("_fetch_share", "1")
	values.Add("_fetch_total", "1")
	values.Add("_sort", "file_type:asc,updated_at:desc")
	values.Add("__t", GetTimestamp())

	err = c.invoker.Get("https://drive-h.quark.cn", "/1/clouddrive/share/sharepage/detail", values, &resp)

	return
}
func (c core) shareSave(fileInfo fileInfoResp, pwdId, stoken, toPdirFid string) (taskInfo TaskInfo, err error) {
	values := url.Values{}
	values.Add("__t", GetTimestamp())
	values.Add("uc_param_str", "")

	body := make(map[string]any)
	body["fid_list"] = fileInfo.FidList
	body["fid_token_list"] = fileInfo.FidTokenList
	body["pdir_fid"] = "0"
	body["pwd_id"] = pwdId
	body["scene"] = "link"
	body["stoken"] = stoken
	body["to_pdir_fid"] = toPdirFid
	var resp saveInfoResp
	err = c.invoker.Post("https://drive-pc.quark.cn", "/1/clouddrive/share/sharepage/save", values, body, &resp)
	if err != nil {
		return
	}
	taskInfo, err = c.task(resp.Data.TaskId, 3)
	return
}
