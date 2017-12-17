package main

import "github.com/MsloveDl/HuobiProAPI/services"
import "fmt"

func main() {
	fmt.Println(services.GetTicker("bcxbtc"))
}
