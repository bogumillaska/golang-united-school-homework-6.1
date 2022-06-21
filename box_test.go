package golang_united_school_homework

import (
	"testing"
)

func TestAddShape(t *testing.T) {
	box := NewBox(2)
	circle := Circle{Radius: 3}
	circle2 := Circle{Radius: 3}
	circle3 := Circle{Radius: 3}
	err := box.AddShape(circle)
	if err != nil {
		t.Error("err should not produce errors")
	}
	err2 := box.AddShape(circle2)
	if err2 != nil {
		t.Error("err2 should not produce errors")
	}
	err3 := box.AddShape(circle3)
	if err3 == nil {
		t.Error("err3 should produce errors")
	}

}

func TestGetShape(t *testing.T) {
	box := NewBox(2)
	circle := Circle{Radius: 3}
	circle2 := Circle{Radius: 4}
	box.AddShape(circle)
	box.AddShape(circle2)

	shape, err := box.GetByIndex(0)
	if shape != circle {
		t.Errorf("wanted %v, got %v", circle, shape)
	}
	if err != nil {
		t.Error("Should not produce errors")
	}

	shape2, err2 := box.GetByIndex(1)
	if shape2 != circle2 {
		t.Errorf("wanted %v, got %v", circle, shape)
	}
	if err2 != nil {
		t.Error("Should not produce errors")
	}

	shape3, err3 := box.GetByIndex(2)
	if shape3 != nil {
		t.Errorf("shape3 should be nil")
	}
	if err3 == nil {
		t.Error("Should produce errors")
	}
}

func Test_RemoveAllCirlces(t *testing.T) {
	box := NewBox(10)
	rect := Rectangle{Height: 1, Weight: 3}
	tr := Triangle{Side: 3}
	circle := &Circle{Radius: 3}
	circle2 := &Circle{Radius: 4}

	box.AddShape(circle)
	box.AddShape(rect)
	box.AddShape(&Circle{Radius: 1})
	box.AddShape(tr)
	box.AddShape(&Circle{Radius: 2})
	box.AddShape(circle2)

	err := box.RemoveAllCircles()
	if err != nil {
		t.Error("Should find cirlces")
	}

	if len(box.shapes) != 2 {
		t.Errorf("should have %v elements, got %v", 2, len(box.shapes))
	}

}
