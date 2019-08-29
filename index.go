package main

import (
	"fmt"
	"math/rand"
	"time"

	bobsmap "./bob"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(bobsmap.New())
}
