package main
import "fmt"

//Declaramos Dinero basandonos en un entero
type Dinero int

func (d Dinero) String() string {
  return fmt.Sprintf("$%d", d)
}

func main()  {
  var sueldo Dinero
  sueldo = 25000
  fmt.Println(sueldo)
  fmt.Println("El sueldo es ",sueldo)

  aumento := 10000
  sueldo += Dinero(aumento)
  fmt.Println("Sueldo + aumento: ",sueldo)  

}
