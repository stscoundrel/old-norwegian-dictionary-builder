package parser

type DictionaryEntry struct {
	Headword     string `json:"word"`
	PartOfSpeech string `json:"partofspeech"`
	Definition   string `json:"definition"`
}

type DictionaryEntries = []DictionaryEntry
