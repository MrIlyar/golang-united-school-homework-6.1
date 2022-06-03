package golang_united_school_homework

import "errors"

var (
	errorOutOfShapesCapacity = errors.New("out of the shapesCapacity range")
	errorOutOfRange          = errors.New("index doesn't exist or index went out of the range")
	errorNoCircles           = errors.New("circles are not exist in the list")
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == b.shapesCapacity {
		return errorOutOfShapesCapacity
	}

	b.shapes = append(b.shapes, shape)

	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) == 0 || i < 0 || i >= len(b.shapes) {
		return nil, errorOutOfRange
	}

	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if len(b.shapes) == 0 || i < 0 || i >= len(b.shapes) {
		return nil, errorOutOfRange
	}

	var result = b.shapes[i]

	copy(b.shapes[i:], b.shapes[i+1:])
	b.shapes = b.shapes[:len(b.shapes)-1]

	return result, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if len(b.shapes) == 0 || i < 0 || i >= len(b.shapes) {
		return nil, errorOutOfRange
	}

	var result = b.shapes[i]
	b.shapes[i] = shape

	return result, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var result float64

	for _, shape := range b.shapes {
		result += shape.CalcPerimeter()
	}

	return result
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var result float64

	for _, shape := range b.shapes {
		result += shape.CalcArea()
	}

	return result
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var found bool = false
	var notCircles = make([]Shape, 0)

	for _, shape := range b.shapes {
		_, okPointer := shape.(*Circle)
		_, okValue := shape.(Circle)

		if !(okPointer || okValue) {
			notCircles = append(notCircles, shape)
		} else {
			found = true
		}
	}

	b.shapes = notCircles

	if !found {
		return errorNoCircles
	} else {
		return nil
	}
}
