# spa monitor management script
#

# systemd service configuration file
script=spa_monitor.service

all: clean init compile run

# create csv file
init:
	touch history.csv

# remove compiled file
clean:
	rm spa_monitor

# compile application
compile:
	go build *.go

# SystemD - optional

# install as systemd service
install:
	sudo cp $(script) /etc/systemd/system/

# start daemon service
start:
	sudo systemctl start $(script)

# stop daemon service
stop:
	sudo systemctl stop $(script)

# enable daemon service
enable:
	sudo systemctl enable  $(script)
