#!/bin/bash

# Function to create the initial content for file.go
create_pkg_go() {
  cat <<EOL >"$1/$2.go"
package $2

// SOLUTION

func Solve(input string) string {
    return input
}
EOL
}

# Function to create the initial content for file_test.go
create_pkg_test_go() {
  cat <<EOL >"$1/$2_test.go"
package $2

import "testing"

func TestSolve(t *testing.T) {
    input := "123"
    expected := "123"
    if result := Solve(input); result != expected {
        t.Errorf("Solve() = %v, want %v", result, expected)
    }
}
EOL
}

# Function to create the initial content for main.go
create_main_go() {
  cat <<EOL >"$1/main.go"
package main

import (
	"fmt"

	"github.com/ayushgpt01/advent-of-code-golang/input"
	"github.com/ayushgpt01/advent-of-code-golang/pkg/$2"
)

func main() {
	input, err := input.Read("$2.txt")
	if err != nil {
		panic(err)
	}
	result := $2.Solve(input)
	fmt.Println(result)
}
EOL
}

# Check if a file name was provided
if [ -z "$1" ]; then
  echo "Usage: $0 <file_name>"
  exit 1
fi

# File name
file=$1

# Create pkg/file directory and files
mkdir -p pkg/$file
create_pkg_go "pkg/$file" $file
create_pkg_test_go "pkg/$file" $file

# Create cmd/file directory and files
mkdir -p cmd/$file
create_main_go "cmd/$file" $file

echo "Files for $file created successfully."
