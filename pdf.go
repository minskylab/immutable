package immutable

import (
	"path"

	"github.com/k0kubun/pp"
	"github.com/pkg/errors"
	"github.com/thecodingmachine/gotenberg-go-client/v7"
)

func generatePDF(cfg *Config) error {
	pp.Println(cfg)

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

	resultFilename := path.Join(cfg.ImmutableDir, "result.pdf")

	if err := c.Store(req, resultFilename); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func GeneratePDF(cfg *Config) error {
	return generatePDF(cfg)
}
