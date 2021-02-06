package main
import "fmt"

func imprimir()  {
  fmt.Println("Hola enfermeraa :v")
/*  panic("Error")
  cadena := recover()
  fmt.Println(cadena)*/
  defer func ()  {
    cadena := recover()
    fmt.Println(cadena)
  }()
  panic("ERROR")
}

func main()  {
  imprimir()
  fmt.Println("MEEEEEEEIN")
}
