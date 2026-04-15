package ex2

import (
	"errors"
)

type UnsupportedUnitError struct {
	Name string
}

func (ue *UnsupportedUnitError) Error() string {
	return "unsupported unit: " + ue.Name
}

func (ue *UnsupportedUnitError) Is(err error) bool {
	if e, ok := errors.AsType[*UnsupportedUnitError](err); ok {
		return e.Name == ue.Name
	}
	return false
}

// Convert converts a temperature value from one unit to another.
// Supported units: "C" (Celsius), "F" (Fahrenheit), "K" (Kelvin).
func Convert(value float64, from string, to string) (float64, error) {
	if from == to {
		return value, nil
	}

	// Normalize to Celsius first
	var celsius float64
	switch from {
	case "C":
		celsius = value
	case "F":
		celsius = (value - 32.0) * 5.0 / 9.0
	case "K":
		celsius = value - 273.15
	default:
		return 0, &UnsupportedUnitError{Name: from}
	}

	// Convert from Celsius to target
	switch to {
	case "C":
		return celsius, nil
	case "F":
		return celsius*9.0/5.0 + 32.0, nil
	case "K":
		return celsius + 273.15, nil
	default:
		return 0, &UnsupportedUnitError{Name: to}
	}
}
