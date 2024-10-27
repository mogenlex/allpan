package quark

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
func (c Quark) GetMyDirNode(pdirFid string) (resp MyNodeResp, err error) {
	resp, err = c.core.getMyNode(pdirFid)
	return
}
