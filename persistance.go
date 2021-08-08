package immutable

import (
	"os"
	"sort"
	"strings"

	"github.com/pkg/errors"
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

func readRecords(filePath string) ([]ImmutableRecord, error) {
	records := []ImmutableRecord{}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		record, err := parseFromLine(line)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		records = append(records, record)
	}

	return records, nil
}

func AddRecordToLogs(cfg *Config, record *ImmutableRecord) error {
	return appendLineToFile(recordsResultPath(cfg), record.RecordLine())
}

func ReadSortedRecords(cfg *Config) ([]ImmutableRecord, error) {
	records, err := readRecords(recordsResultPath(cfg))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].Date.After(records[j].Date)
	})

	return records, nil
}
