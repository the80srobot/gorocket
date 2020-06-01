/*
Package rocket provides functions for modeling rockets at relativistic speeds.

Scope

This package models an idealized rocket under special relativity. It is NOT complete,
accurate or suited for any serious applications. It provides both algebraic and
numerical solutions to some specific problems, like calculating travel time under
constant acceleration.

Limitations

All algorithms in this package are limited by 64-bit floating point precision.
Numerical solutions appreciably deviate from algebraic predictions after a few
thousand iterations.

General relativity is not modeled. Numerical solutions over-estimate the effect of
time dilation, because they compute the lorentz factor based on the velocity after
applying acceleration in each step.

Notation, units and nomenclature

We refer to two reference frames: proper and coordinate. The proper reference frame
is accelerated and local to the rocket. The coordinate reference frame is the
stationary observer.

All units are SI, and usually omitted for the sake of brevity. (We define some common
constants for light years, etc.)

The notation used throughout the package is standard, but adapted for ASCII.

Quantities:

 a  // acceleration, proper frame (both scalar and vector)
 v  // coordinate velocity (both scalar and vector)
 w  // proper velocity (both scalar and vector)
 t  // coordinate (observer) time
 tau  // proper (shipboard) time
 d  // distance
 lorentz  // the Lorentz factor (gamma)
 dt  // delta time

Other:

 C  // the speed of light
 G  // 9.8
 Year  // 365 * 24 * 3600 seconds
 Velocity  // Used to refer to both velocity and speed

References

Classical rocket equation: https://en.wikipedia.org/wiki/Tsiolkovsky_rocket_equation

Relativistic rocket: https://en.wikipedia.org/wiki/Relativistic_rocket

The Relativistic Rocket by Phillip Gibbs:
http://math.ucr.edu/home/baez/physics/Relativity/SR/Rocket/rocket.html

*/
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

// Rocket is a numerical model for a rocket at relativistic speeds.
//
// Each physics step, call Accelerate, then read out the updated parameters.
// See Accelerate for notes about precision (loss thereof) and other notes.
//
// To simulate time's passage on a physics rame with no acceleration, call
// Accelerate with a zero Vector3.
type Rocket struct {
	// Proper velocity (can exceed C).
	W Vector3
	// Coordinate (Earth/observer) and proper (shipboard) time.
	T, Tau float64
}

// Accelerate the rocket by applying constant proper acceleration for dt seconds.
//
// Acceleration is in the proper frame (as felt by people in the rocket), and
// dt is in coordinate time. This may seem counterintuitive - proper
// acceleration is related to coordinate acceleration by the cube of the lorentz
// factor. See also AccelerateOnProperTime.
//
// This is the main physics step function for the rocket. It updates the time
// and velocity. Accelerate over-estimates time dilation, because it computes
// the Lorentz factor after applying acceleration. Call it with smaller dt
// to make the effect neglibible.
func (r *Rocket) Accelerate(a Vector3, dt float64) {
	r.W.x += a.x * dt
	r.W.y += a.y * dt
	r.W.z += a.z * dt
	r.T += dt
	r.Tau += dt / r.LorentzFactor()
}

// AccelerateOnProperTime is an alternative step function to Accelerate.
//
// The change in proper velocity of an accelerating observer is the integral of
// the proper acceleration (as experienced by the observer) over the coordinate
// time. This may seem counter intuitive. Among other things, it means that the
// change in proper velocity depends on the current velocity, everything else
// being constant.
//
// This function can be helpful for simulating physics using the time
// experienced by an observer inside the accelerating (proper) frame.
func (r *Rocket) AccelerateOnProperTime(a Vector3, dtau float64) {
	dt := r.LorentzFactor() * dtau
	r.W.x += a.x * dt
	r.W.y += a.y * dt
	r.W.z += a.z * dt
	r.T += dt
	r.Tau += dtau
}

func (r *Rocket) LorentzFactor() float64 {
	v := r.V()
	return 1 / math.Sqrt(1-(v*v)/(C*C))
}

func (r *Rocket) V() float64 {
	return CoordinateVelocity(r.W.Magnitude())
}

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
