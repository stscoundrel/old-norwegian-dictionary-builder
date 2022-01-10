package parser

type DictionaryEntry struct {
	Headword     string
	PartOfSpeech string
	Definition   string
}

type DictionaryEntries = []DictionaryEntry
