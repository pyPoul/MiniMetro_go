package internal

import "fmt"

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

type Vec2 struct {
	X int
	Y int
}
