package immutable

import (
	"os"
	"strings"
)

func fileExists(filePath string) bool {
	fileinfo, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		return false
	}

	return !fileinfo.IsDir()
}

func appendLineToFile(filePath string, line string) error {
	existsBefore := fileExists(filePath)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	line = strings.Trim(line, " \n")

	if existsBefore {
		line = "\n" + line
	}

	_, err = file.WriteString(line)
	return err
}

func AddRecordToLogs(cfg *Config, record *ImmutableRecord) error {
	return appendLineToFile(recordsResultPath(cfg), record.RecordLine())
}
