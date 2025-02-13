package main

import (
	"fmt"
	"os"

	"github.com/edjubert/resume-generator/internal/generator"
	"github.com/spf13/cobra"
)

var config, lang, template string

var rootCmd = &cobra.Command{
	Use: "resume-generator",
	Short: "Generate a resume in multiple languages and formats",
	Long: "A CLI tool to generate a resume in multiple languages and formats from a config file",
	Run: func(cmd *cobra.Command, args []string) {
		err := generator.GenerateResume(config, template, lang)
		if err != nil {
			fmt.Println("Error generating resume: ", err)
			os.Exit(1)
		}

		fmt.Println("Resume generated successfully")
	},
}

func init() {
	rootCmd.Flags().StringVarP(&config, "config", "c", "default", "Config name (folder name in configs/)")
	rootCmd.Flags().StringVarP(&lang, "lang", "l", "en", "Language to generate the resume in (e.g., en, fr, es)")
	rootCmd.Flags().StringVarP(&template, "template", "t", "default", "Template name (folder name in templates/)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing command: ", err)
		os.Exit(1)
	}
}

