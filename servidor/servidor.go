package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

type Server struct {
	Cursos      map[string]map[string]float32
	Estudiantes map[string]map[string]float32
}
type Grado struct {
	Estudiante   string
	Curso        string
	Calificacion float32
}

func (this *Server) Run(a bool, b *bool) error {
	this.Cursos = make(map[string]map[string]float32)
	this.Estudiantes = make(map[string]map[string]float32)
	return nil
}

func Print(estudiantes, cursos map[string]map[string]float32) {
	fmt.Println("Estudiantes:")
	for s := range estudiantes {
		fmt.Println(s + " {")
		for c := range estudiantes[s] {
			fmt.Println("  " + c + ": " + fmt.Sprintf("%.2f", estudiantes[s][c]))
		}
		fmt.Println("}")
	}

	fmt.Println("\nCursos:")
	for c := range cursos {
		fmt.Println(c + " {")
		for s := range cursos[c] {
			fmt.Println("  " + s + ": " + fmt.Sprintf("%.2f", cursos[c][s]))
		}
		fmt.Println("}")
	}
	fmt.Println("-----------------------")
}

func (this *Server) RegistrarGrado(g Grado, response *string) error {
	if _, ok := this.Cursos[g.Curso]; ok {
		if _, ok := this.Cursos[g.Curso][g.Estudiante]; ok {
			*response = "error"
			return errors.New(g.Estudiante + " ya tiene una calificacion para " + g.Curso)
		}
		this.Cursos[g.Curso][g.Estudiante] = g.Calificacion
	} else {
		this.Cursos[g.Curso] = map[string]float32{g.Estudiante: g.Calificacion}
	}

	if _, ok := this.Estudiantes[g.Estudiante]; ok {
		this.Estudiantes[g.Estudiante][g.Curso] = g.Calificacion
	} else {
		this.Estudiantes[g.Estudiante] = map[string]float32{g.Curso: g.Calificacion}
	}

	Print(this.Estudiantes, this.Cursos)

	*response = "completo"
	return nil
}

func (this *Server) PromedioEstudiante(estudiante string, promedio *float32) error {
	if _, ok := this.Estudiantes[estudiante]; !ok {
		return errors.New(estudiante + " no tiene coincidencias válidas")
	}
	var temp float32
	for _, v := range this.Estudiantes[estudiante] {
		temp += v
	}
	*promedio = temp / float32(len(this.Estudiantes[estudiante]))
	return nil
}

func (this *Server) PromedioCurso(curso string, promedio *float32) error {
	if _, ok := this.Cursos[curso]; !ok {
		return errors.New(curso + " does not exist")
	}
	var temp float32
	for _, v := range this.Cursos[curso] {
		temp += v
	}
	*promedio = temp / float32(len(this.Cursos[curso]))
	return nil
}

func (this *Server) PromedioGeneral(_ string, promedio *float32) error {
	if len(this.Estudiantes) == 0 {
		return errors.New("No hay estudiantes registrados en el sistema")
	}
	var sg, gg float32
	for estudiante := range this.Estudiantes {
		this.PromedioEstudiante(estudiante, &sg)
		gg += sg
	}
	*promedio = gg / float32(len(this.Estudiantes))
	return nil
}
func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":5400")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Servidor conectado al puerto 5400")
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}
