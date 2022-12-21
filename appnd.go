package main
import "fmt"
import "os"
import "log"

func main(){
   file,err:=os.OpenFile("file.txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0777)
   if err!=nil{
      log.Fatal(err)
      os.Exit(1)
   }
   defer file.Close()
   fmt.Println("Appending to a file:")
   if _,err=file.WriteString("\nLast append");err!=nil{
      log.Fatal(err)
      os.Exit(1)
   }
}