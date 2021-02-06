package main
import "fmt"
func main() {
  header := `
  *******************************
  * Contador de numeros impares *
  *******************************
  `
  var total int
  var arreglo int
  var impares int
  fmt.Println(header)
  fmt.Println("Cuantos numeros quiere ingresar?")
  fmt.Scanln(&total)
  for i:=0;i<total;i++{
    fmt.Println("Escribe el ",i+1," numero:")
    fmt.Scanln(&arreglo)
    if(arreglo%2 != 0){
      impares++
      }
    }
    fmt.Println("Hay ",impares," numeros impares")
}
