package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type shape interface {
	Area() float64
	String() string
	AreaString() string
}
type square struct {
	Height float64
	Width  float64
}

func NewSquare(h, w float64) *square {
	return &square{
		Height: h,
		Width:  w,
	}
}

// Area calculates the area only
func (s *square) Area() float64 {
	return s.Height * s.Width
}

// String formats the area ino string
func (s *square) String() string {
	return fmt.Sprintf("The area of the square is %2f \n", s.Area())
}

// wrong way to do it
func (s *square) AreaString() string {
	return fmt.Sprintf("The area of the square is %2f \n", s.Height*s.Width)
}

func main() {
	square := NewSquare(5, 5)
	fmt.Println(square.String())
	fmt.Println(square.AreaString())
}

// Violates SRP: this function has two responsibilities - reading a file and processing data
func readFileAndProcess(path string) error {
	_, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return nil
}

// Adheres to SRP: responsibilities are separated into distinct functions
func readFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func processFileData(data []byte) error {
	// Processing logic here
	return nil
}
