package rocket

import "testing"

type testCase struct {
	a, tau, t, d, v, lorentz float64
}

var knownValues = []testCase{
	{G, 1 * Year, 1.19 * Year, 0.56 * LightYear, 0.77 * C, 1.58},
	{G, 2 * Year, 3.75 * Year, 2.90 * LightYear, 0.97 * C, 3.99},
	{G, 5 * Year, 83.7 * Year, 82.7 * LightYear, 0.99993 * C, 86.2},
	{G, 8 * Year, 1840 * Year, 1839 * LightYear, 0.9999998 * C, 1895},
	{G, 12 * Year, 113243 * Year, 113242 * LightYear, 0.99999999996 * C, 116641},
}

func approximately(x, y float64) bool {
	const tolerance = 0.01
	return x > y*(1-tolerance) && x < y*(1+tolerance)
}

// Tests that the effects of proper acceleration match up to the hyperbolic solution
// over the same (coordinate) time span.
func TestProperAcceleration(t *testing.T) {
	for _, tc := range knownValues {
		w := tc.a * tc.t
		if !approximately(CoordinateVelocity(w), tc.v) {
			t.Errorf(
				"w(%f G, %f y) = %f c (proper), %f c (coordinate); wanted %f c (proper), %f c (coordinate)",
				tc.a/G, tc.tau/Year, w/C, CoordinateVelocity(w)/C, ProperVelocity(tc.v)/C, tc.v/C)
		}
	}
}

func TestVelocity(t *testing.T) {
	for _, tc := range knownValues {
		v := Velocity(tc.a, tc.t)
		if !approximately(v, tc.v) {
			t.Errorf(
				"Velocity(%f g, %f y) = %f c (wanted %f c)",
				tc.a/G, tc.t/Year, v/C, tc.v/C)
		}
	}
}

func TestLorentzFactor(t *testing.T) {
	for _, tc := range knownValues {
		lorentz := LorentzFactor(tc.a, tc.t)
		if !approximately(lorentz, tc.lorentz) {
			t.Errorf(
				"LorentzFactor(%f g, %f y) = %f (wanted %f)",
				tc.a/G, tc.t/Year, lorentz, tc.lorentz)
		}
	}
}

func TestProperVelocity(t *testing.T) {
	for _, tc := range knownValues {
		w := ProperVelocity(tc.v)

		if v := CoordinateVelocity(w); !approximately(v, tc.v) {
			t.Errorf(
				"CoordinateVelocity(ProperVelocity(%f c)) = %f c (should be the same)",
				tc.v/C, v/C)
		}
	}
}
