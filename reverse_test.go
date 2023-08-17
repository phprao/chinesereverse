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

// go test -v -run TestWithExtraDictFile
func TestWithExtraDictFile(t *testing.T) {
	WithExtraDictFile("dict2.txt")

	fmt.Println(SimplifiedToTraditional("中"))
	fmt.Println(TraditionalToSimplified("种"))
}
