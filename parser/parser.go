package parser

import (
	"strings"
)

const VALID_ENTRIES_IN_DICTIONARY = 42021

func formatLineToEntry(line string) (bool, DictionaryEntry) {
	parts := strings.Split(line, "|")

	if len(parts) < 3 {
		return false, DictionaryEntry{}
	}

	return true, DictionaryEntry{
		Headword:     parts[0],
		PartOfSpeech: parts[1],
		Definition:   parts[2],
	}
}

func ParseLines(lines []string) DictionaryEntries {
	entries := make(DictionaryEntries, 0, VALID_ENTRIES_IN_DICTIONARY)

	for _, line := range lines {
		isValidEntry, entry := formatLineToEntry(line)

		if isValidEntry {
			entries = append(entries, entry)
		}
	}

	return entries
}
