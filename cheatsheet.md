
## go cheatsheet

```
/*
  This is a multiline comment
*/
```

``` // One line comment ```

### packages

write package name in the beginning of the code. this will be the name of the package to export.

``` package main ```								 -   Only one main package for each app.

``` package packageName ``` 				 -	 Use to import to other packages

### import packages

``` import pkg ```

```
import (
  fmt
  io/ioutil
  net/http

  github.com/some/package
  nameOfPackage github.com/some/package
)
```

### using packages

```fmt.Println("I'm using fmt package")``` use packges by the name (according to `import` packages)

### data types:

in go, everything is a TYPE!

primitive data type \ built-in types

``` bool ```      -     boolean: `true` or `false`)

``` string ```    -     string: ```"hello world"``` 

``` int  int8  int16  int32  int64 ```  -     integer: numeric value, a whole number `45`

``` uint uint8 uint16 uint32 uint64 uintptr ```     -     unsigned integer

``` byte ```							-		      alias for uint8

``` rune ```							-			    alias for int32 represents a Unicode code point

``` float32 float64 ```   -     float: numeric value with decimal `0.123`

``` complex64 complex128 ``` 

composite data-types:

array, slice, struct, pointer, function, interface, map (examples below)

### variadic parameters

every value is also type of ```interface{}```

```...interface{}``` means, this will accept as many values of any type.

```fmt.Println("String", 53, 123.123, "Another string")``` 
the definition of Println is `func Println(a ...interface{}) (n int, err error)`

### operators:

``` +,-,*,/,%	```								    -  Addtion, subtraction, multiplication, division,remainder

``` &&, || ,!	```								    -  And, or, not

##### logical operators in use

```
	fmt.Printf("true and false\t %v\n", true && false)
	fmt.Printf("true or true\t %v\n", true || true)
	fmt.Printf("not true\t\t %v\n", !true)
```

### variables

``` var varName type ``` 							-             define global variable (data types above)

``` var var1,var2,var3 string ```			-	            define 3 variables. the values of these vars is zero, since not initialized

```var vanName int = 325 ```					-	            define and initialize variable (this will add value instead of zero value)

``` var var1,var2,var3 int = 1, 2, 3 ```	-	    	  initialize 3 int vars

``` var var1,var2,var3 = "Pita", "Kebab", "Chips" ``` -  initialize 3 vars

short declaration of a variable used only within `{}`:
```
func main() {
	varName := value
	_, varName := 34, 35
}
```

```_``` is a blank variable, value that can return nothing

group of variables:

```
var (
  str string
  i int
  flo float32
)
```

in go, everything is a TYPE! to create your own type:
```
type kebab int
var b kebab
```
this will create a type called `kebab` which is based on `int`

### conversions

what's known as casting in other languages, is conversion in go - convert a type to another type

conversion written like this: `a = string(b)`. full example of using conversion:
```
var a int

type kebab int
var b kebab

func main() {
	a = 42
	b = 43
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	a = int(b)
	fmt.Printf("%T\n", a)
}
```

### constants:

``` const constansName = value ```		    - 	 this value cannot change during compile time

``` const Pi float32 = 3.1415926 ```

``` const bestMovieEver = "Die Hard 1" ``` 

group of constants

```
const (
  i = 24
  suffix = ".conf"
  pi = 3.1415
)	
```

#### iota and bit shifting

iota and shifting binary digit to the left `<<` or right `>>` side
```
const (
  _ = iota
  kb = 1 << (iota * 10)
  mb = 1 << (iota * 10)
  gb = 1 << (iota * 10)
)

func main() {
  fmt.Printf("%d\t\t\t%b\n%d\t\t\t%b\n%d\t\t\t%b\n", kb ,kb, mb,mb,gb,gb)
}

```

### flow control: loops and conditional statements

In computer science, control flow (or flow of control) is the order in which individual statements, instructions or function calls of an imperative program are executed or evaluated.
Reading from top to bottom, left to right, excuting conditional statements, loops and vars.

### for loops 

initialize, condition, post:

```
	sum := 10
	for i=0; i < 10; i++ {
  		sum += i
	}											

```

