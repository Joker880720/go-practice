package tool

import "fmt"

func SliceDelete(a int) []int {

	var slice = []int {1,2,3,4,5,6,7,8,9,10}
	slice = append( slice[:a], slice[a+1:]...)
	fmt.Println(slice)
	return slice

}