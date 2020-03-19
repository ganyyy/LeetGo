package main

import "fmt"

func longestPalindrome(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := [58]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'A']++
	}
	res, n := 0, 0
	for i := 0; i < len(count); i++ {
		res += count[i]
		if (count[i] & 1) == 1 {
			n++
		}
	}
	if n != 0 {
		return res - n + 1
	} else {
		return res
	}
}

func main() {
	res := "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"
	fmt.Println(longestPalindrome(res))
}
