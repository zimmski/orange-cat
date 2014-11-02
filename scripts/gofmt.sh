#!/usr/bin/env sh
NO_COLOR="\033[0m"
OK_COLOR="\033[32;01m"
ERROR_COLOR="\033[31;01m"
WARN_COLOR="\033[33;01m"

files_to_fmt=`gofmt -l .`

if [[ -n "$files_to_fmt" ]]; then
  echo $ERROR_COLOR"Following files haven't yet been gofmt'ed."$NO_COLOR"\n"
  for file in $files_to_fmt; do
    echo "  "$file
  done
  echo
  exit 1
fi
