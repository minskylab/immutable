package immutable

import (
	"html/template"
	"os"

	"github.com/pkg/errors"
)

type ReadmeRecord struct {
	SerialNumber int
	CID          string
	Date         string
	DocumentName string
	DocumentPath string
}

type ReadmeData struct {
	DocumentTitle string
	DocumentName  string
	Records       []ReadmeRecord
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

	totalRecords := len(records)

	for i, record := range records {
		rRecords = append(rRecords, ReadmeRecord{
			SerialNumber: totalRecords - i,
			CID:          record.CID,
			Date:         record.Date.Format("Jan 02, 2006 at 15:04:05"),
			DocumentName: cfg.Document.OutputFilename,
			DocumentPath: finalResultPath(cfg),
		})
	}

	data := &ReadmeData{
		DocumentTitle: cfg.Document.Title,
		DocumentName:  cfg.Document.OutputFilename,
		Records:       rRecords,
	}

	return renderREADME(data, outputFilepath)
}

// GenerateREADME generates a README document into given filepath.
func GenerateREADME(cfg *Config, readmeFilepath string) error {
	return generateREADME(cfg, readmeFilepath)
}

// GenerateREADMEFile generates a README.md or _README.md file based on production config flag.
// If production flag is true, then the file is named README.md, otherwise it is named _README.md.
func GenerateREADMEFile(cfg *Config) error {
	readmeFilename := "_README.md"

	if cfg.Production {
		readmeFilename = "README.md"
	}

	return generateREADME(cfg, readmeFilename)
}
