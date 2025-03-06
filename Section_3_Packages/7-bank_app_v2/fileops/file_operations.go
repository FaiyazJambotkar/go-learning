package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetValueFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValue, errors.New("file not found")
	}

	outputText := string(data)
	value, err := strconv.ParseFloat(outputText, 64)

	if err != nil {
		return 10000, errors.New("file not found")
	}

	return value, nil
}

func WriteValueToFile(value float64) {
	outputText := fmt.Sprint(value)
	os.WriteFile("balance.txt", []byte(outputText), 0644)
}