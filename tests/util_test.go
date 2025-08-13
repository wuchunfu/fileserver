package tests

import (
	"fmt"
	"testing"

	"github.com/wuchunfu/fileserver/utils/filex"
)

func TestIsAbsolutePath(t *testing.T) {
	fmt.Println(filex.IsAbsolutePath("/a/b/c"))
	fmt.Println(filex.IsAbsolutePath("C:\\Windows\\system"))
	fmt.Println(filex.IsAbsolutePath("C://Windows//system"))
	fmt.Println(filex.IsAbsolutePath("Windows\\system"))
	fmt.Println(filex.IsAbsolutePath("a/b/c"))
}
