package drum

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strings"
)

// DecodeFile decodes the drum machine file found at the provided path
// and returns a pointer to a parsed pattern which is the entry point to the
// rest of the data.
func DecodeFile(path string) (*Pattern, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return DecodeReader(file)
}

// DecodeReader decodes the drum machine file from an io.Reader
func DecodeReader(reader io.Reader) (*Pattern, error) {
	header := make([]byte, 6)
	_, err := io.ReadFull(reader, header)
	if err != nil {
		return nil, err
	}
	if string(header) != "SPLICE" {
		return nil, fmt.Errorf("not a splice file")
	}

	var size uint64
	err = binary.Read(reader, binary.BigEndian, &size)
	if err != nil {
		return nil, err
	}
	reader = io.LimitReader(reader, int64(size))

	var pattern Pattern

	version := make([]byte, 32)
	_, err = io.ReadFull(reader, version)
	if err != nil {
		return nil, err
	}
	pattern.HWVersion = strings.TrimRight(string(version), "\x00")

	err = binary.Read(reader, binary.LittleEndian, &pattern.Tempo)
	if err != nil {
		return nil, err
	}

	for {
		track, err := decodeTrack(reader)
		if err == io.EOF {
			// no more tracks to read
			break
		} else if err != nil {
			return nil, err
		}

		pattern.Tracks = append(pattern.Tracks, track)
	}

	return &pattern, nil
}

func decodeTrack(reader io.Reader) (Track, error) {
	var track Track
	err := binary.Read(reader, binary.BigEndian, &track.ID)
	if err != nil {
		return Track{}, err
	}

	var nameLength uint32
	err = binary.Read(reader, binary.BigEndian, &nameLength)
	if err != nil {
		return Track{}, err
	}

	name := make([]byte, nameLength)
	_, err = io.ReadFull(reader, name)
	if err != nil {
		return Track{}, err
	}
	track.Name = string(name)

	for len(track.Steps) < 16 {
		var step uint8
		err = binary.Read(reader, binary.BigEndian, &step)
		if err != nil {
			return Track{}, err
		}

		track.Steps = append(track.Steps, step == 1)
	}

	return track, nil
}
