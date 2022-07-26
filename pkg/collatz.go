package pkg

import (
	"fmt"
	"github.com/marcellribeiro/tunity_test/models"
	"math"
)

func even(n float64) (b bool) {
	b = false
	if math.Mod(n, 2) == 0 {
		b = true
	}
	return b
}

func odd(n float64) (b bool) {
	b = true
	if even(n) {
		b = false
	}
	return b
}

func collatz(collatzData []models.Collatz, lastNumber float64) []models.Collatz {
	lastStep := (collatzData)[len(collatzData)-1].Step
	nextStep := lastStep + 1

	switch {
	case lastNumber == 1:
		return collatzData
	case even(lastNumber):
		lastNumber = lastNumber / 2
		collatzData = appendCollatzToList(nextStep, lastNumber, collatzData)
		return collatz(collatzData, lastNumber)
	case odd(lastNumber):
		lastNumber = lastNumber*3 + 1
		collatzData = appendCollatzToList(nextStep, lastNumber, collatzData)
		return collatz(collatzData, lastNumber)
	}
	return collatzData
}

func appendCollatzToList(nextStep int, lastNumber float64, collatzData []models.Collatz) []models.Collatz {
	var collatzRow models.Collatz
	collatzRow.Step = nextStep
	collatzRow.Result = lastNumber
	collatzData = append(collatzData, collatzRow)

	return collatzData
}

func GetCollatz(rootNumber float64) ([]models.Collatz, error) {
	var collatzData []models.Collatz
	var collatzRow = models.Collatz{Result: rootNumber, Step: 1}

	if rootNumber <= 0 {
		return collatzData, fmt.Errorf("given number must be greater than 0")
	}

	collatzData = append(collatzData, collatzRow)

	result := collatz(collatzData, rootNumber)

	return result, nil
}
