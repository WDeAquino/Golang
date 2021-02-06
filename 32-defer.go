package main
import "fmt"
//Nos da acceso a datos del sistema operativo
import "os"

func primera()  {
  fmt.Println("primera")
}
func segunda()  {
  fmt.Println("segunda")
}
func main()  {
  defer primera()
  defer segunda()
//Abrimos el archivo, si esta en otra carpeta especificamos la ruta
f,err := os.Open("texto.txt")
//Verificamos que no haya errores
if err != nil{
  panic(err)
}
defer f.Close()
  //Creamos un slice para almacenar lo que leemos en el archivo
  x := make([]byte, 175)
  c, err := f.Read(x)
  //Verificamos que no haya errores
  if err != nil{
    panic(err)
  }
fmt.Printf("Cantidad de bits leidos: %d \nTexto leido: %q \nError: %v",c,x,err)
}
