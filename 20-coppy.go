package main
import "fmt"

func main()  {
  var origen = []int{0, 1, 2}
  var destino= make([]int, 2)
  copy(destino,origen)
  fmt.Println(origen, destino)
}
