package immutable

import (
	"html/template"
	"os"

	"github.com/pkg/errors"
)

type ReadmeRecord struct {
	CID          string
	Date         string
	DocumentName string
	DocumentPath string
}

type ReadmeData struct {
	DocumentName string
	Records      []ReadmeRecord
}

func renderREADME(data *ReadmeData, outputFilepath string) error {
	tpl, err := template.ParseFiles("readme/template.md.tpl")
	if err != nil {
		return errors.WithStack(err)
	}

	f, err := os.OpenFile(outputFilepath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.WithStack(err)
	}

	defer f.Close()

	return tpl.Execute(f, data)
}

func generateREADME(cfg *Config, outputFilepath string) error {
	records, err := ReadSortedRecords(cfg)
	if err != nil {
		return errors.WithStack(err)
	}

	rRecords := []ReadmeRecord{}

	for _, record := range records {
		rRecords = append(rRecords, ReadmeRecord{
			CID:          record.CID,
			Date:         record.Date.Format("Jan 02, 2006 - 15:04:05"),
			DocumentName: cfg.DocumentFileName,
			DocumentPath: finalResultPath(cfg),
		})
	}

	data := &ReadmeData{
		DocumentName: cfg.DocumentFileName,
		Records:      rRecords,
	}

	return renderREADME(data, outputFilepath)
}

// GenerateREADME generates a README document into given filepath.
func GenerateREADME(cfg *Config, readmeFilepath string) error {
	return generateREADME(cfg, readmeFilepath)
}

// GenerateREADMEFile generates a _README.md file.
func GenerateREADMEFile(cfg *Config) error {
	return generateREADME(cfg, "_README.md")
}
