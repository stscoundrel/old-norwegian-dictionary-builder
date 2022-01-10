package writer

import (
	"encoding/xml"
	"io/ioutil"
	"log"

	"github.com/stscoundrel/old-norwegian-dictionary-builder/parser"
)

func WriteXml(filePath string, entries parser.DictionaryEntries) {
	file, _ := xml.MarshalIndent(entries, "", "  ")

	err := ioutil.WriteFile(filePath, file, 0644)

	if err != nil {
		log.Fatalf("Error creating JSON file: %s", err)
	}
}
