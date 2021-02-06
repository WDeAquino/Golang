package main
import "fmt"

func main()  {
  y :=  make(map[string]string, 7)
  fmt.Println(y)

  x :=  make(map[string]string)
  fmt.Println(x)

  x["Nombre"] = "Willi"
  x["Edad"] = "29"
  fmt.Println(x)
  fmt.Println(x["Nombre"])
  fmt.Println("a ver si jala esta madre")
  fmt.Println(x["Edad"])

  edades := map[string]int{
    "Willi":  23,
    "Diana":  25,
    "Romina": 5,
  }
  fmt.Println(edades["Willi"])

}
