// Refer README.md for project instructions and YAMLinfo.md for CI/CD pipeline basics with GitHub Actions.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countLines bool, countBytes bool) int {
	// A scanner is used to read text from a reader(such as files)]
	scanner := bufio.NewScanner(r)

	// Define the scanner splite type to words(default is split by lines)
	// scanner.Split(bufio.ScanWords)(v.1.0)

	// If the count lines flag is not set, we want to count words so we define
	// the scanner split type to words(default is split by lines)
	if countLines {
		scanner.Split(bufio.ScanLines)
	} else if countBytes {
		scanner.Split(bufio.ScanBytes)
	} else {
		scanner.Split(bufio.ScanWords)
	}

	// Defining a counter
	wCtr := 0

	// For every word scanned, increment the counter
	for scanner.Scan() {
		wCtr++
	}

	// Return the total
	return wCtr
}

func main() {

	// Defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")

	// Defining a boolean flag -b to count bytes instead of words or lines
	bytes := flag.Bool("b", false, "Count bytes")

	// Parsing the flags provided by the user
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *bytes))
}
