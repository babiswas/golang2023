package main
import "fmt"
import "sync"
import "os"
import "bufio"

func main(){
   ch:=make(chan string,100)
   var wg sync.WaitGroup
   count:=0
   file,err:=os.Open("prefix.txt")
   if err!=nil{
      os.Exit(1)
   }
   scanner:=bufio.NewScanner(file)
   for scanner.Scan(){
       count+=1
       wg.Add(1)
       go func(wg *sync.WaitGroup,str1 string,ch chan<- string){
          defer wg.Done()
          ch<-str1
       }(&wg,scanner.Text(),ch)
   }
   fmt.Println(count)
   for{
        if count==0{
             select{
                  case msg:=<-ch:
                       fmt.Println(msg)
                       count-=1
                       if count!=0{
                          for value:= range ch{
                              fmt.Println(value)
                              count-=1
                              if count==0{
                                 break
                              }
                          }
                       }
             }
        }else{
           break
        }
   }
   wg.Wait()
}