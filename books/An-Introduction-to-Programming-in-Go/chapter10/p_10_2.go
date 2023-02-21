package main

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	start := time.Now()
	fmt.Println("Sleeping for 2 seconds...")
	Sleep(2 * time.Second)
	fmt.Printf("Slept for %v\n", time.Since(start))
}
