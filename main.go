package main

import (
	"github.com/zer0go/kilo-in-go/pkg/kilo"
	"os"
)

func main() {
	var fileName string
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	kilo.EditWithKilo(kilo.InputConfig{
		Data:     []byte{},
		FileName: fileName,
	})
}
