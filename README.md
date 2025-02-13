# Resume Generator

A command-line tool written in Go that generates professional PDF resumes from YAML configurations using customizable LaTeX templates.

## Features

- Generate PDF resumes from YAML configuration files
- Support for multiple languages
- Customizable LaTeX templates
- Profile picture support
- Automatic PDF generation via GitHub Actions
- Cache support for TeX Live packages

## Prerequisites

- Go 1.21 or higher
- TeX Live with XeTeX
- Pandoc

## Installation

```bash
git clone https://github.com/edjubert/resume-generator
cd resume-generator
go mod download
```

## Usage

```bash
go run cmd/main.go --config=default --lang=en --template=default
```

### Command Line Options

- `--config, -c`: Config name (folder name in configs/) (default: "default")
- `--lang, -l`: Language to generate the resume in (default: "en")
- `--template, -t`: Template name (folder name in templates/) (default: "default")

## Project Structure

```
.
├── .github/
│ └── workflows/ # GitHub Actions workflow definitions
├── cmd/
│ └── main.go # Main entry point
├── configs/ # Resume configurations
│ └── default/
│ ├── assets/ # Images and other assets
│ └── config.yaml # Resume content configuration
├── internal/
│ └── generator/ # Core generation logic
├── templates/ # LaTeX templates
│ └── default/
│ ├── template.md # Template file
│ └── config.yaml # Template configuration
└── tmp/ # Temporary files
```

## Configuration

Create a YAML configuration file in the `configs/` directory. See the example configuration:

```yaml:README.md
info:
    name: "John Doe"
    title: "Senior Software Engineer"
    email: "john.doe@example.com"
# ... other fields
experiences:
    company: "Tech Corp"
    roles:
        name: "Senior Software Engineer"
        date: "2022 - Present"
        tasks:
            - "Led backend development"
        technologies:
            - "Golang"
            - "PostgreSQL"
# ... other fields
education:
  - school: "Technical University"
    degree: "Bachelor of Science in Computer Science"
    date: "2016 - 2020"

languages:
  - language: "English"
    level: "Native"
  - language: "French"
    level: "Native"
```

Then select the config you want to use in the command line.

```bash
go run cmd/main.go --config=<config_name> --lang=<language> --template=<template_name>
```

You can also create a new template by creating a new folder in the `templates/` directory and adding a `template.md` file.

## GitHub Actions

The project includes automatic PDF generation using GitHub Actions. On every push to the main branch, it will:

1. Set up the Go environment
2. Install TeX Live dependencies
3. Generate the PDF
4. Upload the result as an artifact

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
