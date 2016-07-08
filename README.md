# Boobot

Boobs-powered robot enslaved by online pirates, trapped in a random group chat!

## Why?

- Because BOOBS!
- Bewbs!
- Jugs!
- brlbrlbrlbrl!!!
- ( 🤖 Y 🤖 )

## Setup

	# Install [Homebrew](http://brew.sh/)
	/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

	# Install the latest [Go](https://golang.org/dl/) distribution
	brew install go --cross-compile-common

	# Set the GOPATH variable, e.g. config.fish:
	set -gx GOPATH $HOME/.go
	set -gx PATH $GOPATH/bin $PATH

	# Download and install packages and dependencies
	make install

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
