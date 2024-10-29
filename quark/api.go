package quark

func NewQuark(cookie string) (quark Quark, err error) {
	quark, err = core{}.login(cookie)
	return
}

// GetShare 获取分享文件信息
func (q Quark) GetShare(url, passcode, pdirFid string) (resp ShareDetailResp, err error) {
	shareCode := ParseShareCode(url)
	stoken, err := q.core.sharePage(shareCode, passcode)
	if err != nil {
		return
	}
	if pdirFid == "" {
		pdirFid = "0"
	}

	resp, err = q.core.shareDetail(shareCode, pdirFid, stoken)
	if err != nil {
		return
	}
	return
}

// GetMyFolderNodes 获取我的目录节点
func (q Quark) GetMyFolderNodes(pdirFid string) (resp MyNodeResp, err error) {
	resp, err = q.core.getObjectFolderNodes(pdirFid)
	return
}

// SaveCurrentCatalogue 分享文件转存
func (q Quark) SaveCurrentCatalogue(info fileInfoResp, pwdId, stoken, toPdirFid string) (taskInfo TaskInfo, err error) {
	taskInfo, err = q.core.shareSave(info, pwdId, stoken, toPdirFid)
	return
}
