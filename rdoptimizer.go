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
	time.Sleep(time.Second * 1)

	// Ensure no more than constants.RPS are fired with every second.
	guard := make(chan struct{}, constants.RPS)

	cycle := 1
	for {
		for i := 1; i <= constants.RPS; i++ {
			guard <- struct{}{}
			go func() {
				helpers.Benchmark(client)
				<-guard
			}()
		}

		time.Sleep(time.Second * 1)
		PrintStats(cycle)
		cycle += 1
	}

}

func PrintStats(cycle int) {
	// fmt.Println("\033[2J")
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Host Alloc = %v MiB", m.Alloc/1024/1024)
	fmt.Printf("\tCycle = %v", cycle)
	fmt.Printf("\n")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
