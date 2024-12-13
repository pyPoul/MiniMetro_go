package main

import (
	et "github.com/hajimehoshi/ebiten/v2"
	it "github.com/pyPoul/MiniMetro_go/internal"
	"github.com/pyPoul/MiniMetro_go/internal/options"
)

func openDisplayOptions() {

	d := options.Display{}
	d.Load()

	if !d.Show {
		return
	}
	et.SetFullscreen(d.Fullscreen)
	et.SetWindowSize(int(d.Resolution.W), int(d.Resolution.H))
	if d.Resizable {
		et.SetWindowResizingMode(et.WindowResizingModeEnabled)
	}
	et.SetWindowTitle("Mini Metro")
	et.SetWindowDecorated(d.Show)
}

func main() {

	game := it.NewGame()
	game.AddRandomStation()
	game.AddRandomStation()
	game.AddRandomStation()
	game.AddRandomStation()

	if err := et.RunGame(game); err != nil {
		panic(err)
	}

}
