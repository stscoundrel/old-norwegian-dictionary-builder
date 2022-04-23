package writer

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/stscoundrel/old-norwegian-dictionary-builder/parser"
)

func formatCSharpEntry(entry parser.DictionaryEntry) []string {
	objectLines := []string{
		"		new DictionaryEntry(\"" + entry.Headword + "\", \"" + entry.PartOfSpeech + "\", \"" + strings.Replace(entry.Definition, "\"", "\\\"", -1) + "\"),",
	}
	return objectLines
}

func formatCSharpClass(entries parser.DictionaryEntries) []string {
	cSharpClass := []string{
		"namespace OldNorwegianDictionary;",
		"",
		"public readonly struct DictionaryEntry",
		"{",
		"	public string Headword { get; init; }",
		"	public string PartOfSpeech { get; init; }",
		"	public string Definition { get; init; }",
		"",
		"	public DictionaryEntry(string headword, string partOfSpeech, string definition)",
		"	{",
		"		Headword = headword;",
		"		PartOfSpeech = partOfSpeech;",
		"		Definition = definition;",
		"	}",
		"}",
		"",
		"public class Dictionary",
		"{",
		"",
		"	private static readonly DictionaryEntry[] Entries = {",
	}

	for _, entry := range entries {
		entryStruct := formatCSharpEntry(entry)
		cSharpClass = append(cSharpClass, entryStruct...)
	}

	classEnd := []string{
		"	};",
		"",
		"	public DictionaryEntry[] GetEntries() {",
		"		return Entries;",
		"	}",
		"}",
	}

	cSharpClass = append(cSharpClass, classEnd...)

	return cSharpClass
}

func WriteCSharpClass(filePath string, entries parser.DictionaryEntries) {
	goLines := formatCSharpClass(entries)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Error creating C# class file(s): %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, line := range goLines {
		_, _ = datawriter.WriteString(line + "\n")
	}

	datawriter.Flush()
	file.Close()
}
