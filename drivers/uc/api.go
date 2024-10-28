package uc

// 公开
import "fmt"

func NewClient(cookie string) (info Info, client UCloud, err error) {
	info, client, err = core{}.isCookie(cookie)
	return
}

// GetShareLinkList 获取分享链接列表
func (u UCloud) GetShareLinkList(url, passcode string) {
	pwdId := ParseShareCode(url)

	resp, err := u.core.getShareLink(pwdId, passcode)
	fmt.Println(resp, err)
	return
}
