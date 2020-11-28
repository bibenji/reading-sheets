# 1 An Overview in Five Examples

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

1.7. Polar to Cartesian - Concurrency

P. 40


messages := make(chan string, 10)
messages <- "Leader"
messages <- "Follower"
message1 := <-messages
message2 := <-messages

1.8. Exercice

P. 48

# 2 Booleans and Numbers

2.1. Constants and Variables

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

2.2. Boolean Values and Expressions

reflect.DeepEqual()

2.3. Numeric Types

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

2.4. Example Statistics

P. 72

2.4.1. Implementing Simple Statistics Functions

2.4.2. Implementing a Basic HTTP Server

see packages : html, net/http, html/template, text/template

2.5. Exercises

# 3 Strings

P. 81

3.1. Literals, Operators, and Escapes

With " (double quotes) or ` (backticks)

double quotes supports escape sequences ( \\, \ooo, \', \", \a, \b, \f, \n, \r, \t, \uhhh, \Uhhhhhhhh, \v, \xhh)
backsticks support multiple lines

+ concatenation operator

you can do strings comparisons too ("Josey" < "José", "Josey" == "José")

- len([]rune(s)) The number of characters in string s —use the faster utf8. RuneCountInString() instead
- []rune(s) Converts string s into a slice of Unicode code points
- []byte(s) Converts string s into a slice of raw bytes without copying; there’s no guarantee that the bytes are valid UTF-8
- string(bytes) Converts a []byte or []uint8 into a string without copying; there’s no guarantee that the bytes are valid UTF-8
- string(i) Converts i of any integer type into a string ; assumes that i is a Unicode code point; e.g., if i is 65 , it returns "A" ★
- strconv.Itoa(i) The string representation of i of type int and an error ; e.g., if i is 65 , it returns ( "65" , nil )

3.2. Comparing Strings

Some problems can occur

[...]

3.3. Characters and Strings

character, code point, Unicode character, Unicode code point interchangeably to refer to a rune (or int32) that holds a single character

we can convert a single character into a one-character string using Go's standard conversion syntax (string(char)) (P. 87 for example)

An entire string can be converted to a slice of rune s (i.e., code points) using the syntax chars := []rune(s) where s is of type string .

And reverse conversion is equally simple using the syntax s := string(chars) where chars is of type []rune or []int32

Better to make []string and strings.Join() than use string +=

And even better :
```
var buffer bytes.Buffer
for {
	if piece, ok := getNextValidString(); ok {
		buffer.WriteString(piece)
	} else {
		break
	}
}
fmt.Print(buffer.String(), "\n")
``` 

Accumulating strings in a bytes.Buffer is potentially much more memory- and CPU-efficient than using the += operator, especially if the number of strings to concatenate is large.

for ...range loop to iterate over a string character by character

Big-O Notation [...] P. 89
- O(1) means constant time, that is, the fastest possible time no matter what n’s size.
- O(log n) means logarithmic time; this is very fast and in proportion to log n.
- O(n) means linear time; this is fast and in proportion to n.
- O( n 2 ) means quadratic time; this is starting to get slow and is in proportion to n 2 .
- O( n m ) means polynomial time which quickly becomes slow as n grows, especially if m ≥ 3.
- O(n!) means factorial time; even for small values of n this can become too slow to be practical.

3.4. Indexing and Slicing Strings
