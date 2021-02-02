all: build test

# 运行和测试标准库
run.test.net: build.net test.net

build.net:
	go build bench_net.go
test.net:
	./bench_net

build:
	go build bench.go

test:
	./bench
