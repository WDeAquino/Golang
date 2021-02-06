package main
import "fmt"
func main() {
    var cadena string
    cadena = "Hola bb"
    cadena2 := "Hola bbx2"
    fmt.Println(cadena, " MAS ", cadena2,"\n")
    fmt.Println(len(cadena2))
    fmt.Println(cadena[0:3])
    cadena2 += " \"nuevamente\" "
    fmt.Println(cadena2)
}
