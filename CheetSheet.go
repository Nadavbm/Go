/*
Go cheetsheet
=============
This is a multiline comment
*/

// One line comment

Packages:
=========
// Every code defined by a package name
package main 								// Only one main package for each app
package packageName 						// Use this for import to main

Import modules | libraries:
===========================
import lib
import (
  http
  fmt
  io/ioutil
  net/http
)

Data types:
===========
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte 										// alias for uint8
rune 										// alias for int32 represents a Unicode code point
float32 float64
complex64 complex128

Operators:
==========
+,-,*,/,%									// Addtion, subtraction, multiplication, division,remainder
&&, || ,!									// And, or, not

Fucntions in short:
===================
func main() {
	x := "Hello World"
	fmt.Println(x)
}											// Running main function here returns "Hello World"

Assign Variables:
=================
var varName type 							// Define global variable (data types above)
var var1,var2,var3 string 					// define 3 variables
var vanName int = 325 						// define and initial variable
var var1,var2,var3 int = 1, 2, 3 			// initialize 3 int vars
var var1,var2,var3 = "Pita", "Kebab", "Chips" // initialize 3 vars

func main() {
	varName := value 						// Short form - only within functions
	_, varName := 34, 35 					// Black variable, 34 will be ignored and 35 will assign to varName
}

var (
  str string
  i int
  flo float32
)											// Define a group of variables

Constants:
==========
const constansName = value 					// this value cannot change during compile time
const Pi float32 = 3.1415926
const bestMovieEver = "Die Hard 1" 			// Yes, this is correct until the end of all times and worlds together

const (
  i = 24
  suffix = ".conf"
  pi = 3.1415
)											// Group of constants

Arrays, Slices and Maps:
========================
var arr [n]type
var arr [10]int							 	// Array of integers with 10 elements
arr[0] = 12 								// The first element has the value of 12
arr[5] = 123 								// The sixth element has the value of 123
arr := [3]int {1, 2, 3}
twoArrays := [2][3]int{[3]int{3,6,9},[3]int{1,2,3}} // two dimesional array set
twoArrays := [2][3]int{{3,6,9},{1,2,3}} 	// as 2 arrays of 3 integers each

// Slices:
var slice []int 							// slice of an array
slice := []byte {'a', 'b', 'c'} 			// initialize data in slice
var arrayName = [4]byte {'a', 'b', 'c', 'd'} // Array of 4 elements
a = arrayName[1:3] 							// slice of second, third and forth elements
b = arrayName[:1] 							// The first two elements of the array
c = arrayName[:] 							// All array elements in a Slice

// Maps:
var people map[sting] int 					// Map String to integer - example: name to age
people := make(map[string]int)
people["Maya"] = 25 						// mapping Key-Value
people["Ezra"] = 46
mapName := map[string]float32 {"someThing": 5.432, "otherThing": 43.234, "anotherThing": 123.321} //Initialize map with keys and values

Conditional statements:
=======================

if x > 4 {
  fmt.println("x is greater than 4")
} else {
  fmt.println(" x is less than 4")
}												// if statement

if x:= someInt; x == 3 {
  fmt.println("x is equal to 3")
} else if x < 3 {
  fmt.println("x is less than 3")
} else {
  fmt.println("x is greater than 3")
}												// If, else if, else statement

Loops:
======

func forLoop(){
	sum := 10
	for i=0; i < 10; i++ {
  		sum += i
	}											
}												// For loop

for sum < 100 {
  sum += sum
}												// For which is actually similar to while

func forLoopWithBreak(){
	for i :=10; i>0 ;i-- {
  		if i == 5{
    		break 								// Break or continue for loop (depeneds of condition)
  	}
  	fmt.println(i)
	}
}

for i,g := range map {
  fmt.println ("Key: ", k)
  fmt.println =("Value: ", v)
}												// For loop using a map

for _, v := range map {
  fmt.println("I want to print only values so: ", v)
}												// For loop with _

// A while true statement:
for {
  // Statement logic
}

Switch statements:
==================
func caseExample() {
	i := 10
	switch i {
	case 1:
 		fmt.println("i is equal to 1")
	}
	case 2,3,4:
  		fmt.println("i is equal to 2,3 or 4")
	case 10:
  		fmt.println("i is equal to 10")
	default:
  		fmt.println("i is an integer that doesn't comply with the other CASE statement")
}

func funcWithFallThrough() {
	i := 4
	switch i {
	case 2:
	  fmt.println("i <= 2")
 	fallthrough
	case 4:
 		fmt.println("i <= 4")
	fallthrough
	case 6:
  		fmt.println("i <= 6")
 	fallthrough
	default:
 		fmt.println("default case")
}

Functions:
==========

func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
  return value1, value2
}
package main

import "fmt"

func maxNumber(a,b int) int {
  if a > b {
    return a
  }
  return b
}										// This function will return the bigger number

func main() {
  x := 2
  y := 5

  max_xy := maxNumber(x, y)
  fmt.Printf("%d is the bigger number", maxNumber(x, y))
  fmt.Printf("%d is the bigger number", max_xy)
  fmt.Printf("the numbers were %d and %d", x, y)
}										// Calling the function maxNumber and return it's value

// full simple function:
package main

import "fmt"

func main(){
  sum := 0;
  for i:=0; i < 10; i++ {
    if i%2 == 0 {
      fmt.Println(i, " is even")
    } else {
      sum += i
    }
  }
  fmt.Println("sum of all odd numbers and less than 11 is: ", sum)
}

// Pass a copy of Pointer to a funcion:
package main

import "fmt"

func add1(a *int) int{
  *a = *a + 1 // add 1 to a
  return *a // return the new a
}

func main() {
  x := 3
  fmt.Println("x = ", x) // x is still 3

  x1 := add1(&x) // use add1 function
  fmt.Println("x is now = ", x1) // x should be 4 here
  fmt.Println("x is now = ", x) // x should be 4 here
}

Struct:
=======
package main

import "fmt"

type Person struct {
    name string
    age int
    hobby string
}

func main() {
  Eli := Person{"Eli", 34, "Likes kebab be pita"}
  fmt.Println("Eli is ", Eli.age, "and he ", Eli.hobby)
}

// Method - "a method is a function with an implicit first argument, called a receiver"
func (r ReceiverType) funcName(parameter) (results)

// Method example:
package main

import (
   "fmt"
   "math"
)

type Circle struct {
  radius float64
}

type Rectangle struct {
  width, height float64
}
// method - use Area twice to calculate area of both structs - Circle and Rectangle
func (c Circle) Area() float64 {
  return c.radius * math.Pi
}

func (r Rectangle) Area() float64 {
  return r.width + r.height
}

func main() {
  c1 := Circle{2}
  r1 := Rectangle{3,4}
  r2 := Rectangle{1,2}

  fmt.Println("Area of c1: ", c1.Area())
  fmt.Println("Area of r1: ", r1.Area())
  fmt.Println("Area of r2: ", r2.Area())
}

Concurrency and goroutine:
==========================
go hello(a, b, c)
// Example of goroutine:
package main

import (
    "fmt"
    "runtime"
)

func falafel(str string) {
  for i := 0; i < 5; i++ {
    runtime.Gosched()
    fmt.Println(str)
  }
}

func main() {
  go falafel("Humus") // create a new goroutine
  go falafel("Chips") //
  falafel("Salat") // current go routine
}

