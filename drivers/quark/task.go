package quark

import (
	"errors"
	"net/url"
	"time"
)

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
