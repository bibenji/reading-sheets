# 3 Strings

P. 81

## 3.1. Literals, Operators, and Escapes

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

## 3.2. Comparing Strings

Some problems can occur

[...]

## 3.3. Characters and Strings

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

## 3.4. Indexing and Slicing Strings

One way to be sure to use slice indexes that slice on character boundaries is to use functions from Go’s strings package, such as strings.Index() or strings.LastIndex()

with naïve

```asp
n	a	ï		v	e
0	1	2	3	4	5
```


We should use utf8.DecodeRuneInString() to get the first character (as a rune , along with the number of UTF-8 bytes used to represent it), and utf8.DecodeLastRuneInString() to get the last character

For strings that contain only 7-bit ASCII we can simply use the [] index operator which gives us very fast (O(1)) lookups. For non-ASCII strings we can convert the string to a []rune and use the [] index operator.<br />
This delivers very fast (O(1)) lookup performance, but at the expense of the one-off conversion which costs both CPU and memory (O(n)).

In the case of our example, if we wrote chars := []rune(s) , the chars variable would be created as a rune (i.e., int32 ) slice with the five code points—compared with six bytes.<br />
Recall that we can easily convert any rune (code point) back to a string —containing one character—using the string(char) syntax.

it isn’t suitable for working with arbitrary Unicode whitespace characters such as U+2028 (Line Separator, LS ) or U+2029 (Paragraph Separator, PS ). (string slicing)

i := strings.IndexFunc(line, unicode.IsSpace)

strings.LastIndexFunc(line, unicode.IsSpace)

## 3.5. String Formatting with the Fmt Package

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

- %% 	A literal % character
- %b 	An integer value as a binary (base 2) number, or (advanced) a floating-
		point number in scientific notation with a power of 2 exponent
- %c 	An integer code point value as a Unicode character
- %d 	An integer value as a decimal (base 10) number
- %e 	A floating-point or complex value in scientific notation with e
- %E 	A floating-point or complex value in scientific notation with E
- %f 	A floating-point or complex value in standard notation
- %g 	A floating-point or complex value using %e or %f , whichever produces the
		most compact output
- %G 	A floating-point or complex value using %E or %f , whichever produces the
		most compact output
- %o 	An integer value as an octal (base 8) number
- %p 	A value’s address as a hexadecimal (base 16) number with a prefix of 0x
		and using lowercase for the digits a – f (for debugging)
- %q 	The string or []byte as a double-quoted string, or the integer as a single-
		quoted string, using Go syntax and using escapes where necessary
- %s 	The string or []byte as raw UTF-8 bytes; this will produce correct
		Unicode output for a text file or on a UTF-8-savvy console
- %t 	A bool value as true or false
- %T 	A value’s type using Go syntax
- %U 	An integer code point value using Unicode notation defaulting to four
		digits; e.g., fmt.Printf("%U", '¶' ) outputs U+00B6
- %v 	A built-in or custom type’s value using a default format, or a custom
		value using its type’s String() method if it exists
- %x 	An integer value as a hexadecimal (base 16) number or a string or
		[]byte value as hexadecimal digits (two per byte), using lowercase for
		the digits a – f
- %X 	An integer value as a hexadecimal (base 16) number or a string or
		[]byte value as hexadecimal digits (two per byte), using uppercase for
		the digits A – F

Modifier Description/result

