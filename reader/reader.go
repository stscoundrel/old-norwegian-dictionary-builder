package reader

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

const LINES_IN_DICTIONARY = 42022

func getDictionaryPath() string {
	_, base, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(base)

	return basepath + "/resources/dictionary.txt"
}

func ReadLinesFromDictionary() []string {
	result := make([]string, 0, LINES_IN_DICTIONARY)
	file, err := os.Open(getDictionaryPath())

	if err != nil {
		fmt.Println("Could not read dictionary file: ", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	/**
	 * Increase default buf capacity.
	 * Dictionar contains some _long_ lines,
	 * which would otherwise break.
	 */
	const maxCapacity = 70000
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	file.Close()

	return result
}
