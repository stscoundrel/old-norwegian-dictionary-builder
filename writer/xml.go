package writer

import (
	"encoding/xml"
	"io/ioutil"
	"log"

	"github.com/stscoundrel/old-norwegian-dictionary-builder/parser"
)

func WriteXml(filePath string, entries parser.DictionaryEntries) {
	entriesWithRoot := parser.Dictionary{DictionaryRoot: entries}
	file, _ := xml.MarshalIndent(entriesWithRoot, "", "  ")

	err := ioutil.WriteFile(filePath, file, 0644)

	if err != nil {
		log.Fatalf("Error creating JSON file: %s", err)
	}
}
