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

func ToGoPackage() {
	lines := reader.ReadLinesFromDictionary()
	entries := parser.ParseLines(lines)

	writer.WriteGoPackage("build/dictionary.go", entries)
}

func ToRustModule() {
	lines := reader.ReadLinesFromDictionary()
	entries := parser.ParseLines(lines)

	writer.WriteRustModule("build/dictionary.rs", entries)
}

func ToTypeScriptModule() {
	lines := reader.ReadLinesFromDictionary()
	entries := parser.ParseLines(lines)

	writer.WriteTypeScriptModule("build/dictionary.ts", entries)
}

func ToPythonClass() {
	lines := reader.ReadLinesFromDictionary()
	entries := parser.ParseLines(lines)

	writer.WritePythonClass("build/dictionary.py", entries)
}

func ToCSharpClass() {
	lines := reader.ReadLinesFromDictionary()
	entries := parser.ParseLines(lines)

	writer.WriteCSharpClass("build/Dictionary.cs", entries)
}

func main() {
	ToJson()
	ToDsl()
	ToXml()
	ToGoPackage()
	ToRustModule()
	ToTypeScriptModule()
	ToPythonClass()
	ToCSharpClass()
}
