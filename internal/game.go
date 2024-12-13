package internal

import (
	et "github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

func init() {

}

type Game struct {
	stations   []*Station
	passengers []*Passenger
	trains     []*Train
}

func NewGame() *Game {
	return &Game{
		stations:   make([]*Station, 0),
		passengers: make([]*Passenger, 0),
		trains:     make([]*Train, 0),
	}
}

func (g *Game) AddStation(s *Station) {
	g.stations = append(g.stations, s)
}

func (g *Game) AddRandomStation() {
	g.stations = append(g.stations, NewRandomStation())
}

func (g *Game) AddPassenger(p *Passenger) {
	g.passengers = append(g.passengers, p)
}

func (g *Game) AddTrain(t *Train) {
	g.trains = append(g.trains, t)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

func (g *Game) Draw(screen *et.Image) {

	// white background
	screen.Fill(color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff})

	/* TODO: go to the station position and draw passengers
	for _, p := range g.passengers {
		op := &et.DrawImageOptions{}
		op.GeoM.Translate(float64(p.position.x), float64(p.position.y))
		screen.DrawImage(p.Image(), op)
	}
	*/

	for _, s := range g.stations {
		op := &et.DrawImageOptions{}
		op.GeoM.Scale(0.5, 0.5)
		op.GeoM.Translate(float64(s.position.X), float64(s.position.y))
		screen.DrawImage(s.Image(), op)
	}

	op := &et.DrawImageOptions{}
	op.GeoM.Scale(0.1, 0.1)
	op.GeoM.Translate(float64(g.stations[0].position.X), float64(g.stations[0].position.y))
	screen.DrawImage(g.stations[0].Image(), op)

}

func (g *Game) Update() error {
	return nil
}
