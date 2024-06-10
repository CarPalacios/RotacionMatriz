# Establecer la imagen base de Go
FROM golang:1.20-alpine3.17

ARG APP_DIR=/matriz

# Crear el directorio de trabajo
WORKDIR ${APP_DIR}

# Copiar los archivos go.mod y go.sum
COPY go.mod go.sum ./

# Descargar todas las dependencias
RUN go mod download

# Copiar el resto del código fuente
COPY . main

# Compilar la aplicación
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Exponer el puerto por el cual la aplicación será accesible
EXPOSE 8080

RUN chmod a+x ./main

# Ejecutar la aplicación
CMD ["./main"]