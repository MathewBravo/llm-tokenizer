package main

import "fmt"

func main() {
	test_text := "The original BPE algorithm operates by iteratively replacing the most common contiguous sequences of characters in a target text with unused 'placeholder' bytes. The iteration ends when no sequences can be found, leaving the target text effectively compressed. Decompression can be performed by reversing this process, querying known placeholder terms against their corresponding denoted sequence, using a lookup table. In the original paper, this lookup table is encoded and stored alongside the compressed text."

	freq := make(map[[2]byte]int)

	for i, _ := range test_text[:len(test_text)-1] {
		pair := [2]byte{test_text[i], test_text[i+1]}
		freq[pair]++
	}

	for k, v := range freq {
		fmt.Printf("%c%c: %d\n", k[0], k[1], v)
	}

}
