package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	dst := make([]byte, 32)
	md5Ctx := md5.New()
	md5Ctx.Write([]byte("恪慎克孝"))
	hex.Encode(dst, md5Ctx.Sum(nil))
	for i := 1; i < 100000000; i++ {
		md5Ctx.Reset()
		md5Ctx.Write(dst)
		hex.Encode(dst, md5Ctx.Sum(nil))
	}
	fmt.Println(string(dst))
}
