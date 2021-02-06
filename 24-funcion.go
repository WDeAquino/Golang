package main
import "fmt"

func aver(nombre...string){
  for _, x :=range nombre{
  fmt.Println(x)
  }
}

func main(){
  a := "Willi"
  b := "Diana"
  c := "Romina"
  aver(a,b,c)
}
