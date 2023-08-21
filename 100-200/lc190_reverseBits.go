package main

func reverseBits(num uint32) uint32 {
	var ret uint32

	// for i := 0; i < 32; i++ {
	//     if num & (1<<i) != 0 {
	//         ret |= 1 <<(31-i)
	//     }
	// }

	// for i := 0; i < 32; i++ {
	//     ret <<= 1
	//     ret |= (num&1)
	//     num >>= 1
	// }

	for i := 0; i < 32; i++ {
		ret = ret<<1 | num&1
		num >>= 1
	}

	return ret
}
