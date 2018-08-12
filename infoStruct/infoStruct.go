package infoStruct

import "time"

//coin metadata
/**
https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?start=1&limit=100&convert=USD
{
	"id": 1,
	"name": "Bitcoin",
	"symbol": "BTC",
	"slug": "bitcoin",
	"circulating_supply": 17204712,
	"total_supply": 17204712,
	"max_supply": 21000000,
	"date_added": "2013-04-28T00:00:00.000Z",
	"num_market_pairs": 331,
	"cmc_rank": 1,
	"last_updated": "2018-08-12T04:04:29.000Z",
	"quote": {
		"USD": {
			"price": 6316.28168389,
			"volume_24h": 4100193263.15743,
			"percent_change_1h": -0.00699175,
			"percent_change_24h": 2.67556,
			"percent_change_7d": -9.86776,
			"market_cap": 108669807282.20248,
			"last_updated": "2018-08-12T04:04:29.000Z"
		}
	}
}
 */
type STCoinMetadata struct {
	ID int32
	Name string
	Symbol string
	Slug string
	Total_supply int32
	Max_supply int32
	Date_added time.Time
	Cmc_rank int32
	Last_updated time.Time
}


//Get market quotes (historical)
/*
https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/historical

"quotes": [{
		"timestamp": "2018-06-22T19:29:37.000Z",
		"quote": {
			"USD": {
				"price": 6242.29,
				"volume_24h": 4681670000,
				"market_cap": 106800038746.48,
				"last_updated": "2018-06-22T20:15:17.651Z"
			}
		}
	},
	{
		"timestamp": "2018-06-22T19:34:33.000Z",
		"quote": {
			"USD": {
				"price": 6242.82,
				"volume_24h": 4682330000,
				"market_cap": 106809106575.84,
				"last_updated": "2018-06-22T20:15:17.651Z"
			}
		}
	}
]
*/
type STCoinHistoryDataItem struct {
	Price float64
	Volume_24h int64
	Market_cap float64
	Last_updated time.Time
	Timestamp time.Time
}





















