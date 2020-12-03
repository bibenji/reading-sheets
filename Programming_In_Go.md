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

One way to be sure to use slice indexes that slice on character boundaries is to use functions from Go’s strings package, such as strings.Index() or strings.LastIndex()

with naïve

n	a	ï		v	e
0	1	2	3	4	5

We should use utf8.DecodeRuneInString() to get the first character (as a rune , along with the number of UTF-8 bytes used to represent it), and utf8.DecodeLastRuneInString() to get the last character

For strings that contain only 7-bit ASCII we can simply use the [] index operator which gives us very fast (O(1)) lookups. For non-ASCII strings we can convert the string to a []rune and use the [] index operator.<br />
This delivers very fast (O(1)) lookup performance, but at the expense of the one-off conversion which costs both CPU and memory (O(n)).

In the case of our example, if we wrote chars := []rune(s) , the chars variable would be created as a rune (i.e., int32 ) slice with the five code points—compared with six bytes.<br />
Recall that we can easily convert any rune (code point) back to a string —containing one character—using the string(char) syntax.

it isn’t suitable for working with arbitrary Unicode whitespace characters such as U+2028 (Line Separator, LS ) or U+2029 (Paragraph Separator, PS ). (string slicing)

i := strings.IndexFunc(line, unicode.IsSpace)<br />
strings.LastIndexFunc(line, unicode.IsSpace)

3.5. String Formatting with the Fmt Package

The fmt package also provides various scan functions (such as fmt.Scan() , fmt.Scanf() , and fmt.Scanln() ) for reading data from the console, from files, and from strings.

strings.Fields()

the strconv package

we can read input typed at the keyboard by creating a bufio.Reader to read from os.Stdin and use the bufio.Reader.ReadString() function to read each line entered

to print to os.Stdout (i.e., to the console)
fmt.Print()
fmt.Printf()
fmt.Println()

to output to a given io.Writer (e.g., to a file)
fmt.Fprint()
fmt.Fprintf()
fmt.Fprintf()

to output as a string
fmt.Sprint()
fmt.Sprintf()
fmt.Sprintln()

fmt.Errorf()

Verb Description/result
%% 	A literal % character
%b 	An integer value as a binary (base 2) number, or (advanced) a floating-
	point number in scientific notation with a power of 2 exponent
%c 	An integer code point value as a Unicode character
%d 	An integer value as a decimal (base 10) number
%e 	A floating-point or complex value in scientific notation with e
%E 	A floating-point or complex value in scientific notation with E
%f 	A floating-point or complex value in standard notation
%g 	A floating-point or complex value using %e or %f , whichever produces the
	most compact output
%G 	A floating-point or complex value using %E or %f , whichever produces the
	most compact output
%o 	An integer value as an octal (base 8) number
%p 	A value’s address as a hexadecimal (base 16) number with a prefix of 0x
	and using lowercase for the digits a – f (for debugging)
%q 	The string or []byte as a double-quoted string, or the integer as a single-
	quoted string, using Go syntax and using escapes where necessary
%s 	The string or []byte as raw UTF-8 bytes; this will produce correct
	Unicode output for a text file or on a UTF-8-savvy console
%t 	A bool value as true or false
%T 	A value’s type using Go syntax
%U 	An integer code point value using Unicode notation defaulting to four
	digits; e.g., fmt.Printf("%U", '¶' ) outputs U+00B6
%v 	A built-in or custom type’s value using a default format, or a custom
	value using its type’s String() method if it exists
%x 	An integer value as a hexadecimal (base 16) number or a string or
	[]byte value as hexadecimal digits (two per byte), using lowercase for
	the digits a – f
%X 	An integer value as a hexadecimal (base 16) number or a string or
	[]byte value as hexadecimal digits (two per byte), using uppercase for
	the digits A – F

Modifier Description/result
space	Makes the verb output “ - ” before negative numbers and a space before
		positive numbers or to put spaces between the bytes printed when
		using the %x or %X verbs; e.g., fmt.Printf("% X", "←" ) outputs E2 86 92
