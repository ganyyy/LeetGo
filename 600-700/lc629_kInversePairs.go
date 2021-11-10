package main

func kInversePairs(n, k int) int {
	const mod int = 1e9 + 7
	// f[i][j] 表示长度为i的数组, 恰好包含j个逆序对的方案数量
	// 现在添加一个新的数字进行状态转移
	// 假设第i个元素选取的数字k
	// 将整个数组分为两部分:
	// 1...k-1, k+1...i
	// 那么整个数组的逆序对数为 (i-k) + 逆序(i-1)
	// 要使得总共有j个逆序对, 逆序(i-1)这部分的逆序对数量应该是(j-(i-k))
	// 逆序对的数量仅和元素的相对大小有关, 所以
	// 1...k-1这部分不变, k+1...i这部分整体-1变成 1...i-1
	// 等同于sum(f[i-1][j-(i-k)]), k∈[1,i] = sum(f[i-1][j-k]), k∈[0,i-1]
	//          为啥呢? 当k=i时, j-i+k = j, 对应右边j-k, k=0的情况
	//                  当k=i-1时, j-i+i-1 = j-1, 对应右边j-k, k=1的情况. 依此类推
	// f[i][j-1]和f[i][j]之间的关系为
	// f[i][j] = f[i][j-1]-f[i-1][j-i]+f[i-1][j](带入上边的计算公式)
	// 此时, f[i]只和f[i-1]相关, 所以可以进行压缩
	f := [2][]int{make([]int, k+1), make([]int, k+1)}
	f[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= k; j++ {
			cur := i & 1
			prev := cur ^ 1
			f[cur][j] = 0
			if j > 0 {
				f[cur][j] = f[cur][j-1]
			}
			if j >= i {
				f[cur][j] -= f[prev][j-i]
			}
			f[cur][j] += f[prev][j]
			if f[cur][j] >= mod {
				f[cur][j] -= mod
			} else if f[cur][j] < 0 {
				f[cur][j] += mod
			}
		}
	}
	return f[n&1][k]
}
