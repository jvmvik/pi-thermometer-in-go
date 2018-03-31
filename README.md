Raspberry Pi thermometer reading 
 implementation in Go.
---

For ease of use, this application can be executed as root by setting port number to 80.
````http.ListenAndServe(":80", h)```
 you can leave the port to 8080 if you do not want to run this application as root.
 also you can start a frontend web server like Apache2 and redirect the traffic from port 80.

How to run / compile?
---
On my Raspbian runnning on Pi Zero.
Switch to command line.
$>sudo -s

Run go as script.
$>go run spa_monitor.go

For more speed the application can be compiled then executed:
$>go build spa_monitor.go
$>./spa_monitor

Note
---
this process is not detached so I'm running a tmux session on the background.
 tmux enables to run terminal session on the background on a linux system.
https://en.wikipedia.org/wiki/Tmux

$>apt-get install tmux
$>tmux new -s spa

Reference
---
 https://www.modmypi.com/blog/ds18b20-one-wire-digital-temperature-sensor-and-the-raspberry-pi
