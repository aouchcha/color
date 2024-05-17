package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	InputFile, Paint, lettertocollored, text, OutputFile := HandleTheTerminalCommands()
	
	fmt.Printf("%#v,%#v,%#v,%#v,%#v", InputFile, Paint, lettertocollored, text, OutputFile)


	var result string
	var temp []rune

	if Paint != "" {
		if lettertocollored == "" {
			lettertocollored = text
		}
		Colors := []rune(Paint)
		if Colors[0] >= 'a' && Colors[0] <= 'z' {
			Colors[0] -= 32
		}
		fmt.Println(string(Colors))
		Paint = string(Colors)
		switch Paint {
		case "Red":
			Paint = "\033[31m"
		case "Green":
			Paint = "\033[32m"
		case "Yellow":
			Paint = "\033[33m"
		case "Blue": 
			Paint = "\033[34m"
		case "Magenta": 
			Paint = "\033[35m"
		case "Cyan":
			Paint = "\033[36m"
		case "Gray":
			Paint = "\033[37m"
		case "White":
			Paint = "\033[97m"
		default:
			log.Fatalln("Unknown color:", Paint)
		}
	}

	Reset := "\033[0m"

	sli := []rune(lettertocollored)
	for i := 0; i < len(sli); i++ {
		for j := i + 1; j < len(sli); j++ {
			if sli[i] == sli[j] {
				sli[i] = '0'
			}
		}
	}

	for i := 0; i < len(sli); i++ {
		if sli[i] != '0' {
			temp = append(temp, sli[i])
		}
	}
	fmt.Println(len(temp))

	data, err := os.ReadFile(InputFile)
	if err != nil {
		log.Fatalln("There is a problem with the reading file")
	}

	var sep string
	if InputFile == "standard.txt" || InputFile == "shadow.txt" {
		sep = "\n"
	} else {
		sep = "\r\n"
	}

	slice := RemoveEmptyStrings(strings.Split(string(data[1:]), sep))

	slicedArgs := strings.Split(text, "\\n")

	var start int

	for _, word := range slicedArgs {
		if word != "" {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					if char < 32 || char > 126 {
						log.Fatalln("You entered an inprintabale character !!!")
					} else {
						if len(temp) != 0 {
							var IsItDone bool
						for k := 0; k < len(temp); k++ {

							start = int(char-32)*8 + i
							if char == temp[k] {
								result += fmt.Sprint(Paint, slice[start], Reset)
								IsItDone = true

							}

						}
						if !IsItDone {
							result += (slice[start])
						}
						}else{
							start = int(char-32)*8 + i
							
								result += slice[start]
								

							}
						}

					}
					result += "\n"
				}

			} else {
			result += "\n"
		}
	}

	if OutputFile != "" {
		file, err := os.Create(OutputFile)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()

		FinalResult := []byte(result)
		

		err = os.WriteFile(OutputFile, FinalResult, 0o644)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Println("Go check the file that you creat !!!!!")
	} else {
		fmt.Print(result)
	}
}

func HandleTheTerminalCommands() (string, string, string, string, string) {
	if len(os.Args[1:]) > 5 {
		log.Fatalln("Usage: go run . [OPTION] [STRING]\n", "EX: go run . --color=<color> <letters to be colored> ", "something")
	}
	var InputFile string
	var Color string
	var text string
	var lettertocollored string
	var OutputFile string
	slice := os.Args[1:]
	TwoFlags := 0
	
		for i := 0; i < len(slice); i++ {
			if slice[i] == "standard" || slice[i] == "standard.txt" || slice[i] == "shadow" || slice[i] == "shadow.txt" || slice[i] == "thinkertoy" || slice[i] == "thinkertoy.txt" {
				if slice[i] == "standard" || slice[i] == "shadow" || slice[i] == "thinkertoy" {
					InputFile = slice[i] + ".txt"
				} else {
					InputFile = slice[i]
				}
			} else if strings.Contains(slice[i], "--color=") {
				Color = slice[i][8:]
				TwoFlags++

				if i < len(slice)-1 {
					if slice[i+1] == "standard.txt" || slice[i+1] == "shadow.txt" || slice[i+1] == "thinkertoy.txt" {
						lettertocollored = text
					} else {
						lettertocollored = slice[i+1]
					}
				}

			} else if strings.Contains(slice[i], "--output=") {
				OutputFile = slice[i][9:]
				TwoFlags++
			} else {
				text = slice[i]
			}
		
		if InputFile == "" {
			InputFile = "standard.txt"
		}else if len(slice) == 1 {
		text = slice[0]
		InputFile = "standard.txt"
	}

	if TwoFlags == 2 {
		log.Fatalln("You entered two flags and it's not allowed !!!")
	}
		}
	return InputFile, Color, lettertocollored, text, OutputFile
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
