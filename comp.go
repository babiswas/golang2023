package main
import "fmt"

func main(){
  arr:=[5]int{1,2,3,4,5}
  brr:=[5]int{1,2,3,4,5}
  fmt.Println(arr==brr)
  crr:=[5]int{1,2,3,7,8}
  fmt.Println(crr==brr)
  drr:=[]int{100,200,300,400,500}
  krr:=[]int{100,200,300,400,500}
  if len(krr)!=len(drr){
     fmt.Println(false)
  }
  for index,value:= range drr{
      if value!=krr[index]{
         fmt.Println(false)
      }
  }
  fmt.Println(true)
}