/*
A tuple is a finite sorted list of elements. It is a data structure that groups data. Tuples are
typically immutable sequential collections. The element has related fields of different
datatypes. The only way to modify a tuple is to change the fields. Operators such as + and *
can be applied to tuples.
*/
package main

import "fmt"

func powerSeries(a int) (int, int) {
	return a * 2, a * 3
}

// Tuples can be named inside the func, as shown below
func powerSeriesTwo(a int) (square int, cube int) {
	square = a * a
	cube = square * a
	return
}

func main() {
	var square int
	var cube int
	square, cube = powerSeries(3)
	fmt.Println("Square: ", square, "Cube: ", cube)
}
