package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader) int {
	// A scanner is used to read text from a reader(such as files)]
	scanner := bufio.NewScanner(r)

	// Define rhe scanner splite type to words(default is split by lines)
	scanner.Split(bufio.ScanWords)

	// Defining w counter
	wCtr := 0
	
	// For every word scanned, increment the counter
	for scanner.Scan() {
		wCtr++
	}

	// Return the total
	return wCtr
}

func main() {

	fmt.Println(count(os.Stdin))
}
