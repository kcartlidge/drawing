package drawing

// Polygon ... Defines an auto-closing shape.
type Polygon []Point

// NewPolygon ... Returns a new polygon with the given points.
func NewPolygon(points ...Point) Polygon {
	return Polygon(points)
}
