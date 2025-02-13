package generator

func GenerateResume(configName, templateName, lang string) error {
	config, err := LoadConfig(configName, templateName)
	if err != nil {
		return err
	}

	mdFile, err := RenderMarkdown(config, templateName, lang, configName)
	if err != nil {
		return err
	}

	err = GeneratePDF(mdFile, configName, lang)
	if err != nil {
		return err
	}

	return nil
}
