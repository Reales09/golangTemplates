package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// // Estructuras
type Usuarios struct {
	UserName string
	Edad     int
	// Activo   bool
	// Admin    bool
	// Cursos   []Curso
}

// type Curso struct {
// 	Nombre string
// }

// Funciones

func Saludar(nombre string) string {
	return "Hola " + nombre + " desde una función"
}

var templates = template.Must(template.New("T").ParseGlob("templates/**/*.html"))
var errorTemplate = template.Must(template.ParseFiles("templates/error/error.html"))

//Handler error

func handlerError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	errorTemplate.Execute(rw, nil)
}

// Funcion de render template
func renderTemplate(rw http.ResponseWriter, name string, data interface{}) {
	err := templates.ExecuteTemplate(rw, name, data)

	if err != nil {
		// http.Error(rw, "No es posible retornar el template", http.StatusInternalServerError)
		handlerError(rw, http.StatusInternalServerError)

	}
}

// Handler
func Index(rw http.ResponseWriter, r *http.Request) {

	// funciones := template.FuncMap{
	// 	"saludar": Saludar,
	// }

	// fmt.Fprintf(rw, "Hola mundo")

	// template, err := template.New("index.html").Funcs(funciones).ParseFiles("index.html")

	// template := template.Must(template.New("index.html").ParseFiles("index.html", "base.html"))

	usuario := Usuarios{"Reales", 32}

	// template.Execute(rw, usuario)

	renderTemplate(rw, "index.html", usuario)

	// c1 := Curso{"Go"}
	// c2 := Curso{"Python"}
	// c3 := Curso{"Java"}
	// c4 := Curso{"JavaScript"}

	// cursos := []Curso{c1, c2, c3, c4}

	// usuario := Usuarios{"Reales", 32, true, false, cursos}
}

func Registro(rw http.ResponseWriter, r *http.Request) {

	renderTemplate(rw, "registr.html", nil)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/registro", Registro)

	//servidor
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	fmt.Println("El servidor esta coorriendo en el puerto 3000")
	fmt.Printf("Run Server: http://localhost:3000/")
	log.Fatal(server.ListenAndServe())
}
