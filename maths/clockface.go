package clockface

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

// A Point is a Cartesian coordinate. They are used in the package
// to represent the unit vector from the origin of a clock hand.
type Point struct {
	X float64
	Y float64
}

// SecondsInRadians returns the angle of the second hand of a clock
func SecondsInRadians(t time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / (float64(t.Second()))))
}

// SecondHandPoint returns the unit vector of the second hand of a clock
func SecondHandPoint(t time.Time) Point {
	return angleToPoint(SecondsInRadians(t))
}

// MinutesInRadians returns the angle of the minute hand of a clock
func MinutesInRadians(t time.Time) float64 {
	return (SecondsInRadians(t) / minutesInClock) + (math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

// MinuteHandPoint returns the unit vector of the minute hand of a clock
func MinuteHandPoint(t time.Time) Point {
	return angleToPoint(MinutesInRadians(t))
}

// HoursInRadians returns the angle of the hour hand of a clock
func HoursInRadians(t time.Time) float64 {
	return (MinutesInRadians(t) / hoursInClock) + (math.Pi / (hoursInHalfClock / float64(t.Hour()%12)))
}

// HourHandPoint returns the unit vector of the hour hand of a clock
func HourHandPoint(t time.Time) Point {
	return angleToPoint(HoursInRadians(t))
}

func angleToPoint(angle float64) Point {
	// because SOH, CAH, TOA
	// and assuming and H = 1
	// x which is the opposite of the angle of a clock; O = S(a) / H(1)
	// y which is the adjacent of the angle of a clock; A = C(a) / H(1)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
