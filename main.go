package main
import ( 
	"fmt"
	tool "myProject/tool"
)

func main(){
	
	var a int
	fmt.Println(tool.Add(5,3))
	fmt.Println(tool.GetStructure())
	fmt.Println(tool.DeferTest())
	fmt.Println("please input index you want to delete:")
	fmt.Scanf("%d",&a)
	fmt.Println(tool.SliceDelete(a))


	var slice = make([] int, 5,10)//長度為5，空間為10
	fmt.Println("value:",slice, "length:",len(slice), "capacity:",cap(slice))
}