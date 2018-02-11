package main

import (
	"encoding/json"
	//"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Group of measurement
type Datapoint []Thermometer

// Single thermoter
type Thermometer struct {
	ID    string
	Value float32
}

func ReadDatapoint(root string) Datapoint {
	files, err := filepath.Glob(root + "/*/w1_slave")
	if err != nil {
		log.Fatal(err)
	}

	datapoint := make(Datapoint, len(files), len(files))
	for i, f := range files {
		parts := strings.Split(f, "/")
		id := parts[len(parts)-2]
		datapoint[i] = *ReadTemperature(f, id)
	}
	return datapoint
}

// read thermometer data and return the value in celcius
func ReadTemperature(path string, id string) *Thermometer {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	regex, err := regexp.Compile(`\s+t=(\d+)`)
	found := regex.FindStringSubmatch(string(data))

	temperature, err := strconv.ParseFloat(found[1], 32)
	if err != nil {
		log.Fatal("fail to parse temperature")
	}
	return &Thermometer{id, float32(temperature / 1000)}
}

func handler(w http.ResponseWriter, r *http.Request) {
	root := "/sys/devices/w1_bus_master1/"
	datapoint := ReadDatapoint(root)
	data, err := json.Marshal(datapoint)
	if err != nil {
		log.Fatal("fail encoding datapoint to JSON")
	}
	fmt.Fprintf(w, string(data))
}

func main() {
	fmt.Println("start web server.\n http://spa")
	//fmt.Println(ReadDatapoint("data"))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
