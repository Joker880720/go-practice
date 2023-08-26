package tool

import "fmt"

func DeferTest() (a int) {
	//var a int 
	fmt.Println("a的地址 %p",&a)
	defer func () {
		fmt.Println("a的地址 %p 值為 %d",&a, a)
		a = 100
		fmt.Println("a的地址 %p 值為 %d",&a, a)
	}()
	return a
}
