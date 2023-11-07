package modifications

import (
	"regexp"
	"strconv"
	"strings"
)

func toDecimal(num, sys string) string {
	numSys := map[string]int{
		"(hex)": 16,
		"(oct)": 8,
		"(bin)": 2,
	}
	decimal, _ := strconv.ParseInt(num, numSys[sys], 64)
	return strconv.Itoa(int(decimal))
}

func ReplaceNumbers(input string) string {
	re := regexp.MustCompile(`\w*\s*\((hex|bin|oct)\)`)
	replaced := re.ReplaceAllStringFunc(input, func(match string) string {
		validMatch, _ := regexp.MatchString(`[0-9A-Fa-f]+\s*\(hex\)|[0-1]+\s*\(bin\)|[0-7]+\s*\(oct\)`, match)
		if validMatch {
			number := regexp.MustCompile(`[\dA-Fa-f]+`).FindString(match)
			numSystem := regexp.MustCompile(`\((hex|bin|oct)\)`).FindString(match)
			return re.ReplaceAllString(match, toDecimal(number, numSystem))
		} else {
			toRemove := regexp.MustCompile(`\((hex|bin|oct)\)`).FindString(match)
			return strings.TrimSuffix(match, toRemove)
		}
	})
	return replaced
}