- space			Makes the verb output “ - ” before negative numbers and a space before positive numbers or to put spaces between the bytes printed when using the %x or %X verbs; e.g., fmt.Printf("% X", "←" ) outputs E2 86 92
- "#"			Makes the verb use an “alternative” output format:
- %#o 			outputs octal with a leading 0
- %#p			outputs a pointer without the leading 0x
- %#q 			outputs a string or []byte as a raw string (using backticks) if possible—otherwise outputs a double-quoted string
- %#v 			outputs a value as itself using Go syntax
- %#x 			outputs hexadecimal with a leading 0x
- %#X			outputs hexadecimal with a leading 0X
- "+"			Makes the verb output + or - for numbers, ASCII characters (with others escaped) for strings, and field names for struct s
- "-"			Makes the verb left-justify the value (the default is to right-justify)
- 0				Makes the verb pad with leading 0 s instead of spaces
- n.m, n .m 	For numbers, makes the verb output a floating-point or complex value using n (of type int ) characters (or more if necessary to avoid truncation) and with m (of type int ) digits after the decimal point(s). For strings n specifies the minimum field width, and will result in space padding if the string has too few characters, and .m specifies the maximum number of the string’s characters to use (going from left to right), and will result in the string being truncated if it is too long. Either or both of m and n can be replaced with * in which case their values are taken from the arguments. Either n or .m may be omitted.

The way that fmt.Print() and fmt.Fprint() handle whitespace is subtly different from the fmt.Println() and fmt.Fprintln() functions. As a rule of thumb the former are most useful for printing a single value or for “converting” a value to a string without error checking

### 3.5.1. Formatting Booleans

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

### 3.5.2. Formatting Integers

```asp
fmt.Printf("|%b|%9b|%-9b|%09b|% 9b|\n", 37, 37, 37, 37, 37)
|100101|···100101|100101···|000100101|···100101|
```

```asp
fmt.Printf("|%o|%#o|%# 8o|%#+ 8o|%+08o|\n", 41, 41, 41, 41, -41)
|51|051|·····051|····+051|-0000051|
```

```asp
i := 3931
fmt.Printf("|%x|%X|%8x|%08x|%#04X|0x%04X|\n", i, i, i, i, i, i)
|f5b|F5B|·····f5b|00000f5b|0X0F5B|0x0F5B|
```

```asp
i = 569
fmt.Printf("|$%d|$%06d|$%+06d|$%s|\n", i, i, i, Pad(i, 6, '*'))
|$569|$000569|$+00569|$***569|
```

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

### 3.5.3. Formatting Characters

fmt.Printf("%d %#04x %U '%c'\n", 0x3A6, 934, '\u03A6', '\U000003A6')

934·0x03a6·U+03A6·'Φ'

### 3.5.4. Formatting Floating-Point Numbers

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

### 3.5.5. Formatting Strings and Slices

P. 101

```
slogan := "End Óréttlæti♥"
fmt.Printf("%s\n%q\n%+q\n%#q\n", slogan, slogan, slogan, slogan)

chars := []rune(slogan)
fmt.Printf("%x\n%#x\n%#X\n", chars, chars, chars)

bytes := []byte(slogan)
fmt.Printf("%s\n%x\n%X\n% X\n%v\n", bytes, bytes, bytes, bytes, bytes)
```

### 3.5.6. Formatting for Debugging

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

## 3.6. Other String-Related Packages

P. 106

### 3.6.1. The Strings Package

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

### 3.6.2. The Strconv Package

P. 113

strconv.ParseBool() = convert string representation of a truth value into a bool

The Strconv Package’s Functions #1

- strconv.AppendBool(bs, b)						bs with "true" or "false" appended dependingon bool b
- strconv.AppendFloat(bs, f, fmt, prec, bits) 	bs with float64 f appended; see strconv.FormatFloat() for the other parameters
- strconv.AppendInt(bs, i, base)					bs with int64 i appended using the given base
- strconv.AppendQuote(bs, s)						bs with s appended using strconv.Quote()
- strconv.AppendQuoteRune(bs, char)				bs with rune char appended using strconv.QuoteRune()
- strconv.AppendQuoteRuneToASCII(bs, char) 		bs with rune char appended using strconv.QuoteRuneToASCII()
- strconv.AppendQuoteToASCII(bs, s)				bs with s appended using strconv.QuoteToASCII()
- strconv.AppendUInt(bs, u, base)					bs with uint64 u appended using the given base
- strconv.Atoi(s)									string s converted to an int , and an error or nil ; see also strconv.ParseInt()
- strconv.CanBackquote(s)							true if s can be represented in Go syntax using backticks
- strconv.FormatBool(tf)							"true" or "false" depending on bool tf							
- strconv.FormatFloat(f, fmt, prec, bits)			float64 f as a string . The fmt is a byte corresponding to an fmt.Print() verb, 'b' for %b , 'e' for %e , etc. (see Table 3.4, 95 ➤). The prec is thenumber of digits after the decimal point for an fmt of 'e' , 'E' , and 'f' ; or the total number of digits for a fmt of 'g' and 'G' —use -1 to request the smallest number of digits that can be used while preserving accuracy going the other way (e.g., using strconv.ParseFloat() ). The bits affects rounding and is usually 64.
- strconv.FormatInt(i, base)						int64 i as a string in base base
- strconv.FormatUInt(u, base)						uint64 u as a string in base base
- strconv.IsPrint(c)								true if rune c is a printable character
- strconv.Itoa(i)									int i as a string using base 10; see also strconv.FormatInt()

