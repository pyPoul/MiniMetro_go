package options

import (
	"encoding/json"
	"fmt"
	"os"
)

type Display struct {
	Show       bool `json:"display"`
	Fullscreen bool `json:"fullscreen"`
	Resizable  bool `json:"resizable"`
	Resolution struct {
		W float64 `json:"width"`
		H float64 `json:"height"`
	} `json:"resolution"`
}

func (d *Display) Load() {

	base := func(d *Display) {
		d.Show = true
		d.Fullscreen = false
		d.Resizable = true
		d.Resolution = struct {
			W float64 `json:"width"`
			H float64 `json:"height"`
		}{W: 1280, H: 720}
	}

	jsonFile, err := os.Open("shared/options/display.json")
	if err != nil {
		fmt.Println("Error opening file:\n", err)
		base(d)
		return
	}

	defer func() {
		err = jsonFile.Close()
		if err != nil {
			fmt.Println("Error closing file:\n", err)
		}
	}()

	if err = json.NewDecoder(jsonFile).Decode(d); err != nil {
		fmt.Println("Error decoding json:\n", err)
		base(d)
		return
	}
}
