package main

import (
	"testing"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	celsius := Celsius(25.0)
	expected := Fahrenheit(77.0)

	fahrenheit := CelsiusToFahrenheit(celsius)

	if fahrenheit != expected {
		t.Errorf("CelsiusToFahrenheit(%f) = %f; want %f", celsius, fahrenheit, expected)
	}
}

func TestFahrenheitToCelsius(t *testing.T) {
	fahrenheit := Fahrenheit(77.0)
	expected := Celsius(25.0)

	celsius := FahrenheitToCelsius(fahrenheit)

	if celsius != expected {
		t.Errorf("FahrenheitToCelsius(%f) = %f; want %f", fahrenheit, celsius, expected)
	}
}
