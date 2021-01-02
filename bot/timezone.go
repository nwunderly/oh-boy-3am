package bot

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Timezone struct {
	Name string `json:"name"`
	Offset float64 `json:"offset"`
}

var Timezones = LoadTimezoneData()

func LoadTimezoneData() []Timezone {
	file, err := os.Open("data/timezones.json")
	if err != nil {panic(err)}
	data, err := ioutil.ReadAll(file)

	var result []Timezone

	err = json.Unmarshal(data, &result)
	if err != nil {panic(err)}

	return result
}

func (tz Timezone) Is3Am() bool {
	return false
}
