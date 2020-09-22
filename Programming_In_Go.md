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