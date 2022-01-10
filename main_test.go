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
