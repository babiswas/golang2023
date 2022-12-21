package main
import "fmt"
import "errors"

func divide(num1,num2 int)error{
   if num2==0{
      return errors.New("Divison by zero not allowed:")
   }
   fmt.Println("Result:",num1/num2)
   return nil
}

func main(){
  divide(4,2)
  divide(2,4)
  divide(4,0)
}