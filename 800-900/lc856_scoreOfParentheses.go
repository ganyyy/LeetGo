package main

// func scoreOfParentheses(s string) int {

//     const (
//         L = '('
//         R = ')'
//     )

//     var calc func(string) int

//     calc = func(s string) int {
//         if len(s) == 0 {
//             return 0
//         }
//         if s[0] != L {
//             return 0 // ?
//         }
//         var cnt int
//         var ret int
//         var start int
//         for i := 0; i < len(s); i++ {
//             if s[i] == L {
//                 cnt++
//             } else {
//                 cnt--
//             }
//             if cnt != 0 {
//                 continue
//             }
//             if i-start == 1 {
//                 // ()
//                 ret += 1
//             } else {
//                 ret += 2 * calc(s[start+1:i])
//             }
//             start = i+1
//         }
//         return ret
//     }

//     return calc(s)
// }

func scoreOfParentheses(s string) (ans int) {
	bal := 0 // 表示层级, 每一个 ( 都会加深一层
	for i, c := range s {
		if c == '(' {
			bal++
		} else {
			bal--
			if s[i-1] == '(' {
				ans += 1 << bal
			}
		}
	}
	return
}
