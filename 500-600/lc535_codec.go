package main

import (
	"math/rand"
	"strconv"
	"strings"
	"unsafe"
)

type Codec struct {
	dataBase map[int]string
	id       int
}

func Constructor3() Codec {
	return Codec{map[int]string{}, 0}
}

func (c *Codec) encode(longUrl string) string {
	c.id++
	c.dataBase[c.id] = longUrl
	return "http://tinyurl.com/" + strconv.Itoa(c.id)
}

func (c *Codec) decode(shortUrl string) string {
	i := strings.LastIndexByte(shortUrl, '/')
	id, _ := strconv.Atoi(shortUrl[i+1:])
	return c.dataBase[id]
}

const KEY = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const LEN = 6
const PREFIX = "http://tinyurl.com/"

func RandUrl() string {
	var ret = make([]byte, LEN)
	for i := range ret {
		ret[i] = KEY[rand.Intn(len(KEY))]
	}
	return *(*string)(unsafe.Pointer(&ret))
}

type Codec2 struct {
	tinyToOrigin map[string]string
	originToTiny map[string]string
}

func Constructor2() Codec2 {
	return Codec2{
		tinyToOrigin: map[string]string{},
		originToTiny: map[string]string{},
	}
}

// Encodes a URL to a shortened URL.
func (c *Codec2) encode(longUrl string) string {
	var url string
	for {
		url = PREFIX + RandUrl()
		if c.originToTiny[url] == "" {
			break
		}
	}
	c.originToTiny[longUrl] = url
	c.tinyToOrigin[url] = longUrl
	return url
}

// Decodes a shortened URL to its original URL.
func (c *Codec2) decode(shortUrl string) string {
	return c.tinyToOrigin[shortUrl]
}
