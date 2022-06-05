//go:build ignore

package main

import (
	"math"
	"math/rand"
	"time"
)

type Solution struct {
	radius, x, y float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	rand.Seed(time.Now().UnixNano())
	return Solution{
		radius: radius,
		x:      x_center,
		y:      y_center,
	}
}

func (s *Solution) RandPoint() []float64 {
	var ret = make([]float64, 2)
	// 随机半径
	var d = math.Sqrt(rand.Float64()) * s.radius
	// 随机弧度对应的正余弦
	var sin, cos = math.Sincos(rand.Float64() * 2 * math.Pi)
	ret[0], ret[1] = d*sin+s.x, d*cos+s.y
	return ret
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
