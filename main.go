package main

import (
	"fmt"

	"github.com/MsloveDl/HuobiProAPI/services"
)

func main() {
	fmt.Println(services.GetTrade("bcxbtc", 100))
}
