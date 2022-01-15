package dashnum

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Task 1: Difficulty - 1/10; Estimate - 30 minutes; Actual - 27 minutes
// Task 2: Difficulty - 1/10; Estimate - 10 minutes; Actual - 6 minutes

// generate takes boolean flag and generates random correct strings if the
// parameter is `true` and random incorrect strings if the flag is `false`.
func generate(valid bool) string {

	n := rand.Intn(100)
	randWordL := rand.Intn(32-1) + 1
	randN := rand.Intn(100-1) + 1

	var str string

	if !valid {
		str = fmt.Sprintf("%v-%v", randString(randWordL), randN)
	} else {
		str = fmt.Sprintf("%v-%v", randN, randString(randWordL))
	}

	for i := 0; i < n; i++ {
		randWordL := rand.Intn(32-1) + 1
		randN := rand.Intn(100-1) + 1

		str = fmt.Sprintf("%v-%v-%v", str, randN, randString(randWordL))
	}

	if !valid {
		str = fmt.Sprintf("%v", str)
	}

	return str
}

// randString generates random string of the given `n` length
func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// testValidity takes the string as an input, and returns boolean flag `true` if the given
// string complies with the format, or `false` if the string does not comply
func testValidity(str string) bool {
	var re = regexp.MustCompile(`^([0-9]+-[a-zA-Z]+-?)+$`)
	return re.MatchString(str)
}

// averageNumber takes the string, and returns the average number from all the numbers
func averageNumber(str string) float64 {
	if !testValidity(str) {
		return 0
	}

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
	if !testValidity(str) {
		return ""
	}

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
	if !testValidity(str) {
		return "", "", 0, []string{}
	}

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
