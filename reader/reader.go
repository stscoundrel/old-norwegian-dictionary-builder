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

	const maxCapacity = 165536
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	file.Close()

	return result
}
