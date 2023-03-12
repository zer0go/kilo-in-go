package main

import "os"

func main() {
	var fileName string
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	EditWithKilo(InputConfig{
		Data:     []byte{},
		FileName: fileName,
	})
}
