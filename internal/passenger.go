package internal

import et "github.com/hajimehoshi/ebiten/v2"

type Passenger struct {
	dst   *Shape
	image *et.Image
}

func (p *Passenger) Image() *et.Image {

	if p.image != nil {
		return p.image
	}

	var err error
	p.image, err = ObjToImg(p)
	if err != nil {
		panic(err)
	}

	return p.image
}
