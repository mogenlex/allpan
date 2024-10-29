package cloud189

// 返回结构体
// 验证密码链接Resp
type checkAccessCode struct {
	ShareId int64 `json:"shareId"`
}

type ShareInfoResp struct {
	ShareId        int64  `json:"shareId"`
	FileId         string `json:"fileId"`
	AccessCode     string `json:"accessCode"`
	ShareMode      int32  `json:"shareMode"`
	ShareDirFileId string `json:"shareDirFileId"`
	FileName       string `json:"fileName"`
}

type ListShareDirResp struct {
	ResCode    int    `json:"res_code"`
	ResMessage string `json:"res_message"`
	ExpireTime int    `json:"expireTime"`
	ExpireType int    `json:"expireType"`
	FileListAO struct {
		Count        int          `json:"count"`
		FileList     []FileList   `json:"fileList"`
		FileListSize int          `json:"fileListSize"`
		FolderList   []FolderList `json:"folderList"`
	} `json:"fileListAO"`
	LastRev int64 `json:"lastRev"`
}
type FileList []struct {
	CreateDate   string `json:"createDate"`
	FileCata     int    `json:"fileCata"`
	FileListSize int    `json:"fileListSize"`
	Id           int64  `json:"id"`
	LastOpTime   string `json:"lastOpTime"`
	Name         string `json:"name"`
	ParentId     int64  `json:"parentId"`
	Rev          string `json:"rev"`
	StarLabel    int    `json:"starLabel"`
}
type FolderList struct {
	CreateDate   string `json:"createDate"`
	FileCata     int    `json:"fileCata"`
	FileListSize int    `json:"fileListSize"`
	Id           int64  `json:"id"`
	LastOpTime   string `json:"lastOpTime"`
	Name         string `json:"name"`
	ParentId     int64  `json:"parentId"`
	Rev          string `json:"rev"`
	StarLabel    int    `json:"starLabel"`
}

type GetObjectFolderNodesResp struct {
	IsParent string `json:"isParent"`
	Name     string `json:"name"`
	PId      string `json:"pId"`
	Id       string `json:"id"`
}

type CreateFolderResp struct {
	ResCode      int    `json:"res_code"`
	ResMessage   string `json:"res_message"`
	CreateDate   string `json:"createDate"`
	FileCata     int    `json:"fileCata"`
	FileListSize int    `json:"fileListSize"`
	Id           string `json:"id"`
	LastOpTime   string `json:"lastOpTime"`
	Name         string `json:"name"`
	ParentId     int    `json:"parentId"`
	Rev          string `json:"rev"`
}
type CreateBatchTaskResp struct {
	ResCode        int    `json:"res_code"`
	ResMessage     string `json:"res_message"`
	FailedCount    int    `json:"failedCount"`
	Process        int    `json:"process"`
	SkipCount      int    `json:"skipCount"`
	SubTaskCount   int    `json:"subTaskCount"`
	SuccessedCount int    `json:"successedCount"`
	TaskId         string `json:"taskId"`
	TaskStatus     int    `json:"taskStatus"`
}
