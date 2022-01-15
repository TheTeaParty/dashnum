package main

import (
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func Test_averageNumber(t *testing.T) {
	avg1 := averageNumber("23-ab-48-caba-56-haha")
	expectAvg := (float64(23) + float64(48) + float64(56)) / float64(3)
	assertEqual(t, avg1, expectAvg, "%v != %v")

	avg2 := averageNumber("23-ab-48-caba-150-haha")
	assertEqual(t, avg2, expectAvg, "%v != %v")
}

func Test_storyStats(t *testing.T) {
	testStr := "23-ab-48-longest-56-haha"
	avgs1 := float64(len("longest")) +
	long1, short1, avg1, avgStrs := storyStats(testStr)
	assertEqual(t, long1, len("longest"), "Longest Word Test Failed: %v != %v")
	assertEqual(t, short1, len("ab"), "Shortest Word Test Failed: %v != %v")
	assertEqual(t, avg1, len("ab"), "Avg Word Length Test Failed: %v != %v")
}

func Test_testValidity(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testValidity(tt.args.str); got != tt.want {
				t.Errorf("testValidity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wholeStory(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wholeStory(tt.args.str); got != tt.want {
				t.Errorf("wholeStory() = %v, want %v", got, tt.want)
			}
		})
	}
}
