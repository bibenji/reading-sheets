# 1 An Overview in Five Examples

#### 1.1. Getting Going
#### 1.2. Editing, Compiling, and Running
#### 1.3. Hello Who?
#### 1.4. Big Digits - Two Dimensional Slices
#### 1.5. Stack - Custom Types with Methods
#### 1.6. Americanise - Files, Maps, and Closures

slices and map

`godoc -http=:8000`

`godoc image NewRGBA`

`godoc image/png`

gonow and gorun

slce[n], slice[n:]

longWeekend := []string{"Friday", "Saturday", "Sunday", "Monday"}<br />
var lowPrimes = []int{2, 3, 5, 7, 11, 13, 17, 19}

In UTF-8 (and 7-bit ASCII) the character '0' is code point (character) 48 decimal, the character '1' is code point 49, and so on. So if, for example, we have the character '3' (code point 51), we can get its integer value by doing the subtraction '3' - '0' (i.e., 51 − 48) which results in an integer (of type byte ) of value 3.

Go uses single quotes for character literals

byte, int32, int64

P. 21

```
// without * means it's made on a copy of the instance
func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}

// with * means it's made on the instance itself
func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
}
```

Dereferencing (the pointer) is done by preceding the variable name with a star.

References
behave very much like pointers in that when they are passed to functions any
changes made to them inside the function affect the original channel, map, or
slice. However, references don’t need to be dereferenced, so in most cases there’s
no need to use stars with them.

P. 29

## 1.7. Polar to Cartesian - Concurrency

P. 40

messages := make(chan string, 10)
messages <- "Leader"
messages <- "Follower"
message1 := <-messages
message2 := <-messages

## 1.8. Exercice

P. 48

# 2 Booleans and Numbers

## 2.1. Constants and Variables

P. 51

Go’s Predefined Identifiers
append copy int8 nil true
bool delete int16 panic uint
byte error int32 print uint8
cap false int64 println uint16
close float32 iota real uint32
complex float64 len recover uint64
complex64 imag make rune uintptr
complex128 int new string

Typed numeric constants (e.g., top ) can only be used in expressions with other numbers of the same type (unless converted). Untyped numeric constants can be used in expressions with numbers of any built-in type

```asp
const (
	Cyan = 0
	Magenta = 1
	Yellow = 2
)

const (
	Cyan = iota // 0
	Magenta // 1
	Yellow // 2
)
```

Similarly, if Cyan was set to 9, then they would all be set to 9; or if Magenta was set to 5, Cyan would be set to 0 (first in the group and not assigned an explicit value or iota ), Magenta would be 5 (explicitly set), and Yellow would be 5 (the previous constant’s value).

```asp
type BitFlag int

const (
	Active BitFlag = 1 << iota
	Send	// Implicitly BitFlag = 1 << iota
	Receive	// Implicitly BitFlag = 1 << iota
)

flag := Active | Send
```

```asp
func (flag BitFlag) String() string {
	var flags []string
	if flag&Active == Active {
		flags = append(flags, "Active")
	}
	if flag&Send == Send {
		flags = append(flags, "Send")
	}
	if flag&Receive == Receive {
		flags = append(flags, "Receive")
	}
	if len(flags) > 0 { // int(flag) is vital to avoid infinite recursion!
		return fmt.Sprintf("%d(%s)", int(flag), strings.Join(flags, "|"))
	}
	return "0()"
}
```

## 2.2. Boolean Values and Expressions

reflect.DeepEqual()

## 2.3. Numeric Types

Rat which are of unbounded size (i.e., limited only by the machine’s memory).

If we need to perform arithmetic or comparisons on typed numbers of different types we must perform conversions—usually to the biggest type to avoid loss of accuracy.

If we want to perform safe downsizing conversions we can always create suitable functions. For example:
```asp
func Uint8FromInt(x int) (uint8, error) {
	if 0 <= x && x <= math.MaxUint8 {
		return uint8(x), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", x)
}
```

