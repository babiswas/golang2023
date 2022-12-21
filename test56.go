package main
import "fmt"
import "os"
import "bufio"

func main(){
  file,err:=os.Open("prefix.txt")
  if err!=nil{
    os.Exit(1)
  }
  defer file.Close()
  scanner:=bufio.NewScanner(file)
  for scanner.Scan(){
    chr:=scanner.Text()
    fmt.Println(len(chr))
  }
}