package parser

type DictionaryEntry struct {
	Headword     string `json:"word" xml:"word"`
	PartOfSpeech string `json:"partofspeech" xml:"partOfSpeech,attr"`
	Definition   string `json:"definition" xml:"definition"`
}

type DictionaryEntries = []DictionaryEntry

// Root struct for XML ouput.
type Dictionary struct {
	DictionaryRoot DictionaryEntries `xml:"DictionaryEntry"`
}
