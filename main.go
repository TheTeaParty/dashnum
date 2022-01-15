package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Task 1: Difficulty - 1/10; Estimate - 30 minutes; Actual - 27 minutes
// Task 2: Difficulty - 1/10; Estimate - 10 minutes

func main() {
}

// testValidity takes the string as an input, and returns boolean flag `true` if the given
// string complies with the format, or `false` if the string does not comply
func testValidity(str string) bool {
	var re = regexp.MustCompile(`^([0-9]+-[a-zA-Z]+-?)+$`)
	return re.MatchString(str)
}

// averageNumber takes the string, and returns the average number from all the numbers
func averageNumber(str string) float64 {
	var re = regexp.MustCompile(`([0-9]+)-?`)
	matches := re.FindAllStringSubmatch(str, -1)
	var sum int

	for _, match := range matches {
		i, _ := strconv.Atoi(match[1])
		sum = sum + i
	}

	return float64(sum) / float64(len(matches))
}

// wholeStory takes the string, and returns a text that is composed of all the text words separated by
// spaces, e.g. `story` called for the string `1-hello-2-world` should return text: `"hello world"
func wholeStory(str string) string {
	var re = regexp.MustCompile(`([a-zA-Z]+)-?`)
	matches := re.FindAllStringSubmatch(str, -1)
	var story []string

	for _, match := range matches {
		story = append(story, match[1])
	}

	return strings.Join(story, " ")
}

// storyStats returns four things:
//   * the shortest word
//   * the longest word
//   * the average word length
//   * the list (or empty list) of all words from the story that have the length the same as the average length rounded up and down.
func storyStats(str string) (string, string, float64, []string) {
	var re = regexp.MustCompile(`([a-zA-Z]+)-?`)
	matches := re.FindAllStringSubmatch(str, -1)
	words := make(map[int][]string, 0)
	longestLen := 0
	shortestLen := math.MaxInt
	total := 0

	for _, match := range matches {
		wLen := len(match[1])
		total += wLen

		_, ok := words[wLen]
		if !ok {
			words[wLen] = make([]string, 0)
		}

		words[wLen] = append(words[wLen], match[1])

		if wLen > longestLen {
			longestLen = wLen
		}

		if wLen < shortestLen {
			shortestLen = wLen
		}
	}

	avgF := float64(total) / float64(len(matches))
	avgUp := int(math.Ceil(avgF))
	avgDown := int(math.Floor(avgF))

	avgUps, _ := words[avgUp]
	avgDowns, _ := words[avgDown]

	return words[longestLen][0], words[shortestLen][0], avgF, append(avgUps, avgDowns...)
}