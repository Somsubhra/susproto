package sus

import (
	"bufio"
	logger "log"
	"os"
)

func ReadNamesFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.Fatalf("failed to close file: %v", err)
		}
	}(file)

	var names []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		name := scanner.Text()
		names = append(names, name)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return names, nil
}
