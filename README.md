Package rocket provides functions for modeling rockets at relativistic speeds.

# Scope

This package models an idealized rocket under special relativity. It is NOT complete,
accurate or suited for any serious applications. It provides both algebraic and
numerical solutions to some specific problems, like calculating travel time under
constant acceleration.

# Limitations

All algorithms in this package are limited by 64-bit floating point precision.
Numerical solutions appreciably deviate from algebraic predictions after a few
thousand iterations.

General relativity is not modeled. Numerical solutions over-estimate the effect of
time dilation, because they compute the lorentz factor based on the velocity after
applying acceleration in each step.

# Notation, units and nomenclature

We refer to two reference frames: proper and coordinate. The proper reference frame
is accelerated and local to the rocket. The coordinate reference frame is the
stationary observer.

All units are SI, and usually omitted for the sake of brevity. (We define some common
constants for light years, etc.)

The notation used throughout the package is standard, but adapted for ASCII.

Quantities:

```
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
```

# References

Classical rocket equation: https://en.wikipedia.org/wiki/Tsiolkovsky_rocket_equation

Relativistic rocket: https://en.wikipedia.org/wiki/Relativistic_rocket

The Relativistic Rocket by Phillip Gibbs:
http://math.ucr.edu/home/baez/physics/Relativity/SR/Rocket/rocket.html
