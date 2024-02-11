package main

import (
	"Tucil1_13522011/lib/initialize"
	"Tucil1_13522011/lib/input"
	"Tucil1_13522011/lib/output"
	"Tucil1_13522011/lib/solve"
	"time"
)

func main() {
	// Welcome message
	initialize.PrintWelcome()

	// Get Input Data
	var inputData input.InputData
	input.GetInputData(&inputData)

	// Solve
	timeStart := time.Now()
	optimalSolution := solve.GetOptimalSolution(inputData)
	timeEnd := time.Now()
	deltaTime := timeEnd.Sub(timeStart)

	// Print Result
	output.PrintResult(optimalSolution, deltaTime, inputData)
}
