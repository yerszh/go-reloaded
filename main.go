package main

import (
	"fmt"
	"os"
	"path/filepath"

	"go-reloaded/mods"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(`Error! go run . "(filename to read).txt" "(filename to write).txt"`)
		return
	}

	inputFileName, outputFileName := os.Args[1], os.Args[2]

	if filepath.Ext(inputFileName) != ".txt" || filepath.Ext(outputFileName) != ".txt" {
		fmt.Println("Error! Arguments extension must be .txt")
		return
	}

	text, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	formattedText := string(text)

	for range "12" {
		formattedText = mods.FormatText(mods.FormatPunctuation(formattedText))
	}

	err = os.WriteFile(outputFileName, []byte(formattedText), 0o777)
	if err != nil {
		fmt.Println(err)
		return
	}
}
