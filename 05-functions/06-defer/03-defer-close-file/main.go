package main

import (
	"io"
	"os"
)

func main() {
	src, err := os.Open("src.txt")

	if err != nil {
		panic(err)
	}

	defer src.Close()

	dst, err := os.Create("dst.txt")
	if err != nil {
		panic(err)
	}

	defer dst.Close()
	// bs := make([]byte, 5)
	// src.Read(bs)
	// dst.Write(bs)

	_, err = io.Copy(dst, src)
	if err != nil {
		panic(err)
	}
}
