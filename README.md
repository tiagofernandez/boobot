# Boobot

Boobs-powered robot enslaved by online pirates, programmed to serve in random group chats. ( 🤖 Y 🤖 )


## Setup

### Mac OS

	# Install Homebrew
	/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

	# Install the latest Go distribution
	brew install go --cross-compile-common

	# Set the GOPATH variable in your profile, e.g. config.fish:
	set -gx GOPATH $HOME/.go
	set -gx PATH $GOPATH/bin $PATH

	# Download packages, build and run!
	make

### Docker

Install [Docker for Mac](https://docs.docker.com/docker-for-mac/).

	make docker-image
	make docker-run


## Platforms

### Raspberry Pi

	# Configuration
	ssh pi@192.168.1.xx
	sudo raspi-config

	# Install Docker
	curl -s https://packagecloud.io/install/repositories/Hypriot/Schatzkiste/script.deb.sh | sudo bash
	apt-get install docker-hypriot

	# Shutdown
	sudo shutdown -h now
