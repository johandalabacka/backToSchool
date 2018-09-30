package main

// patternType given the first line it will generate the
// next lines and see if any line has a certain pattern.
// vanishing: next generated line is empty
// blinking: next generated line matches a previous line on the same positon
// gliding: next generated line matches a previous line but with another position
// none: if no pattern is found after 100
func patternType(firstLine string) string {
	// All previous lines mapping to which position
	// the first filled square is on
	previousLines := make(map[string]int)

	// The first line has always position 0 for first filled square
	line, firstFilled := firstLine, 0
	previousLines[line] = firstFilled

	// Examine at most 100 lines. First is already given
	for lineCount := 2; lineCount <= 100; lineCount++ {
		// Make next line from the previous one
		line, firstFilled = makeNextLine(line, firstFilled)

		// Check if the new line matches a pattern
		if line == "" {
			return "vanishing"
		}
		previousFirstFilled, found := previousLines[line]
		if found {
			if previousFirstFilled == firstFilled {
				return "blinking"
			}
			return "gliding"
		}

		// If no pattern found. Store it and contine
		previousLines[line] = firstFilled
	}
	// No pattern found
	return "other"
}

// makeNextLine given a line and the position where it starts
// generates the next line and its position
func makeNextLine(line string, firstPos int) (string, int) {
	len := len(line)

	// New line can be at most two chars longer
	// one square at each end
	nextLen := len + 2
	b := make([]byte, nextLen)

	firstFilled := 0
	firstFilledFound := false
	lastFilled := 0

	// Index of first square in (old) line
	j := -1
	// Iterate over squares in new line
	for i := 0; i < nextLen; i++ {

		// Is the square above filled
		filled := j >= 0 && j < len && line[j] == '#'
		// Should current square be filled in the new line
		shouldFill := false

		// Count the filled squares in the (old) line
		count := 0
		if j-1 >= 0 && line[j-1] == '#' {
			count++
		}
		if j-2 >= 0 && line[j-2] == '#' {
			count++
		}
		if j+1 < len && line[j+1] == '#' {
			count++
		}
		if j+2 < len && line[j+2] == '#' {
			count++
		}

		if !filled {
			/*
				Rule #1, the square above is blank:
				If there are 2 or 3 filled squares in total next to it
				(taking into account 4 squares, 2 on each sides)
				it will be filled. If not, it will be left blank.
			*/
			if count == 2 || count == 3 {
				shouldFill = true
			}
		} else {
			/*
				Rule #2, the square above is filled:
				If there are 2 or 4 squares filled in total next to it
				(taking into account 4 squares, 2 on each sides)
				it will be filled. If not, it will be left blank.
			*/
			if count == 2 || count == 4 {
				shouldFill = true
			}
		}

		if shouldFill {
			b[i] = '#'
			if !firstFilledFound {
				firstFilledFound = true
				firstFilled = i
			}
			lastFilled = i
		} else {
			b[i] = '.'
		}

		j++ // Advance in old line
	}

	if !firstFilledFound {
		return "", 0
	}

	nextLine := string(b[firstFilled : lastFilled+1])

	// First position is
	nextFirstPos := firstPos - 1 + firstFilled

	return nextLine, nextFirstPos
}
