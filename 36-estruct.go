package main
import "fmt"

type Persona struct{
  Nombre string
  Edad int
}
func older(p1,p2 Persona) (Persona, int) {
  if p1.Edad>p2.Edad {
      return p1,p1.Edad-p2.Edad
    }
  return p2,p2.Edad-p1.Edad
}

func main()  {
  var a Persona
  a.Nombre = "Willinton"
  a.Edad = 28
  fmt.Println("Estructura de a de tipo Persona: ", a)
  b := Persona{Nombre: "Diana", Edad: 25}
  fmt.Println("Estructura de b de tipo Persona: ", b)
  fmt.Println("Quien es mas grande? A o B")
  fmt.Println(older(a,b))

}
