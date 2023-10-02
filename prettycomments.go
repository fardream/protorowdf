package protorowdf

import "strings"

func processLineComment(line string) string {
	if strings.TrimSpace(line) == "" {
		return ""
	}

	l := strings.TrimRight(line, " \t\n\v\f\r")
	return " " + l
}

func prettyComments[T interface{ GetComment() string }](a T) []string {
	lines := strings.Split(strings.TrimSpace(a.GetComment()), "\n")
	result := make([]string, 0, len(lines))
	for _, l := range lines {
		result = append(result, processLineComment(l))
	}

	if len(result) == 1 && result[0] == "" {
		return nil
	}

	return result
}
