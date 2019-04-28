// Package collatzconjecture implements CollatzConjecture function
package collatzconjecture

import "fmt"

// CollatzConjecture returns number of iterations required to reach 1
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("collatzconjecture: %d less than or equal to zero", n)
	}
	var it = 0
	for n > 1 {
		if n%2 == 1 {
			n = 3*n + 1
		} else {
			n = n / 2
		}
		it++
	}

	return it, nil
}
