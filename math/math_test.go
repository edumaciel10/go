package math

import (
	"testing"
)

func TestPi(t *testing.T) {
	result := Pi()
	expected := 3.141592653589793
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestSqrt(t *testing.T) {
	result := Sqrt(4);

	expected := 2.0

	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}
