package cloud189

// NewCloud189 创建一个Client
func NewCloud189(account, password string) (cloud189 *Cloud189, err error) {
	cloud189, err = core{}.login(account, password)
	if err != nil {
		return
	}
	return
}

// GetShare 获取分享链接信息
func (c Cloud189) GetShare(url, accessCode string) (info ShareInfoResp, err error) {
	// 解析分享url 获取shareCode
	shareCode := ParseShareCode(url)

	// 获取链接信息
	info, err = c.core.getShareInfoByCodeV2(shareCode)
	info.AccessCode = accessCode
	if err != nil {
		return
	}

	if info.AccessCode != "" {
		resp, err := c.core.checkAccessCode(shareCode, accessCode)
		if err != nil {
			return info, err
		}
		info.ShareId = resp.ShareId
	}
	return
}

// GetListShareDir 获取链接目录列表
func (c Cloud189) GetListShareDir(req ShareInfoResp) (resp ListShareDirResp, err error) {
	resp, err = c.core.listShareDir(req)
	return
}

// GetMyFolderNodes  获取我的文件夹节点
func (c Cloud189) GetMyFolderNodes(id string) (resp []GetObjectFolderNodesResp, err error) {
	resp, err = c.core.getObjectFolderNodes(id)
	return
}

// CreateFolder 创建文件夹
func (c Cloud189) CreateFolder(parentFolderId, folderName string) (resp CreateFolderResp, err error) {
	resp, err = c.core.createFolder(parentFolderId, folderName)
	return
}

// SaveCurrentCatalogue 转存
func (c Cloud189) SaveCurrentCatalogue(taskInfos []TaskInfosReq, targetFolderId, shareId string) (resp CreateBatchTaskResp, err error) {
	resp, err = c.core.createBatchTask(taskInfos, targetFolderId, shareId, 3)
	return
}

// Sign 签到
func (c Cloud189) Sign() {

}
