package main

import (
	"fmt"
	"unsafe"
)

type NameStruct struct {
	Name string
}

type T2 struct {
	a int
	b int
}

func main() {

	//1
	v := (-2) % 4
	fmt.Println(v)

	//2,3,4,5
	a := uint8(255)
	b := uint8(127)
	c := a + b
	c2 := a & b
	c3 := a ^ b
	c4 := a &^ b
	fmt.Println(a, "+", b, "=", c)
	fmt.Println(a, "&", b, "=", c2)
	fmt.Println(a, "^", b, "=", c3)
	fmt.Println(a, "&^", b, "=", c4)

	//6
	ns := &NameStruct{}
	fmt.Printf("%T", ns)

	//7,8
	fmt.Println()
	t := T2{a: 123, b: 234}
	fmt.Println(unsafe.Sizeof(t))
	fmt.Println(unsafe.Sizeof(&t))

	//9
	const k int = 1
	//k = 2

	//10
	const d = 5
	const l complex128 = d

	//11
	//var string int = 1
	//var ak string = "a"
	//fmt.Printf("%T %T", string, ak)

	//14
	const (
		w1 = 1
		w2
		w3
	)
	fmt.Println(w1, w2, w3)

	//15
	const(
		a1 = iota + 1
		a2
		a3
	)
	fmt.Println(a1,a2,a3)

	//16
	const(
		b1 = iota + 1
		b2 = iota + 2
		b3 = iota + 3
	)
	fmt.Println(b1,b2,b3)

	//17
	const (
		n1 = 1 << iota
		n2 = 1 << iota
		n3 = 1 << iota
	)
	fmt.Println(n1, n2, n3)


	//18
	type ByteSize int64

	const (
		_ = iota
		kb ByteSize = 1 << (10 * iota)
		mb
		gb
	)

	fmt.Println(kb, mb, gb)


	//19
	const( max = 10)
	const(
		u1 = max - iota
		u2
		u3
	)
	fmt.Println(u1,u2,u3)

	//20

	const(
		x1 = string(iota + 'a')
		y1
		z1
	)
	fmt.Println(x1,y1,z1)
}

