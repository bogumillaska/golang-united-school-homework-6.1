package golang_united_school_homework

import (
	"errors"
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
	if len(b.shapes) >= b.shapesCapacity {
		return errors.New("shapes capacity reached")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	print(len(b.shapes))
	if i > len(b.shapes)-1 {
		return nil, errors.New("index out of bounds")
	}
	shape := b.shapes[i]
	return shape, nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	shape, err := b.GetByIndex(i)

	if shape != nil {
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	}
	return shape, err
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	shapeToReplace, err := b.GetByIndex(i)

	if shapeToReplace != nil {
		b.shapes[i] = shape
	}
	return shapeToReplace, err

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
	var foundCircle bool = false
	for i := 0; i < len(b.shapes); i++ {
		if _, ok := b.shapes[i].(*Circle); ok {
			foundCircle = true
			copy(b.shapes[i:], b.shapes[i+1:])
			b.shapes = b.shapes[:len(b.shapes)-1]
			i = i - 1
		}
	}

	if !foundCircle {
		return errors.New("no circles found")
	}
	return nil
}
