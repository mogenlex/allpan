package _89

import (
	"fmt"
	"testing"
)

var c Cloud189

func newLogin() Cloud189 {
	client, err := Login("15315496321", "mogen123")
	if err != nil {
		panic(err)
	}

	return Cloud189{core{invoker{client}}}

}

func TestLogin(t *testing.T) {
	newLogin()

}
func TestGetShare(t *testing.T) {
	c := newLogin()
	info, err := c.GetShare("https://cloud.189.cn/web/share?code=yqUzEzea2mam", "x39p")
	fmt.Println(info, err)
	resp, err := c.GetListShareDir(info)
	fmt.Println(resp, err)
}

func TestNodes(t *testing.T) {
	c := newLogin()
	resp, err := c.GetObjectFolderNodes("524853159516486717")
	fmt.Println(resp, err)
}

func TestName(t *testing.T) {
	c := newLogin()
	resp, err := c.CreateFolder("-13", "test111")
	fmt.Println(resp, err)
}
