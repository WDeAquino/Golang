package main
import "fmt"

type Persona struct{
  nombre string
  edad int
}

type Estudiante struct{
  Persona
  carrera string
}

func main()  {
  Willinton := Estudiante{
    Persona{
      nombre: "Willinton",
      edad: 23,
    },
    "Mecatronica",
  }
  fmt.Println(Willinton)
  fmt.Println(Willinton.nombre)
}
