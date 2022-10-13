.PHONY: all

all: build install

build:
	go build -o towebp main.go

install:
	cp towebp /usr/local/bin
