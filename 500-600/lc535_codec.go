package main

import (
	"strconv"
	"strings"
)

type Codec struct {
	dataBase map[int]string
	id       int
}

func Constructor() Codec {
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
