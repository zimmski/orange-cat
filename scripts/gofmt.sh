#!/usr/bin/env sh
files_to_fmt=`go fmt ./...`

if [ -n "$files_to_fmt" ]; then
  for file in $files_to_fmt; do
    echo "  "$file
  done
  exit 1
fi
