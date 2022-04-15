package main

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestCreatesJsonBuild(t *testing.T) {
	ToJson()

	_, err := os.Stat("./build/dictionary.json")

	fmt.Println(err)

	if errors.Is(err, os.ErrNotExist) {
		t.Error("JSON output file not found in build directory")
	}
}

func TestCreatesDslBuild(t *testing.T) {
	ToDsl()

	_, err := os.Stat("./build/dictionary.dsl")

	fmt.Println(err)

	if errors.Is(err, os.ErrNotExist) {
		t.Error("DSL output file not found in build directory")
	}
}

func TestCreatesXmlBuild(t *testing.T) {
	ToXml()

	_, err := os.Stat("./build/dictionary.xml")

	fmt.Println(err)

	if errors.Is(err, os.ErrNotExist) {
		t.Error("XML output file not found in build directory")
	}
}

func TestCreatesGoPackageBuild(t *testing.T) {
	ToGoPackage()

	_, err := os.Stat("./build/dictionary.go")

	fmt.Println(err)

	if errors.Is(err, os.ErrNotExist) {
		t.Error("Go package output file not found in build directory")
	}
}

func TestCreatesTypeScriptModuleBuild(t *testing.T) {
	ToTypeScriptModule()

	_, err := os.Stat("./build/dictionary.ts")

	fmt.Println(err)

	if errors.Is(err, os.ErrNotExist) {
		t.Error("TypeScript module output file not found in build directory")
	}
}
