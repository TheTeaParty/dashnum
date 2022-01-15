package dashnum

import (
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	t.Fatalf("%v %v != %v", message, a, b)
}

func assertNotEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a != b {
		return
	}
	t.Fatalf("%v %v == %v", message, a, b)
}

func Test_averageNumber(t *testing.T) {
	avg1 := averageNumber("23-ab-48-caba-56-haha")
	expectAvg := (float64(23) + float64(48) + float64(56)) / float64(3)
	assertEqual(t, avg1, expectAvg, "Error Testing Success Case")

	avg2 := averageNumber("23-ab-48-caba-150-haha")
	assertNotEqual(t, avg2, expectAvg, "Error Testing Fail Case")
}

func Test_storyStats(t *testing.T) {
	testStr := "23-ab-48-longest-56-haha"
	avgs1 := (float64(len("ab")) +float64(len("longest")) +float64(len("haha"))) / float64(3)
	long1, short1, avg1, avgsR1 := storyStats(testStr)
	assertEqual(t, long1, "longest", "Longest Word Test Failed")
	assertEqual(t, short1, "ab", "Shortest Word Test Failed")
	assertEqual(t, avg1, avgs1, "Avg Word Length Test Failed")
	assertEqual(t, len(avgsR1), 1, "Avg Word Length Test Failed")
}

func Test_testValidity(t *testing.T) {
	valid1 := testValidity("23-ab-48-caba-150-haha")
	if !valid1 {
		t.Fatalf("Error Checking Success Validity")
	}
	valid2 := testValidity("23-55-ab-48-caba-150-haha")
	if valid2 {
		t.Fatalf("Error Checking Fail Validity")
	}
}
//
//func Test_wholeStory(t *testing.T) {
//	type args struct {
//		str string
//	}
//	tests := []struct {
//		name string
//		args args
//		want string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := wholeStory(tt.args.str); got != tt.want {
//				t.Errorf("wholeStory() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
