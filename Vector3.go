package math3d

import (
	"math"
)

type Vector3 struct {
	X float64
	Y float64
	Z float64
}


func (v Vector3) Magnitude () (float64) {

	sqrt := math.Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))
	return sqrt

}

func (v Vector3) Normalise () {

	mag := v.Magnitude()

	v.X /= mag
	v.Y /= mag
	v.Z /= mag

}

func (v Vector3) Subtract (vsub *Vector3) (Vector3) {

	vnew := Vector3 {}

	vnew.X = v.X - vsub.X
	vnew.Y = v.Y - vsub.Y
	vnew.Z = v.Z - vsub.Z

	return vnew
}

func (v Vector3) Add (vadd *Vector3) (Vector3) {
	vnew := Vector3 {}

	vnew.X = v.X + vadd.X
	vnew.Y = v.Y + vadd.Y
	vnew.Z = v.Z + vadd.Z

	return vnew
}

func (v Vector3) Cross (vec Vector3) (Vector3) {
	vnew := Vector3{}

	vnew.X = (v.Y * vec.Z) - (v.Z * vec.Y)
	vnew.Y = (v.Z * vec.X) - (v.X * vec.Z)
	vnew.Z = (v.X * vec.Y) - (v.Y * vec.X)

	return vnew

}

func (this Vector3) Dot (vec Vector3) (float64) {

	return this.X * vec.X + this.Y * vec.Y + this.Z * vec.Z
}