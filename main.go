package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//go run . [OPTION] [STRING] [BANNER].
	//go run . --output=<fileName.txt> something standard

	InputFile, OutputFile, text := HandleTheTerminalCommande()

	if text == "" {
		log.Fatalln("You didn't Enter any word to print !!!")
	}
	slice := FixTheInputFormat(InputFile)

	slicedtext := strings.Split(text, "\\n")

	result := DrawAscii(slice, slicedtext)

	if IsItAllNewLines(result) {
		result = result[1:]
	}

	CreateOutputFile(result, OutputFile)

}

func HandleTheTerminalCommande() (string, string, string) {
	if len(os.Args[1:]) > 3 {
		log.Fatalln("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
	}
	var InputFile string
	var OutputFile string
	var text string

	if len(os.Args[1:]) == 3 {
		count := 3
		for _, arg := range os.Args[1:] {
			if arg == "standard" || arg == "standard.txt" || arg == "shadow" || arg == "shadow.txt" || arg == "thinkertoy" || arg == "thinkertoy.txt" {
				if arg == "standard" || arg == "shadow" || arg == "thinkertoy" {
					InputFile = arg + ".txt"
					count--

				} else {
					InputFile = arg
					count--

				}
			} else if strings.Contains(arg, "--output=") {
				OutputFile = arg[9:]
				count++

			} else {
				text = arg
			}
		}
	} else if len(os.Args[1:]) == 2 {
		count := 2
		for _, arg := range os.Args[1:] {
			if arg == "standard" || arg == "standard.txt" || arg == "shadow" || arg == "shadow.txt" || arg == "thinkertoy" || arg == "thinkertoy.txt" {
				if arg == "standard" || arg == "shadow" || arg == "thinkertoy" {
					InputFile = arg + ".txt"
					count--
				} else {
					InputFile = arg
					count--
				}
			} else if strings.Contains(arg, "--output=") {
				OutputFile = arg[9:]
				count--
			} else {
				text = arg
			}

		}
		if InputFile == "" {
			InputFile = "standard.txt"
		} else if OutputFile == "" {
			OutputFile = "banner.txt"
		}
		fmt.Println(text, InputFile, OutputFile)
	} else if len(os.Args[1:]) == 1 {
		InputFile = "standard.txt"
		OutputFile = "banner.txt"
		text = os.Args[1]
	}

	return InputFile, OutputFile, text
}

func FixTheInputFormat(InputFile string) []string {
	var sep string
	if InputFile == "standard.txt" || InputFile == "shadow.txt" {
		sep = "\n"
	} else {
		sep = "\r\n"
	}
	data, err := os.ReadFile(InputFile)
	if err != nil {
		log.Fatalln("There's q problem with the input file", InputFile)
	}
	slice := RemoveEmptyStrings(strings.Split(string(data), sep))
	return slice
}

func RemoveEmptyStrings(slice []string) []string {
	var temp []string
	for i := range slice {
		if slice[i] != "" {
			temp = append(temp, slice[i])
		}
	}
	return temp
}

func DrawAscii(slice, slicedtext []string) string {
	var result string

	for _, word := range slicedtext {
		if word != "" {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					if char < 32 || char > 126 {
						log.Fatalln("You entered an inprintabale caharcter !!!!")
					} else {
						start := int(char-32)*8 + i
						result += slice[start]
					}
				}
				result += "\n"
			}
		} else {
			result += "\n"
		}
	}
	return result
}

func IsItAllNewLines(result string) bool {
	for _, char := range result {
		if char != '\n' {
			return false
		}
	}
	return true
}

func CreateOutputFile(result, OutputFile string) {
	file, err := os.Create(OutputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	FinalResult := []byte(result)

	err = os.WriteFile(OutputFile, FinalResult, 0644)

	if err != nil {
		fmt.Println("Error: ", err)
	}
}
