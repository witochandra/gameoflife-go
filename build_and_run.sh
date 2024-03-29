#!/bin/bash

inputPath=$1
generations=$2
outputPath=$3

go run main.go "$inputPath" "$generations" "$outputPath"
