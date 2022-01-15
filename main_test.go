package main

import (
	"reflect"
	"testing"
)

func Test_averageNumber(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageNumber(tt.args.str); got != tt.want {
				t.Errorf("averageNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_storyStats(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
		want2 float64
		want3 []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := storyStats(tt.args.str)
			if got != tt.want {
				t.Errorf("storyStats() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("storyStats() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("storyStats() got2 = %v, want %v", got2, tt.want2)
			}
			if !reflect.DeepEqual(got3, tt.want3) {
				t.Errorf("storyStats() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
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