#		Makes the verb use an “alternative” output format:
%#o 	outputs octal with a leading 0
%#p		outputs a pointer without the leading 0x
%#q 	outputs a string or []byte as a raw string (using backticks) if
		possible—otherwise outputs a double-quoted string
%#v 	outputs a value as itself using Go syntax
%#x 	outputs hexadecimal with a leading 0x
%#X		outputs hexadecimal with a leading 0X
+		Makes the verb output + or - for numbers, ASCII characters (with
		others escaped) for strings, and field names for struct s
-		Makes the verb left-justify the value (the default is to right-justify)
0		Makes the verb pad with leading 0 s instead of spaces

n.m, n .m
		For numbers, makes the verb output a floating-point or complex
		value using n (of type int ) characters (or more if necessary to avoid
		truncation) and with m (of type int ) digits after the decimal point(s).
		For strings n specifies the minimum field width, and will result in
		space padding if the string has too few characters, and .m specifies the
		maximum number of the string’s characters to use (going from left to
		right), and will result in the string being truncated if it is too long.
		Either or both of m and n can be replaced with * in which case their
		values are taken from the arguments.
		Either n or .m may be omitted.

The way that fmt.Print() and fmt.Fprint() handle whitespace is subtly different from the fmt.Println() and fmt.Fprintln() functions. As a rule of thumb the former are most useful for printing a single value or for “converting” a value to a string without error checking

3.5.1. Formatting Booleans

```
fmt.Printf("%t %t\n", true, false)=
=> true false

fmt.Printf("%d %d\n", IntForBool(true), IntForBool(false))
=> 1 0

func IntForBool(b bool) int {
	if b {
		return 1
	}
	return 0
}
```

strconv.ParseBool()

3.5.2. Formatting Integers

fmt.Printf("|%b|%9b|%-9b|%09b|% 9b|\n", 37, 37, 37, 37, 37)
|100101|···100101|100101···|000100101|···100101|

fmt.Printf("|%o|%#o|%# 8o|%#+ 8o|%+08o|\n", 41, 41, 41, 41, -41)
|51|051|·····051|····+051|-0000051|

i := 3931
fmt.Printf("|%x|%X|%8x|%08x|%#04X|0x%04X|\n", i, i, i, i, i, i)
|f5b|F5B|·····f5b|00000f5b|0X0F5B|0x0F5B|

i = 569
fmt.Printf("|$%d|$%06d|$%+06d|$%s|\n", i, i, i, Pad(i, 6, '*'))
|$569|$000569|$+00569|$***569|

```
func Pad(number, width int, pad rune) string {
	s := fmt.Sprint(number)
	gap := width - utf8.RuneCountInString(s)
	if gap > 0 {
		return strings.Repeat(string(pad), gap) + s
	}
	return s
}
```

3.5.3. Formatting Characters

fmt.Printf("%d %#04x %U '%c'\n", 0x3A6, 934, '\u03A6', '\U000003A6')
934·0x03a6·U+03A6·'Φ'

3.5.4. Formatting Floating-Point Numbers

```
for _, x := range []float64{-.258, 7194.84, -60897162.0218, 1.500089e-8} {
	fmt.Printf("|%20.5e|%20.5f|%s|\n", x, x, Humanize(x, 20, 5, '*', ','))
}

|········-2.58000e-01|············-0.25800|************-0.25800|
|·········7.19484e+03|··········7194.84000|*********7,194.84000|
|········-6.08972e+07|·····-60897162.02180|***-60,897,162.02180|
|·········1.50009e-08|·············0.00000|*************0.00000|

func Humanize(amount float64, width, decimals int, pad, separator rune) string {
	dollars, cents := math.Modf(amount)
	whole := fmt.Sprintf("%+.0f", dollars)[1:] // Strip "±"
	fraction := ""
	if decimals > 0 {
		fraction = fmt.Sprintf("%+.*f", decimals, cents)[2:] // Strip "±0"
	}
	sep := string(separator)
	for i := len(whole) - 3; i > 0; i -= 3 {
		whole = whole[:i] + sep + whole[i:]
	}
	if amount < 0.0 {
		whole = "-" + whole
	}
	number := whole + fraction
	gap := width - utf8.RuneCountInString(number)
	if gap > 0 {
		return strings.Repeat(string(pad), gap) + number
	}
	return number
}
```

