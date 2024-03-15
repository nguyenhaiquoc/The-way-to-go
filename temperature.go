package main

type Celsius float32
type Fahrenheit float32

func CelsiusToFahrenheit(celsius Celsius) Fahrenheit {
	return Fahrenheit(celsius*9.0/5.0 + 32.0)
}

func FahrenheitToCelsius(fahrenheit Fahrenheit) Celsius {
	return Celsius((fahrenheit - 32.0) * 5.0 / 9.0)
}
