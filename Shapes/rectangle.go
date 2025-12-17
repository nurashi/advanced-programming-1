package Shapes

type Rectangle struct {
	width, height float64
}

func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{width: width, height: height}
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}
