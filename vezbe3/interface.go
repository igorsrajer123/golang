package main

type Shape interface {
	calculateSurface() float64
	calculateLength() float64
}

type Triangle struct {
	a float64
}

type Square struct {
	a float64
}

func (triangle Triangle) calculateSurface() float64 {
	return triangle.a * triangle.a * triangle.a
}

func (square Square) calculateSurface() float64 {
	return square.a * square.a
}
