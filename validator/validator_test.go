package validator_test

import (
	"testing"

	"github.com/geofox-coding/first-go-app/validator"
)

func TestPackageTestsInCustomOrder(t *testing.T) {

	tests := []struct {
		testLabel string
		test func(*testing.T)	
	}{
		{"Dissimilarity Test", TestStringsAreDifferent},
		{"Contains Digit Test", TestStringContainsDigit},
		{"Simmilarity Test", TestStringsAreEqual},		
	}

	for _, test := range tests {
		t.Run(test.testLabel, test.test)
	}
}

func TestStringsAreEqual(t *testing.T) {

	firstInput := "mazedonien"
	secondInput := "mazedonien"

	actual := validator.StringsAreEqual(firstInput, secondInput)
	expected := true

	if actual != expected {
		t.Fatalf("%s doesn't match %s, expected a match", firstInput, secondInput)
	}
}

func TestStringsAreDifferent(t *testing.T) {

	firstInput := "mazedonien"
	secondInput := "albanien"

	actual := validator.StringsAreDifferent(firstInput, secondInput)
	expected := true

	if actual != expected {
		t.Fatalf("%s matches %s, expected don't match", firstInput, secondInput)
	}
}

func TestStringContainsDigit(t *testing.T) {

	firstInput := "maz3donien"

	actual := validator.StringContainsDigit(firstInput)
	expected := true

	if actual != expected {
		t.Fatalf("%s doesn't contain digit, expected contains", firstInput)
	}
}
