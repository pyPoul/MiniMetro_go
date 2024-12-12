package internal

import (
	"fmt"
	et "github.com/hajimehoshi/ebiten/v2"
	"image"
	"os"
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

	f, err := os.Open(fmt.Sprintf("assets/%s/%s.png", typeObj, name))
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return et.NewImageFromImage(img), nil
}
