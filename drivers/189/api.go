package _89

type Cloud189 struct {
	core core
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
	if err != nil {
		return
	}
	return
}

// GetObjectFolderNodes 获取对象文件夹节点
func (c Cloud189) GetObjectFolderNodes(id string) (resp []GetObjectFolderNodesResp, err error) {
	resp, err = c.core.getObjectFolderNodes(id)
	if err != nil {
		return
	}
	return
}

// CreateFolder 创建文件夹
func (c Cloud189) CreateFolder(parentFolderId, folderName string) (resp CreateFolderResp, err error) {
	resp, err = c.core.createFolder(parentFolderId, folderName)
	if err != nil {
		return
	}
	return
}
