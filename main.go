package main

import (
	"github.com/zer0go/kilo-in-go/pkg"
	"os"
)

func main() {
	var fileName string
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	pkg.EditWithKilo(pkg.InputConfig{
		Data:     []byte{},
		FileName: fileName,
	})
}
