package utils

import (
	"os"
	"regexp"
	"strconv"
)

func LoadData() string {
	file := "input"

	if os.Getenv("TEST") != "" {
		file = "test_input"
	}

	data, error := os.ReadFile(file)

	if error != nil {
		return ""
	}

	return string(data)
}

func SafeStringToInt(input string) int {
	marks, err := strconv.Atoi(input)
	if err != nil {
		return 0
	}
	return marks
}

func GetRegexGroups(regEx regexp.Regexp, matches []string) (groups map[string]string) {
	groups = map[string]string{}
	for i, name := range regEx.SubexpNames() {
		if i > 0 && i <= len(matches) {
			groups[name] = matches[i]
		}
	}
	return groups
}
