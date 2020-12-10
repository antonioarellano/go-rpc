package main

import (
	"fmt"
	"net/rpc"
)

type Grado struct {
	Estudiante   string
	Curso        string
	Calificacion float32
}

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:5400")
	if err != nil {
		fmt.Println(err)
		return
	}
	n := true
	c.Call("Server.Run", n, &n)
	var op int
	for {
		fmt.Println("1.- Agregar calificacion a materia")
		fmt.Println("2.- Mostrar promedio de estudiante")
		fmt.Println("3-. Mostrar promedio general")
		fmt.Println("4.- Mostrar promedio del curso")
		fmt.Println("0.- Salir")
		fmt.Scanln(&op)

		switch op {
		case 1:
			var nombre, curso, respuesta string
			var calificacion float32

			fmt.Print("Nombre: ")
			fmt.Scanln(&nombre)
			fmt.Print("Curso: ")
			fmt.Scanln(&curso)
			fmt.Print("Calificacion: ")
			fmt.Scanln(&calificacion)

			newGrado := Grado{nombre, curso, calificacion}
			var response string
			err = c.Call("Server.RegistrarGrado", newGrado, &respuesta)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(response)
		case 2:
			var nombre string
			var respuesta float32

			fmt.Print("Ingrese el nombre del estudiante: ")
			fmt.Scanln(&nombre)

			err = c.Call("Server.PromedioEstudiante", nombre, &respuesta)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(nombre+":", respuesta)
			}
		case 3:
			var respuesta float32
			err = c.Call("Server.PromedioGeneral", "", &respuesta)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Promedio general :", respuesta)
			}
		case 4:
			var curso string
			var respuesta float32

			fmt.Print("Curso: ")
			fmt.Scanln(&curso)

			err = c.Call("Server.PromedioCurso", curso, &respuesta)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(curso+":", respuesta)
			}
		case 0:
			return
		}
	}
}

func main() {
	client()
}
