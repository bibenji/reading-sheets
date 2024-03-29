## 1 Introducing Go

[...]

May goroutines execute on a single OS thread

```
func log(msg string){
... some logging code here
}
// Elsewhere in our code after we've discovered an error.
go log("something dire happened")
```

channels = data structures that enable safe data communication between goroutines

Inheritance & Composition
Interfaces

import "fmt"

## 2 Go quick-start

P. 30

2.3.2 feed.go

P. 43

[...] Relire toute cette partie

P. 53

## 3 Packaging and tooling

```
net/http/
    cgi/
    cookiejar/
        testdata/
    fcgi/
    httptest/
    httputil/
    pprof/
    testdata/
```

godoc fmt
go build
go get

```
import (
    "fmt"
    myfmt "mylib/fmt"
)
```

func init()

go with no arguments to have all it can do
(build, clean, doc, env, fix, fmt, generate, get, install, list, run , test, tool, version, vet)

go build hello.go
go clean hello.go

go build github.com/goinaction/code/chapter3/wordcount

go build github.com/goinaction/code/chapter3/...

go build wordcount.go

go build .

go run wordcount.go

go vet -> to check errors

go fmt -> format the code

go doc 
godoc
godoc -http=:6060

godep

#### gb too

gb build all

revoir gb

## 4 Arrays, slice, and maps

Arrays

```asp
// Declare an integer array of five elements.
var array [5]int

// Declare an integer array of five elements.
// Initialize each element with a specific value.
array := [5]int{10, 20, 30, 40, 50}

// Declare an integer array.
// Initialize each element with a specific value.
// Capacity is determined based on the number of values initialized.
array := [...]int{10, 20, 30, 40, 50}

// Declare an integer array of five elements.
// Initialize index 1 and 2 with specific values.
// The rest of the elements contain their zero value.
array := [5]int{1: 10, 2: 20}

// Declare an integer array of five elements.
// Initialize each element with a specific value.
array := [5]int{10, 20, 30, 40, 50}
// Change the value at index 2.
array[2] = 35

// Declare an integer pointer array of five elements.
// Initialize index 0 and 1 of the array with integer pointers.
array := [5]*int{0: new(int), 1: new(int)}
// Assign values to index 0 and 1.
*array[0] = 10
*array[1] = 20

// Declare a string array of five elements.
var array1 [5]string
// Declare a second string array of five elements.
// Initialize the array with colors.
array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
// Copy the values from array2 into array1.
array1 = array2

// Declare a string pointer array of three elements.
var array1 [3]*string
// Declare a second string pointer array of three elements.
// Initialize the array with string pointers.
array2 := [3]*string{new(string), new(string), new(string)}
// Add colors to each element
*array2[0] = "Red"
*array2[1] = "Blue"
*array2[2] = "Green"
// Copy the values from array2 into array1.
array1 = array2
```

Multidimensional arrays
```
// Declare a two dimensional integer array of four elements
// by two elements.
var array [4][2]int
// Use an array literal to declare and initialize a two
// dimensional integer array.
array := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
// Declare and initialize index 1 and 3 of the outer array.
array := [4][2]int{1: {20, 21}, 3: {40, 41}}
// Declare and initialize individual elements of the outer
// and inner array.
array := [4][2]int{1: {0: 20}, 3: {1: 41}}

// Declare a two dimensional integer array of two elements.
var array [2][2]int
// Set integer values to each individual element.
array[0][0] = 10
array[0][1] = 20
array[1][0] = 30
array[1][1] = 40

// Declare two different two dimensional integer arrays.
var array1 [2][2]int
var array2 [2][2]int
// Add integer
array2[0][0] =
array2[0][1] =
array2[1][0] =
array2[1][1] =
values to each individual element.
10
20
30
40
// Copy the values from array2 into array1.
array1 = array2

// Copy index 1 of array1 into a new array of the same type.
var array3 [2]int = array1[1]
// Copy the integer found in index 1 of the outer array
// and index 0 of the interior array into a new variable of
// type integer.
var value int = array1[1][0]
```

Passing arrays between functions
```asp
// Declare an array of 8 megabytes.
var array [1e6]int
// Pass the array to the function foo.
foo(array)
// Function foo accepts an array of one million integers.
func foo(array [1e6]int) {
...
}

// Allocate an array of 8 megabytes.
var array [1e6]int
// Pass the address of the array to the function foo.
foo(&array)
// Function foo accepts a pointer to an array of one million integers.
func foo(array *[1e6]int) {
...
}

```

Slice internals and fundamentals

A slice is a data structure that provides a way for you to work with and manage collections of data.

append built-in function

with : pointer, length and capacity

