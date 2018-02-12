package main

import (
	"encoding/json"
    "encoding/csv"
	//"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
    "html/template"
    "time"
    "os"
    "bufio"
)

// Group of measurement
type Datapoint []Thermometer

// Single thermoter
type Thermometer struct {
	ID    string
	Value float32
}

type DisplayThermometer struct {
    Air float32
    Water float32
    HistoryAir string
    HistoryWater string
}

func GetRoot() string {
    // For debug
    //return "data/"
    // On the raspberry pi
    return "/sys/devices/w1_bus_master1/"
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

func jsonHandler(w http.ResponseWriter, r *http.Request) {
    datapoint := ReadDatapoint(GetRoot())
	data, err := json.Marshal(datapoint)
	if err != nil {
		log.Fatal("fail encoding datapoint to JSON")
	}
	fmt.Fprintf(w, string(data))
}

func Round(x, unit float32) float32 {
	return float32(int32(x/unit+0.5)) * unit
}

func ReadDisplay(root string) DisplayThermometer {
    datapoint := ReadDatapoint(root) 
    var display DisplayThermometer
    for _, item := range datapoint {
        if(item.ID == "28-0517c1a38eff") {
            display.Water = Round(item.Value, 0.1)
        }
        if(item.ID == "28-0517c207d1ff") {
            display.Air = Round(item.Value, 0.1)
        }
    }
    
    return display
}

func ReadHistory(path string, display *DisplayThermometer) {
    
    file, err := os.Open("history.csv")
    if err != nil {
        log.Fatal("cannot open file")
    }
    defer file.Close()
    
    r := csv.NewReader(bufio.NewReader(file))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

    offset := len(records) - 7
    display.HistoryAir   = "0,223 "
    display.HistoryWater = "0,223 "
    for i := offset; i < len(records); i++ {
        x := (i - offset) * 80
        t_air, err := strconv.ParseFloat(records[i][0], 64)
        if err != nil {
            log.Fatal(err)
        }
        t_water, _ := strconv.ParseFloat(records[i][1], 64)
        if err != nil {
            log.Fatal(err)
        }
        y_air := (50 - t_air) * 4 + 25
        y_water := (50 - t_water) * 4 + 25
        
        display.HistoryAir += i2s(x) + "," + f2s(y_air) + " "
        display.HistoryWater += i2s(x) + "," + f2s(y_water) + " "
    }
    
    display.HistoryAir   += " 480,223"
    display.HistoryWater += " 480,223"
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    if(r.URL.Path != "/") {
        http.NotFound(w, r)
        return
    }
    // Current value to display
    display := ReadDisplay(GetRoot())
    ReadHistory("history.csv", &display)
    t, _ := template.ParseFiles("index.html")
    t.Execute(w, display)
}

func i2s(v int) string {
    return strconv.FormatInt(int64(v), 10)
}

func f2s(v float64) string {
    return strconv.FormatFloat(v, 'f', -1, 32)
}

func recordHandler(w http.ResponseWriter, r *http.Request) {
    
    p := ReadDisplay(GetRoot())
    
    // header {"air", "water", "timestamp"}
    record := []string{strconv.FormatFloat(float64(p.Air), 'f', -1, 32),
                       strconv.FormatFloat(float64(p.Water), 'f', -1, 32),
                       strconv.FormatInt(time.Now().Unix(), 10)}

    file, err := os.OpenFile("history.csv", os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        log.Fatal("cannot open file")
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    if err := writer.Write(record); err != nil {
        log.Fatalln("error writing record to csv:", err)
    }
    writer.Flush()
    
    if err := writer.Error(); err != nil {
		log.Fatal(err)
        fmt.Fprintf(w, "{\"status\":\"fail\"}")
        return
	}
    
    fmt.Fprintf(w, "{\"status\":\"ok\"}")
}

func main() {
	fmt.Println("start web server.\n http://spa")
    h := http.NewServeMux()
	h.HandleFunc("/json", jsonHandler)
    
    h.HandleFunc("/record", recordHandler)
    
    fs := http.FileServer(http.Dir("static"))
    h.Handle("/static/", http.StripPrefix("/static/", fs))
    
    h.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8888", h)
}
