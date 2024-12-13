package internal

import (
	"fmt"
	et "github.com/hajimehoshi/ebiten/v2"
	etil "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func ObjToImg(obj interface{}) (*et.Image, error) {

	var typeObj string
	var shape *Shape

	switch o := obj.(type) {

	case *Station:
		typeObj = "stations"
		shape = o.shape

	case *Passenger:
		typeObj = "passengers"
		shape = o.dst

	default:
		return nil, fmt.Errorf("unknown object type: %T", obj)
	}

	name, err := shape.ToString()
	if err != nil {
		return nil, err
	}

	var img *et.Image
	img, _, err = etil.NewImageFromFile(fmt.Sprintf("assets/%s/%s.png", typeObj, name))
	if err != nil {
		return nil, err
	}

	return img, nil
}