```asp
// Create a slice of strings.
// Contains a length and capacity of 5 elements.
slice := make([]string, 5)

// Create a slice of integers.
// Contains a length of 3 and has a capacity of 5 elements.
slice := make([]int, 3, 5)

// Create a slice of strings.
// Contains a length and capacity of 5 elements.
slice := []string{"Red", "Blue", "Green", "Yellow", "Pink"}

// Create a slice of integers.
// Contains a length and capacity of 3 elements.
slice := []int{10, 20, 30}

// Create a slice of strings.
// Initialize the 100th element with an empty string.
slice := []string{99: ""}
```

!!!
Remember, if you specify a value inside the [ ] operator, you’re creating an array. If
you don’t specify a value, you’re creating a slice.

```asp
// Create a nil slice of integers.
var slice []int

// Use make to create an empty slice of integers.
slice := make([]int, 0)
// Use a slice literal to create an empty slice of integers.
slice := []int{}
```

Taking the slice of a slice
```
// Create a slice of integers.
// Contains a length and capacity of 5 elements.
slice := []int{10, 20, 30, 40, 50}

// Create a new slice.
// Contains a length of 2 and capacity of 4 elements.
newSlice := slice[1:3]
```

For slice[i:j] with an underlying array of capacity k
Length:
j - i
Capacity: k - i

#### Growing Slices

```asp
// Create a slice of integers.
// Contains a length and capacity of 5 elements.
slice := []int{10, 20, 30, 40, 50}
// Create a new slice.
// Contains a length of 2 and capacity of 4 elements.
newSlice := slice[1:3]
// Allocate a new element from capacity.
// Assign the value of 60 to the new element.
newSlice = append(newSlice, 60)
```

```asp
// Slice the third element and restrict the capacity.
// Contains a length of 1 element and capacity of 2 elements.
slice := source[2:3:4]
```

For slice[i:j:k] or [2:3:4]
Length: j - i or 3 - 2 = 1
Capacity: k - i or 4 - 2 = 2

By having the option to set the capacity of a new slice to be the same as the length,
you can force the first append operation to detach the new slice from the underlying
array. Detaching the new slice from its original source array makes it safe to change.

```asp
// Create two slices each initialized with two integers.
s1 := []int{1, 2}
s2 := []int{3, 4}
// Append the two slices together and display the results.
fmt.Printf("%v\n", append(s1, s2...))
Output:
[1 2 3 4]
```

Iterating over slices
```asp
// Create a slice of integers.
// Contains a length and capacity of 4 elements.
slice := []int{10, 20, 30, 40}
// Iterate over each element and display each value.
for index, value := range slice {
fmt.Printf("Index: %d Value: %d\n", index, value)
}
```

!!! range is making a copy of the value, not returning a reference. If you use the address of the value variable as a pointer to each element, you’ll be making a mistake

```asp
// Create a slice of integers.
// Contains a length and capacity of 4 elements.
slice := []int{10, 20, 30, 40}
// Iterate over each element and display the value and addresses.
for index, value := range slice {
fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n",
value, &value, &slice[index])
}
Output:
Value: 10 Value-Addr: 10500168 ElemAddr: 1052E100
Value: 20 Value-Addr: 10500168 ElemAddr: 1052E104
Value: 30 Value-Addr: 10500168 ElemAddr: 1052E108
Value: 40 Value-Addr: 10500168 ElemAddr: 1052E10C
```

If you don’t need the index value, you can use the underscore character to discard
the value.

```asp
// Create a slice of integers.
// Contains a length and capacity of 4 elements.
slice := []int{10, 20, 30, 40}
// Iterate over each element starting at element 3.
for index := 2; index < len(slice); index++ {
fmt.Printf("Index: %d Value: %d\n", index, slice[index])
}

...
```

len(slice) and cap(slice)

#### Multidimensional slices

P. 100

```asp
// Create a slice of a slice of integers.
slice := [][]int{{10}, {100, 200}}

// Append the value of 20 to the first slice of integers.
slice[0] = append(slice[0], 20)
```

Passing slices between functions

```asp
// Allocate a slice of 1 million integers.
slice := make([]int, 1e6)
// Pass the slice to the function foo.
slice = foo(slice)
// Function foo accepts a slice of integers and returns the slice back.
func foo(slice []int) []int {
...
return slice
}
```

On a 64-bit architecture, a slice requires 24 bytes of memory. The pointer field
requires 8 bytes, and the length and capacity fields require 8 bytes respectively. Since
the data associated with a slice is contained in the underlying array, there are no prob-
lems passing a copy of a slice to any function. Only the slice is being copied, not the
underlying array

### 4.3 Map internals and fundamentals

P. 102

A map is a data structure that provides you with an unordered collection of key/value
pairs.

map is implemented using a hash table
so no real order

