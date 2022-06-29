package main

const mod int = 1e9 + 7

func numPrimeArrangements(n int) int {
	numPrimes := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			numPrimes++
		}
	}
	return factorial(numPrimes) * factorial(n-numPrimes) % mod
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func factorial(n int) int {
	f := 1
	for i := 1; i <= n; i++ {
		f = f * i % mod
	}
	return f
}
