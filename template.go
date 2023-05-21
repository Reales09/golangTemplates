package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// Estructuras
type Usuarios struct {
	UserName string
	Edad     int
	Activo   bool
	Admin    bool
	Cursos   []Curso
}

type Curso struct {
	Nombre string
}

// Handler
func Index(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(rw, "Hola mundo")
	template, err := template.ParseFiles("index.html")

	c1 := Curso{"Go"}
	c2 := Curso{"Python"}
	c3 := Curso{"Java"}
	c4 := Curso{"JavaScript"}

	cursos := []Curso{c1, c2, c3, c4}

	usuario := Usuarios{"Reales", 32, true, false, cursos}

	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, usuario)
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	//servidor
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	fmt.Println("El servidor esta coorriendo en el puerto 3000")
	fmt.Printf("Run Server: http://localhost:3000/")
	log.Fatal(server.ListenAndServe())
}
