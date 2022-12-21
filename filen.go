package main
import "fmt"
import "io/ioutil"
import "os"
import "log"
import "path/filepath"

func main(){
  files,err:=ioutil.ReadDir("./")
  if err!=nil{
     log.Fatal(err)
     os.Exit(1)
  }
  fn:=func(path string,info os.FileInfo,err error)error{
     if err!=nil{
        return err
     }
     fmt.Println(path,info.Size())
     return nil
  }
  for index,file:= range files{
      fmt.Println(index," ",file.Name())
  }
  err=filepath.Walk(".",fn)
  if err!=nil{
    log.Fatal(err)
  } 	 	
}
