package main

import (
	"fmt"

	"homework/solution"
)

func main() {
	// Example 1
	job := &solution.MergeDictsJob{}
	result, err := solution.ExecuteMergeDictsJob(job)
	fmt.Printf("Example 1: result=%+v, error=%v\n", result, err)

	// Example 2
	job = &solution.MergeDictsJob{Dicts: []map[string]string{{"a": "b"}, nil}}
	result, err = solution.ExecuteMergeDictsJob(job)
	fmt.Printf("Example 2: result=%+v, error=%v\n", result, err)

	// Example 3
	job = &solution.MergeDictsJob{Dicts: []map[string]string{{"a": "b"}, {"b": "c"}}}
	result, err = solution.ExecuteMergeDictsJob(job)
	fmt.Printf("Example 3: result=%+v, error=%v\n", result, err)
}
