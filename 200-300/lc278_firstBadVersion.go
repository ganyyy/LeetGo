package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(int) bool {
	return true
}

func firstBadVersion(n int) int {
	left, right := 1, n
	for left <= right {
		// 核心是这个: 为了防止溢出, 通过left+(right-left)/2 求中值. 等同于
		// (right+left)/2
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {

}
