package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	snr := bufio.NewScanner(os.Stdin)
	for {
		snr.Scan()
		line := snr.Text()
		if len(line) == 0 || line == "exit" {
			break
		}
		fields := strings.Fields(line)
		fmt.Printf("Fields: %q\n", fields)
	}
}
