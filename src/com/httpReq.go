package com

import (
	"github.com/kirinlabs/HttpRequest"
	"github.com/golang/glog"
	"fmt"
)



var (
	CMC_PRO_API_KEY = "0a5fcb4e-2b4e-4681-9814-731af7a5222b"
	URL_COMMON = "https://pro-api.coinmarketcap.com/v1/"
)

//methon: "cryptocurrency/listings/latest"
func GetDataByHttp(methon string, param map[string]string) (string, error){
	strParam := ""
	for k, v := range param{
		if 0 == len(strParam){
			strParam = fmt.Sprintf("%v%v=%v", strParam, k, v)
		}else{
			strParam = fmt.Sprintf("%v&%v=%v", strParam, k, v)
		}
	}

	url := URL_COMMON + methon + "?" + strParam
	fmt.Println(url)

	req := HttpRequest.NewRequest()
	req.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"X-CMC_PRO_API_KEY": CMC_PRO_API_KEY,
	})

	fmt.Println(req)

	res, err := req.Get(url, nil)
	if nil != err{
		glog.Error("get failed!", err.Error())
		return "", err
	}

	body, err := res.Body()
	if nil != err{
		glog.Error("res.Body failed!", err.Error(), res)
		return "", err
	}

	return string(body), nil
}