- strconv.ParseBool(s)					true and nil if s is "1" , "t" , "T" , "true" , "True" , or "TRUE" ; false and nil if s is "0" , "f" , "F" , "false" , "False" , or "FALSE" ; false and an error otherwise
- strconv.ParseFloat(s, bits)				A float64 and nil if s is parseable as a floating-point number, or 0 and an error ; bits should be 64; but use 32 if converting to a float32
- strconv.ParseInt( s, base, bits)		An int64 and nil if s is parseable as an integer, or 0 and an error ; a base of 0 means the base will be deduced from s (a leading "0x" or "0X" means base 16, a leading "0" means base 8; otherwise base 10), or a specific base (2–36) can be given; bits should be 0 if converting to an int or the bit size if converting to a sized integer (e.g., 16 for an int16)
- strconv.ParseUint( s, base, bits)		A uint64 and nil or 0 and an error —just the same as strconv.ParseInt() apart from being unsigned
- strconv.Quote(s) 						A string using Go’s double-quoted string syntax to represent string s ; see also Table 3.1 (83 ➤)
- strconv.QuoteRune(char)					A string using Go’s single-quoted string syntax to represent Unicode code point char of type rune
- strconv.QuoteRuneToASCII(char) 			A string using Go’s single-quoted string syntax to represent Unicode code point char of type rune , using an escape sequence for a non-ASCII character
- strconv.QuoteToASCII(s)					A string using Go’s double-quoted string syntax to represent string s , using escape sequences for non-ASCII characters
- strconv.Unquote(s)						A string that contains the Go syntax single-quoted character or double-quoted or backtick-quoted string in string s and an error
- strconv.UnquoteChar(s, b)				A rune (the first character), a bool (whether the first character’s UTF-8 representation needs more than one byte), a string (the rest of the string), and an error ; if b is set to a single or double quote that quote must be escaped

strings.TrimSpace()
scan functions

return i in a binary base :<br />
fmt.Println(strconv.FormatInt(int64(i), 2))

### 3.6.3. The Utf8 Package

utf8.DecodeRuneInString()
utf.DecodeLastRuneInString() (to get first and last characters in a string)

The Utf8 Package’s Functions

- utf8.DecodeLastRune(b)				The last rune in b and the number of bytes it occupies, or U+FFFD (the Unicode replacement character, ? ) and 0, if b doesn’t end with a valid rune
- utf8.DecodeLastRuneInString(s)		The same as utf8.DecodeLastRune(), only it takes a string as input
- utf8.DecodeRune(b)		 			The first rune in b and the number of bytes it occupies, or U+FFFD (the Unicode replacement character, ? ) and 0, if b doesn’t start with a valid rune
- utf8.DecodeRuneInString(s) 			The same as utf8.DecodeRune() , only it takes a string as input
- utf8.EncodeRune( b, c)				Writes c into b as UTF-8 bytes and returns the number of bytes written ( b must have enough space)
- utf8.FullRune(b) 					true if b begins with a UTF-8-encoded rune
- utf8.FullRuneInString(b) 			true if s begins with a UTF-8-encoded rune
- utf8.RuneCount(b) 					Same as utf8.RuneCountInString() but works on a []byte
- utf8.RuneCountInString(s)			The number of rune s in s ; this may be less than len(s) if s contains non-ASCII characters
- utf8.RuneLen(c)						The number of bytes needed to encode c
- utf8.RuneStart(x) 					true if byte x could be the first byte of a rune
- utf8.Valid(b) 						true if b ’s bytes represent valid UTF-8-encoded rune s
- utf8.ValidString(s)					true if s ’s bytes represent valid UTF-8-encoded rune s

