package main

func cleanInput(text string) []string {
	var cleanedText []string

	if len(text) <= 0 {
		return cleanedText
	}

	parsedWord := ""

	for _, c := range text {
		if c == ' ' {

			if len(parsedWord) != 0 {
				cleanedText = append(cleanedText, parsedWord)
			}
			parsedWord = ""
			continue
		}

		parsedWord += string(c)
	}

	return cleanedText
}