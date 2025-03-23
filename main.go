package main

import (
	"fmt"
	"log"

	"golang.org/x/sys/unix"
)

func main()  {
	totalMem, err := unix.SysctlUint64("hw.memsize")
	if err != nil {
		log.Fatalf("Error retrieving total memory: %v\n", err)
	}
	totalGB := float64(totalMem) / (1024 * 1024 * 1024)
	fmt.Printf("Total Memory: %d bytes\n (%.2f GB)", totalMem, totalGB)
}

func getSystemResults()  {
	
}