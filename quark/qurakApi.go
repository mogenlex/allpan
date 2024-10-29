package quark

import (
	"errors"
	"fmt"
	"github.com/imroc/req/v3"
	"net/url"
	"time"
)

// 获取我的目录
func (c core) getObjectFolderNodes(pdirFid string) (resp MyNodeResp, err error) {
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

// 获取分享页的stoken
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

// 获取分享页的详情
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

// 创建转存任务ID
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

// 执行任务
func (c core) task(taskId string, maxRetries int) (taskInfo TaskInfo, err error) {
	values := url.Values{}
	values.Add("task_id", taskId)
	values.Add("uc_param_str", "")
	values.Add("retry_index", "0")
	values.Add("__t", GetTimestamp())
	err = c.invoker.Get("https://drive-pc.quark.cn", "/1/clouddrive/task", values, &taskInfo)

	switch taskInfo.Data.Status {
	case 2:
		return taskInfo, nil
	}

	if taskInfo.Code == 32003 && taskInfo.Status == 400 {
		err = errors.New("空间不足")
		return
	} else if taskInfo.Code == 0 && taskInfo.Status == 200 && maxRetries > 0 {
		time.Sleep(1 * time.Second)
		return c.task(taskId, maxRetries-1)
	}

	return
}

// NewQuark 函数用于创建一个新的 Quark 实例
func (c core) login(cookie string) (Quark, error) {
	// 初始化 URL 参数
	values := url.Values{}
	values.Set("fr", "pc")
	values.Set("platform", "pc")

	// 创建一个 HTTP 客户端
	client := req.C()
	client.QueryParams = values
	client.SetCommonHeader("Cookie", cookie)

	// 发送 GET 请求获取账户信息
	path := "/account/info"
	resp, err := client.R().Get(baseUrl + path) // 替换为实际的 baseUrl
	if err != nil {
		return Quark{}, fmt.Errorf("failed to get account info: %w", err)
	}

	// 解析响应
	var r infoResp
	err = resp.Into(&r)
	if err != nil {
		return Quark{}, fmt.Errorf("failed to parse response: %w", err)
	}

	// 检查用户是否登录成功
	if r.Data.Nickname == "" {
		return Quark{}, errors.New("login error")
	}

	// 创建并返回 Quark 实例
	return Quark{
		core: core{
			invoker: invoker{
				client: client,
			},
		},
	}, nil
}