```asp
// Create a map with a key of type string and a value of type int.
dict := make(map[string]int)

// Create a map with a key and value of type string.
// Initialize the map with 2 key/value pairs.
dict := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}

// Create a map using a slice of strings as the value.
dict := map[int][]string{}

// Create an empty map to store colors and their color codes.
colors := map[string]string{}
// Add the Red color code to the map.
colors["Red"] = "#da1337"
```

You can create a nil map by declaring a map without any initialization. A nil map
can’t be used to store key/value pairs. Trying will produce a runtime error.

```asp
// Retrieve the value for the key "Blue".
value, exists := colors["Blue"]
// Did this key exist?
if exists {
    fmt.Println(value)
}

// Retrieve the value for the key "Blue".
value := colors["Blue"]
// Did this key exist?
if value != "" {
    fmt.Println(value)
}
```

iterate over a map
```asp
// Create a map of colors and color hex codes.
colors := map[string]string{
"AliceBlue":
"#f0f8ff",
"Coral":
"#ff7F50",
"DarkGray":
"#a9a9a9",
"ForestGreen": "#228b22",
}
// Display all the colors in the map.
for key, value := range colors {
fmt.Printf("Key: %s Value: %s\n", key, value)
}
```

removing an item from a map : delete
`delete(colors, "Coral")`

Passing a map between two functions doesn’t make a copy of the map.

## 5 Go's type system

P. 109

struct

```asp
// user defines a user in the program.
type user struct {
    name        string
    email       string
    ext         int
    privileged  bool
}
```

eserve the use of the keyword var as a way to indicate that a variable
is being set to its zero value. If the variable will be initialized to something other than
its zero value, then use the short variable declaration operator with a struct literal.

```asp
// Declare a variable of type user and initialize all the fields.
lisa := user{
    name: "Lisa",
    email: "lisa@email.com",
    ext: 123,
    privileged: true,
}

// Declare a variable of type user.
lisa := user{"Lisa", "lisa@email.com", 123, true}
```

```asp
// Declare a variable of type admin.
fred := admin{
    person: user{
        name: "Lisa",
        email: "lisa@email.com",
        ext: 123,
        privileged: true,
    },
    level: "super",
}
```

Declaration of a new type based on an int64

`type Duration int64`

5.2 Methods

```asp
// user defines a user in the program.
type user struct {
    name string
    email string
}

// notify implements a method with a value receiver.
func (u user) notify() {
    fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email=
}

func (u *user) changeEmail(email string) {
    u.email = email
}

[...]

in main :

// Pointers of type user can also be used to call methods
// declared with a value receiver.
lisa := &user{"Lisa", "lisa@email.com"}
lisa.notify()
(you can call methods usign a pointer)

```

The parameter between
the keyword func and the function name is called a receiver

two types of receivers : value and pointer receivers

= makes a copy

(*lisa).notify() [to support the method call]

...

(&bill).notify() [to support the method call]

5.3 The Nature of types

Reference types

...

net package
type IP []byte

In the end, reference type values are treated like primitive data values.
= copy

Struct types

```asp
func Now() Time {
    sec, nsec := now()
    return Time{sec + unixToInternal, nsec, Local}
}
```

```asp
func (t Time) Add(d Duration) Time {
    t.sec += int64(d / 1e9)
    nsec := int32(t.nsec) + int32(d%1e9)
    if nsec >= 1e9 {
        t.sec++
        nsec -= 1e9
    } else if nsec < 0 {
        t.sec--
        nsec += 1e9
    }
    t.nsec = nsec
    return t
}
```

In most cases, struct types don’t exhibit a primitive nature, but a nonprimitive one.
In these cases, adding or removing something from the value of the type should
mutate the value.

```asp
// File represents an open file descriptor.
type File struct {
    *file
}
```

Since there’s no way to pre-
vent programmers from making copies, the implementation of the File type uses an
embedded pointer of an unexported type. We’ll talk about embedding types later in
this chapter, but this extra level of indirection provides protection from copies.

```asp
func Open(name string) (file *File, err error) {
    return OpenFile(name, O_RDONLY, 0)
}
```

```asp
func (f *File) Chdir() error {
    if f == nil {
        return ErrInvalid
    }

    if e := syscall.Fchdir(f.fd); e != nil {
        return &PathError{"chdir", f.name, e}
    }
    
    return nil
}
```

The decision to use a value or pointer receiver should not be based on whether the
method is mutating the receiving value. The decision should be based on the nature
of the type. One exception to this guideline is when you need the flexibility that value
type receivers provide when working with interface values. In these cases, you may
choose to use a value receiver even though the nature of the type is nonprimitive. It’s
entirely based on the mechanics behind how interface values call methods for the val-
ues stored inside of them.

5.4 Interfaces

P. 122


