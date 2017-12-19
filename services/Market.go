package services

import (
	"encoding/json"
	"strconv"

	"github.com/MsloveDl/HuobiProAPI/config"
	"github.com/MsloveDl/HuobiProAPI/models"
	"github.com/MsloveDl/HuobiProAPI/untils"
)

//------------------------------------------------------------------------------------------
// 交易API

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

// 获取Market Detail 24小时成交量数据
// strSymbol: 交易对, btcusdt, bccbtc......
// return: MarketDetailReturn对象
// 未进行回测
func GetMarketDetail(strSymbol string) models.MarketDetailReturn {
	marketDetailReturn := models.MarketDetailReturn{}

	strParams := "symbol=" + strSymbol
	strUrl := config.MARKET_URL + "/detail"

	jsonMarketDetailReturn := untils.HttpGetRequest(strUrl, strParams)
	json.Unmarshal([]byte(jsonMarketDetailReturn), &marketDetailReturn)

	return marketDetailReturn
}

//------------------------------------------------------------------------------------------
// 公共API

// 查询系统支持的所有交易及精度
// return: SymbolsReturn对象
// 未回测
func GetSymbols() models.SymbolsReturn {
	symbolsReturn := models.SymbolsReturn{}

	strUrl := config.TRADE_URL + "/common/symbols"

	jsonSymbolsReturn := untils.HttpGetRequest(strUrl, "")
	json.Unmarshal([]byte(jsonSymbolsReturn), &symbolsReturn)

	return symbolsReturn
}

// 查询系统支持的所有币种
// return: CurrencysReturn对象
// 未回测
func GetCurrencys() models.CurrencysReturn {
	currencysReturn := models.CurrencysReturn{}

	strUrl := config.TRADE_URL + "/common/currencys"

	jsonCurrencysReturn := untils.HttpGetRequest(strUrl, "")
	json.Unmarshal([]byte(jsonCurrencysReturn), &currencysReturn)

	return currencysReturn
}

// 查询系统当前时间戳
// return: TimestampReturn对象
// 未回测
func GetTimestamp() models.TimestampReturn {
	timestampReturn := models.TimestampReturn{}

	strUrl := config.TRADE_URL + "/common/timestamp"

	jsonTimestampReturn := untils.HttpGetRequest(strUrl, "")
	json.Unmarshal([]byte(jsonTimestampReturn), &timestampReturn)

	return timestampReturn
}
