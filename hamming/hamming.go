// Package hamming implements Distance function
package hamming

import "fmt"

// Distance returns Hamming distance two DNA strands or error if they have different length
func Distance(a, b string) (int, error) {
	var l = len(a)
	if l != len(b) {
		return 0, fmt.Errorf("distance: %s %s have different length", a, b)
	}

	var dist int
	for i := 0; i < l; i++ {
		if a[i] != b[i] {
			dist++
		}
	}

	return dist, nil
}
