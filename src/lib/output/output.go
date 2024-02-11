package output

import (
	"Tucil1_13522011/lib/input"
	"Tucil1_13522011/lib/solve"
	"fmt"
	"os"
	"time"
)

func getPathCode(path solve.Path, inputData input.InputData) string {
	// Function to get the path code
	var pathCode string
	for i, p := range path {
		pathCode += string(inputData.Matrix.Buffer[p.Y][p.X])
		if i != len(path)-1 {
			pathCode += " "
		}
	}
	return pathCode
}

func saveToFile(solution solve.SolutionData, deltaTime time.Duration, inputData input.InputData) {
	// Function to save solution to file
	// Get file name
	fmt.Println("Masukkan nama file untuk menyimpan solusi (di folder /test/output beserta ekstensi txt):")
	var fileName string
	fmt.Scan(&fileName)

	// Check if file already exists
	for {
		_, err := os.Stat("../test/output/" + fileName)
		if os.IsNotExist(err) {
			break
		}
		fmt.Println("File sudah ada, apakah ingin menimpanya? (y/n)")

		var overwrite string
		fmt.Scan(&overwrite)
		for overwrite != "y" && overwrite != "n" {
			fmt.Println("Input tidak valid")
			fmt.Scan(&overwrite)
		}

		if overwrite == "y" {
			break
		}

		fmt.Println("Masukkan nama file untuk menyimpan solusi (di folder /test/output beserta ekstensi txt):")
		fmt.Scan(&fileName)
	}

	// Create file
	file, err := os.Create("../test/output/" + fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Write to file
	// Reward
	file.WriteString(fmt.Sprintf("Optimal Reward: %d\n", solution.Reward))
	// Path
	pathCode := getPathCode(solution.Path, inputData)
	file.WriteString(fmt.Sprintf("Path: %s\n", pathCode))
	// Coordinate
	file.WriteString("Coordinate:\n")
	for _, p := range solution.Path {
		file.WriteString(fmt.Sprintf("%d %d\n", p.X+1, p.Y+1))
	}
	// Delta Time
	file.WriteString("\n")
	file.WriteString(fmt.Sprintf("Time: %d ms\n", deltaTime.Milliseconds()))
}

func PrintResult(solution solve.SolutionData, deltaTime time.Duration, inputData input.InputData) {
	fmt.Println("===================================================================================")
	if solution.Reward == 0 {
		// Reward
		fmt.Println("Optimal Reward: 0")
		fmt.Println("No solution found")
	} else {
		// Reward
		fmt.Println("Optimal Reward:", solution.Reward)
		// Path
		pathCode := getPathCode(solution.Path, inputData)
		fmt.Printf("Path: %s\n", pathCode)
		// Coordinate
		fmt.Println("Coordinate:")
		for _, p := range solution.Path {
			fmt.Printf("%d %d\n", p.X+1, p.Y+1)
		}
	}

	// Delta Time
	fmt.Println()
	fmt.Println("Time:", deltaTime.Milliseconds(), "ms")
	fmt.Println()

	// Ingin menyimpan solusi
	fmt.Println("Apakah ingin menyimpan solusi? (y/n)")
	var saveSolution string
	fmt.Scan(&saveSolution)
	for saveSolution != "y" && saveSolution != "n" {
		fmt.Println("Input simpan solusi tidak valid")
		fmt.Scan(&saveSolution)
	}

	if saveSolution == "y" {
		saveToFile(solution, deltaTime, inputData)
	}

	fmt.Println("===================================================================================")
	fmt.Println("===================================================================================")
}