### 3.6.4. The Unicode Package

The Unicode Package’s Functions

- unicode.Is(table, c)			true if c is in the table (see text)
- unicode.IsControl(c)			true if c is a control character
- unicode.IsDigit(c)			true if c is a decimal digit
- unicode.IsGraphic(c)			true if c is a “graphic” character such as a letter, number, punctuation mark, symbol, or space
- unicode.IsLetter(c) 			true if c is a letter
- unicode.IsLower(c) 			true if c is a lowercase letter
- unicode.IsMark(c) 			true if c is a mark character
- unicode.IsOneOf(tables, c) 	true if c is in one of the tables
- unicode.IsPrint(c) 			true if c is a printable character
- unicode.IsPunct(c) 			true if c is a punctuation character
- unicode.IsSpace(c) 			true if c is a whitespace character
- unicode.IsSymbol(c) 			true if c is a symbol character
- unicode.IsTitle(c) 			true if c is a title-case letter
- unicode.IsUpper(c) 			true if c is an uppercase letter
- unicode.SimpleFold(c) 		A case-folded copy of the given c
- unicode.To(case, c) 			The case version of c where case is unicode.LowerCase, unicode.TitleCase, or unicode.UpperCase
- unicode.ToLower(c) 			The lowercase version of c
- unicode.ToTitle(c) 			The title-case version of c
- unicode.ToUpper(c) 			The uppercase version of c

### 3.6.5. The Regexp Package

Although replacements can be referred to by number or by name (e.g., $2 , $filename ), it is safest to use braces as delimiters (e.g., ${2} , ${filename} ).

The Regexp Package’s Functions

- regexp.Match(p, b)			true and nil if p matches b of type []byte
- regexp.MatchReader(p, r) 		true and nil if p matches the text read by r of type io.RuneReader
- regexp.MatchString(p, s) 		true and nil if p matches s
- regexp.QuoteMeta(s)			A string with all regexp metacharacters safely quoted
- regexp.MustCompile(p) 		A *regexp.Regexp and nil if p compiles successfully; see Tables 3.18 and 3.19 (➤ 124–125)
- regexp.CompilePOSIX(p)		A *regexp.Regexp and nil if p compiles successfully; see Tables 3.18 and 3.19 (➤ 124–125)
- regexp.MustCompile(p)			A *regexp.Regexp if p compiles successfully, otherwise panics; see Tables 3.18 and 3.19 (➤ 124–125)
- regexp.MustCompilePOSIX(p) 	A *regexp.Regexp if p compiles successfully, otherwise panics; see Tables 3.18 and 3.19 (➤ 124–125)

The Regexp Package’s Escape Sequences

- \c Literal character c ; e.g., \* is a literal * rather than a quantifier
- \000 Character with the given octal code point
- \xHH Character with the given 2-digit hexadecimal code point
- \x{HHHH} Character with the given 1–6-digit hexadecimal code point
- \a ASCII bell (BEL) ≡ \007
- \f ASCII formfeed (FF) ≡ \014
- \n ASCII linefeed (LF) ≡ \012
- \r ASCII carriage return (CR) ≡ \015
- \t ASCII tab (TAB) ≡ \011
- \v ASCII vertical tab (VT) ≡ \013
- \Q...\E Matches the ... text literally even if it contains characters like *

The Regexp Package’s Character Classes

