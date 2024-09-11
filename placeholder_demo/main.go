package main

import "fmt"

func main() {
	comp := Comp{
		Boss:   "boss",
		Worker: "worker",
	}

	fmt.Printf("%v\n", comp)
	fmt.Printf("%+v\n", comp)
	fmt.Printf("%#v\n", comp)
	fmt.Printf("%T\n", comp)
	fmt.Printf("%q\n", comp)
	fmt.Printf("%t\n", false)

	fmt.Printf("%v\n", -2e-2)

	fmt.Printf("%v\n", 0x5p-2)
	fmt.Printf("%b\n", 0x5p-2)
	fmt.Printf("%x\n", 0x5p-2)
	fmt.Printf("%X\n", 0x5p-2)

	fmt.Printf("%X\n", 255)
	fmt.Printf("%+d\n", 255)
	fmt.Printf("% d\n", 255)
	fmt.Printf("%#o\n", 255)
	fmt.Printf("%U\n", 255)
	fmt.Printf("%q\n", 255)
}

type Comp struct {
	Boss   string `json:"boss"`
	Worker string `json:"worker"`
}
