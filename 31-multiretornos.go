package main
import "fmt"

func multiplicar(numero int)(n1,n2,n3 int)  {
  n1=numero*10
  n2=numero*20
  n3=numero*30
  return
}

func main() {
  fmt.Println(multiplicar(25))
  c1,c2,c3:=multiplicar(1);
  fmt.Println(c1,c2,c3)
}
