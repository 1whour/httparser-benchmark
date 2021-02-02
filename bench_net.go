package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"time"
)

var data = []byte(
	"POST /joyent/http-parser HTTP/1.1\r\n" +
		"Host: github.com\r\n" +
		"DNT: 1\r\n" +
		"Accept-Encoding: gzip, deflate, sdch\r\n" +
		"Accept-Language: ru-RU,ru;q=0.8,en-US;q=0.6,en;q=0.4\r\n" +
		"User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) " +
		"Chrome/39.0.2171.65 Safari/537.36\r\n" +
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9," +
		"image/webp,*/*;q=0.8\r\n" +
		"Referer: https://github.com/joyent/http-parser\r\n" +
		"Connection: keep-alive\r\n" +
		"Transfer-Encoding: chunked\r\n" +
		"Cache-Control: max-age=0\r\n\r\nb\r\nhello world\r\n0\r\n")

var kBytes = int64(8) << 30

func bench(iterCount int64, silent bool) {
	var start time.Time
	if !silent {
		start = time.Now()
	}

	fmt.Printf("req_len=%d\n", len(data))
	for i := int64(0); i < iterCount; i++ {
		_, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(data)))
		if err != nil {
			panic(err.Error())
		}
	}

	if !silent {
		end := time.Now()

		fmt.Printf("Benchmark result:\n")

		elapsed := end.Sub(start) / time.Second

		total := iterCount * int64(len(data))
		bw := float64(total) / float64(elapsed)

		fmt.Printf("%.2f mb | %.2f mb/s | %.2f req/sec | %.2f s\n",
			float64(total)/(1024*1024),
			bw/(1024*1024),
			float64(iterCount)/float64(elapsed),
			float64(elapsed))

	}
}

func main() {
	iterations := kBytes / int64(len(data))
	bench(iterations, false)
}
