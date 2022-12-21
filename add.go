package main
import "fmt"

func add(a,b,c int)int{
  return a+b+c
}

func myclosure()func()int{
  num:=1
  return func()int{
    num+=1
    return num
  }
}

func main(){
  nums:=add(2,3,4)
  fmt.Println("The sum of the numbers is:",nums)
  f1:=myclosure()
  fmt.Println("Executing Closure:")
  fmt.Println(f1())
  fmt.Println(f1())
}