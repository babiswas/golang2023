package main
import "fmt"

func main(){
   fmt.Println("Map in golang:")
   m:=make(map[string]int)
   m["test1"]=1
   m["test2"]=2
   m["test3"]=3
   l:=make([]string,1)
   for key,_:=range m{
      l=append(l,key)
   }
   fmt.Println(l)
   m2:=map[string]int{"bapan":1,"tapan":2,"madan":3,"jiban":4}
   fmt.Println(m2)
   fmt.Println("After Deletion:")
   delete(m2,"madan")
   fmt.Println(m2)
   value,ok:=m2["bapan"]
   if ok{
      fmt.Println(value)
   }
   value,ok=m2["test"]
   fmt.Println(ok)
}