Constant expressions are evaluated at compile time; they may use any of the arithmetic, Boolean, and comparison operators.

```asp
const (
	efri int64		= 10000000000							// type: int64
	hlutföllum		= 16.0 / 9.0							// type: float64
	mælikvarða		= complex(-2, 3.5) * hlutföllum			// type: complex128
	erGjaldgengur 	= 0.0 <= hlutföllum && hlutföllum < 2.0	// type: bool
)
```

2.3.1. Integer Types

P. 59

11 separate integer types, five signed and five unsigned, plus an integer type for storing pointers!

byte for synonym of unsigned uint8
rune for synonym of int32 type

int => represented by signed 32-bit integer

byte 	Synonym for uint8
int 	The int32 or int64 range depending on the implementation
int8 	[−128, 127]
int16 	[−32 768, 32 767]
int32 	[−2 147 483 648, 2 147 483 647]
int64 	[−9 223 372 036 854 775 808, 9 223 372 036 854 775 807]
rune 	Synonym for int32
uint 	The uint32 or uint64 range depending on the implementation
uint8 	[0, 255]
uint16 	[0, 65 535]
uint32 	[0, 4 294 967 295]
uint64 	[0, 18 446 744 073 709 551 615]
uintptr An unsigned integer capable of storing a pointer value (advanced)

^x 		The bitwise complement of x
x %= y 	Sets x to be the remainder of dividing x by y ; division by zero causes
		a runtime panic
x &= y 	Sets x to the bitwise AND of x and y
x |= y 	Sets x to the bitwise OR of x and y
x ^= y 	Sets x to the bitwise XOR of x and y
x &^= y Sets x to the bitwise clear ( AND NOT ) of x and y
x >>= u Sets x to the result of right-shifting itself by unsigned int u shifts
x <<= u Sets x to the result of left-shifting itself by unsigned int u shifts
x % y 	The remainder of dividing x by y ; division by zero causes a runtime
		panic
x & y 	The bitwise AND of x and y
x | y 	The bitwise OR of x and y
x ^ y 	The bitwise XOR of x and y
x &^ y 	The bitwise clear ( AND NOT ) of x and y
x << u 	The result of left-shifting x by unsigned int u shifts
x >> u 	The result of right-shifting x by unsigned int u shifts

2.3.1.1. Big Integers

big.Ints
big.Rats

math/big package

[...] P. 63

2.3.2. Floating-Point Types

float32 	±3.402 823 466 385 288 598 117 041 834 845 169 254 40 × 10^38
			The mantissa is reliably accurate to about 7 decimal places.
float64 	±1.797 693 134 862 315 708 145 274 237 317 043 567 981 × 10^308
			The mantissa is reliably accurate to about 15 decimal places.
complex64	The real and imaginary parts are both of type float32 .
complex128 	The real and imaginary parts are both of type float64 .

The Math Package’s Constants and Functions [...] P. 65-67

2.3.2.1. Complex Types

P. 70

[...]

match/cmplx package

The Complex Math Package’s Functions [...] P. 71

## 2.4. Example Statistics

P. 72

2.4.1. Implementing Simple Statistics Functions

2.4.2. Implementing a Basic HTTP Server

see packages : html, net/http, html/template, text/template

## 2.5. Exercises

# <a href="/programming-in-go-rs/rs-programming-in-go-chapter003.md">3 Strings</a>

#### 3.1. Literals, Operators, and Escapes
#### 3.2. Comparing Strings
#### 3.3. Characters and Strings
#### 3.4. Indexing and Slicing Strings
#### 3.5. String Formatting with the Fmt Package
#### 3.6. Other String-Related Packages
#### 3.7. Example: M3u2pls
#### 3.8. Exercices

# 4 Collection Types

P. 139

Go does not support pointer arithmetic, thus eliminating a whole category of potential bugs that can affect C and C++ programs.

container/heap , container/list , and container/ring

## 4.1. Values, Pointers, and Reference Types

