package main
import "fmt"
import "strconv"

type Person struct{
  name string
  id int
}

type Human interface{
    detail()string
}

type Janitor struct{
  name string
  id int
}

func (p *Person)set_name(name string){
   p.name=name
}

func (p *Person)set_id(id int){
   p.id=id
}

func (p Person)detail()string{
   return p.name+" "+strconv.Itoa(p.id)+" "+"Person"
}

func (j *Janitor)set_name(name string){
    j.name=name
}

func (j *Janitor)set_id(id int){
   j.id=id
}

func (j Janitor)detail()string{
  return j.name+" "+strconv.Itoa(j.id)+" "+"Janitor"
}


func main(){
   p:=Person{}
   p.set_name("Bapan")
   p.set_id(12)
   fmt.Println(p)
   h:=Janitor{}
   h.set_name("Tanwar")
   h.set_id(21)
   fmt.Println(h)
   fmt.Println("===============================================")
   func(i interface{}){
      switch v:=i.(type){
         case Janitor:
             fmt.Println("Janitor type executed:")
             fmt.Println(v.detail())
             v.set_name("madan")
             fmt.Println(v.detail())
         case Person:
             fmt.Println("Person type executed:")
             fmt.Println(v.detail())
             v.set_name("madan")
             fmt.Println(v.detail())
         default:
              fmt.Println("Unknown Type:")
      }
   }(h)
   fmt.Println("=============================================")
   func(i interface{}){
      switch v:=i.(type){
         case Janitor:
             fmt.Println("Janitor type executed:")
             fmt.Println(v.detail())
             v.set_name("madan")
             fmt.Println(v.detail())
         case Person:
             fmt.Println("Person type executed:")
             fmt.Println(v.detail())
             v.set_name("madan")
             fmt.Println(v.detail())
         default:
              fmt.Println("Unknown Type:")
      }
   }(p)
}