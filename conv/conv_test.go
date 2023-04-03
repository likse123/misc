package conv

import (
	"math"
	"testing"
	//"reflect"
)

/*
*
	Mal for testfunksjoner
	Du skal skrive alle funksjonene basert på denne malen
	For alle konverteringsfunksjonene (tilsammen 6)
	kan du bruke malen som den er (du må selvsagt endre
	funksjonsnavn og testverdier)
*/
func TestFahrenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 32, want: 0},
	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
			//if !reflect.DeepEqual(tc.want, got) {
			//	t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// De andre testfunksjonene implementeres her
// ...
func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: -89.4, want: -128.92},
	}

	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
			//if !reflect.DeepEqual(tc.want, got) {
			//	t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: -89.4, want: 183.75},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
			//if !reflect.DeepEqual(tc.want, got) {
			//	t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 273.15, want: 0},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
			//if !reflect.DeepEqual(tc.want, got) {
			//	t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestFahrenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 32, want: 273.15},
	}

	for _, tc := range tests {
		got := FahrenheitToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
			//if !reflect.DeepEqual(tc.want, got) {
			//	t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 183.75, want: -128.92},
	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
			//if !reflect.DeepEqual(tc.want, got) {
			//	t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func withinTolerance(a, b, error float64) bool {
	// Først sjekk om tallene er nøyaktig like
	if a == b {
		return true
	}

	difference := math.Abs(a - b)

	// Siden vi skal dele med b, må vi sjekke om den er 0
	// Hvis b er 0, returner avgjørelsen om d er mindre enn feilmarginen
	// som vi aksepterer
	if b == 0 {
		return difference < error
	}

	// Tilslutt sjekk den relative differanse mot feilmargin
	return (difference / math.Abs(b)) < error
}
