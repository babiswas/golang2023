package main
import "fmt"
import "os"
import "log"

func main(){
   path,err:=os.Getwd()
   if err!=nil{
      log.Fatal(err)
      os.Exit(1)
   }
   fmt.Println(path)
   path,err=os.Executable()
   if err!=nil{
      log.Fatal(err)
      os.Exit(1)
   }
   fmt.Println(path)
}