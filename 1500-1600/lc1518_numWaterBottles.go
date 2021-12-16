package main

func numWaterBottles(numBottles int, numExchange int) int {
	var ret = numBottles

	for numBottles >= numExchange {
		var exchange = numBottles / numExchange
		ret += exchange
		numBottles = exchange + (numBottles % numExchange)
	}

	return ret
}
