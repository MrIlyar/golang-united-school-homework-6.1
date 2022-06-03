package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (shape Rectangle) CalcPerimeter() float64 {
	return shape.Height*2.0 + shape.Weight*2.0
}

func (shape Rectangle) CalcArea() float64 {
	return shape.Height * shape.Weight
}
