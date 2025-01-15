package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_MEAT_PRICE float32 = 10

func main() {
	var chickenChannel = make(chan string)
	var meatChannel = make(chan string)
	var websites = []string{"amazon", "ebay", "flipkart", "snapdeal", "shopclues"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkMeatPrices(websites[i], meatChannel)
	}
	sendMessage(chickenChannel, meatChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var price = rand.Float32() * 20
		if price < MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func checkMeatPrices(website string, meatChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var price = rand.Float32() * 20
		if price < MAX_CHICKEN_PRICE {
			meatChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, meatChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Println("Chicken is available on", website)
	case website := <-meatChannel:
		fmt.Println("Meat is available on", website)
	}
}
