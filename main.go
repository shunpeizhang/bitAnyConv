package main

import (
	"bitAnyConv/CMCDataInterface"
	"bitAnyConv/infoStruct"
	"time"
)

func main(){
	//param := make(map[string]string)
	//param["start"] = "1"
	//param["limit"] = "100"
	//param["convert"] = "USD"
	//
	//res, err := com.GetDataByHttp("cryptocurrency/listings/latest", param)
	//fmt.Println(res, err)


	info := infoStruct.STCoinMetadata{}
	info.Symbol = "BTC"
	info.ID = 1
	info.Cmc_rank = 1
	info.Date_added = time.Now()
	info.Max_supply = 1
	info.Name = "Bitcoin"
	info.Slug = "bitcoin"
	info.Total_supply = 1

	CMCDataInterface.GetCoinHistory(&info, time.Unix(1534052593, 0), time.Unix(1534053193, 0), 10)


}




