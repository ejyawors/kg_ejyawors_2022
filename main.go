package main
import "fmt"
import "os"

// Question: Convert an array of integers into an array of strings representing the phonetic equivalent of the
// integer.

func main() {
  x := []string{"Zero","One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"} // list used for indexing; ex: x[1] == One
  num_args := len(os.Args[1:]) // length of arguments inputted

  a := os.Args[1] // Getting command line argument
  fmt.Println(num_args, x[a[0] - 48]) // Test if string, or rune, index is ascii

}

// references:
// https://gobyexample.com/
