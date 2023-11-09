package main

import (
	"fmt"
	"os"

	"go-reloaded/modifications"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error")
		return
	}

	inputName := os.Args[1]
	outputName := os.Args[2]

	if !(inputName[len(inputName)-4:] == ".txt" && outputName[len(outputName)-4:] == ".txt") {
		fmt.Println("Error")
		return
	}

	inputBytes, err := os.ReadFile(inputName)
	if err != nil {
		fmt.Println(err)
		return
	}
	outputText := modifications.FormatPunctuation(string(inputBytes))
	outputText = modifications.ReplaceAWithAn(outputText)
	outputText = modifications.ReplaceNumbers(outputText)
	outputText = modifications.ConvertWords(outputText)
	outputText = modifications.FormatPunctuation(outputText)
	err = os.WriteFile(outputName, []byte(outputText), 0o777)
	if err != nil {
		fmt.Println(err)
		return
	}
}
