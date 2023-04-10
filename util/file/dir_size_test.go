package file

import (
	"fmt"
	"testing"
)

func TestDirSize(t *testing.T) {
	size := DirSize("/Users/daile/go/src/hamster/util")
	fmt.Println(size)
}
