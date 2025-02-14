package main

import (
	"fmt"
	"log"
	"os"

	"github.com/edjubert/yaml-to-resume/internal/generator"

	"github.com/spf13/cobra"
)

var (
	configName string
	lang       string
	template string
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "yaml-to-resume",
		Short: "Generate a multilingual resume",
		Run: func(cmd *cobra.Command, args []string) {
			if configName == "" || lang == "" {
				log.Fatal("Please provide --config and --lang")
			}

			err := generator.GenerateResume(configName, template, lang)
			if err != nil {
				fmt.Println("Error generating resume: ", err)
				os.Exit(1)
			}
		},
	}

	rootCmd.Flags().StringVarP(&configName, "config", "c", "", "Configuration name (e.g., 'senior-golang')")
	rootCmd.Flags().StringVarP(&lang, "lang", "l", "en", "Language (e.g., 'en', 'fr')")
	rootCmd.Flags().StringVarP(&template, "template", "t", "default", "Template name (folder name in templates/)")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
