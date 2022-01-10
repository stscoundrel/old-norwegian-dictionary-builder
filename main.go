package main

import (
	"github.com/stscoundrel/old-norwegian-dictionary-builder/parser"
	"github.com/stscoundrel/old-norwegian-dictionary-builder/reader"
	"github.com/stscoundrel/old-norwegian-dictionary-builder/writer"
)

func dummyMethod() string {
	return "Lorem ipsum dolor sit amet"
}

func ToJson() {
	lines := reader.ReadLinesFromDictionary()
	entries := parser.ParseLines(lines)

	writer.WriteJson("build/dictionary.json", entries)
}

func main() {
	ToJson()
}
