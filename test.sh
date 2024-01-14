#!/bin/bash

echo "Testing against input files..."
for file in ./examples/*; do
    echo "Testing with input file: $file"
    
    TIME_TAKEN=$( { time go run . "$file" ; } 2>&1 | grep real | awk '{print $2}')
    output=$(go run . "$file")
    echo "$output"
    echo "----------------------------------"
    
    spaces_count=$(echo "$output" | grep -o '\.' | wc -l)
    echo "Number of spaces in output: $spaces_count"
    echo "Time taken: $TIME_TAKEN"
    echo "----------------------------------"
done

echo "All tests completed."