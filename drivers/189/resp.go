package _89

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
		Count    int `json:"count"`
		FileList []struct {
			CreateDate string `json:"createDate"`
			FileCata   int    `json:"fileCata"`
			Icon       struct {
				LargeUrl string `json:"largeUrl"`
				SmallUrl string `json:"smallUrl"`
			} `json:"icon"`
			Id         string `json:"id"`
			LastOpTime string `json:"lastOpTime"`
			Md5        string `json:"md5"`
			MediaType  int    `json:"mediaType"`
			Name       string `json:"name"`
			Rev        string `json:"rev"`
			Size       int64  `json:"size"`
			StarLabel  int    `json:"starLabel"`
		} `json:"fileList"`
		FileListSize int `json:"fileListSize"`
		FolderList   [][]struct {
			CreateDate string `json:"createDate"`
			FileCata   int    `json:"fileCata"`
			Icon       struct {
				LargeUrl string `json:"largeUrl"`
				SmallUrl string `json:"smallUrl"`
			} `json:"icon"`
			Id         string `json:"id"`
			LastOpTime string `json:"lastOpTime"`
			Md5        string `json:"md5"`
			MediaType  int    `json:"mediaType"`
			Name       string `json:"name"`
			Rev        string `json:"rev"`
			Size       int64  `json:"size"`
			StarLabel  int    `json:"starLabel"`
		} `json:"folderList"`
	} `json:"fileListAO"`
	LastRev int64 `json:"lastRev"`
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
