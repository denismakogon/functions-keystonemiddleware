# Just builds
.PHONY: build dep

DIR := ${CURDIR}

dep:
	glide install --strip-vendor

build:
	go build -o functions-with-keystone-auth
