package modifications

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func convertMatch(match, submatch string) string {
	wordCase := regexp.MustCompile(`cap|up|low`).FindString(submatch)
	wordNum, _ := strconv.Atoi(regexp.MustCompile(`\d+`).FindString(submatch))
	formattedWords := regexp.MustCompile(`(\s+)$`).ReplaceAllString(strings.TrimSuffix(match, submatch), "")
	words := strings.Split(formattedWords, " ")
	counter := 0

	for index := len(words) - 1; index >= 0; index-- {
		if counter == wordNum {
			break
		}
		switch wordCase {
		case "up":
			words[index] = strings.ToUpper(words[index])
		case "low":
			words[index] = strings.ToLower(words[index])
		case "cap":
			if words[index] != "" {
				firstLetter := strings.ToUpper(words[index][0:1])
				restOfString := words[index][1:]
				words[index] = firstLetter + restOfString
			}
		}
		counter++
	}
	return strings.Join(words, " ")
}

func ConvertWords(input string) string {
	reForConvert := regexp.MustCompile(`\((cap|up|low), \d+\)|\((cap|up|low)\)`)
	submatches := reForConvert.FindAllString(input, -1)
	text := input
	for _, submatch := range submatches {
		rePrevWords := regexp.MustCompile(fmt.Sprintf(`(\S*\s*)*\(%s\)`, submatch))
		allPrevWords := rePrevWords.FindString(text)
		converted := convertMatch(allPrevWords, submatch)
		text = strings.Replace(text, allPrevWords, converted, -1)
	}
	return text
}
