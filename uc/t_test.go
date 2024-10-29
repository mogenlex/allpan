package uc

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cast"
	"testing"
)

var cookie = "__wpkreporterwid_=46056e13-2eda-4d09-850c-c41cfaac5d10; ctoken=cMNJxOxNwzjXy0wKqIU1_S9Z; UDRIVE_TRANSFER_SESS=3GAMPP9YU8_a6FEywghLkkBMDxC2LC0jaYP5ZrA0oECfZs8FV67MUaz7Z8bkDMCeqp2wCEqNXhoM_Gny42kAqFV_-eRm7sKJiM-fliABrMD6jHNCmFNL5HUJoyX6pA4dY82oI2RH0VFc8lxhEoGCyIlem9E20h2QEIBq8jYvzao3Xwpdd_L67Pqt8EZSkwRm; b-user-id=9b6a43dc-9b9d-d0d2-c7ed-f16b3f4aecc2; __itrace_wid=f67bab5e-bcec-4d8a-a9ba-3486df173ecf; _UP_28A_52_=381; _UP_A4A_11_=wb96a12d3c1c4da38a92a2c2fbcc296c; _UP_D_=pc; _UP_F7E_8D_=i74i77uRTztsDFx1Aoa40qObll9KxKcTA9xfjsE6HO1NlFv12Q97EAGp4hxaNCZCdqwdLw3shJNLHpvA%2Fhicy1HUTu2LBlCPom4qTWeMFqCgN55FQx3lIyu%2B1OsWIKjG1w4pOGWcZjvuo4m2jHof9eRj66wpeTPO4r7NBD%2F4wEE0IpjIBHWretgcndtvmRjOaHNJLiTcrsXjhfgOBiOFPbzk4nczcPLP1rSgCMVm5ws2Z%2BPyTAVLm9UiDHFvPYOerR3OjYIJySBY1%2FFbsfLBenM3hDSUBvwcFj64bNwVNDsVOMEdD6uzZJXMxBnUatpyAHLu79tlMNqP8TGNMQXXgvSqK5ufzR58ZeivnehV0qE%2FWt1yDEDt%2BfWrmT4mVs6zZWXvqpzmoV3MeygIUCEakh2GAn6rsLT1sf4GAak5y386F8u7yQFbh%2F0Q7RCSfK2U6tAXQttwc%2FtDK7HYGyvolg%3D%3D; __pus=757b950e7731de6fb25cbaddec6f1864AATtYUEShybofmpXnGsgSvOUpKbBXSVltk9fU3vfU4DN/nRH86UuL7rK1G9GFzHi8OEKrG1gLzGkv/OwRjnguo8K; __kp=7f6b9eb0-9526-11ef-9d84-6b4791befde0; __kps=AATIHBT57Vsexifzsp5J6lr9; __ktd=kz7Z0hA0VannXAJqzKaaTA==; __uid=AATIHBT57Vsexifzsp5J6lr9; tfstk=fVnmdeX6JqzfO4QuEUrXSNagcMT8ltZ_OfIT6lFwz7PWh53xQfDZ9bUNkIMasVVr1KsY0Isa_vG7kSlTD1xiZozsk5FYslDstLdpvHHjhlZ29BKKycR8Qkw4QR5TzzyUv4cAlHHjhT7XaXv9v1cvfB3a_fz4zuybaRSq_SkyERyOgNSq_LDzNRNN7-SZUay3BsrabfkyERGA0GPnb0iyGiFYC94EG020zUH8aV36B8a0TxPkTBolbrVE37jwOVpDV5VSYQLoFfgqM8GMxQqsNVlaIb5y6rMZrb2mMBf_6mMxj8McpMMuSS0r0zXNbYr0ZX0iiIbbgqhuR8zPIHw-97kj04v1axlLioyzyHJnUPu-cyivaiqERxEx7fTdPoDZuujyUa7ezts_Udnl5Na4F8VLeC7y1BvA_KpkEwTbu8wv9Lvl5Na4F8VpELbCGry7HBC..; __puus=43db424d2304f40f8033d31d11204e0dAAS99f6/5aFo4333EN8PcmvXy7CEfwIdCGe2dIuDsVis+b6zD7dCJlV7CpG4G5PQJIhALNpGPdeQN6yhC1888v8mpCiKbtHWROODNnIB8MgY8oYjDjbB5BgWV7UEM6oGz1tvUY32aaWiPahZGeZxWwPiXDiXU57F0sS6iE8J46G0K0DNpSbeUpuCVg/8iKwo8UY="

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
	c.GetShareLinkList("https://drive.uc.cn/s/db7c5443b7f34#/list/share", "")
}
