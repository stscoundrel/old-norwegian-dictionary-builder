package writer

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/stscoundrel/old-norwegian-dictionary-builder/parser"
)

func formatStructEntry(entry parser.DictionaryEntry) []string {
	structLines := []string{
		"        {",
		"            Headword: \"" + entry.Headword + "\",",
		"            Definition: \"" + strings.Replace(entry.Definition, "\"", "'", -1) + "\",",
		"            PartOfSpeech: \"" + entry.PartOfSpeech + "\",",
		"        },",
	}
	return structLines
}

func formatGoPackage(entries parser.DictionaryEntries) []string {
	goPackage := []string{
		"package dictionary",
		"",
		"",
		"type DictionaryEntry struct {",
		"	Headword     string",
		"	PartOfSpeech string",
		"	Definition   string",
		"}",
		"",
		"func GetDictionary() []DictionaryEntry {",
		"	return []DictionaryEntry{",
	}

	for _, entry := range entries {
		entryStruct := formatStructEntry(entry)
		goPackage = append(goPackage, entryStruct...)
	}

	goPackage = append(goPackage, "    }")
	goPackage = append(goPackage, "}")

	return goPackage
}

func WriteGoPackage(filePath string, entries parser.DictionaryEntries) {
	goLines := formatGoPackage(entries)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Error creating Go package file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, line := range goLines {
		_, _ = datawriter.WriteString(line + "\n")
	}

	datawriter.Flush()
	file.Close()
}
