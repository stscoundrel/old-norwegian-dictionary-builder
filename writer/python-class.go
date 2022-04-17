package writer

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/stscoundrel/old-norwegian-dictionary-builder/parser"
)

func formatPythonClassEntry(entry parser.DictionaryEntry) []string {
	structLines := []string{
		"		DictionaryEntry(",
		"			word='" + entry.Headword + "',",
		"			part_of_speech='" + entry.PartOfSpeech + "',",
		"			definition='" + strings.Replace(entry.Definition, "'", "\\'", -1) + "',",
		"		),",
	}
	return structLines
}

func formatPythonClass(entries parser.DictionaryEntries) []string {
	pythonClass := []string{
		"from typing import NamedTuple, Tuple",
		"",
		"",
		"class DictionaryEntry(NamedTuple):",
		"	word: str",
		"	part_of_speech: str",
		"	definition: str",
		"",
		"",
		"class OldNorwegianDictionary:",
		"	entries: Tuple[DictionaryEntry, ...] = (",
	}

	for _, entry := range entries {
		entryStruct := formatPythonClassEntry(entry)
		pythonClass = append(pythonClass, entryStruct...)
	}

	classMethod := []string{
		"	)",
		"",
		"	@classmethod",
		"	def get_dictionary(cls) -> Tuple[DictionaryEntry, ...]:",
		"		return cls.entries",
		"",
	}

	pythonClass = append(pythonClass, classMethod...)

	return pythonClass
}

func WritePythonClass(filePath string, entries parser.DictionaryEntries) {
	goLines := formatPythonClass(entries)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Error creating Python class file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, line := range goLines {
		_, _ = datawriter.WriteString(line + "\n")
	}

	datawriter.Flush()
	file.Close()
}
