package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FahrenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFahrenheit
    ...
*/
// Konverterer Fahrenheit til Celsius
func FahrenheitToCelsius(value float64) float64 {
	// Her skal du implementere funksjonen
	// Celsius = (Fahrenheit -32) * (5/9)
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	//return 56.67
	return (value - 32.0) * 5 / 9
}

// De andre konverteringsfunksjonene implementere her
// ...
func CelsiusToFahrenheit(value float64) float64 {
	return value*9/5 + 32
}

func KelvinToCelsius(value float64) float64 {
	return value - 273.15
}

func CelsiusToKelvin(value float64) float64 {
	return value + 273.15
}

func FahrenheitToKelvin(value float64) float64 {
	return (value-32)*5/9 + 273.15
}

func KelvinToFahrenheit(value float64) float64 {
	return (value-273.15)*9/5 + 32
}
