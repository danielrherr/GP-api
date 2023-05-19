# Usa una imagen base de Go
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el archivo go.mod y go.sum al contenedor
COPY go.mod .
#COPY go.sum .

# Ejecuta 'go mod download' para descargar las dependencias
RUN go mod download


# Copia todo el contenido del directorio actual al contenedor
COPY . .

# Compila la aplicación
RUN go build -o app .

# Expone el puerto en el que se ejecutará la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación cuando el contenedor se inicie
CMD ["./app"]


