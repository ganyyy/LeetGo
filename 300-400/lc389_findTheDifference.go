package main

func findTheDifference(s string, t string) byte {
	// var set [26]int

	// for i := 0; i < len(s); i++ {
	//     set[s[i]-'a']++
	// }

	// var idx int
	// for i := 0; i < len(t); i++ {
	//     idx = int(t[i]-'a')
	//     set[idx]--
	//     if set[idx] < 0 {
	//         return t[i]
	//     }
	// }

	// var sum int64
	// for i := 0; i < len(s); i++ {
	//     sum += int64(t[i]-s[i])
	// }
	// return byte(sum) + t[len(t)-1]

	var diff byte
	for i := range s {
		diff ^= s[i] ^ t[i]
	}
	return diff ^ t[len(t)-1]
}
