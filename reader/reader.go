package reader

import (
	"bufio"
	"fmt"
	"os"
)

const DICTIONARY_PATH = "./resources/dictionary.txt"

func ReadLinesFromDictionary() []string {
	result := []string{}
	file, err := os.Open(DICTIONARY_PATH)

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
