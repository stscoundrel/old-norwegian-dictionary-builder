package writer

import (
	"encoding/json"
	"io/ioutil"

	"github.com/stscoundrel/old-norwegian-dictionary-builder/parser"
)

func WriteJson(filePath string, entries parser.DictionaryEntries) {
	file, _ := json.MarshalIndent(entries, "", " ")

	_ = ioutil.WriteFile(filePath, file, 0644)
}
