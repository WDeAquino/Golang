package main
import "fmt"

func main()  {
  nombre := []string{
    "Willi",
    "Diana",
    "Romina",
    "Irene",
    "Williberto",
  }

  for index, nombres := range nombre {
    fmt.Printf("El nombre %q esta en el indice %d \n",nombres,index)
  }
  for index := range nombre{
    fmt.Println(index)
  }
}
