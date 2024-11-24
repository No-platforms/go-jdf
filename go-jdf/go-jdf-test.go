package go-jdf

import (
	"testing"
)

func TestConversions(t *testing.T) {
	jalaliDate := JalaliDate{Year:1402 , Month:9 , Day:25}
	
	gregorianExpected := GregorianDate{Year:2023 , Month:11 , Day:15} // Replace with actual expected value based on your calculations.

	gregorianResult , err := ToGregorian(jalaliDate)
	if err != nil || gregorianResult != gregorianExpected {
		t.Fatalf("Expected %v but got %v; error: %v", gregorianExpected , gregorianResult , err)
	}

	jalaliResult , err := ToJalali(gregorianResult)
	if err != nil || jalaliResult != jalaliDate {
		t.Fatalf("Expected %v but got %v; error: %v", jalaliDate , jalaliResult , err)
	}
}