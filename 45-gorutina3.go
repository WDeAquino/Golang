package main
import (
  "io"
  "log"
  "net"
  "time"
  "strings"
)

func main()  {
  //Dos variables que reciben los resultados de la funcion net.listen
  listener, err := net.Listen("tcp","localhost:8000")//metodo y conexion
  //Si la funcion de arriba devuelve un error se ejecuta la funcion Fatal
  if err != nil{
    log.Print("Ni se conecto esta cosa")
    log.Fatal(err)
  }
  //Se crea un bucle for infinito que espera alguna conexion
  for {
    conn, err := listener.Accept()
    if err != nil{
      log.Print(err)//por si se va alv el cliente
      log.Print("Ya se fue alv el cliente")
      continue
    }
  go manejarCon(conn)
  }
}

func manejarCon(c net.Conn) {
  defer c.Close()
      io.WriteString(c, "Hola bb :v \n\r")
      done := time.After(20 * time.Second)
      eco := make(chan []byte)
      leerEntrada(eco)
      for {
        select {
        case datos := <-eco:
            log.Print(datos)
        case <-done:
            time.Sleep(1*time.Second)
        }
      }
    }
func leerEntrada(out chan<- []byte) {
	for {
    r := strings.NewReader("some io.Reader stream to be read\n")
		datos := make([]byte, 1024)
		n, _ := io.ReadFull(r, datos)
		if n > 0 {
			out <- datos
		}
	}
}
