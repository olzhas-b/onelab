package main

import (
	"fmt"
	"unsafe"
)
type Changeable0 struct {
 	a byte
	b int
	c bool
	d int
 
}
type Changeable1 struct {
 	a byte
	b int
	d int
	c bool
 
}
type Changeable2 struct {
 	a byte
	c bool
	b int
	d int
 
}
type Changeable3 struct {
 	a byte
	c bool
	d int
	b int
 
}
type Changeable4 struct {
 	a byte
	d int
	b int
	c bool
 
}
type Changeable5 struct {
 	a byte
	d int
	c bool
	b int
 
}
type Changeable6 struct {
 	b int
	a byte
	c bool
	d int
 
}
type Changeable7 struct {
 	b int
	a byte
	d int
	c bool
 
}
type Changeable8 struct {
 	b int
	c bool
	a byte
	d int
 
}
type Changeable9 struct {
 	b int
	c bool
	d int
	a byte
 
}
type Changeable10 struct {
 	b int
	d int
	a byte
	c bool
 
}
type Changeable11 struct {
 	b int
	d int
	c bool
	a byte
 
}
type Changeable12 struct {
 	c bool
	a byte
	b int
	d int
 
}
type Changeable13 struct {
 	c bool
	a byte
	d int
	b int
 
}
type Changeable14 struct {
 	c bool
	b int
	a byte
	d int
 
}
type Changeable15 struct {
 	c bool
	b int
	d int
	a byte
 
}
type Changeable16 struct {
 	c bool
	d int
	a byte
	b int
 
}
type Changeable17 struct {
 	c bool
	d int
	b int
	a byte
 
}
type Changeable18 struct {
 	d int
	a byte
	b int
	c bool
 
}
type Changeable19 struct {
 	d int
	a byte
	c bool
	b int
 
}
type Changeable20 struct {
 	d int
	b int
	a byte
	c bool
 
}
type Changeable21 struct {
 	d int
	b int
	c bool
	a byte
 
}
type Changeable22 struct {
 	d int
	c bool
	a byte
	b int
 
}
type Changeable23 struct {
 	d int
	c bool
	b int
	a byte
 
}
func main() {
	fmt.Print(string(unsafe.Sizeof(Changeable0{})))
	fmt.Print(string(unsafe.Sizeof(Changeable1{})))
	fmt.Print(string(unsafe.Sizeof(Changeable2{})))
	fmt.Print(string(unsafe.Sizeof(Changeable3{})))
	fmt.Print(string(unsafe.Sizeof(Changeable4{})))
	fmt.Print(string(unsafe.Sizeof(Changeable5{})))
	fmt.Print(string(unsafe.Sizeof(Changeable6{})))
	fmt.Print(string(unsafe.Sizeof(Changeable7{})))
	fmt.Print(string(unsafe.Sizeof(Changeable8{})))
	fmt.Print(string(unsafe.Sizeof(Changeable9{})))
	fmt.Print(string(unsafe.Sizeof(Changeable10{})))
	fmt.Print(string(unsafe.Sizeof(Changeable11{})))
	fmt.Print(string(unsafe.Sizeof(Changeable12{})))
	fmt.Print(string(unsafe.Sizeof(Changeable13{})))
	fmt.Print(string(unsafe.Sizeof(Changeable14{})))
	fmt.Print(string(unsafe.Sizeof(Changeable15{})))
	fmt.Print(string(unsafe.Sizeof(Changeable16{})))
	fmt.Print(string(unsafe.Sizeof(Changeable17{})))
	fmt.Print(string(unsafe.Sizeof(Changeable18{})))
	fmt.Print(string(unsafe.Sizeof(Changeable19{})))
	fmt.Print(string(unsafe.Sizeof(Changeable20{})))
	fmt.Print(string(unsafe.Sizeof(Changeable21{})))
	fmt.Print(string(unsafe.Sizeof(Changeable22{})))
	fmt.Print(string(unsafe.Sizeof(Changeable23{})))
}