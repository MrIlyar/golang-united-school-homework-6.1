package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (shape Triangle) CalcPerimeter() float64 {
	return shape.Side * 3.0
}

func (shape Triangle) CalcArea() float64 {
	return math.Sqrt(3) / 4 * shape.Side * shape.Side
}
