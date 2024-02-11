package input

import (
	"Tucil1_13522011/lib/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Matrix struct {
	Width  int
	Height int
	Buffer [][]string
}

type Sequence struct {
	Sequence []string
	Reward   int
}

type ArraySequence struct {
	Buffer []Sequence
	Size   int
}

type InputData struct {
	BufferSize int
	Matrix     Matrix
	Sequences  ArraySequence
}

func printInputData(inputData InputData) {
	// Function to print input data
	fmt.Println("===================================================================================")
	fmt.Println("Hasil input data yang dilakukan:")
	fmt.Println("Buffer size:", inputData.BufferSize)
	fmt.Println("Matrix width:", inputData.Matrix.Width)
	fmt.Println("Matrix height:", inputData.Matrix.Height)
	fmt.Println("Matrix data:")
	for i := 0; i < inputData.Matrix.Height; i++ {
		for j := 0; j < inputData.Matrix.Width; j++ {
			fmt.Print(inputData.Matrix.Buffer[i][j])
			if j != inputData.Matrix.Width-1 {
				fmt.Print(" ")
			} else {
				fmt.Println()
			}
		}
	}
	fmt.Println("Sequence size:", inputData.Sequences.Size)
	fmt.Println("Sequence data:")
	for i := 0; i < inputData.Sequences.Size; i++ {
		for j := 0; j < len(inputData.Sequences.Buffer[i].Sequence); j++ {
			fmt.Print(inputData.Sequences.Buffer[i].Sequence[j])
			if j != len(inputData.Sequences.Buffer[i].Sequence)-1 {
				fmt.Print(" ")
			} else {
				fmt.Println()
			}
		}
		fmt.Println(inputData.Sequences.Buffer[i].Reward)
	}
	fmt.Println("===================================================================================")
}

func getInputMethod(inputMethod *string) {
	fmt.Println("===================================================================================")
	fmt.Println("Pilih metode input:")
	fmt.Println("1. File")
	fmt.Println("2. Keyboard")

	fmt.Scan(inputMethod)
	for *inputMethod != "1" && *inputMethod != "2" {
		fmt.Println("Pilihan tidak valid, silakan masukkan pilihan yang valid")
		fmt.Scan(inputMethod)
	}
	fmt.Println("===================================================================================")
}

func getFromFile(inputData *InputData) {
	fmt.Println("===================================================================================")
	fmt.Println("Masukkan nama file (di folder /test/input beserta ekstensi txt):")
	var filename string
	// Input filename
	fmt.Scan(&filename)

	// Check if file doesnt exists
	for {
		if _, err := os.Stat("../test/input/" + filename); os.IsNotExist(err) {
			fmt.Println("File tidak ditemukan, silakan masukkan nama file yang valid (di folder /test/input beserta ekstensi txt):")
			fmt.Scan(&filename)
		} else {
			break
		}
	}
	fmt.Println("===================================================================================")

	// Open file
	file, err := os.Open("../test/input/" + filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Read file & parse input
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		if i == 0 {
			// Buffer size
			_, err := fmt.Sscanf(scanner.Text(), "%d", &inputData.BufferSize)
			if err != nil {
				fmt.Println("Error: buffer size bukan angka yang valid")
				os.Exit(1)
			}
		} else if i == 1 {
			// Matrix width and height
			_, err := fmt.Sscanf(scanner.Text(), "%d %d", &inputData.Matrix.Width, &inputData.Matrix.Height)
			if err != nil {
				fmt.Println("Error: width atau height matriks bukan angka yang valid")
				os.Exit(1)
			}
		} else if i >= 2 && i < 2+inputData.Matrix.Height {
			// Matrix data
			rowString := scanner.Text()

			// Trim beginning and trailing spaces
			rowString = strings.TrimSpace(rowString)

			// Check if row length is not equal to matrix width
			rows := strings.Split(rowString, " ")
			if len(rows) != inputData.Matrix.Width {
				fmt.Println("Error: input matriks tidak sesuai dengan input panjang width matriks")
				os.Exit(1)
			}
			inputData.Matrix.Buffer = append(inputData.Matrix.Buffer, rows)
		} else if i == 2+inputData.Matrix.Height {
			// Sequence size
			_, err := fmt.Sscanf(scanner.Text(), "%d", &inputData.Sequences.Size)
			if err != nil {
				fmt.Println("Error: sequence size bukan angka yang valid")
				os.Exit(1)
			}
		} else if i > 2+inputData.Matrix.Height && i <= 2+inputData.Matrix.Height+inputData.Sequences.Size*2 {
			if (i-2-inputData.Matrix.Height)%2 == 1 {
				// Sequence data
				rowString := scanner.Text()

				// Trim beginning and trailing spaces
				rowString = strings.TrimSpace(rowString)

				sequence := strings.Split(rowString, " ")
				inputData.Sequences.Buffer = append(inputData.Sequences.Buffer, Sequence{sequence, 0})
			} else {
				// Reward data
				_, err := fmt.Sscanf(scanner.Text(), "%d", &inputData.Sequences.Buffer[(i-3-inputData.Matrix.Height)/2].Reward)
				if err != nil {
					fmt.Println("Error: reward bukan angka yang valid")
					os.Exit(1)
				}
			}
		}
	}
}

