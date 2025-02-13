package generator

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func RenderMarkdown(config *ResumeConfig, templateName, lang, configName string) (string, error) {
	tmplPath := filepath.Join("templates", templateName, "template.md")
	
	// Create new template with functions first
	tmpl := template.New(filepath.Base(tmplPath)).Funcs(template.FuncMap{
		"join": strings.Join,
	})
	
	// Parse the template file
	tmpl, err := tmpl.ParseFiles(tmplPath)
	if err != nil {
		return "", fmt.Errorf("error loading template: %w", err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, config)

	if err != nil {
		return "", fmt.Errorf("error rendering template: %w", err)
	}

	outputDir := filepath.Join("tmp", configName)
	os.MkdirAll(outputDir, os.ModePerm)

	outputFile := filepath.Join(outputDir, lang+".md")
	err = os.WriteFile(outputFile, buf.Bytes(), 0644)
	if err != nil {
		return "", fmt.Errorf("error writing markdown file: %w", err)
	}

	return outputFile, nil
}