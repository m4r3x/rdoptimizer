package helpers

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis"
	pb "github.com/golang/protobuf/proto"
	"github.com/m4r3x/rdoptimizer/proto"
)

// Benchmark runs benchmark based on MODE defined inside constants.
func Benchmark(client *redis.Client, rps int, mode int) {
	message := ""
	switch mode {
	case 1:
		message = rawEventsBenchmark(client, false)
	case 2:
		message = mapsEventsBenchmark(client, false)
	case 3:
		message = rawEventsBenchmark(client, true)
	case 4:
		message = mapsEventsBenchmark(client, true)
	case 5:
		message = rawTextBenchmark(client)
	case 6:
		message = protoTextBenchmark(client)
	}

	if message != "" && rps < 5 {
		fmt.Println(message)
	}
}

func rawEventsBenchmark(client *redis.Client, withIP bool) string {
	device := Device()
	country, city := Location()
	event := Event()
	ip := IPV6Address()

	if withIP {
		client.SAdd(makeTimestampKey(), fmt.Sprintf("%s.%s.%s.%s.%s", device, country, city, event, ip))
		return fmt.Sprintf("%s %s %s %s %s %s \n", makeTimestampKey(), device, country, city, event, ip)
	}

	client.SAdd(makeTimestampKey(), fmt.Sprintf("%s.%s.%s.%s", device, country, city, event))
	return fmt.Sprintf("%s %s %s %s %s \n", makeTimestampKey(), device, country, city, event)
}

var ipsMap = map[string]int{}
var ipsMapMutex = sync.RWMutex{}

func mapsEventsBenchmark(client *redis.Client, withIP bool) string {
	device := DeviceKey()
	country, city := LocationKey()
	event := EventKey()
	ip := IPV6Address()

	if withIP {
		ipsMapMutex.RLock()
		ipKey, isset := ipsMap[ip]
		ipsMapMutex.RUnlock()
		if !isset {
			ipKey = len(ipsMap) + 1
			ipsMapMutex.Lock()
			ipsMap[ip] = ipKey
			ipsMapMutex.Unlock()
		}

		client.SAdd(makeTimestampKey(), fmt.Sprintf("%d.%d.%d.%d.%d", device, country, city, event, ipKey))
		return fmt.Sprintf("%s %d %d %d %d %d \n", makeTimestampKey(), device, country, city, event, ipKey)
	}

	client.SAdd(makeTimestampKey(), fmt.Sprintf("%d.%d.%d.%d", device, country, city, event))
	return fmt.Sprintf("%s %d %d %d %d \n", makeTimestampKey(), device, country, city, event)
}

type message struct {
	Body   string
	Header string
}

func rawTextBenchmark(client *redis.Client) string {
	t := Text()
	m := message{
		Body:   t,
		Header: t,
	}
	jm, err := json.Marshal(m)

	if err != nil {
		fmt.Println("Json failed to Marshal data: %v", err)
		os.Exit(1)
	}

	client.SAdd(makeTimestampKey(), jm)

	return t
}

func protoTextBenchmark(client *redis.Client) string {
	t := Text()
	m := &proto.Message{
		Body:   pb.String(t),
		Header: pb.String(t),
	}
	pm, err := pb.Marshal(m)

	if err != nil {
		fmt.Println("Proto failed to Marshal data: %v", err)
		os.Exit(1)
	}

	client.SAdd(makeTimestampKey(), pm)

	return t
}

// Return timestamp rounded to `bucketSizeInSeconds` seconds bucket.
func makeTimestampKey() string {
	bucketSizeInSeconds := 10
	exactTimestamp := time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))

	return strconv.FormatInt(exactTimestamp/(1000*int64(bucketSizeInSeconds)), 10)
}
