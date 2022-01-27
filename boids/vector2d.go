package main

import "math"

type Vector2D struct {
	x float64
	y float64
}

// Add adds vector values
func (v1 Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x + v2.x, y: v1.y + v2.y}
}

// Subtract subtracts vector values
func (v1 Vector2D) Subtract(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x - v2.x, y: v1.y - v2.y}
}

// Multiply multiplies vector values
func (v1 Vector2D) Multiply(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x * v2.x, y: v1.y * v2.y}
}

// AddV adds a value to a number
func (v1 Vector2D) AddV(d float64) Vector2D {
	return Vector2D{x: v1.x + d, y: v1.y + d}
}

// MultiplyV multiplies a value to a number
func (v1 Vector2D) MultiplyV(d float64) Vector2D {
	return Vector2D{x: v1.x * d, y: v1.y * d}
}

// DivideV divides a value from a number
func (v1 Vector2D) DivideV(d float64) Vector2D {
	return Vector2D{x: v1.x / d, y: v1.y / d}
}

func (v1 Vector2D) limit(lower, upper float64) Vector2D {
	return Vector2D{x: math.Min(math.Max(v1.x, lower), upper), y: math.Min(math.Max(v1.y, lower), upper)}
}

// Distance represents distance between to two points
func (v1 Vector2D) Distance(v2 Vector2D) float64 {
	return math.Sqrt(math.Pow(v1.x-v2.x, 2) + math.Pow(v1.y-v2.y, 2))
}
