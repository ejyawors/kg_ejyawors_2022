package main
import "fmt"
import "os"

// Question: Convert an array of integers into an array of strings representing the phonetic equivalent of the
// integer.

func main() {
  x := []string{"Zero","One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"} // list used for indexing; ex: x[1] == One
  result := ""

  for _, i := range os.Args[1:]{ // read each argument in command line
    for j := 0; j < len(i); j++{ // for each element in array, convert to numerical word
      result = result + (x[i[j] - 48]) // convert from ascii to integer, access numerical word from list x, append to string 
    }
    result = result + "," // appebnd comma to string
  }
  result = result[:len(result)-1] // slice away final comma in string
  fmt.Println(result) // print result
}

// references:
// https://gobyexample.com/

// test code: ***Works!***
// go run main.go 10 300 5
// OneZero,ThreeZeroZero,Five
// go run main.go 3 25 209
// Three,TwoFive,TwoZeroNine
