# Just builds
.PHONY: build

DIR := ${CURDIR}

build:
	go build -o keystonemiddleware