the `while` keyword doesn't exist in go, instead use `for`:
```
  for sum < 100 {
    sum += sum
  }
```

##### nested loops

```
	for i := 0; i <= 10; i++ {
    fmt.Printf("outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("inner loop: \t\t%d\n", i, j)
		}
	}
```

##### the for statement

```
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("for a condintion of x less than 10, print x and add x += 1")
```

```
	x := 1
	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("we will break the for loop after the condition fulfilled - x won't be bigger than 9")
```

##### break and continue

```
x := 1
	for {
		x++
		if x > 100 {
			break
		}
		if x%2 != 1 {
			continue
		}
		fmt.Println(x is not an equal number)	
	}
```

#### conditional statement 

##### if statement

```
  if x > 4 {
    fmt.println("x is greater than 4")
  } else {
    fmt.println(" x is less than 4")
  }
```

##### if, else if, else

```
  if x:= someInt; x == 3 {
    fmt.println("x is equal to 3")
  } else if x < 3 {
    fmt.println("x is less than 3")
  } else {
    fmt.println("x is greater than 3")
  }	
```

#### switch statements:

switch based on true or false

```
	switch {
	case false:
		fmt.Println("wont print")
	case (2 == 4):
		fmt.Println("wont print again")
	case (3 == 3):
		fmt.Println("prints")
	case (12 == 12):
		fmt.Println("this is true")
		fallthrough
	case (9 != 9):
		fmt.Println("not true")
		fallthrough
	default:
		fmt.Println("print always default")
	}
```

switch based on value:

```
	k := "kebab"
	switch k {
	case "hummus":
		fmt.Println("not kebab")
	case "shishlik":
		fmt.Println("still not kebab")
	case "kebab":
		fmt.Println("kebab")
	default:
		fmt.Println("default food - like prison")
	}
```


### arrays, slices and maps:

arrays cannot be changed in size

``` var arr [n]type ```

``` var arr [10]int	```				          - 		array of integers with 10 elements

``` arr[0] = 12 ```						     	    -   	the first element has the value of 12

``` arr[5] = 123 	```					    	    -   	the sixth element has the value of 123

``` arr := [3]int {1, 2, 3} ```

``` twoArrays := [2][3]int{[3]int{3,6,9},[3]int{1,2,3}} ```     -    two dimesional array set

``` twoArrays := [2][3]int{{3,6,9},{1,2,3}} ``` 	-    as 2 arrays of 3 integers each

``` var weekdays []string ```

``` weekdays := [...]string{"Monday","Tuesday","Wednesday","Thursday","Friday"} ```

``` fmt.Println(weekdays[0]) ```			-		print Monday

``` fmt.Println(weekdays[3]) ```			-		print Thursday

array of integers, adding 2 in last position
```
	var x [5]int
	fmt.Println(x)
	x[4] = 2
	fmt.Println(x)
	fmt.Println(len(x))
```  

### slices:

slice use underline array and can expand. always use slice instead of arrays

``` var slice []int ```		-		slice of an array

``` slice := []byte {'a', 'b', 'c'} ```	    - initialize data in slice

``` var arrayName = [4]byte {'a', 'b', 'c', 'd'} ```      -     aray of 4 elements

``` a = arrayName[1:3] ```				-		slice of second, third and forth elements

``` b = arrayName[:1] 	```				-		the first two elements of the array

``` c = arrayName[:] 		```				-	all array elements in a Slice

slice of strings
```
  menu := [...]string{"Falafel", "Hummus", "Shawarma"}
  veggimenu := menu[0:1]
  fmt.Println(veggimenu)
```

slice of integers
```
	slc := []int{4, 5, 7, 8, 9}
	fmt.Println(len(slc))
	fmt.Println(slc)
	fmt.Println(slc[1])
	fmt.Println(slc[1:])
	fmt.Println(slc[1:3])
	
  slc = append(slc, 77, 88, 99, 1014)

	slc2 := []int{234, 456, 678, 987}
	slc = append(slc, slc2...)
	fmt.Println(slc)

	slc = append(slc[:2], slc[4:]...)
	
	for i, v := range slc {
		fmt.Println(i, v)	
	}
```

