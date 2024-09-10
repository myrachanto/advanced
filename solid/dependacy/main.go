package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// ReadFromSource reads data from an io.Reader and prints it
func ReadFromSource(r io.Reader) {
	// Create a buffer to hold the read data
	buffer := make([]byte, 1024)

	// Read the data into the buffer
	n, err := r.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading: %v\n", err)
		return
	}

	// Print the data read from the reader
	fmt.Printf("Data read: %s\n", string(buffer[:n]))
}

func main() {
	// Low-level module 1: String reader
	strReader := strings.NewReader("Hello, String Reader!")
	fmt.Println("Using strings.Reader:")
	ReadFromSource(strReader)

	// Low-level module 2: Buffer reader (bytes.Buffer)
	buf := bytes.NewBufferString("Hello, Buffer Reader!")
	fmt.Println("\nUsing bytes.Buffer:")
	ReadFromSource(buf)

	// Low-level module 3: File reader
	file, err := os.Open("testfile.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Println("\nUsing os.File:")
	ReadFromSource(file)
}
