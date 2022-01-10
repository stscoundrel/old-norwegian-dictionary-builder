package main

import (
	"github.com/stscoundrel/old-norwegian-dictionary-builder/parser"
	"github.com/stscoundrel/old-norwegian-dictionary-builder/reader"
	"github.com/stscoundrel/old-norwegian-dictionary-builder/writer"
)

func ToJson() {
	lines := reader.ReadLinesFromDictionary()
	entries := parser.ParseLines(lines)

	writer.WriteJson("build/dictionary.json", entries)
}

func ToDsl() {
	lines := reader.ReadLinesFromDictionary()
	entries := parser.ParseLines(lines)

	writer.WriteDsl("build/dictionary.dsl", entries)
}

func ToXml() {
	lines := reader.ReadLinesFromDictionary()
	entries := parser.ParseLines(lines)

	writer.WriteXml("build/dictionary.xml", entries)
}

func main() {
	ToJson()
	ToDsl()
	ToXml()
}
