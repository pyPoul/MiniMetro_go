package internal

import (
	et "github.com/hajimehoshi/ebiten/v2"
	"image"
	"os"
)

type Station struct {
	nb       uint8
	shape    *Shape
	position Vec2
	rarity   uint8
	image    *et.Image
	inter    bool
	pQueue   []*Passenger
}

func (s *Station) Image() *et.Image {

	if s.image != nil {
		return s.image
	}

	name, err := s.shape.ToString()
	if err != nil {
		panic(err)
	}

	f, err := os.Open("assets/stations/" + name + ".png")
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	s.image = et.NewImageFromImage(img)

	return s.image
}

func (s *Station) AddPassenger(p *Passenger) {
	s.pQueue = append(s.pQueue, p)
}

func (s *Station) RemovePassenger() *Passenger {
	if len(s.pQueue) == 0 {
		return nil
	}

	p := s.pQueue[0]
	s.pQueue = s.pQueue[1:]

	return p
}
