package main

import (
	"fmt"

	"github.com/MsloveDl/HuobiProAPI/untils"
)

func main() {
	chexiaoceshi()
}

func chexiaoceshi() {
	fmt.Println(untils.ApiKeyPost(make(map[string]string), "/v1/order/orders/405798779/submitcancel"))
}

func xiadanceshi() {
	mapParams := make(map[string]string)
	mapParams["account-id"] = "706355"
	mapParams["amount"] = "0.001"
	mapParams["price"] = "0.200000"
	mapParams["source"] = "api"
	mapParams["symbol"] = "bchbtc"
	mapParams["type"] = "buy-limit"

	fmt.Println(untils.ApiKeyPost(mapParams, "/v1/order/orders/place"))
}
