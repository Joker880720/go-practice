package main
import ( 
	"fmt"
	//tool "myProject/tool"
	interfacePractice "myProject/interface_practice"
)

func main(){
	/*
	//Function 模組化練習
	var a int
	fmt.Println(tool.Add(5,3))
	fmt.Println(tool.GetStructure())
	fmt.Println(tool.DeferTest())
	fmt.Println("please input index you want to delete:")
	fmt.Scanf("%d",&a)
	fmt.Println(tool.SliceDelete(a))
	shape := tool.Shape{ShapeName: "rectangle"}
	fmt.Println(shape.ShapeDisplay())*/

	//方法接收器練習
	//var kit interfacePractice.Kit
	fmt.Println(interfacePractice.NumberCount())

	//
	slice := make([] int, 5,10)//長度為5，空間為10
	fmt.Println("value:",slice, "length:",len(slice), "capacity:",cap(slice))
}