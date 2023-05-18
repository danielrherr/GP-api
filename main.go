package main

import (
	"fmt"
//	"log"
	"net/http"
	"github.com/danielrherr/Utils"
	"github.com/danielrherr/Models/logs"
)

var lo *Utils.Queue
var levelLog int =4
func helloHandler(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprint(w, "Hola, mundo daniel")
	lo.SetLogs(logs.Logs{Log:"Se va a cargar la configuracion.1",Level:levelLog})		
}

func statusHandler (w http.ResponseWriter,r *http.Request){	
	fmt.Fprint(w, "estoy bien")
	lo.SetLogs(logs.Logs{Log:"Se va a cargar la configuracion.",Level:levelLog})		
}

func main() {
	lo = Utils.GetIntancia(levelLog)
	http.HandleFunc("/", helloHandler) // Manejador para la ruta raíz "/"
	http.HandleFunc("/status", statusHandler)
	//log.Fatal(http.ListenAndServe("	:8080", nil)) // Inicia el servidor en el puerto 8080	
	lo.SetLogs(logs.Logs{Log:"Se va a cargar la configuracion.",Level:levelLog})		
	http.ListenAndServe(":8080", nil) // Inicia el servidor en el puerto 8080	
	
}

