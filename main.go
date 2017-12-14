package drawing

import (
	"image"
	"image/color"
	"time"
)

// TimingLogFunc ... Logs timings. Sample LogDuration and NoDurationLog implementations are provided.
type TimingLogFunc func(string, time.Duration)

// NewSurface ... Returns a new drawing surface.
func NewSurface(width, height int, background color.NRGBA, logger TimingLogFunc) Surface {
	b := NewRect(0, 0, width-1, height-1)
	s := Surface{
		log:        logger,
		Bounds:     b,
		Image:      image.NewNRGBA(image.Rect(0, 0, width, height)),
		Width:      width,
		Height:     height,
		Background: background,
	}

	s.Clear()
	return s
}
