package main
import "fmt"
import "sort"

func main(){
  a:=[]string{"hello","aello","bello","mello"}
  fmt.Println(a)
  sort.Strings(a)
  fmt.Println(a)
  b:=[]int{11,21,10,6,17,15,14}
  sort.Ints(b)
  fmt.Println(b)
}