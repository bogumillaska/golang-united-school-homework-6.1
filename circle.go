package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (cir Circle) CalcPerimeter() float64 {
	return 2 * math.Pi * cir.Radius
}

func (cir Circle) CalcArea() float64 {
	return math.Pi * math.Pow(cir.Radius, 2)
}
