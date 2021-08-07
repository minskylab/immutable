package immutable

import (
	"path"

	// "github.com/k0kubun/pp"
	"github.com/pkg/errors"
	"github.com/thecodingmachine/gotenberg-go-client/v7"
)

const (
	resultDocumentName = "result.pdf"
	recordsName        = "records"
)

func finalResultPath(cfg *Config) string {
	return path.Join(cfg.ImmutableDir, resultDocumentName)
}

func recordsResultPath(cfg *Config) string {
	return path.Join(cfg.ImmutableDir, recordsName)
}

func generatePDF(cfg *Config) error {
	c := &gotenberg.Client{Hostname: cfg.Gotenberg.Hostname}

	indexFilename := "index.html"
	indexPath := path.Join(cfg.TemplatesDir, indexFilename)
	index, err := gotenberg.NewDocumentFromPath(indexFilename, indexPath)
	if err != nil {
		return errors.WithStack(err)
	}

	markdownFilename := "DOCUMENT.md"
	markdownPath := path.Join(cfg.TemplatesDir, markdownFilename)
	markdown, err := gotenberg.NewDocumentFromPath(markdownFilename, markdownPath)
	if err != nil {
		return errors.WithStack(err)
	}

	req := gotenberg.NewMarkdownRequest(index, markdown)

	if err := c.Store(req, finalResultPath(cfg)); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func GeneratePDF(cfg *Config) error {
	return generatePDF(cfg)
}
