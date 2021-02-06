package main
import "fmt"
import "math"


type rectangulo struct{
  base,altura float64
}
type circulo struct{
  radio float64
}

func (r rectangulo) area() float64 {
    return r.base*r.altura
}
func (c circulo) area() float64 {
    return c.radio * c.radio * math.Pi
}

func main()  {
  r := rectangulo{15,1}
  c := circulo{1}
  fmt.Println("El area del rectangulo es: ", r.area())
  fmt.Println("El area del cirulo es: ", c.area())
}
