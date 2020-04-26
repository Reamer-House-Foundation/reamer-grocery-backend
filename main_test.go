package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {

	if importJSONDataFromFile("data.json", &data) != true {
		t.Error("mock data is bad")
	}
}
