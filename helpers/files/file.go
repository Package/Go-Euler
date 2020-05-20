package files

import (
	"bufio"
	"os"
	"strings"
)

// Returns a single string containing all the data in the file
func ReadAllAsString(filePath string) (string, error) {

	handle, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	// like a finally {} block. closes file at the end of function even if error occurs
	defer handle.Close()

	// Scan the file in line by line
	str := ""
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		line := scanner.Text()
		str += line
	}

	return str, nil
}

// Returns a slice of strings. Each line in the file is a new entry in the slice.
func ReadAsStringArray(filePath string) ([] string, error) {

	handle, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer handle.Close()

	stringData := []string {}
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		line := scanner.Text()
		stringData = append(stringData, line)
	}

	return stringData, nil
}

// Returns the contents of the file parsed into a multi dimensional array.
// Each line is a new entry within the array and then each word separated by space is a new entry within that.
func ReadAsMultiArray(filePath string) ([][]string, error) {

	handle, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer handle.Close()

	data := [][]string {}
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, strings.Split(line, " "))
	}

	return data, nil
}