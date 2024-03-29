## What is this?

Conway's Game of Life in Golang

## How to run?

```sh
go mod tidy

go run main.go <input_path> <generation> <outputpath>

# Example
go run main.go ./examples/oscillators/beacon_0.txt 10 ./output.txt
```

## Integration tests

`./integration_test.sh` contains a script to test the code with the examples in the `examples` folder.
