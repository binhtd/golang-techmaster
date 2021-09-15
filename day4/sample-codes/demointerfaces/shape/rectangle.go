package shape

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (rec Rectangle) Area() float64 {
	return rec.Width * rec.Height
}

func (rec Rectangle) Perimeter() float64 {
	return 2 * (rec.Height + rec.Width)
}

func (rec Rectangle) String() string {
	return fmt.Sprintf("Rectangle w=%.2f, h=%.2f", rec.Width, rec.Height)
}
