package parser

import "strings"

func formatLineToEntry(line string) DictionaryEntry {
	parts := strings.Split(line, "|")

	return DictionaryEntry{
		headword:     parts[0],
		partOfSpeech: parts[1],
		definition:   parts[2],
	}
}

func ParseLines(lines []string) DictionaryEntries {
	entries := DictionaryEntries{}

	for _, line := range lines {
		entries = append(entries, formatLineToEntry(line))
	}

	return entries
}
