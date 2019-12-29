// Package dog is you best friend
package dog

import "fmt"
//Years converts from human years to dog years
func Years(years int) (int, error) {
	if years < 1 {
		return 0, fmt.Errorf("value provided must be positive: %v given", years)
	}

	return years * 7, nil
}