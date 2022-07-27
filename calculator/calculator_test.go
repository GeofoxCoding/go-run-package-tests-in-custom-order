package calculator_test

import (
	"testing"

	"github.com/geofox-coding/first-go-app/calculator"
)

func TestSum(t *testing.T) {
	actual := calculator.Sum(5, 10)
	expected := 15

	if actual != expected {
		t.Fatalf("Wrong sum got %d, wanted %d", actual, expected)
	}
}