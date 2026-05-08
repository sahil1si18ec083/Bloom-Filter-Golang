# Bloom Filter in Golang — Complete Beginner-Friendly Explanation

## What is a Bloom Filter?

A Bloom Filter is a probabilistic data structure.

It is used to quickly check:

* Whether an item may exist
* Or definitely does not exist

It is memory efficient and very fast.

Bloom Filter can return:

* `Definitely NOT present`
* `Maybe present`

It never gives false negatives.
But it can give false positives.

---

# Real World Usage

Bloom Filters are used in:

* Databases
* Caching systems
* Distributed systems
* URL filtering
* Duplicate checking
* Search engines

Companies and systems using Bloom Filters:

* Redis
* Cassandra
* Bigtable
* HBase

---

# How Bloom Filter Works

Bloom Filter internally uses:

* Bit array
* Multiple hash functions

Example:

```text
Index:
0 1 2 3 4 5 6 7 8 9

Values:
F F F F F F F F F F
```

Initially all values are false.

---

# Add Operation

Suppose we add:

```text
cat
```

Hash functions generate indexes:

```text
2 5 8
```

Now these indexes become true.

```text
F F T F F T F F T F
```

---

# Search Operation

Suppose we search:

```text
cat
```

Again hashes generate:

```text
2 5 8
```

All are true.

So:

```text
Maybe present
```

---

# Full Code with Explanation

```go
package main

import (
	"fmt"
	"hash/fnv"
)

// BloomFilter structure
// bitArray stores true/false values
// Size stores total size of filter

type BloomFilter struct {
	bitArray []bool
	Size     uint32
}

// Number of hash functions
const hashFunctionCount uint32 = 5

// Constructor function
// Creates bloom filter of given size

func NewBloomFilter(size uint32) *BloomFilter {
	return &BloomFilter{
		bitArray: make([]bool, size),
		Size:     size,
	}
}

// Add function
// Adds item into bloom filter

func (bf *BloomFilter) Add(s string) {

	// Run multiple hash functions
	for i := uint32(0); i < hashFunctionCount; i++ {

		// Generate index
		index := bf.Hash(s, i)

		// Mark index as true
		bf.bitArray[index] = true
	}
}

// Hash function
// Converts string into index

func (bf *BloomFilter) Hash(s string, i uint32) uint32 {

	// Create FNV hash object
	h := fnv.New32a()

	// Convert string into bytes
	h.Write([]byte(s))

	// Generate final index
	return (h.Sum32() + i) % bf.Size
}

// Contains function
// Checks whether item may exist

func (bf *BloomFilter) Contains(s string) bool {

	// Check all hash indexes
	for i := uint32(0); i < hashFunctionCount; i++ {

		// Generate index
		index := bf.Hash(s, i)

		// If even one index is false
		// item definitely does not exist
		if !bf.bitArray[index] {
			return false
		}
	}

	// All indexes are true
	// item may exist
	return true
}

func main() {

	fmt.Println("Bloom Filter")

	// Create bloom filter of size 100
	bf := NewBloomFilter(100)

	// Add items
	bf.Add("cat")
	bf.Add("dog")

	// Search items
	fmt.Println(bf.Contains("cat"))
	fmt.Println(bf.Contains("dog"))
	fmt.Println(bf.Contains("elephant"))
	fmt.Println(bf.Contains("myself"))
}
```

---

# Understanding Important Parts

## 1. Why Multiple Hash Functions?

If only one hash function is used:

* More collisions happen
* False positives increase

If too many hash functions are used:

* Operations become slow

So usually:

```text
3 to 7 hash functions
```

are used.

---

# 2. Why uint32?

Because:

* Hash values are never negative
* FNV hash returns uint32
* Array indexes should be positive

---

# 3. Why make()?

```go
make([]bool, size)
```

creates memory for the slice.

Without make:

```go
var arr []bool
```

memory is not allocated properly.

---

# 4. Why modulo (%)?

Hash values are very large.

Example:

```text
981273981273
```

But array size may be:

```text
100
```

So we use:

```go
hash % size
```

to keep index within range.

---

# Time Complexity

| Operation | Complexity |
| --------- | ---------- |
| Add       | O(k)       |
| Search    | O(k)       |

Where:

```text
k = number of hash functions
```

---

# Limitations

Bloom Filter:

* Can give false positives
* Cannot give false negatives
* Cannot delete items in basic implementation

---

# Future Improvements

You can improve this project using:

* Counting Bloom Filter
* Bitset optimization
* Better hashing
* Redis Bloom Filter
* Scalable Bloom Filter

---

# Sample Output

```text
Bloom Filter
true
true
false
false
```

---

# Key Learning Concepts

This project teaches:

* Hashing
* Golang structs
* Methods
* Slices
* Loops
* Probabilistic data structures
* Backend system optimization

---

# Final Summary

Bloom Filter is:

* Fast
* Memory efficient
* Scalable
* Used in real backend systems

It is one of the most important probabilistic data structures used in distributed systems and caching.
