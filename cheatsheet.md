
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

### conversions

what's known as casting in other languages, is conversion in go - convert a type to another type



### variadic parameters

every value is also type of ```interface{}```

```...interface{}``` means, this will accept as many values of any type.

```fmt.Println("String", 53, 123.123, "Another string")``` 
the definition of Println is `func Println(a ...interface{}) (n int, err error)`

### operators:

``` +,-,*,/,%	```								    -  Addtion, subtraction, multiplication, division,remainder

``` &&, || ,!	```								    -  And, or, not

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

arrays:

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

### slices:

``` var slice []int ```		-		slice of an array

``` slice := []byte {'a', 'b', 'c'} ```	    - initialize data in slice

``` var arrayName = [4]byte {'a', 'b', 'c', 'd'} ```      -     aray of 4 elements

``` a = arrayName[1:3] ```				-		slice of second, third and forth elements

``` b = arrayName[:1] 	```				-		the first two elements of the array

``` c = arrayName[:] 		```				-	all array elements in a Slice

```
menu := [...]string{"Falafel", "Hummus", "Shawarma"}
veggimenu := menu[0:1]
fmt.Println(veggimenu)
```

### maps:

``` var people map[sting] int ```			- 	map string to integer - example: name to age

``` people := make(map[string]int) ```

``` people["Maya"] = 25 	 ```	-  	mapping Key-Value


``` people["Ezra"] = 46 ```

``` mapName := map[string]float32 {"someThing": 5.432, "otherThing": 43.234, "anotherThing": 123.321} ``` -  initialize map with keys and values

```
menu := map[string]int{
	"Falafel":	15,
	"Hummus":	18,
	"Shawarma":	32,
}
fmt.Println(menu["Hummus"])
```

### structs:
```
type Preson struct {
	name string
	age int
}
```

``` var ish Person ```

``` shimon := Person{name: "Shimon", age: 54} ```

### conditional statements:

#### switch statements:
```
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
```
```
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
```
### loops:

for loop example (init, condition, post):
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

```
func forLoopWithBreak(){
	for i :=10; i>0 ;i-- {
  		if i == 5{
    		break 								// Break or continue for loop (depeneds of condition)
  	}
  	fmt.println(i)
	}
}
```

```
for i,g := range map {
  fmt.println ("Key: ", k)
  fmt.println =("Value: ", v)
}												// For loop using a map
``` 

```
for _, v := range map {
  fmt.println("I want to print only values so: ", v)
}												// For loop with _
```

#### switch statements:
```
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
```
```
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
```
### functions:

```
func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
  return value1, value2
}

```
```
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
```

##### full simple function:
```
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
```

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

#### struct:
```
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
```

### method 

a method is a function with an implicit first argument, called a receiver

```
func (r ReceiverType) funcName(parameter) (results)
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