func getFromKeyboard(inputData *InputData) {
	fmt.Println("===================================================================================")
	fmt.Println("Masukkan jumlah token unik:")
	var tokenUnikCount int
	fmt.Scan(&tokenUnikCount)

	fmt.Println("Masukkan token unik:")
	tokenArray := make([]string, tokenUnikCount)
	for i := 0; i < tokenUnikCount; i++ {
		fmt.Scan(&tokenArray[i])
	}

	fmt.Println("Masukkan buffer size:")
	fmt.Scan(&inputData.BufferSize)

	fmt.Println("Masukkan matriks width:")
	fmt.Scan(&inputData.Matrix.Width)

	fmt.Println("Masukkan matriks height:")
	fmt.Scan(&inputData.Matrix.Height)

	fmt.Println("Masukkan jumlah sequence:")
	fmt.Scan(&inputData.Sequences.Size)

	var ukuranMaxSequence int
	fmt.Println("Masukkan ukuran max sequence:")
	fmt.Scan(&ukuranMaxSequence)

	// Generate random matrix
	for i := 0; i < inputData.Matrix.Height; i++ {
		// Generate row
		row := make([]string, inputData.Matrix.Width)
		for j := 0; j < inputData.Matrix.Width; j++ {
			randomTokenIndex := utils.GetRandom(0, tokenUnikCount-1)
			row[j] = tokenArray[randomTokenIndex]
		}
		inputData.Matrix.Buffer = append(inputData.Matrix.Buffer, row)
	}

	// Generate random sequence
	for i := 0; i < inputData.Sequences.Size; i++ {
		// Generate sequence
		// Sequence minimal 2 token (dari spec)
		sequenceLength := utils.GetRandom(2, ukuranMaxSequence)
		sequence := make([]string, sequenceLength)
		for j := 0; j < sequenceLength; j++ {
			randomTokenIndex := utils.GetRandom(0, tokenUnikCount-1)
			sequence[j] = tokenArray[randomTokenIndex]
		}
		// Generate sequence reward (diatur min = 0 dan max = 50)
		randomReward := utils.GetRandom(0, 50)
		inputData.Sequences.Buffer = append(inputData.Sequences.Buffer, Sequence{sequence, randomReward})
	}
	fmt.Println("===================================================================================")
}

func GetInputData(inputData *InputData) {
	var inputMethod string
	getInputMethod(&inputMethod)

	if inputMethod == "1" {
		getFromFile(inputData)
	} else if inputMethod == "2" {
		getFromKeyboard(inputData)
	}

	printInputData(*inputData)
}
