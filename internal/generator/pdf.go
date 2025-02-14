package generator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func GeneratePDF(markdownFile, configName, templateName, lang string) error {
	outputDir := filepath.Join("output", configName, templateName)
	os.MkdirAll(outputDir, os.ModePerm)

	pdfFile := filepath.Join(outputDir, lang+".pdf")
	args := []string{markdownFile, "-o", pdfFile, "--pdf-engine", "xelatex", "--from", "markdown"}	

	cmd := exec.Command("pandoc", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error generating PDF: %w", err)
	}

	return nil
}
