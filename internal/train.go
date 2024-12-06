package internal

type Wagon struct {
	passengers []*Passenger
}

type Train struct {
	lastPos    *Station
	passengers []*Passenger
	wagons     []*Wagon
	line       *TrainLine
	dir        int8
}

func (t *Train) Connect(w *Wagon) {
	t.wagons = append(t.wagons, w)
}

// TODO: change passengers from this wagon
func (t *Train) ChangeWagon(t2 *Train) {

	lastIdx := len(t.wagons) - 1
	if len(t.wagons) > 0 && len(t.wagons[lastIdx].passengers) != 0 {
		lastWagon := t.wagons[lastIdx]
		t.wagons = t.wagons[:lastIdx]
		t2.wagons = append(t2.wagons, lastWagon)
	}
}

func (t *Train) UpdatePos(s *Station) {

	t.lastPos = s
	next := len(t.line.stations) + int(t.dir)
	if next < 0 || next >= len(t.line.stations) {
		t.dir *= -1
	}
}

// TODO: verify train & wagons capacity
func (t *Train) AddPassenger(p *Passenger) {
	t.passengers = append(t.passengers, p)
}

// TODO: verify wagons
func (t *Train) RemovePassengers(s *Shape) []*Passenger {

	var newPassengers []*Passenger
	var removed []*Passenger
	for _, p := range t.passengers {
		if p.dst == s {
			removed = append(removed, p)
			continue
		}
		newPassengers = append(newPassengers, p)
	}
	t.passengers = newPassengers
	return removed
}

func (t *Train) ChangePassengers(sh *Shape, st *Station) {

	for _, p := range t.RemovePassengers(sh) {
		st.AddPassenger(p)
	}
}

func (t *Train) VerifyAndAdd(ps []*Passenger) []*Passenger {

	// passengers that can't be added
	var remain []*Passenger

	// sum of passengers in the train
	sum := len(t.passengers)
	for _, w := range t.wagons {
		sum += len(w.passengers)
	}
	// remaining places
	limit := (len(t.wagons)+1)*6 - sum

	for _, p := range ps {
		if limit <= 0 {
			break
		}
		already := false

		// direct line TODO: verify train direction
		for _, s := range t.line.direct {
			if s == p.dst {
				t.AddPassenger(p)
				limit--
				already = true
				break
			}
		}
		if already {
			continue
		}
		// non-direct line TODO: verify train direction
		for _, s := range t.line.connection {
			if s == p.dst {
				t.AddPassenger(p)
				limit--
				already = true
				break
			}
		}
		// if the passenger can't be added
		if !already {
			remain = append(remain, p)
		}
	}
	return remain
}
