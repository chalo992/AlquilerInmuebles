package postgresql

import (
	"AlquilerInmuebles/internal/domain"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	// Obtiene la configuración de la base de datos desde el paquete config
	configDB := GetConfigDataBase()

	// Construye la cadena DSN para PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configDB.Host,
		configDB.Port,
		configDB.User,
		configDB.Pass,
		configDB.Name,
	)

	// Conecta a la base de datos PostgreSQL
	//var err error
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("¡Conectado a la base de datos!")

	// Migración automática del esquema
	err = DB.AutoMigrate(&domain.Usuario{}, &domain.Tarjeta{}, &domain.Inquilino{}, &domain.Inmueble{}, &domain.ImagenInmueble{}, &domain.Reserva{}, &domain.Calificacion{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Esquemas migrados correctamente")
	return DB, nil
}
