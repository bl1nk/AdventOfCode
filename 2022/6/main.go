package main

import (
	"fmt"
	"os"

	"github.com/bl1nk/adventofcode/2022/6/comms"
)

func main() {
	c, err := comms.New("./2022/6/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	m := c.FindStartOfPacketMarker()
	fmt.Printf("Part 1: %d packets need to be processed before the first start-of-packet marker is detected.\n", m)
	m = c.FindStartOfMessageMarker()
	fmt.Printf("Part 2: %d packets need to be processed before the first start-of-message marker is detected.\n", m)
}
