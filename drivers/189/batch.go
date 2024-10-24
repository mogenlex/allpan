package cloud189

import (
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"net/url"
	"time"
)

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
