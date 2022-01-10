package writer

import (
	"bufio"
	"log"
	"os"

	"github.com/stscoundrel/old-norwegian-dictionary-builder/parser"
)

func formatEntry(entry parser.DictionaryEntry) (string, string) {
	return entry.Headword, "    " + entry.Definition
}

func formatDSL(entries parser.DictionaryEntries) []string {
	dsl := []string{
		"#NAME Dictionary of the Old Norwegian Language",
		"#INDEX_LANGUAGE \"Old Norse\"",
		"#CONTENTS_LANGUAGE \"Norwegian\"",
	}

	for _, entry := range entries {
		headwordLine, definitionLine := formatEntry(entry)
		dsl = append(dsl, headwordLine)
		dsl = append(dsl, definitionLine)
	}

	return dsl
}

func WriteDsl(filePath string, entries parser.DictionaryEntries) {
	dslLines := formatDSL(entries)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Error creating DSL file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, line := range dslLines {
		_, _ = datawriter.WriteString(line + "\n")
	}

	datawriter.Flush()
	file.Close()
}
