package rocket

import "math"

type Vector3 struct {
	x, y, z float64
}

func (v Vector3) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v Vector3) Normalized() Vector3 {
	m := v.Magnitude()
	if m <= 9.99999974737875e-06 {
		return Vector3{}
	}

	return Vector3{v.x / m, v.y / m, v.z / m}
}

func (v Vector3) Add(w Vector3) Vector3 {
	return Vector3{v.x + w.x, v.y + w.y, v.z + w.z}
}

func (v Vector3) MultiplyByScalar(s float64) Vector3 {
	return Vector3{v.x * s, v.y * s, v.z * s}
}
