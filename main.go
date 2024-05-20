package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	InputFile, Paint, lettertocollored, text, OutputFile := HandleTheTerminalCommands()
	if len(os.Args[1:]) == 0 {
		return
	}
	Paint = ChoseTheColor(Paint)
	slice, slicedArgs := Format(InputFile, text)
	if len(slice) != 760 {
		log.Fatalln("You don't have the right to modify in the template file", InputFile)
	}

	result := DrawAsciiArtColor(slice, slicedArgs, Paint, lettertocollored)
	if OutputFile != "NILL" {
		if OutputFile == "" {
			OutputFile = "Banner.txt"
		} else {
			OutputFile += ".txt"
		}

		file, err := os.Create(OutputFile)
		if err != nil {
			log.Fatalln("Error :", err)
		}
		defer file.Close()

		err = os.WriteFile(OutputFile, []byte(result), 0644)
		if err != nil {
			log.Fatalln("Error :", err)
		}

	} else {
		fmt.Print(result)
	}

}

func HandleTheTerminalCommands() (string, string, string, string, string) {
	var InputFile string
	var Color string
	var text string
	var lettertocollored string
	var OutputFile string

	slice := os.Args[1:]
	TwoFlags := 0
	//the case when we have a lot of arguments or the flag isn't correct
	if !CheckTheFlags(slice) || len(os.Args[1:]) > 4 {
		log.Fatalln("Usage: go run . [OPTION] [STRING]\n", "EX: go run . --color=<color> <letters to be colored> ", "something")
	} else {
		for i := 0; i < len(slice); i++ {
			if slice[i] == "standard" || slice[i] == "standard.txt" || slice[i] == "shadow" || slice[i] == "shadow.txt" || slice[i] == "thinkertoy" || slice[i] == "thinkertoy.txt" {
				if slice[i] == "standard" || slice[i] == "shadow" || slice[i] == "thinkertoy" {
					InputFile = slice[i] + ".txt"
				} else {
					InputFile = slice[i]
				}
			} else if strings.Contains(slice[i], "--output=") {
				if len(slice) == 4 {
					log.Fatalln("You didn't enter the proply format you give the the flag of the output and two strings !!! or you give two flags")
				}
				OutputFile = slice[i][9:]
				TwoFlags++
			} else if strings.Contains(slice[i], "--color=") {
				if TwoFlags != 0 {
					log.Fatalln("You entered two flags !!!")
				}
				Color = slice[i][8:]
				if i < len(slice)-1 {
					lettertocollored = slice[i+1]
				}
				if i == len(slice)-3 {
					text = slice[i+2]
				}

			} else if text == "" {
				text = slice[i]
			}
			if OutputFile == "" && Color != "" {
				OutputFile = "NILL"
			}
			if InputFile == "" {
				InputFile = "standard.txt"
			}
			if text == "" && lettertocollored != "" {
				text = lettertocollored
			}
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

func CheckTheFlags(slice []string) bool {
	for i := 0; i < len(slice); i++ {
		if len(slice[i]) < 2 {
			continue
		} else if slice[i][:2] == "--" {
			sli := strings.Split(slice[i], "=")
			if len(sli) > 2 {
				return false
			} else {
				if sli[0] != "--color" && sli[0] != "--output" {
					return false
				}
			}
		}
	}
	return true
}

func ChoseTheColor(Paint string) string {
	if Paint != "" {
		Colors := []rune(Paint)

		if Colors[0] >= 'a' && Colors[0] <= 'z' {
			Colors[0] = Colors[0] - 32
		}
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
		case "Orange":
			Paint = "\033[38;5;208m"
		case "Purple":
			Paint = "\033[38;5;93m"
		case "Pink":
			Paint = "\033[38;5;205m"
		case "Brown":
			Paint = "\033[38;5;94m"
		case "LightBlue":
			Paint = "\033[38;5;153m"
		case "LightGreen":
			Paint = "\033[38;5;120m"
		case "LightYellow":
			Paint = "\033[38;5;228m"
		case "LightCyan":
			Paint = "\033[38;5;195m"
		case "LightMagenta":
			Paint = "\033[38;5;207m"
		case "LightGray":
			Paint = "\033[38;5;250m"
		default:
			log.Fatalln("Unknown color:", Paint)
		}
	} else {
		log.Fatalln("You dind't enter the color")
	}
	return Paint
}

func Format(InputFile, text string) ([]string, []string) {
	data, err := os.ReadFile(InputFile)
	if err != nil {
		log.Fatalln(err)
	}

	var sep string
	if InputFile == "standard.txt" || InputFile == "shadow.txt" {
		sep = "\n"
	} else {
		sep = "\r\n"
	}

	slice := RemoveEmptyStrings(strings.Split(string(data[1:]), sep))

	slicedArgs := strings.Split(text, "\\n")
	return slice, slicedArgs
}

func DrawAsciiArtColor(slice, slicedArgs []string, Paint, lettertocollored string) string {
	var result string
	var start int
	Reset := "\033[0m"
	for _, word := range slicedArgs {
		if word != "" {
			for i := 0; i < 8; i++ {
				if Paint != "" {
					//the case when we have a word to colore it
					if strings.Contains(word, lettertocollored) {
						temp := []rune(word)
						for k := 0; k < len(temp); k++ {
							if string(temp[k:k+len(lettertocollored)]) == lettertocollored {
								for _, char := range lettertocollored {
									if char < 32 || char > 126 {
										log.Fatalln("You entered an inprintabale charactwer")
									} else {
										start = int(char-32)*8 + i
										if k < k+len(lettertocollored) {
											result += fmt.Sprint(Paint, slice[start], Reset)
										}
									}
								}
								k += len(lettertocollored) - 1
							} else {
								start = int(temp[k]-32)*8 + i
								result += slice[start]
							}
						}
						result += "\n"
					} else {
						log.Fatalln("The word that you want to colored isn't exist")
					}
				} else {
					for _, char := range word {
						if char < 32 || char > 126 {
							log.Fatalln("You entered an inprintabale charactwer")
						} else {
							start = int(char-32)*8 + i
							result += slice[start]
						}
					}
					result += "\n"
				}
			}
		} else {
			result += "\n"
		}
	}
	return result
}
