package main

import (
	"AlquilerInmuebles/cmd/api/common"
	"AlquilerInmuebles/cmd/api/handlers/calificacion"
	"AlquilerInmuebles/cmd/api/handlers/credenciales"
	"AlquilerInmuebles/cmd/api/handlers/inmueble"
	"AlquilerInmuebles/cmd/api/handlers/reservas"
	"AlquilerInmuebles/cmd/api/handlers/tarjetaCredito"
	"AlquilerInmuebles/cmd/api/handlers/usuario"
	"AlquilerInmuebles/cmd/api/routes"
	"AlquilerInmuebles/internal/repositories/postgresql"
	calificacionPostgre "AlquilerInmuebles/internal/repositories/postgresql/calificacion"
	credencialesPostgre "AlquilerInmuebles/internal/repositories/postgresql/credenciales"
	inmueblePostgre "AlquilerInmuebles/internal/repositories/postgresql/inmueble"
	reservasPostgre "AlquilerInmuebles/internal/repositories/postgresql/reservas"
	tarjetaPostgre "AlquilerInmuebles/internal/repositories/postgresql/tarjetaCredito"
	usuarioPostgre "AlquilerInmuebles/internal/repositories/postgresql/usuario"
	"AlquilerInmuebles/internal/services"
	calificacionService "AlquilerInmuebles/internal/services/calificacion"
	credencialesService "AlquilerInmuebles/internal/services/credenciales"
	inmuebleService "AlquilerInmuebles/internal/services/inmueble"
	reservasService "AlquilerInmuebles/internal/services/reservas"
	tarjetaService "AlquilerInmuebles/internal/services/tarjetaCredito"
	usuarioService "AlquilerInmuebles/internal/services/usuario"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	common.GetConfig()
	services.GetConfigService()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Static("/static/imagenesInmuebles", "./imagenesInmuebles")
	router.Static("/static/imagenInquilinos", "./imagenInquilinos")

	// Configuración CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // origen permitido
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	db, err := postgresql.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	//usuario
	userRepo := &usuarioPostgre.Repository{DataBase: db}
	usuarioSrv := &usuarioService.Service{Repo: userRepo}
	usuarioHandler := &usuario.HandlerUsuario{UsuarioService: usuarioSrv}

	//credenciales
	credenRepo := &credencialesPostgre.Repository{DataBase: db}
	credencialesSrv := &credencialesService.ServiceCredenciales{Repo: credenRepo, UserRepo: userRepo}
	credencialesHandler := &credenciales.CredencialesHanlder{CredencialesService: credencialesSrv}

	//tarjeta
	tarjetaRepo := &tarjetaPostgre.Repository{DataBase: db}
	tarjetaSrv := &tarjetaService.ServiceTarjeta{Repo: tarjetaRepo}
	tarjetaHandler := &tarjetaCredito.HandlerTarjeta{TarjetaService: tarjetaSrv}

	//inmueble
	inmuebleRepo := &inmueblePostgre.Repository{DataBase: db}
	inmuebleSrv := &inmuebleService.ServiceInmueble{Repo: inmuebleRepo}
	inmuebleHandler := &inmueble.InmuebleHandler{InmuebleService: inmuebleSrv}

	//reservas
	reservasRepo := &reservasPostgre.Repository{DataBase: db}
	reservasSrv := &reservasService.ServiceReserva{Repo: reservasRepo, InmRepo: inmuebleRepo}
	reservasHandler := &reservas.ReservaHandler{ReservaService: reservasSrv}
	inmuebleSrv.ResRepo = reservasRepo

	//calificacion
	calificacionRepo := &calificacionPostgre.Repository{DataBase: db}
	calificacionSrv := &calificacionService.ServiceCalificacion{Repo: calificacionRepo}
	calificacionHandler := &calificacion.CalificacionHandler{CalificacionService: calificacionSrv}

	//rutas
	routes.UsuarioRutas(router, usuarioHandler)
	routes.CredencialesRutas(router, credencialesHandler)
	routes.TarjetaRutas(router, tarjetaHandler)
	routes.InmuebleRutas(router, inmuebleHandler)
	routes.CalificacionRutas(router, calificacionHandler)
	routes.ReservasRutas(router, reservasHandler)
	puerto := ":5138"
	fmt.Println("Servidor ejecutándose en el puerto", puerto)
	log.Fatal(router.Run(puerto))
}
