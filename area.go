package main
import "fmt"

type Func func(int,int)int

type Rectangle struct{
  length int
  breadth int
  perimeter Func
  area Func
}

func main(){
    rect:=Rectangle{3,4,func(length int,breadth int)int{ return length*breadth },func(length int,breadth int)int{return 2*(length+breadth)}}
    fmt.Println("The perimeter of the rectangle is:",rect.perimeter(rect.length,rect.breadth))
    fmt.Println("The area of the rectangle is:",rect.area(rect.length,rect.breadth))
}