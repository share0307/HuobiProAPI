package services

import (
	"encoding/json"
	"strconv"

	"github.com/MsloveDl/HuobiProAPI/config"
	"github.com/MsloveDl/HuobiProAPI/models"
	"github.com/MsloveDl/HuobiProAPI/untils"
)

// 获取K线数据
// strSymbol: 交易对, btcusdt, bccbtc......
// strPeriod: K线类型, 1min, 5min, 15min......
// nSize: 获取数量, [1-2000]
// return: KLineReturn 对象
func GetKLine(strSymbol, strPeriod string, nSize int) models.KLineReturn {
	kLineReturn := models.KLineReturn{}

	strParams := "symbol=" + strSymbol + "&period=" + strPeriod +
		"&size=" + strconv.Itoa(nSize)
	strUrl := config.MARKET_URL + "/history/kline"

	jsonKLineReturn := untils.HttpGetRequest(strUrl, strParams)
	json.Unmarshal([]byte(jsonKLineReturn), &kLineReturn)

	return kLineReturn
}

// 获取聚合行情
// strSymbol: 交易对, btcusdt, bccbtc......
// return: TickReturn对象
func GetTicker(strSymbol string) models.TickerReturn {
	tickerReturn := models.TickerReturn{}

	strParams := "symbol=" + strSymbol
	strUrl := config.MARKET_URL + "/detail/merged"

	jsonTickReturn := untils.HttpGetRequest(strUrl, strParams)
	json.Unmarshal([]byte(jsonTickReturn), &tickerReturn)

	return tickerReturn
}
