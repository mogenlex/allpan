package quark

import (
	"errors"
	"github.com/imroc/req/v3"
	jsoniter "github.com/json-iterator/go"
	"net/url"
)

const (
	baseUrl = "https://pan.quark.cn"
)

type invoker struct {
	client *req.Client
}

// handle400Error 处理400错误
func handle400Error(i *invoker, path string, params url.Values, data interface{}) error {
	return nil
}

//const cookie1 = "ctoken=nMbRshRXN0trRjrX0xhu-97V; b-user-id=a2c9ea31-e534-0d8a-675b-160af4b2fb2b; grey-id=ec950091-e39c-0fb4-ac36-75f62c830762; grey-id.sig=Z7Wf49PHyos32ntFZ5ZL4Amp_DyqcZ_BdH8IyHJUgbc; isQuark=true; isQuark.sig=hUgqObykqFom5Y09bll94T1sS9abT1X-4Df_lzgl8nM; __wpkreporterwid_=8c358d29-b6db-4019-28f9-76d0b9c21c91; _UP_A4A_11_=wb96a128ce55457ea2cc95945d87e4b6; _UP_D_=pc; _UP_30C_6A_=st96a620197nb6xgklsq2ryx895h3egw; _UP_TS_=sg1be132b1cfb78bef9af0eeb9bf0a0c25e; _UP_E37_B7_=sg1be132b1cfb78bef9af0eeb9bf0a0c25e; _UP_TG_=st96a620197nb6xgklsq2ryx895h3egw; _UP_335_2B_=1; __pus=c11e16d728126816a758207921c0248cAATsRU8Ae+xOOieWPH5HWahIpStieaZjLKct/x4E1pmg833DW7yizu3bNG9X6k47cpHJMxuKAxRUj5sOVgpqkRGn; __kp=a50d7ee0-9459-11ef-be06-fdb184d33df9; __kps=AATrbC2XxtcPZEXZtA22lSZZ; __ktd=uGRoTefKBPGQ1VcV7XIuLw==; __uid=AATrbC2XxtcPZEXZtA22lSZZ; tfstk=fUMEC6vdOppeGMJ5Fkyz05e9SZ2LUJYjtYa7q0muAy43AJio_0i2vugIqzlzS4KddJYdEzmxWWanPJ_yUVy3wXa7Rzyr2q-6lK9jvDe8KETXhzjf91ef-ab7Z5jiavqhvK9jvDcG-8ryhvGCbMz4r8V3ZO2gWk2lKW0uS5q0cw2or8xasoZ8r943-l4gDP4uE80ojjBqx0XacWxX0awJme8btrm3-DiiIx_brD4NET4UYWoZxPWlEAc_HRA_-II_yJaKfkuJH9e3ajn742JFK4oxBVrioLWb-cnjGWHHH_nLKyyiKjSlEcyU-SDiSw8aovniHRV6nnoiB2iKQ0skElgbSDH3UKxYLJ43QlHWW94rsbDQ98QHlWno_VrUng71ylxFPYhFr_V3XlzXbhlm4a3O-nak__C8s-EalHZCw_F3plzXbpCRwWd8brtpA; __itrace_wid=b9ccf2df-a115-4324-1bc6-87f4beb6458a; __puus=2e0b21d36e3dae2497e6bea15534c053AASv10YgRjNpvh13sJD7xgSLP0dgKAb5Sg3P4tpY7AvJAFHpEoPxNw86/NiIRqeU4VVRSwhiMWJEatoldyNhqHxrVHq8rPiEbaTfoyPb4sNRv5WnCZ8xmbAExRleQp9kKS5ZDyQRo88f0mcs50qUmpOTRbVGk/hsrRsVpkVq+EO7jMwQ5SDfnyxvXytHsZrOLEei3p060BkZqvmUTptGvoMn"

func (i *invoker) Get(url, path string, params url.Values, data interface{}) error {

	client := i.client.SetBaseURL(url).DevMode().R()

	params.Add("pr", "ucpro")
	params.Add("fr", "pc")

	client.QueryParams = params

	err := i.do(client, "GET", path, &data)
	return err
}
func (i *invoker) do(client *req.Request, method string, path string, data interface{}) (err error) {
	resp, err := client.Send(method, path)

	if resp.StatusCode == 400 {
		resMessage := jsoniter.Get(resp.Bytes(), "res_message").ToString()
		return errors.New(resMessage)
	}

	if err != nil {
		return
	}
	err = resp.Into(&data)
	if err != nil {
		return
	}
	return
}

func (i *invoker) Post(url, path string, params url.Values, body map[string]interface{}, data interface{}) error {

	client := i.client.DevMode().SetBaseURL(url).R()

	params.Add("pr", "ucpro")
	params.Add("fr", "pc")

	client.QueryParams = params

	client.SetBody(body)

	client.SetHeader("Accept", "application/json, text/plain, */*")
	client.SetHeader("Content-Type", "application/json")
	err := i.do(client, "POST", path, &data)
	return err
}
