package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func init() {
	log.SetFlags(0)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <number>", os.Args[0])
	}

	n, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time.Unix(0, n))
}
