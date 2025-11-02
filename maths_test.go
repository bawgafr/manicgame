package main

import "testing"

var level = NewLevel(1600, 600)

func TestXCoorMap(t *testing.T) {

	values := []float64{0.0, 200.0, 800.0, 1600.0}
	expected := []float64{0.0, 100.0, 400.0, 800.0}

	for i, v := range values {
		mapped := XCoordMap(v, level)
		if mapped != expected[i] {
			t.Errorf("For input %f, expected %f but got %f", v, expected[i], mapped)
		}
	}
	// mappedX := XCoordMap(400, level)
	// if mappedX != 800.0 {
	// 	t.Errorf("Expected mappedX to be 800.0, got %f", mappedX)
	// }

	// mappedX = XCoordMap(800, level)
	// if mappedX != 1600.0 {
	// 	t.Errorf("Expected mappedX to be 1600.0, got %f", mappedX)
	// }
}

func TestYCoordMap(t *testing.T) {

	values := []float64{0.0, 300.0, 600.0}
	expected := []float64{50.0, 275.0, 500.0}

	for i, v := range values {
		y := YCoordMap(v, level)
		if y != expected[i] {
			t.Errorf("For input %f, expected %f but got %f", v, expected[i], y)
		}
	}

}
