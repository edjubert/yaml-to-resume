package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

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

type Skill struct {
	Label string `mapstructure:"label"`
	Items []string `mapstructure:"items"`
}

type ResumeConfig struct {
	Info        Info            `mapstructure:"info"`
	Experiences []Experience    `mapstructure:"experiences"`
	Skills      map[string]Skill `mapstructure:"skills"`
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

type Role struct {
	Name        string   `mapstructure:"name"`
	Date         string   `mapstructure:"date"`
	Tasks        []string `mapstructure:"tasks"`
	Technologies []string `mapstructure:"technologies"`
}

type Experience struct {
	Company string `mapstructure:"company"`
	Roles   []Role `mapstructure:"roles"`
}

// LoadConfig loads the base config and merges translations
func LoadConfig(configName, templateName, lang string) (*ResumeConfig, error) {
	configPath := filepath.Join("configs", configName)

	// Load base config.yaml
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading base config: %w", err)
	}

	var config ResumeConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error parsing base config: %w", err)
	}

		// Load template-specific config
		templatePath := filepath.Join("templates", templateName)
		viper.SetConfigFile(filepath.Join(templatePath, "config.yaml"))


	// Load language-specific overrides
	langFile := filepath.Join(configPath, lang+".yaml")
	if _, err := os.Stat(langFile); err == nil {
		viper.SetConfigFile(langFile)
		if err := viper.MergeInConfig(); err != nil {
			return nil, fmt.Errorf("error merging translation file: %w", err)
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error parsing merged config: %w", err)
	}

	return &config, nil
}
