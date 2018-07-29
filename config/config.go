package config

import (
	"encoding/json"
	"io/ioutil"
)

// DocumentMargins has all the margins that will be enforced for a document
type DocumentMargins struct {
	Top    int `json:"top"`
	Bottom int `json:"bottom"`
	Left   int `json:"left"`
	Right  int `json:"right"`
}

// OutputProfile is the container for output device characteristics
type OutputProfile struct {
	Margins DocumentMargins `json:"margins"`
}

// Read will convert a jsonFile to an OutputProfile
func Read(path string) OutputProfile {
	file, err := ioutil.ReadFile(path)
	if nil != err {
		panic(err)
	}

	var profile OutputProfile
	json.Unmarshal(file, &profile)
	return profile
}
