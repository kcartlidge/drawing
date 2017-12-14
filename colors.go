package drawing

import (
	"image/color"
)

var (
	// Transparent ... Fully transparent.
	Transparent = color.RGBA{R: 0, G: 0, B: 0, A: 0}
	// Black ... Returns opaque Black as RGBA.
	Black = color.RGBA{R: 0, G: 0, B: 0, A: 255}
	// Blue ... Returns opaque Blue as RGBA.
	Blue = color.RGBA{R: 32, G: 0, B: 192, A: 255}
	// Red ... Returns opaque Red as RGBA.
	Red = color.RGBA{R: 176, G: 0, B: 0, A: 255}
	// Magenta ... Returns opaque Magenta as RGBA.
	Magenta = color.RGBA{R: 192, G: 0, B: 176, A: 255}
	// Green ... Returns opaque Green as RGBA.
	Green = color.RGBA{R: 0, G: 176, B: 16, A: 255}
	// Cyan ... Returns opaque Cyan as RGBA.
	Cyan = color.RGBA{R: 0, G: 176, B: 176, A: 255}
	// Yellow ... Returns opaque Yellow as RGBA.
	Yellow = color.RGBA{R: 176, G: 176, B: 0, A: 255}
	// Gray ... Returns opaque Gray as RGBA.
	Gray = color.RGBA{R: 176, G: 176, B: 176, A: 255}
	// LightBlue ... Returns opaque LightBlue as RGBA.
	LightBlue = color.RGBA{R: 64, G: 0, B: 255, A: 255}
	// LightRed ... Returns opaque LightRed as RGBA.
	LightRed = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	// LightMagenta ... Returns opaque LightMagenta as RGBA.
	LightMagenta = color.RGBA{R: 255, G: 0, B: 255, A: 255}
	// LightGreen ... Returns opaque LightGreen as RGBA.
	LightGreen = color.RGBA{R: 0, G: 255, B: 32, A: 255}
	// LightCyan ... Returns opaque LightCyan as RGBA.
	LightCyan = color.RGBA{R: 0, G: 255, B: 255, A: 255}
	// LightYellow ... Returns opaque LightYellow as RGBA.
	LightYellow = color.RGBA{R: 255, G: 255, B: 0, A: 255}
	// White ... Returns opaque White as RGBA.
	White = color.RGBA{R: 255, G: 255, B: 255, A: 255}
)
