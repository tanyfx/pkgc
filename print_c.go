package pkgc

import (
	"fmt"
	"github.com/tanyfx/pkgd"
)

func PrintInfo() {
	fmt.Printf("pkgc branch forkc:\t%s\n", pkgd.GetInfo())
}
