package main

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	x := 1 % k
	for i := 1; ; i++ { // 一定有解
		// 这里的i是位数!
		if x == 0 {
			return i
		}
		/*
		   前置知识点:
		       (a+b)%m = ((a%m)+(b%m))%m
		       (a*b)%m = ((a%m)*(b%m))%m
		*/
		// n = 1,  x = n % k
		// n = 11, x = n % k
		// n = 111,x = n % k
		// ...
		// n_new = (n_old * 10 + 1)
		// x_new = n_new % k
		//       = (n_old * 10 + 1) % k
		//       = ((n_old % k) * 10 + 1) % k  带入公式1
		//       = (x_old * 10 + 1) % k
		x = (x*10 + 1) % k
	}
}
