package main
import "fmt"

func main(){
   mp:=[]string{"hello","bello","nello","illo","tello"}
   l:=make([]string,len(mp))
   copy(l,mp)
   m:=[]string{"bello","nello"}
   m=append(m,mp...)
   fmt.Println(l)
   fmt.Println(m)
}