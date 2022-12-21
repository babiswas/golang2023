package main
import "fmt"
import "strings"
import "sort"

func main(){
  str1:="hello"
  str2:="bello"
  fmt.Println(strings.Compare(str1,str2))
  fmt.Println(strings.Compare(str2,str1))
  fmt.Println(strings.Compare(str1,"hello"))
  test:=[]string{"Bapan","Madan","Apan","Hari","Manu"}
  sort.Strings(test)
  fmt.Println(test)
  test1:="hello,bello,mello,tello"
  s:=strings.Split(test1,",")
  fmt.Println(s)
  str1="Go programming"
  str2="programming"
  if strings.Contains(str1,str2){
     fmt.Println(str2,"Contains in the string:",str1)
  }
  str4:="hello world"
  fmt.Println(strings.ToUpper(str4))
  fmt.Println(strings.ToLower("HELLO WORLD"))
  str5:="hello world"
  for index,value:=range str5{
    fmt.Printf("%d,%c\n",index,value)
  }
  str6:="Cat"
  fmt.Println(strings.Replace(str6,"t","hello",5))
  fmt.Println(str6)
}