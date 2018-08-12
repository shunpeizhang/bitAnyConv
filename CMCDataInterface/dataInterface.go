package CMCDataInterface

import (
	"container/list"
	"time"
	"bitAnyConv/infoStruct"
	"bitAnyConv/com"
	"fmt"
	"github.com/golang/glog"
)

//获取某种币的历史数据
func GetCoinHistory(metaInfo *infoStruct.STCoinMetadata, time_start time.Time, time_end time.Time, count int32) (*list.List, error){
	//数据请求
	var retData string
	{
		param := make(map[string]string)
		param["id"] = fmt.Sprintf("%v", metaInfo.ID)
		param["symbol"] = metaInfo.Symbol
		param["time_start"] = fmt.Sprintf("%v", time_start.Unix())
		param["time_end"] = fmt.Sprintf("%v", time_end.Unix())
		param["count"] = fmt.Sprintf("%v", count)
		param["convert"] = "USD"
		param["interval"] = "5m"

		var err error
		retData, err = com.GetDataByHttp("cryptocurrency/quotes/historical", param)
		if nil != err{
			glog.Error("com.GetDataByHttp failed!", param)
			return nil, err
		}

		fmt.Println(retData)
	}

	dataList := list.New()


	return dataList, nil
}























