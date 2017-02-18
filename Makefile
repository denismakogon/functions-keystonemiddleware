# Just builds
.PHONY: all dep build

DIR := ${CURDIR}

dep:
	glide install --strip-vendor

build:
	go build -o keystonemiddleware

all: dep build
