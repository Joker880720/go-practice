package interface_practice

import "fmt"

type Kit interface {
	Add(num1 int, num2 int) (num int, err error)
	Delete(index int) (arr any, err error)
	Revise(index int) (arr any, err error)
	Survey(index int) (arr any, err error)
}

type AddCount struct{
	num int
	num2 int
}

//想改變值就要用指標(結構接收器內部修改欄位值不會生效,因方法調用本身為值傳遞)
func (t AddCount) Add(num1 int, num2 int) int{
	t.num = num1;
	t.num2 = num2;
	return t.num + t.num2
} 

func NumberCount() int {
	a := AddCount{num: 10, num2: 20}
	b := a.Add( 64, 87)
	fmt.Println(b)
	return b
}

