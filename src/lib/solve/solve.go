package solve

import (
	"Tucil1_13522011/lib/input"
)

type Coordinate struct {
	X int
	Y int
}

type Path []Coordinate

type SolutionData struct {
	Reward int
	Path   Path
}

func getReward(path Path, inputData input.InputData) int {
	// Function to calculate the reward from a pathforo

	// Initialize reward
	reward := 0

	// Iterate through all sequences
	for _, seq := range inputData.Sequences.Buffer {
		// Iterate through all possible position
		for i := 0; i < len(path)-len(seq.Sequence)+1; i++ {
			allMatch := true

			for j := 0; j < len(seq.Sequence); j++ {
				x := path[i+j].X
				y := path[i+j].Y
				// Check if the sequence match
				if inputData.Matrix.Buffer[y][x] != seq.Sequence[j] {
					allMatch = false
					break
				}
			}

			// If all match, add the reward
			if allMatch {
				reward += seq.Reward

				// Each sequence can only be used once
				break
			}
		}
	}

	return reward
}

func isCoordinateInPath(coordinate Coordinate, path Path) bool {
	// Function to check if a coordinate is in a path
	for _, p := range path {
		if p.X == coordinate.X && p.Y == coordinate.Y {
			return true
		}
	}

	return false
}

func GetOptimalSolution(inputData input.InputData) SolutionData {
	// Initialize solution
	solution := SolutionData{0, nil}

	// Iterate through all possible starting point
	for i := 0; i < inputData.Matrix.Width; i++ {
		newPath := getOptimalPathRecursive(Path{{i, 0}}, true, inputData)
		newReward := getReward(newPath, inputData)

		// Update solution if new reward is better
		if newReward > solution.Reward {
			solution.Reward = newReward
			solution.Path = newPath
		}
	}

	return solution
}

func getOptimalPathRecursive(currentPath Path, isCurrentVertical bool, inputData input.InputData) Path {
	// Base case
	if inputData.BufferSize <= inputData.Matrix.Width*inputData.Matrix.Height {
		if len(currentPath) == inputData.BufferSize {
			return nil
		}
	} else {
		if len(currentPath) == inputData.Matrix.Width*inputData.Matrix.Height {
			return nil
		}
	}

	var mostOptimalPath Path
	mostReward := 0

	// Generate path using recursive
	if isCurrentVertical {
		for i := 0; i < inputData.Matrix.Height; i++ {
			initialX := currentPath[len(currentPath)-1].X
			targetCoordinate := Coordinate{initialX, i}
			if !isCoordinateInPath(targetCoordinate, currentPath) {
				// Get new path
				newPath1 := make([]Coordinate, len(currentPath))
				copy(newPath1, currentPath)
				newPath1 = append(newPath1, targetCoordinate)

				// Recurrence
				newPath2 := getOptimalPathRecursive(newPath1, false, inputData)

				reward1 := getReward(newPath1, inputData)
				reward2 := getReward(newPath2, inputData)

				if reward1 >= reward2 && reward1 > mostReward {
					mostReward = reward1
					mostOptimalPath = newPath1
				}

				if reward2 > reward1 && reward2 > mostReward {
					mostReward = reward2
					mostOptimalPath = newPath2
				}
			}
		}
	} else {
		for i := 0; i < inputData.Matrix.Width; i++ {
			initialY := currentPath[len(currentPath)-1].Y
			targetCoordinate := Coordinate{i, initialY}
			if !isCoordinateInPath(targetCoordinate, currentPath) {
				// Get new path
				newPath1 := make([]Coordinate, len(currentPath))
				copy(newPath1, currentPath)
				newPath1 = append(newPath1, targetCoordinate)

				// Recurrence
				newPath2 := getOptimalPathRecursive(newPath1, true, inputData)

				reward1 := getReward(newPath1, inputData)
				reward2 := getReward(newPath2, inputData)

				if reward1 >= reward2 && reward1 > mostReward {
					mostReward = reward1
					mostOptimalPath = newPath1
				}

				if reward2 > reward1 && reward2 > mostReward {
					mostReward = reward2
					mostOptimalPath = newPath2
				}
			}
		}
	}

	return mostOptimalPath
}