```asp
for _, x := range []complex128{2 + 3i, 172.6 - 58.3019i, -.827e2 + 9.04831e-3i} {
	fmt.Printf("|%15s|%9.3f|%.2f|%.1e|\n", fmt.Sprintf("%6.2f%+.3fi", real(x), imag(x)), x, x, x)
}

|····2.00+3.000i|(····2.000···+3.000i)|(2.00+3.00i)|(2.0e+00+3.0e+00i)|
|·172.60-58.302i|(··172.600··-58.302i)|(172.60-58.30i)|(1.7e+02-5.8e+01i)|
|··-82.70+0.009i|(··-82.700···+0.009i)|(-82.70+0.01i)|(-8.3e+01+9.0e-03i)|
```

3.5.5. Formatting Strings and Slices

P. 101

```
slogan := "End Óréttlæti♥"
fmt.Printf("%s\n%q\n%+q\n%#q\n", slogan, slogan, slogan, slogan)

chars := []rune(slogan)
fmt.Printf("%x\n%#x\n%#X\n", chars, chars, chars)

bytes := []byte(slogan)
fmt.Printf("%s\n%x\n%X\n% X\n%v\n", bytes, bytes, bytes, bytes, bytes)
```

3.5.6. Formatting for Debugging

```
p := polar{-83.40, 71.60}
fmt.Printf("|%T|%v|%#v|\n", p, p, p)
fmt.Printf("|%T|%v|%t|\n", false, false, false)
fmt.Printf("|%T|%v|%d|\n", 7607, 7607, 7607)
fmt.Printf("|%T|%v|%f|\n", math.E, math.E, math.E)
fmt.Printf("|%T|%v|%f|\n", 5+7i, 5+7i, 5+7i)
s := "Relativity"
fmt.Printf("|%T|\"%v\"|\"%s\"|%q|\n", s, s, s, s)
```

Two of Go’s types have synonyms: byte for uint8 and rune for int32 .

%p (pointer) verb

i := 5
f := -48.3124
s := "Tomás Bretón"
fmt.Printf("|%p → %d|%p → %f|%#p → %s|\n", &i, i, &f, f, &s, s)

& address of operator is explained in the next chapter

ability to outpu slices and maps, and even channels!

```asp
fmt.Println([]float64{math.E, math.Pi, math.Phi})
fmt.Printf("%v\n", []float64{math.E, math.Pi, math.Phi})
fmt.Printf("%#v\n", []float64{math.E, math.Pi, math.Phi})
fmt.Printf("%.5f\n", []float64{math.E, math.Pi, math.Phi})
```

```asp
fmt.Printf("%q\n", []string{"Software patents", "kill", "innovation"})
fmt.Printf("%v\n", []string{"Software patents", "kill", "innovation"})
fmt.Printf("%#v\n", []string{"Software patents", "kill", "innovation"})
fmt.Printf("%17s\n", []string{"Software patents", "kill", "innovation"})
```

```asp
fmt.Printf("%v\n", map[int]string{1: "A", 2: "B", 3: "C", 4: "D"})
fmt.Printf("%#v\n", map[int]string{1: "A", 2: "B", 3: "C", 4: "D"})
fmt.Printf("%v\n", map[int]int{1: 1, 2: 2, 3: 4, 4: 8})
fmt.Printf("%#v\n", map[int]int{1: 1, 2: 2, 3: 4, 4: 8})
fmt.Printf("%04b\n", map[int]int{1: 1, 2: 2, 3: 4, 4: 8})
```

only thing not possible with fmt package's print functions is padding with a particular character (other than zeros or spaces)

but it's easy to do with functions

3.6. Other String-Related Packages

P. 106

3.6.1. The Strings Package

