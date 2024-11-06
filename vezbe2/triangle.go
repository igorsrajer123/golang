package main

import (
	"math"
)

type Triangle struct {
	a float64
	b float64
	c float64
}

func calculateTriangleSurface(triangle Triangle) float64 {
	triangle.a = 6
	s := (triangle.a + triangle.b + triangle.c) / 2
	return math.Sqrt(s * (s - triangle.a) * (s - triangle.b) * (s - triangle.c))
}
