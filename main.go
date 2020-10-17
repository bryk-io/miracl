package main

import (
	"fmt"

	"go.bryk.io/miracl/core"
)

func main() {
	fmt.Printf("upstream commit: %s\n", core.Version)
}