strings.Split()
strings.SplitN()
strings.SplitAfter()
strings.SplitAfterN()

strings.FieldsFunc()
```asp
for _, record := range []string{"László Lajtha*1892*1963", "Édouard Lalo\t1823\t1892", "José Ángel Lamas|1775|1814"} {
	fmt.Println(strings.FieldsFunc(record, func(char rune) bool {
		switch char {
			case '\t', '*', '|':
				return true
		}
		return false
	}))
}
```

The Strings Package's Functions

- strings.Contains(s, t) 		true if t occurs in s
- strings.Count(s, t) 			How many (nonoverlapping) times t occurs in s
- strings.EqualFold(s, t) 		true if the strings are case-insensitively equal
- strings.Fields(s) 			The []string that results in splitting s on white-space strings.
- FieldsFunc(s, f)				The []string that results in splitting s at every character where f returns true
- strings.HasPrefix(s, t) 		true if s starts with t 
- strings.HasSuffix(s, t) 		true if s ends with t
- strings.Index(s, t) 			The index of the first occurrence of t in s
- strings.IndexAny(s, t)		The first index in s of any character that is in t
- strings.IndexFunc(s, f)		The index of the first character in s for which f returns true
- strings.IndexRune(s, char)	The index of the first occurrence of character char of type rune in s
- strings.Join(xs, t)			A string containing the concatenation of all the strings in xs , each separated by t (which can be "" )
- strings.LastIndex(s, t)		The index of the last occurrence of t in s
- strings.LastIndexAny(s, t)	The last index in s of any character that is in t
- strings.LastIndexFunc(s, f)	The index of the last character in s for which f returns true
- strings.Map(mf, t)			A copy of t with every character replaced or delet-ed according to the mapping function mf with the signature func(rune) rune (see text)
- strings.NewReader(s) 			A pointer to a value that provides Read() , Read-Byte() , and ReadRune() methods that operate on s
- strings.NewReplacer(...)		A pointer to a value that has methods for replacing each pair of old, new strings it is given
- strings.Repeat(s, i) 			A string consisting of i concatenations of s

- strings.Replace(s, old, new, i) 	A copy of s with every nonoverlapping occurrence of string old replaced by string new if i is -1 , or with at most i replacements otherwise
- strings.Split(s, t)				The []string that results in splitting s on t as many times as t occurs in s
- strings.SplitAfter(s, t)			Works like strings.Split() only the separator is kept in the resultant strings (see text)			
- SplitAfterN(s, t, i) 				Works like strings.SplitN() only the separator is kept in the resultant strings
- strings.SplitN(s, t, i)			The []string that results in splitting s on t , i -1 times
- strings.Title(s)					A copy of s with the first letter of every word title-cased
- strings.ToLower(s)				A lowercased copy of s
- strings.ToLowerSpecial(r, s)		A lowercased copy of s , prioritizing the rules in r (advanced)
- strings.ToTitle(s)				A title-cased copy of s
- strings.ToTitleSpecial(r, s)		A title-cased copy of s , prioritizing the rules in r (advanced)
- strings.ToUpper(s)				An uppercased copy of s
- strings.ToUpperSpecial(r, s)		An uppercased copy of s , prioritizing the rules in r (advanced)
- strings.Trim(s, t)				A copy of s with the characters in t removed from both ends
- strings.TrimFunc(s, f)			A copy of s with the characters for which f returns true removed from both ends
- strings.TrimLeft(s, t)			A copy of s with the characters in t removed from the start
- strings.TrimLeftFunc(s, f)		A copy of s with the characters for which f returns true removed from the start
- strings.TrimRight(s, t)			A copy of s with the characters in t removed from the end
- strings.TrimRightFunc(s, f)		A copy of s with the characters for which f returns true removed from the end
- strings.TrimSpace(s)				A copy of s with whitespace removed from both ends

Thing to simplify whitespace... P. 111

[...] About reading data from types that implement the ReadRune() function such as the bufio.Reader
Also, bufio.NewReader()
`reader := strings.NewReader("Café)`

3.6.2. The Strconv Package

P. 113
