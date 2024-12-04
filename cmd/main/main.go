package main

import (
	et "github.com/hajimehoshi/ebiten/v2"
	"github.com/pyPoul/MiniMetro_go/internal"
)

func main() {

	et.SetFullscreen(true)
	et.SetWindowTitle("Mini Metro")

	if err := et.RunGame(&internal.Game{}); err != nil {
		panic(err)
	}

}
