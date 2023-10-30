package mods

import "regexp"

func FormatPunctuation(input string) string {
	removedExtraSpace := regexp.MustCompile(` {1,}`).ReplaceAllString(input, " ")
	bracketWithoutSpace := regexp.MustCompile(`\(\s*(.*?)\s*\)`).ReplaceAllString(removedExtraSpace, "($1)")
	reForJoin := regexp.MustCompile(` *([.,!?:;]) *`)
	formated := reForJoin.ReplaceAllString(bracketWithoutSpace, "$1")
	reForLastPunctuation := regexp.MustCompile(`([.,!?:;]+)([^ ]|$)`)
	formated = reForLastPunctuation.ReplaceAllString(formated, "$1 $2")
	reForApostrophe := regexp.MustCompile(`'\s*(.*?)\s*'`)
	formated = reForApostrophe.ReplaceAllString(formated, "'$1'")
	reForquote := regexp.MustCompile(`(.*?) +' +(.*?)`)
	formated = reForquote.ReplaceAllString(formated, "$1'$2")
	reForQuotes := regexp.MustCompile(`"\s*(.*?)\s*"`)
	formated = reForQuotes.ReplaceAllString(formated, `"$1"`)
	return formated
}