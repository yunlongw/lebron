package main

import "fmt"

// Writer interface wraps Write method.
type Writer interface {
	Write(v interface{})
}

// ConsoleWriter is an implementation of Writer, used for writing to console.
type ConsoleWriter struct{}

// Write writes a value to the console.
func (w ConsoleWriter) Write(v interface{}) {
	fmt.Println(v)
}

// FileWriter is an implementation of Writer, used for writing to file.
type FileWriter struct{}

// Write writes a value to the file.
func (w FileWriter) Write(v interface{}) {
	// Write value to file here...
	fmt.Printf("Writing %v to file...\n", v)
}

func main() {
	// Create a new console writer
	consoleWriter := ConsoleWriter{}

	// Write a string to console using console writer
	consoleWriter.Write("Hello, world!")

	// Create a new file writer
	fileWriter := FileWriter{}

	// Write a string to file using file writer
	fileWriter.Write("Hello, world!")
}
