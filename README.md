# Boobot

Boobs-powered robot enslaved by online pirates, programmed to serve in random group chats. ( 🤖 Y 🤖 )

## Setup

### Mac OS

	# Install Homebrew and some handy tools
	/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
	brew install ssh-copy-id

	# Install the latest Go distribution
	brew install go --cross-compile-common

	# Set the GOPATH variable in your profile, e.g. config.fish:
	set -gx GOPATH $HOME/.go
	set -gx PATH $GOPATH/bin $PATH

	# Install Ansible
	sudo easy_install pip
	sudo pip install ansible

	# Install Docker for Mac
	# https://docs.docker.com/docker-for-mac/

	# Start having fun!
	make dev

### Raspberry Pi

	# Copy your SSH key to enable passwordless authentication
	echo "
	Host pi
	    HostName 192.168.1.32
	    User pi" >> ~/.ssh/config
	ssh-copy-id ~/.ssh/id_rsa pi

	# Install Docker Hypriot
	ssh pi
	curl -s https://packagecloud.io/install/repositories/Hypriot/Schatzkiste/script.deb.sh | sudo bash
	apt-get install docker-hypriot
	sudo usermod -a -G docker pi
