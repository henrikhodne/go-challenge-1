// Package drum is supposed to implement the decoding of .splice drum machine files.
// See golang-challenge.com/go-challenge1/ for more information
package drum

type Pattern struct {
	HWVersion string
	Tempo     float32
	Tracks    []Track
}

type Track struct {
	ID    uint8
	Name  string
	Steps []bool
}
