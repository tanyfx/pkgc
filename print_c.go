package pkgc

import (
	"fmt"
	"github.com/tanyfx/pkgd"
)

func PrintInfo() {
	fmt.Printf("pkgc branch master:\t%s\n", pkgd.GetInfo())
}
