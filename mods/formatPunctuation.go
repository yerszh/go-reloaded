package mods

import (
	"regexp"
	"strings"
)

func replaceAtoAn(text string) string {
	re := regexp.MustCompile(`\b[aA]\b \b[aAeEiIoOuUhH]\w*`)
	replaced := re.ReplaceAllStringFunc(text, func(match string) string {
		splitted := strings.Split(match, " ")
		if splitted[0] == "a" {
			splitted[0] = "an"
		} else if splitted[0] == "A" {
			splitted[0] = "An"
		}
		return re.ReplaceAllString(match, strings.Join(splitted, " "))
	})
	return replaced
}

func FormatPunctuation(text string) string {

	replacedAtoAn := replaceAtoAn(text)

	caseLow := regexp.MustCompile(`(?i)low`).ReplaceAllString(replacedAtoAn, "low")
	caseCap := regexp.MustCompile(`(?i)cap`).ReplaceAllString(caseLow, "cap")
	caseUp := regexp.MustCompile(`(?i)up`).ReplaceAllString(caseCap, "up")
	caseHex := regexp.MustCompile(`(?i)hex`).ReplaceAllString(caseUp, "hex")
	caseBin := regexp.MustCompile(`(?i)bin`).ReplaceAllString(caseHex, "bin")

	replaceAllBacktick := regexp.MustCompile("`").ReplaceAllString(caseBin, "'")

	removedExtraSpace := regexp.MustCompile(` {1,}`).ReplaceAllString(replaceAllBacktick, " ")

	bracketWithoutSpace := regexp.MustCompile(`\(\s*(.*?)\s*\)`).ReplaceAllString(removedExtraSpace, "($1)")

	formated := regexp.MustCompile(` *([.,!?:;]) *`).ReplaceAllString(bracketWithoutSpace, "$1")

	formated = regexp.MustCompile(`([.,!?:;]+)([^ ]|$)`).ReplaceAllString(formated, "$1 $2")

	formatedApostrophe := regexp.MustCompile(`'\s*(.*?)\s*'`).ReplaceAllString(formated, "'$1'")

	formated = regexp.MustCompile(`(.*?) +' +(.*?)`).ReplaceAllString(formatedApostrophe, "$1'$2")

	formatedQuotes := regexp.MustCompile(`"\s*(.*?)\s*"`).ReplaceAllString(formated, `"$1"`)

	spaceStartLine := regexp.MustCompile(`^\s*`).ReplaceAllString(formatedQuotes, "")

	dumbCase := regexp.MustCompile(`\(lower case\)`).ReplaceAllString(spaceStartLine, "(low)")

	return dumbCase
}
