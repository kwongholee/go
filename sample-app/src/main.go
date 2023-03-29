package main

import "fmt"

func main() {
	var a int = 1
	const b int = 32
	var c, d, e float32 = 3.14, 31.4, 5.9
	const (
		f = iota // 0
		g        // 1
		h        // 2
	)

	rawLiteral := `아리랑\n아리랑`
	interLiteral := "아리랑\n아리랑"

	fmt.Println(a, b, c, d, e, f, g, h)
	fmt.Println()
	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)
}
