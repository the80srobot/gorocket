package rocket

import (
	"math"
)

const (
	LightYear = 9460730472580800
	C         = 299792458
	Year      = float64(3600 * 24 * 365)
	G         = 9.8
)

func CoordinateTime(d, a float64) float64 {
	q := d / C
	return math.Sqrt(q*q + 2*d/a)
}

func Velocity(a, t float64) float64 {
	at := (a * t)
	return at / math.Sqrt(1+(at/C)*(at/C))
}

func VelocityWithV0(a, t, v0 float64) float64 {
	if v0 == 0 {
		return Velocity(a, t)
	}

	t0 := CoordinateTimeToReachVelocity(a, v0)
	return Velocity(a, t0+t)
}

func CoordinateTimeToReachVelocity(a, v float64) float64 {
	return (C * v) / (a * math.Sqrt((C*C)-(v*v)))
}

func ProperVelocity(v float64) float64 {
	r := v / C
	return v / math.Sqrt(1-(r*r))
}

func CoordinateVelocity(w float64) float64 {
	return (C * w) / math.Sqrt(C*C+w*w)
}

func ProperTime(d, a float64) float64 {
	return (C / a) * math.Acosh((a*d)/(C*C)+1)
}

func LorentzFactor(a, t float64) float64 {
	x := (a * t) / C
	return math.Sqrt(1 + x*x)
}
