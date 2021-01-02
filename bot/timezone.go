package bot

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type Timezone struct {
	Name      string `json:"Name"`
	OffsetHr  int    `json:"OffsetHr"`
	OffsetMin int    `json:"OffsetMin"`
}

var Timezones = LoadTimezoneData()

func LoadTimezoneData() []Timezone {
	file, err := os.Open("data/timezones.json")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(file)

	var result []Timezone

	err = json.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (tz Timezone) Is3Am() bool {
	timestamp := time.Now().UTC()
	offset := time.Hour*time.Duration(tz.OffsetHr) + time.Minute*time.Duration(tz.OffsetMin)
	localTime := timestamp.Add(offset)
	return localTime.Hour() == 3 && localTime.Minute() == 0
}