use `make` to define the slice and the underline array by size (of the slice) and capacity (of the underline array)

```
	slc := make([]int, 10, 12)
	fmt.Println(slc)
	fmt.Println(len(slc))
	fmt.Println(cap(slc))
	
	slc = append(slc, 13, 99 ,8021)
	
	fmt.Println(slc)
	fmt.Println(len(slc))
	fmt.Println(cap(slc))
```

```shish := [][]string{kebab, cebab}``` - multi-dimensional slice

### maps:

map is a key-value store. syntax for a map is map[key]value. maps are unordered

``` var people map[string]int ```			- 	map string to integer - example: name to age

``` people := make(map[string]int) ```

``` people["Maya"] = 25 	 ```	-  	mapping Key-Value

``` people["Ezra"] = 46 ```

``` mapName := map[string]float32 {"someThing": 5.432, "otherThing": 43.234, "anotherThing": 123.321} ``` -  

using `for` over map:
```
  for i,g := range map {
    fmt.println ("Key: ", k)
    fmt.println =("Value: ", v)
  }
``` 

```
  for _, v := range map {
    fmt.println("I want to print only values so: ", v)
  }	
```

initialize map with keys and values

```
  menu := map[string]int{
  	"Falafel":	15,
  	"Hummus":	18,
  	"Shawarma":	32,
  }

  fmt.Println(menu["Hummus"])
  
  v, ok := menu["Hummus"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := menu["Hummus"]; ok {
		fmt.Println(v)

  menu["Chips"] = 12

  delete(menu, "Shawarma")
```

### structs:

struct is a composite data type which allow us to composing values of different types.

```
  type Preson struct {
  	name string
  	age int
  }
  
  var ish Person

  shimon := Person{
      name: "Shimon", 
      age: 54
  }

  ruhama := Person {
      name: "Ruhama",
      age: 57
  }

  fmt.Println(shimon)
  fmt.Println(shimon.age, ruhama.age)
```

embedded structs:

```
type Person struct {
	name string
	age  int
}

type TruckDriver struct {
	Person
	truckd 	bool
}

func main() {
	truckDriver := TruckDriver{
		Person: Person{
			name: "Sami HaKabai",
			age:  36,
		},
		truckd: true,
	}
	fmt.Println(truckDriver, truckDriver.age, truckDriver.truckd)
}
```

##### anonymous struct

```
	person := struct {
		name    string
		age		int
		friends map[string]int
		food	[]string
	}{
		name: 	 "Shimon",
		age:	 32,
		friends: map[string]int{
			"Simha": 	55,
			"Izik":     27,
			"Betuel":   38,
		},
		food: []string{
			"Kebab",
			"Hummus",
		},
	}
	fmt.Println("the name is ", person.first, ", age of ", person.age, " likes to eat - ", person.food)

	for k, v := range person.friends {
		fmt.Println("friend of %s", person.name, "is ", k, v, " years old")
	}

	for i, val := range person.food {
		fmt.Println(i, val)
	}
```

### functions:

a function is a group of statement that execute in the program. function gets an input value and return a different value according the function. We create function to arrange code and for code reusability.
function will come in this syntax:
```
func (r receiver) funcName(parameters) (return value) {
  code ..
  return value
}
```
`func` keyword, receiver (usually type struct in methods), input parameters (of any type) and return value (of any type).

basic function (no parameter nor return values):
```
func main() {
  sum := 0
  for i:=0; i < 10; i++ {
    if i%2 == 0 {
      fmt.Println(i, " is even")
    } else {
      sum += i
    }
  }
  fmt.Println("sum of all odd numbers and less than 11 is: ", sum)
}
```

function with parameter and return value, execution in main function:
```
func main() {
	str := pho("falafel")
	fmt.Println(str)
}

func pho(str string) string {
	return fmt.Sprint("eat more ", str)
}
```

