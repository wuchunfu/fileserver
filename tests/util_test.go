package tests

import (
	"fileserver/utils"
	"fmt"
	"testing"
)

func TestIsAbsolutePath(t *testing.T) {
	fmt.Println(utils.IsAbsolutePath("/a/b/c"))
	fmt.Println(utils.IsAbsolutePath("C:\\Windows\\system"))
	fmt.Println(utils.IsAbsolutePath("C://Windows//system"))
	fmt.Println(utils.IsAbsolutePath("Windows\\system"))
	fmt.Println(utils.IsAbsolutePath("a/b/c"))
}
