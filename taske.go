package main
import "fmt"
import "strconv"

func main(){
   msg:=make(chan string)
   go func(){
      for i:=0;i<10;i++{
         fmt.Println("task"+strconv.Itoa(i))
      }
      msg<-"hello"
   }()
   go func(){
      for i:=0;i<10;i++{
         fmt.Println("task"+strconv.Itoa(i))
      }
   }()
   fmt.Println(<-msg)
}