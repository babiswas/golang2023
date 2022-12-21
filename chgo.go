package main
import "fmt"

func main(){
  ch:=make(chan string,4)
  ch<-"hello"
  ch<-"bello"
  fmt.Println(<-ch)
}