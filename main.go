package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/danielrherr/Utils"
)

func main() {
	http.HandleFunc("/", helloHandler) // Manejador para la ruta raíz "/"
	log.Fatal(http.ListenAndServe(":8080", nil)) // Inicia el servidor en el puerto 8080
	Utils.log("pruebas")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hola, mundo")
}
