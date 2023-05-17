package main

import (
	"fmt"
	"log"
	"net/http"
	"go"
	//"github.com/danielrherr/Utils"
)

func main() {
	http.HandleFunc("/", helloHandler) // Manejador para la ruta raíz "/"
	log.Fatal(http.ListenAndServe(":8080", nil)) // Inicia el servidor en el puerto 8080
	//Utils.log("pruebas")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hola, mundo")
}

/*func LoginHandler(w http.ResponseWriter,r *http.Request){
	var x =  {"codigo":1,"descripcion":"hola mundo"}
	return (X)
}*/