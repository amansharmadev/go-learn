package main

import (
	"fmt"
	"os"
)

type inputFile struct {
	filepath  string
	separator string
	pretty    bool
}

func main() {
	fmt.Println(len(os.Args))
	fmt.Println(os.Args)
}
