package drawing

import (
	"image/color"
)

// Transparent ... Fully transparent.
func Transparent() color.NRGBA {
	return color.NRGBA{R: 0, G: 0, B: 0, A: 0}
}

// Black ... Returns the color (transparency is 0 to 255).
func Black(transparency uint8) color.NRGBA {
	return color.NRGBA{R: 0, G: 0, B: 0, A: transparency}
}

// Blue ... Returns the color (transparency is 0 to 255).
func Blue(transparency uint8) color.NRGBA {
	return color.NRGBA{R: 32, G: 0, B: 192, A: transparency}
}

// Red ... Returns the color (transparency is 0 to 255).
func Red(transparency uint8) color.NRGBA {
	return color.NRGBA{R: 176, G: 0, B: 0, A: transparency}
}

// Magenta ... Returns the color (transparency is 0 to 255).
func Magenta(transparency uint8) color.NRGBA {
	return color.NRGBA{R: 192, G: 0, B: 176, A: transparency}
}

// Green ... Returns the color (transparency is 0 to 255).
func Green(transparency uint8) color.NRGBA {
	return color.NRGBA{R: 0, G: 176, B: 16, A: transparency}
}

// Cyan ... Returns the color (transparency is 0 to 255).
func Cyan(transparency uint8) color.NRGBA {
	return color.NRGBA{R: 0, G: 176, B: 176, A: transparency}
}

// Yellow ... Returns the color (transparency is 0 to 255).
func Yellow(transparency uint8) color.NRGBA {
	return color.NRGBA{R: 176, G: 176, B: 0, A: transparency}
}

// Gray ... Returns the color (transparency is 0 to 255).
func Gray(transparency uint8) color.NRGBA {
	return color.NRGBA{R: 176, G: 176, B: 176, A: transparency}
}

// LightBlue ... Returns the color (transparency is 0 to 255).
func LightBlue(transparency uint8) color.NRGBA {
	return color.NRGBA{R: 64, G: 0, B: 255, A: transparency}
}

// LightRed ... Returns the color (transparency is 0 to 255).
func LightRed(transparency uint8) color.NRGBA {
	return color.NRGBA{R: 255, G: 0, B: 0, A: transparency}
}

// LightMagenta ... Returns the color (transparency is 0 to 255).
func LightMagenta(transparency uint8) color.NRGBA {
	return color.NRGBA{R: 255, G: 0, B: 255, A: transparency}
}

// LightGreen ... Returns the color (transparency is 0 to 255).
func LightGreen(transparency uint8) color.NRGBA {
	return color.NRGBA{R: 0, G: 255, B: 32, A: transparency}
}

// LightCyan ... Returns the color (transparency is 0 to 255).
func LightCyan(transparency uint8) color.NRGBA {
	return color.NRGBA{R: 0, G: 255, B: 255, A: transparency}
}

// LightYellow ... Returns the color (transparency is 0 to 255).
func LightYellow(transparency uint8) color.NRGBA {
	return color.NRGBA{R: 255, G: 255, B: 0, A: transparency}
}

// White ... Returns the color (transparency is 0 to 255).
func White(transparency uint8) color.NRGBA {
	return color.NRGBA{R: 255, G: 255, B: 255, A: transparency}
}
