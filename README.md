Raspberry Pi thermometer reading 
 implementation in Go.
---

This mini project enables to measure the precisely the temperature at a low cost using the temperature probe: Maxim 1820.
You can find a nicely package probe on Amazon, Newegg marketplace.

Component required.
---
  * Raspberry Pi (any model) with SD card + Raspbian installed or equivalent
  * Resistor 4.7kΩ
  * Maxim 1820 : Temperature probe
  
Schematic
---
![schematic pi thermometer](/pi-thermometer-in-go.jpg)

Quick start
---
`$ sudo make all`

The steps are details below.
For ease of use, this application can be executed as root by setting port number to 80.
```http.ListenAndServe(":80", h)```

 you can leave the port to 8080 if you do not want to run this application as root.
 also you can start a frontend web server like Apache2 and redirect the traffic from port 80.

How to run / compile?
---
On my Raspbian runnning on Pi Zero.
Switch to command line.

`$ sudo -s`

Run go as script.

`$ go run spa_monitor.go`

For more speed the application can be compiled then executed:

`$ go build spa_monitor.go`

`$ ./spa_monitor`

How to run this a daemon using systemd ?
---
You need to edit the path in spa_monitor.service

To test the systemd configuration file

`$ make install start `

Enable the service

`$ make enable`

To stop the service

`$ make stop`

To restart

`$ make stop start`

Note
---
This process is not detached so I'm running a tmux session on the background.
 tmux enables to run terminal session on the background on a linux system.
https://en.wikipedia.org/wiki/Tmux

`$ apt install tmux`

`$ tmux new -s spa`

Extension
---

It's possible to add more probe to capture more datapoint.
Each probe (Maxim 1820) should be added in parallel and a new resistor must be added for each probe.
My current production application has two probes.

More details about implementation.
---
 https://www.modmypi.com/blog/ds18b20-one-wire-digital-temperature-sensor-and-the-raspberry-pi
