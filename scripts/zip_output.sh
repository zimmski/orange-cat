#!/usr/bin/env sh
ROOT=$PWD
for dir in `ls out`; do
  cd out/$dir
  zip "../orange_"$dir".zip" *
  cd $ROOT
done
