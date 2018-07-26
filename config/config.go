package config

import (
	"encoding/json"
	"io/ioutil"
)

// OutputProfile is the container for output device characteristics
type OutputProfile struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

// Read will convert a jsonFile to an OutputProfile
func Read(path string) OutputProfile {
	file, err := ioutil.ReadFile("./config.json")
	if nil != err {
		panic(err)
	}

	var profile OutputProfile
	json.Unmarshal(file, &profile)
	return profile
}
