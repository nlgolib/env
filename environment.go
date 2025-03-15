package env

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
)

type Environment struct {
	Key string `json:"key"`
}

// Must checks if the environment variable is set.
// If it is not set, it will log a fatal error and exit the program.
func (e *Environment) Must() {
	value := os.Getenv(e.Key)
	if value == "" {
		log.Fatalf("Environment variable %s is not set", e.Key)
	}
}

// String returns the environment variable as a string.
func (e *Environment) String() string {
	return os.Getenv(e.Key)
}

// Bool returns the environment variable as a boolean.
func (e *Environment) Bool() bool {
	if e.String() == "" {
		return false
	}
	value, err := strconv.ParseBool(e.String())
	if err != nil {
		log.Fatalf("Environment variable %s is not a boolean", e.Key)
	}
	return value
}

// Int returns the environment variable as an integer.
func (e *Environment) Int() int {
	if e.String() == "" {
		return 0
	}
	value, err := strconv.Atoi(e.String())
	if err != nil {
		log.Fatalf("Environment variable %s is not an integer", e.Key)
	}
	return value
}

// Float64 returns the environment variable as a float64.
func (e *Environment) Float64() float64 {
	if e.String() == "" {
		return 0
	}
	value, err := strconv.ParseFloat(e.String(), 64)
	if err != nil {
		log.Fatalf("Environment variable %s is not a float64", e.Key)
	}
	return value
}

// JSON returns the environment variable as a JSON object.
func (e *Environment) JSON() any {
	value := e.String()
	if value == "" {
		return nil
	}
	var result any
	err := json.Unmarshal([]byte(value), &result)
	if err != nil {
		log.Fatalf("Environment variable %s is not a valid JSON", e.Key)
	}
	return result
}

// StringSlice returns the environment variable as a slice of strings.
func (e *Environment) StringSlice() []string {
	value := e.String()
	if value == "" {
		return []string{}
	}
	value = strings.ReplaceAll(value, " ", "")
	return strings.Split(value, ",")
}

// IntSlice returns the environment variable as a slice of integers.
func (e *Environment) IntSlice() []int {
	values := e.StringSlice()
	if len(values) == 0 {
		return []int{}
	}
	result := make([]int, len(values))
	for i, value := range values {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("Environment variable %s is not an integer", e.Key)
		}
		result[i] = intValue
	}
	return result
}

// Float64Slice returns the environment variable as a slice of float64.
func (e *Environment) Float64Slice() []float64 {
	values := e.StringSlice()
	if len(values) == 0 {
		return []float64{}
	}
	result := make([]float64, len(values))
	for i, value := range values {
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatalf("Environment variable %s is not a float64", e.Key)
		}
		result[i] = floatValue
	}
	return result
}
