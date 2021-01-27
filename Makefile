all: build test

build:
	go build bench.go

test:
	./bench
