package core

import "math"

type Vector2 struct {
	X, Y float64
}

func VectorAdd(v1, v2 Vector2) Vector2 {
	return Vector2{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

func VectorSubtract(v1, v2 Vector2) Vector2 {
	return Vector2{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
}

func VectorDot(v1, v2 Vector2) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

func VectorCross(v1, v2 Vector2) float64 {
	return v1.X*v2.Y - v1.Y*v2.X
}

func VectorScale(v Vector2, factor float64) Vector2 {
	return Vector2{v.X * factor, v.Y * factor}
}

func VectorNormalize(v Vector2) Vector2 {
	length := math.Sqrt(v.X*v.X + v.Y*v.Y)
	return Vector2{v.X / length, v.Y / length}
}

func VectorZero() Vector2 {
	return Vector2{}
}
