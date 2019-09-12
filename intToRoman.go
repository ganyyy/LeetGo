package main

import "fmt"

func intToRoman(num int) string {
	k := []int{
		1, 4, 5, 9,
		10, 40, 50, 90,
		100, 400, 500, 900,
		1000,
	}
	v := [][]byte{
		[]byte("I"), []byte("IV"), []byte("V"), []byte("IX"),
		[]byte("X"), []byte("XL"), []byte("L"), []byte("XC"),
		[]byte("C"), []byte("CD"), []byte("D"), []byte("CM"),
		[]byte("M"),
	}
	res := make([]byte, 0)

	i := len(k) - 1
	for num != 0 {
		tK := k[i]
		if num >= tK {
			d := num / tK
			tV := v[i]
			for j := 0; j < d; j++ {
				res = append(res, tV...)
			}
			num = num % tK
		} else {
			i--
		}
	}
	return string(res)
}

func main() {
	fmt.Println(intToRoman(1994))
}
