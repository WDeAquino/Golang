package main
import "fmt"

func print(nombre string)  {
  fmt.Println(nombre)
}

func main()  {
  cadena := "Willi"
  imprimir := print
  imprimir(cadena)
}
