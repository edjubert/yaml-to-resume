# Resume Generator

Automatically generate PDF versions of my resume in multiple languages using Pandoc and LaTeX.

## Overview

This project contains my resume in both French and English, written in Markdown with LaTeX formatting. It uses GitHub Actions to automatically generate PDF versions whenever changes are pushed to the main branch.

## Structure

```
.
├── src/
│ ├── FR.md # French version
│ └── EN.md # English version
├── assets/ # Images and icons
│ ├── email.png
│ ├── github.png
│ ├── linkedin.png
│ ├── malt.png
│ └── me.png
├── output/ # Generated PDFs
│ ├── CV - Edouard Jubert - FR.pdf
│ └── CV - Edouard Jubert - EN.pdf
└── .github/
└── workflows/ # GitHub Actions workflow
```

## Features

- Multi-language support (FR/EN)
- Automatic PDF generation
- Professional LaTeX formatting
- Consistent styling across versions
- Version control for both source and output

## Local Development

### Prerequisites

- Pandoc
- XeLaTeX
- DejaVu Sans font

### Generate PDFs Locally

```bash
# Clone the repository
git clone https://github.com/edjubert/resume.git
cd resume

# Generate PDFs
./generate.sh
```

## GitHub Actions

The repository includes a GitHub Actions workflow that automatically:

1. Generates PDF versions of both resumes
2. Uploads the PDFs as artifacts
3. Commits the generated PDFs back to the repository

## License

This project is licensed under the [MIT License](LICENSE) - Feel free to use this template for your own resume!

## Contact

Edouard Jubert

- Email: edouard.jubert@proton.me
- GitHub: [@edjubert](https://github.com/edjubert)
- LinkedIn: [Edouard Jubert](https://www.linkedin.com/in/edouard-jubert-9a348b58/)
