package writer

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/stscoundrel/old-norwegian-dictionary-builder/parser"
)

func WriteJson(filePath string, entries parser.DictionaryEntries) {
	file, _ := json.MarshalIndent(entries, "", " ")

	err := ioutil.WriteFile(filePath, file, 0644)

	if err != nil {
		log.Fatalf("Error creating JSON file: %s", err)
	}
}
