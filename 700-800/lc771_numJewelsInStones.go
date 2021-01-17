package main

func numJewelsInStones(J string, S string) int {
	src := [52]byte{}
	for _, v := range J {
		src[getPos(byte(v))] = 1
	}
	var cnt int
	for i := 0; i < len(S); i++ {
		if src[getPos(S[i])] != 0 {
			cnt++
		}
	}

	return cnt
}

func getPos(b byte) byte {
	if b > 'Z' {
		return b - 'a' + 26
	}
	return b - 'A'
}

func main() {

}
