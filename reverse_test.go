package chinesereverse

import (
	"fmt"
	"testing"
)

// go test -v -run TestSToT
func TestSToT(t *testing.T) {
	fmt.Println(SimplifiedToTraditional("连续"))
	fmt.Println(TraditionalToSimplified("連續"))
}
