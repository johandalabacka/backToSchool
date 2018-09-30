package main

import (
	"bufio"
	"os"
	"testing"
)

func TestMakeNextLine(t *testing.T) {
	testData := []struct {
		oldFirstFilled int
		oldLine        string
		newFirstFilled int
		newLine        string
	}{
		{0, "", 0, ""},
		{0, "#", 0, ""},
		{0, "##", -1, "#..#"},
		{0, "##..##", -1, "#..##..#"},
		{-2, "##..##", -3, "#..##..#"},
	}

	for _, d := range testData {
		resultLine, resultFirstFilled := makeNextLine(d.oldLine, d.oldFirstFilled)
		if resultFirstFilled != d.newFirstFilled || resultLine != d.newLine {
			t.Errorf("Wanted %d %q got %d %q\n",
				d.newFirstFilled, d.newLine,
				resultFirstFilled, resultLine)
		}
	}
}

func TestPatternType(t *testing.T) {
	testData := []struct {
		line        string
		patternType string
	}{
		{"", "vanishing"},
		{"#", "vanishing"},
		{"####", "blinking"},
		{"# ## #", "blinking"},
		{"# # # #", "blinking"},
		{"# ##", "blinking"},
		{"# ###", "gliding"},
		{"# #", "vanishing"},
	}

	for _, d := range testData {
		result := patternType(d.line)
		if result != d.patternType {
			t.Errorf("Wanted %s got %s for %q\n", d.patternType, result, d.line)
		}
	}
}

func BenchmarkMain(b *testing.B) {
	patterns := []string{}
	f, err := os.Open("patterns.txt")
	if err != nil {
		panic("Could not open patterns file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		pattern := scanner.Text()
		patterns = append(patterns, pattern)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, pattern := range patterns {
			_ = patternType(pattern)
		}
	}
}
