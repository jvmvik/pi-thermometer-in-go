package main

import (
 "fmt"
 "io/ioutil"
 "log"
 "regexp"
 "strconv"
)

// read thermometer data and return the value in celcius
func main() {

 path := "/sys/devices/w1_bus_master1/28-0517c1a38eff/w1_slave"
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
 fmt.Println(temperature/1000)
}
