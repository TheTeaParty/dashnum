package main

import "regexp"

// Task 1: Difficulty - 1/10; Estimate - 30 minutes

func main() {

}

// testValidity` takes the string as an input, and returns boolean flag `true` if the given
// string complies with the format, or `false` if the string does not comply
func testValidity(str string) bool {
	var re = regexp.MustCompile(`^([0-9]+-[a-zA-Z]+-?)+$`)
	return re.MatchString(str)
}