- [chars] Any character in chars
- [^chars] Any character not in chars
- [:name:]	Any ASCII character in the name character class:
			[[:alnum:]]	≡ [0-9A-Za-z]			[[:lower:]]	≡ [a-z]  
			[[:alpha:]]	≡ [A-Za-z]				[[:print:]]	≡ [ -~]  
			[[:ascii:]]	≡ [\x00-\x7F]			[[:punct:]]	≡ [!-/:-@[-`{-~]  
			[[:blank:]]	≡ [ \t]					[[:space:]]	≡ [ \t\n\v\f\r]  
			[[:cntrl:]]	≡ [\x00-\x1F\x7F]		[[:upper:]]	≡ [A-Z]  
			[[:digit:]]	≡ [0-9]					[[:word:]]	≡ [0-9A-Za-z_]  
			[[:graph:]]	≡ [!-~]					[[:xdigit:]] ≡ [0-9A-Fa-z]			 
- [:^name:] Any ASCII character not in the name character class
- . Any character (including newline if flag s is set)
- \d Any ASCII digit: [0-9]
- \D Any ASCII nondigit: [^0-9]
- \s Any ASCII whitespace: [ \t\n\f\r]
- \S Any ASCII nonwhitespace: [^ \t\n\f\r]
- \w Any ASCII “word” character: [0-9A-Za-z_]
- \W Any ASCII non-“word” character: [^0-9A-Za-z_]
- \pN Any Unicode character in the N one-letter character class; e.g., \pL to match a Unicode letter
- \PN Any Unicode character not in the N one-letter character class; e.g., \PL to match a Unicode nonletter
- \p{Name} Any Unicode character in the Name character class; e.g., \p{Ll} matches lowercase letters, \p{Lu} matches uppercase letters, and \p{Greek} matches Greek characters
- \P{Name} Any Unicode character not in the Name character class

The Regexp Package’s Zero-Width Assertions

- ^ Start of text (or start of line if flag m is set)
- $ End of text (or end of line if flag m is set)
- \A Start of text
- \z End of text
- \b Word boundary ( \w followed by \W or \A or \z ; or vice versa)
- \B Not a word boundary
- e? or e{0,1} Greedily match zero or one occurrence of expression e
- e+ or e{1,} Greedily match one or more occurrences of expression e
- e* or e{0,} Greedily match zero or more occurrences of expression e
- e{m,} Greedily match at least m occurrences of expression e
- e{,n} Greedily match at most n occurrences of expression e
- e{m,n} Greedily match at least m and at most n occurrences of expression e
- e{m} or e{m}? Match exactly m occurrences of expression e
- e?? or e{0,1}? Nongreedily match zero or one occurrence of expression e
- e+? or e{1,}? Nongreedily match one or more occurrences of expression e
- e*? or e{0,}? Nongreedily match zero or more occurrences of expression e
- e{m,}? Nongreedily match at least m occurrences of expression e
- e{,n}? Nongreedily match at most n occurrences of expression e
- e{m,n}? Nongreedily match at least m and at most n occurrences of expression e

The Regexp Package’s Flags and Groups

- i Match case-insensitively (the default is case-sensitive matching)
- m Multiline mode makes ^ and $ match at the start and end of every line (the default is single-line mode)
- s Make . match any character including newlines (the default is for . to match any character except newlines)
- U Make greedy matches nongreedy and vice versa; i.e., swap the meaning of ? after a quantifier (the default is for matches to be greedy unless their quantifier is followed by ? to make them nongreedy)
- (?flags) Apply the given flags from this point on (precede the flag or flags with - to negate)
- (?flags:e) Apply the given flags to expression e (precede the flag or flags with - to negate)
- (e) Group and capture the match for expression e
- (?P<name>e) Group and capture the match for expression e using the capture name name
- (?:e) Group but don’t capture the match for expression e 

P. 124

The * regexp.Regexp Type’s Methods

rx.Expand(...) 					Performs the $ replacements done by the ReplaceAll() method—rarely used directly (advanced)
rx.ExpandString(...) 			Performs the $ replacements done by the ReplaceAllString() method—rarely used directly (advanced)
rx.Find(b) 						A []byte with the leftmost match or nil
rx.FindAll(b, n) 				A [][]byte of all nonoverlapping matches or nil
rx.FindAllIndex(b, n) 			An [][]int (a slice of 2-item slices) each identifying a match or nil ; e.g., b[pos[0]:pos[1]] where pos is one of the 2-item slices
rx.FindAllString(s, n) 			A []string of all nonoverlapping matches or nil
rx.FindAllStringIndex(s, n) 	An [][]int (a slice of 2-item slices) each identifying a match or nil ; e.g., s[pos[0]:pos[1]] where pos is one of the 2-item slices
rx.FindAllStringSubmatch(s, n) 	A [][]string (a slice of string slices where each string corresponds to a capture) or nil
rx.FindAllStringSubmatchIndex(s, n) An [][]int (a slice of 2-item int slices that correspond to captures) or nil
rx.FindAllSubmatch(b, n) 		A [][][]byte (a slice of slices of []byte s where each []byte corresponds to a capture) or nil
rx.FindAllSubmatchIndex(b, n) 	An [][]int (a slice of 2-item int slices that correspond to captures) or nil
rx.FindIndex(b) 				A 2-item []int identifying the leftmost match; e.g., b[pos[0]:pos[1]] where pos is the 2-item slice, or nil
rx.FindReaderIndex(r) 			A 2-item []int identifying the leftmost match or nil
rx.FindReaderSubmatchIndex(r) 	An []int identifying the leftmost match and captures or nil
rx.FindString(s)				The leftmost match or an empty string
rx.FindStringIndex(s)			A 2-item []int identifying the leftmost match or nil
rx.FindStringSubmatch(s)		A []string with the leftmost match and captures or nil
rx.FindStringSubmatchIndex(s)	An []int identifying the leftmost match and captures or nil

rx.FindSubmatch(b)					A [][]byte with the leftmost match and captures or nil
rx.FindSubmatchIndex(b)				A [][]byte with the leftmost match and captures or nil
rx.LiteralPrefix()					The possibly empty prefix string that the regexp must begin with and a bool indicating whether the whole regexp is a literal string match
rx.Match(b) 						true if the regexp matches b
rx.MatchReader(r) 					true if the regexp matches r of type io.RuneReader
rx.MatchString(s) 					true if the regexp matches s
rx.NumSubexp() 						How many parenthesized groups the regexp has
rx.ReplaceAll(b, br) 				A []byte that is a copy of b with every match replaced with br of type []byte with $ replacements (see text)
rx.ReplaceAllFunc(b, f) 			A []byte that is a copy of b with every match replaced with the return value of a call to function f of type func([]byte) []byte and whose argument is a match
rx.ReplaceAllLiteral(b, br)			A []byte that is a copy of b with every match replaced with br of type []byte
rx.ReplaceAllLiteralString(s, sr)	A string that is a copy of s with every match replaced with sr of type string replacements 
rx.ReplaceAllString(s, sr) 			A string that is a copy of s with every match replaced with sr of type string with $ replacements (see text)
rx.ReplaceAllStringFunc(s, f) 		A string that is a copy of s with every match replaced with the return value of a call to function f of type func(string) string and whose argument is a match
rx.String()							A string containing the regexp pattern
rx.SubexpNames()					A []string (which must not be modified), containing the names of all the named subexpressions

regexp to find duplicate words... P. 126-127

for key:value

```
keyValueRx := regexp.MustCompile(`\s*([[:alpha:]]\w*)\s*:\s*(.+)`)
```

for name="value" or name='value'

```
attrValueRx := regexp.MustCompile(regexp.QuoteMeta(attrName) + `=(?:"([^"]+)"|'([^']+)')`)
```

how to simplify whitespaces with regexp

[...]

unaccented := UnaccentedLatin1(latin1) to perform the conversion

for every "String" regexp function there is a function that operates on []bytes

## 3.7. Example: M3u2pls

P. 130

done the m3u2pls prog

## 3.8. Exercices

1. prog to do the conversion from .pls to .m3u

2. Soundex algorithm




