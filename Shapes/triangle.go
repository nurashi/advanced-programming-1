package Shapes

import "math"

type Triangle struct {
	Side1 float64
	Side2 float64
	Side3 float64
}

func NewTriangle(side1, side2, side3 float64) *Triangle {
	return &Triangle{side1, side2, side3}
}

// I didn't write validation for other shapes just because they are can exist anyway, but in case with triangle we have absolute condition for existance of traingle
func (t *Triangle) Validate() bool {
	// 1. Check if any side is non-positive, NaN, or Infinite
	if t.Side1 <= 0 || t.Side2 <= 0 || t.Side3 <= 0 || math.IsNaN(t.Side1) || math.IsNaN(t.Side2) || math.IsNaN(t.Side3) || math.IsInf(t.Side1, 0) || math.IsInf(t.Side2, 0) || math.IsInf(t.Side3, 0) {
		return false
	}

	// 2. Check the Triangle Inequality Theorem
	if t.Side1+t.Side2 <= t.Side3 || t.Side1+t.Side3 <= t.Side2 || t.Side2+t.Side3 <= t.Side1 {
		return false
	}

	// If all conditions are met, it is Side1 valid triangle
	return true
}

func (t *Triangle) Area() float64 {
	if t.Validate() {
		return 0
	}
	semi := (t.Side1 + t.Side2 + t.Side3) / 2 // semi-perimeter
	return math.Sqrt(semi * (semi - t.Side1) * (semi - t.Side2) * (semi - t.Side3))
}

func (t *Triangle) Perimeter() float64 {
	if !t.Validate() {
		return 0
	}
	return t.Side1 + t.Side2 + t.Side3
}