func that takes two input parameters and return two values:
```
func main() {
	i, boo := calculator([]int{3,2,4,5,6,7,9,12}, 3)
	fmt.Println(i, boo)
}

func calculator(numbers []int, mod int) (int, bool) {
	var moo bool
	sum := 0
	for i := range numbers {
		if i%mod == 0 {
			fmt.Println(i, "divided well by", mod)
		}
		sum += i
	}
	if mod%2 == 0 {
		fmt.Println(moo, "is even number")
		moo = true
	} else {
		moo = false
	}
	return sum, moo
}
```

function with variadic parameter:
```
func pho(x ...int) int {
	fmt.Println(x)
	sum := 0
	for i := range x {
		sum += i
	}
	return sum
}
```

#### defer

A defer statement defers the execution of a function until the surrounding function returns.

```
func main() {
	defer someFunc()
	otherFunc()
}
```

##### anonymous function

anonymous func can be used inside a func
```
func main() {
	func() {
		fmt.Println("Anonymous")
	}()
}
```

assign func to a variable, print all odd numbers
```
func main() {
	g := func() {
		for i:=0; i<100; i++ {
			if i % 2 != 0 {
				fmt.Println(i)
			}
		}
	}
	g()
	fmt.Printf("%T", g)
}
```

anonymous func that accept integer
```
func main() {
	f := 2
	g := func(i int) int {
		return i * 2
	}(f)
	fmt.Println("I add double: ", g)
}
```

##### a function which return function

```
func main() {
	g := funciturn()
	fmt.Println(g())
	fmt.Printf("%T", g)
}

func funciturn() func() string {
	return func() string {
		return "func from a func. funci"
	}
}
```

#### callback

passing a func as an arguement



### method 

a method is a function with an implicit first argument, called a receiver

```func (r ReceiverType) funcName(parameter) (results)```

method to calculate area of circle\square:

```
package main

import (
	"fmt"
)

const pi = 3.14

type Circle struct {
	radius float64
}

type Square struct {
	length float64
	wide   float64
}

func main() {
	circle := Circle{radius: 3.2}
	square := Square{length: 4, wide: 6}
	c := circle.area()
	s := square.area()
	fmt.Println(c, "\t", s)
}

func (c Circle) area() float64 {
	return (c.radius / 2) * pi
}

func (s Square) area() float64 {
	return s.length * s.wide
}
```

### interface

interfaces in Go provide a way to specify the behavior of an object. 
Interfaces are named collections of method signatures.

```
package main

import (
	"fmt"
)

const pi = 3.14

type Circle struct {
	radius float64
}

type Square struct {
	length float64
	wide   float64
}

type Shape interface {
	area() float64
}

func main() {
	circle := Circle{radius: 3.2}
	square := Square{length: 4, wide: 6}
	outputSize(circle)
	outputSize(square)
}

func outputSize(s Shape) {
	fmt.Println(s.area())
}

func (c Circle) area() float64 {
	return (c.radius * c.radius) * pi
}

func (s Square) area() float64 {
	return s.length * s.wide
}
```

### pointers

values stored in memory. a pointer is a memory address and a point is a value stored in memory.

```
	x := 12
	fmt.Println("the value of x: ", x, "\nx var in memory address: ", &x)
	fmt.Printf("type of x: %T\nx pointer type is %T",x , &x)

	var y = &x
	fmt.Println("\nvalue of y", y, "\nvalue of *y: ", *y)
	fmt.Printf("what happend to y when assign to a pointer: %T\n", y)
	fmt.Printf("what is the type of &x now: %T\n", &x)
```

full example (look at output of fmt):

```
package main

import "fmt"

func add1(a *int) int{
  	*a = *a + 1
	fmt.Println("memory address of a: ", a)
	fmt.Printf("a type is: %T\n*a type is %T\n", *a, a)
  	return *a
}

func main() {
  	x := 3
  	fmt.Println("the value if x is: ", x, "the memory address of &x: ", &x)
  	x1 := add1(&x)
  	fmt.Println("after using x as a pointer in add one function: ", x1)
	fmt.Printf("the type of x: %T\nthe type of x1: %T", x, x1)
}
```










Some more examples:

#### pass a copy of Pointer to a funcion:
```
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
```


##### method example:
```
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
```

### concurrency and goroutine:

``` go hello(a, b, c) ```

Example of goroutine:
```
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
```
