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






