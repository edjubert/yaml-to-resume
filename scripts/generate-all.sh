#!/usr/bin/env bash

# Generate all resumes
for config in $(ls configs); do
    for template in $(ls templates); do
        for lang in en fr; do
            echo "Generating ./output/$config/$template/$lang.pdf"
            ./bin/yaml-to-resume cmd/main.go --config=$config --lang=$lang --template=$template 2>&1 >/dev/null || exit 1
        done
    done
done
