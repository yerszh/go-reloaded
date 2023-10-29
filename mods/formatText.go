package mods

import (
	"regexp"
	"strconv"
	"strings"
)

func replaceNum(number, match string) string {
	numSystems := map[string]int{
		"hex": 16,
		"bin": 2,
	}

	decimal, _ := strconv.ParseInt(number, numSystems[match], 64)
	convertedNum := strconv.Itoa(int(decimal))

	if convertedNum == "0" {
		return number
	} else {
		return convertedNum
	}
}
func capitalizeWord(s string) string {
	runes := []rune(s)
	capitalize := true

	for i, r := range runes {
		if isAlpha := (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z'); isAlpha {
			if capitalize && r >= 'a' && r <= 'z' {
				runes[i] -= 32
			} else if !capitalize && r >= 'A' && r <= 'Z' {
				runes[i] += 32
			}
			capitalize = false
		} else {
			capitalize = true
		}
	}

	return string(runes)
}

func convertMatch(allPrevWords, match, reMatchSlash string) string {
	isNewLine := regexp.MustCompile(`\n` + reMatchSlash)
	if isNewLine.MatchString(allPrevWords) {
		return strings.TrimSuffix(allPrevWords, match)
	}

	wordCase := regexp.MustCompile(`cap|up|low|bin|hex`).FindString(match)
	wordNum, err := strconv.Atoi(regexp.MustCompile(`\d+`).FindString(match))
	if err != nil {
		wordNum = 1
	}

	onlyPrevWords := regexp.MustCompile(`(\s+)$`).ReplaceAllString(strings.TrimSuffix(allPrevWords, match), "")
	onlyPrevWords = strings.Replace(onlyPrevWords, "\n", " \n", -1)
	words := regexp.MustCompile(` `).Split(onlyPrevWords, -1)
	wordCount := 0

	for index := len(words) - 1; index >= 0; index-- {
		wordTemp := words[index]

		if wordNum == wordCount {
			break
		}

		switch wordCase {
		case "bin":
			words[index] = replaceNum(wordTemp, wordCase) + " "
		case "hex":
			words[index] = replaceNum(wordTemp, wordCase) + " "
		case "up":
			words[index] = strings.ToUpper(wordTemp) + " "
		case "low":
			words[index] = strings.ToLower(wordTemp) + " "
		case "cap":
			words[index] = capitalizeWord(wordTemp) + " "
		}

		wordCount++
	}

	return strings.Join(words, " ")
}

func FormatText(inputText string) string {
	reForConvert := regexp.MustCompile(`\((cap|up|low|hex|bin)((, \w+)|(, \w+\W+))?\) ?`)
	matches := reForConvert.FindAllString(inputText, -1)

	for _, match := range matches {
		reMatchSlash := regexp.MustCompile(`([()])`).ReplaceAllString(match, `\$1`)
		reAllPrevWords := regexp.MustCompile(`(\S+\s+|\s+\S+)*` + reMatchSlash) //`(.+?)\(%s\)`
		allPrevWords := reAllPrevWords.FindString(inputText)
		converted := convertMatch(allPrevWords, match, reMatchSlash)
		inputText = strings.Replace(inputText, allPrevWords, converted, 1)
	}

	return inputText
}
