package helpers

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/m4r3x/rdoptimizer/constants"
)

func Benchmark(client *redis.Client) {
	switch constants.MODE {
	case 1:
		rawEventsBenchmark()
	case 2:
		mapsEventsBenchmark()
	case 3:
		rawTextBenchmark()
	case 4:
		protoTextBenchmark()
	}
}

func rawEventsBenchmark() {
	device := Device()
	country, city := Location()
	event := Event()
	ip := IPV6Address()

	if constants.RPS < constants.RPS_VERBOSE {
		fmt.Printf("%s %s %s %s %s \n", device, country, city, event, ip)
	}
	return
}

func mapsEventsBenchmark() {

}

func rawTextBenchmark() {

}

func protoTextBenchmark() {

}
