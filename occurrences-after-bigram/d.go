package occurrences_after_bigram

import "strings"

func findOcurrences(text string, first string, second string) []string {
	words := strings.Fields(text)
	var result []string
	for i := 0; i < len(words)-2; i++ {
		if words[i] == first && words[i+1] == second {
			result = append(result, words[i+2])
		}
	}
	return result
}

func findOcurrences1(text string, first string, second string) []string {
	var result []string
	dest := first + " " + second + " "
	index := strings.Index(text, dest)
	for index != -1 {
		newText := text[index+len(dest):]
		if index == 0 || text[index-1] == ' ' {
			result = append(result, getFirstWord(newText))
		}
		text = newText
		index = strings.Index(text, dest)
	}
	return result
}

func getFirstWord(s string) string {
	index := strings.Index(s, " ")
	if index == -1 {
		return s
	}
	return s[:index]
}
