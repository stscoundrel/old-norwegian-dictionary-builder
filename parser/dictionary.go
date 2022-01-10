package parser

type DictionaryEntry struct {
	Headword     string `json:"word" xml:"word"`
	PartOfSpeech string `json:"partofspeech" xml:"partOfSpeech,attr"`
	Definition   string `json:"definition" xml:"definition"`
}

type DictionaryEntries = []DictionaryEntry
