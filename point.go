package drawing

// Point ... An X/Y coordinate.
type Point struct {
	X, Y int
}

// NewPoint ... Returns a new point.
func NewPoint(x1, y1 int) Point {
	return Point{X: x1, Y: y1}
}
