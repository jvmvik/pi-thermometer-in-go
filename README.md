Raspberry Pi thermometer reading 
 implementation in Go.
---

For ease of use, this application can be executed as root by setting port number to 80.
```http.ListenAndServe(":80", h)```
 you can leave the port to 8080 if you do not want to run this application as root.
 also you can start a frontend web server like Apache2 and redirect the traffic from port 80.

Quick start
---
`$ sudo make all`

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

Reference
---
 https://www.modmypi.com/blog/ds18b20-one-wire-digital-temperature-sensor-and-the-raspberry-pi
