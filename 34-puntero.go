package main
import "fmt"

func main()  {
  a:=25
  fmt.Println("El valor de a es: ",a)
  fmt.Println("La direccion de a es: ",&a)
  b:=&a
  fmt.Println("B vale: ", b)
  fmt.Println("La direccion de b es:", &b)
  fmt.Println("B esta apuntando hacia: ",*b)
  *b=32
  fmt.Println("El valor de a es: ",a)
  c:=&a
  if (c==b) {
    fmt.Println("c es igual a b"  )
  }
  d:=new(int)
  fmt.Println("La direccion de d es: ",d)
  fmt.Println("El valor de d es: ", *d)
  d=b
  fmt.Println("La direccion de d es: ",d)
  fmt.Println("El valor de d es: ", *d)

}
