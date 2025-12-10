package utils

import "math"

type Vec3 struct {
	X int
	Y int
	Z int
}

func (v Vec3) Add(other Vec3) Vec3 {
	return Vec3{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

func (v Vec3) Sub(other Vec3) Vec3 {
	return Vec3{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}

func (v Vec3) Distance(other Vec3) float64 {
	d := other.Sub(v)
	return math.Sqrt(float64(d.X*d.X + d.Y*d.Y + d.Z*d.Z))
}

func (v Vec3) ManhattanDistance(other Vec3) int {
	d := other.Sub(v)
	return Abs(d.X) + Abs(d.Y) + Abs(d.Z)
}
