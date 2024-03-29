#!/bin/bash

mkdir -p ./tmp

run_test() {
    local name="$1"
    local input="$2"
    local expected="$3"
    local generation="$4"

    if [ -z "$generation" ]; then
        generation=1
    fi

    echo "-----------------------------------"
    echo "Running test: $name"
    ./build_and_run.sh "$input" "$generation" ./tmp/output.txt

    if diff "$expected" "./tmp/output.txt" >/dev/null 2>&1; then
        echo "Pass"
    else
        echo "Fail"
    fi
}


# still lifes
run_test "block" "./examples/still_lives/block.txt" "./examples/still_lives/block.txt"
run_test "bee-hive" "./examples/still_lives/bee-hive.txt" "./examples/still_lives/bee-hive.txt"
run_test "loaf" "./examples/still_lives/loaf.txt" "./examples/still_lives/loaf.txt"
run_test "boat" "./examples/still_lives/boat.txt" "./examples/still_lives/boat.txt"
run_test "tub" "./examples/still_lives/tub.txt" "./examples/still_lives/tub.txt"

# oscillators
run_test "blinker 0" "./examples/oscillators/blinker_0.txt" "./examples/oscillators/blinker_1.txt"
run_test "blinker 1" "./examples/oscillators/blinker_1.txt" "./examples/oscillators/blinker_0.txt"

run_test "toad 0" "./examples/oscillators/toad_0.txt" "./examples/oscillators/toad_1.txt"
run_test "toad 1" "./examples/oscillators/toad_1.txt" "./examples/oscillators/toad_0.txt"

run_test "beacon 0" "./examples/oscillators/beacon_0.txt" "./examples/oscillators/beacon_1.txt"
run_test "beacon 1" "./examples/oscillators/beacon_1.txt" "./examples/oscillators/beacon_0.txt"

# spaceships
run_test "glider 0" "./examples/spaceships/glider_0.txt" "./examples/spaceships/glider_1.txt"
run_test "glider 1" "./examples/spaceships/glider_1.txt" "./examples/spaceships/glider_2.txt"
run_test "glider 2" "./examples/spaceships/glider_2.txt" "./examples/spaceships/glider_3.txt"
run_test "glider 3" "./examples/spaceships/glider_3.txt" "./examples/spaceships/glider_4.txt"
run_test "glider 100" "./examples/spaceships/glider_0.txt" "./examples/spaceships/glider_100.txt" 100


run_test "lwws 0" "./examples/spaceships/lwws_0.txt" "./examples/spaceships/lwws_1.txt"
run_test "lwws 1" "./examples/spaceships/lwws_1.txt" "./examples/spaceships/lwws_2.txt"
run_test "lwws 2" "./examples/spaceships/lwws_2.txt" "./examples/spaceships/lwws_3.txt"
run_test "lwws 3" "./examples/spaceships/lwws_3.txt" "./examples/spaceships/lwws_4.txt"
run_test "lwws 100" "./examples/spaceships/lwws_0.txt" "./examples/spaceships/lwws_100.txt" 100

rm -rf ./tmp
