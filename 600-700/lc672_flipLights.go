package main

var P = [4]int{
	0b11111111_11111111,
	0b10101010_10101010,
	0b01010101_01010101,
	0b10010010_01001001,
}

func flipLights(n int, presses int) int {
	if presses == 0 {
		return 1
	}
	if n == 1 {
		return 2
	}
	if presses == 1 {
		if n == 2 {
			return 3
		}
		return 4
	}
	if n == 2 {
		return 4
	}
	if presses == 2 {
		return 7
	}
	return 8

	// 一行一行试出来的...

	// var ret = make(map[int]bool)

	// var src int
	// queue := []int{src}

	// show := func(src []int) {
	//     var sb = make([]string, 0, len(src))
	//     for _, v := range src {
	//         sb = append(sb, fmt.Sprintf("%08b", v))
	//     }
	//     fmt.Println(strings.Join(sb, ","))
	// }

	// show(queue)

	// for len(queue) != 0 && presses > 0 {
	//     ln := len(queue)
	//     var count = make(map[int]bool)
	//     for i := 0; i < ln; i++ {
	//         for j, v := range P {
	//             s := queue[i]
	//             sv := v
	//             v ^= queue[i]
	//             if presses == 1 {
	//                 if ret[v] {
	//                     continue
	//                 }
	//                 ret[v] = true
	//                 fmt.Printf("%016b, %d[%016b], %016b\n",s, j, sv, v)
	//             } else {
	//                 if count[v] {
	//                     continue
	//                 }
	//             }
	//             count[v] = true
	//             queue = append(queue, v)
	//         }
	//     }
	//     queue = queue[ln:]
	//     presses--
	//     // show(queue)
	// }

	// return len(ret)
}
