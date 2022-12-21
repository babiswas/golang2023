package main
import "fmt"
import "sync"
import "strconv"

func main(){
   var wg sync.WaitGroup
   tasklist:=[]string{"task1"}
   ch1:=make(chan string,4)
   for i:=0;i<20;i++{
      task_number:=strconv.Itoa(i)
      tasklist=append(tasklist,"task"+task_number)
   }
   snd:=func(wg *sync.WaitGroup,ch1 chan<- string){
       defer wg.Done()
       for _,value:=range tasklist{
          ch1<-value
       }
       close(ch1)
   }
   rcv:=func(wg *sync.WaitGroup,ch1 <-chan string){
        defer wg.Done()
        for {
           val,ok:=<-ch1
           if ok==true{
              fmt.Println(val)
           }else{
             fmt.Println("Channel is closed:")
             break
           }
           for val:=range ch1{
              fmt.Println(val)
           }
        }
   }
   wg.Add(1)
   go snd(&wg,ch1)
   wg.Add(1)
   go rcv(&wg,ch1)
   wg.Wait()
}