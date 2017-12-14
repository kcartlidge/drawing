package drawing

import (
	"log"
	"math"
	"time"
)

// TrackDuration ... Deferred execution logs duration of the current scope.
func TrackDuration(name string, started time.Time) {
	sofar := time.Since(started)
	log.Printf("- %s - %s", name, sofar)
}

func intAbs(v int) int {
	if v < 0 {
		return 0 - v
	}
	return v
}

func swapFloats(a, b float64) (float64, float64) {
	return b, a
}

func round(x float64) float64 {
	return math.Floor(x + 0.5)
}

func mantissa(x float64) float64 {
	return x - math.Floor(x)
}

func inverseMantissa(x float64) float64 {
	return 1 - mantissa(x)
}
