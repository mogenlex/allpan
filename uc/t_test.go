package uc

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cast"
	"testing"
)

var cookie = "_UP_28A_52_=529; _UP_A4A_11_=wb96a13845b848c1baf7b2628c97aa21; _UP_D_=pc; tfstk=fI2-Qn6G8ZbowkKvMB1mtV0OkqscsJEykzr6K20kOrUY-kZoOYDHpW3YX8cCzk4xJyqEP8morDaLoqknr8Vu93yYWpvlz_rBpvDCSNXGIu-rLv_MJb8ahUonbL6nOX9XaO4w4NXGIu_GcTNOS8XJ8-giAviSdpTfD00jADaIPjTjXckIdvaIlxiofetSNXOfhqoId2aIdNMCJqtS-pFFijS1ULf2Id2xcw3z23pKm8m-wVZ7ep9I_0h-5ugV7g9ykjZE9J-ep-IrkR2T7pYiWDA5DideTbiVFo_6dJ0eBloxSgBpTBlV0m3GDlAeTbJKDVj27BREwI5..; __uid=AATIHBT57Vsexifzsp5J6lr9; __kps=AATIHBT57Vsexifzsp5J6lr9; __ktd=kz7Z0hA0VannXAJqzKaaTA==; __pus=1a2b4712886418913cfb5245bb677a66AATy3FIJLWf6VN01ikfgrueQql8PM7nZsW32yt089ZLw/jQTwmRzlDIh163gARw9QioXXU0DqF86SDL2vi+5ENzkoMDqDOW+L8GFZs65+ap7J4Sn/lHEH1ZX9yAv5lDUd/M=; __kp=41ba5520-9544-11ef-9825-17457c22fcf4; __puus=363d51577434c5e71462e0a3362ba9b6AAS99f6/5aFo4333EN8PcmvXJIyvFijn7EgNqBIO6EPM7ooXujOscSFX93bplNGVtfWe8EGRsd1akTb9HieDI3O/OfnrPDB3ApfNuCGteJYQ/Ynb8KRWkfoh9gIBnoeE36tNjrMShEYc9eHcHdVoci1UFSa+/BF8mpwYYc+/0y3PdvDNe92MfUlhyzIVV3uo0sw="

func login() UCloud {
	info, client, err := NewClient(cookie)
	if err != nil {
		panic(err)
	}
	marshal, err := jsoniter.Marshal(info)
	fmt.Println("======info=====", cast.ToString(marshal))
	return client
}

func TestInfo(t *testing.T) {
	login()
}
func TestShareLink(t *testing.T) {
	c := login()
	c.GetShare("https://drive.uc.cn/s/db7c5443b7f34#/list/share", "")
}

// 只获取我的目录
func TestGetFolderNodes(t *testing.T) {
	c := login()
	c.GetFolderNodes("")
}
