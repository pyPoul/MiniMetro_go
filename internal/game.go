package internal

import (
	et "github.com/hajimehoshi/ebiten/v2"
	etu "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

func (g *Game) Draw(screen *et.Image) {
	return
}

func (g *Game) Update() error {
	return nil
}
