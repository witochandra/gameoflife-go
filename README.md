## What is this?

Conway's Game of Life in Golang

## How to run?

```sh
go mod tidy

go run main.go

# run with input file
go run main.go \
    -render=true \
    -input=examples/input.txt

# run without rendering & output the last state to a file
go run main.go \
    -render=false \
    -input=examples/input.txt \
    -output=output.txt \
    -steps=10
```
