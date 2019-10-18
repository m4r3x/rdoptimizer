package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/go-redis/redis"
	"github.com/m4r3x/rdoptimizer/constants"
	"github.com/m4r3x/rdoptimizer/helpers"
)

func main() {
	connectionString := getEnv("REDIS_CONNECTION", "localhost:6379")
	fmt.Println("Redis connection:", connectionString)
	client := redis.NewClient(&redis.Options{
		Addr:     connectionString,
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()

	if err != nil {
		fmt.Println("Redis connection: NOK")
		fmt.Println(pong, err)
		os.Exit(1)
	}

	fmt.Println("Redis connection: OK")
	client.FlushAll()
	fmt.Println("Redis flush: OK")
	time.Sleep(time.Second * 1)

	rps := constants.RPS()
	mode := constants.MODE()
	// Ensure no more than constants.RPS() are fired with every second.
	guard := make(chan struct{}, rps)

	cycle := 1
	for {
		for i := 1; i <= rps; i++ {
			guard <- struct{}{}
			go func() {
				helpers.Benchmark(client, rps, mode)
				<-guard
			}()
		}

		time.Sleep(time.Second * 1)
		printStats(cycle, client)
		cycle++
	}

}

func printStats(cycle int, client *redis.Client) {
	// Print Redis Memory Usage
	result, err := client.Do("MEMORY", "STATS").Result()
	if err != nil {
		fmt.Printf("Redis Alloc = UNKNOWN")
	} else {
		if stats, ok := result.([]interface{}); ok {
			for i, max := 0, len(stats)-1; i < max; i += 2 {
				if stats[i] == "total.allocated" {
					allocated := stats[i+1].(int64)
					fmt.Printf("Redis Alloc = %d KB", allocated/1024)
				}
			}
		}
	}

	// Print Host Memory Usage
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("\tHost Alloc = %v MB", m.Alloc/1024/1024)

	// Print Cycle
	fmt.Printf("\tCycle = %v", cycle)
	fmt.Printf("\n")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
