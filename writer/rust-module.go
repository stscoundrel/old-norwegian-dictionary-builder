package writer

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/stscoundrel/old-norwegian-dictionary-builder/parser"
)

func formatRustStructEntry(entry parser.DictionaryEntry) []string {
	structLines := []string{
		"	DictionaryEntry{",
		"    	word: \"" + entry.Headword + "\",",
		"    	definition: \"" + strings.Replace(entry.Definition, "\"", "'", -1) + "\",",
		"    	part_of_speech: \"" + entry.PartOfSpeech + "\",",
		"	},",
	}
	return structLines
}

func formatRustModule(entries parser.DictionaryEntries) []string {
	rustModule := []string{
		"pub struct DictionaryEntry {",
		"	pub word: &'static str,",
		"	pub part_of_speech: &'static str,",
		"	pub definition: &'static str,",
		"}",
		"",
		"static DICTIONARY: [DictionaryEntry; 42021] = [",
	}
	for _, entry := range entries {
		entryStruct := formatRustStructEntry(entry)
		rustModule = append(rustModule, entryStruct...)
	}

	rustModule = append(rustModule, "];")
	rustModule = append(rustModule, "")
	rustModule = append(rustModule, "pub fn get_dictionary() -> [DictionaryEntry; 42021] {")
	rustModule = append(rustModule, "	return dictionary;")
	rustModule = append(rustModule, "}")

	return rustModule
}

func WriteRustModule(filePath string, entries parser.DictionaryEntries) {
	goLines := formatRustModule(entries)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Error creating Rust module file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, line := range goLines {
		_, _ = datawriter.WriteString(line + "\n")
	}

	datawriter.Flush()
	file.Close()
}
