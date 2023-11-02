package modifications

import (
	"regexp"
	"strings"
)

func ReplaceAWithAn(input string) string {
	re := regexp.MustCompile(`\b[aA]\b \b[aAeEiIoOuUhH]\w*`)
	replaced := re.ReplaceAllStringFunc(input, func(match string) string {
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
