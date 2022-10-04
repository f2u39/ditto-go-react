package err

import (
	"fmt"
)

// Check error
func E(err error) bool {
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// Panic error
func P(err error) {
	if err != nil {
		panic(err)
	}
}
