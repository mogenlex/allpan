package quark

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

// GetMyDirNode 获取我的目录节点
func (c Quark) GetMyDirNode(pdirFid string) (resp MyNodeResp, err error) {
	resp, err = c.core.getMyNode(pdirFid)
	return
}

// ShareSave 分享文件转存
func (c Quark) ShareSave(info fileInfoResp, pwdId, stoken, toPdirFid string) (taskInfo TaskInfo, err error) {
	taskInfo, err = c.core.shareSave(info, pwdId, stoken, toPdirFid)
	return
}