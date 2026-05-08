package main

import (
	"fmt"
	"hash/fnv"
)

type BloomFilter struct {
	bitArray []bool
	Size     uint32
}

const hashFunctionCount uint32 = 5

func NewBloomFilter(size uint32) *BloomFilter {
	return &BloomFilter{
		bitArray: make([]bool, size),
		Size:     size,
	}
}
func (bf *BloomFilter) Add(s string) {
	for i := uint32(0); i < hashFunctionCount; i++ {
		index := bf.Hash(s, i)
		bf.bitArray[index] = true
	}

}
func (bf *BloomFilter) Hash(s string, i uint32) uint32 {
	h := fnv.New32a()

	h.Write([]byte(s))

	return (h.Sum32() + i) % bf.Size

}
func (bf *BloomFilter) Contains(s string) bool {
	for i := uint32(0); i < hashFunctionCount; i++ {
		index := bf.Hash(s, i)
		if !bf.bitArray[index] {
			return false
		}
	}
	return true

}
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
