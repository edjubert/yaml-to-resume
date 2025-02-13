package generator

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

type Experience struct {
	Company  string    `mapstructure:"company"`
	Roles    []Role    `mapstructure:"roles"`
}

type Role struct {
	Name     string   `mapstructure:"name"`
	Date     string   `mapstructure:"date"`
	Tasks    []string   `mapstructure:"tasks"`
	Technologies []string `mapstructure:"technologies"`
}


type TemplateConfig struct {
	Font             string `mapstructure:"font"`
	TitleStyle       string `mapstructure:"title_style"`
	DateFormat       string `mapstructure:"date_format"`
	ShowTechnologies bool   `mapstructure:"show_technologies"`
}

type Education struct {
	School     string `mapstructure:"school"`
	Degree     string `mapstructure:"degree"`
	Date       string `mapstructure:"date"`
}

type Language struct {
	Language string `mapstructure:"language"`
	Level    string `mapstructure:"level"`
}

type ResumeConfig struct {
	Info        Info            `mapstructure:"info"`
	Experiences []Experience    `mapstructure:"experiences"`
	Skills      map[string][]string `mapstructure:"skills"`
	Education   []Education       `mapstructure:"education"`
	Languages   []Language        `mapstructure:"languages"`
	Template    TemplateConfig  `mapstructure:"template"`
}

type Info struct {
	Name string `mapstructure:"name"`
	Title string `mapstructure:"title"`
	Email string `mapstructure:"email"`
	ProfilePicture string `mapstructure:"profile_picture"`
	Linkedin string `mapstructure:"linkedin"`
	Malt string `mapstructure:"malt"`
	GitHub string `mapstructure:"github"`
	Location string `mapstructure:"location"`
}

// LoadConfig loads both resume and template configs
func LoadConfig(configName, templateName string) (*ResumeConfig, error) {
	configPath := filepath.Join("configs", configName)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading resume config: %w", err)
	}

	var config ResumeConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error parsing resume config: %w", err)
	}

	// Load template-specific config
	templatePath := filepath.Join("templates", templateName)
	viper.SetConfigFile(filepath.Join(templatePath, "config.yaml"))

	if err := viper.MergeInConfig(); err != nil {
		fmt.Println("No template-specific config found, using defaults")
	}

	var templateConfig TemplateConfig
	if err := viper.Unmarshal(&templateConfig); err != nil {
		return nil, fmt.Errorf("error parsing template config: %w", err)
	}

	config.Template = templateConfig

	return &config, nil
}
