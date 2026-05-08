# Bloom Filter in Go

This project is a simple implementation of a Bloom filter in Go.

A Bloom filter is a probabilistic data structure used to check whether an item may exist in a set. It is designed to be very fast and memory efficient, especially when compared with storing every item directly in a map, list, or database.

This implementation is written for learning purposes. It shows the basic idea behind Bloom filters using a boolean array, a hash function, and multiple hash positions.

## Table of Contents

- [What Is a Bloom Filter?](#what-is-a-bloom-filter)
- [Why Use a Bloom Filter?](#why-use-a-bloom-filter)
- [Project Structure](#project-structure)
- [How This Program Works](#how-this-program-works)
- [Code Explanation](#code-explanation)
- [Running the Program](#running-the-program)
- [Example Output](#example-output)
- [Time and Space Complexity](#time-and-space-complexity)
- [False Positives](#false-positives)
- [Limitations](#limitations)
- [Possible Improvements](#possible-improvements)

## What Is a Bloom Filter?

A Bloom filter is used to test whether an element is a member of a set.

Unlike a normal set or map, a Bloom filter does not store the actual values. Instead, it stores information about those values in a bit array.

When you ask a Bloom filter whether an item exists, it can answer in two ways:

- `false`: the item is definitely not present
- `true`: the item may be present

This means a Bloom filter can produce false positives, but it does not produce false negatives.

For example:

- If `"cat"` was added, `Contains("cat")` should return `true`
- If `"elephant"` was not added, `Contains("elephant")` will usually return `false`
- In some cases, `Contains("elephant")` may return `true` even though it was never added

That last case is called a false positive.

## Why Use a Bloom Filter?

Bloom filters are useful when you want a fast and memory-efficient way to check membership.

Common use cases include:

- Checking whether a username might already exist
- Avoiding unnecessary database lookups
- Detecting duplicate URLs in web crawlers
- Checking whether an email address may be in a blocklist
- Filtering cache misses
- Large-scale systems where storing all values directly would use too much memory

The tradeoff is accuracy. Bloom filters save memory by accepting a small chance of false positives.

## Project Structure

```text
.
├── main.go
└── README.md
```

The project currently contains one Go source file:

- `main.go`: contains the Bloom filter implementation and a small demo in the `main` function

## How This Program Works

This implementation uses:

- A boolean array called `bitArray`
- A fixed Bloom filter size
- The FNV-1a hash function from Go's standard library
- Five hash positions for each string

When a string is added to the filter:

1. The string is passed into a hash function.
2. The hash result is converted into an array index.
3. The program repeats this process five times.
4. Each resulting index in the boolean array is set to `true`.

When checking whether a string exists:

1. The same five indexes are calculated for the string.
2. The program checks whether all five indexes are `true`.
3. If any index is `false`, the string is definitely not present.
4. If all indexes are `true`, the string may be present.

## Code Explanation

### Package and Imports

```go
package main

import (
    "fmt"
    "hash/fnv"
)
```

The program uses:

- `fmt` to print output to the terminal
- `hash/fnv` to create a hash value for each string

FNV is a fast non-cryptographic hash function. It is useful for demonstrations like this, but it is not intended for cryptographic security.

### BloomFilter Struct

```go
type BloomFilter struct {
    bitArray []bool
    Size     uint32
}
```

The `BloomFilter` struct has two fields:

- `bitArray`: stores the filter state
- `Size`: stores the number of available positions in the array

Each position in `bitArray` is either:

- `false`: no inserted item has marked this position yet
- `true`: at least one inserted item has marked this position

### Hash Function Count

```go
const hashFunctionCount uint32 = 5
```

This constant controls how many positions are marked for each inserted string.

In this project, each string affects five positions in the bit array.

Using more hash positions can reduce false positives up to a point, but it also makes insert and lookup operations slightly more expensive.

### Creating a New Bloom Filter

```go
func NewBloomFilter(size uint32) *BloomFilter {
    return &BloomFilter{
        bitArray: make([]bool, size),
        Size:     size,
    }
}
```

`NewBloomFilter` creates and returns a pointer to a new Bloom filter.

For example:

```go
bf := NewBloomFilter(100)
```

This creates a Bloom filter with 100 boolean positions.

At the start, all positions are `false`.

### Adding an Item

```go
func (bf *BloomFilter) Add(s string) {
    for i := uint32(0); i < hashFunctionCount; i++ {
        index := bf.Hash(s, i)
        bf.bitArray[index] = true
    }
}
```

The `Add` method inserts a string into the Bloom filter.

For each value of `i` from `0` to `4`, the method:

1. Calculates an index using `bf.Hash(s, i)`
2. Sets that position in `bitArray` to `true`

The actual string is not stored anywhere. Only the calculated positions are stored.

### Hashing an Item

```go
func (bf *BloomFilter) Hash(s string, i uint32) uint32 {
    h := fnv.New32a()

    h.Write([]byte(s))

    return (h.Sum32() + i) % bf.Size
}
```

The `Hash` method turns a string into an array index.

It does this by:

1. Creating a new FNV-1a 32-bit hash
2. Writing the string bytes into the hash
3. Taking the hash result
4. Adding `i`
5. Using modulo `% bf.Size` to keep the result inside the array bounds

The modulo operation ensures the returned index is always between `0` and `Size - 1`.

For example, if the Bloom filter size is `100`, the index will always be in this range:

```text
0 to 99
```

### Checking an Item

```go
func (bf *BloomFilter) Contains(s string) bool {
    for i := uint32(0); i < hashFunctionCount; i++ {
        index := bf.Hash(s, i)
        if !bf.bitArray[index] {
            return false
        }
    }
    return true
}
```

The `Contains` method checks whether a string may exist in the Bloom filter.

It calculates the same five indexes used by `Add`.

If any of those positions is `false`, the item definitely was not added.

If all positions are `true`, the item may have been added.

### Main Function

```go
func main() {
    fmt.Println("NewBloomFilter")

    bf := NewBloomFilter(100)
    bf.Add("cat")
    bf.Add("dog")

    fmt.Println(bf.Contains("cat"))
    fmt.Println(bf.Contains("dog"))
    fmt.Println(bf.Contains("elephant"))
    fmt.Println(bf.Contains("myself"))
}
```

The `main` function demonstrates the Bloom filter.

It:

1. Creates a Bloom filter with size `100`
2. Adds `"cat"`
3. Adds `"dog"`
4. Checks whether `"cat"` exists
5. Checks whether `"dog"` exists
6. Checks whether `"elephant"` exists
7. Checks whether `"myself"` exists

Since `"cat"` and `"dog"` were added, they should return `true`.

Since `"elephant"` and `"myself"` were not added, they will usually return `false`. However, because Bloom filters can have false positives, an item that was not added may still sometimes return `true`.

## Running the Program

Make sure Go is installed on your machine.

To check your Go version:

```bash
go version
```

Then run the program:

```bash
go run main.go
```

## Example Output

```text
NewBloomFilter
true
true
false
false
```

Output meaning:

- `true` for `"cat"` because it was added
- `true` for `"dog"` because it was added
- `false` for `"elephant"` because it was not added
- `false` for `"myself"` because it was not added

The output for values that were not added can change depending on filter size, hash behavior, and inserted values.

## Time and Space Complexity

Let:

- `k` be the number of hash functions
- `m` be the size of the bit array

In this project:

- `k = 5`
- `m = 100` in the demo

### Add Operation

```text
Time complexity: O(k)
```

Each inserted item is hashed `k` times.

Since `k` is fixed at `5`, insertion is effectively constant time for this implementation.

### Contains Operation

```text
Time complexity: O(k)
```

Each lookup checks `k` positions.

Since `k` is fixed, lookup is also effectively constant time.

### Space Complexity

```text
Space complexity: O(m)
```

The filter stores a fixed-size boolean array.

It does not store the inserted strings themselves.

## False Positives

A false positive happens when the Bloom filter says an item may exist even though it was never added.

This can happen because different strings may mark the same positions in the bit array.

For example:

```text
Add("cat") marks indexes:      10, 11, 12, 13, 14
Add("dog") marks indexes:      40, 41, 42, 43, 44
Contains("fish") checks:       10, 11, 12, 13, 14
```

If `"fish"` happens to check positions that are already `true`, the Bloom filter returns `true` even though `"fish"` was never added.

False positives become more likely when:

- The bit array is too small
- Too many items are inserted
- Too many indexes become `true`
- The hash functions do not distribute values well

## Limitations

This implementation is intentionally simple, but it has some limitations.

### It Uses `[]bool`

The filter uses a boolean slice:

```go
bitArray []bool
```

This is easy to understand, but it is not the most memory-efficient representation.

A production Bloom filter would usually use a compact bitset so each position uses one bit instead of a full boolean value.

### Hashing Is Simplified

The code uses:

```go
return (h.Sum32() + i) % bf.Size
```

This creates five nearby hash positions by adding `i` to the same base hash value.

That is useful for learning, but real Bloom filters usually use multiple independent hash functions or double hashing to spread values more evenly.

### No Delete Operation

Standard Bloom filters do not support safe deletion.

If you set a position back to `false`, you might accidentally remove information needed by another item.

To support deletion, you would need a counting Bloom filter.

### Fixed Size

The Bloom filter size is chosen when it is created:

```go
bf := NewBloomFilter(100)
```

The current implementation does not resize automatically.

If too many values are added, the false positive rate increases.

## Possible Improvements

This project could be improved in several ways:

- Replace `[]bool` with a compact bitset
- Use better hash generation, such as double hashing
- Add a configurable number of hash functions
- Add tests using Go's `testing` package
- Add benchmarks for insert and lookup speed
- Add a false positive rate calculator
- Validate that the filter size is greater than zero
- Create a reusable package instead of keeping everything in `main.go`
- Add support for other data types, not just strings
- Implement a counting Bloom filter to support deletion

## Summary

This project demonstrates the core idea of a Bloom filter:

- It uses less memory than storing all values directly
- It can quickly check whether a value may exist
- It can say when a value definitely does not exist
- It may return false positives
- It does not return false negatives for inserted values

The code is small, readable, and useful for understanding how probabilistic membership checking works in Go.
#   B l o o m - F i l t e r - G o l a n g  
 