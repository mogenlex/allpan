package cloud189

import (
	"errors"
	"fmt"
	"github.com/imroc/req/v3"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cast"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var (
	mu sync.Mutex
)

// 登录
func (c core) login(account, password string) (cc *req.Client, err error) {
	mu.Lock()
	defer mu.Unlock()
	client := req.C()
	tempUrl := "https://cloud.189.cn/api/portal/loginUrl.action?redirectURL=https%3A%2F%2Fcloud.189.cn%2Fmain.action"
	var lt, reqId string

	resp, err := client.
		SetRedirectPolicy(req.RedirectPolicy(func(req *http.Request, via []*http.Request) error {
			if req.URL.Query().Get("lt") != "" {
				lt = req.URL.Query().Get("lt")
			}
			if req.URL.Query().Get("reqId") != "" {
				reqId = req.URL.Query().Get("reqId")
			}
			return nil
		})).R().Get(tempUrl)

	if err != nil {
		return
	}

	cookies := resp.Cookies()
	client.SetRedirectPolicy(req.MaxRedirectPolicy(10))

	appConfResp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:76.0) Gecko/20100101 Firefox/74.0",
			"Referer":    resp.Request.URL.String(),
			"Cookie":     cookiesToString(cookies),
			"lt":         lt,
			"reqId":      reqId,
		}).
		SetFormData(map[string]string{
			"appKey":  "cloud",
			"version": "2.0",
		}).
		Post("https://open.e.189.cn/api/logbox/oauth2/appConf.do")
	if err != nil {
		return
	}

	data := jsoniter.Get(appConfResp.Bytes(), "data")
	accountType := data.Get("accountType").ToString()
	clientType := data.Get("clientType").ToString()
	paramId := data.Get("paramId").ToString()
	mailSuffix := data.Get("mailSuffix").ToString()
	returnUrl := data.Get("returnUrl").ToString()

	encryptConfResp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:76.0) Gecko/20100101 Firefox/74.0",
			"Referer":    "https://open.e.189.cn/api/logbox/separate/web/index.html",
			"Cookie":     cookiesToString(cookies),
		}).
		SetFormData(map[string]string{
			"appId": "cloud",
		}).
		Post("https://open.e.189.cn/api/logbox/config/encryptConf.do")
	if err != nil {
		return
	}

	if resCode := jsoniter.Get(encryptConfResp.Bytes(), "result").ToInt(); resCode != 0 {
		err = fmt.Errorf("Failed to get encrypt config")
		return
	}

	encryptData := jsoniter.Get(encryptConfResp.Bytes(), "data")
	pubKey := encryptData.Get("pubKey").ToString()
	pre := encryptData.Get("pre").ToString()
	vCodeRS := ""
	userRsa := RsaEncode([]byte(account), pubKey)
	passwordRsa := RsaEncode([]byte(password), pubKey)

	loginResp, err := client.R().
		SetHeaders(map[string]string{
			"lt":         lt,
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:74.0) Gecko/20100101 Firefox/76.0",
			"Referer":    "https://open.e.189.cn/",
		}).
		SetFormData(map[string]string{
			"version":      "v2.0",
			"appKey":       "cloud",
			"accountType":  accountType,
			"userName":     pre + userRsa,
			"epd":          pre + passwordRsa,
			"validateCode": vCodeRS,
			"captchaToken": "",
			"returnUrl":    returnUrl,
			"mailSuffix":   mailSuffix,
			"paramId":      paramId,
			"clientType":   clientType,
			"dynamicCheck": "FALSE",
			"cb_SaveName":  "1",
			"isOauth2":     "false",
		}).
		Post("https://open.e.189.cn/api/logbox/oauth2/loginSubmit.do")
	if err != nil {
		return
	}

	if restCode := jsoniter.Get(loginResp.Bytes(), "result").ToInt(); restCode == 0 {
		toUrl := jsoniter.Get(loginResp.Bytes(), "toUrl").ToString()

		resp, err = client.SetRedirectPolicy(req.MaxRedirectPolicy(10)).R().Get(toUrl)
		if err != nil {
			return
		}

		resp, err = client.R().Get("https://cloud.189.cn/v2/getUserBriefInfo.action?noCache=" + Random())
		if err != nil {
			return
		}

		_ = jsoniter.Get(resp.Bytes(), "sessionKey").ToString()
		resp, err = client.SetRedirectPolicy(req.MaxRedirectPolicy(5)).R().Get("https://api.cloud.189.cn/open/oauth2/ssoH5.action")
		if err != nil {
			return
		}

		//fmt.Println(resp.HeaderToString())

		//v, _ := url.ParseQuery(resp.Header.Get("location"))
		//accessToken := v.Get("https://h5.cloud.189.cn/index.html?accessToken")

		return client, err
		//fmt.Println(sessionKey, "login cloud189", "accessToken", accessToken)
	} else if restCode == -2 {
		err = errors.New(jsoniter.Get(loginResp.Bytes(), "msg").ToString())
		return
	} else {
		err = errors.New(jsoniter.Get(loginResp.Bytes(), "msg").ToString())
		return
	}
}

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

// 获取带密码的shareId
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

// 只获取我的目录 目录id
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

// CreateBatchTask 创建批量任务Id接口
func (c core) createBatchTask(taskInfos []TaskInfosReq, targetFolderId, shareId string, maxRetries int) (resp CreateBatchTaskResp, err error) {
	marshal, err := jsoniter.Marshal(taskInfos)
	fmt.Println(string(marshal), taskInfos)
	path := "/open/batch/createBatchTask.action"
	values := url.Values{}
	values.Set("type", "SHARE_SAVE")
	values.Set("taskInfos", string(marshal))
	values.Set("targetFolderId", targetFolderId)
	values.Set("shareId", shareId)
	fmt.Println(values)
	err = c.invoker.Post(path, values, &resp)
	if err != nil {
		return
	}
	resp, err = c.checkBatchTask(resp.TaskId, maxRetries)
	return
}

// checkBatchTask 执行任务
func (c core) checkBatchTask(taskId string, maxRetries int) (resp CreateBatchTaskResp, err error) {
	path := "/open/batch/checkBatchTask.action"
	values := url.Values{}
	values.Set("taskId", taskId)
	values.Set("type", "SHARE_SAVE")
	err = c.invoker.Post(path, values, &resp)

	fmt.Println(resp, taskId)
	if err != nil {
		return
	}
	switch resp.TaskStatus {
	case -1:
		err = errors.New(resp.ResMessage)
		return
	case 2:
		return
	case 4:
		return
	}

	if resp.SubTaskCount != resp.SuccessedCount && maxRetries > 0 {
		time.Sleep(1 * time.Second)
		return c.checkBatchTask(taskId, maxRetries-1)
	}
	return
}
