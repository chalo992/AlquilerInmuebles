🏠 AlquilerInmuebles
AlquilerInmuebles es una API REST desarrollada en Go (Golang) para la gestión de inmuebles, usuarios y reservas.
Implementa una arquitectura hexagonal (Hexagonal Architecture) que mantiene una separación clara entre el dominio, los servicios, los repositorios y las interfaces externas.


🚀 Tecnologías utilizadas
Lenguaje: Go 1.24
Framework web: Gin Gonic
Autenticación: JWT (JSON Web Tokens)
ORM: GORM
Base de datos: PostgreSQL
Arquitectura: Hexagonal Architecture  


## 📁 Estructura del proyecto
AlquilerInmuebles/
│
├── .env
├── go.mod
├── go.sum
│
├── cmd/
│   ├── main.go
│   └── api/
│       ├── common/
│       │   ├── config.go
│       │   ├── error_mapper.go
│       │   └── obtenerIdusuario.go
│       │
│       ├── handlers/
│       │   ├── calificacion/
│       │   │   ├── calificacion.go
│       │   │   └── handler.go
│       │   ├── credenciales/
│       │   │   ├── credenciales.go
│       │   │   └── handler.go
│       │   ├── inmueble/
│       │   │   ├── inmueble.go
│       │   │   └── handler.go
│       │   ├── reservas/
│       │   │   ├── reservas.go
│       │   │   └── handler.go
│       │   ├── tarjetaCredito/
│       │   │   ├── tarjetaCredito.go
│       │   │   └── handler.go
│       │   └── usuario/
│       │       ├── usuario.go
│       │       └── handler.go
│       │
│       ├── middleware/
│       │   ├── autorizacion.go
│       │   └── rolAutorizacion.go
│       │
│       └── routes/
│           ├── calificacion.go
│           ├── credenciales.go
│           ├── inmueble.go
│           ├── reservas.go
│           ├── tarjetaCredito.go
│           └── usuario.go
│
└── internal/
    ├── domain/
    │   ├── calificacion.go
    │   ├── credenciales.go
    │   ├── email.go
    │   ├── error.go
    │   ├── imagenInmueble.go
    │   ├── inmueble.go
    │   ├── inquilino.go
    │   ├── reserva.go
    │   ├── tarjetaCredito.go
    │   └── usuario.go
    │
    ├── ports/
    │   ├── calificacion.go
    │   ├── credenciales.go
    │   ├── inmueble.go
    │   ├── reservas.go
    │   ├── tarjetaCredito.go
    │   └── usuario.go
    │
    ├── repositories/
    │   └── postgresql/
    │       ├── config.go
    │       ├── connect.go
    │       ├── calificacion/
    │       │   ├── calificacion.go
    │       │   └── repository.go
    │       ├── credenciales/
    │       │   ├── credenciales.go
    │       │   └── repository.go
    │       ├── inmueble/
    │       │   ├── inmueble.go
    │       │   └── repository.go
    │       ├── reservas/
    │       │   ├── reservas.go
    │       │   └── repository.go
    │       ├── tarjetaCredito/
    │       │   ├── tarjetaCredito.go
    │       │   └── repository.go
    │       └── usuario/
    │           ├── usuario.go
    │           └── repository.go
    │
    └── services/
        ├── config.go
        ├── calificacion/
        │   ├── calificacion.go
        │   └── service.go
        ├── common/
        │   ├── password.go
        │   ├── transformarFecha.go
        │   ├── validarTarjeta.go
        │   └── verificarFechaReserva.go
        ├── credenciales/
        │   ├── credenciales.go
        │   └── service.go
        ├── inmueble/
        │   ├── inmueble.go
        │   └── service.go
        ├── reservas/
        │   ├── reservas.go
        │   └── service.go
        ├── tarjetaCredito/
        │   ├── tarjetaCredito.go
        │   └── service.go
        └── usuario/
            ├── usuario.go
            └── service.go

## ⚙️ Configuración y ejecución

1️⃣ Clonar el repositorio
git clone https://github.com/chalo992/AlquilerInmuebles.git
cd AlquilerInmuebles

2️⃣ Instalar dependencias
go mod tidy

3️⃣ Configurar variables de entorno (.env)
DB_HOST=localhost
DB_USER=tu_usuario_postgresql
DB_PASSWORD=tu_contraseña
DB_NAME=tu_db_name
JWT_SECRET=clave_secreta


## 🧩 Middlewares

Autenticación JWT: valida tokens para proteger rutas privadas.
Autorización por rol: controla el acceso según el rol del usuario.


## 🧱 Arquitectura Hexagonal

El proyecto se organiza en tres capas principales:
Domain: contiene las entidades y modelos del negocio.
Ports: define las interfaces de comunicación entre capas.
Services / Repositories / Handlers: implementan la lógica de negocio, acceso a datos y exposición HTTP.
Esto permite independencia entre el dominio y las interfaces externas (como Gin o PostgreSQL).


## 🧩 Manejo de errores
El proyecto implementa un sistema centralizado de manejo de errores mediante la estructura ErrorNegocio, definida en el paquete domain.

type ErrorNegocio struct {
    Mensaje    string
    HTTPStatus int
}

func (e *ErrorNegocio) Error() string {
    return e.Mensaje
}

Cada error de negocio se define como una función que devuelve una instancia de ErrorNegocio, permitiendo mantener mensajes consistentes y códigos de estado HTTP apropiados.
Por ejemplo:
func ErrorEmailRegistrado() *ErrorNegocio {
    return &ErrorNegocio{
        Mensaje:    "El email ya está registrado",
        HTTPStatus: http.StatusConflict,
    }
}

Estos errores se devuelven desde los repositorios o servicios, y son propagados hasta los handlers.
En la capa de handlers, se utiliza un error mapper (definido en cmd/api/common/error_mapper.go) para convertir los errores de dominio en respuestas HTTP adecuadas.
