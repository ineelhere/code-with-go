This Go program demonstrates how to work with Unicode characters in a string.

It includes functionality to print the length of a string in bytes, iterate over each byte to print its hexadecimal representation, count the number of runes in the string, and iterate over each rune to print its Unicode code point and starting index.

Additionally, it uses the `utf8.DecodeRuneInString` function to decode runes and provides an example function, `examineRune`, to examine specific runes.

```go
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    const s = "สวัสดี"

    // Print the length of the string in bytes
    fmt.Println("Len:", len(s))

    // Iterate over each byte and print its hexadecimal representation
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
    fmt.Println()

    // Print the count of runes in the string
    fmt.Println("Rune count:", utf8.RuneCountInString(s))

    // Iterate over each rune and print its Unicode code point and starting index
    for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }

    fmt.Println("\nUsing DecodeRuneInString")

    // Iterate over each rune using DecodeRuneInString and print its Unicode code point and starting index
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width

        // Examine the rune using the examineRune function
        examineRune(runeValue)
    }
}

// examineRune examines specific runes and prints custom messages
func examineRune(r rune) {
    if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
}
```

### Output
```
Len: 18
e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b8 aa e0 b8 94 e0 b8 b5
Rune count: 6
U+0E2A 'ส' starts at 0
U+0E27 'ว' starts at 3
U+0E31 'ั' starts at 6
U+0E2A 'ส' starts at 9
U+0E14 'ด' starts at 12
U+0E35 'ี' starts at 15

Using DecodeRuneInString
U+0E2A 'ส' starts at 0
found so sua
U+0E27 'ว' starts at 3
U+0E31 'ั' starts at 6
U+0E2A 'ส' starts at 9
found so sua
U+0E14 'ด' starts at 12
U+0E35 'ี' starts at 15
```

Explanation:

1. The program initializes a constant string `s` with the value "สวัสดี", which means "hello" in Thai.

2. The length of the string in bytes is printed using `len(s)`.

3. A loop iterates over each byte of the string, printing its hexadecimal representation.

4. The `utf8.RuneCountInString` function is used to count the number of runes in the string.

5. A loop using `range` iterates over each rune in the string, printing its Unicode code point and starting index.

6. Another loop uses `utf8.DecodeRuneInString` to iterate over each rune and print its Unicode code point and starting index.

7. The `examineRune` function is called to examine specific runes and print custom messages based on their values.