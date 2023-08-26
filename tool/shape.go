package tool

type Shape struct{
	ShapeName string
}

func (s Shape) ShapeDisplay() string{
	return s.ShapeName
}