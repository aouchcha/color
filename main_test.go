package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const InputFile string = "thinkertoy.txt"

// const InputFile string = "standard.txt"
// const InputFile string = "shadow.txt"
func Test_main(t *testing.T) {
	var want string
	tests := LoadTests()
	for _, test := range tests {
		if test == "" {
			continue
		}
		// run the tests in the main and stock the result from stdout.
		got, err := exec.Command("go", "run", ".", test, InputFile).Output()
		if err != nil {
			log.Fatalln(err)
		}
		want = FS(test)
		// Compare the result that the main.go give and the test give if they are the same
		if want == string(got) {
			t.Logf(test)
		} else {
			t.Fatal(test)
		}
	}
}

func FS(text string) string {
	slice := TamplateFormat(InputFile)
	if len(slice) != 760 {
		log.Fatalln("You have changed the Input file !!!!")
	}
	want := DrawFS(slice, text)
	return want
}

func LoadTests() []string {
	data, err := os.ReadFile("TESTFILE.txt")
	if err != nil {
		log.Fatalln("Error :", err)
	}
	text := strings.Split(string(data), "\n")
	return text
}

func TamplateFormat(InputFile string) []string {
	data, err := os.ReadFile(InputFile)
	if err != nil {
		log.Fatalln("Error :", err)
	}
	slice := DeleteEmptySlices(strings.Split(strings.ReplaceAll(string(data), "\r", ""), "\n"))
	return slice
}

func DrawFS(slice []string, text string) string {
	var result string
	if text != "" {
		slicedArg := strings.Split(text, "\\n")
		for _, word := range slicedArg {
			if word != "" {
				for j := 0; j < 8; j++ {
					for _, char := range word {
						if char < 32 || char > 126 {
							log.Fatalln("You netered an inprintabal charactere !!!!")
						} else {
							start := int(char-32)*8 + j

							result += slice[start]

						}
					}
					result += "\n"
				}
			} else {
				result += "\n"
			}
		}
	} else {
		result += "\n"
	}
	result = IsItNewLine(result)
	return result
}

func DeleteEmptySlices(slice []string) []string {
	var temp []string
	for i := range slice {
		if slice[i] != "" {
			temp = append(temp, slice[i])
		}
	}
	return temp
}

func IsItNewLine(result string) string {
	for _, char := range result {
		if char != '\n' {
			return result
		}
	}
	return result[1:]
}
