package drum

import "fmt"

func (p *Pattern) String() string {
	var tracks string
	for _, track := range p.Tracks {
		tracks = tracks + track.String() + "\n"
	}

	return fmt.Sprintf("Saved with HW Version: %s\nTempo: %v\n%s", p.HWVersion, p.Tempo, tracks)
}

func (t Track) String() string {
	var steps string
	for i, step := range t.Steps {
		if i%4 == 0 {
			steps = steps + "|"
		}
		if step {
			steps = steps + "x"
		} else {
			steps = steps + "-"
		}
	}
	steps = steps + "|"

	return fmt.Sprintf("(%d) %s\t%s", t.ID, t.Name, steps)
}
