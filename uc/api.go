package uc

// 公开
import "fmt"

func NewClient(cookie string) (info Info, client UCloud, err error) {
	info, client, err = core{}.isCookie(cookie)
	return
}

// GetShare 获取分享链接列表
func (u UCloud) GetShare(url, passcode string) {
	pwdId := ParseShareCode(url)

	resp, err := u.core.getShare(pwdId, passcode)
	fmt.Println(resp, err)
	return
}

// GetFolderNodes 获取我的文件夹
func (u UCloud) GetFolderNodes(pdirFid string) (resp FolderNodes, err error) {
	if pdirFid == "" {
		pdirFid = "0"
	}
	resp, err = u.core.getFolderNodes(pdirFid)
	return
}
