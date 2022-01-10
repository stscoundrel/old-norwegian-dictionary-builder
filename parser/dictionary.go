package parser

type DictionaryEntry struct {
	headword     string
	partOfSpeech string
	definition   string
}

type DictionaryEntries = []DictionaryEntry
