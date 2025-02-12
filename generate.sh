#!/bin/bash

# List of resume files to process
resumes=("FR" "EN")
srcFolder="./src"
outputFolder="./output"
name="Edouard Jubert"
prefix="CV - $name -"

# Process each resume
for lang in "${resumes[@]}"; do
    # Create output folder if it doesn't exist
    if [ ! -d "$outputFolder" ]; then
        mkdir -p "$outputFolder"
    fi

    filename="$prefix $lang"
    echo "Generating $filename.pdf..."
    pandoc "$srcFolder/$lang.md" \
        -f markdown \
        -t pdf \
        --pdf-engine=xelatex \
        -o "$outputFolder/$filename.pdf"
done

echo "Done!"
