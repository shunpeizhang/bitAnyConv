package main

import (
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"com"
)



func main(){
	param := make(map[string]string)
	param["start"] = "1"
	param["limit"] = "100"
	param["convert"] = "USD"

	res, err := com.GetDataByHttp("cryptocurrency/listings/latest", param)
	fmt.Println(res)
	return

	url := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?start=1&limit=100&convert=USD"

	req := HttpRequest.NewRequest()
	req.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"X-CMC_PRO_API_KEY": "0a5fcb4e-2b4e-4681-9814-731af7a5222b",
	})

	fmt.Println(req)

	ress, err := req.Get(url, nil)
	if nil != err{
		fmt.Println("get failed!", err.Error())
		return
	}
	body, err := ress.Body()

	fmt.Println(string(body))



}




