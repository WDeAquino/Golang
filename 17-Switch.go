package main
import "fmt"
func main() {
  encabezado:=`
    ***************************************************
    ************* Dia de la semana switch**************
    ***************************************************
  `
  fmt.Println(encabezado)
  fmt.Println("Que dia es?")
  var cont int
  fmt.Scanln(&cont)
  switch cont{
  case 1:
    fmt.Println("Lunes")
  case 2:
    fmt.Println("Martes")
  case 3:
    fmt.Println("Miercoles")
  default:
    fmt.Println("No chinges y dime un puto dia")
    break
  }

}
