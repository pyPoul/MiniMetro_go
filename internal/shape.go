package internal

import (
	"fmt"
	"math/rand"
)

type Shape int

const (
	Circle Shape = iota
	Triangle
	Square
	Star
	Pentagon
	Gem
	Cross
	Wedge
	Diamond
	Oval
)

func (s *Shape) ToString() (string, error) {
	switch *s {
	case Circle:
		return "circle", nil
	case Triangle:
		return "triangle", nil
	case Square:
		return "square", nil
	case Star:
		return "star", nil
	case Pentagon:
		return "pentagon", nil
	case Gem:
		return "gem", nil
	case Cross:
		return "cross", nil
	case Wedge:
		return "wedge", nil
	case Diamond:
		return "dia_imonde", nil
	case Oval:
		return "oval", nil
	default:
		return "", fmt.Errorf("shape not found")
	}
}

func NewRandomShape() *Shape {
	s := Shape(rand.Intn(10))
	return &s
}

type Vec2 struct {
	X int
	y int
}

func NewVec2(x, y int) *Vec2 {
	return &Vec2{X: x, y: y}
}

func NewRandomVec2(maxX, maxY int) *Vec2 {

	return &Vec2{
		X: rand.Intn(maxX),
		y: rand.Intn(maxY),
	}
}
