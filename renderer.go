package immutable

import (
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/thecodingmachine/gotenberg-go-client/v7"
)

const (
	recordsName = "records"
)

func finalResultPath(cfg *Config) string {
	return path.Join(cfg.ImmutableDir, cfg.Document.OutputFilename)
}

func recordsResultPath(cfg *Config) string {
	return path.Join(cfg.ImmutableDir, recordsName)
}

// list files in directory
func filesFromDirectory(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result []string
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		result = append(result, f.Name())
	}

	return result, nil
}

func obtainMDFiles(cfg *Config) ([]gotenberg.Document, error) {
	files, err := filesFromDirectory(cfg.TemplatesDir)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// var result []string
	documents := []gotenberg.Document{}
	for _, f := range files {
		if strings.HasSuffix(strings.ToLower(f), ".md") {
			filepath := path.Join(cfg.TemplatesDir, f)

			doc, err := gotenberg.NewDocumentFromPath(f, filepath)
			if err != nil {
				return nil, errors.WithStack(err)
			}

			documents = append(documents, doc)
		}
	}

	return documents, nil
}

func generatePDF(cfg *Config) error {
	c := &gotenberg.Client{Hostname: cfg.Gotenberg.Hostname}

	indexFilename := "index.html"
	indexPath := path.Join(cfg.TemplatesDir, indexFilename)
	index, err := gotenberg.NewDocumentFromPath(indexFilename, indexPath)
	if err != nil {
		return errors.WithStack(err)
	}

	mdDocuments, err := obtainMDFiles(cfg)
	if err != nil {
		return errors.WithStack(err)
	}

	req := gotenberg.NewMarkdownRequest(index, mdDocuments...)

	if err := c.Store(req, finalResultPath(cfg)); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// GeneratePDF generates a PDF from the given config.
func GeneratePDF(cfg *Config) error {
	return generatePDF(cfg)
}
