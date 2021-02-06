package main
import "fmt"
func sumar(numeros...int) int{
    resultado :=0
    for _,numero:=range numeros{
      resultado += numero
    }
    return resultado;
}
func main() {
  fmt.Println(sumar(1,2,3,1000))
}
