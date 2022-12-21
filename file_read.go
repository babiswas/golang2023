package main
import "fmt"
import "os"
import "bufio"
import "log"

func main(){
  file,err:=os.Open("file.txt")
  if err!=nil{
      os.Exit(1)
      return
  }
  defer file.Close()
  scanner:=bufio.NewScanner(file)
  for scanner.Scan(){
     fmt.Println(scanner.Text())
  }
  if err:=scanner.Err();err!=nil{
     log.Fatal(err)
  }
}