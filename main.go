package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://github.com/wunderdogsw/wunderpahkina-vol9/blob/master/README.md
// I give Wunderdog the right to publish this source code.
// Author: Johan Dahl johan.dalabacka@gmail.com

// I also tried two different variants
// 1) Using go-routines, one for each line. But my portable has only 1 cpu with 2 cores
// so I gained nothing on this small input set but making the program slightly more
// complicated
// 2) Reusing buffers instead of allocating new ones for each line. It saved some memory
// but it didn't make any difference in execution speed and just introduced a global
// variable.

// main iterates over the lines in patterns.txt
// and finds the patternType
func main() {
	f, err := os.Open("patterns.txt")
	if err != nil {
		panic("Could not open patterns file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		pattern := scanner.Text()

		patternType := patternType(pattern)
		fmt.Println(patternType)
	}
}
