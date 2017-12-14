package drawing

import (
	"image"
	"time"
)

// TimingLogFunc ... Logs timings. Sample LogDuration and NoDurationLog implementations are provided.
type TimingLogFunc func(string, time.Duration)

// NewSurface ... Returns a new drawing surface.
func NewSurface(width, height int, logger TimingLogFunc) Surface {
	b := NewRect(0, 0, width-1, height-1)
	s := Surface{
		log:        logger,
		Bounds:     b,
		Image:      image.NewRGBA(image.Rect(0, 0, width, height)),
		Width:      width,
		Height:     height,
		Background: Black,
	}

	s.Clear(s.Background)
	return s
}
