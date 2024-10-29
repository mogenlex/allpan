package cloud189

import (
	"fmt"
	"testing"
)

func newLogin() *Cloud189 {
	cloud189, err := NewCloud189("15315496321", "mogen123")
	if err != nil {
		panic(err)
	}
	return cloud189
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
	resp, err := c.GetMyFolderNodes("-13")
	fmt.Println(resp, err)
}

func TestCreateFolder(t *testing.T) {
	c := newLogin()
	resp, err := c.CreateFolder("-13", "test111")
	fmt.Println(resp, err)
}
func TestName(t *testing.T) {
	c := newLogin()
	var taskInfos []TaskInfosReq

	taskInfos = append(taskInfos, TaskInfosReq{
		FileId:   "223121159946435791",
		FileName: "胆大党（2024）",
		IsFolder: 1,
	})

	resp, err := c.SaveCurrentCatalogue(taskInfos, "12423112875931", "-13")
	fmt.Println(resp, err, taskInfos)
}
