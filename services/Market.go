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

// 获取交易深度信息
// strSymbol: 交易对, btcusdt, bccbtc......
// strType: Depth类型, step0、step1......stpe5 (合并深度0-5, 0时不合并)
// return: MarketDepthReturn对象
func GetMarketDepth(strSymbol, strType string) models.MarketDepthReturn {
	marketDepthReturn := models.MarketDepthReturn{}

	strParams := "symbol=" + strSymbol + "&type=" + strType
	strUrl := config.MARKET_URL + "/depth"

	jsonMarketDepthReturn := untils.HttpGetRequest(strUrl, strParams)
	json.Unmarshal([]byte(jsonMarketDepthReturn), &marketDepthReturn)

	return marketDepthReturn
}

// 获取交易细节信息
// strSymbol: 交易对, btcusdt, bccbtc......
// return: TradeDetailReturn对象
func GetTradeDetail(strSymbol string) models.TradeDetailReturn {
	tradeDetailReturn := models.TradeDetailReturn{}

	strParams := "symbol=" + strSymbol
	strUrl := config.MARKET_URL + "/trade"

	jsonTradeDetailReturn := untils.HttpGetRequest(strUrl, strParams)
	json.Unmarshal([]byte(jsonTradeDetailReturn), &tradeDetailReturn)

	return tradeDetailReturn
}

// 批量获取最近的交易记录
// strSymbol: 交易对, btcusdt, bccbtc......
// nSize: 获取交易记录的数量, 范围1-2000
// return: TradeReturn对象
func GetTrade(strSymbol string, nSize int) models.TradeReturn {
	tradeReturn := models.TradeReturn{}

	strParams := "symbol=" + strSymbol + "&size=" + strconv.Itoa(nSize)
	strUrl := config.MARKET_URL + "/history/trade"

	jsonTradeReturn := untils.HttpGetRequest(strUrl, strParams)
	json.Unmarshal([]byte(jsonTradeReturn), &tradeReturn)

	return tradeReturn
}
