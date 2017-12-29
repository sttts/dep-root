package main

import (
	"fmt"

	"github.com/sttts/dep-a"
)

func main() {
	fmt.Printf("release of dep-root calling dep-a: %s\n", a.A())
}
