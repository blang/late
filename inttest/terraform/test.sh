#!/bin/bash
cat data.json | late render -t terraform.tmpl
cat data.yaml | late render -f yaml -t terraform.tmpl
cat data.yaml | late render -f yaml -t ./terraform.tmpl
