package drawing

// Rect ... Defines a rectangular area.
type Rect struct {
	TopLeft, TopRight       Point
	BottomLeft, BottomRight Point
	Width, Height           int
}

// NewRect ... Returns a new rectangle.
func NewRect(x1, y1, x2, y2 int) Rect {
	tl := NewPoint(x1, y1)
	tr := NewPoint(x2, y1)
	bl := NewPoint(x1, y2)
	br := NewPoint(x2, y2)
	return Rect{
		TopLeft:     tl,
		TopRight:    tr,
		BottomLeft:  bl,
		BottomRight: br,
		Width:       x2 - x1 + 1,
		Height:      y2 - y1 + 1,
	}
}
