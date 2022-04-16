package writer

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/stscoundrel/old-norwegian-dictionary-builder/parser"
)

func formatInterfaceEntry(entry parser.DictionaryEntry) []string {
	structLines := []string{
		"  {",
		"    word: '" + entry.Headword + "',",
		"    definition: '" + strings.Replace(entry.Definition, "'", "\\'", -1) + "',",
		"    partOfSpeech: '" + entry.PartOfSpeech + "',",
		"  },",
	}
	return structLines
}

func formatTypeScriptModule(entries parser.DictionaryEntries) []string {
	typeScriptModule := []string{
		"export interface DictionaryEntry{",
		"  word: string,",
		"  partOfSpeech: string,",
		"  definition: string,",
		"}",
		"",
		"const entries = [",
	}

	for _, entry := range entries {
		entryStruct := formatInterfaceEntry(entry)
		typeScriptModule = append(typeScriptModule, entryStruct...)
	}

	typeScriptModule = append(typeScriptModule, "]")
	typeScriptModule = append(typeScriptModule, "")
	typeScriptModule = append(typeScriptModule, "export const getDictionary = (): DictionaryEntry[] => entries")
	typeScriptModule = append(typeScriptModule, "")
	typeScriptModule = append(typeScriptModule, "export default getDictionary")

	return typeScriptModule
}

func WriteTypeScriptModule(filePath string, entries parser.DictionaryEntries) {
	goLines := formatTypeScriptModule(entries)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Error creating TypeScript module file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, line := range goLines {
		_, _ = datawriter.WriteString(line + "\n")
	}

	datawriter.Flush()
	file.Close()
}
