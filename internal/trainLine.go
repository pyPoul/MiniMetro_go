package internal

type TrainLine struct {
	stations   []*Station
	direct     []*Shape
	connection []*Shape
}

func (tl *TrainLine) AddStation(s *Station) {
	tl.stations = append(tl.stations, s)
}

func (tl *TrainLine) InsertStation(s *Station, idx int) {
	tl.stations = append(tl.stations[:idx], append([]*Station{s}, tl.stations[idx:]...)...)
}

func (tl *TrainLine) RemoveStation(idx int) {
	tl.stations = append(tl.stations[:idx], tl.stations[idx+1:]...)
}
