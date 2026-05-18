package core

import "math"

type Vector2 struct {
	X, Y float64
}

func Add(v1, v2 Vector2) Vector2 {
	return Vector2{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

func Subtract(v1, v2 Vector2) Vector2 {
	return Vector2{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
}

func Dot(v1, v2 Vector2) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

func Cross(v1, v2 Vector2) float64 {
	return v1.X*v2.Y - v1.Y*v2.X
}

func Scale(v Vector2, factor float64) Vector2 {
	return Vector2{v.X * factor, v.Y * factor}
}

func Normalize(v Vector2) Vector2 {
	length := math.Sqrt(v.X*v.X + v.Y*v.Y)
	return Vector2{v.X / length, v.Y / length}
}