Go variables hold values
channels, functions, methods, maps and slices hold references
and variables that hold pointers

& used as a binary operator performs a bitwise AND

& used as a unary operator returns memory address of its operand

*int = pointer to int

* operator
* multiplies operands when used as binary operator
* provides access to value pointed to by the variable it is applied to as unary operator

x := 3
pi := &x (pi est un *int)

si on fait x++, ça change aussi la valeur de *pi (pas la valeur de pi qui est une adresse quoi)

si on fait *pi++, ça change la valeur de x du coup

& sert à obtenir pointeur sur une variable
et * pi valeur de la variable pointée

and pointers to pointers, and pointers to pointers to pointers (**int i.e.)

swapAndProduct1(&x, &y, &product), call with:
```asp
func swapAndProduct1(x, y, product *int) {
	if *x > *y {
		*x, *y = *y, *x
	}
	*product = *x * *y // The compiler would be happy with: *product=*x**y
}
```

type composer struct { }

to have a value, init with :
- composer{}

to have a pointer to composer :
- agnes := new(composer)
- julia := &composer{}
- augusta := &composer{"Blibli", 1847}

julia.name, julia.birthYear = "Blabla", 1819

go has reference types (and interfaces)

passing slice or pointer to a struct is cheap

maps and slices are reference types

for index, item ... range loop make a copy of each item

but not for i := range

```asp
func resizeRect(rect *rectangle, Δwidth, Δheight int) {
	(*rect).x1 += Δwidth // Ugly explicit dereference
	rect.y1 += Δheight // . automatically dereferences structs
}
```

## 4.2. Arrays and Slices

P. 148

Arrays are created using the syntaxes:
[length]Type
[N]Type{value1, value2, ... , valueN}
[ ... ]Type{value1, value2, ... , valueN}

array length is fix and unchangeable

```
var buffer [20]byte
var grid1 [3][3]int
grid1[1][0], grid1[1][1], grid1[1][2] = 8, 6, 2
grid2 := [3][3]int{{4, 3}, {8, 6, 2}}
cities := [...]string{"Shanghai", "Mumbai", "Istanbul", "Beijing"}
cities[len(cities)-1] = "Karachi"
fmt.Println("Type
Len Contents")
fmt.Printf("%-8T %2d %v\n", buffer, len(buffer), buffer)
fmt.Printf("%-8T %2d %q\n", cities, len(cities), cities)
fmt.Printf("%-8T %2d %v\n", grid1, len(grid1), grid1)
fmt.Printf("%-8T %2d %v\n", grid2, len(grid2), grid2)
```

Arrays are of fixed size whereas slices can be resized.

Better to use slices unless very specific need.

We could store any type of items with the same interface.

Slices are created using the syntaxes:
make([]Type, length, capacity)
make([]Type, length)
[]Type{}
[]Type{value1, value2, ... , valueN}

in second, third and fourth syntaxes, capacity and length are the same

s[n] The item at index position n in slice s
s[n:m] A slice taken from slice s from index positions n to m - 1
s[n:] A slice taken from slice s from index positions n to len(s) - 1
s[:m] A slice taken from slice s from index positions 0 to m - 1
s[:] A slice taken from slice s from index positions 0 to len(s) - 1
cap(s) The capacity of slice s ; always ≥ len(s)
len(s) The number of items in slice s ; always ≤ cap(s)
s = s[:cap(s)] Increase slice s ’s length to its capacity if they are different

The syntax []Type{} is equivalent to make([]Type, 0) ;

append() to increase capacity

s := new([7]string)[:]
s[0], s[1], s[2], s[3], s[4], s[5], s[6] = "A", "B", "C", "D", "E", "F","G"

...

The buffer ’s contents are only the first len(buffer) items; the other items are
inaccessible unless we reslice the buffer —something we will see how to do later
on in this section.

...

4.2.1. Indexing and Slicing Slices

P. 153 - Relire un peu au-dessus

4.3.

4.4.
