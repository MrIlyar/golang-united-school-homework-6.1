package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (shape Circle) CalcPerimeter() float64 {
	return shape.Radius * 2.0 * math.Pi
}

func (shape Circle) CalcArea() float64 {
	return shape.Radius * shape.Radius * math.Pi
}
